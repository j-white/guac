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
	"strconv"
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
)

const (
	Source    Label = "source"
	Scorecard Label = "scorecard"
)

func validateSourceInputSpec(source model.SourceInputSpec) error {
	if source.Commit != nil && source.Tag != nil {
		if *source.Commit != "" && *source.Tag != "" {
			return gqlerror.Errorf("Passing both commit and tag selectors is an error")
		}
	}
	return nil
}

// CertifyScorecard used to ingest scorecards
func (c *tinkerpopClient) CertifyScorecard(ctx context.Context, source model.SourceInputSpec, scorecard model.ScorecardInputSpec) (*model.CertifyScorecard, error) {
	// TODO: Can we push this validation up a layer, so that the storage engines don't need to worry about it?
	err := validateSourceInputSpec(source)
	if err != nil {
		return nil, err
	}

	// map to vertices and edges
	sourceProperties := map[interface{}]interface{}{
		gremlingo.T.Label: string(Source),
		name:              source.Name,
		sourceType:        source.Type,
		namespace:         source.Namespace,
		tag:               source.Tag,
		commit:            source.Commit,
	}

	checks := toChecks(scorecard.Checks)
	scorecardProperties := map[interface{}]interface{}{
		gremlingo.T.Label: string(Scorecard),
		aggregateScore:    scorecard.AggregateScore,
		timeScanned:       scorecard.TimeScanned.UTC(),
		scorecardVersion:  scorecard.ScorecardVersion,
		scorecardCommit:   scorecard.ScorecardCommit,
		origin:            scorecard.Origin,
		collector:         scorecard.Collector,
		checksJson:        json.Marshal(checks),
	}

	edgeProperties := map[interface{}]interface{}{
		gremlingo.T.Label:       string(Scorecard),
		gremlingo.Direction.In:  gremlingo.Merge.InV,
		gremlingo.Direction.Out: gremlingo.Merge.OutV,
	}

	// upsert (source -> scorecard)
	g := gremlingo.Traversal_().WithRemote(c.remote)
	r, err := g.MergeV(sourceProperties).As("source").
		MergeV(scorecardProperties).As("scorecard").
		MergeE(edgeProperties).
		// late bind
		Option(gremlingo.Merge.InV, gremlingo.T__.Select("source")).
		Option(gremlingo.Merge.OutV, gremlingo.T__.Select("scorecard")).
		As("edge").
		Select("scorecard").
		Id().Next()
	if err != nil {
		return nil, err
	}
	id, err := r.GetInt64()
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
		ID:        strconv.FormatInt(id, 10),
		Source:    src,
		Scorecard: &modelScorecard,
	}
	return &certification, nil
}

func (c *tinkerpopClient) Sources(ctx context.Context, sourceSpec *model.SourceSpec) ([]*model.Source, error) {
	// build the query
	g := gremlingo.Traversal_().WithRemote(c.remote)
	v := g.V().HasLabel(string(Source))
	if sourceSpec != nil {
		if sourceSpec.ID != nil {
			id, err := strconv.ParseInt(*sourceSpec.ID, 10, 64)
			if err != nil {
				return nil, err
			}
			v = g.V(id).HasLabel(string(Artifact))
		}
		if sourceSpec.Name != nil {
			v = v.Has(name, *sourceSpec.Name)
		}
		if sourceSpec.Type != nil {
			v = v.Has(sourceType, *sourceSpec.Type)
		}
		if sourceSpec.Namespace != nil {
			v = v.Has(namespace, *sourceSpec.Namespace)
		}
		if sourceSpec.Tag != nil {
			v = v.Has(tag, *sourceSpec.Tag)
		}
		if sourceSpec.Commit != nil {
			v = v.Has(commit, *sourceSpec.Commit)
		}
	}
	v = v.ValueMap(true)

	// execute the query
	results, err := v.Limit(c.config.MaxLimit).ToList()
	if err != nil {
		return nil, err
	}

	// generate the model objects from the resultset
	var sources []*model.Source
	for _, result := range results {
		resultMap := result.GetInterface().(map[interface{}]interface{})
		id := strconv.FormatInt(resultMap[string(gremlingo.T.Id)].(int64), 10)
		tagValue := (resultMap[tag].([]interface{}))[0].(string)
		commitValue := (resultMap[commit].([]interface{}))[0].(string)
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

	return sources, nil
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

func (c *tinkerpopClient) IngestSources(ctx context.Context, sources []*model.SourceInputSpec) ([]*model.Source, error) {
	//TODO implement me
	panic("implement me")
}

func (c *tinkerpopClient) Scorecards(ctx context.Context, certifyScorecardSpec *model.CertifyScorecardSpec) ([]*model.CertifyScorecard, error) {
	/*Select("scorecard").
		ValueMap(true).
		Limit(1).
		Next()
	resultMap := r.GetInterface().(map[interface{}]interface{})
	checks, err := readArrayFromVertexProperties(resultMap)

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
	}*/

	panic("implement me")
}
