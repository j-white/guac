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
		DependencyType:   model.DependencyType(edgeValues[dependencyType].(string)),
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
	// Note: depPkgSpec only takes up to the pkgName as IsDependency does not allow for the attestation
	// to be made at the pkgVersion level. Version range for the dependent package is defined as a property
	// on IsDependency.
	depPkgSpec := model.PkgInputSpec{
		Type:       depPkg.Type,
		Namespace:  depPkg.Namespace,
		Name:       depPkg.Name,
		Version:    nil,
		Subpath:    nil,
		Qualifiers: nil,
	}

	return ingestModelObjectsWithRelation[*model.PkgInputSpec, *model.IsDependencyInputSpec, *model.IsDependency](
		c, &pkg, &depPkgSpec, &dependency, getPackageQueryValues, getDependencyQueryValues, getDependencyObjectFromEdge)
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
		if isDependencySpec.DependencyType != nil {
			query.has[dependencyType] = isDependencySpec.DependencyType.String()
		}
		if isDependencySpec.Justification != nil {
			query.has[justification] = *isDependencySpec.Justification
		}
		if isDependencySpec.Origin != nil {
			query.has[origin] = *isDependencySpec.Origin
		}
		if isDependencySpec.Collector != nil {
			query.has[collector] = *isDependencySpec.Collector
		}
		if isDependencySpec.DependentPackage != nil {
			inVQuery := createGraphQuery(Package)
			if isDependencySpec.DependentPackage.ID != nil {
				inVQuery.id = *isDependencySpec.DependentPackage.ID
			}
			if isDependencySpec.DependentPackage.Type != nil {
				inVQuery.has[typeStr] = *isDependencySpec.DependentPackage.Type
			}
			if isDependencySpec.DependentPackage.Namespace != nil {
				inVQuery.has[namespace] = *isDependencySpec.DependentPackage.Namespace
			}
			if isDependencySpec.DependentPackage.Name != nil {
				inVQuery.has[name] = *isDependencySpec.DependentPackage.Name
			}
			query.inVQuery = inVQuery
		}
	}
	// FIXME: Should this be done for all?
	if query.isEmpty() {
		return nil, nil
	}
	return queryModelObjectsFromEdge[*model.IsDependency](c, query, getDependencyObjectFromEdge)
}
