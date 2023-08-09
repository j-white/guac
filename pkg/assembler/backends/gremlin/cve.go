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
	"strings"

	"github.com/guacsec/guac/pkg/assembler/graphql/model"
)

const (
	CVE Label = "cve"
)

func getCVEQueryValues(cve *model.CVEInputSpec) *GraphQuery {
	q := createGraphQuery(CVE)
	q.has[year] = int64(cve.Year)
	q.has[cveId] = strings.ToLower(cve.CveID)
	return q
}

func getCVEObject(id string, values map[interface{}]interface{}) *model.Cve {
	return &model.Cve{
		ID:    id,
		Year:  int(values[year].(int64)),
		CveID: values[cveId].(string),
	}
}

func (c *gremlinClient) IngestCve(ctx context.Context, cve *model.CVEInputSpec) (*model.Cve, error) {
	return ingestModelObject[*model.CVEInputSpec, *model.Cve](c, cve, getCVEQueryValues, getCVEObject)
}

func (c *gremlinClient) IngestCVEs(ctx context.Context, cves []*model.CVEInputSpec) ([]*model.Cve, error) {
	return bulkIngestModelObjects[*model.CVEInputSpec, *model.Cve](c, cves, getCVEQueryValues, getCVEObject)
}

func (c *gremlinClient) Cve(ctx context.Context, cveSpec *model.CVESpec) ([]*model.Cve, error) {
	query := createGraphQuery(CVE)
	if cveSpec != nil {
		if cveSpec.ID != nil {
			query.id = *cveSpec.ID
		}
		if cveSpec.Year != nil {
			query.has[year] = *cveSpec.Year
		}
		if cveSpec.CveID != nil {
			query.has[cveId] = strings.ToLower(*cveSpec.CveID)
		}
	}
	return queryModelObjectsFromVertex[*model.Cve](c, query, getCVEObject)
}
