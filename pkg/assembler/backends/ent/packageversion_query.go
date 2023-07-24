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
	"github.com/guacsec/guac/pkg/assembler/backends/ent/billofmaterials"
	"github.com/guacsec/guac/pkg/assembler/backends/ent/occurrence"
	"github.com/guacsec/guac/pkg/assembler/backends/ent/packagename"
	"github.com/guacsec/guac/pkg/assembler/backends/ent/packageversion"
	"github.com/guacsec/guac/pkg/assembler/backends/ent/pkgequal"
	"github.com/guacsec/guac/pkg/assembler/backends/ent/predicate"
)

// PackageVersionQuery is the builder for querying PackageVersion entities.
type PackageVersionQuery struct {
	config
	ctx               *QueryContext
	order             []packageversion.OrderOption
	inters            []Interceptor
	predicates        []predicate.PackageVersion
	withName          *PackageNameQuery
	withOccurrences   *OccurrenceQuery
	withSbom          *BillOfMaterialsQuery
	withEqualPackages *PkgEqualQuery
	// intermediate query (i.e. traversal path).
	gremlin *dsl.Traversal
	path    func(context.Context) (*dsl.Traversal, error)
}

// Where adds a new predicate for the PackageVersionQuery builder.
func (pvq *PackageVersionQuery) Where(ps ...predicate.PackageVersion) *PackageVersionQuery {
	pvq.predicates = append(pvq.predicates, ps...)
	return pvq
}

// Limit the number of records to be returned by this query.
func (pvq *PackageVersionQuery) Limit(limit int) *PackageVersionQuery {
	pvq.ctx.Limit = &limit
	return pvq
}

