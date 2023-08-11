package gremlin

import (
	"context"
	"github.com/guacsec/guac/pkg/assembler/graphql/model"
)

const (
	IsDependency Label = "isDependency"
)

func getDependencyQueryValues(pkg *model.PkgInputSpec, depPkg *model.PkgInputSpec, dependency *model.IsDependencyInputSpec) *GraphQuery {
	q := createGraphQuery(IsDependency)

	// add guac keys
	pkgId := guacPkgId(*pkg)
	depPkgId := guacPkgId(*depPkg)
	q.has["pkgVersionGuacKey"] = pkgId.VersionId
	q.has["secondPkgNameGuacKey"] = depPkgId.NameId

	// isDependency
	q.has[versionRange] = dependency.VersionRange
	q.has[dependencyType] = dependency.DependencyType.String()
	q.has[justification] = dependency.Justification
	q.has[origin] = dependency.Origin
	q.has[collector] = dependency.Collector

	return q
}

func getDependencyObjectFromEdgeMuted(id string, outValues map[interface{}]interface{}, edgeValues map[interface{}]interface{}, inValues map[interface{}]interface{}) *model.IsDependency {
	// remove other values
	inValuesSpec := make(map[interface{}]interface{})
	inValuesSpec[typeStr] = inValues[typeStr]
	inValuesSpec[name] = inValues[name]
	inValuesSpec[namespace] = inValues[namespace]
	return getDependencyObjectFromEdge(id, outValues, edgeValues, inValuesSpec)
}

func getDependencyObjectFromEdge(id string, outValues map[interface{}]interface{}, edgeValues map[interface{}]interface{}, inValues map[interface{}]interface{}) *model.IsDependency {
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

func getPackageQueryValuesForDep(pkg *model.PkgInputSpec) *GraphQuery {
	q := createGraphQuery(Package)
	q.has[typeStr] = pkg.Type
	q.has[name] = pkg.Name
	if pkg.Namespace != nil {
		q.has[namespace] = *pkg.Namespace
	}
	return q
}

// IngestDependency
//
//	pkg ->isDependency-> depPkg
func (c *gremlinClient) IngestDependency(ctx context.Context, pkg model.PkgInputSpec, depPkg model.PkgInputSpec, dependency model.IsDependencyInputSpec) (*model.IsDependency, error) {
	return ingestModelObjectsWithRelation[*model.PkgInputSpec, *model.IsDependencyInputSpec, *model.IsDependency](
		c, &pkg, &depPkg, &dependency, getPackageQueryValues, getPackageQueryValuesForDep, getDependencyQueryValues, getDependencyObjectFromEdgeMuted)
}

func (c *gremlinClient) IngestDependencies(ctx context.Context, pkgs []*model.PkgInputSpec, depPkgs []*model.PkgInputSpec, dependencies []*model.IsDependencyInputSpec) ([]*model.IsDependency, error) {
	return bulkIngestModelObjectsWithRelation[*model.PkgInputSpec, *model.IsDependencyInputSpec, *model.IsDependency](
		c, pkgs, depPkgs, dependencies, getPackageQueryValues, getPackageQueryValuesForDep, getDependencyQueryValues, getDependencyObjectFromEdgeMuted)
}

func (c *gremlinClient) IsDependency(ctx context.Context, isDependencySpec *model.IsDependencySpec) ([]*model.IsDependency, error) {
	q := createQueryForEdge(IsDependency).
		withId(isDependencySpec.ID).
		withPropDependencyType(dependencyType, isDependencySpec.DependencyType).
		withPropString(justification, isDependencySpec.Justification).
		withPropString(origin, isDependencySpec.Origin).
		withPropString(collector, isDependencySpec.Collector).
		withPropString(versionRange, isDependencySpec.VersionRange)
	if isDependencySpec.Package != nil {
		q = q.withOutVertex(createQueryToMatchPackage(isDependencySpec.Package))
	}
	if isDependencySpec.DependentPackage != nil {
		q = q.withInVertex(createQueryToMatchPackageName(isDependencySpec.DependentPackage))
	}
	return queryEdge[*model.IsDependency](c, q, getDependencyObjectFromEdgeMuted)
}
