// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/guacsec/guac/pkg/assembler/backends/ent/predicate"
	"github.com/guacsec/guac/pkg/assembler/backends/ent/source"
	"github.com/guacsec/guac/pkg/assembler/backends/ent/sourcename"
	"github.com/guacsec/guac/pkg/assembler/backends/ent/sourcenamespace"
)

// SourceNamespaceQuery is the builder for querying SourceNamespace entities.
type SourceNamespaceQuery struct {
	config
	ctx        *QueryContext
	order      []sourcenamespace.OrderOption
	inters     []Interceptor
	predicates []predicate.SourceNamespace
	withSource *SourceQuery
	withNames  *SourceNameQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the SourceNamespaceQuery builder.
func (snq *SourceNamespaceQuery) Where(ps ...predicate.SourceNamespace) *SourceNamespaceQuery {
	snq.predicates = append(snq.predicates, ps...)
	return snq
}

// Limit the number of records to be returned by this query.
func (snq *SourceNamespaceQuery) Limit(limit int) *SourceNamespaceQuery {
	snq.ctx.Limit = &limit
	return snq
}

// Offset to start from.
func (snq *SourceNamespaceQuery) Offset(offset int) *SourceNamespaceQuery {
	snq.ctx.Offset = &offset
	return snq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (snq *SourceNamespaceQuery) Unique(unique bool) *SourceNamespaceQuery {
	snq.ctx.Unique = &unique
	return snq
}

// Order specifies how the records should be ordered.
func (snq *SourceNamespaceQuery) Order(o ...sourcenamespace.OrderOption) *SourceNamespaceQuery {
	snq.order = append(snq.order, o...)
	return snq
}

// QuerySource chains the current query on the "source" edge.
func (snq *SourceNamespaceQuery) QuerySource() *SourceQuery {
	query := (&SourceClient{config: snq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := snq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := snq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(sourcenamespace.Table, sourcenamespace.FieldID, selector),
			sqlgraph.To(source.Table, source.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, sourcenamespace.SourceTable, sourcenamespace.SourceColumn),
		)
		fromU = sqlgraph.SetNeighbors(snq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryNames chains the current query on the "names" edge.
func (snq *SourceNamespaceQuery) QueryNames() *SourceNameQuery {
	query := (&SourceNameClient{config: snq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := snq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := snq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(sourcenamespace.Table, sourcenamespace.FieldID, selector),
			sqlgraph.To(sourcename.Table, sourcename.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, sourcenamespace.NamesTable, sourcenamespace.NamesColumn),
		)
		fromU = sqlgraph.SetNeighbors(snq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first SourceNamespace entity from the query.
// Returns a *NotFoundError when no SourceNamespace was found.
func (snq *SourceNamespaceQuery) First(ctx context.Context) (*SourceNamespace, error) {
	nodes, err := snq.Limit(1).All(setContextOp(ctx, snq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{sourcenamespace.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (snq *SourceNamespaceQuery) FirstX(ctx context.Context) *SourceNamespace {
	node, err := snq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first SourceNamespace ID from the query.
// Returns a *NotFoundError when no SourceNamespace ID was found.
func (snq *SourceNamespaceQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = snq.Limit(1).IDs(setContextOp(ctx, snq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{sourcenamespace.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (snq *SourceNamespaceQuery) FirstIDX(ctx context.Context) int {
	id, err := snq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single SourceNamespace entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one SourceNamespace entity is found.
// Returns a *NotFoundError when no SourceNamespace entities are found.
func (snq *SourceNamespaceQuery) Only(ctx context.Context) (*SourceNamespace, error) {
	nodes, err := snq.Limit(2).All(setContextOp(ctx, snq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{sourcenamespace.Label}
	default:
		return nil, &NotSingularError{sourcenamespace.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (snq *SourceNamespaceQuery) OnlyX(ctx context.Context) *SourceNamespace {
	node, err := snq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only SourceNamespace ID in the query.
// Returns a *NotSingularError when more than one SourceNamespace ID is found.
// Returns a *NotFoundError when no entities are found.
func (snq *SourceNamespaceQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = snq.Limit(2).IDs(setContextOp(ctx, snq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{sourcenamespace.Label}
	default:
		err = &NotSingularError{sourcenamespace.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (snq *SourceNamespaceQuery) OnlyIDX(ctx context.Context) int {
	id, err := snq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of SourceNamespaces.
func (snq *SourceNamespaceQuery) All(ctx context.Context) ([]*SourceNamespace, error) {
	ctx = setContextOp(ctx, snq.ctx, "All")
	if err := snq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*SourceNamespace, *SourceNamespaceQuery]()
	return withInterceptors[[]*SourceNamespace](ctx, snq, qr, snq.inters)
}

// AllX is like All, but panics if an error occurs.
func (snq *SourceNamespaceQuery) AllX(ctx context.Context) []*SourceNamespace {
	nodes, err := snq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of SourceNamespace IDs.
func (snq *SourceNamespaceQuery) IDs(ctx context.Context) (ids []int, err error) {
	if snq.ctx.Unique == nil && snq.path != nil {
		snq.Unique(true)
	}
	ctx = setContextOp(ctx, snq.ctx, "IDs")
	if err = snq.Select(sourcenamespace.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (snq *SourceNamespaceQuery) IDsX(ctx context.Context) []int {
	ids, err := snq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (snq *SourceNamespaceQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, snq.ctx, "Count")
	if err := snq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, snq, querierCount[*SourceNamespaceQuery](), snq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (snq *SourceNamespaceQuery) CountX(ctx context.Context) int {
	count, err := snq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (snq *SourceNamespaceQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, snq.ctx, "Exist")
	switch _, err := snq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (snq *SourceNamespaceQuery) ExistX(ctx context.Context) bool {
	exist, err := snq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the SourceNamespaceQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (snq *SourceNamespaceQuery) Clone() *SourceNamespaceQuery {
	if snq == nil {
		return nil
	}
	return &SourceNamespaceQuery{
		config:     snq.config,
		ctx:        snq.ctx.Clone(),
		order:      append([]sourcenamespace.OrderOption{}, snq.order...),
		inters:     append([]Interceptor{}, snq.inters...),
		predicates: append([]predicate.SourceNamespace{}, snq.predicates...),
		withSource: snq.withSource.Clone(),
		withNames:  snq.withNames.Clone(),
		// clone intermediate query.
		sql:  snq.sql.Clone(),
		path: snq.path,
	}
}

// WithSource tells the query-builder to eager-load the nodes that are connected to
// the "source" edge. The optional arguments are used to configure the query builder of the edge.
func (snq *SourceNamespaceQuery) WithSource(opts ...func(*SourceQuery)) *SourceNamespaceQuery {
	query := (&SourceClient{config: snq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	snq.withSource = query
	return snq
}

// WithNames tells the query-builder to eager-load the nodes that are connected to
// the "names" edge. The optional arguments are used to configure the query builder of the edge.
func (snq *SourceNamespaceQuery) WithNames(opts ...func(*SourceNameQuery)) *SourceNamespaceQuery {
	query := (&SourceNameClient{config: snq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	snq.withNames = query
	return snq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Namespace string `json:"namespace,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.SourceNamespace.Query().
//		GroupBy(sourcenamespace.FieldNamespace).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (snq *SourceNamespaceQuery) GroupBy(field string, fields ...string) *SourceNamespaceGroupBy {
	snq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &SourceNamespaceGroupBy{build: snq}
	grbuild.flds = &snq.ctx.Fields
	grbuild.label = sourcenamespace.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Namespace string `json:"namespace,omitempty"`
//	}
//
//	client.SourceNamespace.Query().
//		Select(sourcenamespace.FieldNamespace).
//		Scan(ctx, &v)
func (snq *SourceNamespaceQuery) Select(fields ...string) *SourceNamespaceSelect {
	snq.ctx.Fields = append(snq.ctx.Fields, fields...)
	sbuild := &SourceNamespaceSelect{SourceNamespaceQuery: snq}
	sbuild.label = sourcenamespace.Label
	sbuild.flds, sbuild.scan = &snq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a SourceNamespaceSelect configured with the given aggregations.
func (snq *SourceNamespaceQuery) Aggregate(fns ...AggregateFunc) *SourceNamespaceSelect {
	return snq.Select().Aggregate(fns...)
}

func (snq *SourceNamespaceQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range snq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, snq); err != nil {
				return err
			}
		}
	}
	for _, f := range snq.ctx.Fields {
		if !sourcenamespace.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if snq.path != nil {
		prev, err := snq.path(ctx)
		if err != nil {
			return err
		}
		snq.sql = prev
	}
	return nil
}

func (snq *SourceNamespaceQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*SourceNamespace, error) {
	var (
		nodes       = []*SourceNamespace{}
		_spec       = snq.querySpec()
		loadedTypes = [2]bool{
			snq.withSource != nil,
			snq.withNames != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*SourceNamespace).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &SourceNamespace{config: snq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, snq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := snq.withSource; query != nil {
		if err := snq.loadSource(ctx, query, nodes, nil,
			func(n *SourceNamespace, e *Source) { n.Edges.Source = e }); err != nil {
			return nil, err
		}
	}
	if query := snq.withNames; query != nil {
		if err := snq.loadNames(ctx, query, nodes,
			func(n *SourceNamespace) { n.Edges.Names = []*SourceName{} },
			func(n *SourceNamespace, e *SourceName) { n.Edges.Names = append(n.Edges.Names, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (snq *SourceNamespaceQuery) loadSource(ctx context.Context, query *SourceQuery, nodes []*SourceNamespace, init func(*SourceNamespace), assign func(*SourceNamespace, *Source)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*SourceNamespace)
	for i := range nodes {
		fk := nodes[i].SourceID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(source.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "source_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (snq *SourceNamespaceQuery) loadNames(ctx context.Context, query *SourceNameQuery, nodes []*SourceNamespace, init func(*SourceNamespace), assign func(*SourceNamespace, *SourceName)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*SourceNamespace)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(sourcename.FieldNamespaceID)
	}
	query.Where(predicate.SourceName(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(sourcenamespace.NamesColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.NamespaceID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "namespace_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (snq *SourceNamespaceQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := snq.querySpec()
	_spec.Node.Columns = snq.ctx.Fields
	if len(snq.ctx.Fields) > 0 {
		_spec.Unique = snq.ctx.Unique != nil && *snq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, snq.driver, _spec)
}

func (snq *SourceNamespaceQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(sourcenamespace.Table, sourcenamespace.Columns, sqlgraph.NewFieldSpec(sourcenamespace.FieldID, field.TypeInt))
	_spec.From = snq.sql
	if unique := snq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if snq.path != nil {
		_spec.Unique = true
	}
	if fields := snq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, sourcenamespace.FieldID)
		for i := range fields {
			if fields[i] != sourcenamespace.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if snq.withSource != nil {
			_spec.Node.AddColumnOnce(sourcenamespace.FieldSourceID)
		}
	}
	if ps := snq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := snq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := snq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := snq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (snq *SourceNamespaceQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(snq.driver.Dialect())
	t1 := builder.Table(sourcenamespace.Table)
	columns := snq.ctx.Fields
	if len(columns) == 0 {
		columns = sourcenamespace.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if snq.sql != nil {
		selector = snq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if snq.ctx.Unique != nil && *snq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range snq.predicates {
		p(selector)
	}
	for _, p := range snq.order {
		p(selector)
	}
	if offset := snq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := snq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// SourceNamespaceGroupBy is the group-by builder for SourceNamespace entities.
type SourceNamespaceGroupBy struct {
	selector
	build *SourceNamespaceQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (sngb *SourceNamespaceGroupBy) Aggregate(fns ...AggregateFunc) *SourceNamespaceGroupBy {
	sngb.fns = append(sngb.fns, fns...)
	return sngb
}

// Scan applies the selector query and scans the result into the given value.
func (sngb *SourceNamespaceGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, sngb.build.ctx, "GroupBy")
	if err := sngb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*SourceNamespaceQuery, *SourceNamespaceGroupBy](ctx, sngb.build, sngb, sngb.build.inters, v)
}

func (sngb *SourceNamespaceGroupBy) sqlScan(ctx context.Context, root *SourceNamespaceQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(sngb.fns))
	for _, fn := range sngb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*sngb.flds)+len(sngb.fns))
		for _, f := range *sngb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*sngb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := sngb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// SourceNamespaceSelect is the builder for selecting fields of SourceNamespace entities.
type SourceNamespaceSelect struct {
	*SourceNamespaceQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (sns *SourceNamespaceSelect) Aggregate(fns ...AggregateFunc) *SourceNamespaceSelect {
	sns.fns = append(sns.fns, fns...)
	return sns
}

// Scan applies the selector query and scans the result into the given value.
func (sns *SourceNamespaceSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, sns.ctx, "Select")
	if err := sns.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*SourceNamespaceQuery, *SourceNamespaceSelect](ctx, sns.SourceNamespaceQuery, sns, sns.inters, v)
}

func (sns *SourceNamespaceSelect) sqlScan(ctx context.Context, root *SourceNamespaceQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(sns.fns))
	for _, fn := range sns.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*sns.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := sns.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
