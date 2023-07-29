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

func getDependencyObject(id string, values map[interface{}]interface{}) *model.IsDependency {
	isDependency := &model.IsDependency{
		ID: id,
		//Package:          pkg,
		//DependentPackage: depPkg,
		//VersionRange:     createdValue.VersionRange,
		//DependencyType:   dependencyTypeEnum,
		Justification: values[justification].(string),
		Origin:        values[collector].(string),
		Collector:     values[origin].(string),
	}
	return isDependency
}

func (c *tinkerpopClient) IngestDependency(ctx context.Context, pkg model.PkgInputSpec, depPkg model.PkgInputSpec, dependency model.IsDependencyInputSpec) (*model.IsDependency, error) {
	pkgVertexProperties := getPackageQueryValues(&pkg)
	depPkgVertexProperties := getPackageQueryValues(&depPkg)
	dependencyEdgeProperties := getDependencyQueryValues(&pkg, &depPkg, &dependency)
	dependencyEdgeProperties[gremlingo.Direction.In] = gremlingo.Merge.InV
	dependencyEdgeProperties[gremlingo.Direction.Out] = gremlingo.Merge.OutV

	// upsert (pkg -- (dep) --> deppkg)
	g := gremlingo.Traversal_().WithRemote(c.remote)
	// FIXME: No usert here, should use v.Has() instead
	fmt.Printf("MOO upsert for edge: %v\n", dependencyEdgeProperties)

	r, err := g.MergeV(pkgVertexProperties).As("pkg").
		MergeV(depPkgVertexProperties).As("depPkg").
		MergeE(dependencyEdgeProperties).As("edge").
		// late bind
		Option(gremlingo.Merge.InV, gremlingo.T__.Select("pkg")).
		Option(gremlingo.Merge.OutV, gremlingo.T__.Select("depPkg")).
		Select("edge").Id().Next()
	fmt.Printf("MOO upsert for edge has results: %v %v\n", r, err)
	if err != nil {
		return nil, err
	}
	id := r.GetInterface() //.(*gremlingo.JanusRelationIdentifier)
	fmt.Printf("MOO upsert returned id: %v\n", id)

	return getDependencyObject(strconv.FormatInt(1, 10), dependencyEdgeProperties), nil
}

func (c *tinkerpopClient) IngestDependencies(ctx context.Context, pkgs []*model.PkgInputSpec, depPkgs []*model.PkgInputSpec, dependencies []*model.IsDependencyInputSpec) ([]*model.IsDependency, error) {
	// FIXME: Implement bulk insert
	var isDependencies []*model.IsDependency
	for k, depSpec := range dependencies {
		dep, err := c.IngestDependency(ctx, *pkgs[k], *depPkgs[k], *depSpec)
		if err != nil {
			return isDependencies, err
		}
		isDependencies = append(isDependencies, dep)
	}
	return isDependencies, nil
}

func (c *tinkerpopClient) IsDependency(ctx context.Context, isDependencySpec *model.IsDependencySpec) ([]*model.IsDependency, error) {
	return []*model.IsDependency{}, fmt.Errorf("not implemented: IsDependency")
}
