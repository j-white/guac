// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/guacsec/guac/pkg/assembler/backends/ent/hassourceat"
	"github.com/guacsec/guac/pkg/assembler/backends/ent/packagename"
	"github.com/guacsec/guac/pkg/assembler/backends/ent/packageversion"
	"github.com/guacsec/guac/pkg/assembler/backends/ent/predicate"
	"github.com/guacsec/guac/pkg/assembler/backends/ent/sourcename"
)

// HasSourceAtUpdate is the builder for updating HasSourceAt entities.
type HasSourceAtUpdate struct {
	config
	hooks    []Hook
	mutation *HasSourceAtMutation
}

// Where appends a list predicates to the HasSourceAtUpdate builder.
func (hsau *HasSourceAtUpdate) Where(ps ...predicate.HasSourceAt) *HasSourceAtUpdate {
	hsau.mutation.Where(ps...)
	return hsau
}

// SetPackageVersionID sets the "package_version_id" field.
func (hsau *HasSourceAtUpdate) SetPackageVersionID(i int) *HasSourceAtUpdate {
	hsau.mutation.SetPackageVersionID(i)
	return hsau
}

// SetNillablePackageVersionID sets the "package_version_id" field if the given value is not nil.
func (hsau *HasSourceAtUpdate) SetNillablePackageVersionID(i *int) *HasSourceAtUpdate {
	if i != nil {
		hsau.SetPackageVersionID(*i)
	}
	return hsau
}

// ClearPackageVersionID clears the value of the "package_version_id" field.
func (hsau *HasSourceAtUpdate) ClearPackageVersionID() *HasSourceAtUpdate {
	hsau.mutation.ClearPackageVersionID()
	return hsau
}

// SetPackageNameID sets the "package_name_id" field.
func (hsau *HasSourceAtUpdate) SetPackageNameID(i int) *HasSourceAtUpdate {
	hsau.mutation.SetPackageNameID(i)
	return hsau
}

// SetNillablePackageNameID sets the "package_name_id" field if the given value is not nil.
func (hsau *HasSourceAtUpdate) SetNillablePackageNameID(i *int) *HasSourceAtUpdate {
	if i != nil {
		hsau.SetPackageNameID(*i)
	}
	return hsau
}

// ClearPackageNameID clears the value of the "package_name_id" field.
func (hsau *HasSourceAtUpdate) ClearPackageNameID() *HasSourceAtUpdate {
	hsau.mutation.ClearPackageNameID()
	return hsau
}

// SetSourceID sets the "source_id" field.
func (hsau *HasSourceAtUpdate) SetSourceID(i int) *HasSourceAtUpdate {
	hsau.mutation.SetSourceID(i)
	return hsau
}

// SetKnownSince sets the "known_since" field.
func (hsau *HasSourceAtUpdate) SetKnownSince(t time.Time) *HasSourceAtUpdate {
	hsau.mutation.SetKnownSince(t)
	return hsau
}

// SetJustification sets the "justification" field.
func (hsau *HasSourceAtUpdate) SetJustification(s string) *HasSourceAtUpdate {
	hsau.mutation.SetJustification(s)
	return hsau
}

// SetOrigin sets the "origin" field.
func (hsau *HasSourceAtUpdate) SetOrigin(s string) *HasSourceAtUpdate {
	hsau.mutation.SetOrigin(s)
	return hsau
}

// SetCollector sets the "collector" field.
func (hsau *HasSourceAtUpdate) SetCollector(s string) *HasSourceAtUpdate {
	hsau.mutation.SetCollector(s)
	return hsau
}

// SetPackageVersion sets the "package_version" edge to the PackageVersion entity.
func (hsau *HasSourceAtUpdate) SetPackageVersion(p *PackageVersion) *HasSourceAtUpdate {
	return hsau.SetPackageVersionID(p.ID)
}

// SetAllVersionsID sets the "all_versions" edge to the PackageName entity by ID.
func (hsau *HasSourceAtUpdate) SetAllVersionsID(id int) *HasSourceAtUpdate {
	hsau.mutation.SetAllVersionsID(id)
	return hsau
}

