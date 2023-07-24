// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"time"

	"entgo.io/ent/dialect/gremlin"
	"entgo.io/ent/dialect/gremlin/graph/dsl"
	"entgo.io/ent/dialect/gremlin/graph/dsl/__"
	"entgo.io/ent/dialect/gremlin/graph/dsl/g"
	"entgo.io/ent/dialect/gremlin/graph/dsl/p"
	"github.com/guacsec/guac/pkg/assembler/backends/ent/scorecard"
	"github.com/guacsec/guac/pkg/assembler/graphql/model"
)

// ScorecardCreate is the builder for creating a Scorecard entity.
type ScorecardCreate struct {
	config
	mutation *ScorecardMutation
	hooks    []Hook
}

// SetChecks sets the "checks" field.
func (sc *ScorecardCreate) SetChecks(mc []*model.ScorecardCheck) *ScorecardCreate {
	sc.mutation.SetChecks(mc)
	return sc
}

// SetAggregateScore sets the "aggregate_score" field.
func (sc *ScorecardCreate) SetAggregateScore(f float64) *ScorecardCreate {
	sc.mutation.SetAggregateScore(f)
	return sc
}

// SetNillableAggregateScore sets the "aggregate_score" field if the given value is not nil.
func (sc *ScorecardCreate) SetNillableAggregateScore(f *float64) *ScorecardCreate {
	if f != nil {
		sc.SetAggregateScore(*f)
	}
	return sc
}

// SetTimeScanned sets the "time_scanned" field.
func (sc *ScorecardCreate) SetTimeScanned(t time.Time) *ScorecardCreate {
	sc.mutation.SetTimeScanned(t)
	return sc
}

// SetNillableTimeScanned sets the "time_scanned" field if the given value is not nil.
func (sc *ScorecardCreate) SetNillableTimeScanned(t *time.Time) *ScorecardCreate {
	if t != nil {
		sc.SetTimeScanned(*t)
	}
	return sc
}

// SetScorecardVersion sets the "scorecard_version" field.
func (sc *ScorecardCreate) SetScorecardVersion(s string) *ScorecardCreate {
	sc.mutation.SetScorecardVersion(s)
	return sc
}

// SetScorecardCommit sets the "scorecard_commit" field.
func (sc *ScorecardCreate) SetScorecardCommit(s string) *ScorecardCreate {
	sc.mutation.SetScorecardCommit(s)
	return sc
}

// SetOrigin sets the "origin" field.
func (sc *ScorecardCreate) SetOrigin(s string) *ScorecardCreate {
	sc.mutation.SetOrigin(s)
	return sc
}

// SetCollector sets the "collector" field.
func (sc *ScorecardCreate) SetCollector(s string) *ScorecardCreate {
	sc.mutation.SetCollector(s)
	return sc
}

// AddCertificationIDs adds the "certifications" edge to the CertifyScorecard entity by IDs.
func (sc *ScorecardCreate) AddCertificationIDs(ids ...int) *ScorecardCreate {
	sc.mutation.AddCertificationIDs(ids...)
	return sc
}

// AddCertifications adds the "certifications" edges to the CertifyScorecard entity.
func (sc *ScorecardCreate) AddCertifications(c ...*CertifyScorecard) *ScorecardCreate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return sc.AddCertificationIDs(ids...)
}

// Mutation returns the ScorecardMutation object of the builder.
func (sc *ScorecardCreate) Mutation() *ScorecardMutation {
	return sc.mutation
}

