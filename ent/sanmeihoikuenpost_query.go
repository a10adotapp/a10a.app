// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/a10adotapp/a10a.app/ent/predicate"
	"github.com/a10adotapp/a10a.app/ent/sanmeihoikuenpost"
)

// SanmeiHoikuenPostQuery is the builder for querying SanmeiHoikuenPost entities.
type SanmeiHoikuenPostQuery struct {
	config
	ctx        *QueryContext
	order      []sanmeihoikuenpost.OrderOption
	inters     []Interceptor
	predicates []predicate.SanmeiHoikuenPost
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the SanmeiHoikuenPostQuery builder.
func (shpq *SanmeiHoikuenPostQuery) Where(ps ...predicate.SanmeiHoikuenPost) *SanmeiHoikuenPostQuery {
	shpq.predicates = append(shpq.predicates, ps...)
	return shpq
}

// Limit the number of records to be returned by this query.
func (shpq *SanmeiHoikuenPostQuery) Limit(limit int) *SanmeiHoikuenPostQuery {
	shpq.ctx.Limit = &limit
	return shpq
}

// Offset to start from.
func (shpq *SanmeiHoikuenPostQuery) Offset(offset int) *SanmeiHoikuenPostQuery {
	shpq.ctx.Offset = &offset
	return shpq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (shpq *SanmeiHoikuenPostQuery) Unique(unique bool) *SanmeiHoikuenPostQuery {
	shpq.ctx.Unique = &unique
	return shpq
}

// Order specifies how the records should be ordered.
func (shpq *SanmeiHoikuenPostQuery) Order(o ...sanmeihoikuenpost.OrderOption) *SanmeiHoikuenPostQuery {
	shpq.order = append(shpq.order, o...)
	return shpq
}

// First returns the first SanmeiHoikuenPost entity from the query.
// Returns a *NotFoundError when no SanmeiHoikuenPost was found.
func (shpq *SanmeiHoikuenPostQuery) First(ctx context.Context) (*SanmeiHoikuenPost, error) {
	nodes, err := shpq.Limit(1).All(setContextOp(ctx, shpq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{sanmeihoikuenpost.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (shpq *SanmeiHoikuenPostQuery) FirstX(ctx context.Context) *SanmeiHoikuenPost {
	node, err := shpq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first SanmeiHoikuenPost ID from the query.
// Returns a *NotFoundError when no SanmeiHoikuenPost ID was found.
func (shpq *SanmeiHoikuenPostQuery) FirstID(ctx context.Context) (id uint32, err error) {
	var ids []uint32
	if ids, err = shpq.Limit(1).IDs(setContextOp(ctx, shpq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{sanmeihoikuenpost.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (shpq *SanmeiHoikuenPostQuery) FirstIDX(ctx context.Context) uint32 {
	id, err := shpq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single SanmeiHoikuenPost entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one SanmeiHoikuenPost entity is found.
// Returns a *NotFoundError when no SanmeiHoikuenPost entities are found.
func (shpq *SanmeiHoikuenPostQuery) Only(ctx context.Context) (*SanmeiHoikuenPost, error) {
	nodes, err := shpq.Limit(2).All(setContextOp(ctx, shpq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{sanmeihoikuenpost.Label}
	default:
		return nil, &NotSingularError{sanmeihoikuenpost.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (shpq *SanmeiHoikuenPostQuery) OnlyX(ctx context.Context) *SanmeiHoikuenPost {
	node, err := shpq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only SanmeiHoikuenPost ID in the query.
// Returns a *NotSingularError when more than one SanmeiHoikuenPost ID is found.
// Returns a *NotFoundError when no entities are found.
func (shpq *SanmeiHoikuenPostQuery) OnlyID(ctx context.Context) (id uint32, err error) {
	var ids []uint32
	if ids, err = shpq.Limit(2).IDs(setContextOp(ctx, shpq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{sanmeihoikuenpost.Label}
	default:
		err = &NotSingularError{sanmeihoikuenpost.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (shpq *SanmeiHoikuenPostQuery) OnlyIDX(ctx context.Context) uint32 {
	id, err := shpq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of SanmeiHoikuenPosts.
func (shpq *SanmeiHoikuenPostQuery) All(ctx context.Context) ([]*SanmeiHoikuenPost, error) {
	ctx = setContextOp(ctx, shpq.ctx, "All")
	if err := shpq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*SanmeiHoikuenPost, *SanmeiHoikuenPostQuery]()
	return withInterceptors[[]*SanmeiHoikuenPost](ctx, shpq, qr, shpq.inters)
}

// AllX is like All, but panics if an error occurs.
func (shpq *SanmeiHoikuenPostQuery) AllX(ctx context.Context) []*SanmeiHoikuenPost {
	nodes, err := shpq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of SanmeiHoikuenPost IDs.
func (shpq *SanmeiHoikuenPostQuery) IDs(ctx context.Context) (ids []uint32, err error) {
	if shpq.ctx.Unique == nil && shpq.path != nil {
		shpq.Unique(true)
	}
	ctx = setContextOp(ctx, shpq.ctx, "IDs")
	if err = shpq.Select(sanmeihoikuenpost.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (shpq *SanmeiHoikuenPostQuery) IDsX(ctx context.Context) []uint32 {
	ids, err := shpq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (shpq *SanmeiHoikuenPostQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, shpq.ctx, "Count")
	if err := shpq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, shpq, querierCount[*SanmeiHoikuenPostQuery](), shpq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (shpq *SanmeiHoikuenPostQuery) CountX(ctx context.Context) int {
	count, err := shpq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (shpq *SanmeiHoikuenPostQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, shpq.ctx, "Exist")
	switch _, err := shpq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (shpq *SanmeiHoikuenPostQuery) ExistX(ctx context.Context) bool {
	exist, err := shpq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the SanmeiHoikuenPostQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (shpq *SanmeiHoikuenPostQuery) Clone() *SanmeiHoikuenPostQuery {
	if shpq == nil {
		return nil
	}
	return &SanmeiHoikuenPostQuery{
		config:     shpq.config,
		ctx:        shpq.ctx.Clone(),
		order:      append([]sanmeihoikuenpost.OrderOption{}, shpq.order...),
		inters:     append([]Interceptor{}, shpq.inters...),
		predicates: append([]predicate.SanmeiHoikuenPost{}, shpq.predicates...),
		// clone intermediate query.
		sql:  shpq.sql.Clone(),
		path: shpq.path,
	}
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
//	client.SanmeiHoikuenPost.Query().
//		GroupBy(sanmeihoikuenpost.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (shpq *SanmeiHoikuenPostQuery) GroupBy(field string, fields ...string) *SanmeiHoikuenPostGroupBy {
	shpq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &SanmeiHoikuenPostGroupBy{build: shpq}
	grbuild.flds = &shpq.ctx.Fields
	grbuild.label = sanmeihoikuenpost.Label
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
//	client.SanmeiHoikuenPost.Query().
//		Select(sanmeihoikuenpost.FieldCreatedAt).
//		Scan(ctx, &v)
func (shpq *SanmeiHoikuenPostQuery) Select(fields ...string) *SanmeiHoikuenPostSelect {
	shpq.ctx.Fields = append(shpq.ctx.Fields, fields...)
	sbuild := &SanmeiHoikuenPostSelect{SanmeiHoikuenPostQuery: shpq}
	sbuild.label = sanmeihoikuenpost.Label
	sbuild.flds, sbuild.scan = &shpq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a SanmeiHoikuenPostSelect configured with the given aggregations.
func (shpq *SanmeiHoikuenPostQuery) Aggregate(fns ...AggregateFunc) *SanmeiHoikuenPostSelect {
	return shpq.Select().Aggregate(fns...)
}

func (shpq *SanmeiHoikuenPostQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range shpq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, shpq); err != nil {
				return err
			}
		}
	}
	for _, f := range shpq.ctx.Fields {
		if !sanmeihoikuenpost.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if shpq.path != nil {
		prev, err := shpq.path(ctx)
		if err != nil {
			return err
		}
		shpq.sql = prev
	}
	return nil
}

func (shpq *SanmeiHoikuenPostQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*SanmeiHoikuenPost, error) {
	var (
		nodes = []*SanmeiHoikuenPost{}
		_spec = shpq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*SanmeiHoikuenPost).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &SanmeiHoikuenPost{config: shpq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, shpq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (shpq *SanmeiHoikuenPostQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := shpq.querySpec()
	_spec.Node.Columns = shpq.ctx.Fields
	if len(shpq.ctx.Fields) > 0 {
		_spec.Unique = shpq.ctx.Unique != nil && *shpq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, shpq.driver, _spec)
}

func (shpq *SanmeiHoikuenPostQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(sanmeihoikuenpost.Table, sanmeihoikuenpost.Columns, sqlgraph.NewFieldSpec(sanmeihoikuenpost.FieldID, field.TypeUint32))
	_spec.From = shpq.sql
	if unique := shpq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if shpq.path != nil {
		_spec.Unique = true
	}
	if fields := shpq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, sanmeihoikuenpost.FieldID)
		for i := range fields {
			if fields[i] != sanmeihoikuenpost.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := shpq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := shpq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := shpq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := shpq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (shpq *SanmeiHoikuenPostQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(shpq.driver.Dialect())
	t1 := builder.Table(sanmeihoikuenpost.Table)
	columns := shpq.ctx.Fields
	if len(columns) == 0 {
		columns = sanmeihoikuenpost.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if shpq.sql != nil {
		selector = shpq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if shpq.ctx.Unique != nil && *shpq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range shpq.predicates {
		p(selector)
	}
	for _, p := range shpq.order {
		p(selector)
	}
	if offset := shpq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := shpq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// SanmeiHoikuenPostGroupBy is the group-by builder for SanmeiHoikuenPost entities.
type SanmeiHoikuenPostGroupBy struct {
	selector
	build *SanmeiHoikuenPostQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (shpgb *SanmeiHoikuenPostGroupBy) Aggregate(fns ...AggregateFunc) *SanmeiHoikuenPostGroupBy {
	shpgb.fns = append(shpgb.fns, fns...)
	return shpgb
}

// Scan applies the selector query and scans the result into the given value.
func (shpgb *SanmeiHoikuenPostGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, shpgb.build.ctx, "GroupBy")
	if err := shpgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*SanmeiHoikuenPostQuery, *SanmeiHoikuenPostGroupBy](ctx, shpgb.build, shpgb, shpgb.build.inters, v)
}

func (shpgb *SanmeiHoikuenPostGroupBy) sqlScan(ctx context.Context, root *SanmeiHoikuenPostQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(shpgb.fns))
	for _, fn := range shpgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*shpgb.flds)+len(shpgb.fns))
		for _, f := range *shpgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*shpgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := shpgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// SanmeiHoikuenPostSelect is the builder for selecting fields of SanmeiHoikuenPost entities.
type SanmeiHoikuenPostSelect struct {
	*SanmeiHoikuenPostQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (shps *SanmeiHoikuenPostSelect) Aggregate(fns ...AggregateFunc) *SanmeiHoikuenPostSelect {
	shps.fns = append(shps.fns, fns...)
	return shps
}

// Scan applies the selector query and scans the result into the given value.
func (shps *SanmeiHoikuenPostSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, shps.ctx, "Select")
	if err := shps.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*SanmeiHoikuenPostQuery, *SanmeiHoikuenPostSelect](ctx, shps.SanmeiHoikuenPostQuery, shps, shps.inters, v)
}

func (shps *SanmeiHoikuenPostSelect) sqlScan(ctx context.Context, root *SanmeiHoikuenPostQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(shps.fns))
	for _, fn := range shps.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*shps.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := shps.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
