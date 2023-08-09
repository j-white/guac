package gremlin

import (
	"context"
	"github.com/guacsec/guac/pkg/assembler/graphql/model"
	"strings"
)

const (
	OSV Label = "osv"
)

func getOSVQueryValues(osv *model.OSVInputSpec) *GraphQuery {
	q := createGraphQuery(OSV)
	q.has[osvId] = strings.ToLower(osv.OsvID)
	return q
}

func getOSVObject(id string, values map[interface{}]interface{}) *model.Osv {
	return &model.Osv{
		ID:    id,
		OsvID: values[osvId].(string),
	}
}

func (c *gremlinClient) IngestOsv(ctx context.Context, osv *model.OSVInputSpec) (*model.Osv, error) {
	return ingestModelObject[*model.OSVInputSpec, *model.Osv](c, osv, getOSVQueryValues, getOSVObject)
}

func (c *gremlinClient) IngestOSVs(ctx context.Context, osvs []*model.OSVInputSpec) ([]*model.Osv, error) {
	return bulkIngestModelObjects[*model.OSVInputSpec, *model.Osv](c, osvs, getOSVQueryValues, getOSVObject)
}

func (c *gremlinClient) Osv(ctx context.Context, osvSpec *model.OSVSpec) ([]*model.Osv, error) {
	query := createGraphQuery(OSV)
	if osvSpec != nil {
		if osvSpec.ID != nil {
			query.id = *osvSpec.ID
		}
		if osvSpec.OsvID != nil {
			query.has[osvId] = strings.ToLower(*osvSpec.OsvID)
		}
	}
	return queryModelObjectsFromVertex[*model.Osv](c, query, getOSVObject)
}
