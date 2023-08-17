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
	"strings"
)

const (
	Artifact Label = "artifact"
)

func getArtifactQueryValues(artifact *model.ArtifactInputSpec) *GraphQuery {
	q := createGraphQuery(Artifact)
	q.has[algorithm] = strings.ToLower(artifact.Algorithm)
	q.has[digest] = strings.ToLower(artifact.Digest)
	return q
}

func getArtifactObject(result *gremlinQueryResult) (*model.Artifact, error) {
	return &model.Artifact{
		ID:        result.vertexId,
		Algorithm: result.vertex[algorithm].(string),
		Digest:    result.vertex[digest].(string),
	}, nil
}

func (c *gremlinClient) IngestArtifact(ctx context.Context, artifact *model.ArtifactInputSpec) (*model.Artifact, error) {
	return ingestModelObject[*model.ArtifactInputSpec, *model.Artifact](c, artifact, getArtifactQueryValues, getArtifactObject)
}

func (c *gremlinClient) IngestArtifacts(ctx context.Context, artifacts []*model.ArtifactInputSpec) ([]*model.Artifact, error) {
	return bulkIngestModelObjects[*model.ArtifactInputSpec, *model.Artifact](c, artifacts, getArtifactQueryValues, getArtifactObject)
}

func (c *gremlinClient) Artifacts(ctx context.Context, artifactSpec *model.ArtifactSpec) ([]*model.Artifact, error) {
	query := createGraphQuery(Artifact)
	if artifactSpec != nil {
		if artifactSpec.ID != nil {
			query.id = *artifactSpec.ID
		}
		if artifactSpec.Algorithm != nil {
			query.has[algorithm] = strings.ToLower(*artifactSpec.Algorithm)
		}
		if artifactSpec.Digest != nil {
			query.has[digest] = strings.ToLower(*artifactSpec.Digest)
		}
	}
	return queryModelObjectsFromVertex[*model.Artifact](c, query, getArtifactObject)
}
