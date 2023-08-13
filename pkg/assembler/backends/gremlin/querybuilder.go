package gremlin

import (
	"github.com/guacsec/guac/pkg/assembler/graphql/model"
	"time"
)

type gremlinQueryBuilder[M any] struct {
	query   *GraphQuery
	mapper  func(*gremlinQueryResult) M
	queries []*gremlinQueryBuilder[M]
}

type gremlinQueryResult struct {
	id   string
	out  map[interface{}]interface{}
	edge map[interface{}]interface{}
	in   map[interface{}]interface{}
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

func (gqb *gremlinQueryBuilder[M]) withMapper(mapper func(*gremlinQueryResult) M) *gremlinQueryBuilder[M] {
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

func (gqb *gremlinQueryBuilder[M]) findAll(c *gremlinClient) ([]M, error) {
	return queryModelObjectsFromEdge[M](c, gqb.query, gqb.mapper)
}
