// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/gremlin"
	"entgo.io/ent/dialect/gremlin/graph/dsl"
	"entgo.io/ent/dialect/gremlin/graph/dsl/__"
	"entgo.io/ent/dialect/gremlin/graph/dsl/g"
	"github.com/guacsec/guac/pkg/assembler/backends/ent/predicate"
	"github.com/guacsec/guac/pkg/assembler/backends/ent/slsaattestation"
)

// SLSAAttestationQuery is the builder for querying SLSAAttestation entities.
type SLSAAttestationQuery struct {
	config
	ctx           *QueryContext
	order         []slsaattestation.OrderOption
	inters        []Interceptor
	predicates    []predicate.SLSAAttestation
	withBuiltFrom *ArtifactQuery
	withBuiltBy   *BuilderQuery
	withSubject   *ArtifactQuery
	// intermediate query (i.e. traversal path).
	gremlin *dsl.Traversal
	path    func(context.Context) (*dsl.Traversal, error)
}

// Where adds a new predicate for the SLSAAttestationQuery builder.
func (saq *SLSAAttestationQuery) Where(ps ...predicate.SLSAAttestation) *SLSAAttestationQuery {
	saq.predicates = append(saq.predicates, ps...)
	return saq
}

// Limit the number of records to be returned by this query.
func (saq *SLSAAttestationQuery) Limit(limit int) *SLSAAttestationQuery {
	saq.ctx.Limit = &limit
	return saq
}

