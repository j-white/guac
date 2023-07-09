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
	"github.com/guacsec/guac/pkg/assembler/graphql/model"
	"github.com/vektah/gqlparser/v2/gqlerror"
	"sort"
	"strings"
)

const (
	timeScanned      string = "timeScanned"
	aggregateScore   string = "aggregateScore"
	checkKeys        string = "checkKeys"
	checkValues      string = "checkValues"
	scorecardVersion string = "scorecardVersion"
	scorecardCommit  string = "scorecardCommit"
)

// Ingest Scorecards

func (c *tinkerpopClient) CertifyScorecard(ctx context.Context, source model.SourceInputSpec, scorecard model.ScorecardInputSpec) (*model.CertifyScorecard, error) {
	values := map[string]any{}
	values["sourceType"] = source.Type
	values["namespace"] = source.Namespace
	values["name"] = source.Name

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

	values[timeScanned] = scorecard.TimeScanned.UTC()
	values[aggregateScore] = scorecard.AggregateScore
	values[scorecardVersion] = scorecard.ScorecardVersion
	values[scorecardCommit] = scorecard.ScorecardCommit

	// Cannot use getScorecardChecks due to type mismatch
	// Generics would be really helpful here :)
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
	values[checkKeys] = checkKeysList
	values[checkValues] = checkValuesList

	// TODO(mihaimaruseac): Should we put origin/collector on the edge instead?
	values["origin"] = scorecard.Origin
	values["collector"] = scorecard.Collector

	//g := gremlingo.Traversal_().WithRemote(c.remote)
	//tx := g.Tx()
	//gtx, _ := tx.Begin()
	//
	//result, err := session.WriteTransaction(
	//	func(tx neo4j.Transaction) (interface{}, error) {
	//		query := `
	//MATCH (root:Src) -[:SrcHasType]-> (type:SrcType) -[:SrcHasNamespace]-> (ns:SrcNamespace) -[:SrcHasName] -> (name:SrcName)
	//WHERE type.type = $sourceType AND ns.namespace = $namespace AND name.name = $name AND name.commit = $commit AND name.tag = $tag
	//MERGE (name) <-[:subject]- (certifyScorecard:CertifyScorecard{timeScanned:$timeScanned,aggregateScore:$aggregateScore,scorecardVersion:$scorecardVersion,scorecardCommit:$scorecardCommit,checkKeys:$checkKeys,checkValues:$checkValues,origin:$origin,collector:$collector})
	//RETURN type.type, ns.namespace, name.name, name.commit, name.tag, certifyScorecard`
	//		result, err := tx.Run(query, values)
	//		if err != nil {
	//			return nil, err
	//		}
	//
	//		// query returns a single record
	//		record, err := result.Single()
	//		if err != nil {
	//			return nil, err
	//		}
	//
	//		// TODO(mihaimaruseac): Profile to compare returning node vs returning list of properties
	//		certifyScorecardNode := record.Values[5].(dbtype.Node)
	//		checks, err := getCollectedChecks(
	//			certifyScorecardNode.Props[checkKeys].([]interface{}),
	//			certifyScorecardNode.Props[checkValues].([]interface{}))
	//		if err != nil {
	//			return nil, err
	//		}
	//
	//		scorecard := model.Scorecard{
	//			TimeScanned:      certifyScorecardNode.Props[timeScanned].(time.Time),
	//			AggregateScore:   certifyScorecardNode.Props[aggregateScore].(float64),
	//			Checks:           checks,
	//			ScorecardVersion: certifyScorecardNode.Props[scorecardVersion].(string),
	//			ScorecardCommit:  certifyScorecardNode.Props[scorecardCommit].(string),
	//			Origin:           certifyScorecardNode.Props[origin].(string),
	//			Collector:        certifyScorecardNode.Props[collector].(string),
	//		}
	//
	//		tag := record.Values[4]
	//		commit := record.Values[3]
	//		nameStr := record.Values[2].(string)
	//		namespaceStr := record.Values[1].(string)
	//		srcType := record.Values[0].(string)
	//
	//		src := generateModelSource(srcType, namespaceStr, nameStr, commit, tag)
	//
	//		certification := model.CertifyScorecard{
	//			Source:    src,
	//			Scorecard: &scorecard,
	//		}
	//
	//		return &certification, nil
	//	})
	//if err != nil {
	//	return nil, err
	//}
	//
	//return result.(*model.CertifyScorecard), nil
	return nil, nil
}

func removeInvalidCharFromProperty(key string) string {
	// neo4j does not accept "." in its properties. If the qualifier contains a "." that must
	// be replaced by an "-"
	return strings.ReplaceAll(key, ".", "_")
}
