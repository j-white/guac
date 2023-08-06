package tinkerpop

import (
	"context"
	"fmt"
	gremlingo "github.com/apache/tinkerpop/gremlin-go/v3/driver"
	"github.com/guacsec/guac/pkg/assembler/graphql/model"
	"strconv"
)

const (
	IsDependency Label = "isDependency"
)

func getDependencyQueryValues(pkg *model.PkgInputSpec, depPkg *model.PkgInputSpec, dependency *model.IsDependencyInputSpec) map[interface{}]interface{} {
	values := make(map[interface{}]interface{})
	values[gremlingo.T.Label] = string(IsDependency)

	// add guac keys
	pkgId := guacPkgId(*pkg)
	depPkgId := guacPkgId(*depPkg)
	values["pkgVersionGuacKey"] = pkgId.VersionId
	values["secondPkgNameGuacKey"] = depPkgId.NameId

	// isDependency
	values[versionRange] = dependency.VersionRange
	values[dependencyType] = dependency.DependencyType.String()
	values[justification] = dependency.Justification
	values[origin] = dependency.Origin
	values[collector] = dependency.Collector

	return values
}

func getDependencyObject(id int64, values map[interface{}]interface{}) *model.IsDependency {
	isDependency := &model.IsDependency{
		ID:               strconv.FormatInt(id, 10),
		Package:          &model.Package{},
		DependentPackage: &model.Package{},
		VersionRange:     "",
		DependencyType:   model.DependencyTypeDirect,
		Justification:    values[justification].(string),
		Origin:           values[collector].(string),
		Collector:        values[origin].(string),
	}
	return isDependency
}

func (c *tinkerpopClient) IngestDependency(ctx context.Context, pkg model.PkgInputSpec, depPkg model.PkgInputSpec, dependency model.IsDependencyInputSpec) (*model.IsDependency, error) {
	return ingestModelObjectsWithRelation[*model.PkgInputSpec, *model.IsDependencyInputSpec, *model.IsDependency](
		c, &pkg, &depPkg, &dependency, getPackageQueryValues, getDependencyQueryValues, getDependencyObject)
}

func (c *tinkerpopClient) IngestDependencies(ctx context.Context, pkgs []*model.PkgInputSpec, depPkgs []*model.PkgInputSpec, dependencies []*model.IsDependencyInputSpec) ([]*model.IsDependency, error) {
	return bulkIngestModelObjectsWithRelation[*model.PkgInputSpec, *model.IsDependencyInputSpec, *model.IsDependency](
		c, pkgs, depPkgs, dependencies, getPackageQueryValues, getDependencyQueryValues, getDependencyObject)
}

func (c *tinkerpopClient) IsDependency(ctx context.Context, isDependencySpec *model.IsDependencySpec) ([]*model.IsDependency, error) {
	return []*model.IsDependency{}, fmt.Errorf("not implemented: IsDependency")
}
