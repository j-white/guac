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
	"fmt"
	gremlingo "github.com/apache/tinkerpop/gremlin-go/v3/driver"
	"github.com/guacsec/guac/pkg/assembler/graphql/model"
	"github.com/vektah/gqlparser/v2/gqlerror"
	"strconv"
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
	sourceType       string = "type"
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

// IngestScorecard
//
//	scorecard -> ScorecardToSource -> src
func (c *gremlinClient) IngestScorecard(ctx context.Context, source model.SourceInputSpec, scorecard model.ScorecardInputSpec) (*model.CertifyScorecard, error) {
	// scorecard -> scorecard vertex

	// source to match criteria

	// MergeV().as(scorecard).V().has().limit(1).MergeE().project()
	//

	// 	q := createQueryForEdge(IsDependency).

	//
	//q := upsertModelObject(Scorecard, scorecard)
	//		.withMapper()
	//	.relateToEdge(Source)

	return nil, nil
	//return ingestModelObjectsWithRelation[*model.ScorecardInputSpec, *model.IsDependencyInputSpec, *model.IsDependency](
	//	c, &pkg, &depPkg, &dependency, getPackageQueryValues, getPackageQueryValuesForDep, getDependencyQueryValues, getDependencyObjectFromEdgeMuted)
}

func (c *gremlinClient) IngestScorecard2(ctx context.Context, source model.SourceInputSpec, scorecard model.ScorecardInputSpec) (*model.CertifyScorecard, error) {
	// TODO: Can we push this validation up a layer, so that the storage engines don't need to worry about it?
	err := validateSourceInputSpec(source)
	if err != nil {
		return nil, err
	}

	// map to vertices and edges
	sourceQ := createGraphQuery(Source)
	sourceQ.has[name] = source.Name
	sourceQ.has[sourceType] = source.Type
	sourceQ.has[namespace] = source.Namespace

	// optional values, at least one of these must exist
	if source.Tag != nil {
		sourceQ.has[tag] = *source.Tag
	}
	if source.Commit != nil {
		sourceQ.has[commit] = *source.Commit
	}

	checks := toChecks(scorecard.Checks)
	checksJsonValue, err := json.Marshal(checks)
	if err != nil {
		return nil, err
	}

	scorecardQ := createGraphQuery(Scorecard)
	scorecardQ.has[aggregateScore] = scorecard.AggregateScore
	scorecardQ.has[timeScanned] = scorecard.TimeScanned.UTC()
	scorecardQ.has[scorecardVersion] = scorecard.ScorecardVersion
	scorecardQ.has[scorecardCommit] = scorecard.ScorecardCommit
	scorecardQ.has[origin] = scorecard.Origin
	scorecardQ.has[collector] = scorecard.Collector
	scorecardQ.has[checksJson] = string(checksJsonValue)

	edgeQ := createGraphQuery(ScorecardToSource)

	relation := &Relation{
		outV: sourceQ,
		inV:  scorecardQ,
		edge: edgeQ,
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
		Checks:           checks,
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

func (c *gremlinClient) Scorecards(ctx context.Context, certifyScorecardSpec *model.CertifyScorecardSpec) ([]*model.CertifyScorecard, error) {
	// build the query
	g := gremlingo.Traversal_().WithRemote(c.remote)
	fmt.Println("spec", certifyScorecardSpec)

	v := g.V().HasLabel(string(Scorecard))
	if certifyScorecardSpec != nil {
		if certifyScorecardSpec.ID != nil {
			id, err := strconv.ParseInt(*certifyScorecardSpec.ID, 10, 64)
			if err != nil {
				return nil, err
			}
			v = g.V(id).HasLabel(string(Scorecard))
		}
		if certifyScorecardSpec.ScorecardVersion != nil {
			v = v.Has(scorecardVersion, certifyScorecardSpec.ScorecardVersion)
		}
		if certifyScorecardSpec.ScorecardCommit != nil {
			v = v.Has(scorecardVersion, certifyScorecardSpec.ScorecardCommit)
		}
		if certifyScorecardSpec.Collector != nil {
			v = v.Has(collector, certifyScorecardSpec.Collector)
		}
		if certifyScorecardSpec.Origin != nil {
			v = v.Has(origin, certifyScorecardSpec.Origin)
		}
		if certifyScorecardSpec.TimeScanned != nil {
			v = v.Has(timeScanned, certifyScorecardSpec.TimeScanned)
		}
		if certifyScorecardSpec.AggregateScore != nil {
			v = v.Has(aggregateScore, certifyScorecardSpec.AggregateScore)
		}
		if certifyScorecardSpec.Checks != nil && len(certifyScorecardSpec.Checks) > 0 {
			// match checks 1:1
			checksJsonValue, err := json.Marshal(certifyScorecardSpec.Checks)
			if err != nil {
				return nil, err
			}
			v = v.Has(checksJson, string(checksJsonValue))
		}
		v = v.As("scorecard")
		// all scorecards should have at least one source
		v = v.Out().HasLabel(string(Source))
		if certifyScorecardSpec.Source != nil {
			if certifyScorecardSpec.Source.ID != nil {
				id, err := strconv.ParseInt(*certifyScorecardSpec.Source.ID, 10, 64)
				if err != nil {
					return nil, err
				}
				v = v.Out(id).HasLabel(string(Source))
			}
			if certifyScorecardSpec.Source.Name != nil {
				v = v.Has(name, certifyScorecardSpec.Source.Name)
			}
			if certifyScorecardSpec.Source.Type != nil {
				v = v.Has(sourceType, certifyScorecardSpec.Source.Type)
			}
			if certifyScorecardSpec.Source.Namespace != nil {
				v = v.Has(namespace, certifyScorecardSpec.Source.Namespace)
			}
			if certifyScorecardSpec.Source.Commit != nil {
				v = v.Has(commit, certifyScorecardSpec.Source.Commit)
			}
			if certifyScorecardSpec.Source.Tag != nil {
				v = v.Has(tag, certifyScorecardSpec.Source.Tag)
			}
		}
		v = v.As("source")
	}
	v = v.Select("scorecard", "source").Select(gremlingo.Column.Values).Limit(c.config.MaxResultsPerQuery).Unfold().ValueMap(true)

	// execute the query
	results, err := v.ToList()
	if err != nil {
		return nil, err
	}
	fmt.Println("results", results)

	// generate the model objects from the resultset
	var scorecards []*model.CertifyScorecard
	id := ""
	var scorecard *model.CertifyScorecard
	var sources []*model.Source
	for _, result := range results {
		resultMap := result.GetInterface().(map[interface{}]interface{})
		id = strconv.FormatInt(resultMap[string(gremlingo.T.Id)].(int64), 10)
		if resultMap[sourceType] != nil {
			var tagValue string
			if resultMap[tag] != nil {
				tagValue = (resultMap[tag].([]interface{}))[0].(string)
			}
			var commitValue string
			if resultMap[commit] != nil {
				commitValue = (resultMap[commit].([]interface{}))[0].(string)
			}
			source := &model.Source{
				ID:   id,
				Type: (resultMap[sourceType].([]interface{}))[0].(string),
				Namespaces: []*model.SourceNamespace{{
					ID:        id,
					Namespace: (resultMap[namespace].([]interface{}))[0].(string),
					Names: []*model.SourceName{{
						ID:     id,
						Name:   (resultMap[name].([]interface{}))[0].(string),
						Tag:    &tagValue,
						Commit: &commitValue,
					}},
				}},
			}
			sources = append(sources, source)
		}
		if resultMap[checksJson] != nil {
			var checks []*model.ScorecardCheck
			err := json.Unmarshal([]byte(resultMap[checksJson].([]interface{})[0].(string)), &checks)
			if err != nil {
				return nil, err
			}
			scorecard = &model.CertifyScorecard{
				ID: id,
				Scorecard: &model.Scorecard{
					TimeScanned:      (resultMap[timeScanned].([]interface{}))[0].(time.Time),
					AggregateScore:   (resultMap[aggregateScore].([]interface{}))[0].(float64),
					Checks:           checks,
					ScorecardVersion: (resultMap[scorecardVersion].([]interface{}))[0].(string),
					ScorecardCommit:  (resultMap[scorecardCommit].([]interface{}))[0].(string),
					Origin:           (resultMap[origin].([]interface{}))[0].(string),
					Collector:        (resultMap[collector].([]interface{}))[0].(string),
				},
			}
			scorecards = append(scorecards, scorecard)
		}
	}

	for i, scorecard := range scorecards {
		// FIXME: This is not necessarily true... they may not be returned in the same order they were paired
		scorecard.Source = sources[i]
	}

	return scorecards, nil
}

func toChecks(inputCheck []*model.ScorecardCheckInputSpec) []*model.ScorecardCheck {
	var checks []*model.ScorecardCheck
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
