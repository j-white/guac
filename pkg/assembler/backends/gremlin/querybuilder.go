package gremlin

import (
	gremlingo "github.com/apache/tinkerpop/gremlin-go/v3/driver"
	"github.com/guacsec/guac/pkg/assembler/graphql/model"
	"strings"
	"time"
)

type gremlinQueryBuilder[M any] struct {
	query   *GraphQuery
	mapper  func(*gremlinQueryResult) (M, error)
	queries []*gremlinQueryBuilder[M]
}

type gremlinQueryResult struct {
	vertex      map[interface{}]interface{}
	vertexId    string
	vertexLabel Label

	out      map[interface{}]interface{}
	outId    string
	outLabel Label

	edge map[interface{}]interface{}
	// FIXME: Refactor to edge id
	id        string
	edgeLabel Label

	in      map[interface{}]interface{}
	inId    string
	inLabel Label

	// the source query
	query *GraphQuery
}

func createQueryForVertex[M any](label Label) *gremlinQueryBuilder[M] {
	return &gremlinQueryBuilder[M]{
		query: createGraphQuery(label),
	}
}

func createQueryForEdge[M any](label Label) *gremlinQueryBuilder[M] {
	return &gremlinQueryBuilder[M]{
		query: createGraphQuery(label),
	}
}

func createUpsertForEdge[M any](label Label) *gremlinQueryBuilder[M] {
	return &gremlinQueryBuilder[M]{
		query: createGraphQuery(label),
	}
}

func createUpsertForVertex[M any](label Label) *gremlinQueryBuilder[M] {
	gqb := &gremlinQueryBuilder[M]{
		query: createGraphQuery(label),
	}
	gqb.query.isUpsert = true
	return gqb
}

func createBulkUpsertForVertex[M any](label Label) *gremlinQueryBuilder[M] {
	return &gremlinQueryBuilder[M]{
		query: createGraphQuery(label),
	}
}

func createBulkUpsertForEdge[M any](label Label) *gremlinQueryBuilder[M] {
	return &gremlinQueryBuilder[M]{
		query: createGraphQuery(label),
	}
}

func (gqb *gremlinQueryBuilder[M]) withId(id *string) *gremlinQueryBuilder[M] {
	if id != nil {
		gqb.query.id = *id
	}
	return gqb
}

func (gqb *gremlinQueryBuilder[M]) withPropString(key string, value *string) *gremlinQueryBuilder[M] {
	if value != nil {
		gqb.query.has[key] = *value
	}
	return gqb
}

func (gqb *gremlinQueryBuilder[M]) withPropStringOrEmpty(key string, value *string) *gremlinQueryBuilder[M] {
	if value != nil {
		gqb.query.has[key] = *value
	} else {
		gqb.query.has[key] = ""
	}
	return gqb
}

func (gqb *gremlinQueryBuilder[M]) withPropStringToLower(key string, value *string) *gremlinQueryBuilder[M] {
	if value != nil {
		gqb.query.has[key] = strings.ToLower(*value)
	} else {
		gqb.query.has[key] = ""
	}
	return gqb
}

func (gqb *gremlinQueryBuilder[M]) withPropDependencyType(key string, value *model.DependencyType) *gremlinQueryBuilder[M] {
	if value != nil {
		gqb.query.has[key] = value.String()
	}
	return gqb
}

func (gqb *gremlinQueryBuilder[M]) withPropTime(key string, value *time.Time) *gremlinQueryBuilder[M] {
	if value != nil {
		gqb.query.has[key] = *value
	}
	return gqb
}

func (gqb *gremlinQueryBuilder[M]) withPropTimeGreaterOrEqual(key string, value *time.Time) *gremlinQueryBuilder[M] {
	if value != nil {
		gqb.query.has[key] = gremlingo.P.Gte(*value)
	}
	return gqb
}

func (gqb *gremlinQueryBuilder[M]) withPropFloat64(key string, value *float64) *gremlinQueryBuilder[M] {
	if value != nil {
		gqb.query.has[key] = *value
	}
	return gqb
}

func (gqb *gremlinQueryBuilder[M]) withInVertex(q *gremlinQueryBuilder[M]) *gremlinQueryBuilder[M] {
	gqb.query.inVQuery = q.query
	return gqb
}

func (gqb *gremlinQueryBuilder[M]) withOutVertex(q *gremlinQueryBuilder[M]) *gremlinQueryBuilder[M] {
	gqb.query.outVQuery = q.query
	return gqb
}

func (gqb *gremlinQueryBuilder[M]) withMapper(mapper func(*gremlinQueryResult) (M, error)) *gremlinQueryBuilder[M] {
	gqb.mapper = mapper
	return gqb
}

func (gqb *gremlinQueryBuilder[M]) withQueries(queries []*gremlinQueryBuilder[M]) *gremlinQueryBuilder[M] {
	gqb.queries = queries
	return gqb
}

func (gqb *gremlinQueryBuilder[M]) withOrderByKey(orderBy string) *gremlinQueryBuilder[M] {
	gqb.query.orderByKey = orderBy
	return gqb
}

func (gqb *gremlinQueryBuilder[M]) withOrderByDirection(direction interface{}) *gremlinQueryBuilder[M] {
	gqb.query.orderByDirection = direction
	return gqb
}

// terminal steps
func (gqb *gremlinQueryBuilder[M]) isEmpty() bool {
	return gqb.query.isEmpty()
}

func (gqb *gremlinQueryBuilder[M]) upsert(c *gremlinClient) (M, error) {
	return ingestModelObjectsWithRelation[M](c, gqb)
}

func (gqb *gremlinQueryBuilder[M]) upsertBulk(c *gremlinClient) ([]M, error) {
	return bulkIngestModelObjectsWithRelation[M](c, gqb)
}

func (gqb *gremlinQueryBuilder[M]) findAllEdges(c *gremlinClient) ([]M, error) {
	return queryModelObjectsFromEdge[M](c, gqb.query, gqb.mapper)
}

func (gqb *gremlinQueryBuilder[M]) findAllVertices(c *gremlinClient) ([]M, error) {
	return queryModelObjectsFromVertex[M](c, gqb.query, gqb.mapper)
}
