//
// Copyright 2023 The GUAC Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package gremlin

import (
	"encoding/json"
	"errors"
	"fmt"
	gremlingo "github.com/apache/tinkerpop/gremlin-go/v3/driver"
	"golang.org/x/sync/errgroup"
	"math"
	"strconv"
)

type Label string

const MaxVertexUpsertChunkSize = 200
const MaxEdgeUpsertChunkSize = 10 // breaks with anything over 10

func (c *gremlinClient) upsertVertex(q *GraphQuery) (string, error) {
	g := gremlingo.Traversal_().WithRemote(c.remote)
	r, err := g.MergeV(q.toVertexMap()).Id().Next()
	if err != nil {
		return "", err
	}
	if c.config.Flavor == JanusGraph {
		id, err := r.GetInt64()
		if err != nil {
			return "", err
		}
		return strconv.FormatInt(id, 10), nil
	} else {
		return r.GetString(), nil
	}
}

func (c *gremlinClient) bulkUpsertVertices(qs []*GraphQuery) ([]string, error) {
	var ids []string
	var vertexRefs []interface{}
	var gt *gremlingo.GraphTraversal
	g := gremlingo.Traversal_().WithRemote(c.remote)
	// chain the upserts
	for i, q := range qs {
		vertexRef := fmt.Sprintf("v%d", i)
		vertexRefs = append(vertexRefs, vertexRef)
		if i == 0 {
			gt = g.MergeV(q.toVertexMap()).As(vertexRef)
		} else {
			gt = gt.MergeV(q.toVertexMap()).As(vertexRef)
		}
	}
	results, err := gt.Select(vertexRefs...).Select(gremlingo.Column.Values).Unfold().Id().ToList()
	if err != nil {
		return nil, err
	}

	for _, result := range results {
		if c.config.Flavor == JanusGraph {
			id, err := result.GetInt64()
			if err != nil {
				return nil, err
			}
			ids = append(ids, strconv.FormatInt(id, 10))
		} else {
			ids = append(ids, result.GetString())
		}
	}

	// verify the results are somewhat sane
	if len(ids) != len(qs) {
		return nil, fmt.Errorf("bulkUpsertVertices: number of results(%d) gathered does not match number of inputs(%d)",
			len(ids), len(qs))
	}

	return ids, nil
}

type MapSerializer[M any] func(model M) (result *GraphQuery)

type ObjectDeserializer[M any] func(id string, values map[interface{}]interface{}) (model M)

type EdgeMapSerializer[VInput any, EInput any] func(v1 VInput, v2 VInput, edge EInput) (result *GraphQuery)

type EdgeObjectDeserializer[M any] func(id string, out map[interface{}]interface{}, edge map[interface{}]interface{}, in map[interface{}]interface{}) (model M)

/*
C is typically an InputSpec
D is model object w/ id after upsert
*/
func ingestModelObject[C any, D any](c *gremlinClient, modelObject C, serializer MapSerializer[C], deserializer func(*gremlinQueryResult) (D, error)) (D, error) {
	var object D
	values := serializer(modelObject)

	id, err := c.upsertVertex(values)
	if err != nil {
		return object, err
	}
	result := &gremlinQueryResult{
		vertexId: id,
		vertex:   values.toReadMap(),
	}

	object, err = deserializer(result)
	if err != nil {
		return object, err
	}

	return object, nil
}

