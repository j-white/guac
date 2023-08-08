package gremlin

import (
	gremlingo "github.com/apache/tinkerpop/gremlin-go/v3/driver"
	"reflect"
	"strconv"
)

type VertexQuery struct {
	label        Label
	id           string
	partitionKey string
	has          map[string]interface{}
}

func createVertexQuery(label Label) *VertexQuery {
	q := &VertexQuery{label: label}
	q.has = make(map[string]interface{})
	return q
}

func queryModelObjects[M any](c *gremlinClient, query *VertexQuery, deserializer ObjectDeserializer[M]) ([]M, error) {
	// build the query
	g := gremlingo.Traversal_().WithRemote(c.remote)
	var v *gremlingo.GraphTraversal
	if query.id != "" {
		// if we have the id, use it at the start of the query instead of matching later
		v = g.V(query.id)
	} else {
		v = g.V()
	}
	// always match on label
	v = v.HasLabel(string(query.label))
	// match on partition key if set
	if query.partitionKey != "" {
		v = v.Has(guacPartitionKey, query.partitionKey)
	}
	// all filters
	for key, value := range query.has {
		v = v.Has(key, value)
	}
	// retrieve all values
	v = v.ValueMap(true)

	// execute the query (blocking)
	results, err := v.Limit(c.config.MaxResultsPerQuery).ToList()
	if err != nil {
		return nil, err
	}

	// generate the model objects from the resultset
	var objects []M
	for _, result := range results {
		resultMap := flattenResultMap(result.GetInterface().(map[interface{}]interface{}))

		var vertexId string
		if c.config.Flavor == JanusGraph {
			vertexId = strconv.FormatInt(resultMap[string(gremlingo.T.Id)].(int64), 10)
		} else {
			vertexId = resultMap[string(gremlingo.T.Id)].(string)
		}

		object := deserializer(vertexId, resultMap)
		objects = append(objects, object)
	}

	return objects, nil
}

/*
*
in responses, values are in arrays, even single values like so:

	namespace: ["somenamespace"]

we convert this to:

	namespace: "somenamespace"

for single values, and keep arrays otherwise
*/
func flattenResultMap(resultMap map[interface{}]interface{}) map[interface{}]interface{} {
	newResultMap := make(map[interface{}]interface{})
	for k, v := range resultMap {
		flattenedValue := v
		rt := reflect.TypeOf(v)
		switch rt.Kind() {
		case reflect.Slice:
			fallthrough
		case reflect.Array:
			valueAsArray := v.([]interface{})
			if len(valueAsArray) == 1 {
				flattenedValue = valueAsArray[0]
			}
		}
		newResultMap[k] = flattenedValue
	}
	return newResultMap
}
