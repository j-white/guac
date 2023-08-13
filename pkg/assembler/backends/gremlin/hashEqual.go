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
	"github.com/guacsec/guac/pkg/assembler/graphql/model"
)

const (
	HashEqual Label = "hashEqual"
)

func getHashEqualObjectFromEdge(result *gremlinQueryResult) *model.HashEqual {
	return &model.HashEqual{
		ID:            result.id,
		Justification: result.edge[justification].(string),
		Origin:        result.edge[collector].(string),
		Collector:     result.edge[origin].(string),
	}
}

func createUpsertForHashEqual(artifact *model.ArtifactInputSpec, equalArtifact *model.ArtifactInputSpec, hashEqual *model.HashEqualInputSpec) *gremlinQueryBuilder[*model.HashEqual] {
	return createUpsertForEdge[*model.HashEqual](HashEqual).
		withPropString(justification, &hashEqual.Justification).
		withPropString(origin, &hashEqual.Origin).
		withPropString(collector, &hashEqual.Collector).
		withOutVertex(createQueryToMatchArtifactInput[*model.HashEqual](artifact)).
		withInVertex(createQueryToMatchArtifactInput[*model.HashEqual](equalArtifact)).
		withMapper(getHashEqualObjectFromEdge)
}

// IngestHashEqual
//
//	artifact -> hashEqual  -> artifact
func (c *gremlinClient) IngestHashEqual(ctx context.Context, artifact model.ArtifactInputSpec, equalArtifact model.ArtifactInputSpec, hashEqual model.HashEqualInputSpec) (*model.HashEqual, error) {
	return createUpsertForHashEqual(&artifact, &equalArtifact, &hashEqual).upsert(c)
}

func (c *gremlinClient) IngestHashEquals(ctx context.Context, artifacts []*model.ArtifactInputSpec, otherArtifacts []*model.ArtifactInputSpec, hashEquals []*model.HashEqualInputSpec) ([]*model.HashEqual, error) {
	// build the queries
	var queries []*gremlinQueryBuilder[*model.HashEqual]
	for k := range artifacts {
		queries = append(queries, createUpsertForHashEqual(artifacts[k], otherArtifacts[k], hashEquals[k]))
	}
	return createBulkUpsertForEdge[*model.HashEqual](HashEqual).
		withQueries(queries).
		upsertBulk(c)
}

func (c *gremlinClient) HashEqual(ctx context.Context, hashEqualSpec *model.HashEqualSpec) ([]*model.HashEqual, error) {
	q := createQueryForEdge[*model.HashEqual](HashEqual).
		withId(hashEqualSpec.ID).
		withPropString(justification, hashEqualSpec.Justification).
		withPropString(origin, hashEqualSpec.Justification).
		withPropString(collector, hashEqualSpec.Justification)
	if len(hashEqualSpec.Artifacts) > 0 {
		// FIXME: More work to do here
		q = q.withInVertex(createQueryToMatchArtifact[*model.HashEqual](hashEqualSpec.Artifacts[0]))
	}
	return q.findAll(c)
}
