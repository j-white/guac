package gremlin

import (
	"github.com/guacsec/guac/pkg/assembler/graphql/model"
	"time"
)

type gremlinQueryBuilder struct {
	query *GraphQuery
}

func createQueryForVertex(label Label) *gremlinQueryBuilder {
	return &gremlinQueryBuilder{
		query: createGraphQuery(label),
	}
}

func createQueryForEdge(label Label) *gremlinQueryBuilder {
	return &gremlinQueryBuilder{
		query: createGraphQuery(label),
	}
}

func createQueryToMatchPackage(pkg *model.PkgSpec) *gremlinQueryBuilder {
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
	return &gremlinQueryBuilder{query: query}
}

func createQueryToMatchPackageDependency(pkg *model.PkgSpec) *gremlinQueryBuilder {
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
	return &gremlinQueryBuilder{query: query}
}

func createQueryToMatchPackageName(pkg *model.PkgNameSpec) *gremlinQueryBuilder {
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
	return &gremlinQueryBuilder{query: query}
}

func (gqb *gremlinQueryBuilder) withId(id *string) *gremlinQueryBuilder {
	if id != nil {
		gqb.query.id = *id
	}
	return gqb
}

func (gqb *gremlinQueryBuilder) withPropString(key string, value *string) *gremlinQueryBuilder {
	if value != nil {
		gqb.query.has[key] = *value
	}
	return gqb
}

func (gqb *gremlinQueryBuilder) withPropDependencyType(key string, value *model.DependencyType) *gremlinQueryBuilder {
	if value != nil {
		gqb.query.has[key] = value.String()
	}
	return gqb
}

func (gqb *gremlinQueryBuilder) withPropTime(key string, value *time.Time) *gremlinQueryBuilder {
	if value != nil {
		gqb.query.has[key] = *value
	}
	return gqb
}

func (gqb *gremlinQueryBuilder) withPropFloat64(key string, value *float64) *gremlinQueryBuilder {
	if value != nil {
		gqb.query.has[key] = *value
	}
	return gqb
}

func (gqb *gremlinQueryBuilder) withInVertex(q *gremlinQueryBuilder) *gremlinQueryBuilder {
	gqb.query.inVQuery = q.query
	return gqb
}

func (gqb *gremlinQueryBuilder) withOutVertex(q *gremlinQueryBuilder) *gremlinQueryBuilder {
	gqb.query.outVQuery = q.query
	return gqb
}

func (gqb *gremlinQueryBuilder) isEmpty() bool {
	return gqb.query.isEmpty()
}

func queryEdge[M any](c *gremlinClient, q *gremlinQueryBuilder, deserializer EdgeObjectDeserializer[M]) ([]M, error) {
	return queryModelObjectsFromEdge[M](c, q.query, deserializer)
}
