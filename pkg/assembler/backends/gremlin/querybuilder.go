package gremlin

import (
	"github.com/guacsec/guac/pkg/assembler/graphql/model"
	"time"
)

type gremlinQueryBuilder[M any] struct {
	query *GraphQuery
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

func createBulkUpsertForEdge[M any](label Label) *gremlinQueryBuilder[M] {
	return &gremlinQueryBuilder[M]{
		query: createGraphQuery(label),
	}
}

func createQueryToMatchPackage[M any](pkg *model.PkgSpec) *gremlinQueryBuilder[M] {
	query := createGraphQuery(Package)
	if pkg.ID != nil {
		query.id = *pkg.ID
	}
	if pkg.Type != nil {
		query.has[typeStr] = *pkg.Type
	}
	if pkg.Namespace != nil {
		query.has[namespace] = *pkg.Namespace
	}
	if pkg.Name != nil {
		query.has[name] = *pkg.Name
	}
	if pkg.Subpath != nil {
		query.has[subpath] = *pkg.Subpath
	}
	if pkg.Version != nil {
		// *filter.Version != v.version ||
		//	noMatchInput(filter.Subpath, v.subpath) ||
		//	noMatchQualifiers(filter, v.qualifiers) {
	}
	return &gremlinQueryBuilder[M]{query: query}
}

func createQueryToMatchPackageDependency[M any](pkg *model.PkgSpec) *gremlinQueryBuilder[M] {
	query := createGraphQuery(Package)
	if pkg.ID != nil {
		query.id = *pkg.ID
	}
	if pkg.Type != nil {
		query.has[typeStr] = *pkg.Type
	}
	if pkg.Namespace != nil {
		query.has[namespace] = *pkg.Namespace
	}
	if pkg.Name != nil {
		query.has[name] = *pkg.Name
	}
	if pkg.Version != nil {
		// *filter.Version != v.version ||
		//	noMatchInput(filter.Subpath, v.subpath) ||
		//	noMatchQualifiers(filter, v.qualifiers) {
	}
	return &gremlinQueryBuilder[M]{query: query}
}

func createQueryToMatchPackageName[M any](pkg *model.PkgNameSpec) *gremlinQueryBuilder[M] {
	query := createGraphQuery(Package)
	if pkg.ID != nil {
		query.id = *pkg.ID
	}
	if pkg.Type != nil {
		query.has[typeStr] = *pkg.Type
	}
	if pkg.Namespace != nil {
		query.has[namespace] = *pkg.Namespace
	}
	if pkg.Name != nil {
		query.has[name] = *pkg.Name
	}
	return &gremlinQueryBuilder[M]{query: query}
}

func (gqb *gremlinQueryBuilder[M]) withOrderByKey(orderBy string) *gremlinQueryBuilder[M] {
	gqb.query.orderByKey = orderBy
	return gqb
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

func (gqb *gremlinQueryBuilder[M]) isEmpty() bool {
	return gqb.query.isEmpty()
}

func (gqb *gremlinQueryBuilder[M]) withMapper(mapper func(*gremlinQueryResult) M) *gremlinQueryBuilder[M] {
	return gqb
}

func (gqb *gremlinQueryBuilder[M]) withQueries(queries []*gremlinQueryBuilder[M]) *gremlinQueryBuilder[M] {
	return gqb
}

// terminal steps
func (gqb *gremlinQueryBuilder[M]) upsertBulk() ([]M, error) {
	return nil, nil
}

func (gqb *gremlinQueryBuilder[M]) upsert() (M, error) {
	var object M
	return object, nil
}

func (q *gremlinQueryBuilder[M]) findAll() ([]M, error) {
	return nil, nil
}

/*
func queryEdge[M any](c *gremlinClient, q *gremlinQueryBuilder, deserializer EdgeObjectDeserializer[M]) ([]M, error) {
	return queryModelObjectsFromEdge[M](c, q.query, deserializer)
}

func upsertEdge[M any](c *gremlinClient, q *gremlinQueryBuilder, deserializer EdgeObjectDeserializer[M]) (M, error) {
	var object M
	relation := &Relation{
		outV: q.query.outVQuery,
		inV:  q.query.inVQuery,
		edge: q.query,
	}
	relationWithId, err := c.upsertRelationDirect(relation)
	if err != nil {
		return object, err
	}
	return deserializer(relationWithId.edgeId, relationWithId.outV, relation.edge.toReadMap(), relationWithId.inV), nil
}
*/