// SetNillableAllVersionsID sets the "all_versions" edge to the PackageName entity by ID if the given value is not nil.
func (hsau *HasSourceAtUpdate) SetNillableAllVersionsID(id *int) *HasSourceAtUpdate {
	if id != nil {
		hsau = hsau.SetAllVersionsID(*id)
	}
	return hsau
}

// SetAllVersions sets the "all_versions" edge to the PackageName entity.
func (hsau *HasSourceAtUpdate) SetAllVersions(p *PackageName) *HasSourceAtUpdate {
	return hsau.SetAllVersionsID(p.ID)
}

// SetSource sets the "source" edge to the SourceName entity.
func (hsau *HasSourceAtUpdate) SetSource(s *SourceName) *HasSourceAtUpdate {
	return hsau.SetSourceID(s.ID)
}

// Mutation returns the HasSourceAtMutation object of the builder.
func (hsau *HasSourceAtUpdate) Mutation() *HasSourceAtMutation {
	return hsau.mutation
}

// ClearPackageVersion clears the "package_version" edge to the PackageVersion entity.
func (hsau *HasSourceAtUpdate) ClearPackageVersion() *HasSourceAtUpdate {
	hsau.mutation.ClearPackageVersion()
	return hsau
}

// ClearAllVersions clears the "all_versions" edge to the PackageName entity.
func (hsau *HasSourceAtUpdate) ClearAllVersions() *HasSourceAtUpdate {
	hsau.mutation.ClearAllVersions()
	return hsau
}

