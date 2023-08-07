package tinkerpop

import (
	"context"
	gremlingo "github.com/apache/tinkerpop/gremlin-go/v3/driver"
	"github.com/guacsec/guac/pkg/assembler/graphql/model"
	"strings"
)

const (
	OSV Label = "osv"
)

func getOSVQueryValues(osv *model.OSVInputSpec) map[interface{}]interface{} {
	values := make(map[interface{}]interface{})
	values[gremlingo.T.Label] = string(OSV)
	values[osvId] = strings.ToLower(osv.OsvID)
	return values
}

func getOSVObject(id string, values map[interface{}]interface{}) *model.Osv {
	return &model.Osv{
		ID:    id,
		OsvID: values[osvId].(string),
	}
}

func (c *tinkerpopClient) IngestOsv(ctx context.Context, osv *model.OSVInputSpec) (*model.Osv, error) {
	return ingestModelObject[*model.OSVInputSpec, *model.Osv](c, osv, getOSVQueryValues, getOSVObject)
}

func (c *tinkerpopClient) IngestOSVs(ctx context.Context, osvs []*model.OSVInputSpec) ([]*model.Osv, error) {
	return bulkIngestModelObjects[*model.OSVInputSpec, *model.Osv](c, osvs, getOSVQueryValues, getOSVObject)
}

func (c *tinkerpopClient) Osv(ctx context.Context, osvSpec *model.OSVSpec) ([]*model.Osv, error) {
	var osvc []*model.Osv
	return osvc, nil
}