// Save creates the Scorecard in the database.
func (sc *ScorecardCreate) Save(ctx context.Context) (*Scorecard, error) {
	sc.defaults()
	return withHooks(ctx, sc.gremlinSave, sc.mutation, sc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (sc *ScorecardCreate) SaveX(ctx context.Context) *Scorecard {
	v, err := sc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sc *ScorecardCreate) Exec(ctx context.Context) error {
	_, err := sc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sc *ScorecardCreate) ExecX(ctx context.Context) {
	if err := sc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (sc *ScorecardCreate) defaults() {
	if _, ok := sc.mutation.AggregateScore(); !ok {
		v := scorecard.DefaultAggregateScore
		sc.mutation.SetAggregateScore(v)
	}
	if _, ok := sc.mutation.TimeScanned(); !ok {
		v := scorecard.DefaultTimeScanned()
		sc.mutation.SetTimeScanned(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (sc *ScorecardCreate) check() error {
	if _, ok := sc.mutation.Checks(); !ok {
		return &ValidationError{Name: "checks", err: errors.New(`ent: missing required field "Scorecard.checks"`)}
	}
	if _, ok := sc.mutation.AggregateScore(); !ok {
		return &ValidationError{Name: "aggregate_score", err: errors.New(`ent: missing required field "Scorecard.aggregate_score"`)}
	}
	if _, ok := sc.mutation.TimeScanned(); !ok {
		return &ValidationError{Name: "time_scanned", err: errors.New(`ent: missing required field "Scorecard.time_scanned"`)}
	}
	if _, ok := sc.mutation.ScorecardVersion(); !ok {
		return &ValidationError{Name: "scorecard_version", err: errors.New(`ent: missing required field "Scorecard.scorecard_version"`)}
	}
	if _, ok := sc.mutation.ScorecardCommit(); !ok {
		return &ValidationError{Name: "scorecard_commit", err: errors.New(`ent: missing required field "Scorecard.scorecard_commit"`)}
	}
	if _, ok := sc.mutation.Origin(); !ok {
		return &ValidationError{Name: "origin", err: errors.New(`ent: missing required field "Scorecard.origin"`)}
	}
	if _, ok := sc.mutation.Collector(); !ok {
		return &ValidationError{Name: "collector", err: errors.New(`ent: missing required field "Scorecard.collector"`)}
	}
	return nil
}

func (sc *ScorecardCreate) gremlinSave(ctx context.Context) (*Scorecard, error) {
	if err := sc.check(); err != nil {
		return nil, err
	}
	res := &gremlin.Response{}
	query, bindings := sc.gremlin().Query()
	if err := sc.driver.Exec(ctx, query, bindings, res); err != nil {
		return nil, err
	}
	if err, ok := isConstantError(res); ok {
		return nil, err
	}
	rnode := &Scorecard{config: sc.config}
	if err := rnode.FromResponse(res); err != nil {
		return nil, err
	}
	sc.mutation.id = &rnode.ID
	sc.mutation.done = true
	return rnode, nil
}

func (sc *ScorecardCreate) gremlin() *dsl.Traversal {
	type constraint struct {
		pred *dsl.Traversal // constraint predicate.
		test *dsl.Traversal // test matches and its constant.
	}
	constraints := make([]*constraint, 0, 1)
	v := g.AddV(scorecard.Label)
	if value, ok := sc.mutation.Checks(); ok {
		v.Property(dsl.Single, scorecard.FieldChecks, value)
	}
	if value, ok := sc.mutation.AggregateScore(); ok {
		v.Property(dsl.Single, scorecard.FieldAggregateScore, value)
	}
	if value, ok := sc.mutation.TimeScanned(); ok {
		v.Property(dsl.Single, scorecard.FieldTimeScanned, value)
	}
	if value, ok := sc.mutation.ScorecardVersion(); ok {
		v.Property(dsl.Single, scorecard.FieldScorecardVersion, value)
	}
	if value, ok := sc.mutation.ScorecardCommit(); ok {
		v.Property(dsl.Single, scorecard.FieldScorecardCommit, value)
	}
	if value, ok := sc.mutation.Origin(); ok {
		v.Property(dsl.Single, scorecard.FieldOrigin, value)
	}
	if value, ok := sc.mutation.Collector(); ok {
		v.Property(dsl.Single, scorecard.FieldCollector, value)
	}
	for _, id := range sc.mutation.CertificationsIDs() {
		v.AddE(scorecard.CertificationsLabel).To(g.V(id)).OutV()
		constraints = append(constraints, &constraint{
			pred: g.E().HasLabel(scorecard.CertificationsLabel).InV().HasID(id).Count(),
			test: __.Is(p.NEQ(0)).Constant(NewErrUniqueEdge(scorecard.Label, scorecard.CertificationsLabel, id)),
		})
	}
	if len(constraints) == 0 {
		return v.ValueMap(true)
	}
	tr := constraints[0].pred.Coalesce(constraints[0].test, v.ValueMap(true))
	for _, cr := range constraints[1:] {
		tr = cr.pred.Coalesce(cr.test, tr)
	}
	return tr
}

// ScorecardCreateBulk is the builder for creating many Scorecard entities in bulk.
type ScorecardCreateBulk struct {
	config
	builders []*ScorecardCreate
}
