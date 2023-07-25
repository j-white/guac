package tinkerpop

import (
	"context"
	gremlingo "github.com/apache/tinkerpop/gremlin-go/v3/driver"
	"github.com/guacsec/guac/pkg/assembler/graphql/model"
	"strconv"
	"strings"
)

const (
	Builder Label = "builder"
)

func getBuilderQueryValues(builder *model.BuilderInputSpec) map[interface{}]interface{} {
	values := make(map[interface{}]interface{})
	values[gremlingo.T.Label] = string(Builder)
	values[uri] = strings.ToLower(builder.URI)
	return values
}

func getBuilderObject(id int64, values map[interface{}]interface{}) *model.Builder {
	return &model.Builder{
		ID:  strconv.FormatInt(id, 10),
		URI: (values[uri].([]interface{}))[0].(string),
	}
}

func (c *tinkerpopClient) IngestBuilder(ctx context.Context, builder *model.BuilderInputSpec) (*model.Builder, error) {
	return ingestModelObject[*model.BuilderInputSpec, *model.Builder](c, builder, getBuilderQueryValues, getBuilderObject)
}

func (c *tinkerpopClient) IngestBuilders(ctx context.Context, builders []*model.BuilderInputSpec) ([]*model.Builder, error) {
	return bulkIngestModelObjects[*model.BuilderInputSpec, *model.Builder](c, builders, getBuilderQueryValues, getBuilderObject)
}
