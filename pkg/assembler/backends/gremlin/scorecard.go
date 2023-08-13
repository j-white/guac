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
	"context"
	"encoding/json"
	gremlingo "github.com/apache/tinkerpop/gremlin-go/v3/driver"
	"github.com/guacsec/guac/internal/testing/ptrfrom"
	"github.com/guacsec/guac/pkg/assembler/graphql/model"
	"time"
)

const (
	commit string = "commit"
	tag    string = "tag"
)

const (
	aggregateScore   string = "aggregateScore"
	checksJson       string = "checksJson"
	collector        string = "collector"
	name             string = "name"
	namespace        string = "namespace"
	origin           string = "origin"
	scorecardVersion string = "scorecardVersion"
	scorecardCommit  string = "scorecardCommit"
	timeScanned      string = "timeScanned"
	version          string = "version"
	subpath          string = "subpath"
	versionRange     string = "versionRange"
	dependencyType   string = "dependencyType"
)

const (
	Source            Label = "source"
	Scorecard         Label = "scorecard"
	ScorecardToSource Label = "scorecard-to-source"
)

func createUpsertForScorecardVertex(scorecard *model.ScorecardInputSpec) *gremlinQueryBuilder[*model.CertifyScorecard] {
	q := createUpsertForVertex[*model.CertifyScorecard](Scorecard).
		withPropString(scorecardVersion, &scorecard.ScorecardVersion).
		withPropString(scorecardCommit, &scorecard.ScorecardCommit).
		withPropString(collector, &scorecard.Collector).
		withPropString(origin, &scorecard.Origin).
		withPropTime(timeScanned, &scorecard.TimeScanned).
		withPropFloat64(aggregateScore, &scorecard.AggregateScore)
	if scorecard.Checks != nil {
		checksJsonValue, err := json.Marshal(scorecard.Checks)
		if err != nil {
			checksJsonValue = nil
		}
		q = q.withPropString(checksJson, ptrfrom.String(string(checksJsonValue)))
	} else {
		q = q.withPropString(checksJson, ptrfrom.String("[]"))
	}
	return q
}

func createUpsertForScorecard(source *model.SourceInputSpec, scorecard *model.ScorecardInputSpec) *gremlinQueryBuilder[*model.CertifyScorecard] {
	return createUpsertForEdge[*model.CertifyScorecard](ScorecardToSource).
		// used for sorting
		withPropFloat64(aggregateScore, &scorecard.AggregateScore).
		withOutVertex(createUpsertForScorecardVertex(scorecard)).
		withInVertex(createQueryToMatchSourceInput[*model.CertifyScorecard](source)).
		withMapper(getScorecardObjectFromEdge)
}

func (c *gremlinClient) IngestScorecard(ctx context.Context, source model.SourceInputSpec, scorecard model.ScorecardInputSpec) (*model.CertifyScorecard, error) {
	return createUpsertForScorecard(&source, &scorecard).upsert(c)
}

func (c *gremlinClient) CertifyScorecard(ctx context.Context, source model.SourceInputSpec, scorecard model.ScorecardInputSpec) (*model.CertifyScorecard, error) {
	return c.IngestScorecard(ctx, source, scorecard)
}

func (c *gremlinClient) IngestScorecards(ctx context.Context, sources []*model.SourceInputSpec, scorecards []*model.ScorecardInputSpec) ([]*model.CertifyScorecard, error) {
	var queries []*gremlinQueryBuilder[*model.CertifyScorecard]
	for k := range sources {
		queries = append(queries, createUpsertForScorecard(sources[k], scorecards[k]))
	}

	return createBulkUpsertForEdge[*model.CertifyScorecard](Scorecard).
		withQueries(queries).
		upsertBulk(c)
}

func (c *gremlinClient) Scorecards(ctx context.Context, certifyScorecardSpec *model.CertifyScorecardSpec) ([]*model.CertifyScorecard, error) {
	q := createQueryForEdge[*model.CertifyScorecard](ScorecardToSource).
		withOrderByKey(aggregateScore).
		withOrderByDirection(gremlingo.Order.Asc).
		withMapper(getScorecardObjectFromEdge)
	if certifyScorecardSpec != nil {
		if certifyScorecardSpec.Source != nil {
			q = q.withInVertex(createQueryToMatchSource[*model.CertifyScorecard](certifyScorecardSpec.Source))
		}
		scorecardQ := createQueryForVertex[*model.CertifyScorecard](Scorecard).
			withId(certifyScorecardSpec.ID).
			withPropString(scorecardVersion, certifyScorecardSpec.ScorecardVersion).
			withPropString(scorecardCommit, certifyScorecardSpec.ScorecardCommit).
			withPropString(collector, certifyScorecardSpec.Collector).
			withPropString(origin, certifyScorecardSpec.Origin).
			withPropTime(timeScanned, certifyScorecardSpec.TimeScanned).
			withPropFloat64(aggregateScore, certifyScorecardSpec.AggregateScore)
		if certifyScorecardSpec.Checks != nil {
			checksJsonValue, err := json.Marshal(certifyScorecardSpec.Checks)
			if err != nil {
				checksJsonValue = nil
			}
			scorecardQ = scorecardQ.withPropString(checksJson, ptrfrom.String(string(checksJsonValue)))
		}
		q = q.withOutVertex(scorecardQ)
	}
	return q.findAll(c)
}

func getScorecardObjectFromEdge(result *gremlinQueryResult) *model.CertifyScorecard {
	var checks []*model.ScorecardCheck
	if result.out[checksJson] != nil {
		err := json.Unmarshal([]byte(result.out[checksJson].(string)), &checks)
		if err != nil {
			checks = nil
		}
	}

	scorecard := &model.CertifyScorecard{
		ID: result.id,
		Scorecard: &model.Scorecard{
			TimeScanned:      result.out[timeScanned].(time.Time),
			AggregateScore:   result.out[aggregateScore].(float64),
			Checks:           checks,
			ScorecardVersion: result.out[scorecardVersion].(string),
			ScorecardCommit:  result.out[scorecardCommit].(string),
			Origin:           result.out[origin].(string),
			Collector:        result.out[collector].(string),
		},
		Source: &model.Source{
			Type: result.in[typeStr].(string),
			Namespaces: []*model.SourceNamespace{{Namespace: result.in[namespace].(string), Names: []*model.SourceName{
				{Name: result.in[name].(string), Tag: ptrfrom.String(result.in[tag].(string)), Commit: ptrfrom.String(result.in[commit].(string))}}}},
		},
	}

	return scorecard
}
