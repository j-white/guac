package gremlin

import (
	"context"
	"github.com/guacsec/guac/pkg/assembler/graphql/model"
)

const (
	IsDependency Label = "isDependency"
)

//
//func getPackageQueryValuesForDepMatching(pkg *model.PkgInputSpec) *GraphQuery {
//	q := createGraphQuery(Package)
//	q.has[name] = pkg.Name
//	if pkg.Namespace != nil {
//		q.has[namespace] = *pkg.Namespace
//	} else {
//		q.has[namespace] = ""
//	}
//	return q
//}

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

func getDependencyObjectFromEdge(id string, outValues map[interface{}]interface{}, edgeValues map[interface{}]interface{}, inValues map[interface{}]interface{}) *model.IsDependency {
	pkg := getPackageObject("", outValues)
	depPkg := getPackageObject("", inValues)

	isDependency := &model.IsDependency{
		ID:               id,
		Package:          pkg,
		DependentPackage: depPkg,
		VersionRange:     "",
		DependencyType:   "",
		Justification:    edgeValues[justification].(string),
		Origin:           edgeValues[collector].(string),
		Collector:        edgeValues[origin].(string),
	}
	return isDependency
}

// IngestDependency
//
//	pkg ->isDependency-> depPkg
func (c *gremlinClient) IngestDependency(ctx context.Context, pkg model.PkgInputSpec, depPkg model.PkgInputSpec, dependency model.IsDependencyInputSpec) (*model.IsDependency, error) {
	return ingestModelObjectsWithRelation[*model.PkgInputSpec, *model.IsDependencyInputSpec, *model.IsDependency](
		c, &pkg, &depPkg, &dependency, getPackageQueryValues, getDependencyQueryValues, getDependencyObjectFromEdge)
}

func (c *gremlinClient) IngestDependencies(ctx context.Context, pkgs []*model.PkgInputSpec, depPkgs []*model.PkgInputSpec, dependencies []*model.IsDependencyInputSpec) ([]*model.IsDependency, error) {
	return bulkIngestModelObjectsWithRelation[*model.PkgInputSpec, *model.IsDependencyInputSpec, *model.IsDependency](
		c, pkgs, depPkgs, dependencies, getPackageQueryValues, getDependencyQueryValues, getDependencyObjectFromEdge)
}

func (c *gremlinClient) IsDependency(ctx context.Context, isDependencySpec *model.IsDependencySpec) ([]*model.IsDependency, error) {
	query := createGraphQuery(IsDependency)
	if isDependencySpec != nil {
		if isDependencySpec.ID != nil {
			query.id = *isDependencySpec.ID
		}
	}
	return queryModelObjectsFromEdge[*model.IsDependency](c, query, getDependencyObjectFromEdge)
}
