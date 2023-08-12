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

type Relation struct {
	outV *GraphQuery
	inV  *GraphQuery
	edge *GraphQuery

	upsertOutV bool
}

type RelationWithId struct {
	edgeId   string
	relation *Relation

	outV map[interface{}]interface{}
	inV  map[interface{}]interface{}
}

/*
C is typically an InputSpec
D is model object w/ id after upsert
*/
func ingestModelObject[C any, D any](c *gremlinClient, modelObject C, serializer MapSerializer[C], deserializer ObjectDeserializer[D]) (D, error) {
	var object D
	values := serializer(modelObject)

	id, err := c.upsertVertex(values)
	if err != nil {
		return object, err
	}
	object = deserializer(id, values.toReadMap())
	return object, nil
}

func bulkIngestModelObjects[C any, D any](c *gremlinClient, modelObjects []C, serializer MapSerializer[C], deserializer ObjectDeserializer[D]) ([]D, error) {
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
		object := deserializer(ids[k], qs[k].toReadMap())
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

func ingestModelObjectsWithRelation[VInput any, EInput any, EOutput any](c *gremlinClient,
	v1InputObject VInput,
	v2InputObject VInput,
	edgeInputObject EInput,
	v1InputSerializer MapSerializer[VInput],
	v2InputSerializer MapSerializer[VInput],
	edgeInputSerializer EdgeMapSerializer[VInput, EInput],
	edgeOutputDeserializer EdgeObjectDeserializer[EOutput]) (EOutput, error) {

	var edgeObject EOutput

	relation := &Relation{
		outV: v1InputSerializer(v1InputObject),
		inV:  v2InputSerializer(v2InputObject),
		edge: edgeInputSerializer(v1InputObject, v2InputObject, edgeInputObject),
	}

	ch := make(chan []*RelationWithId, 1)
	err := c.upsertRelation(ch, relation)
	if err != nil {
		return edgeObject, err
	}
	close(ch)
	// return the first element from the first list in the channel
	for _, relationWithId := range <-ch {
		return edgeOutputDeserializer(relationWithId.edgeId,
			flattenResultMap(relationWithId.outV),
			relationWithId.relation.edge.toReadMap(),
			flattenResultMap(relationWithId.inV)), nil
	}
	return edgeObject, errors.New("no results returned by upsert")
}

func bulkIngestModelObjectsWithRelation[VInput any, EInput any, EOutput any](c *gremlinClient,
	v1InputObjects []VInput,
	v2InputObjects []VInput,
	edgeInputObjects []EInput,
	v1InputSerializer MapSerializer[VInput],
	v2InputSerializer MapSerializer[VInput],
	edgeInputSerializer EdgeMapSerializer[VInput, EInput],
	edgeOutputDeserializer EdgeObjectDeserializer[EOutput]) ([]EOutput, error) {

	var objects []EOutput
	// basic input validation
	if len(v1InputObjects) < 1 {
		// nothing to do
		return objects, nil
	}
	if len(v1InputObjects) != len(v2InputObjects) || len(v1InputObjects) != len(edgeInputObjects) {
		return objects, fmt.Errorf("size of inputs do not match outV(%d) inV(%d) e(%d)",
			len(v1InputObjects), len(v2InputObjects), len(edgeInputObjects))
	}

	// serialize all the values
	var allRelations []*Relation
	for k, v1InputObject := range v1InputObjects {
		v2InputObject := v2InputObjects[k]
		edgeInputObject := edgeInputObjects[k]
		relation := &Relation{
			outV: v1InputSerializer(v1InputObject),
			inV:  v2InputSerializer(v2InputObject),
			edge: edgeInputSerializer(v1InputObject, v2InputObject, edgeInputObject),
		}
		allRelations = append(allRelations, relation)
	}

	// split into chunks and run upserts in parallel
	var g errgroup.Group
	numChunksUpper := int(math.Ceil(float64(len(allRelations))/MaxEdgeUpsertChunkSize)) + 1
	resultChan := make(chan []*RelationWithId, numChunksUpper)
	for _, chunk := range chunkSlice(allRelations, MaxEdgeUpsertChunkSize) {
		if len(chunk) == 1 {
			// if there's only 1, the result handling is different, so do a normal upsert
			localRelationRef := chunk[0]
			g.Go(func() error {
				err := c.upsertRelation(resultChan, localRelationRef)
				return err
			})
		} else {
			localChunkRef := chunk
			g.Go(func() error {
				err := c.bulkUpsertRelations(resultChan, localChunkRef)
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
	for result := range resultChan {
		for _, relationWithId := range result {
			object := edgeOutputDeserializer(relationWithId.edgeId,
				flattenResultMap(relationWithId.outV),
				relationWithId.relation.edge.toReadMap(),
				flattenResultMap(relationWithId.inV))
			objects = append(objects, object)
		}
	}
	// verify the results are somewhat sane
	if len(objects) != len(allRelations) {
		return nil, fmt.Errorf("bulkIngestModelObjectsWithRelation: number of objects(%d) gathered does not match number of inputs(%d)",
			len(objects), len(allRelations))
	}

	return objects, nil
}

func (c *gremlinClient) upsertRelation(queue chan []*RelationWithId, relation *Relation) error {
	var t *gremlingo.GraphTraversal
	g := gremlingo.Traversal_().WithRemote(c.remote)
	if !relation.upsertOutV {
		// match from
		t = g.V().HasLabel(string(relation.outV.label))
		for k, v := range relation.outV.has {
			t = t.Has(k, v)
		}
		t = t.Limit(1).As("from")
	} else {
		// upsert from
		t = g.MergeV(relation.outV.toVertexMap()).As("from")
	}

	// match to
	t = t.V().HasLabel(string(relation.inV.label))
	for k, v := range relation.inV.has {
		t = t.Has(k, v)
	}
	t = t.Limit(1).As("to")
	// upsert edge
	t = t.MergeE(relation.edge.toEdgeMap()).
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
	if c.config.Flavor == JanusGraph {
		edgeId = strconv.FormatInt(edgeValue.(*janusgraphRelationIdentifier).RelationId, 10)
	} else {
		edgeMap := edgeValue.(map[interface{}]interface{})
		edgeId = edgeMap["id"].(string)
	}

	relationWithId := &RelationWithId{
		edgeId:   edgeId,
		relation: relation,
		outV:     fromMap,
		inV:      toMap,
	}

	// ship it
	queue <- []*RelationWithId{relationWithId}
	return nil
}

func (c *gremlinClient) upsertRelationDirect(relation *Relation) (*RelationWithId, error) {
	ch := make(chan []*RelationWithId, 1)
	err := c.upsertRelation(ch, relation)
	if err != nil {
		return nil, err
	}
	close(ch)
	// return the first element from the first list in the channel
	for _, relationWithId := range <-ch {
		return relationWithId, nil
	}
	return nil, errors.New("no results returned by upsert")
}

func (c *gremlinClient) bulkUpsertRelations(queue chan []*RelationWithId, relations []*Relation) error {
	var edgeRefs []interface{}
	var t *gremlingo.GraphTraversal
	g := gremlingo.Traversal_().WithRemote(c.remote)
	// chain the upserts
	for k, relation := range relations {
		v1Ref := fmt.Sprintf("from-%d", k)
		v2Ref := fmt.Sprintf("to-%d", k)
		edgeRef := fmt.Sprintf("e-%d", k)
		edgeRefs = append(edgeRefs, edgeRef)

		// match from
		if k == 0 {
			t = g.V().HasLabel(string(relation.outV.label))
		} else {
			t = t.V().HasLabel(string(relation.outV.label))
		}
		for k, v := range relation.outV.has {
			t = t.Has(k, v)
		}
		t = t.Limit(1).As(v1Ref)
		// match to
		t = t.V().HasLabel(string(relation.inV.label))
		for k, v := range relation.inV.has {
			t = t.Has(k, v)
		}
		t = t.Limit(1).As(v2Ref)
		// upsert edge
		t = t.MergeE(relation.edge.toEdgeMap()).
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
	if len(results) != len(relations) {
		return fmt.Errorf("bulkUpsertRelations: number of results(%d) gathered does not match number of inputs(%d)",
			len(results), len(relations))
	}

	var relationsWithIds []*RelationWithId
	for k, r := range results {
		resultMap := r.GetInterface().(map[interface{}]interface{})
		edgeValue := resultMap["edge"]
		fromMap := resultMap["from"].(map[interface{}]interface{})
		toMap := resultMap["to"].(map[interface{}]interface{})

		var edgeId string
		if c.config.Flavor == JanusGraph {
			edgeId = strconv.FormatInt(edgeValue.(*janusgraphRelationIdentifier).RelationId, 10)
		} else {
			edgeMap := edgeValue.(map[interface{}]interface{})
			edgeId = edgeMap["id"].(string)
		}

		relationWithId := &RelationWithId{
			edgeId:   edgeId,
			relation: relations[k],
			outV:     fromMap,
			inV:      toMap,
		}
		relationsWithIds = append(relationsWithIds, relationWithId)
	}

	queue <- relationsWithIds
	return nil
}