func bulkIngestModelObjects[C any, D any](c *gremlinClient, modelObjects []C, serializer MapSerializer[C], deserializer func(*gremlinQueryResult) (D, error)) ([]D, error) {
	var objects []D
	if len(modelObjects) < 1 {
		// nothing to do
		return objects, nil
	}

	// serialize
	var qs []*GraphQuery
	for _, modelObject := range modelObjects {
		values := serializer(modelObject)
		qs = append(qs, values)
	}

	// split in chunk to limit websocket frame size
	// FIXME: make this configurable
	// FIXME: given these span multiple requests, we should wrap them in a single transaction
	// FIXME: can we do these in parallel? (probably not if they are in a single transaction)
	var ids = make([]string, 0)
	for _, chunk := range chunkSlice(qs, MaxVertexUpsertChunkSize) {
		if len(chunk) == 1 {
			// if there's only 1, the query handling is different, so do a normal upsert
			idForChunk, err := c.upsertVertex(chunk[0])
			if err != nil {
				return objects, err
			}
			ids = append(ids, idForChunk)
		} else {
			idsForChunk, err := c.bulkUpsertVertices(chunk)
			if err != nil {
				return objects, err
			}
			ids = append(ids, idsForChunk...)
		}
	}

	if len(ids) != len(modelObjects) {
		return nil, fmt.Errorf("the lengths dont match, I am sad. Num ids is %d, vs expected is %d", len(ids), len(modelObjects))
	}

	for k := range modelObjects {
		result := &gremlinQueryResult{
			vertexId:    ids[k],
			vertexLabel: qs[k].label,
			vertex:      qs[k].toReadMap(),
		}

		object, err := deserializer(result)
		if err != nil {
			return objects, err
		}
		objects = append(objects, object)
	}
	return objects, nil
}

func chunkSlice[T any](slice []T, chunkSize int) [][]T {
	var chunks [][]T
	for i := 0; i < len(slice); i += chunkSize {
		end := i + chunkSize
		if end > len(slice) {
			end = len(slice)
		}
		chunks = append(chunks, slice[i:end])
	}
	return chunks
}

func storeArrayInVertexProperties(q *GraphQuery, propertyName string, mapToStore []string) {
	mapToStoreJson, _ := json.Marshal(mapToStore)
	q.has[propertyName] = string(mapToStoreJson)
}

func ingestModelObjectsWithRelation[M any](c *gremlinClient, gqb *gremlinQueryBuilder[M]) (M, error) {
	var object M

	if gqb.query.inVQuery != nil || gqb.query.outVQuery != nil {
		result, err := upsertRelationDirect(c, gqb)
		if err != nil {
			return object, err
		}
		return gqb.mapper(result)
	} else {
		result, err := upsertVertexDirect(c, gqb)
		if err != nil {
			return object, err
		}
		return gqb.mapper(result)
	}
}

func upsertVertexDirect[M any](c *gremlinClient, q *gremlinQueryBuilder[M]) (*gremlinQueryResult, error) {
	id, err := c.upsertVertex(q.query)
	if err != nil {
		return nil, err
	}

	return &gremlinQueryResult{
		vertexId:    id,
		vertexLabel: q.query.label,
		vertex:      q.query.toReadMap(),
	}, nil
}

func bulkIngestModelObjectsWithRelation[M any](c *gremlinClient, gqb *gremlinQueryBuilder[M]) ([]M, error) {
	var objects []M

	// split into chunks and run upserts in parallel
	var g errgroup.Group
	numChunksUpper := int(math.Ceil(float64(len(gqb.queries))/MaxEdgeUpsertChunkSize)) + 1
	resultChan := make(chan []*gremlinQueryResult, numChunksUpper)
	for _, chunk := range chunkSlice(gqb.queries, MaxEdgeUpsertChunkSize) {
		if len(chunk) == 1 {
			// if there's only 1, the result handling is different, so do a normal upsert
			localRelationRef := chunk[0]
			g.Go(func() error {
				err := upsertRelation[M](c, resultChan, localRelationRef)
				return err
			})
		} else {
			localChunkRef := chunk
			g.Go(func() error {
				err := bulkUpsertRelations[M](c, resultChan, localChunkRef)
				return err
			})
		}
		// FIXME: Only one at time - concurrency handling is limited in JanusGraph
		if err := g.Wait(); err != nil {
			return objects, err
		}
	}

	// all upserts are done, we can close the channel
	close(resultChan)
	// map back out to the target object
	for results := range resultChan {
		for _, result := range results {
			object, err := gqb.queries[0].mapper(result)
			if err != nil {
				return objects, err
			}
			objects = append(objects, object)
		}
	}
	// verify the results are somewhat sane
	if len(objects) != len(gqb.queries) {
		return nil, fmt.Errorf("bulkIngestModelObjectsWithRelation: number of objects(%d) gathered does not match number of inputs(%d)",
			len(objects), len(gqb.queries))
	}

	return objects, nil
}

