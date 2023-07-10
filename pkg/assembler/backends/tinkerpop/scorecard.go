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
	"context"
	"encoding/json"
	gremlingo "github.com/apache/tinkerpop/gremlin-go/driver"
	"github.com/guacsec/guac/pkg/assembler/graphql/model"
	"github.com/vektah/gqlparser/v2/gqlerror"
	"sort"
	"strings"
	"time"
)

const (
	aggregateScore   string = "aggregateScore"
	checkKeys        string = "checkKeys"
	checkValues      string = "checkValues"
	collector        string = "collector"
	name             string = "name"
	namespace        string = "namespace"
	origin           string = "origin"
	scorecardVersion string = "scorecardVersion"
	scorecardCommit  string = "scorecardCommit"
	sourceType       string = "sourceType"
	timeScanned      string = "timeScanned"
)

// CertifyScorecard used to ingest scorecards
func (c *tinkerpopClient) CertifyScorecard(ctx context.Context, source model.SourceInputSpec, scorecard model.ScorecardInputSpec) (*model.CertifyScorecard, error) {
	values := map[string]any{}
	// FIXME: dedup
	if source.Commit != nil && source.Tag != nil {
		if *source.Commit != "" && *source.Tag != "" {
			return nil, gqlerror.Errorf("Passing both commit and tag selectors is an error")
		}
	}

	if source.Commit != nil {
		values["commit"] = *source.Commit
	} else {
		values["commit"] = ""
	}

	if source.Tag != nil {
		values["tag"] = *source.Tag
	} else {
		values["tag"] = ""
	}

	// flatten checks into a list keys (check name) and a list of values (the scores)
	checksMap := map[string]int{}
	checkKeysList := []string{}
	checkValuesList := []int{}
	for _, check := range scorecard.Checks {
		key := removeInvalidCharFromProperty(check.Check)
		checksMap[key] = check.Score
		checkKeysList = append(checkKeysList, key)
	}
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

	// find the source vertex, upsert the scorecard, and make sure there's an edge created between Source and ScoreCard
	g := gremlingo.Traversal_().WithRemote(c.remote)
	tx := g.Tx()
	gtx, _ := tx.Begin()

	r, err := gtx.V().HasLabel(source.Name).Has(sourceType, source.Type).
		Has(namespace, source.Namespace).
		As("source").
		AddV().
		Property(checkKeys, string(checkKeysJson)).
		Property(checkValues, string(checkValuesJson)).
		Property(aggregateScore, scorecard.AggregateScore).
		Property(timeScanned, scorecard.TimeScanned.UTC()).
		Property(scorecardVersion, scorecard.ScorecardVersion).
		Property(scorecardCommit, scorecard.ScorecardCommit).
		Property(origin, scorecard.Origin).
		Property(collector, scorecard.Collector).
		As("scorecard").
		AddE("scorecard").From("source").To("scorecard").
		Select("scorecard").
		ElementMap().Next()
	if err != nil {
		return nil, err
	}
	resultMap := r.GetInterface().(map[interface{}]interface{})

	checks, err := getCollectedChecks(
		resultMap[checkKeys].(string),
		resultMap[checkValues].(string))
	if err != nil {
		return nil, err
	}

	src := generateModelSource(source.Type, source.Namespace, source.Name, nil, nil)

	modelScorecard := model.Scorecard{
		TimeScanned:      resultMap[timeScanned].(time.Time),
		AggregateScore:   resultMap[aggregateScore].(float64),
		Checks:           checks,
		ScorecardVersion: resultMap[scorecardVersion].(string),
		ScorecardCommit:  resultMap[scorecardCommit].(string),
		Origin:           resultMap[origin].(string),
		Collector:        resultMap[collector].(string),
	}

	certification := model.CertifyScorecard{
		Source:    src,
		Scorecard: &modelScorecard,
	}

	return &certification, nil
}

func getCollectedChecks(keyListJson string, valueListJson string) ([]*model.ScorecardCheck, error) {

	var keyList []string
	var valueList []int64

	err := json.Unmarshal([]byte(keyListJson), &keyList)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal([]byte(valueListJson), &valueList)
	if err != nil {
		return nil, err
	}

	if len(keyList) != len(valueList) {
		return nil, gqlerror.Errorf("length of scorecard checks do not match")
	}
	checks := []*model.ScorecardCheck{}
	for i := range keyList {
		check := &model.ScorecardCheck{
			Check: keyList[i],
			Score: int(valueList[i]),
		}
		checks = append(checks, check)
	}
	return checks, nil
}

// FIXME: We should be able to store this verbatim
func removeInvalidCharFromProperty(key string) string {
	// neo4j does not accept "." in its properties. If the qualifier contains a "." that must
	// be replaced by an "-"
	return strings.ReplaceAll(key, ".", "_")
}
