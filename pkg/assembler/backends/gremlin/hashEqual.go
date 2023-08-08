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
	gremlingo "github.com/apache/tinkerpop/gremlin-go/v3/driver"
	"github.com/guacsec/guac/pkg/assembler/graphql/model"
)

const (
	HashEqual Label = "hashEqual"
)

func getHashEqualQueryValues(artifact *model.ArtifactInputSpec, equalArtifact *model.ArtifactInputSpec, hashEqual *model.HashEqualInputSpec) map[interface{}]interface{} {
	values := make(map[interface{}]interface{})
	values[gremlingo.T.Label] = string(HashEqual)
	values[justification] = hashEqual.Justification
	values[origin] = hashEqual.Origin
	values[collector] = hashEqual.Collector
	return values
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

// IngestHashEqual
//
//	artifact ->hashEqualSubjectArtEdges-> hashEqual  ->hashEqualArtEdges-> artifact
func (c *gremlinClient) IngestHashEqual(ctx context.Context, artifact model.ArtifactInputSpec, equalArtifact model.ArtifactInputSpec, hashEqual model.HashEqualInputSpec) (*model.HashEqual, error) {
	return ingestModelObjectsWithRelation[*model.ArtifactInputSpec, *model.HashEqualInputSpec, *model.HashEqual](
		c, &artifact, &equalArtifact, &hashEqual, getArtifactQueryValues, getHashEqualQueryValues, getHashEqualObject)
}

func (c *gremlinClient) IngestHashEquals(ctx context.Context, artifacts []*model.ArtifactInputSpec, otherArtifacts []*model.ArtifactInputSpec, hashEquals []*model.HashEqualInputSpec) ([]*model.HashEqual, error) {
	return bulkIngestModelObjectsWithRelation[*model.ArtifactInputSpec, *model.HashEqualInputSpec, *model.HashEqual](
		c, artifacts, otherArtifacts, hashEquals, getArtifactQueryValues, getHashEqualQueryValues, getHashEqualObject)
}

func (c *gremlinClient) HashEqual(ctx context.Context, hashEqualSpec *model.HashEqualSpec) ([]*model.HashEqual, error) {
	query := createVertexQuery(HashEqual)
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
	return queryModelObjects[*model.HashEqual](c, query, getHashEqualObject)
}