func upsertRelation[M any](c *gremlinClient, queue chan []*gremlinQueryResult, q *gremlinQueryBuilder[M]) error {
	var t *gremlingo.GraphTraversal
	g := gremlingo.Traversal_().WithRemote(c.remote)

	if !q.query.outVQuery.isUpsert {
		// match from
		t = g.V().HasLabel(string(q.query.outVQuery.label))
		for k, v := range q.query.outVQuery.has {
			t = t.Has(k, v)
		}
		t = t.Limit(1).As("from")
	} else {
		// upsert from
		t = g.MergeV(q.query.outVQuery.toVertexMap()).As("from")
	}

	if !q.query.inVQuery.isUpsert {
		// match to
		t = t.V().HasLabel(string(q.query.inVQuery.label))
		for k, v := range q.query.inVQuery.has {
			t = t.Has(k, v)
		}
		t = t.Limit(1).As("to")
	} else {
		// upsert to
		t = t.MergeV(q.query.inVQuery.toVertexMap()).As("to")
	}

	// upsert edge
	t = t.MergeE(q.query.toEdgeMap()).
		Option(gremlingo.Merge.OutV, gremlingo.T__.Select("from")).
		Option(gremlingo.Merge.InV, gremlingo.T__.Select("to")).
		As("edge")
	// project for output
	t = t.Project("from", "edge", "to").
		By(gremlingo.T__.OutV().ValueMap(true)).
		By(gremlingo.T__.Select("edge").Id()).
		By(gremlingo.T__.InV().ValueMap(true))
	r, err := t.Next()
	if err != nil {
		return err
	}

	resultMap := r.GetInterface().(map[interface{}]interface{})
	edgeValue := resultMap["edge"]
	fromMap := resultMap["from"].(map[interface{}]interface{})
	toMap := resultMap["to"].(map[interface{}]interface{})

	var edgeId string
	var outId string
	var inId string
	if c.config.Flavor == JanusGraph {
		edgeId = strconv.FormatInt(edgeValue.(*janusgraphRelationIdentifier).RelationId, 10)
		outId = strconv.FormatInt(fromMap[string(gremlingo.T.Id)].(int64), 10)
		inId = strconv.FormatInt(toMap[string(gremlingo.T.Id)].(int64), 10)
	} else {
		edgeMap := edgeValue.(map[interface{}]interface{})
		edgeId = edgeMap["id"].(string)

		outId = fromMap[string(gremlingo.T.Id)].(string)
		inId = toMap[string(gremlingo.T.Id)].(string)
	}

	result := &gremlinQueryResult{
		id:        edgeId,
		edge:      q.query.toReadMap(),
		edgeLabel: q.query.label,
		out:       flattenResultMap(fromMap),
		outLabel:  q.query.outVQuery.label,
		outId:     outId,
		in:        flattenResultMap(toMap),
		inLabel:   q.query.inVQuery.label,
		inId:      inId,
		query:     q.query,
	}

	// ship it
	queue <- []*gremlinQueryResult{result}
	return nil
}

func upsertRelationDirect[M any](c *gremlinClient, q *gremlinQueryBuilder[M]) (*gremlinQueryResult, error) {
	ch := make(chan []*gremlinQueryResult, 1)
	err := upsertRelation(c, ch, q)
	if err != nil {
		return nil, err
	}
	close(ch)
	// return the first element from the first list in the channel
	for _, result := range <-ch {
		return result, nil
	}
	return nil, errors.New("no results returned by upsert")
}

