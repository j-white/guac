package gremlin

import (
	"context"
	"github.com/guacsec/guac/pkg/assembler/graphql/model"
)

const (
	IsDependency Label = "isDependency"
)

func getDependencyObjectFromEdgeMuted(result *gremlinQueryResult) *model.IsDependency {
	// remove other values
	inValuesSpec := make(map[interface{}]interface{})
	inValuesSpec[typeStr] = result.in[typeStr]
	inValuesSpec[name] = result.in[name]
	inValuesSpec[namespace] = result.in[namespace]
	result.in = inValuesSpec
	return getDependencyObjectFromEdge(result)
}

func getDependencyObjectFromEdge(result *gremlinQueryResult) *model.IsDependency {
	id := result.id
	outValues := result.out
	edgeValues := result.edge
	inValues := result.in

	pkg := getPackageObject("", outValues)
	depPkg := getPackageObject("", inValues)

	isDependency := &model.IsDependency{
		ID:               id,
		Package:          pkg,
		DependentPackage: depPkg,
		VersionRange:     edgeValues[versionRange].(string),
		DependencyType:   model.DependencyType(edgeValues[dependencyType].(string)),
		Justification:    edgeValues[justification].(string),
		Origin:           edgeValues[collector].(string),
		Collector:        edgeValues[origin].(string),
	}
	return isDependency
}

func createQueryToMatchPackageNameFromInput[M any](pkg *model.PkgInputSpec) *gremlinQueryBuilder[M] {
	return createQueryForVertex[M](Package).
		withPropString(typeStr, &pkg.Type).
		withPropString(name, &pkg.Name).
		withPropString(namespace, pkg.Namespace)
}

func createUpsertForIsDependency(pkg *model.PkgInputSpec, depPkg *model.PkgInputSpec, dependency *model.IsDependencyInputSpec) *gremlinQueryBuilder[*model.IsDependency] {
	// compute guac keys
	pkgId := guacPkgId(*pkg)
	depPkgId := guacPkgId(*depPkg)

	return createUpsertForEdge[*model.IsDependency](IsDependency).
		withPropString(guacPkgVersionKey, &pkgId.VersionId).
		withPropString(guacSecondPkgNameKey, &depPkgId.NameId).
		withPropString(versionRange, &dependency.VersionRange).
		withPropDependencyType(dependencyType, &dependency.DependencyType).
		withPropString(justification, &dependency.Justification).
		withPropString(origin, &dependency.Origin).
		withPropString(collector, &dependency.Collector).
		withOutVertex(createQueryToMatchPackageInput[*model.IsDependency](pkg)).
		withInVertex(createQueryToMatchPackageNameFromInput[*model.IsDependency](depPkg)).
		withMapper(getDependencyObjectFromEdgeMuted)
}

// IngestDependency
//
//	pkg ->isDependency-> depPkg
func (c *gremlinClient) IngestDependency(ctx context.Context, pkg model.PkgInputSpec, depPkg model.PkgInputSpec, dependency model.IsDependencyInputSpec) (*model.IsDependency, error) {
	return createUpsertForIsDependency(&pkg, &depPkg, &dependency).upsert(c)
}

func (c *gremlinClient) IngestDependencies(ctx context.Context, pkgs []*model.PkgInputSpec, depPkgs []*model.PkgInputSpec, dependencies []*model.IsDependencyInputSpec) ([]*model.IsDependency, error) {
	// build the queries
	var queries []*gremlinQueryBuilder[*model.IsDependency]
	for k := range pkgs {
		queries = append(queries, createUpsertForIsDependency(pkgs[k], depPkgs[k], dependencies[k]))
	}

	return createBulkUpsertForEdge[*model.IsDependency](IsDependency).
		withQueries(queries).
		upsertBulk(c)
}

func (c *gremlinClient) IsDependency(ctx context.Context, isDependencySpec *model.IsDependencySpec) ([]*model.IsDependency, error) {
	q := createQueryForEdge[*model.IsDependency](IsDependency).
		withId(isDependencySpec.ID).
		withPropDependencyType(dependencyType, isDependencySpec.DependencyType).
		withPropString(justification, isDependencySpec.Justification).
		withPropString(origin, isDependencySpec.Origin).
		withPropString(collector, isDependencySpec.Collector).
		withPropString(versionRange, isDependencySpec.VersionRange).
		withOrderByKey(guacPkgVersionKey).
		withMapper(getDependencyObjectFromEdgeMuted)
	if isDependencySpec.Package != nil {
		q = q.withOutVertex(createQueryToMatchPackage[*model.IsDependency](isDependencySpec.Package))
	}
	if isDependencySpec.DependentPackage != nil {
		q = q.withInVertex(createQueryToMatchPackageName[*model.IsDependency](isDependencySpec.DependentPackage))
	}
	return q.findAll(c)
}
