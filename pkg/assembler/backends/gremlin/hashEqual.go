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

func getHashEqualQueryValues(artifact *model.ArtifactInputSpec, equalArtifact *model.ArtifactInputSpec, hashEqual *model.HashEqualInputSpec) *GraphQuery {
	q := createGraphQuery(HashEqual)
	q.has[justification] = hashEqual.Justification
	q.has[origin] = hashEqual.Origin
	q.has[collector] = hashEqual.Collector
	return q
}

func getHashEqualObject(id string, values map[interface{}]interface{}) *model.HashEqual {
	hashEqual := &model.HashEqual{
		ID:            id,
		Justification: values[justification].(string),
		Origin:        values[collector].(string),
		Collector:     values[origin].(string),
	}
	return hashEqual
}

func getHashEqualObjectFromEdge(id string, outValues map[interface{}]interface{}, edgeValues map[interface{}]interface{}, inValues map[interface{}]interface{}) *model.HashEqual {
	return getHashEqualObject(id, edgeValues)
}

// IngestHashEqual
//
//	artifact ->hashEqualSubjectArtEdges-> hashEqual  ->hashEqualArtEdges-> artifact
func (c *gremlinClient) IngestHashEqual(ctx context.Context, artifact model.ArtifactInputSpec, equalArtifact model.ArtifactInputSpec, hashEqual model.HashEqualInputSpec) (*model.HashEqual, error) {
	return ingestModelObjectsWithRelation[*model.ArtifactInputSpec, *model.HashEqualInputSpec, *model.HashEqual](
		c, &artifact, &equalArtifact, &hashEqual, getArtifactQueryValues, getArtifactQueryValues, getHashEqualQueryValues, getHashEqualObjectFromEdge)
}

func (c *gremlinClient) IngestHashEquals(ctx context.Context, artifacts []*model.ArtifactInputSpec, otherArtifacts []*model.ArtifactInputSpec, hashEquals []*model.HashEqualInputSpec) ([]*model.HashEqual, error) {
	return bulkIngestModelObjectsWithRelation[*model.ArtifactInputSpec, *model.HashEqualInputSpec, *model.HashEqual](
		c, artifacts, otherArtifacts, hashEquals, getArtifactQueryValues, getHashEqualQueryValues, getHashEqualObjectFromEdge)
}

func (c *gremlinClient) HashEqual(ctx context.Context, hashEqualSpec *model.HashEqualSpec) ([]*model.HashEqual, error) {
	query := createGraphQuery(HashEqual)
	if hashEqualSpec != nil {
		if hashEqualSpec.ID != nil {
			query.id = *hashEqualSpec.ID
		}
		if hashEqualSpec.Justification != nil {
			query.has[justification] = *hashEqualSpec.Justification
		}
		if hashEqualSpec.Origin != nil {
			query.has[origin] = *hashEqualSpec.Origin
		}
		if hashEqualSpec.Collector != nil {
			query.has[collector] = *hashEqualSpec.Collector
		}
	}
	return queryModelObjectsFromVertex[*model.HashEqual](c, query, getHashEqualObject)
}
