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
	"github.com/vektah/gqlparser/v2/gqlerror"
	"time"
)

const (
	commit string = "commit"
	tag    string = "tag"
)

const (
	aggregateScore   string = "aggregateScore"
	checksJson       string = "checksJson"
	checkKeys        string = "checkKeys"
	checkValues      string = "checkValues"
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

func validateSourceInputSpec(source model.SourceInputSpec) error {
	if source.Commit != nil && source.Tag != nil {
		if *source.Commit != "" && *source.Tag != "" {
			return gqlerror.Errorf("Passing both commit and tag selectors is an error")
		}
	}
	return nil
}

func getScorecardQueryValues(scorecard *model.ScorecardInputSpec) *GraphQuery {
	checks := toChecks(scorecard.Checks)
	checksJsonValue, err := json.Marshal(checks)
	if err != nil {
		checksJsonValue = nil
	}
	q := createGraphQuery(Scorecard)
	q.has[aggregateScore] = scorecard.AggregateScore
	q.has[timeScanned] = scorecard.TimeScanned.UTC()
	q.has[scorecardVersion] = scorecard.ScorecardVersion
	q.has[scorecardCommit] = scorecard.ScorecardCommit
	q.has[origin] = scorecard.Origin
	q.has[collector] = scorecard.Collector
	q.has[checksJson] = string(checksJsonValue)
	return q
}

func getSourceMatchQueryValues(source *model.SourceInputSpec) *GraphQuery {
	q := createGraphQuery(Source)
	q.has[name] = source.Name
	q.has[typeStr] = source.Type
	q.has[namespace] = source.Namespace

	if source.Commit != nil {
		q.has[commit] = *source.Commit
	}
	if source.Tag != nil {
		q.has[tag] = *source.Tag
	}

	return q
}

// IngestScorecard
//
//	scorecard -> ScorecardToSource -> src
func (c *gremlinClient) IngestScorecard(ctx context.Context, source model.SourceInputSpec, scorecard model.ScorecardInputSpec) (*model.CertifyScorecard, error) {
	// TODO: Can we push this validation up a layer, so that the storage engines don't need to worry about it?
	err := validateSourceInputSpec(source)
	if err != nil {
		return nil, err
	}

	//	scorecard -> ScorecardToSource -> src
	edgeQ := createGraphQuery(ScorecardToSource)
	// copy to easy for easy sorting
	edgeQ.has[aggregateScore] = scorecard.AggregateScore

	scorecardQ := getScorecardQueryValues(&scorecard)
	sourceQ := getSourceMatchQueryValues(&source)

	relation := &Relation{
		outV:       scorecardQ,
		upsertOutV: true,
		inV:        sourceQ,
		edge:       edgeQ,
	}
	relationWithId, err := c.upsertRelationDirect(relation)
	if err != nil {
		return nil, err
	}

	// build artifact from canonical model after a successful upsert
	src := generateModelSource(source.Type, source.Namespace, source.Name, nil, nil)
	modelScorecard := model.Scorecard{
		TimeScanned:      scorecard.TimeScanned,
		AggregateScore:   scorecard.AggregateScore,
		Checks:           toChecks(scorecard.Checks),
		ScorecardVersion: scorecard.ScorecardVersion,
		ScorecardCommit:  scorecard.ScorecardCommit,
		Origin:           scorecard.Origin,
		Collector:        scorecard.Collector,
	}
	certification := model.CertifyScorecard{
		ID:        relationWithId.edgeId,
		Source:    src,
		Scorecard: &modelScorecard,
	}
	return &certification, nil
}

func (c *gremlinClient) IngestScorecards(ctx context.Context, sources []*model.SourceInputSpec, scorecards []*model.ScorecardInputSpec) ([]*model.CertifyScorecard, error) {
	// FIXME: Implement bulk insert
	var scorecardObjects []*model.CertifyScorecard
	for k, scorecardSpec := range scorecards {
		scorecard, err := c.IngestScorecard(ctx, *sources[k], *scorecardSpec)
		if err != nil {
			return scorecardObjects, err
		}
		scorecardObjects = append(scorecardObjects, scorecard)
	}
	return scorecardObjects, nil
}

// CertifyScorecard an existing alias for ingesting scorecards
func (c *gremlinClient) CertifyScorecard(ctx context.Context, source model.SourceInputSpec, scorecard model.ScorecardInputSpec) (*model.CertifyScorecard, error) {
	return c.IngestScorecard(ctx, source, scorecard)
}

func createQueryToMatchSource(src *model.SourceSpec) *gremlinQueryBuilder {
	query := createGraphQuery(Source)
	if src.ID != nil {
		query.id = *src.ID
	}
	if src.Name != nil {
		query.has[name] = *src.Name
	}
	if src.Type != nil {
		query.has[typeStr] = *src.Type
	}
	if src.Namespace != nil {
		query.has[namespace] = *src.Namespace
	}
	if src.Commit != nil {
		query.has[commit] = *src.Commit
	}
	if src.Tag != nil {
		query.has[tag] = *src.Tag
	}
	return &gremlinQueryBuilder{query: query}
}

func (c *gremlinClient) Scorecards(ctx context.Context, certifyScorecardSpec *model.CertifyScorecardSpec) ([]*model.CertifyScorecard, error) {
	q := createQueryForEdge(ScorecardToSource)
	if certifyScorecardSpec != nil {
		if certifyScorecardSpec.Source != nil {
			q = q.withInVertex(createQueryToMatchSource(certifyScorecardSpec.Source))
		}
		scorecardQ := createQueryForEdge(Scorecard).
			withId(certifyScorecardSpec.ID).
			withPropString(scorecardVersion, certifyScorecardSpec.ScorecardVersion).
			withPropString(commit, certifyScorecardSpec.ScorecardCommit).
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
	q.query.orderByKey = aggregateScore
	q.query.orderByDirection = gremlingo.Order.Asc
	return queryEdge[*model.CertifyScorecard](c, q, getScorecardObjectFromEdge)
}

func getScorecardObjectFromEdge(id string, out map[interface{}]interface{}, edge map[interface{}]interface{}, in map[interface{}]interface{}) *model.CertifyScorecard {
	var checks []*model.ScorecardCheck
	err := json.Unmarshal([]byte(out[checksJson].(string)), &checks)
	if err != nil {
		checks = nil
	}

	scorecard := &model.CertifyScorecard{
		ID: id,
		Scorecard: &model.Scorecard{
			TimeScanned:      out[timeScanned].(time.Time),
			AggregateScore:   out[aggregateScore].(float64),
			Checks:           checks,
			ScorecardVersion: out[scorecardVersion].(string),
			ScorecardCommit:  out[scorecardCommit].(string),
			Origin:           out[origin].(string),
			Collector:        out[collector].(string),
		},
		Source: &model.Source{
			Type: in[typeStr].(string),
			Namespaces: []*model.SourceNamespace{{Namespace: in[namespace].(string), Names: []*model.SourceName{
				{Name: in[name].(string), Tag: ptrfrom.String(in[tag].(string)), Commit: ptrfrom.String(in[commit].(string))}}}},
		},
	}

	return scorecard
}

func toChecks(inputCheck []*model.ScorecardCheckInputSpec) []*model.ScorecardCheck {
	checks := make([]*model.ScorecardCheck, 0)
	for _, check := range inputCheck {
		checks = append(checks, toCheck(check))
	}
	return checks
}

func toCheck(inputCheck *model.ScorecardCheckInputSpec) *model.ScorecardCheck {
	return &model.ScorecardCheck{
		Check: inputCheck.Check,
		Score: inputCheck.Score,
	}
}
