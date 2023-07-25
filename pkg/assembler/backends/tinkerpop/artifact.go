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
	gremlingo "github.com/apache/tinkerpop/gremlin-go/v3/driver"
	"github.com/guacsec/guac/pkg/assembler/graphql/model"
	"strconv"
	"strings"
)

const (
	Artifact Label = "artifact"
)

func (c *tinkerpopClient) IngestArtifact(ctx context.Context, artifact *model.ArtifactInputSpec) (*model.Artifact, error) {
	// all fields are required, and canonicalized to lower case
	values := map[interface{}]interface{}{
		gremlingo.T.Label: string(Artifact),
		algorithm:         strings.ToLower(artifact.Algorithm),
		digest:            strings.ToLower(artifact.Digest),
	}

	id, err := c.upsertVertex(values)
	if err != nil {
		return nil, err
	}

	// build artifact from canonical model after a successful upsert
	a := &model.Artifact{
		ID:        strconv.FormatInt(id, 10),
		Algorithm: values[algorithm].(string),
		Digest:    values[digest].(string),
	}

	return a, nil
}

// IngestArtifacts iterate through the list one by one on a single thread, abort on any failure and return those inserted so far
func (c *tinkerpopClient) IngestArtifacts(ctx context.Context, artifacts []*model.ArtifactInputSpec) ([]*model.Artifact, error) {
	// FIXME: Implement bulk insert
	var artifactObjects []*model.Artifact
	for _, artifactSpec := range artifacts {
		artifact, err := c.IngestArtifact(ctx, artifactSpec)
		if err != nil {
			return artifactObjects, err
		}
		artifactObjects = append(artifactObjects, artifact)
	}
	return artifactObjects, nil
}

func (c *tinkerpopClient) Artifacts(ctx context.Context, artifactSpec *model.ArtifactSpec) ([]*model.Artifact, error) {
	// build the query
	g := gremlingo.Traversal_().WithRemote(c.remote)
	v := g.V().HasLabel(string(Artifact))
	if artifactSpec != nil {
		if artifactSpec.ID != nil {
			id, err := strconv.ParseInt(*artifactSpec.ID, 10, 64)
			if err != nil {
				return nil, err
			}
			v = g.V(id).HasLabel(string(Artifact))
		}
		if artifactSpec.Algorithm != nil {
			v = v.Has(algorithm, strings.ToLower(*artifactSpec.Algorithm))
		}
		if artifactSpec.Digest != nil {
			v = v.Has(digest, strings.ToLower(*artifactSpec.Digest))
		}
	}
	v = v.ValueMap(true)

	// execute the query
	results, err := v.Limit(c.config.MaxLimit).ToList()
	if err != nil {
		return nil, err
	}

	// generate the model objects from the resultset
	var artifacts []*model.Artifact
	for _, result := range results {
		resultMap := result.GetInterface().(map[interface{}]interface{})
		artifact := &model.Artifact{
			ID:        strconv.FormatInt(resultMap[string(gremlingo.T.Id)].(int64), 10),
			Algorithm: (resultMap[algorithm].([]interface{}))[0].(string),
			Digest:    (resultMap[digest].([]interface{}))[0].(string),
		}
		artifacts = append(artifacts, artifact)
	}

	return artifacts, nil
}
