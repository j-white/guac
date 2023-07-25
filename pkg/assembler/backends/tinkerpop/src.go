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
)

func getSourceQueryValues(source *model.SourceInputSpec) map[interface{}]interface{} {
	values := make(map[interface{}]interface{})
	values[gremlingo.T.Label] = string(Source)
	values[typeStr] = source.Type
	values[namespace] = source.Namespace

	if source.Commit != nil {
		values[commit] = *source.Commit
	} else {
		values[commit] = ""
	}

	if source.Tag != nil {
		values[tag] = *source.Tag
	} else {
		values[tag] = ""
	}

	return values
}

func getSourceObject(id int64, values map[interface{}]interface{}) *model.Source {
	return generateModelSource((values[typeStr].([]interface{}))[0].(string),
		(values["namespace"].([]interface{}))[0].(string),
		(values["label"].([]interface{}))[0].(string),
		(values["commit"].([]interface{}))[0].(string),
		(values["tag"].([]interface{}))[0].(string))
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

func (c *tinkerpopClient) IngestSource(ctx context.Context, source model.SourceInputSpec) (*model.Source, error) {
	return ingestModelObject[*model.SourceInputSpec, *model.Source](c, &source, getSourceQueryValues, getSourceObject)
}

func (c *tinkerpopClient) IngestSources(ctx context.Context, sources []*model.SourceInputSpec) ([]*model.Source, error) {
	return bulkIngestModelObjects[*model.SourceInputSpec, *model.Source](c, sources, getSourceQueryValues, getSourceObject)
}
