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

//func (c *tinkerpopClient) bulkUpsertVertex(properties []map[interface{}]interface{}) ([]int64, error) {
//	g := gremlingo.Traversal_().WithRemote(c.remote)
//	r, err := g.MergeV(properties).Id().Next()
//	if err != nil {
//		return -1, err
//	}
//	return r.GetInt64()
//}

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
