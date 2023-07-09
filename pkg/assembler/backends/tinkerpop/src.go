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
	"github.com/vektah/gqlparser/v2/gqlerror"
)

func (c *tinkerpopClient) IngestSource(ctx context.Context, source model.SourceInputSpec) (*model.Source, error) {
	// FIXME: We should be using an upsert pattern
	// FIXME: How do we differentiate vertices of different "types"
	//var __ = gremlingo.T__
	//	v := g.MergeV(__.Id(source.Name).Option()).Property("sourceType", source.Type).Property("namespace", source.Namespace)
	// FIXME: Make sure we close
	g := gremlingo.Traversal_().WithRemote(c.remote)
	tx := g.Tx()
	gtx, _ := tx.Begin()

	v := gtx.AddV(source.Name).Property("sourceType", source.Type).Property("namespace", source.Namespace)

	if source.Commit != nil && source.Tag != nil {
		if *source.Commit != "" && *source.Tag != "" {
			return nil, gqlerror.Errorf("Passing both commit and tag selectors is an error")
		}
	}

	if source.Commit != nil {
		v = v.Property("commit", *source.Commit)
	} else {
		v = v.Property("commit", "")
	}

	if source.Tag != nil {
		v = v.Property("tag", *source.Tag)
	} else {
		v = v.Property("tag", "")
	}

	r, err := v.ElementMap().Next()
	if err != nil {
		return nil, err
	}
	resultMap := r.GetInterface().(map[interface{}]interface{})

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	// Generate the model.Source in the database after process the model.SourceInputSpec
	return generateModelSource(resultMap["sourceType"].(string),
		resultMap["namespace"].(string),
		resultMap["label"].(string),
		resultMap["commit"].(string),
		resultMap["tag"].(string)), nil
}

func generateModelSource(srcType, namespaceStr, nameStr string, commitValue, tagValue interface{}) *model.Source {
	tag := (*string)(nil)
	if tagValue != nil {
		tagStr := tagValue.(string)
		tag = &tagStr
	}
	commit := (*string)(nil)
	if commitValue != nil {
		commitStr := commitValue.(string)
		commit = &commitStr
	}
	name := &model.SourceName{
		Name:   nameStr,
		Tag:    tag,
		Commit: commit,
	}

	namespace := &model.SourceNamespace{
		Namespace: namespaceStr,
		Names:     []*model.SourceName{name},
	}

	src := model.Source{
		Type:       srcType,
		Namespaces: []*model.SourceNamespace{namespace},
	}
	return &src
}
