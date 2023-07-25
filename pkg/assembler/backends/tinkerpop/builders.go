package tinkerpop

import (
	"context"
	"github.com/guacsec/guac/pkg/assembler/graphql/model"
	"strings"
)

// copied from arrango
func getBuilderQueryValues(builder *model.BuilderInputSpec) map[interface{}]interface{} {
	values := make(map[interface{}]interface{})
	values["uri"] = strings.ToLower(builder.URI)
	return values
}

func (c *tinkerpopClient) IngestBuilder(ctx context.Context, builder *model.BuilderInputSpec) (*model.Builder, error) {
	return ingestModelObject[*model.BuilderInputSpec, *model.Builder](c, builder, getBuilderQueryValues, nil)
}

func (c *tinkerpopClient) IngestBuilders(ctx context.Context, builders []*model.BuilderInputSpec) ([]*model.Builder, error) {
	return bulkIngestModelObjects[*model.BuilderInputSpec, *model.Builder](c, builders, getBuilderQueryValues, nil)
}

type MapSerializer[M any] func(model M) (result map[interface{}]interface{})

type ObjectDeserializer[M any] func(id int64, values map[interface{}]interface{}) (model M)

func ingestModelObject[C any, D any](c *tinkerpopClient, modelObject C, serializer MapSerializer[C], deserializer ObjectDeserializer[D]) (D, error) {
	var object D
	values := serializer(modelObject)
	id, err := c.upsertVertex(values)
	if err != nil {
		return object, err
	}
	object = deserializer(id, values)
	return object, nil
}

func bulkIngestModelObjects[C any, D any](c *tinkerpopClient, modelObjects []C, serializer MapSerializer[C], deserializer ObjectDeserializer[D]) ([]D, error) {
	var objects []D
	for _, modelObject := range modelObjects {
		object, err := ingestModelObject(c, modelObject, serializer, deserializer)
		if err != nil {
			return objects, err
		}
		objects = append(objects, object)
	}
	return objects, nil
}
