package gremlin

import (
	"context"
	"github.com/guacsec/guac/pkg/assembler/graphql/model"
)

const (
	Builder Label = "builder"
)

func getBuilderQueryValues(builder *model.BuilderInputSpec) *GraphQuery {
	q := createGraphQuery(Builder)
	q.has[uri] = builder.URI
	return q
}

func getBuilderObject(result *gremlinQueryResult) (*model.Builder, error) {
	return &model.Builder{
		ID:  result.vertexId,
		URI: result.vertex[uri].(string),
	}, nil
}

func (c *gremlinClient) IngestBuilder(ctx context.Context, builder *model.BuilderInputSpec) (*model.Builder, error) {
	return ingestModelObject[*model.BuilderInputSpec, *model.Builder](c, builder, getBuilderQueryValues, getBuilderObject)
}

func (c *gremlinClient) IngestBuilders(ctx context.Context, builders []*model.BuilderInputSpec) ([]*model.Builder, error) {
	return bulkIngestModelObjects[*model.BuilderInputSpec, *model.Builder](c, builders, getBuilderQueryValues, getBuilderObject)
}

func (c *gremlinClient) Builders(ctx context.Context, builderSpec *model.BuilderSpec) ([]*model.Builder, error) {
	query := createGraphQuery(Builder)
	if builderSpec != nil {
		if builderSpec.ID != nil {
			query.id = *builderSpec.ID
		}
		if builderSpec.URI != nil {
			query.has[uri] = *builderSpec.URI
		}
	}
	return queryModelObjectsFromVertex[*model.Builder](c, query, getBuilderObject)
}
