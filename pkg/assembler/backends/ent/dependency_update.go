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
	"github.com/guacsec/guac/pkg/assembler/backends/ent/predicate"
)

// DependencyUpdate is the builder for updating Dependency entities.
type DependencyUpdate struct {
	config
	hooks    []Hook
	mutation *DependencyMutation
}

// Where appends a list predicates to the DependencyUpdate builder.
func (du *DependencyUpdate) Where(ps ...predicate.Dependency) *DependencyUpdate {
	du.mutation.Where(ps...)
	return du
}

// SetPackageID sets the "package_id" field.
func (du *DependencyUpdate) SetPackageID(i int) *DependencyUpdate {
	du.mutation.SetPackageID(i)
	return du
}

// SetDependentPackageID sets the "dependent_package_id" field.
func (du *DependencyUpdate) SetDependentPackageID(i int) *DependencyUpdate {
	du.mutation.SetDependentPackageID(i)
	return du
}

// SetVersionRange sets the "version_range" field.
func (du *DependencyUpdate) SetVersionRange(s string) *DependencyUpdate {
	du.mutation.SetVersionRange(s)
	return du
}

// SetDependencyType sets the "dependency_type" field.
func (du *DependencyUpdate) SetDependencyType(dt dependency.DependencyType) *DependencyUpdate {
	du.mutation.SetDependencyType(dt)
	return du
}

// SetJustification sets the "justification" field.
func (du *DependencyUpdate) SetJustification(s string) *DependencyUpdate {
	du.mutation.SetJustification(s)
	return du
}

// SetOrigin sets the "origin" field.
func (du *DependencyUpdate) SetOrigin(s string) *DependencyUpdate {
	du.mutation.SetOrigin(s)
	return du
}

// SetCollector sets the "collector" field.
func (du *DependencyUpdate) SetCollector(s string) *DependencyUpdate {
	du.mutation.SetCollector(s)
	return du
}

// SetPackage sets the "package" edge to the PackageVersion entity.
func (du *DependencyUpdate) SetPackage(p *PackageVersion) *DependencyUpdate {
	return du.SetPackageID(p.ID)
}

// SetDependentPackage sets the "dependent_package" edge to the PackageName entity.
func (du *DependencyUpdate) SetDependentPackage(p *PackageName) *DependencyUpdate {
	return du.SetDependentPackageID(p.ID)
}

// Mutation returns the DependencyMutation object of the builder.
func (du *DependencyUpdate) Mutation() *DependencyMutation {
	return du.mutation
}

// ClearPackage clears the "package" edge to the PackageVersion entity.
func (du *DependencyUpdate) ClearPackage() *DependencyUpdate {
	du.mutation.ClearPackage()
	return du
}

