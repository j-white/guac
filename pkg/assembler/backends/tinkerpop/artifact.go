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
	gremlingo "github.com/apache/tinkerpop/gremlin-go/driver"
	"github.com/guacsec/guac/pkg/assembler/graphql/model"
	"strconv"
	"strings"
)

const (
	algorithm string = "algorithm"
	digest    string = "digest"
)

func (c *tinkerpopClient) IngestArtifact(ctx context.Context, artifact *model.ArtifactInputSpec) (*model.Artifact, error) {
	// all fields are required, and canonicalized to lower case
	values := map[string]interface{}{
		algorithm: strings.ToLower(string(artifact.Algorithm)),
		digest:    strings.ToLower(string(artifact.Digest)),
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

func (c *tinkerpopClient) IngestArtifacts(ctx context.Context, artifacts []*model.ArtifactInputSpec) ([]*model.Artifact, error) {
	// FIXME: A good opportunity to use async
	var artifactObjects []*model.Artifact
	for _, artifactSpec := range artifacts {
		artifact, err := c.IngestArtifact(ctx, artifactSpec)
		if err != nil {
			return nil, err
		}
		artifactObjects = append(artifactObjects, artifact)
	}
	return artifactObjects, nil
}

func (c *tinkerpopClient) Artifacts(ctx context.Context, artifactSpec *model.ArtifactSpec) ([]*model.Artifact, error) {
	g := gremlingo.Traversal_().WithRemote(c.remote)
	tx := g.Tx()
	gtx, _ := tx.Begin()

	v := gtx.V()
	// FIXME: Return all artifacts?
	if artifactSpec == nil {
		return nil, nil
	}
	// FIXME: do we support looking up by ID as well?
	if artifactSpec.Algorithm != nil {
		v = v.Has(algorithm, artifactSpec.Algorithm)
	}
	if artifactSpec.Digest != nil {
		v = v.Has(digest, artifactSpec.Digest)
	}

	// FIXME: iterate in batches vs retrieving all the results as a list
	results, err := v.ToList()
	if err == nil {
		return nil, err
	}

	var artifacts []*model.Artifact
	for _, result := range results {
		resultMap := result.GetInterface().(map[interface{}]interface{})
		artifact := &model.Artifact{
			ID:        strconv.FormatInt(resultMap[id].(int64), 10),
			Algorithm: resultMap[algorithm].(string),
			Digest:    resultMap[digest].(string),
		}
		artifacts = append(artifacts, artifact)
	}

	return artifacts, nil
}
