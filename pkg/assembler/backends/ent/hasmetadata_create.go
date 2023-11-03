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
	"github.com/guacsec/guac/pkg/assembler/backends/ent/artifact"
	"github.com/guacsec/guac/pkg/assembler/backends/ent/hasmetadata"
	"github.com/guacsec/guac/pkg/assembler/backends/ent/packagename"
	"github.com/guacsec/guac/pkg/assembler/backends/ent/packageversion"
	"github.com/guacsec/guac/pkg/assembler/backends/ent/sourcename"
)

// HasMetadataCreate is the builder for creating a HasMetadata entity.
type HasMetadataCreate struct {
	config
	mutation *HasMetadataMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetSourceID sets the "source_id" field.
func (hmc *HasMetadataCreate) SetSourceID(i int) *HasMetadataCreate {
	hmc.mutation.SetSourceID(i)
	return hmc
}

// SetNillableSourceID sets the "source_id" field if the given value is not nil.
func (hmc *HasMetadataCreate) SetNillableSourceID(i *int) *HasMetadataCreate {
	if i != nil {
		hmc.SetSourceID(*i)
	}
	return hmc
}

// SetPackageVersionID sets the "package_version_id" field.
func (hmc *HasMetadataCreate) SetPackageVersionID(i int) *HasMetadataCreate {
	hmc.mutation.SetPackageVersionID(i)
	return hmc
}

// SetNillablePackageVersionID sets the "package_version_id" field if the given value is not nil.
func (hmc *HasMetadataCreate) SetNillablePackageVersionID(i *int) *HasMetadataCreate {
	if i != nil {
		hmc.SetPackageVersionID(*i)
	}
	return hmc
}

// SetPackageNameID sets the "package_name_id" field.
func (hmc *HasMetadataCreate) SetPackageNameID(i int) *HasMetadataCreate {
	hmc.mutation.SetPackageNameID(i)
	return hmc
}

// SetNillablePackageNameID sets the "package_name_id" field if the given value is not nil.
func (hmc *HasMetadataCreate) SetNillablePackageNameID(i *int) *HasMetadataCreate {
	if i != nil {
		hmc.SetPackageNameID(*i)
	}
	return hmc
}

// SetArtifactID sets the "artifact_id" field.
func (hmc *HasMetadataCreate) SetArtifactID(i int) *HasMetadataCreate {
	hmc.mutation.SetArtifactID(i)
	return hmc
}

// SetNillableArtifactID sets the "artifact_id" field if the given value is not nil.
func (hmc *HasMetadataCreate) SetNillableArtifactID(i *int) *HasMetadataCreate {
	if i != nil {
		hmc.SetArtifactID(*i)
	}
	return hmc
}

// SetTimestamp sets the "timestamp" field.
func (hmc *HasMetadataCreate) SetTimestamp(t time.Time) *HasMetadataCreate {
	hmc.mutation.SetTimestamp(t)
	return hmc
}

// SetKey sets the "key" field.
func (hmc *HasMetadataCreate) SetKey(s string) *HasMetadataCreate {
	hmc.mutation.SetKey(s)
	return hmc
}

// SetValue sets the "value" field.
func (hmc *HasMetadataCreate) SetValue(s string) *HasMetadataCreate {
	hmc.mutation.SetValue(s)
	return hmc
}

// SetJustification sets the "justification" field.
func (hmc *HasMetadataCreate) SetJustification(s string) *HasMetadataCreate {
	hmc.mutation.SetJustification(s)
	return hmc
}

// SetOrigin sets the "origin" field.
func (hmc *HasMetadataCreate) SetOrigin(s string) *HasMetadataCreate {
	hmc.mutation.SetOrigin(s)
	return hmc
}

// SetCollector sets the "collector" field.
func (hmc *HasMetadataCreate) SetCollector(s string) *HasMetadataCreate {
	hmc.mutation.SetCollector(s)
	return hmc
}

// SetSource sets the "source" edge to the SourceName entity.
func (hmc *HasMetadataCreate) SetSource(s *SourceName) *HasMetadataCreate {
	return hmc.SetSourceID(s.ID)
}

// SetPackageVersion sets the "package_version" edge to the PackageVersion entity.
func (hmc *HasMetadataCreate) SetPackageVersion(p *PackageVersion) *HasMetadataCreate {
	return hmc.SetPackageVersionID(p.ID)
}

// SetAllVersionsID sets the "all_versions" edge to the PackageName entity by ID.
func (hmc *HasMetadataCreate) SetAllVersionsID(id int) *HasMetadataCreate {
	hmc.mutation.SetAllVersionsID(id)
	return hmc
}

// SetNillableAllVersionsID sets the "all_versions" edge to the PackageName entity by ID if the given value is not nil.
func (hmc *HasMetadataCreate) SetNillableAllVersionsID(id *int) *HasMetadataCreate {
	if id != nil {
		hmc = hmc.SetAllVersionsID(*id)
	}
	return hmc
}

// SetAllVersions sets the "all_versions" edge to the PackageName entity.
func (hmc *HasMetadataCreate) SetAllVersions(p *PackageName) *HasMetadataCreate {
	return hmc.SetAllVersionsID(p.ID)
}

// SetArtifact sets the "artifact" edge to the Artifact entity.
func (hmc *HasMetadataCreate) SetArtifact(a *Artifact) *HasMetadataCreate {
	return hmc.SetArtifactID(a.ID)
}

// Mutation returns the HasMetadataMutation object of the builder.
func (hmc *HasMetadataCreate) Mutation() *HasMetadataMutation {
	return hmc.mutation
}

// Save creates the HasMetadata in the database.
func (hmc *HasMetadataCreate) Save(ctx context.Context) (*HasMetadata, error) {
	return withHooks(ctx, hmc.sqlSave, hmc.mutation, hmc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (hmc *HasMetadataCreate) SaveX(ctx context.Context) *HasMetadata {
	v, err := hmc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (hmc *HasMetadataCreate) Exec(ctx context.Context) error {
	_, err := hmc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (hmc *HasMetadataCreate) ExecX(ctx context.Context) {
	if err := hmc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (hmc *HasMetadataCreate) check() error {
	if _, ok := hmc.mutation.Timestamp(); !ok {
		return &ValidationError{Name: "timestamp", err: errors.New(`ent: missing required field "HasMetadata.timestamp"`)}
	}
	if _, ok := hmc.mutation.Key(); !ok {
		return &ValidationError{Name: "key", err: errors.New(`ent: missing required field "HasMetadata.key"`)}
	}
	if _, ok := hmc.mutation.Value(); !ok {
		return &ValidationError{Name: "value", err: errors.New(`ent: missing required field "HasMetadata.value"`)}
	}
	if _, ok := hmc.mutation.Justification(); !ok {
		return &ValidationError{Name: "justification", err: errors.New(`ent: missing required field "HasMetadata.justification"`)}
	}
	if _, ok := hmc.mutation.Origin(); !ok {
		return &ValidationError{Name: "origin", err: errors.New(`ent: missing required field "HasMetadata.origin"`)}
	}
	if _, ok := hmc.mutation.Collector(); !ok {
		return &ValidationError{Name: "collector", err: errors.New(`ent: missing required field "HasMetadata.collector"`)}
	}
	return nil
}

func (hmc *HasMetadataCreate) sqlSave(ctx context.Context) (*HasMetadata, error) {
	if err := hmc.check(); err != nil {
		return nil, err
	}
	_node, _spec := hmc.createSpec()
	if err := sqlgraph.CreateNode(ctx, hmc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	hmc.mutation.id = &_node.ID
	hmc.mutation.done = true
	return _node, nil
}

func (hmc *HasMetadataCreate) createSpec() (*HasMetadata, *sqlgraph.CreateSpec) {
	var (
		_node = &HasMetadata{config: hmc.config}
		_spec = sqlgraph.NewCreateSpec(hasmetadata.Table, sqlgraph.NewFieldSpec(hasmetadata.FieldID, field.TypeInt))
	)
	_spec.OnConflict = hmc.conflict
	if value, ok := hmc.mutation.Timestamp(); ok {
		_spec.SetField(hasmetadata.FieldTimestamp, field.TypeTime, value)
		_node.Timestamp = value
	}
	if value, ok := hmc.mutation.Key(); ok {
		_spec.SetField(hasmetadata.FieldKey, field.TypeString, value)
		_node.Key = value
	}
	if value, ok := hmc.mutation.Value(); ok {
		_spec.SetField(hasmetadata.FieldValue, field.TypeString, value)
		_node.Value = value
	}
	if value, ok := hmc.mutation.Justification(); ok {
		_spec.SetField(hasmetadata.FieldJustification, field.TypeString, value)
		_node.Justification = value
	}
	if value, ok := hmc.mutation.Origin(); ok {
		_spec.SetField(hasmetadata.FieldOrigin, field.TypeString, value)
		_node.Origin = value
	}
	if value, ok := hmc.mutation.Collector(); ok {
		_spec.SetField(hasmetadata.FieldCollector, field.TypeString, value)
		_node.Collector = value
	}
	if nodes := hmc.mutation.SourceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   hasmetadata.SourceTable,
			Columns: []string{hasmetadata.SourceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(sourcename.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.SourceID = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := hmc.mutation.PackageVersionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   hasmetadata.PackageVersionTable,
			Columns: []string{hasmetadata.PackageVersionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(packageversion.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.PackageVersionID = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := hmc.mutation.AllVersionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   hasmetadata.AllVersionsTable,
			Columns: []string{hasmetadata.AllVersionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(packagename.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.PackageNameID = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := hmc.mutation.ArtifactIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   hasmetadata.ArtifactTable,
			Columns: []string{hasmetadata.ArtifactColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(artifact.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.ArtifactID = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.HasMetadata.Create().
//		SetSourceID(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.HasMetadataUpsert) {
//			SetSourceID(v+v).
//		}).
//		Exec(ctx)
func (hmc *HasMetadataCreate) OnConflict(opts ...sql.ConflictOption) *HasMetadataUpsertOne {
	hmc.conflict = opts
	return &HasMetadataUpsertOne{
		create: hmc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.HasMetadata.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (hmc *HasMetadataCreate) OnConflictColumns(columns ...string) *HasMetadataUpsertOne {
	hmc.conflict = append(hmc.conflict, sql.ConflictColumns(columns...))
	return &HasMetadataUpsertOne{
		create: hmc,
	}
}

type (
	// HasMetadataUpsertOne is the builder for "upsert"-ing
	//  one HasMetadata node.
	HasMetadataUpsertOne struct {
		create *HasMetadataCreate
	}

	// HasMetadataUpsert is the "OnConflict" setter.
	HasMetadataUpsert struct {
		*sql.UpdateSet
	}
)

// SetSourceID sets the "source_id" field.
func (u *HasMetadataUpsert) SetSourceID(v int) *HasMetadataUpsert {
	u.Set(hasmetadata.FieldSourceID, v)
	return u
}

// UpdateSourceID sets the "source_id" field to the value that was provided on create.
func (u *HasMetadataUpsert) UpdateSourceID() *HasMetadataUpsert {
	u.SetExcluded(hasmetadata.FieldSourceID)
	return u
}

// ClearSourceID clears the value of the "source_id" field.
func (u *HasMetadataUpsert) ClearSourceID() *HasMetadataUpsert {
	u.SetNull(hasmetadata.FieldSourceID)
	return u
}

// SetPackageVersionID sets the "package_version_id" field.
func (u *HasMetadataUpsert) SetPackageVersionID(v int) *HasMetadataUpsert {
	u.Set(hasmetadata.FieldPackageVersionID, v)
	return u
}

// UpdatePackageVersionID sets the "package_version_id" field to the value that was provided on create.
func (u *HasMetadataUpsert) UpdatePackageVersionID() *HasMetadataUpsert {
	u.SetExcluded(hasmetadata.FieldPackageVersionID)
	return u
}

// ClearPackageVersionID clears the value of the "package_version_id" field.
func (u *HasMetadataUpsert) ClearPackageVersionID() *HasMetadataUpsert {
	u.SetNull(hasmetadata.FieldPackageVersionID)
	return u
}

// SetPackageNameID sets the "package_name_id" field.
func (u *HasMetadataUpsert) SetPackageNameID(v int) *HasMetadataUpsert {
	u.Set(hasmetadata.FieldPackageNameID, v)
	return u
}

// UpdatePackageNameID sets the "package_name_id" field to the value that was provided on create.
func (u *HasMetadataUpsert) UpdatePackageNameID() *HasMetadataUpsert {
	u.SetExcluded(hasmetadata.FieldPackageNameID)
	return u
}

// ClearPackageNameID clears the value of the "package_name_id" field.
func (u *HasMetadataUpsert) ClearPackageNameID() *HasMetadataUpsert {
	u.SetNull(hasmetadata.FieldPackageNameID)
	return u
}

// SetArtifactID sets the "artifact_id" field.
func (u *HasMetadataUpsert) SetArtifactID(v int) *HasMetadataUpsert {
	u.Set(hasmetadata.FieldArtifactID, v)
	return u
}

// UpdateArtifactID sets the "artifact_id" field to the value that was provided on create.
func (u *HasMetadataUpsert) UpdateArtifactID() *HasMetadataUpsert {
	u.SetExcluded(hasmetadata.FieldArtifactID)
	return u
}

// ClearArtifactID clears the value of the "artifact_id" field.
func (u *HasMetadataUpsert) ClearArtifactID() *HasMetadataUpsert {
	u.SetNull(hasmetadata.FieldArtifactID)
	return u
}

// SetTimestamp sets the "timestamp" field.
func (u *HasMetadataUpsert) SetTimestamp(v time.Time) *HasMetadataUpsert {
	u.Set(hasmetadata.FieldTimestamp, v)
	return u
}

// UpdateTimestamp sets the "timestamp" field to the value that was provided on create.
func (u *HasMetadataUpsert) UpdateTimestamp() *HasMetadataUpsert {
	u.SetExcluded(hasmetadata.FieldTimestamp)
	return u
}

// SetKey sets the "key" field.
func (u *HasMetadataUpsert) SetKey(v string) *HasMetadataUpsert {
	u.Set(hasmetadata.FieldKey, v)
	return u
}

// UpdateKey sets the "key" field to the value that was provided on create.
func (u *HasMetadataUpsert) UpdateKey() *HasMetadataUpsert {
	u.SetExcluded(hasmetadata.FieldKey)
	return u
}

// SetValue sets the "value" field.
func (u *HasMetadataUpsert) SetValue(v string) *HasMetadataUpsert {
	u.Set(hasmetadata.FieldValue, v)
	return u
}

// UpdateValue sets the "value" field to the value that was provided on create.
func (u *HasMetadataUpsert) UpdateValue() *HasMetadataUpsert {
	u.SetExcluded(hasmetadata.FieldValue)
	return u
}

// SetJustification sets the "justification" field.
func (u *HasMetadataUpsert) SetJustification(v string) *HasMetadataUpsert {
	u.Set(hasmetadata.FieldJustification, v)
	return u
}

// UpdateJustification sets the "justification" field to the value that was provided on create.
func (u *HasMetadataUpsert) UpdateJustification() *HasMetadataUpsert {
	u.SetExcluded(hasmetadata.FieldJustification)
	return u
}

// SetOrigin sets the "origin" field.
func (u *HasMetadataUpsert) SetOrigin(v string) *HasMetadataUpsert {
	u.Set(hasmetadata.FieldOrigin, v)
	return u
}

// UpdateOrigin sets the "origin" field to the value that was provided on create.
func (u *HasMetadataUpsert) UpdateOrigin() *HasMetadataUpsert {
	u.SetExcluded(hasmetadata.FieldOrigin)
	return u
}

// SetCollector sets the "collector" field.
func (u *HasMetadataUpsert) SetCollector(v string) *HasMetadataUpsert {
	u.Set(hasmetadata.FieldCollector, v)
	return u
}

// UpdateCollector sets the "collector" field to the value that was provided on create.
func (u *HasMetadataUpsert) UpdateCollector() *HasMetadataUpsert {
	u.SetExcluded(hasmetadata.FieldCollector)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.HasMetadata.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *HasMetadataUpsertOne) UpdateNewValues() *HasMetadataUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.HasMetadata.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *HasMetadataUpsertOne) Ignore() *HasMetadataUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *HasMetadataUpsertOne) DoNothing() *HasMetadataUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the HasMetadataCreate.OnConflict
// documentation for more info.
func (u *HasMetadataUpsertOne) Update(set func(*HasMetadataUpsert)) *HasMetadataUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&HasMetadataUpsert{UpdateSet: update})
	}))
	return u
}

// SetSourceID sets the "source_id" field.
func (u *HasMetadataUpsertOne) SetSourceID(v int) *HasMetadataUpsertOne {
	return u.Update(func(s *HasMetadataUpsert) {
		s.SetSourceID(v)
	})
}

// UpdateSourceID sets the "source_id" field to the value that was provided on create.
func (u *HasMetadataUpsertOne) UpdateSourceID() *HasMetadataUpsertOne {
	return u.Update(func(s *HasMetadataUpsert) {
		s.UpdateSourceID()
	})
}

// ClearSourceID clears the value of the "source_id" field.
func (u *HasMetadataUpsertOne) ClearSourceID() *HasMetadataUpsertOne {
	return u.Update(func(s *HasMetadataUpsert) {
		s.ClearSourceID()
	})
}

// SetPackageVersionID sets the "package_version_id" field.
func (u *HasMetadataUpsertOne) SetPackageVersionID(v int) *HasMetadataUpsertOne {
	return u.Update(func(s *HasMetadataUpsert) {
		s.SetPackageVersionID(v)
	})
}

// UpdatePackageVersionID sets the "package_version_id" field to the value that was provided on create.
func (u *HasMetadataUpsertOne) UpdatePackageVersionID() *HasMetadataUpsertOne {
	return u.Update(func(s *HasMetadataUpsert) {
		s.UpdatePackageVersionID()
	})
}

// ClearPackageVersionID clears the value of the "package_version_id" field.
func (u *HasMetadataUpsertOne) ClearPackageVersionID() *HasMetadataUpsertOne {
	return u.Update(func(s *HasMetadataUpsert) {
		s.ClearPackageVersionID()
	})
}

// SetPackageNameID sets the "package_name_id" field.
func (u *HasMetadataUpsertOne) SetPackageNameID(v int) *HasMetadataUpsertOne {
	return u.Update(func(s *HasMetadataUpsert) {
		s.SetPackageNameID(v)
	})
}

// UpdatePackageNameID sets the "package_name_id" field to the value that was provided on create.
func (u *HasMetadataUpsertOne) UpdatePackageNameID() *HasMetadataUpsertOne {
	return u.Update(func(s *HasMetadataUpsert) {
		s.UpdatePackageNameID()
	})
}

// ClearPackageNameID clears the value of the "package_name_id" field.
func (u *HasMetadataUpsertOne) ClearPackageNameID() *HasMetadataUpsertOne {
	return u.Update(func(s *HasMetadataUpsert) {
		s.ClearPackageNameID()
	})
}

// SetArtifactID sets the "artifact_id" field.
func (u *HasMetadataUpsertOne) SetArtifactID(v int) *HasMetadataUpsertOne {
	return u.Update(func(s *HasMetadataUpsert) {
		s.SetArtifactID(v)
	})
}

// UpdateArtifactID sets the "artifact_id" field to the value that was provided on create.
func (u *HasMetadataUpsertOne) UpdateArtifactID() *HasMetadataUpsertOne {
	return u.Update(func(s *HasMetadataUpsert) {
		s.UpdateArtifactID()
	})
}

// ClearArtifactID clears the value of the "artifact_id" field.
func (u *HasMetadataUpsertOne) ClearArtifactID() *HasMetadataUpsertOne {
	return u.Update(func(s *HasMetadataUpsert) {
		s.ClearArtifactID()
	})
}

// SetTimestamp sets the "timestamp" field.
func (u *HasMetadataUpsertOne) SetTimestamp(v time.Time) *HasMetadataUpsertOne {
	return u.Update(func(s *HasMetadataUpsert) {
		s.SetTimestamp(v)
	})
}

// UpdateTimestamp sets the "timestamp" field to the value that was provided on create.
func (u *HasMetadataUpsertOne) UpdateTimestamp() *HasMetadataUpsertOne {
	return u.Update(func(s *HasMetadataUpsert) {
		s.UpdateTimestamp()
	})
}

// SetKey sets the "key" field.
func (u *HasMetadataUpsertOne) SetKey(v string) *HasMetadataUpsertOne {
	return u.Update(func(s *HasMetadataUpsert) {
		s.SetKey(v)
	})
}

// UpdateKey sets the "key" field to the value that was provided on create.
func (u *HasMetadataUpsertOne) UpdateKey() *HasMetadataUpsertOne {
	return u.Update(func(s *HasMetadataUpsert) {
		s.UpdateKey()
	})
}

// SetValue sets the "value" field.
func (u *HasMetadataUpsertOne) SetValue(v string) *HasMetadataUpsertOne {
	return u.Update(func(s *HasMetadataUpsert) {
		s.SetValue(v)
	})
}

// UpdateValue sets the "value" field to the value that was provided on create.
func (u *HasMetadataUpsertOne) UpdateValue() *HasMetadataUpsertOne {
	return u.Update(func(s *HasMetadataUpsert) {
		s.UpdateValue()
	})
}

// SetJustification sets the "justification" field.
func (u *HasMetadataUpsertOne) SetJustification(v string) *HasMetadataUpsertOne {
	return u.Update(func(s *HasMetadataUpsert) {
		s.SetJustification(v)
	})
}

// UpdateJustification sets the "justification" field to the value that was provided on create.
func (u *HasMetadataUpsertOne) UpdateJustification() *HasMetadataUpsertOne {
	return u.Update(func(s *HasMetadataUpsert) {
		s.UpdateJustification()
	})
}

// SetOrigin sets the "origin" field.
func (u *HasMetadataUpsertOne) SetOrigin(v string) *HasMetadataUpsertOne {
	return u.Update(func(s *HasMetadataUpsert) {
		s.SetOrigin(v)
	})
}

// UpdateOrigin sets the "origin" field to the value that was provided on create.
func (u *HasMetadataUpsertOne) UpdateOrigin() *HasMetadataUpsertOne {
	return u.Update(func(s *HasMetadataUpsert) {
		s.UpdateOrigin()
	})
}

// SetCollector sets the "collector" field.
func (u *HasMetadataUpsertOne) SetCollector(v string) *HasMetadataUpsertOne {
	return u.Update(func(s *HasMetadataUpsert) {
		s.SetCollector(v)
	})
}

// UpdateCollector sets the "collector" field to the value that was provided on create.
func (u *HasMetadataUpsertOne) UpdateCollector() *HasMetadataUpsertOne {
	return u.Update(func(s *HasMetadataUpsert) {
		s.UpdateCollector()
	})
}

// Exec executes the query.
func (u *HasMetadataUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for HasMetadataCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *HasMetadataUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *HasMetadataUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *HasMetadataUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// HasMetadataCreateBulk is the builder for creating many HasMetadata entities in bulk.
type HasMetadataCreateBulk struct {
	config
	err      error
	builders []*HasMetadataCreate
	conflict []sql.ConflictOption
}

// Save creates the HasMetadata entities in the database.
func (hmcb *HasMetadataCreateBulk) Save(ctx context.Context) ([]*HasMetadata, error) {
	if hmcb.err != nil {
		return nil, hmcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(hmcb.builders))
	nodes := make([]*HasMetadata, len(hmcb.builders))
	mutators := make([]Mutator, len(hmcb.builders))
	for i := range hmcb.builders {
		func(i int, root context.Context) {
			builder := hmcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*HasMetadataMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, hmcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = hmcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, hmcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, hmcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (hmcb *HasMetadataCreateBulk) SaveX(ctx context.Context) []*HasMetadata {
	v, err := hmcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (hmcb *HasMetadataCreateBulk) Exec(ctx context.Context) error {
	_, err := hmcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (hmcb *HasMetadataCreateBulk) ExecX(ctx context.Context) {
	if err := hmcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.HasMetadata.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.HasMetadataUpsert) {
//			SetSourceID(v+v).
//		}).
//		Exec(ctx)
func (hmcb *HasMetadataCreateBulk) OnConflict(opts ...sql.ConflictOption) *HasMetadataUpsertBulk {
	hmcb.conflict = opts
	return &HasMetadataUpsertBulk{
		create: hmcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.HasMetadata.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (hmcb *HasMetadataCreateBulk) OnConflictColumns(columns ...string) *HasMetadataUpsertBulk {
	hmcb.conflict = append(hmcb.conflict, sql.ConflictColumns(columns...))
	return &HasMetadataUpsertBulk{
		create: hmcb,
	}
}

// HasMetadataUpsertBulk is the builder for "upsert"-ing
// a bulk of HasMetadata nodes.
type HasMetadataUpsertBulk struct {
	create *HasMetadataCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.HasMetadata.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *HasMetadataUpsertBulk) UpdateNewValues() *HasMetadataUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.HasMetadata.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *HasMetadataUpsertBulk) Ignore() *HasMetadataUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *HasMetadataUpsertBulk) DoNothing() *HasMetadataUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the HasMetadataCreateBulk.OnConflict
// documentation for more info.
func (u *HasMetadataUpsertBulk) Update(set func(*HasMetadataUpsert)) *HasMetadataUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&HasMetadataUpsert{UpdateSet: update})
	}))
	return u
}

// SetSourceID sets the "source_id" field.
func (u *HasMetadataUpsertBulk) SetSourceID(v int) *HasMetadataUpsertBulk {
	return u.Update(func(s *HasMetadataUpsert) {
		s.SetSourceID(v)
	})
}

// UpdateSourceID sets the "source_id" field to the value that was provided on create.
func (u *HasMetadataUpsertBulk) UpdateSourceID() *HasMetadataUpsertBulk {
	return u.Update(func(s *HasMetadataUpsert) {
		s.UpdateSourceID()
	})
}

// ClearSourceID clears the value of the "source_id" field.
func (u *HasMetadataUpsertBulk) ClearSourceID() *HasMetadataUpsertBulk {
	return u.Update(func(s *HasMetadataUpsert) {
		s.ClearSourceID()
	})
}

// SetPackageVersionID sets the "package_version_id" field.
func (u *HasMetadataUpsertBulk) SetPackageVersionID(v int) *HasMetadataUpsertBulk {
	return u.Update(func(s *HasMetadataUpsert) {
		s.SetPackageVersionID(v)
	})
}

// UpdatePackageVersionID sets the "package_version_id" field to the value that was provided on create.
func (u *HasMetadataUpsertBulk) UpdatePackageVersionID() *HasMetadataUpsertBulk {
	return u.Update(func(s *HasMetadataUpsert) {
		s.UpdatePackageVersionID()
	})
}

// ClearPackageVersionID clears the value of the "package_version_id" field.
func (u *HasMetadataUpsertBulk) ClearPackageVersionID() *HasMetadataUpsertBulk {
	return u.Update(func(s *HasMetadataUpsert) {
		s.ClearPackageVersionID()
	})
}

// SetPackageNameID sets the "package_name_id" field.
func (u *HasMetadataUpsertBulk) SetPackageNameID(v int) *HasMetadataUpsertBulk {
	return u.Update(func(s *HasMetadataUpsert) {
		s.SetPackageNameID(v)
	})
}

// UpdatePackageNameID sets the "package_name_id" field to the value that was provided on create.
func (u *HasMetadataUpsertBulk) UpdatePackageNameID() *HasMetadataUpsertBulk {
	return u.Update(func(s *HasMetadataUpsert) {
		s.UpdatePackageNameID()
	})
}

// ClearPackageNameID clears the value of the "package_name_id" field.
func (u *HasMetadataUpsertBulk) ClearPackageNameID() *HasMetadataUpsertBulk {
	return u.Update(func(s *HasMetadataUpsert) {
		s.ClearPackageNameID()
	})
}

// SetArtifactID sets the "artifact_id" field.
func (u *HasMetadataUpsertBulk) SetArtifactID(v int) *HasMetadataUpsertBulk {
	return u.Update(func(s *HasMetadataUpsert) {
		s.SetArtifactID(v)
	})
}

// UpdateArtifactID sets the "artifact_id" field to the value that was provided on create.
func (u *HasMetadataUpsertBulk) UpdateArtifactID() *HasMetadataUpsertBulk {
	return u.Update(func(s *HasMetadataUpsert) {
		s.UpdateArtifactID()
	})
}

// ClearArtifactID clears the value of the "artifact_id" field.
func (u *HasMetadataUpsertBulk) ClearArtifactID() *HasMetadataUpsertBulk {
	return u.Update(func(s *HasMetadataUpsert) {
		s.ClearArtifactID()
	})
}

// SetTimestamp sets the "timestamp" field.
func (u *HasMetadataUpsertBulk) SetTimestamp(v time.Time) *HasMetadataUpsertBulk {
	return u.Update(func(s *HasMetadataUpsert) {
		s.SetTimestamp(v)
	})
}

// UpdateTimestamp sets the "timestamp" field to the value that was provided on create.
func (u *HasMetadataUpsertBulk) UpdateTimestamp() *HasMetadataUpsertBulk {
	return u.Update(func(s *HasMetadataUpsert) {
		s.UpdateTimestamp()
	})
}

// SetKey sets the "key" field.
func (u *HasMetadataUpsertBulk) SetKey(v string) *HasMetadataUpsertBulk {
	return u.Update(func(s *HasMetadataUpsert) {
		s.SetKey(v)
	})
}

// UpdateKey sets the "key" field to the value that was provided on create.
func (u *HasMetadataUpsertBulk) UpdateKey() *HasMetadataUpsertBulk {
	return u.Update(func(s *HasMetadataUpsert) {
		s.UpdateKey()
	})
}

// SetValue sets the "value" field.
func (u *HasMetadataUpsertBulk) SetValue(v string) *HasMetadataUpsertBulk {
	return u.Update(func(s *HasMetadataUpsert) {
		s.SetValue(v)
	})
}

// UpdateValue sets the "value" field to the value that was provided on create.
func (u *HasMetadataUpsertBulk) UpdateValue() *HasMetadataUpsertBulk {
	return u.Update(func(s *HasMetadataUpsert) {
		s.UpdateValue()
	})
}

// SetJustification sets the "justification" field.
func (u *HasMetadataUpsertBulk) SetJustification(v string) *HasMetadataUpsertBulk {
	return u.Update(func(s *HasMetadataUpsert) {
		s.SetJustification(v)
	})
}

// UpdateJustification sets the "justification" field to the value that was provided on create.
func (u *HasMetadataUpsertBulk) UpdateJustification() *HasMetadataUpsertBulk {
	return u.Update(func(s *HasMetadataUpsert) {
		s.UpdateJustification()
	})
}

// SetOrigin sets the "origin" field.
func (u *HasMetadataUpsertBulk) SetOrigin(v string) *HasMetadataUpsertBulk {
	return u.Update(func(s *HasMetadataUpsert) {
		s.SetOrigin(v)
	})
}

// UpdateOrigin sets the "origin" field to the value that was provided on create.
func (u *HasMetadataUpsertBulk) UpdateOrigin() *HasMetadataUpsertBulk {
	return u.Update(func(s *HasMetadataUpsert) {
		s.UpdateOrigin()
	})
}

// SetCollector sets the "collector" field.
func (u *HasMetadataUpsertBulk) SetCollector(v string) *HasMetadataUpsertBulk {
	return u.Update(func(s *HasMetadataUpsert) {
		s.SetCollector(v)
	})
}

// UpdateCollector sets the "collector" field to the value that was provided on create.
func (u *HasMetadataUpsertBulk) UpdateCollector() *HasMetadataUpsertBulk {
	return u.Update(func(s *HasMetadataUpsert) {
		s.UpdateCollector()
	})
}

// Exec executes the query.
func (u *HasMetadataUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the HasMetadataCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for HasMetadataCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *HasMetadataUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