func bulkUpsertRelations[M any](c *gremlinClient, queue chan []*gremlinQueryResult, queries []*gremlinQueryBuilder[M]) error {
	var edgeRefs []interface{}
	var t *gremlingo.GraphTraversal
	g := gremlingo.Traversal_().WithRemote(c.remote)
	// chain the upserts
	for k, q := range queries {
		v1Ref := fmt.Sprintf("from-%d", k)
		v2Ref := fmt.Sprintf("to-%d", k)
		edgeRef := fmt.Sprintf("e-%d", k)
		edgeRefs = append(edgeRefs, edgeRef)

		// FIXME: cleanup this k==0 business and dedup
		// match from
		if k == 0 {
			if !q.query.outVQuery.isUpsert {
				// match from
				t = g.V().HasLabel(string(q.query.outVQuery.label))
				for k, v := range q.query.outVQuery.has {
					t = t.Has(k, v)
				}
				t = t.Limit(1).As(v1Ref)
			} else {
				// upsert from
				t = g.MergeV(q.query.outVQuery.toVertexMap()).As(v1Ref)
			}
		} else {
			if !q.query.outVQuery.isUpsert {
				// match from
				t = t.V().HasLabel(string(q.query.outVQuery.label))
				for k, v := range q.query.outVQuery.has {
					t = t.Has(k, v)
				}
				t = t.Limit(1).As(v1Ref)
			} else {
				// upsert from
				t = t.MergeV(q.query.outVQuery.toVertexMap()).As(v1Ref)
			}
		}
		// match to
		t = t.V().HasLabel(string(q.query.inVQuery.label))
		for k, v := range q.query.inVQuery.has {
			t = t.Has(k, v)
		}
		t = t.Limit(1).As(v2Ref)
		// upsert edge
		t = t.MergeE(q.query.toEdgeMap()).
			Option(gremlingo.Merge.OutV, gremlingo.T__.Select(v1Ref)).
			Option(gremlingo.Merge.InV, gremlingo.T__.Select(v2Ref)).
			As(edgeRef)
	}

	results, err := t.Select(edgeRefs...).Select(gremlingo.Column.Values).Unfold().Project("from", "edge", "to").
		By(gremlingo.T__.OutV().ValueMap(true)).
		By(gremlingo.T__.Id()).
		By(gremlingo.T__.InV().ValueMap(true)).ToList()
	if err != nil {
		return err
	}

	// verify the results are somewhat sane
	if len(results) != len(queries) {
		return fmt.Errorf("bulkUpsertRelations: number of results(%d) gathered does not match number of inputs(%d)",
			len(results), len(queries))
	}

	var queryResults []*gremlinQueryResult
	for k, r := range results {
		resultMap := r.GetInterface().(map[interface{}]interface{})
		edgeValue := resultMap["edge"]
		fromMap := resultMap["from"].(map[interface{}]interface{})
		toMap := resultMap["to"].(map[interface{}]interface{})

		var edgeId string
		var outId string
		var inId string
		if c.config.Flavor == JanusGraph {
			edgeId = strconv.FormatInt(edgeValue.(*janusgraphRelationIdentifier).RelationId, 10)
			outId = strconv.FormatInt(fromMap[string(gremlingo.T.Id)].(int64), 10)
			inId = strconv.FormatInt(toMap[string(gremlingo.T.Id)].(int64), 10)
		} else {
			edgeMap := edgeValue.(map[interface{}]interface{})
			edgeId = edgeMap["id"].(string)

			outId = fromMap[string(gremlingo.T.Id)].(string)
			inId = toMap[string(gremlingo.T.Id)].(string)
		}

		result := &gremlinQueryResult{
			id:        edgeId,
			edge:      queries[k].query.toReadMap(),
			edgeLabel: queries[k].query.label,
			out:       flattenResultMap(fromMap),
			outLabel:  queries[k].query.outVQuery.label,
			outId:     outId,
			in:        flattenResultMap(toMap),
			inLabel:   queries[k].query.inVQuery.label,
			inId:      inId,
			query:     queries[k].query,
		}

		queryResults = append(queryResults, result)
	}

	queue <- queryResults
	return nil
}
