package gremlin

import (
	gremlingo "github.com/apache/tinkerpop/gremlin-go/v3/driver"
	"reflect"
	"strconv"
)

type GraphQuery struct {
	label            Label
	id               string
	partitionKey     string
	has              map[string]interface{}
	outVQuery        *GraphQuery
	inVQuery         *GraphQuery
	orderByKey       string
	orderByDirection interface{}
	isUpsert         bool
}

func createGraphQuery(label Label) *GraphQuery {
	q := &GraphQuery{label: label}
	q.has = make(map[string]interface{})
	q.orderByDirection = gremlingo.Order.Desc
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

func (query *GraphQuery) isEmpty() bool {
	return query.id == "" && len(query.has) < 1
}

func queryModelObjectsFromVertex[M any](c *gremlinClient, query *GraphQuery, deserializer func(*gremlinQueryResult) (M, error)) ([]M, error) {
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

		gremlinResult := gremlinQueryResult{
			vertexId:    vertexId,
			vertexLabel: query.label,
			vertex:      query.toReadMap(),
		}

		object, err := deserializer(&gremlinResult)
		if err != nil {
			return objects, err
		}
		objects = append(objects, object)
	}

	return objects, nil
}

func queryModelObjectsFromEdge[M any](c *gremlinClient, query *GraphQuery, deserializer func(*gremlinQueryResult) (M, error)) ([]M, error) {
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
	// all filters on the edge
	for key, value := range query.has {
		v = v.Has(key, value)
	}
	v.As("edge")
	if query.outVQuery != nil {
		v = v.OutV()
		if query.outVQuery.id != "" {
			v = v.HasId(query.outVQuery.id)
		}
		// always match on label
		v = v.HasLabel(string(query.outVQuery.label))
		// match on partition key if set
		if query.outVQuery.partitionKey != "" {
			v = v.Has(guacPartitionKey, query.outVQuery.partitionKey)
		}
		// all filters on the vertex
		for key, value := range query.outVQuery.has {
			v = v.Has(key, value)
		}
		v = v.As("from")
	}
	v.Select("edge")
	if query.inVQuery != nil {
		v = v.InV()
		if query.inVQuery.id != "" {
			v = v.HasId(query.inVQuery.id)
		}
		// always match on label
		v = v.HasLabel(string(query.inVQuery.label))
		// match on partition key if set
		if query.inVQuery.partitionKey != "" {
			v = v.Has(guacPartitionKey, query.inVQuery.partitionKey)
		}
		// all filters on the vertex
		for key, value := range query.inVQuery.has {
			v = v.Has(key, value)
		}
		v = v.As("to")
	}

	// retrieve all values
	v = v.Select("edge")
	if query.orderByKey != "" {
		v = v.Order().By(query.orderByKey, query.orderByDirection)
	}
	v = v.Project("from", "edge", "to").
		By(gremlingo.T__.OutV().ValueMap(true)).
		By(gremlingo.T__.ValueMap(true)).
		By(gremlingo.T__.InV().ValueMap(true))

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
		var outId string
		var inId string
		if c.config.Flavor == JanusGraph {
			relationId := edgeMap[string(gremlingo.T.Id)].(*janusgraphRelationIdentifier)
			edgeId = strconv.FormatInt(relationId.RelationId, 10)
			outId = strconv.FormatInt(fromMap[string(gremlingo.T.Id)].(int64), 10)
			inId = strconv.FormatInt(toMap[string(gremlingo.T.Id)].(int64), 10)
		} else {
			edgeId = resultMap[string(gremlingo.T.Id)].(string)
			outId = fromMap[string(gremlingo.T.Id)].(string)
			inId = toMap[string(gremlingo.T.Id)].(string)
		}

		var outLabel Label
		if query.outVQuery != nil {
			outLabel = query.outVQuery.label
		} else {
			outLabel = Label(fromMap[string(gremlingo.T.Label)].(string))
		}

		var inLabel Label
		if query.inVQuery != nil {
			inLabel = query.inVQuery.label
		} else {
			inLabel = Label(toMap[string(gremlingo.T.Label)].(string))
		}

		gResult := &gremlinQueryResult{
			id:       edgeId,
			out:      fromMap,
			outLabel: outLabel,
			outId:    outId,
			edge:     edgeMap,
			in:       toMap,
			inLabel:  inLabel,
			inId:     inId,
			query:    query,
		}

		object, err := deserializer(gResult)
		if err != nil {
			return objects, err
		}
		objects = append(objects, object)
	}

	return objects, nil
}

// flattenResultMap
//
//	in responses, values are in arrays, even single values like so:
//	namespace: ["somenamespace"]
//	we convert this to:
//	namespace: "somenamespace"
//
// for single values, and keep arrays otherwise
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
