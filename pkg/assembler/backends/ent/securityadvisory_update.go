// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"

	"entgo.io/ent/dialect/gremlin"
	"entgo.io/ent/dialect/gremlin/graph/dsl"
	"entgo.io/ent/dialect/gremlin/graph/dsl/__"
	"entgo.io/ent/dialect/gremlin/graph/dsl/g"
	"github.com/guacsec/guac/pkg/assembler/backends/ent/predicate"
	"github.com/guacsec/guac/pkg/assembler/backends/ent/securityadvisory"
)

// SecurityAdvisoryUpdate is the builder for updating SecurityAdvisory entities.
type SecurityAdvisoryUpdate struct {
	config
	hooks    []Hook
	mutation *SecurityAdvisoryMutation
}

// Where appends a list predicates to the SecurityAdvisoryUpdate builder.
func (sau *SecurityAdvisoryUpdate) Where(ps ...predicate.SecurityAdvisory) *SecurityAdvisoryUpdate {
	sau.mutation.Where(ps...)
	return sau
}

// SetGhsaID sets the "ghsa_id" field.
func (sau *SecurityAdvisoryUpdate) SetGhsaID(s string) *SecurityAdvisoryUpdate {
	sau.mutation.SetGhsaID(s)
	return sau
}

// SetNillableGhsaID sets the "ghsa_id" field if the given value is not nil.
func (sau *SecurityAdvisoryUpdate) SetNillableGhsaID(s *string) *SecurityAdvisoryUpdate {
	if s != nil {
		sau.SetGhsaID(*s)
	}
	return sau
}

// ClearGhsaID clears the value of the "ghsa_id" field.
func (sau *SecurityAdvisoryUpdate) ClearGhsaID() *SecurityAdvisoryUpdate {
	sau.mutation.ClearGhsaID()
	return sau
}

// SetCveID sets the "cve_id" field.
func (sau *SecurityAdvisoryUpdate) SetCveID(s string) *SecurityAdvisoryUpdate {
	sau.mutation.SetCveID(s)
	return sau
}

// SetNillableCveID sets the "cve_id" field if the given value is not nil.
func (sau *SecurityAdvisoryUpdate) SetNillableCveID(s *string) *SecurityAdvisoryUpdate {
	if s != nil {
		sau.SetCveID(*s)
	}
	return sau
}

// ClearCveID clears the value of the "cve_id" field.
func (sau *SecurityAdvisoryUpdate) ClearCveID() *SecurityAdvisoryUpdate {
	sau.mutation.ClearCveID()
	return sau
}

// SetCveYear sets the "cve_year" field.
func (sau *SecurityAdvisoryUpdate) SetCveYear(i int) *SecurityAdvisoryUpdate {
	sau.mutation.ResetCveYear()
	sau.mutation.SetCveYear(i)
	return sau
}

// SetNillableCveYear sets the "cve_year" field if the given value is not nil.
func (sau *SecurityAdvisoryUpdate) SetNillableCveYear(i *int) *SecurityAdvisoryUpdate {
	if i != nil {
		sau.SetCveYear(*i)
	}
	return sau
}

// AddCveYear adds i to the "cve_year" field.
func (sau *SecurityAdvisoryUpdate) AddCveYear(i int) *SecurityAdvisoryUpdate {
	sau.mutation.AddCveYear(i)
	return sau
}

// ClearCveYear clears the value of the "cve_year" field.
func (sau *SecurityAdvisoryUpdate) ClearCveYear() *SecurityAdvisoryUpdate {
	sau.mutation.ClearCveYear()
	return sau
}

// SetOsvID sets the "osv_id" field.
func (sau *SecurityAdvisoryUpdate) SetOsvID(s string) *SecurityAdvisoryUpdate {
	sau.mutation.SetOsvID(s)
	return sau
}

// SetNillableOsvID sets the "osv_id" field if the given value is not nil.
func (sau *SecurityAdvisoryUpdate) SetNillableOsvID(s *string) *SecurityAdvisoryUpdate {
	if s != nil {
		sau.SetOsvID(*s)
	}
	return sau
}

// ClearOsvID clears the value of the "osv_id" field.
func (sau *SecurityAdvisoryUpdate) ClearOsvID() *SecurityAdvisoryUpdate {
	sau.mutation.ClearOsvID()
	return sau
}

