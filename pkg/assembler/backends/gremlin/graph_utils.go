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
	"github.com/guacsec/guac/pkg/assembler/graphql/model"
	"golang.org/x/sync/errgroup"
	"math"
	"sort"
	"strconv"
)

type Label string

const MaxVertexUpsertChunkSize = 200
const MaxEdgeUpsertChunkSize = 10 // breaks with anything over 10

func (c *gremlinClient) upsertVertex(properties map[interface{}]interface{}) (string, error) {
	g := gremlingo.Traversal_().WithRemote(c.remote)
	r, err := g.MergeV(properties).Id().Next()
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

func (c *gremlinClient) bulkUpsertVertices(allProperties []map[interface{}]interface{}) ([]string, error) {
	var ids []string
	var vertexRefs []interface{}
	var gt *gremlingo.GraphTraversal
	g := gremlingo.Traversal_().WithRemote(c.remote)
	// chain the upserts
	for i, properties := range allProperties {
		vertexRef := fmt.Sprintf("v%d", i)
		vertexRefs = append(vertexRefs, vertexRef)
		if i == 0 {
			gt = g.MergeV(properties).As(vertexRef)
		} else {
			gt = gt.MergeV(properties).As(vertexRef)
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
	if len(ids) != len(allProperties) {
		return nil, fmt.Errorf("bulkUpsertVertices: number of results(%d) gathered does not match number of inputs(%d)",
			len(ids), len(allProperties))
	}

	return ids, nil
}

type MapSerializer[M any] func(model M) (result map[interface{}]interface{})

type ObjectDeserializer[M any] func(id string, values map[interface{}]interface{}) (model M)

type EdgeMapSerializer[VInput any, EInput any] func(v1 VInput, v2 VInput, edge EInput) (result map[interface{}]interface{})

type Relation struct {
	v1   map[interface{}]interface{}
	v2   map[interface{}]interface{}
	edge map[interface{}]interface{}
}

type RelationWithId struct {
	edgeId   string
	relation *Relation
}

/*
C is typically an InputSpec
D is model object w/ id after upsert
*/
func ingestModelObject[C any, D any](c *gremlinClient, modelObject C, serializer MapSerializer[C], deserializer ObjectDeserializer[D]) (D, error) {
	var object D
	values := serializer(modelObject)
	// Verify that label is present
	if _, ok := values[gremlingo.T.Label]; !ok {
		return object, errors.New("missing label!!! please add it :)")
	}

	id, err := c.upsertVertex(values)
	if err != nil {
		return object, err
	}
	object = deserializer(id, values)
	return object, nil
}

func bulkIngestModelObjects[C any, D any](c *gremlinClient, modelObjects []C, serializer MapSerializer[C], deserializer ObjectDeserializer[D]) ([]D, error) {
	var objects []D
	if len(modelObjects) < 1 {
		// nothing to do
		return objects, nil
	}

	// serialize
	var allValues []map[interface{}]interface{}
	for _, modelObject := range modelObjects {
		values := serializer(modelObject)
		allValues = append(allValues, values)
	}

	// split in chunk to limit websocket frame size
	// FIXME: make this configurable
	// FIXME: given these span multiple requests, we should wrap them in a single transaction
	// FIXME: can we do these in parallel? (probably not if they are in a single transaction)
	var ids = make([]string, 0)
	for _, chunk := range chunkSlice(allValues, MaxVertexUpsertChunkSize) {
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
		object := deserializer(ids[k], allValues[k])
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

func storeMapInVertexProperties(properties map[interface{}]interface{}, propertyName string, mapToStore map[string]string) {
	mapToStoreJson, _ := json.Marshal(mapToStore)
	properties[propertyName] = string(mapToStoreJson)
}

func storeArrayInVertexProperties2(properties map[interface{}]interface{}, propertyName string, mapToStore []string) {
	mapToStoreJson, _ := json.Marshal(mapToStore)
	properties[propertyName] = string(mapToStoreJson)
}

func storeArrayInVertexProperties(checks []*model.ScorecardCheckInputSpec, propertyName string, vertexProperties map[interface{}]interface{}) {
	// flatten checks into a list keys (check name) and a list of values (the scores)
	checksMap := map[string]int{}
	var checkKeysList []string
	var checkValuesList []int
	for _, check := range checks {
		key := check.Check
		checksMap[key] = check.Score
		checkKeysList = append(checkKeysList, key)
	}
	// sort for deterministic outputs
	sort.Strings(checkKeysList)
	for _, k := range checkKeysList {
		checkValuesList = append(checkValuesList, checksMap[k])
	}

	// Support for arrays may vary based on the implementation of Gremlin
	// FIXME: 2023/07/09 02:26:18 %!(EXTRA string=gremlinServerWSProtocol.responseHandler(), string=%!(EXTRA gremlingo.responseStatus={244 Property value [[Binary_Artifacts, Branch_Protection, Code_Review, Contributors]] is of type class java.util.ArrayList is not supported map[exceptions:[java.lang.IllegalArgumentException] stackTrace:java.lang.IllegalArgumentException: Property value [[Binary_Artifacts, Branch_Protection, Code_Review, Contributors]] is of type class java.util.ArrayList is not supported
	//	at org.apache.gremlin.gremlin.structure.Property$Exceptions.dataTypeOfPropertyValueNotSupported(Property.java:159)
	//	at org.apache.gremlin.gremlin.structure.Property$Exceptions.dataTypeOfPropertyValueNotSupported(Property.java:155)
	//	at org.janusgraph.graphdb.transaction.StandardJanusGraphTx.verifyAttribute(StandardJanusGraphTx.java:673)
	//	at org.janusgraph.graphdb.transaction.StandardJanusGraphTx.addProperty(StandardJanusGraphTx.java:888)
	//	at org.janusgraph.graphdb.transaction.StandardJanusGraphTx.addProperty(StandardJanusGraphTx.java:877)
	checkKeysJson, _ := json.Marshal(checkKeysList)
	checkValuesJson, _ := json.Marshal(checkValuesList)
	vertexProperties["checkKeys"] = checkKeysJson
	vertexProperties["checkValues"] = checkValuesJson
}

func readArrayFromVertexProperties(results map[interface{}]interface{}) ([]*model.ScorecardCheck, error) {
	var keyList = results[checkKeys].([]interface{})
	var valueList = results[checkValues].([]interface{})

	var checks []*model.ScorecardCheck
	if len(keyList) != len(valueList) {
		//log.Error("values on vertices do do match - corrupt - need to delete vertex")
		valueList = valueList[:len(keyList)]
	}
	for i := range keyList {
		valueAsJson := valueList[i].(string)
		score, err := strconv.ParseInt(valueAsJson[1:len(valueAsJson)-1], 10, 0)
		if err != nil {
			return nil, err
		}
		check := &model.ScorecardCheck{
			Check: keyList[i].(string),
			Score: int(score),
		}
		checks = append(checks, check)
	}
	return checks, nil
}

func ingestModelObjectsWithRelation[VInput any, EInput any, EOutput any](c *gremlinClient,
	v1InputObject VInput,
	v2InputObject VInput,
	edgeInputObject EInput,
	vInputSerializer MapSerializer[VInput],
	edgeInputSerializer EdgeMapSerializer[VInput, EInput],
	edgeOutputDeserializer ObjectDeserializer[EOutput]) (EOutput, error) {

	var edgeObject EOutput

	relation := &Relation{
		v1:   vInputSerializer(v1InputObject),
		v2:   vInputSerializer(v2InputObject),
		edge: edgeInputSerializer(v1InputObject, v2InputObject, edgeInputObject),
	}
	relation.edge[gremlingo.Direction.In] = gremlingo.Merge.InV
	relation.edge[gremlingo.Direction.Out] = gremlingo.Merge.OutV

	ch := make(chan []*RelationWithId, 1)
	err := c.upsertRelation(ch, relation)
	if err != nil {
		return edgeObject, err
	}
	close(ch)
	// return the first element from the first list in the channel
	for _, relationWithId := range <-ch {
		return edgeOutputDeserializer(relationWithId.edgeId, relationWithId.relation.edge), nil
	}
	return edgeObject, errors.New("no results returned by upsert")
}

func bulkIngestModelObjectsWithRelation[VInput any, EInput any, EOutput any](c *gremlinClient,
	v1InputObjects []VInput,
	v2InputObjects []VInput,
	edgeInputObjects []EInput,
	vInputSerializer MapSerializer[VInput],
	edgeInputSerializer EdgeMapSerializer[VInput, EInput],
	edgeOutputDeserializer ObjectDeserializer[EOutput]) ([]EOutput, error) {

	var objects []EOutput
	// back input validation
	if len(v1InputObjects) < 1 {
		// nothing to do
		return objects, nil
	}
	if len(v1InputObjects) != len(v2InputObjects) || len(v1InputObjects) != len(edgeInputObjects) {
		return objects, fmt.Errorf("size of inputs do not match v1(%d) v2(%d) e(%d)",
			len(v1InputObjects), len(v2InputObjects), len(edgeInputObjects))
	}

	// serialize all the values
	var allRelations []*Relation
	for k, v1InputObject := range v1InputObjects {
		v2InputObject := v2InputObjects[k]
		edgeInputObject := edgeInputObjects[k]
		relation := &Relation{
			v1:   vInputSerializer(v1InputObject),
			v2:   vInputSerializer(v2InputObject),
			edge: edgeInputSerializer(v1InputObject, v2InputObject, edgeInputObject),
		}
		relation.edge[gremlingo.Direction.In] = gremlingo.Merge.InV
		relation.edge[gremlingo.Direction.Out] = gremlingo.Merge.OutV
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
			object := edgeOutputDeserializer(relationWithId.edgeId, relationWithId.relation.edge)
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
	g := gremlingo.Traversal_().WithRemote(c.remote)
	r, err := g.MergeV(relation.v1).As("v1").
		MergeV(relation.v2).As("v2").
		MergeE(relation.edge).As("edge").
		// late bind
		Option(gremlingo.Merge.InV, gremlingo.T__.Select("v1")).
		Option(gremlingo.Merge.OutV, gremlingo.T__.Select("v2")).
		Select("edge").Id().Next()
	if err != nil {
		return err
	}

	var relationsWithIds []*RelationWithId

	var edgeId string
	if c.config.Flavor == JanusGraph {
		edgeId = strconv.FormatInt(r.GetInterface().(*janusgraphRelationIdentifier).RelationId, 10)
	} else {
		edgeId = r.GetString()
	}

	relationsWithIds = append(relationsWithIds, &RelationWithId{
		edgeId:   edgeId,
		relation: relation,
	})

	queue <- relationsWithIds
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
	var gt *gremlingo.GraphTraversal
	g := gremlingo.Traversal_().WithRemote(c.remote)
	// chain the upserts
	for k, relation := range relations {
		v1Ref := fmt.Sprintf("v1-%d", k)
		v2Ref := fmt.Sprintf("v2-%d", k)
		edgeRef := fmt.Sprintf("e-%d", k)
		edgeRefs = append(edgeRefs, edgeRef)

		if k == 0 {
			gt = g.MergeV(relation.v1).As(v1Ref).
				MergeV(relation.v2).As(v2Ref).
				MergeE(relation.edge).As(edgeRef).
				// late bind
				Option(gremlingo.Merge.InV, gremlingo.T__.Select(v1Ref)).
				Option(gremlingo.Merge.OutV, gremlingo.T__.Select(v2Ref))
		} else {
			gt = gt.MergeV(relation.v1).As(v1Ref).
				MergeV(relation.v2).As(v2Ref).
				MergeE(relation.edge).As(edgeRef).
				// late bind
				Option(gremlingo.Merge.InV, gremlingo.T__.Select(v1Ref)).
				Option(gremlingo.Merge.OutV, gremlingo.T__.Select(v2Ref))
		}
	}
	results, err := gt.Select(edgeRefs...).Select(gremlingo.Column.Values).Unfold().Id().ToList()
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
		var edgeId string
		if c.config.Flavor == JanusGraph {
			edgeId = strconv.FormatInt(r.GetInterface().(*janusgraphRelationIdentifier).RelationId, 10)
		} else {
			edgeId = r.GetString()
		}

		relationsWithIds = append(relationsWithIds, &RelationWithId{
			edgeId: edgeId,
			// FIXME: This assumes the IDs are returned in the same order as the input queries
			relation: relations[k],
		})
	}

	queue <- relationsWithIds
	return nil
}
