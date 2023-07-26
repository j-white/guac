package tinkerpop

import (
	"context"
	gremlingo "github.com/apache/tinkerpop/gremlin-go/v3/driver"
	"github.com/guacsec/guac/pkg/assembler/graphql/model"
	"strconv"
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

func getGHSAObject(id int64, values map[interface{}]interface{}) *model.Ghsa {
	return &model.Ghsa{
		ID:     strconv.FormatInt(id, 10),
		GhsaID: values[ghsaId].(string),
	}
}

func (c *tinkerpopClient) IngestGhsa(ctx context.Context, ghsa *model.GHSAInputSpec) (*model.Ghsa, error) {
	return ingestModelObject[*model.GHSAInputSpec, *model.Ghsa](c, ghsa, getGHSAQueryValues, getGHSAObject)
}

func (c *tinkerpopClient) IngestGHSAs(ctx context.Context, ghsas []*model.GHSAInputSpec) ([]*model.Ghsa, error) {
	return bulkIngestModelObjects[*model.GHSAInputSpec, *model.Ghsa](c, ghsas, getGHSAQueryValues, getGHSAObject)
}

func (c *tinkerpopClient) Ghsa(ctx context.Context, ghsaSpec *model.GHSASpec) ([]*model.Ghsa, error) {
	var ghsas []*model.Ghsa
	return ghsas, nil
}
