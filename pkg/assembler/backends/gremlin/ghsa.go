package gremlin

import (
	"context"
	gremlingo "github.com/apache/tinkerpop/gremlin-go/v3/driver"
	"github.com/guacsec/guac/pkg/assembler/graphql/model"
	"strings"
)

const (
	GHSA Label = "ghsa"
)

func getGHSAQueryValues(ghsa *model.GHSAInputSpec) map[interface{}]interface{} {
	values := make(map[interface{}]interface{})
	values[gremlingo.T.Label] = string(GHSA)
	values[ghsaId] = strings.ToLower(ghsa.GhsaID)
	return values
}

func getGHSAObject(id string, values map[interface{}]interface{}) *model.Ghsa {
	return &model.Ghsa{
		ID:     id,
		GhsaID: values[ghsaId].(string),
	}
}

func (c *gremlinClient) IngestGhsa(ctx context.Context, ghsa *model.GHSAInputSpec) (*model.Ghsa, error) {
	return ingestModelObject[*model.GHSAInputSpec, *model.Ghsa](c, ghsa, getGHSAQueryValues, getGHSAObject)
}

func (c *gremlinClient) IngestGHSAs(ctx context.Context, ghsas []*model.GHSAInputSpec) ([]*model.Ghsa, error) {
	return bulkIngestModelObjects[*model.GHSAInputSpec, *model.Ghsa](c, ghsas, getGHSAQueryValues, getGHSAObject)
}

func (c *gremlinClient) Ghsa(ctx context.Context, ghsaSpec *model.GHSASpec) ([]*model.Ghsa, error) {
	var ghsas []*model.Ghsa
	return ghsas, nil
}