// ClearDependentPackage clears the "dependent_package" edge to the PackageName entity.
func (du *DependencyUpdate) ClearDependentPackage() *DependencyUpdate {
	du.mutation.ClearDependentPackage()
	return du
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (du *DependencyUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, du.gremlinSave, du.mutation, du.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (du *DependencyUpdate) SaveX(ctx context.Context) int {
	affected, err := du.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (du *DependencyUpdate) Exec(ctx context.Context) error {
	_, err := du.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (du *DependencyUpdate) ExecX(ctx context.Context) {
	if err := du.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (du *DependencyUpdate) check() error {
	if v, ok := du.mutation.DependencyType(); ok {
		if err := dependency.DependencyTypeValidator(v); err != nil {
			return &ValidationError{Name: "dependency_type", err: fmt.Errorf(`ent: validator failed for field "Dependency.dependency_type": %w`, err)}
		}
	}
	if _, ok := du.mutation.PackageID(); du.mutation.PackageCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Dependency.package"`)
	}
	if _, ok := du.mutation.DependentPackageID(); du.mutation.DependentPackageCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Dependency.dependent_package"`)
	}
	return nil
}

func (du *DependencyUpdate) gremlinSave(ctx context.Context) (int, error) {
	if err := du.check(); err != nil {
		return 0, err
	}
	res := &gremlin.Response{}
	query, bindings := du.gremlin().Query()
	if err := du.driver.Exec(ctx, query, bindings, res); err != nil {
		return 0, err
	}
	if err, ok := isConstantError(res); ok {
		return 0, err
	}
	du.mutation.done = true
	return res.ReadInt()
}

func (du *DependencyUpdate) gremlin() *dsl.Traversal {
	v := g.V().HasLabel(dependency.Label)
	for _, p := range du.mutation.predicates {
		p(v)
	}
	var (
		rv = v.Clone()
		_  = rv

		trs []*dsl.Traversal
	)
	if value, ok := du.mutation.VersionRange(); ok {
		v.Property(dsl.Single, dependency.FieldVersionRange, value)
	}
	if value, ok := du.mutation.DependencyType(); ok {
		v.Property(dsl.Single, dependency.FieldDependencyType, value)
	}
	if value, ok := du.mutation.Justification(); ok {
		v.Property(dsl.Single, dependency.FieldJustification, value)
	}
	if value, ok := du.mutation.Origin(); ok {
		v.Property(dsl.Single, dependency.FieldOrigin, value)
	}
	if value, ok := du.mutation.Collector(); ok {
		v.Property(dsl.Single, dependency.FieldCollector, value)
	}
	if du.mutation.PackageCleared() {
		tr := rv.Clone().OutE(dependency.PackageLabel).Drop().Iterate()
		trs = append(trs, tr)
	}
	for _, id := range du.mutation.PackageIDs() {
		v.AddE(dependency.PackageLabel).To(g.V(id)).OutV()
	}
	if du.mutation.DependentPackageCleared() {
		tr := rv.Clone().OutE(dependency.DependentPackageLabel).Drop().Iterate()
		trs = append(trs, tr)
	}
	for _, id := range du.mutation.DependentPackageIDs() {
		v.AddE(dependency.DependentPackageLabel).To(g.V(id)).OutV()
	}
	v.Count()
	trs = append(trs, v)
	return dsl.Join(trs...)
}

// DependencyUpdateOne is the builder for updating a single Dependency entity.
type DependencyUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *DependencyMutation
}

// SetPackageID sets the "package_id" field.
func (duo *DependencyUpdateOne) SetPackageID(i int) *DependencyUpdateOne {
	duo.mutation.SetPackageID(i)
	return duo
}

// SetDependentPackageID sets the "dependent_package_id" field.
func (duo *DependencyUpdateOne) SetDependentPackageID(i int) *DependencyUpdateOne {
	duo.mutation.SetDependentPackageID(i)
	return duo
}

// SetVersionRange sets the "version_range" field.
func (duo *DependencyUpdateOne) SetVersionRange(s string) *DependencyUpdateOne {
	duo.mutation.SetVersionRange(s)
	return duo
}

// SetDependencyType sets the "dependency_type" field.
func (duo *DependencyUpdateOne) SetDependencyType(dt dependency.DependencyType) *DependencyUpdateOne {
	duo.mutation.SetDependencyType(dt)
	return duo
}

// SetJustification sets the "justification" field.
func (duo *DependencyUpdateOne) SetJustification(s string) *DependencyUpdateOne {
	duo.mutation.SetJustification(s)
	return duo
}

// SetOrigin sets the "origin" field.
func (duo *DependencyUpdateOne) SetOrigin(s string) *DependencyUpdateOne {
	duo.mutation.SetOrigin(s)
	return duo
}

// SetCollector sets the "collector" field.
func (duo *DependencyUpdateOne) SetCollector(s string) *DependencyUpdateOne {
	duo.mutation.SetCollector(s)
	return duo
}

// SetPackage sets the "package" edge to the PackageVersion entity.
func (duo *DependencyUpdateOne) SetPackage(p *PackageVersion) *DependencyUpdateOne {
	return duo.SetPackageID(p.ID)
}

// SetDependentPackage sets the "dependent_package" edge to the PackageName entity.
func (duo *DependencyUpdateOne) SetDependentPackage(p *PackageName) *DependencyUpdateOne {
	return duo.SetDependentPackageID(p.ID)
}

// Mutation returns the DependencyMutation object of the builder.
func (duo *DependencyUpdateOne) Mutation() *DependencyMutation {
	return duo.mutation
}

// ClearPackage clears the "package" edge to the PackageVersion entity.
func (duo *DependencyUpdateOne) ClearPackage() *DependencyUpdateOne {
	duo.mutation.ClearPackage()
	return duo
}

// ClearDependentPackage clears the "dependent_package" edge to the PackageName entity.
func (duo *DependencyUpdateOne) ClearDependentPackage() *DependencyUpdateOne {
	duo.mutation.ClearDependentPackage()
	return duo
}

// Where appends a list predicates to the DependencyUpdate builder.
func (duo *DependencyUpdateOne) Where(ps ...predicate.Dependency) *DependencyUpdateOne {
	duo.mutation.Where(ps...)
	return duo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (duo *DependencyUpdateOne) Select(field string, fields ...string) *DependencyUpdateOne {
	duo.fields = append([]string{field}, fields...)
	return duo
}

// Save executes the query and returns the updated Dependency entity.
func (duo *DependencyUpdateOne) Save(ctx context.Context) (*Dependency, error) {
	return withHooks(ctx, duo.gremlinSave, duo.mutation, duo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (duo *DependencyUpdateOne) SaveX(ctx context.Context) *Dependency {
	node, err := duo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (duo *DependencyUpdateOne) Exec(ctx context.Context) error {
	_, err := duo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (duo *DependencyUpdateOne) ExecX(ctx context.Context) {
	if err := duo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (duo *DependencyUpdateOne) check() error {
	if v, ok := duo.mutation.DependencyType(); ok {
		if err := dependency.DependencyTypeValidator(v); err != nil {
			return &ValidationError{Name: "dependency_type", err: fmt.Errorf(`ent: validator failed for field "Dependency.dependency_type": %w`, err)}
		}
	}
	if _, ok := duo.mutation.PackageID(); duo.mutation.PackageCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Dependency.package"`)
	}
	if _, ok := duo.mutation.DependentPackageID(); duo.mutation.DependentPackageCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Dependency.dependent_package"`)
	}
	return nil
}

func (duo *DependencyUpdateOne) gremlinSave(ctx context.Context) (*Dependency, error) {
	if err := duo.check(); err != nil {
		return nil, err
	}
	res := &gremlin.Response{}
	id, ok := duo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Dependency.id" for update`)}
	}
	query, bindings := duo.gremlin(id).Query()
	if err := duo.driver.Exec(ctx, query, bindings, res); err != nil {
		return nil, err
	}
	if err, ok := isConstantError(res); ok {
		return nil, err
	}
	duo.mutation.done = true
	d := &Dependency{config: duo.config}
	if err := d.FromResponse(res); err != nil {
		return nil, err
	}
	return d, nil
}

func (duo *DependencyUpdateOne) gremlin(id int) *dsl.Traversal {
	v := g.V(id)
	var (
		rv = v.Clone()
		_  = rv

		trs []*dsl.Traversal
	)
	if value, ok := duo.mutation.VersionRange(); ok {
		v.Property(dsl.Single, dependency.FieldVersionRange, value)
	}
	if value, ok := duo.mutation.DependencyType(); ok {
		v.Property(dsl.Single, dependency.FieldDependencyType, value)
	}
	if value, ok := duo.mutation.Justification(); ok {
		v.Property(dsl.Single, dependency.FieldJustification, value)
	}
	if value, ok := duo.mutation.Origin(); ok {
		v.Property(dsl.Single, dependency.FieldOrigin, value)
	}
	if value, ok := duo.mutation.Collector(); ok {
		v.Property(dsl.Single, dependency.FieldCollector, value)
	}
	if duo.mutation.PackageCleared() {
		tr := rv.Clone().OutE(dependency.PackageLabel).Drop().Iterate()
		trs = append(trs, tr)
	}
	for _, id := range duo.mutation.PackageIDs() {
		v.AddE(dependency.PackageLabel).To(g.V(id)).OutV()
	}
	if duo.mutation.DependentPackageCleared() {
		tr := rv.Clone().OutE(dependency.DependentPackageLabel).Drop().Iterate()
		trs = append(trs, tr)
	}
	for _, id := range duo.mutation.DependentPackageIDs() {
		v.AddE(dependency.DependentPackageLabel).To(g.V(id)).OutV()
	}
	if len(duo.fields) > 0 {
		fields := make([]any, 0, len(duo.fields)+1)
		fields = append(fields, true)
		for _, f := range duo.fields {
			fields = append(fields, f)
		}
		v.ValueMap(fields...)
	} else {
		v.ValueMap(true)
	}
	trs = append(trs, v)
	return dsl.Join(trs...)
}
