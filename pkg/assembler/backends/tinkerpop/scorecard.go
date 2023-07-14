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
	gremlingo "github.com/apache/tinkerpop/gremlin-go/v3/driver"
	"github.com/guacsec/guac/pkg/assembler/graphql/model"
	"github.com/vektah/gqlparser/v2/gqlerror"
	"sort"
	"strconv"
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

const (
	Source    Label = "source"
	Scorecard Label = "scorecard"
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

	sourceProperties := map[interface{}]interface{}{
		gremlingo.T.Label: string(Source),
		name:              source.Name,
		"type":            source.Type,
		namespace:         source.Namespace,
	}

	scorecardProperties := map[interface{}]interface{}{
		gremlingo.T.Label: string(Scorecard),
		checkKeys:         string(checkKeysJson),
		checkValues:       string(checkValuesJson),
		aggregateScore:    scorecard.AggregateScore,
		timeScanned:       scorecard.TimeScanned.UTC(),
		scorecardVersion:  scorecard.ScorecardVersion,
		scorecardCommit:   scorecard.ScorecardCommit,
		origin:            scorecard.Origin,
		collector:         scorecard.Collector,
	}

	edgeProperties := map[interface{}]interface{}{
		gremlingo.T.Label:       string(Scorecard),
		gremlingo.Direction.In:  gremlingo.Merge.InV,
		gremlingo.Direction.Out: gremlingo.Merge.OutV,
	}

	r, err := gtx.MergeV(sourceProperties).As("source").
		MergeV(scorecardProperties).As("scorecard").
		MergeE(edgeProperties).
		// late bind
		Option(gremlingo.Merge.InV, gremlingo.T__.Select("source")).
		Option(gremlingo.Merge.OutV, gremlingo.T__.Select("scorecard")).
		As("edge").
		Select("scorecard").
		ValueMap(true).
		Limit(1).
		Next()
	if err != nil {
		return nil, err
	}
	resultMap := r.GetInterface().(map[interface{}]interface{})

	checks, err := getCollectedChecks(
		resultMap[checkKeys].([]interface{}),
		resultMap[checkValues].([]interface{}))
	if err != nil {
		return nil, err
	}

	src := generateModelSource(source.Type, source.Namespace, source.Name, nil, nil)
	modelScorecard := model.Scorecard{
		TimeScanned:      (resultMap[timeScanned].([]interface{}))[0].(time.Time),
		AggregateScore:   (resultMap[aggregateScore].([]interface{}))[0].(float64),
		Checks:           checks,
		ScorecardVersion: (resultMap[scorecardVersion].([]interface{}))[0].(string),
		ScorecardCommit:  (resultMap[scorecardCommit].([]interface{}))[0].(string),
		Origin:           (resultMap[origin].([]interface{}))[0].(string),
		Collector:        (resultMap[collector].([]interface{}))[0].(string),
	}
	certification := model.CertifyScorecard{
		ID:        strconv.FormatInt(resultMap[string(gremlingo.T.Id)].(int64), 10),
		Source:    src,
		Scorecard: &modelScorecard,
	}

	return &certification, nil
}

func getCollectedChecks(keyList []interface{}, valueList []interface{}) ([]*model.ScorecardCheck, error) {
	if len(keyList) != len(valueList) {
		return nil, gqlerror.Errorf("length of scorecard checks do not match")
	}
	var checks []*model.ScorecardCheck
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

// FIXME: We should be able to store this verbatim
func removeInvalidCharFromProperty(key string) string {
	// neo4j does not accept "." in its properties. If the qualifier contains a "." that must
	// be replaced by an "-"
	return strings.ReplaceAll(key, ".", "_")
}