// Offset to start from.
func (saq *SLSAAttestationQuery) Offset(offset int) *SLSAAttestationQuery {
	saq.ctx.Offset = &offset
	return saq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (saq *SLSAAttestationQuery) Unique(unique bool) *SLSAAttestationQuery {
	saq.ctx.Unique = &unique
	return saq
}

// Order specifies how the records should be ordered.
func (saq *SLSAAttestationQuery) Order(o ...slsaattestation.OrderOption) *SLSAAttestationQuery {
	saq.order = append(saq.order, o...)
	return saq
}

// QueryBuiltFrom chains the current query on the "built_from" edge.
func (saq *SLSAAttestationQuery) QueryBuiltFrom() *ArtifactQuery {
	query := (&ArtifactClient{config: saq.config}).Query()
	query.path = func(ctx context.Context) (fromU *dsl.Traversal, err error) {
		if err := saq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		gremlin := saq.gremlinQuery(ctx)
		fromU = gremlin.OutE(slsaattestation.BuiltFromLabel).InV()
		return fromU, nil
	}
	return query
}

// QueryBuiltBy chains the current query on the "built_by" edge.
func (saq *SLSAAttestationQuery) QueryBuiltBy() *BuilderQuery {
	query := (&BuilderClient{config: saq.config}).Query()
	query.path = func(ctx context.Context) (fromU *dsl.Traversal, err error) {
		if err := saq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		gremlin := saq.gremlinQuery(ctx)
		fromU = gremlin.OutE(slsaattestation.BuiltByLabel).InV()
		return fromU, nil
	}
	return query
}

// QuerySubject chains the current query on the "subject" edge.
func (saq *SLSAAttestationQuery) QuerySubject() *ArtifactQuery {
	query := (&ArtifactClient{config: saq.config}).Query()
	query.path = func(ctx context.Context) (fromU *dsl.Traversal, err error) {
		if err := saq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		gremlin := saq.gremlinQuery(ctx)
		fromU = gremlin.OutE(slsaattestation.SubjectLabel).InV()
		return fromU, nil
	}
	return query
}

// First returns the first SLSAAttestation entity from the query.
// Returns a *NotFoundError when no SLSAAttestation was found.
func (saq *SLSAAttestationQuery) First(ctx context.Context) (*SLSAAttestation, error) {
	nodes, err := saq.Limit(1).All(setContextOp(ctx, saq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{slsaattestation.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (saq *SLSAAttestationQuery) FirstX(ctx context.Context) *SLSAAttestation {
	node, err := saq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first SLSAAttestation ID from the query.
// Returns a *NotFoundError when no SLSAAttestation ID was found.
func (saq *SLSAAttestationQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = saq.Limit(1).IDs(setContextOp(ctx, saq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{slsaattestation.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (saq *SLSAAttestationQuery) FirstIDX(ctx context.Context) int {
	id, err := saq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single SLSAAttestation entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one SLSAAttestation entity is found.
// Returns a *NotFoundError when no SLSAAttestation entities are found.
func (saq *SLSAAttestationQuery) Only(ctx context.Context) (*SLSAAttestation, error) {
	nodes, err := saq.Limit(2).All(setContextOp(ctx, saq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{slsaattestation.Label}
	default:
		return nil, &NotSingularError{slsaattestation.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (saq *SLSAAttestationQuery) OnlyX(ctx context.Context) *SLSAAttestation {
	node, err := saq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only SLSAAttestation ID in the query.
// Returns a *NotSingularError when more than one SLSAAttestation ID is found.
// Returns a *NotFoundError when no entities are found.
func (saq *SLSAAttestationQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = saq.Limit(2).IDs(setContextOp(ctx, saq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{slsaattestation.Label}
	default:
		err = &NotSingularError{slsaattestation.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (saq *SLSAAttestationQuery) OnlyIDX(ctx context.Context) int {
	id, err := saq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of SLSAAttestations.
func (saq *SLSAAttestationQuery) All(ctx context.Context) ([]*SLSAAttestation, error) {
	ctx = setContextOp(ctx, saq.ctx, "All")
	if err := saq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*SLSAAttestation, *SLSAAttestationQuery]()
	return withInterceptors[[]*SLSAAttestation](ctx, saq, qr, saq.inters)
}

// AllX is like All, but panics if an error occurs.
func (saq *SLSAAttestationQuery) AllX(ctx context.Context) []*SLSAAttestation {
	nodes, err := saq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of SLSAAttestation IDs.
func (saq *SLSAAttestationQuery) IDs(ctx context.Context) (ids []int, err error) {
	if saq.ctx.Unique == nil && saq.path != nil {
		saq.Unique(true)
	}
	ctx = setContextOp(ctx, saq.ctx, "IDs")
	if err = saq.Select(slsaattestation.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (saq *SLSAAttestationQuery) IDsX(ctx context.Context) []int {
	ids, err := saq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (saq *SLSAAttestationQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, saq.ctx, "Count")
	if err := saq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, saq, querierCount[*SLSAAttestationQuery](), saq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (saq *SLSAAttestationQuery) CountX(ctx context.Context) int {
	count, err := saq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (saq *SLSAAttestationQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, saq.ctx, "Exist")
	switch _, err := saq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (saq *SLSAAttestationQuery) ExistX(ctx context.Context) bool {
	exist, err := saq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the SLSAAttestationQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (saq *SLSAAttestationQuery) Clone() *SLSAAttestationQuery {
	if saq == nil {
		return nil
	}
	return &SLSAAttestationQuery{
		config:        saq.config,
		ctx:           saq.ctx.Clone(),
		order:         append([]slsaattestation.OrderOption{}, saq.order...),
		inters:        append([]Interceptor{}, saq.inters...),
		predicates:    append([]predicate.SLSAAttestation{}, saq.predicates...),
		withBuiltFrom: saq.withBuiltFrom.Clone(),
		withBuiltBy:   saq.withBuiltBy.Clone(),
		withSubject:   saq.withSubject.Clone(),
		// clone intermediate query.
		gremlin: saq.gremlin.Clone(),
		path:    saq.path,
	}
}

// WithBuiltFrom tells the query-builder to eager-load the nodes that are connected to
// the "built_from" edge. The optional arguments are used to configure the query builder of the edge.
func (saq *SLSAAttestationQuery) WithBuiltFrom(opts ...func(*ArtifactQuery)) *SLSAAttestationQuery {
	query := (&ArtifactClient{config: saq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	saq.withBuiltFrom = query
	return saq
}

// WithBuiltBy tells the query-builder to eager-load the nodes that are connected to
// the "built_by" edge. The optional arguments are used to configure the query builder of the edge.
func (saq *SLSAAttestationQuery) WithBuiltBy(opts ...func(*BuilderQuery)) *SLSAAttestationQuery {
	query := (&BuilderClient{config: saq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	saq.withBuiltBy = query
	return saq
}

// WithSubject tells the query-builder to eager-load the nodes that are connected to
// the "subject" edge. The optional arguments are used to configure the query builder of the edge.
func (saq *SLSAAttestationQuery) WithSubject(opts ...func(*ArtifactQuery)) *SLSAAttestationQuery {
	query := (&ArtifactClient{config: saq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	saq.withSubject = query
	return saq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		BuildType string `json:"build_type,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.SLSAAttestation.Query().
//		GroupBy(slsaattestation.FieldBuildType).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (saq *SLSAAttestationQuery) GroupBy(field string, fields ...string) *SLSAAttestationGroupBy {
	saq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &SLSAAttestationGroupBy{build: saq}
	grbuild.flds = &saq.ctx.Fields
	grbuild.label = slsaattestation.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		BuildType string `json:"build_type,omitempty"`
//	}
//
//	client.SLSAAttestation.Query().
//		Select(slsaattestation.FieldBuildType).
//		Scan(ctx, &v)
func (saq *SLSAAttestationQuery) Select(fields ...string) *SLSAAttestationSelect {
	saq.ctx.Fields = append(saq.ctx.Fields, fields...)
	sbuild := &SLSAAttestationSelect{SLSAAttestationQuery: saq}
	sbuild.label = slsaattestation.Label
	sbuild.flds, sbuild.scan = &saq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a SLSAAttestationSelect configured with the given aggregations.
func (saq *SLSAAttestationQuery) Aggregate(fns ...AggregateFunc) *SLSAAttestationSelect {
	return saq.Select().Aggregate(fns...)
}

func (saq *SLSAAttestationQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range saq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, saq); err != nil {
				return err
			}
		}
	}
	if saq.path != nil {
		prev, err := saq.path(ctx)
		if err != nil {
			return err
		}
		saq.gremlin = prev
	}
	return nil
}

func (saq *SLSAAttestationQuery) gremlinAll(ctx context.Context, hooks ...queryHook) ([]*SLSAAttestation, error) {
	res := &gremlin.Response{}
	traversal := saq.gremlinQuery(ctx)
	if len(saq.ctx.Fields) > 0 {
		fields := make([]any, len(saq.ctx.Fields))
		for i, f := range saq.ctx.Fields {
			fields[i] = f
		}
		traversal.ValueMap(fields...)
	} else {
		traversal.ValueMap(true)
	}
	query, bindings := traversal.Query()
	if err := saq.driver.Exec(ctx, query, bindings, res); err != nil {
		return nil, err
	}
	var sas SLSAAttestations
	if err := sas.FromResponse(res); err != nil {
		return nil, err
	}
	for i := range sas {
		sas[i].config = saq.config
	}
	return sas, nil
}

func (saq *SLSAAttestationQuery) gremlinCount(ctx context.Context) (int, error) {
	res := &gremlin.Response{}
	query, bindings := saq.gremlinQuery(ctx).Count().Query()
	if err := saq.driver.Exec(ctx, query, bindings, res); err != nil {
		return 0, err
	}
	return res.ReadInt()
}

func (saq *SLSAAttestationQuery) gremlinQuery(context.Context) *dsl.Traversal {
	v := g.V().HasLabel(slsaattestation.Label)
	if saq.gremlin != nil {
		v = saq.gremlin.Clone()
	}
	for _, p := range saq.predicates {
		p(v)
	}
	if len(saq.order) > 0 {
		v.Order()
		for _, p := range saq.order {
			p(v)
		}
	}
	switch limit, offset := saq.ctx.Limit, saq.ctx.Offset; {
	case limit != nil && offset != nil:
		v.Range(*offset, *offset+*limit)
	case offset != nil:
		v.Range(*offset, math.MaxInt32)
	case limit != nil:
		v.Limit(*limit)
	}
	if unique := saq.ctx.Unique; unique == nil || *unique {
		v.Dedup()
	}
	return v
}

// SLSAAttestationGroupBy is the group-by builder for SLSAAttestation entities.
type SLSAAttestationGroupBy struct {
	selector
	build *SLSAAttestationQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (sagb *SLSAAttestationGroupBy) Aggregate(fns ...AggregateFunc) *SLSAAttestationGroupBy {
	sagb.fns = append(sagb.fns, fns...)
	return sagb
}

// Scan applies the selector query and scans the result into the given value.
func (sagb *SLSAAttestationGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, sagb.build.ctx, "GroupBy")
	if err := sagb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*SLSAAttestationQuery, *SLSAAttestationGroupBy](ctx, sagb.build, sagb, sagb.build.inters, v)
}

func (sagb *SLSAAttestationGroupBy) gremlinScan(ctx context.Context, root *SLSAAttestationQuery, v any) error {
	var (
		trs   []any
		names []any
	)
	for _, fn := range sagb.fns {
		name, tr := fn("p", "")
		trs = append(trs, tr)
		names = append(names, name)
	}
	for _, f := range *sagb.flds {
		names = append(names, f)
		trs = append(trs, __.As("p").Unfold().Values(f).As(f))
	}
	query, bindings := root.gremlinQuery(ctx).Group().
		By(__.Values(*sagb.flds...).Fold()).
		By(__.Fold().Match(trs...).Select(names...)).
		Select(dsl.Values).
		Next().
		Query()
	res := &gremlin.Response{}
	if err := sagb.build.driver.Exec(ctx, query, bindings, res); err != nil {
		return err
	}
	if len(*sagb.flds)+len(sagb.fns) == 1 {
		return res.ReadVal(v)
	}
	vm, err := res.ReadValueMap()
	if err != nil {
		return err
	}
	return vm.Decode(v)
}

// SLSAAttestationSelect is the builder for selecting fields of SLSAAttestation entities.
type SLSAAttestationSelect struct {
	*SLSAAttestationQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (sas *SLSAAttestationSelect) Aggregate(fns ...AggregateFunc) *SLSAAttestationSelect {
	sas.fns = append(sas.fns, fns...)
	return sas
}

// Scan applies the selector query and scans the result into the given value.
func (sas *SLSAAttestationSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, sas.ctx, "Select")
	if err := sas.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*SLSAAttestationQuery, *SLSAAttestationSelect](ctx, sas.SLSAAttestationQuery, sas, sas.inters, v)
}

func (sas *SLSAAttestationSelect) gremlinScan(ctx context.Context, root *SLSAAttestationQuery, v any) error {
	var (
		res       = &gremlin.Response{}
		traversal = root.gremlinQuery(ctx)
	)
	if fields := sas.ctx.Fields; len(fields) == 1 {
		if fields[0] != slsaattestation.FieldID {
			traversal = traversal.Values(fields...)
		} else {
			traversal = traversal.ID()
		}
	} else {
		fields := make([]any, len(sas.ctx.Fields))
		for i, f := range sas.ctx.Fields {
			fields[i] = f
		}
		traversal = traversal.ValueMap(fields...)
	}
	query, bindings := traversal.Query()
	if err := sas.driver.Exec(ctx, query, bindings, res); err != nil {
		return err
	}
	if len(root.ctx.Fields) == 1 {
		return res.ReadVal(v)
	}
	vm, err := res.ReadValueMap()
	if err != nil {
		return err
	}
	return vm.Decode(v)
}
