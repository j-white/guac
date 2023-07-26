package tinkerpop

import (
	"context"
	"fmt"
	"github.com/guacsec/guac/pkg/assembler/graphql/model"
)

func getDependencyQueryValues(pkg *model.PkgInputSpec, depPkg *model.PkgInputSpec, dependency *model.IsDependencyInputSpec) map[string]any {
	values := map[string]any{}

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

func (c *tinkerpopClient) IngestDependency(ctx context.Context, pkg model.PkgInputSpec, depPkg model.PkgInputSpec, dependency model.IsDependencyInputSpec) (*model.IsDependency, error) {
	getDependencyQueryValues(&pkg, &depPkg, &dependency)
	//TODO implement me
	panic("implement me")
}

func (c *tinkerpopClient) IngestDependencies(ctx context.Context, pkgs []*model.PkgInputSpec, depPkgs []*model.PkgInputSpec, dependencies []*model.IsDependencyInputSpec) ([]*model.IsDependency, error) {
	//TODO implement me
	panic("implement me")
}

func (c *tinkerpopClient) IsDependency(ctx context.Context, isDependencySpec *model.IsDependencySpec) ([]*model.IsDependency, error) {
	return []*model.IsDependency{}, fmt.Errorf("not implemented: IsDependency")
}