// Mutation returns the SecurityAdvisoryMutation object of the builder.
func (sau *SecurityAdvisoryUpdate) Mutation() *SecurityAdvisoryMutation {
	return sau.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (sau *SecurityAdvisoryUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, sau.gremlinSave, sau.mutation, sau.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (sau *SecurityAdvisoryUpdate) SaveX(ctx context.Context) int {
	affected, err := sau.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (sau *SecurityAdvisoryUpdate) Exec(ctx context.Context) error {
	_, err := sau.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sau *SecurityAdvisoryUpdate) ExecX(ctx context.Context) {
	if err := sau.Exec(ctx); err != nil {
		panic(err)
	}
}

func (sau *SecurityAdvisoryUpdate) gremlinSave(ctx context.Context) (int, error) {
	res := &gremlin.Response{}
	query, bindings := sau.gremlin().Query()
	if err := sau.driver.Exec(ctx, query, bindings, res); err != nil {
		return 0, err
	}
	if err, ok := isConstantError(res); ok {
		return 0, err
	}
	sau.mutation.done = true
	return res.ReadInt()
}

func (sau *SecurityAdvisoryUpdate) gremlin() *dsl.Traversal {
	v := g.V().HasLabel(securityadvisory.Label)
	for _, p := range sau.mutation.predicates {
		p(v)
	}
	var (
		trs []*dsl.Traversal
	)
	if value, ok := sau.mutation.GhsaID(); ok {
		v.Property(dsl.Single, securityadvisory.FieldGhsaID, value)
	}
	if value, ok := sau.mutation.CveID(); ok {
		v.Property(dsl.Single, securityadvisory.FieldCveID, value)
	}
	if value, ok := sau.mutation.CveYear(); ok {
		v.Property(dsl.Single, securityadvisory.FieldCveYear, value)
	}
	if value, ok := sau.mutation.AddedCveYear(); ok {
		v.Property(dsl.Single, securityadvisory.FieldCveYear, __.Union(__.Values(securityadvisory.FieldCveYear), __.Constant(value)).Sum())
	}
	if value, ok := sau.mutation.OsvID(); ok {
		v.Property(dsl.Single, securityadvisory.FieldOsvID, value)
	}
	var properties []any
	if sau.mutation.GhsaIDCleared() {
		properties = append(properties, securityadvisory.FieldGhsaID)
	}
	if sau.mutation.CveIDCleared() {
		properties = append(properties, securityadvisory.FieldCveID)
	}
	if sau.mutation.CveYearCleared() {
		properties = append(properties, securityadvisory.FieldCveYear)
	}
	if sau.mutation.OsvIDCleared() {
		properties = append(properties, securityadvisory.FieldOsvID)
	}
	if len(properties) > 0 {
		v.SideEffect(__.Properties(properties...).Drop())
	}
	v.Count()
	trs = append(trs, v)
	return dsl.Join(trs...)
}

// SecurityAdvisoryUpdateOne is the builder for updating a single SecurityAdvisory entity.
type SecurityAdvisoryUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *SecurityAdvisoryMutation
}

// SetGhsaID sets the "ghsa_id" field.
func (sauo *SecurityAdvisoryUpdateOne) SetGhsaID(s string) *SecurityAdvisoryUpdateOne {
	sauo.mutation.SetGhsaID(s)
	return sauo
}

// SetNillableGhsaID sets the "ghsa_id" field if the given value is not nil.
func (sauo *SecurityAdvisoryUpdateOne) SetNillableGhsaID(s *string) *SecurityAdvisoryUpdateOne {
	if s != nil {
		sauo.SetGhsaID(*s)
	}
	return sauo
}

// ClearGhsaID clears the value of the "ghsa_id" field.
func (sauo *SecurityAdvisoryUpdateOne) ClearGhsaID() *SecurityAdvisoryUpdateOne {
	sauo.mutation.ClearGhsaID()
	return sauo
}

// SetCveID sets the "cve_id" field.
func (sauo *SecurityAdvisoryUpdateOne) SetCveID(s string) *SecurityAdvisoryUpdateOne {
	sauo.mutation.SetCveID(s)
	return sauo
}

// SetNillableCveID sets the "cve_id" field if the given value is not nil.
func (sauo *SecurityAdvisoryUpdateOne) SetNillableCveID(s *string) *SecurityAdvisoryUpdateOne {
	if s != nil {
		sauo.SetCveID(*s)
	}
	return sauo
}

// ClearCveID clears the value of the "cve_id" field.
func (sauo *SecurityAdvisoryUpdateOne) ClearCveID() *SecurityAdvisoryUpdateOne {
	sauo.mutation.ClearCveID()
	return sauo
}

// SetCveYear sets the "cve_year" field.
func (sauo *SecurityAdvisoryUpdateOne) SetCveYear(i int) *SecurityAdvisoryUpdateOne {
	sauo.mutation.ResetCveYear()
	sauo.mutation.SetCveYear(i)
	return sauo
}

// SetNillableCveYear sets the "cve_year" field if the given value is not nil.
func (sauo *SecurityAdvisoryUpdateOne) SetNillableCveYear(i *int) *SecurityAdvisoryUpdateOne {
	if i != nil {
		sauo.SetCveYear(*i)
	}
	return sauo
}

// AddCveYear adds i to the "cve_year" field.
func (sauo *SecurityAdvisoryUpdateOne) AddCveYear(i int) *SecurityAdvisoryUpdateOne {
	sauo.mutation.AddCveYear(i)
	return sauo
}

// ClearCveYear clears the value of the "cve_year" field.
func (sauo *SecurityAdvisoryUpdateOne) ClearCveYear() *SecurityAdvisoryUpdateOne {
	sauo.mutation.ClearCveYear()
	return sauo
}

// SetOsvID sets the "osv_id" field.
func (sauo *SecurityAdvisoryUpdateOne) SetOsvID(s string) *SecurityAdvisoryUpdateOne {
	sauo.mutation.SetOsvID(s)
	return sauo
}

// SetNillableOsvID sets the "osv_id" field if the given value is not nil.
func (sauo *SecurityAdvisoryUpdateOne) SetNillableOsvID(s *string) *SecurityAdvisoryUpdateOne {
	if s != nil {
		sauo.SetOsvID(*s)
	}
	return sauo
}

// ClearOsvID clears the value of the "osv_id" field.
func (sauo *SecurityAdvisoryUpdateOne) ClearOsvID() *SecurityAdvisoryUpdateOne {
	sauo.mutation.ClearOsvID()
	return sauo
}

// Mutation returns the SecurityAdvisoryMutation object of the builder.
func (sauo *SecurityAdvisoryUpdateOne) Mutation() *SecurityAdvisoryMutation {
	return sauo.mutation
}

// Where appends a list predicates to the SecurityAdvisoryUpdate builder.
func (sauo *SecurityAdvisoryUpdateOne) Where(ps ...predicate.SecurityAdvisory) *SecurityAdvisoryUpdateOne {
	sauo.mutation.Where(ps...)
	return sauo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (sauo *SecurityAdvisoryUpdateOne) Select(field string, fields ...string) *SecurityAdvisoryUpdateOne {
	sauo.fields = append([]string{field}, fields...)
	return sauo
}

// Save executes the query and returns the updated SecurityAdvisory entity.
func (sauo *SecurityAdvisoryUpdateOne) Save(ctx context.Context) (*SecurityAdvisory, error) {
	return withHooks(ctx, sauo.gremlinSave, sauo.mutation, sauo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (sauo *SecurityAdvisoryUpdateOne) SaveX(ctx context.Context) *SecurityAdvisory {
	node, err := sauo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (sauo *SecurityAdvisoryUpdateOne) Exec(ctx context.Context) error {
	_, err := sauo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sauo *SecurityAdvisoryUpdateOne) ExecX(ctx context.Context) {
	if err := sauo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (sauo *SecurityAdvisoryUpdateOne) gremlinSave(ctx context.Context) (*SecurityAdvisory, error) {
	res := &gremlin.Response{}
	id, ok := sauo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "SecurityAdvisory.id" for update`)}
	}
	query, bindings := sauo.gremlin(id).Query()
	if err := sauo.driver.Exec(ctx, query, bindings, res); err != nil {
		return nil, err
	}
	if err, ok := isConstantError(res); ok {
		return nil, err
	}
	sauo.mutation.done = true
	sa := &SecurityAdvisory{config: sauo.config}
	if err := sa.FromResponse(res); err != nil {
		return nil, err
	}
	return sa, nil
}

func (sauo *SecurityAdvisoryUpdateOne) gremlin(id int) *dsl.Traversal {
	v := g.V(id)
	var (
		trs []*dsl.Traversal
	)
	if value, ok := sauo.mutation.GhsaID(); ok {
		v.Property(dsl.Single, securityadvisory.FieldGhsaID, value)
	}
	if value, ok := sauo.mutation.CveID(); ok {
		v.Property(dsl.Single, securityadvisory.FieldCveID, value)
	}
	if value, ok := sauo.mutation.CveYear(); ok {
		v.Property(dsl.Single, securityadvisory.FieldCveYear, value)
	}
	if value, ok := sauo.mutation.AddedCveYear(); ok {
		v.Property(dsl.Single, securityadvisory.FieldCveYear, __.Union(__.Values(securityadvisory.FieldCveYear), __.Constant(value)).Sum())
	}
	if value, ok := sauo.mutation.OsvID(); ok {
		v.Property(dsl.Single, securityadvisory.FieldOsvID, value)
	}
	var properties []any
	if sauo.mutation.GhsaIDCleared() {
		properties = append(properties, securityadvisory.FieldGhsaID)
	}
	if sauo.mutation.CveIDCleared() {
		properties = append(properties, securityadvisory.FieldCveID)
	}
	if sauo.mutation.CveYearCleared() {
		properties = append(properties, securityadvisory.FieldCveYear)
	}
	if sauo.mutation.OsvIDCleared() {
		properties = append(properties, securityadvisory.FieldOsvID)
	}
	if len(properties) > 0 {
		v.SideEffect(__.Properties(properties...).Drop())
	}
	if len(sauo.fields) > 0 {
		fields := make([]any, 0, len(sauo.fields)+1)
		fields = append(fields, true)
		for _, f := range sauo.fields {
			fields = append(fields, f)
		}
		v.ValueMap(fields...)
	} else {
		v.ValueMap(true)
	}
	trs = append(trs, v)
	return dsl.Join(trs...)
}