// ClearSource clears the "source" edge to the SourceName entity.
func (hsau *HasSourceAtUpdate) ClearSource() *HasSourceAtUpdate {
	hsau.mutation.ClearSource()
	return hsau
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (hsau *HasSourceAtUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, hsau.sqlSave, hsau.mutation, hsau.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (hsau *HasSourceAtUpdate) SaveX(ctx context.Context) int {
	affected, err := hsau.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (hsau *HasSourceAtUpdate) Exec(ctx context.Context) error {
	_, err := hsau.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (hsau *HasSourceAtUpdate) ExecX(ctx context.Context) {
	if err := hsau.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (hsau *HasSourceAtUpdate) check() error {
	if _, ok := hsau.mutation.SourceID(); hsau.mutation.SourceCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "HasSourceAt.source"`)
	}
	return nil
}

func (hsau *HasSourceAtUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := hsau.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(hassourceat.Table, hassourceat.Columns, sqlgraph.NewFieldSpec(hassourceat.FieldID, field.TypeInt))
	if ps := hsau.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := hsau.mutation.KnownSince(); ok {
		_spec.SetField(hassourceat.FieldKnownSince, field.TypeTime, value)
	}
	if value, ok := hsau.mutation.Justification(); ok {
		_spec.SetField(hassourceat.FieldJustification, field.TypeString, value)
	}
	if value, ok := hsau.mutation.Origin(); ok {
		_spec.SetField(hassourceat.FieldOrigin, field.TypeString, value)
	}
	if value, ok := hsau.mutation.Collector(); ok {
		_spec.SetField(hassourceat.FieldCollector, field.TypeString, value)
	}
	if hsau.mutation.PackageVersionCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   hassourceat.PackageVersionTable,
			Columns: []string{hassourceat.PackageVersionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(packageversion.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := hsau.mutation.PackageVersionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   hassourceat.PackageVersionTable,
			Columns: []string{hassourceat.PackageVersionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(packageversion.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if hsau.mutation.AllVersionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   hassourceat.AllVersionsTable,
			Columns: []string{hassourceat.AllVersionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(packagename.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := hsau.mutation.AllVersionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   hassourceat.AllVersionsTable,
			Columns: []string{hassourceat.AllVersionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(packagename.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if hsau.mutation.SourceCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   hassourceat.SourceTable,
			Columns: []string{hassourceat.SourceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(sourcename.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := hsau.mutation.SourceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   hassourceat.SourceTable,
			Columns: []string{hassourceat.SourceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(sourcename.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, hsau.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{hassourceat.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	hsau.mutation.done = true
	return n, nil
}

// HasSourceAtUpdateOne is the builder for updating a single HasSourceAt entity.
type HasSourceAtUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *HasSourceAtMutation
}

// SetPackageVersionID sets the "package_version_id" field.
func (hsauo *HasSourceAtUpdateOne) SetPackageVersionID(i int) *HasSourceAtUpdateOne {
	hsauo.mutation.SetPackageVersionID(i)
	return hsauo
}

// SetNillablePackageVersionID sets the "package_version_id" field if the given value is not nil.
func (hsauo *HasSourceAtUpdateOne) SetNillablePackageVersionID(i *int) *HasSourceAtUpdateOne {
	if i != nil {
		hsauo.SetPackageVersionID(*i)
	}
	return hsauo
}

// ClearPackageVersionID clears the value of the "package_version_id" field.
func (hsauo *HasSourceAtUpdateOne) ClearPackageVersionID() *HasSourceAtUpdateOne {
	hsauo.mutation.ClearPackageVersionID()
	return hsauo
}

// SetPackageNameID sets the "package_name_id" field.
func (hsauo *HasSourceAtUpdateOne) SetPackageNameID(i int) *HasSourceAtUpdateOne {
	hsauo.mutation.SetPackageNameID(i)
	return hsauo
}

// SetNillablePackageNameID sets the "package_name_id" field if the given value is not nil.
func (hsauo *HasSourceAtUpdateOne) SetNillablePackageNameID(i *int) *HasSourceAtUpdateOne {
	if i != nil {
		hsauo.SetPackageNameID(*i)
	}
	return hsauo
}

// ClearPackageNameID clears the value of the "package_name_id" field.
func (hsauo *HasSourceAtUpdateOne) ClearPackageNameID() *HasSourceAtUpdateOne {
	hsauo.mutation.ClearPackageNameID()
	return hsauo
}

// SetSourceID sets the "source_id" field.
func (hsauo *HasSourceAtUpdateOne) SetSourceID(i int) *HasSourceAtUpdateOne {
	hsauo.mutation.SetSourceID(i)
	return hsauo
}

// SetKnownSince sets the "known_since" field.
func (hsauo *HasSourceAtUpdateOne) SetKnownSince(t time.Time) *HasSourceAtUpdateOne {
	hsauo.mutation.SetKnownSince(t)
	return hsauo
}

// SetJustification sets the "justification" field.
func (hsauo *HasSourceAtUpdateOne) SetJustification(s string) *HasSourceAtUpdateOne {
	hsauo.mutation.SetJustification(s)
	return hsauo
}

// SetOrigin sets the "origin" field.
func (hsauo *HasSourceAtUpdateOne) SetOrigin(s string) *HasSourceAtUpdateOne {
	hsauo.mutation.SetOrigin(s)
	return hsauo
}

// SetCollector sets the "collector" field.
func (hsauo *HasSourceAtUpdateOne) SetCollector(s string) *HasSourceAtUpdateOne {
	hsauo.mutation.SetCollector(s)
	return hsauo
}

// SetPackageVersion sets the "package_version" edge to the PackageVersion entity.
func (hsauo *HasSourceAtUpdateOne) SetPackageVersion(p *PackageVersion) *HasSourceAtUpdateOne {
	return hsauo.SetPackageVersionID(p.ID)
}

// SetAllVersionsID sets the "all_versions" edge to the PackageName entity by ID.
func (hsauo *HasSourceAtUpdateOne) SetAllVersionsID(id int) *HasSourceAtUpdateOne {
	hsauo.mutation.SetAllVersionsID(id)
	return hsauo
}

// SetNillableAllVersionsID sets the "all_versions" edge to the PackageName entity by ID if the given value is not nil.
func (hsauo *HasSourceAtUpdateOne) SetNillableAllVersionsID(id *int) *HasSourceAtUpdateOne {
	if id != nil {
		hsauo = hsauo.SetAllVersionsID(*id)
	}
	return hsauo
}

// SetAllVersions sets the "all_versions" edge to the PackageName entity.
func (hsauo *HasSourceAtUpdateOne) SetAllVersions(p *PackageName) *HasSourceAtUpdateOne {
	return hsauo.SetAllVersionsID(p.ID)
}

// SetSource sets the "source" edge to the SourceName entity.
func (hsauo *HasSourceAtUpdateOne) SetSource(s *SourceName) *HasSourceAtUpdateOne {
	return hsauo.SetSourceID(s.ID)
}

// Mutation returns the HasSourceAtMutation object of the builder.
func (hsauo *HasSourceAtUpdateOne) Mutation() *HasSourceAtMutation {
	return hsauo.mutation
}

// ClearPackageVersion clears the "package_version" edge to the PackageVersion entity.
func (hsauo *HasSourceAtUpdateOne) ClearPackageVersion() *HasSourceAtUpdateOne {
	hsauo.mutation.ClearPackageVersion()
	return hsauo
}

// ClearAllVersions clears the "all_versions" edge to the PackageName entity.
func (hsauo *HasSourceAtUpdateOne) ClearAllVersions() *HasSourceAtUpdateOne {
	hsauo.mutation.ClearAllVersions()
	return hsauo
}

// ClearSource clears the "source" edge to the SourceName entity.
func (hsauo *HasSourceAtUpdateOne) ClearSource() *HasSourceAtUpdateOne {
	hsauo.mutation.ClearSource()
	return hsauo
}

// Where appends a list predicates to the HasSourceAtUpdate builder.
func (hsauo *HasSourceAtUpdateOne) Where(ps ...predicate.HasSourceAt) *HasSourceAtUpdateOne {
	hsauo.mutation.Where(ps...)
	return hsauo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (hsauo *HasSourceAtUpdateOne) Select(field string, fields ...string) *HasSourceAtUpdateOne {
	hsauo.fields = append([]string{field}, fields...)
	return hsauo
}

// Save executes the query and returns the updated HasSourceAt entity.
func (hsauo *HasSourceAtUpdateOne) Save(ctx context.Context) (*HasSourceAt, error) {
	return withHooks(ctx, hsauo.sqlSave, hsauo.mutation, hsauo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (hsauo *HasSourceAtUpdateOne) SaveX(ctx context.Context) *HasSourceAt {
	node, err := hsauo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (hsauo *HasSourceAtUpdateOne) Exec(ctx context.Context) error {
	_, err := hsauo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (hsauo *HasSourceAtUpdateOne) ExecX(ctx context.Context) {
	if err := hsauo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (hsauo *HasSourceAtUpdateOne) check() error {
	if _, ok := hsauo.mutation.SourceID(); hsauo.mutation.SourceCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "HasSourceAt.source"`)
	}
	return nil
}

func (hsauo *HasSourceAtUpdateOne) sqlSave(ctx context.Context) (_node *HasSourceAt, err error) {
	if err := hsauo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(hassourceat.Table, hassourceat.Columns, sqlgraph.NewFieldSpec(hassourceat.FieldID, field.TypeInt))
	id, ok := hsauo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "HasSourceAt.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := hsauo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, hassourceat.FieldID)
		for _, f := range fields {
			if !hassourceat.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != hassourceat.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := hsauo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := hsauo.mutation.KnownSince(); ok {
		_spec.SetField(hassourceat.FieldKnownSince, field.TypeTime, value)
	}
	if value, ok := hsauo.mutation.Justification(); ok {
		_spec.SetField(hassourceat.FieldJustification, field.TypeString, value)
	}
	if value, ok := hsauo.mutation.Origin(); ok {
		_spec.SetField(hassourceat.FieldOrigin, field.TypeString, value)
	}
	if value, ok := hsauo.mutation.Collector(); ok {
		_spec.SetField(hassourceat.FieldCollector, field.TypeString, value)
	}
	if hsauo.mutation.PackageVersionCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   hassourceat.PackageVersionTable,
			Columns: []string{hassourceat.PackageVersionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(packageversion.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := hsauo.mutation.PackageVersionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   hassourceat.PackageVersionTable,
			Columns: []string{hassourceat.PackageVersionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(packageversion.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if hsauo.mutation.AllVersionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   hassourceat.AllVersionsTable,
			Columns: []string{hassourceat.AllVersionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(packagename.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := hsauo.mutation.AllVersionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   hassourceat.AllVersionsTable,
			Columns: []string{hassourceat.AllVersionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(packagename.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if hsauo.mutation.SourceCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   hassourceat.SourceTable,
			Columns: []string{hassourceat.SourceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(sourcename.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := hsauo.mutation.SourceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   hassourceat.SourceTable,
			Columns: []string{hassourceat.SourceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(sourcename.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &HasSourceAt{config: hsauo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, hsauo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{hassourceat.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	hsauo.mutation.done = true
	return _node, nil
}
