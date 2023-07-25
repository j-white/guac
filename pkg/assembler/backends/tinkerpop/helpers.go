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

package tinkerpop

import (
	"encoding/json"
	"errors"
	"fmt"
	gremlingo "github.com/apache/tinkerpop/gremlin-go/v3/driver"
	"github.com/guacsec/guac/pkg/assembler/graphql/model"
	"sort"
	"strconv"
)

type Label string

func (c *tinkerpopClient) upsertVertex(properties map[interface{}]interface{}) (int64, error) {
	g := gremlingo.Traversal_().WithRemote(c.remote)
	r, err := g.MergeV(properties).Id().Next()
	if err != nil {
		return -1, err
	}
	return r.GetInt64()
}

func (c *tinkerpopClient) bulkUpsertVertices(allProperties []map[interface{}]interface{}) ([]int64, error) {
	var ids []int64
	var gt *gremlingo.GraphTraversal
	g := gremlingo.Traversal_().WithRemote(c.remote)
	// chain all of the upserts
	for i, properties := range allProperties {
		vertexRef := fmt.Sprintf("v:%d", i)
		if i == 0 {
			gt = g.MergeV(properties).As(vertexRef)
		} else {
			gt = gt.MergeV(properties).As(vertexRef)
		}
	}

	results, err := gt.ToList()
	if err != nil {
		return nil, err
	}

	for _, result := range results {
		resultMap := result.GetInterface().(map[interface{}]interface{})
		ids = append(ids, resultMap[string(gremlingo.T.Id)].(int64))
	}

	return ids, nil
}

type MapSerializer[M any] func(model M) (result map[interface{}]interface{})

type ObjectDeserializer[M any] func(id int64, values map[interface{}]interface{}) (model M)

func ingestModelObject[C any, D any](c *tinkerpopClient, modelObject C, serializer MapSerializer[C], deserializer ObjectDeserializer[D]) (D, error) {
	var object D
	values := serializer(modelObject)
	// Verify that label is present
	if _, ok := values[gremlingo.T.Label]; ok {
		return object, errors.New("missing label!!! please add it :)")
	}

	id, err := c.upsertVertex(values)
	if err != nil {
		return object, err
	}
	object = deserializer(id, values)
	return object, nil
}

func bulkIngestModelObjects[C any, D any](c *tinkerpopClient, modelObjects []C, serializer MapSerializer[C], deserializer ObjectDeserializer[D]) ([]D, error) {
	var objects []D

	// serialize
	var allValues []map[interface{}]interface{}
	for _, modelObject := range modelObjects {
		values := serializer(modelObject)
		allValues = append(allValues, values)
	}

	ids, err := c.bulkUpsertVertices(allValues)
	if err != nil {
		return objects, err
	}

	if len(ids) != len(modelObjects) {
		return nil, errors.New("the lengths dont match, I am sad")
	}

	for k := range modelObjects {
		object := deserializer(ids[k], allValues[k])
		objects = append(objects, object)
	}
	return objects, nil
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
	//	at org.apache.tinkerpop.gremlin.structure.Property$Exceptions.dataTypeOfPropertyValueNotSupported(Property.java:159)
	//	at org.apache.tinkerpop.gremlin.structure.Property$Exceptions.dataTypeOfPropertyValueNotSupported(Property.java:155)
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
