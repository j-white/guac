package gremlin

import (
	gremlingo "github.com/apache/tinkerpop/gremlin-go/v3/driver"
	"reflect"
	"strconv"
)

type GraphQuery struct {
	label        Label
	id           string
	partitionKey string
	has          map[string]interface{}
}

func createGraphQuery(label Label) *GraphQuery {
	q := &GraphQuery{label: label}
	q.has = make(map[string]interface{})
	return q
}

func (query *GraphQuery) toVertexMap() map[interface{}]interface{} {
	values := make(map[interface{}]interface{})
	for k, v := range query.has {
		values[k] = v
	}
	values[gremlingo.T.Label] = string(query.label)
	return values
}

func (query *GraphQuery) toEdgeMap() map[interface{}]interface{} {
	values := make(map[interface{}]interface{})
	for k, v := range query.has {
		values[k] = v
	}
	values[gremlingo.T.Label] = string(query.label)
	values[gremlingo.Direction.In] = gremlingo.Merge.InV
	values[gremlingo.Direction.Out] = gremlingo.Merge.OutV
	return values
}

func (query *GraphQuery) toReadMap() map[interface{}]interface{} {
	return query.toVertexMap()
}

func queryModelObjectsFromVertex[M any](c *gremlinClient, query *GraphQuery, deserializer ObjectDeserializer[M]) ([]M, error) {
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

func queryModelObjectsFromEdge[M any](c *gremlinClient, query *GraphQuery, deserializer EdgeObjectDeserializer[M]) ([]M, error) {
	// build the query
	g := gremlingo.Traversal_().WithRemote(c.remote)
	var v *gremlingo.GraphTraversal
	if query.id != "" {
		// if we have the id, use it at the start of the query instead of matching later
		v = g.E(query.id)
	} else {
		v = g.E()
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
	v = v.Project("from", "edge", "to").By(gremlingo.T__.OutV().ValueMap(true)).By(gremlingo.T__.ValueMap(true)).By(gremlingo.T__.InV().ValueMap(true))

	// execute the query (blocking)
	results, err := v.Limit(c.config.MaxResultsPerQuery).ToList()
	if err != nil {
		return nil, err
	}

	// generate the model objects from the resultset
	var objects []M
	for _, result := range results {
		resultMap := result.GetInterface().(map[interface{}]interface{})
		edgeMap := flattenResultMap(resultMap["edge"].(map[interface{}]interface{}))
		fromMap := flattenResultMap(resultMap["from"].(map[interface{}]interface{}))
		toMap := flattenResultMap(resultMap["to"].(map[interface{}]interface{}))

		var edgeId string
		if c.config.Flavor == JanusGraph {
			relationId := edgeMap[string(gremlingo.T.Id)].(*janusgraphRelationIdentifier)
			edgeId = strconv.FormatInt(relationId.RelationId, 10)
		} else {
			edgeId = resultMap[string(gremlingo.T.Id)].(string)
		}

		object := deserializer(edgeId, fromMap, edgeMap, toMap)
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
