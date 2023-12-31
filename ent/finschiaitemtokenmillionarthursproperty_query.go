// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/a10adotapp/a10a.app/ent/finschiaitemtoken"
	"github.com/a10adotapp/a10a.app/ent/finschiaitemtokenmillionarthursproperty"
	"github.com/a10adotapp/a10a.app/ent/predicate"
)

// FinschiaItemTokenMillionArthursPropertyQuery is the builder for querying FinschiaItemTokenMillionArthursProperty entities.
type FinschiaItemTokenMillionArthursPropertyQuery struct {
	config
	ctx                   *QueryContext
	order                 []finschiaitemtokenmillionarthursproperty.OrderOption
	inters                []Interceptor
	predicates            []predicate.FinschiaItemTokenMillionArthursProperty
	withFinschiaItemToken *FinschiaItemTokenQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the FinschiaItemTokenMillionArthursPropertyQuery builder.
func (fitmapq *FinschiaItemTokenMillionArthursPropertyQuery) Where(ps ...predicate.FinschiaItemTokenMillionArthursProperty) *FinschiaItemTokenMillionArthursPropertyQuery {
	fitmapq.predicates = append(fitmapq.predicates, ps...)
	return fitmapq
}

// Limit the number of records to be returned by this query.
func (fitmapq *FinschiaItemTokenMillionArthursPropertyQuery) Limit(limit int) *FinschiaItemTokenMillionArthursPropertyQuery {
	fitmapq.ctx.Limit = &limit
	return fitmapq
}

// Offset to start from.
func (fitmapq *FinschiaItemTokenMillionArthursPropertyQuery) Offset(offset int) *FinschiaItemTokenMillionArthursPropertyQuery {
	fitmapq.ctx.Offset = &offset
	return fitmapq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (fitmapq *FinschiaItemTokenMillionArthursPropertyQuery) Unique(unique bool) *FinschiaItemTokenMillionArthursPropertyQuery {
	fitmapq.ctx.Unique = &unique
	return fitmapq
}

// Order specifies how the records should be ordered.
func (fitmapq *FinschiaItemTokenMillionArthursPropertyQuery) Order(o ...finschiaitemtokenmillionarthursproperty.OrderOption) *FinschiaItemTokenMillionArthursPropertyQuery {
	fitmapq.order = append(fitmapq.order, o...)
	return fitmapq
}

// QueryFinschiaItemToken chains the current query on the "finschia_item_token" edge.
func (fitmapq *FinschiaItemTokenMillionArthursPropertyQuery) QueryFinschiaItemToken() *FinschiaItemTokenQuery {
	query := (&FinschiaItemTokenClient{config: fitmapq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := fitmapq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := fitmapq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(finschiaitemtokenmillionarthursproperty.Table, finschiaitemtokenmillionarthursproperty.FieldID, selector),
			sqlgraph.To(finschiaitemtoken.Table, finschiaitemtoken.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, finschiaitemtokenmillionarthursproperty.FinschiaItemTokenTable, finschiaitemtokenmillionarthursproperty.FinschiaItemTokenColumn),
		)
		fromU = sqlgraph.SetNeighbors(fitmapq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first FinschiaItemTokenMillionArthursProperty entity from the query.
// Returns a *NotFoundError when no FinschiaItemTokenMillionArthursProperty was found.
func (fitmapq *FinschiaItemTokenMillionArthursPropertyQuery) First(ctx context.Context) (*FinschiaItemTokenMillionArthursProperty, error) {
	nodes, err := fitmapq.Limit(1).All(setContextOp(ctx, fitmapq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{finschiaitemtokenmillionarthursproperty.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (fitmapq *FinschiaItemTokenMillionArthursPropertyQuery) FirstX(ctx context.Context) *FinschiaItemTokenMillionArthursProperty {
	node, err := fitmapq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first FinschiaItemTokenMillionArthursProperty ID from the query.
// Returns a *NotFoundError when no FinschiaItemTokenMillionArthursProperty ID was found.
func (fitmapq *FinschiaItemTokenMillionArthursPropertyQuery) FirstID(ctx context.Context) (id uint32, err error) {
	var ids []uint32
	if ids, err = fitmapq.Limit(1).IDs(setContextOp(ctx, fitmapq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{finschiaitemtokenmillionarthursproperty.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (fitmapq *FinschiaItemTokenMillionArthursPropertyQuery) FirstIDX(ctx context.Context) uint32 {
	id, err := fitmapq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single FinschiaItemTokenMillionArthursProperty entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one FinschiaItemTokenMillionArthursProperty entity is found.
// Returns a *NotFoundError when no FinschiaItemTokenMillionArthursProperty entities are found.
func (fitmapq *FinschiaItemTokenMillionArthursPropertyQuery) Only(ctx context.Context) (*FinschiaItemTokenMillionArthursProperty, error) {
	nodes, err := fitmapq.Limit(2).All(setContextOp(ctx, fitmapq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{finschiaitemtokenmillionarthursproperty.Label}
	default:
		return nil, &NotSingularError{finschiaitemtokenmillionarthursproperty.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (fitmapq *FinschiaItemTokenMillionArthursPropertyQuery) OnlyX(ctx context.Context) *FinschiaItemTokenMillionArthursProperty {
	node, err := fitmapq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only FinschiaItemTokenMillionArthursProperty ID in the query.
// Returns a *NotSingularError when more than one FinschiaItemTokenMillionArthursProperty ID is found.
// Returns a *NotFoundError when no entities are found.
func (fitmapq *FinschiaItemTokenMillionArthursPropertyQuery) OnlyID(ctx context.Context) (id uint32, err error) {
	var ids []uint32
	if ids, err = fitmapq.Limit(2).IDs(setContextOp(ctx, fitmapq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{finschiaitemtokenmillionarthursproperty.Label}
	default:
		err = &NotSingularError{finschiaitemtokenmillionarthursproperty.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (fitmapq *FinschiaItemTokenMillionArthursPropertyQuery) OnlyIDX(ctx context.Context) uint32 {
	id, err := fitmapq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of FinschiaItemTokenMillionArthursProperties.
func (fitmapq *FinschiaItemTokenMillionArthursPropertyQuery) All(ctx context.Context) ([]*FinschiaItemTokenMillionArthursProperty, error) {
	ctx = setContextOp(ctx, fitmapq.ctx, "All")
	if err := fitmapq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*FinschiaItemTokenMillionArthursProperty, *FinschiaItemTokenMillionArthursPropertyQuery]()
	return withInterceptors[[]*FinschiaItemTokenMillionArthursProperty](ctx, fitmapq, qr, fitmapq.inters)
}

// AllX is like All, but panics if an error occurs.
func (fitmapq *FinschiaItemTokenMillionArthursPropertyQuery) AllX(ctx context.Context) []*FinschiaItemTokenMillionArthursProperty {
	nodes, err := fitmapq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of FinschiaItemTokenMillionArthursProperty IDs.
func (fitmapq *FinschiaItemTokenMillionArthursPropertyQuery) IDs(ctx context.Context) (ids []uint32, err error) {
	if fitmapq.ctx.Unique == nil && fitmapq.path != nil {
		fitmapq.Unique(true)
	}
	ctx = setContextOp(ctx, fitmapq.ctx, "IDs")
	if err = fitmapq.Select(finschiaitemtokenmillionarthursproperty.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (fitmapq *FinschiaItemTokenMillionArthursPropertyQuery) IDsX(ctx context.Context) []uint32 {
	ids, err := fitmapq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (fitmapq *FinschiaItemTokenMillionArthursPropertyQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, fitmapq.ctx, "Count")
	if err := fitmapq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, fitmapq, querierCount[*FinschiaItemTokenMillionArthursPropertyQuery](), fitmapq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (fitmapq *FinschiaItemTokenMillionArthursPropertyQuery) CountX(ctx context.Context) int {
	count, err := fitmapq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (fitmapq *FinschiaItemTokenMillionArthursPropertyQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, fitmapq.ctx, "Exist")
	switch _, err := fitmapq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (fitmapq *FinschiaItemTokenMillionArthursPropertyQuery) ExistX(ctx context.Context) bool {
	exist, err := fitmapq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the FinschiaItemTokenMillionArthursPropertyQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (fitmapq *FinschiaItemTokenMillionArthursPropertyQuery) Clone() *FinschiaItemTokenMillionArthursPropertyQuery {
	if fitmapq == nil {
		return nil
	}
	return &FinschiaItemTokenMillionArthursPropertyQuery{
		config:                fitmapq.config,
		ctx:                   fitmapq.ctx.Clone(),
		order:                 append([]finschiaitemtokenmillionarthursproperty.OrderOption{}, fitmapq.order...),
		inters:                append([]Interceptor{}, fitmapq.inters...),
		predicates:            append([]predicate.FinschiaItemTokenMillionArthursProperty{}, fitmapq.predicates...),
		withFinschiaItemToken: fitmapq.withFinschiaItemToken.Clone(),
		// clone intermediate query.
		sql:  fitmapq.sql.Clone(),
		path: fitmapq.path,
	}
}

// WithFinschiaItemToken tells the query-builder to eager-load the nodes that are connected to
// the "finschia_item_token" edge. The optional arguments are used to configure the query builder of the edge.
func (fitmapq *FinschiaItemTokenMillionArthursPropertyQuery) WithFinschiaItemToken(opts ...func(*FinschiaItemTokenQuery)) *FinschiaItemTokenMillionArthursPropertyQuery {
	query := (&FinschiaItemTokenClient{config: fitmapq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	fitmapq.withFinschiaItemToken = query
	return fitmapq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.FinschiaItemTokenMillionArthursProperty.Query().
//		GroupBy(finschiaitemtokenmillionarthursproperty.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (fitmapq *FinschiaItemTokenMillionArthursPropertyQuery) GroupBy(field string, fields ...string) *FinschiaItemTokenMillionArthursPropertyGroupBy {
	fitmapq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &FinschiaItemTokenMillionArthursPropertyGroupBy{build: fitmapq}
	grbuild.flds = &fitmapq.ctx.Fields
	grbuild.label = finschiaitemtokenmillionarthursproperty.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at"`
//	}
//
//	client.FinschiaItemTokenMillionArthursProperty.Query().
//		Select(finschiaitemtokenmillionarthursproperty.FieldCreatedAt).
//		Scan(ctx, &v)
func (fitmapq *FinschiaItemTokenMillionArthursPropertyQuery) Select(fields ...string) *FinschiaItemTokenMillionArthursPropertySelect {
	fitmapq.ctx.Fields = append(fitmapq.ctx.Fields, fields...)
	sbuild := &FinschiaItemTokenMillionArthursPropertySelect{FinschiaItemTokenMillionArthursPropertyQuery: fitmapq}
	sbuild.label = finschiaitemtokenmillionarthursproperty.Label
	sbuild.flds, sbuild.scan = &fitmapq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a FinschiaItemTokenMillionArthursPropertySelect configured with the given aggregations.
func (fitmapq *FinschiaItemTokenMillionArthursPropertyQuery) Aggregate(fns ...AggregateFunc) *FinschiaItemTokenMillionArthursPropertySelect {
	return fitmapq.Select().Aggregate(fns...)
}

func (fitmapq *FinschiaItemTokenMillionArthursPropertyQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range fitmapq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, fitmapq); err != nil {
				return err
			}
		}
	}
	for _, f := range fitmapq.ctx.Fields {
		if !finschiaitemtokenmillionarthursproperty.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if fitmapq.path != nil {
		prev, err := fitmapq.path(ctx)
		if err != nil {
			return err
		}
		fitmapq.sql = prev
	}
	return nil
}

func (fitmapq *FinschiaItemTokenMillionArthursPropertyQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*FinschiaItemTokenMillionArthursProperty, error) {
	var (
		nodes       = []*FinschiaItemTokenMillionArthursProperty{}
		_spec       = fitmapq.querySpec()
		loadedTypes = [1]bool{
			fitmapq.withFinschiaItemToken != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*FinschiaItemTokenMillionArthursProperty).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &FinschiaItemTokenMillionArthursProperty{config: fitmapq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, fitmapq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := fitmapq.withFinschiaItemToken; query != nil {
		if err := fitmapq.loadFinschiaItemToken(ctx, query, nodes, nil,
			func(n *FinschiaItemTokenMillionArthursProperty, e *FinschiaItemToken) { n.Edges.FinschiaItemToken = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (fitmapq *FinschiaItemTokenMillionArthursPropertyQuery) loadFinschiaItemToken(ctx context.Context, query *FinschiaItemTokenQuery, nodes []*FinschiaItemTokenMillionArthursProperty, init func(*FinschiaItemTokenMillionArthursProperty), assign func(*FinschiaItemTokenMillionArthursProperty, *FinschiaItemToken)) error {
	ids := make([]uint32, 0, len(nodes))
	nodeids := make(map[uint32][]*FinschiaItemTokenMillionArthursProperty)
	for i := range nodes {
		fk := nodes[i].FinschiaItemTokenID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(finschiaitemtoken.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "finschia_item_token_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (fitmapq *FinschiaItemTokenMillionArthursPropertyQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := fitmapq.querySpec()
	_spec.Node.Columns = fitmapq.ctx.Fields
	if len(fitmapq.ctx.Fields) > 0 {
		_spec.Unique = fitmapq.ctx.Unique != nil && *fitmapq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, fitmapq.driver, _spec)
}

func (fitmapq *FinschiaItemTokenMillionArthursPropertyQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(finschiaitemtokenmillionarthursproperty.Table, finschiaitemtokenmillionarthursproperty.Columns, sqlgraph.NewFieldSpec(finschiaitemtokenmillionarthursproperty.FieldID, field.TypeUint32))
	_spec.From = fitmapq.sql
	if unique := fitmapq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if fitmapq.path != nil {
		_spec.Unique = true
	}
	if fields := fitmapq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, finschiaitemtokenmillionarthursproperty.FieldID)
		for i := range fields {
			if fields[i] != finschiaitemtokenmillionarthursproperty.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if fitmapq.withFinschiaItemToken != nil {
			_spec.Node.AddColumnOnce(finschiaitemtokenmillionarthursproperty.FieldFinschiaItemTokenID)
		}
	}
	if ps := fitmapq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := fitmapq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := fitmapq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := fitmapq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (fitmapq *FinschiaItemTokenMillionArthursPropertyQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(fitmapq.driver.Dialect())
	t1 := builder.Table(finschiaitemtokenmillionarthursproperty.Table)
	columns := fitmapq.ctx.Fields
	if len(columns) == 0 {
		columns = finschiaitemtokenmillionarthursproperty.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if fitmapq.sql != nil {
		selector = fitmapq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if fitmapq.ctx.Unique != nil && *fitmapq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range fitmapq.predicates {
		p(selector)
	}
	for _, p := range fitmapq.order {
		p(selector)
	}
	if offset := fitmapq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := fitmapq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// FinschiaItemTokenMillionArthursPropertyGroupBy is the group-by builder for FinschiaItemTokenMillionArthursProperty entities.
type FinschiaItemTokenMillionArthursPropertyGroupBy struct {
	selector
	build *FinschiaItemTokenMillionArthursPropertyQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (fitmapgb *FinschiaItemTokenMillionArthursPropertyGroupBy) Aggregate(fns ...AggregateFunc) *FinschiaItemTokenMillionArthursPropertyGroupBy {
	fitmapgb.fns = append(fitmapgb.fns, fns...)
	return fitmapgb
}

// Scan applies the selector query and scans the result into the given value.
func (fitmapgb *FinschiaItemTokenMillionArthursPropertyGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, fitmapgb.build.ctx, "GroupBy")
	if err := fitmapgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*FinschiaItemTokenMillionArthursPropertyQuery, *FinschiaItemTokenMillionArthursPropertyGroupBy](ctx, fitmapgb.build, fitmapgb, fitmapgb.build.inters, v)
}

func (fitmapgb *FinschiaItemTokenMillionArthursPropertyGroupBy) sqlScan(ctx context.Context, root *FinschiaItemTokenMillionArthursPropertyQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(fitmapgb.fns))
	for _, fn := range fitmapgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*fitmapgb.flds)+len(fitmapgb.fns))
		for _, f := range *fitmapgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*fitmapgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := fitmapgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// FinschiaItemTokenMillionArthursPropertySelect is the builder for selecting fields of FinschiaItemTokenMillionArthursProperty entities.
type FinschiaItemTokenMillionArthursPropertySelect struct {
	*FinschiaItemTokenMillionArthursPropertyQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (fitmaps *FinschiaItemTokenMillionArthursPropertySelect) Aggregate(fns ...AggregateFunc) *FinschiaItemTokenMillionArthursPropertySelect {
	fitmaps.fns = append(fitmaps.fns, fns...)
	return fitmaps
}

// Scan applies the selector query and scans the result into the given value.
func (fitmaps *FinschiaItemTokenMillionArthursPropertySelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, fitmaps.ctx, "Select")
	if err := fitmaps.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*FinschiaItemTokenMillionArthursPropertyQuery, *FinschiaItemTokenMillionArthursPropertySelect](ctx, fitmaps.FinschiaItemTokenMillionArthursPropertyQuery, fitmaps, fitmaps.inters, v)
}

func (fitmaps *FinschiaItemTokenMillionArthursPropertySelect) sqlScan(ctx context.Context, root *FinschiaItemTokenMillionArthursPropertyQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(fitmaps.fns))
	for _, fn := range fitmaps.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*fitmaps.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := fitmaps.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