// Offset to start from.
func (pvq *PackageVersionQuery) Offset(offset int) *PackageVersionQuery {
	pvq.ctx.Offset = &offset
	return pvq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (pvq *PackageVersionQuery) Unique(unique bool) *PackageVersionQuery {
	pvq.ctx.Unique = &unique
	return pvq
}

// Order specifies how the records should be ordered.
func (pvq *PackageVersionQuery) Order(o ...packageversion.OrderOption) *PackageVersionQuery {
	pvq.order = append(pvq.order, o...)
	return pvq
}

// QueryName chains the current query on the "name" edge.
func (pvq *PackageVersionQuery) QueryName() *PackageNameQuery {
	query := (&PackageNameClient{config: pvq.config}).Query()
	query.path = func(ctx context.Context) (fromU *dsl.Traversal, err error) {
		if err := pvq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		gremlin := pvq.gremlinQuery(ctx)
		fromU = gremlin.InE(packagename.VersionsLabel).OutV()
		return fromU, nil
	}
	return query
}

// QueryOccurrences chains the current query on the "occurrences" edge.
func (pvq *PackageVersionQuery) QueryOccurrences() *OccurrenceQuery {
	query := (&OccurrenceClient{config: pvq.config}).Query()
	query.path = func(ctx context.Context) (fromU *dsl.Traversal, err error) {
		if err := pvq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		gremlin := pvq.gremlinQuery(ctx)
		fromU = gremlin.InE(occurrence.PackageLabel).OutV()
		return fromU, nil
	}
	return query
}

// QuerySbom chains the current query on the "sbom" edge.
func (pvq *PackageVersionQuery) QuerySbom() *BillOfMaterialsQuery {
	query := (&BillOfMaterialsClient{config: pvq.config}).Query()
	query.path = func(ctx context.Context) (fromU *dsl.Traversal, err error) {
		if err := pvq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		gremlin := pvq.gremlinQuery(ctx)
		fromU = gremlin.InE(billofmaterials.PackageLabel).OutV()
		return fromU, nil
	}
	return query
}

// QueryEqualPackages chains the current query on the "equal_packages" edge.
func (pvq *PackageVersionQuery) QueryEqualPackages() *PkgEqualQuery {
	query := (&PkgEqualClient{config: pvq.config}).Query()
	query.path = func(ctx context.Context) (fromU *dsl.Traversal, err error) {
		if err := pvq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		gremlin := pvq.gremlinQuery(ctx)
		fromU = gremlin.InE(pkgequal.PackagesLabel).OutV()
		return fromU, nil
	}
	return query
}

// First returns the first PackageVersion entity from the query.
// Returns a *NotFoundError when no PackageVersion was found.
func (pvq *PackageVersionQuery) First(ctx context.Context) (*PackageVersion, error) {
	nodes, err := pvq.Limit(1).All(setContextOp(ctx, pvq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{packageversion.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (pvq *PackageVersionQuery) FirstX(ctx context.Context) *PackageVersion {
	node, err := pvq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first PackageVersion ID from the query.
// Returns a *NotFoundError when no PackageVersion ID was found.
func (pvq *PackageVersionQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = pvq.Limit(1).IDs(setContextOp(ctx, pvq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{packageversion.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (pvq *PackageVersionQuery) FirstIDX(ctx context.Context) int {
	id, err := pvq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single PackageVersion entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one PackageVersion entity is found.
// Returns a *NotFoundError when no PackageVersion entities are found.
func (pvq *PackageVersionQuery) Only(ctx context.Context) (*PackageVersion, error) {
	nodes, err := pvq.Limit(2).All(setContextOp(ctx, pvq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{packageversion.Label}
	default:
		return nil, &NotSingularError{packageversion.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (pvq *PackageVersionQuery) OnlyX(ctx context.Context) *PackageVersion {
	node, err := pvq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only PackageVersion ID in the query.
// Returns a *NotSingularError when more than one PackageVersion ID is found.
// Returns a *NotFoundError when no entities are found.
func (pvq *PackageVersionQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = pvq.Limit(2).IDs(setContextOp(ctx, pvq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{packageversion.Label}
	default:
		err = &NotSingularError{packageversion.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (pvq *PackageVersionQuery) OnlyIDX(ctx context.Context) int {
	id, err := pvq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of PackageVersions.
func (pvq *PackageVersionQuery) All(ctx context.Context) ([]*PackageVersion, error) {
	ctx = setContextOp(ctx, pvq.ctx, "All")
	if err := pvq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*PackageVersion, *PackageVersionQuery]()
	return withInterceptors[[]*PackageVersion](ctx, pvq, qr, pvq.inters)
}

// AllX is like All, but panics if an error occurs.
func (pvq *PackageVersionQuery) AllX(ctx context.Context) []*PackageVersion {
	nodes, err := pvq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of PackageVersion IDs.
func (pvq *PackageVersionQuery) IDs(ctx context.Context) (ids []int, err error) {
	if pvq.ctx.Unique == nil && pvq.path != nil {
		pvq.Unique(true)
	}
	ctx = setContextOp(ctx, pvq.ctx, "IDs")
	if err = pvq.Select(packageversion.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (pvq *PackageVersionQuery) IDsX(ctx context.Context) []int {
	ids, err := pvq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (pvq *PackageVersionQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, pvq.ctx, "Count")
	if err := pvq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, pvq, querierCount[*PackageVersionQuery](), pvq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (pvq *PackageVersionQuery) CountX(ctx context.Context) int {
	count, err := pvq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (pvq *PackageVersionQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, pvq.ctx, "Exist")
	switch _, err := pvq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (pvq *PackageVersionQuery) ExistX(ctx context.Context) bool {
	exist, err := pvq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the PackageVersionQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (pvq *PackageVersionQuery) Clone() *PackageVersionQuery {
	if pvq == nil {
		return nil
	}
	return &PackageVersionQuery{
		config:            pvq.config,
		ctx:               pvq.ctx.Clone(),
		order:             append([]packageversion.OrderOption{}, pvq.order...),
		inters:            append([]Interceptor{}, pvq.inters...),
		predicates:        append([]predicate.PackageVersion{}, pvq.predicates...),
		withName:          pvq.withName.Clone(),
		withOccurrences:   pvq.withOccurrences.Clone(),
		withSbom:          pvq.withSbom.Clone(),
		withEqualPackages: pvq.withEqualPackages.Clone(),
		// clone intermediate query.
		gremlin: pvq.gremlin.Clone(),
		path:    pvq.path,
	}
}

// WithName tells the query-builder to eager-load the nodes that are connected to
// the "name" edge. The optional arguments are used to configure the query builder of the edge.
func (pvq *PackageVersionQuery) WithName(opts ...func(*PackageNameQuery)) *PackageVersionQuery {
	query := (&PackageNameClient{config: pvq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	pvq.withName = query
	return pvq
}

// WithOccurrences tells the query-builder to eager-load the nodes that are connected to
// the "occurrences" edge. The optional arguments are used to configure the query builder of the edge.
func (pvq *PackageVersionQuery) WithOccurrences(opts ...func(*OccurrenceQuery)) *PackageVersionQuery {
	query := (&OccurrenceClient{config: pvq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	pvq.withOccurrences = query
	return pvq
}

// WithSbom tells the query-builder to eager-load the nodes that are connected to
// the "sbom" edge. The optional arguments are used to configure the query builder of the edge.
func (pvq *PackageVersionQuery) WithSbom(opts ...func(*BillOfMaterialsQuery)) *PackageVersionQuery {
	query := (&BillOfMaterialsClient{config: pvq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	pvq.withSbom = query
	return pvq
}

// WithEqualPackages tells the query-builder to eager-load the nodes that are connected to
// the "equal_packages" edge. The optional arguments are used to configure the query builder of the edge.
func (pvq *PackageVersionQuery) WithEqualPackages(opts ...func(*PkgEqualQuery)) *PackageVersionQuery {
	query := (&PkgEqualClient{config: pvq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	pvq.withEqualPackages = query
	return pvq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		NameID int `json:"name_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.PackageVersion.Query().
//		GroupBy(packageversion.FieldNameID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (pvq *PackageVersionQuery) GroupBy(field string, fields ...string) *PackageVersionGroupBy {
	pvq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &PackageVersionGroupBy{build: pvq}
	grbuild.flds = &pvq.ctx.Fields
	grbuild.label = packageversion.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		NameID int `json:"name_id,omitempty"`
//	}
//
//	client.PackageVersion.Query().
//		Select(packageversion.FieldNameID).
//		Scan(ctx, &v)
func (pvq *PackageVersionQuery) Select(fields ...string) *PackageVersionSelect {
	pvq.ctx.Fields = append(pvq.ctx.Fields, fields...)
	sbuild := &PackageVersionSelect{PackageVersionQuery: pvq}
	sbuild.label = packageversion.Label
	sbuild.flds, sbuild.scan = &pvq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a PackageVersionSelect configured with the given aggregations.
func (pvq *PackageVersionQuery) Aggregate(fns ...AggregateFunc) *PackageVersionSelect {
	return pvq.Select().Aggregate(fns...)
}

func (pvq *PackageVersionQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range pvq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, pvq); err != nil {
				return err
			}
		}
	}
	if pvq.path != nil {
		prev, err := pvq.path(ctx)
		if err != nil {
			return err
		}
		pvq.gremlin = prev
	}
	return nil
}

func (pvq *PackageVersionQuery) gremlinAll(ctx context.Context, hooks ...queryHook) ([]*PackageVersion, error) {
	res := &gremlin.Response{}
	traversal := pvq.gremlinQuery(ctx)
	if len(pvq.ctx.Fields) > 0 {
		fields := make([]any, len(pvq.ctx.Fields))
		for i, f := range pvq.ctx.Fields {
			fields[i] = f
		}
		traversal.ValueMap(fields...)
	} else {
		traversal.ValueMap(true)
	}
	query, bindings := traversal.Query()
	if err := pvq.driver.Exec(ctx, query, bindings, res); err != nil {
		return nil, err
	}
	var pvs PackageVersions
	if err := pvs.FromResponse(res); err != nil {
		return nil, err
	}
	for i := range pvs {
		pvs[i].config = pvq.config
	}
	return pvs, nil
}

func (pvq *PackageVersionQuery) gremlinCount(ctx context.Context) (int, error) {
	res := &gremlin.Response{}
	query, bindings := pvq.gremlinQuery(ctx).Count().Query()
	if err := pvq.driver.Exec(ctx, query, bindings, res); err != nil {
		return 0, err
	}
	return res.ReadInt()
}

func (pvq *PackageVersionQuery) gremlinQuery(context.Context) *dsl.Traversal {
	v := g.V().HasLabel(packageversion.Label)
	if pvq.gremlin != nil {
		v = pvq.gremlin.Clone()
	}
	for _, p := range pvq.predicates {
		p(v)
	}
	if len(pvq.order) > 0 {
		v.Order()
		for _, p := range pvq.order {
			p(v)
		}
	}
	switch limit, offset := pvq.ctx.Limit, pvq.ctx.Offset; {
	case limit != nil && offset != nil:
		v.Range(*offset, *offset+*limit)
	case offset != nil:
		v.Range(*offset, math.MaxInt32)
	case limit != nil:
		v.Limit(*limit)
	}
	if unique := pvq.ctx.Unique; unique == nil || *unique {
		v.Dedup()
	}
	return v
}

// PackageVersionGroupBy is the group-by builder for PackageVersion entities.
type PackageVersionGroupBy struct {
	selector
	build *PackageVersionQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (pvgb *PackageVersionGroupBy) Aggregate(fns ...AggregateFunc) *PackageVersionGroupBy {
	pvgb.fns = append(pvgb.fns, fns...)
	return pvgb
}

// Scan applies the selector query and scans the result into the given value.
func (pvgb *PackageVersionGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, pvgb.build.ctx, "GroupBy")
	if err := pvgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*PackageVersionQuery, *PackageVersionGroupBy](ctx, pvgb.build, pvgb, pvgb.build.inters, v)
}

func (pvgb *PackageVersionGroupBy) gremlinScan(ctx context.Context, root *PackageVersionQuery, v any) error {
	var (
		trs   []any
		names []any
	)
	for _, fn := range pvgb.fns {
		name, tr := fn("p", "")
		trs = append(trs, tr)
		names = append(names, name)
	}
	for _, f := range *pvgb.flds {
		names = append(names, f)
		trs = append(trs, __.As("p").Unfold().Values(f).As(f))
	}
	query, bindings := root.gremlinQuery(ctx).Group().
		By(__.Values(*pvgb.flds...).Fold()).
		By(__.Fold().Match(trs...).Select(names...)).
		Select(dsl.Values).
		Next().
		Query()
	res := &gremlin.Response{}
	if err := pvgb.build.driver.Exec(ctx, query, bindings, res); err != nil {
		return err
	}
	if len(*pvgb.flds)+len(pvgb.fns) == 1 {
		return res.ReadVal(v)
	}
	vm, err := res.ReadValueMap()
	if err != nil {
		return err
	}
	return vm.Decode(v)
}

// PackageVersionSelect is the builder for selecting fields of PackageVersion entities.
type PackageVersionSelect struct {
	*PackageVersionQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (pvs *PackageVersionSelect) Aggregate(fns ...AggregateFunc) *PackageVersionSelect {
	pvs.fns = append(pvs.fns, fns...)
	return pvs
}

// Scan applies the selector query and scans the result into the given value.
func (pvs *PackageVersionSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, pvs.ctx, "Select")
	if err := pvs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*PackageVersionQuery, *PackageVersionSelect](ctx, pvs.PackageVersionQuery, pvs, pvs.inters, v)
}

func (pvs *PackageVersionSelect) gremlinScan(ctx context.Context, root *PackageVersionQuery, v any) error {
	var (
		res       = &gremlin.Response{}
		traversal = root.gremlinQuery(ctx)
	)
	if fields := pvs.ctx.Fields; len(fields) == 1 {
		if fields[0] != packageversion.FieldID {
			traversal = traversal.Values(fields...)
		} else {
			traversal = traversal.ID()
		}
	} else {
		fields := make([]any, len(pvs.ctx.Fields))
		for i, f := range pvs.ctx.Fields {
			fields[i] = f
		}
		traversal = traversal.ValueMap(fields...)
	}
	query, bindings := traversal.Query()
	if err := pvs.driver.Exec(ctx, query, bindings, res); err != nil {
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
