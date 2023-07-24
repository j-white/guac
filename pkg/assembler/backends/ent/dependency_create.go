// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/gremlin"
	"entgo.io/ent/dialect/gremlin/graph/dsl"
	"entgo.io/ent/dialect/gremlin/graph/dsl/g"
	"github.com/guacsec/guac/pkg/assembler/backends/ent/dependency"
)

// DependencyCreate is the builder for creating a Dependency entity.
type DependencyCreate struct {
	config
	mutation *DependencyMutation
	hooks    []Hook
}

// SetPackageID sets the "package_id" field.
func (dc *DependencyCreate) SetPackageID(i int) *DependencyCreate {
	dc.mutation.SetPackageID(i)
	return dc
}

// SetDependentPackageID sets the "dependent_package_id" field.
func (dc *DependencyCreate) SetDependentPackageID(i int) *DependencyCreate {
	dc.mutation.SetDependentPackageID(i)
	return dc
}

// SetVersionRange sets the "version_range" field.
func (dc *DependencyCreate) SetVersionRange(s string) *DependencyCreate {
	dc.mutation.SetVersionRange(s)
	return dc
}

// SetDependencyType sets the "dependency_type" field.
func (dc *DependencyCreate) SetDependencyType(dt dependency.DependencyType) *DependencyCreate {
	dc.mutation.SetDependencyType(dt)
	return dc
}

// SetJustification sets the "justification" field.
func (dc *DependencyCreate) SetJustification(s string) *DependencyCreate {
	dc.mutation.SetJustification(s)
	return dc
}

// SetOrigin sets the "origin" field.
func (dc *DependencyCreate) SetOrigin(s string) *DependencyCreate {
	dc.mutation.SetOrigin(s)
	return dc
}

// SetCollector sets the "collector" field.
func (dc *DependencyCreate) SetCollector(s string) *DependencyCreate {
	dc.mutation.SetCollector(s)
	return dc
}

// SetPackage sets the "package" edge to the PackageVersion entity.
func (dc *DependencyCreate) SetPackage(p *PackageVersion) *DependencyCreate {
	return dc.SetPackageID(p.ID)
}

// SetDependentPackage sets the "dependent_package" edge to the PackageName entity.
func (dc *DependencyCreate) SetDependentPackage(p *PackageName) *DependencyCreate {
	return dc.SetDependentPackageID(p.ID)
}

// Mutation returns the DependencyMutation object of the builder.
func (dc *DependencyCreate) Mutation() *DependencyMutation {
	return dc.mutation
}

// Save creates the Dependency in the database.
func (dc *DependencyCreate) Save(ctx context.Context) (*Dependency, error) {
	return withHooks(ctx, dc.gremlinSave, dc.mutation, dc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (dc *DependencyCreate) SaveX(ctx context.Context) *Dependency {
	v, err := dc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (dc *DependencyCreate) Exec(ctx context.Context) error {
	_, err := dc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dc *DependencyCreate) ExecX(ctx context.Context) {
	if err := dc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (dc *DependencyCreate) check() error {
	if _, ok := dc.mutation.PackageID(); !ok {
		return &ValidationError{Name: "package_id", err: errors.New(`ent: missing required field "Dependency.package_id"`)}
	}
	if _, ok := dc.mutation.DependentPackageID(); !ok {
		return &ValidationError{Name: "dependent_package_id", err: errors.New(`ent: missing required field "Dependency.dependent_package_id"`)}
	}
	if _, ok := dc.mutation.VersionRange(); !ok {
		return &ValidationError{Name: "version_range", err: errors.New(`ent: missing required field "Dependency.version_range"`)}
	}
	if _, ok := dc.mutation.DependencyType(); !ok {
		return &ValidationError{Name: "dependency_type", err: errors.New(`ent: missing required field "Dependency.dependency_type"`)}
	}
	if v, ok := dc.mutation.DependencyType(); ok {
		if err := dependency.DependencyTypeValidator(v); err != nil {
			return &ValidationError{Name: "dependency_type", err: fmt.Errorf(`ent: validator failed for field "Dependency.dependency_type": %w`, err)}
		}
	}
	if _, ok := dc.mutation.Justification(); !ok {
		return &ValidationError{Name: "justification", err: errors.New(`ent: missing required field "Dependency.justification"`)}
	}
	if _, ok := dc.mutation.Origin(); !ok {
		return &ValidationError{Name: "origin", err: errors.New(`ent: missing required field "Dependency.origin"`)}
	}
	if _, ok := dc.mutation.Collector(); !ok {
		return &ValidationError{Name: "collector", err: errors.New(`ent: missing required field "Dependency.collector"`)}
	}
	if _, ok := dc.mutation.PackageID(); !ok {
		return &ValidationError{Name: "package", err: errors.New(`ent: missing required edge "Dependency.package"`)}
	}
	if _, ok := dc.mutation.DependentPackageID(); !ok {
		return &ValidationError{Name: "dependent_package", err: errors.New(`ent: missing required edge "Dependency.dependent_package"`)}
	}
	return nil
}

func (dc *DependencyCreate) gremlinSave(ctx context.Context) (*Dependency, error) {
	if err := dc.check(); err != nil {
		return nil, err
	}
	res := &gremlin.Response{}
	query, bindings := dc.gremlin().Query()
	if err := dc.driver.Exec(ctx, query, bindings, res); err != nil {
		return nil, err
	}
	if err, ok := isConstantError(res); ok {
		return nil, err
	}
	rnode := &Dependency{config: dc.config}
	if err := rnode.FromResponse(res); err != nil {
		return nil, err
	}
	dc.mutation.id = &rnode.ID
	dc.mutation.done = true
	return rnode, nil
}

func (dc *DependencyCreate) gremlin() *dsl.Traversal {
	v := g.AddV(dependency.Label)
	if value, ok := dc.mutation.VersionRange(); ok {
		v.Property(dsl.Single, dependency.FieldVersionRange, value)
	}
	if value, ok := dc.mutation.DependencyType(); ok {
		v.Property(dsl.Single, dependency.FieldDependencyType, value)
	}
	if value, ok := dc.mutation.Justification(); ok {
		v.Property(dsl.Single, dependency.FieldJustification, value)
	}
	if value, ok := dc.mutation.Origin(); ok {
		v.Property(dsl.Single, dependency.FieldOrigin, value)
	}
	if value, ok := dc.mutation.Collector(); ok {
		v.Property(dsl.Single, dependency.FieldCollector, value)
	}
	for _, id := range dc.mutation.PackageIDs() {
		v.AddE(dependency.PackageLabel).To(g.V(id)).OutV()
	}
	for _, id := range dc.mutation.DependentPackageIDs() {
		v.AddE(dependency.DependentPackageLabel).To(g.V(id)).OutV()
	}
	return v.ValueMap(true)
}

// DependencyCreateBulk is the builder for creating many Dependency entities in bulk.
type DependencyCreateBulk struct {
	config
	builders []*DependencyCreate
}
