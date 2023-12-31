// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/a10adotapp/a10a.app/ent/linenft"
	"github.com/a10adotapp/a10a.app/ent/linenftactivity"
	"github.com/a10adotapp/a10a.app/ent/predicate"
)

// LINENFTActivityQuery is the builder for querying LINENFTActivity entities.
type LINENFTActivityQuery struct {
	config
	ctx         *QueryContext
	order       []linenftactivity.OrderOption
	inters      []Interceptor
	predicates  []predicate.LINENFTActivity
	withLineNft *LINENFTQuery
	withFKs     bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the LINENFTActivityQuery builder.
func (laq *LINENFTActivityQuery) Where(ps ...predicate.LINENFTActivity) *LINENFTActivityQuery {
	laq.predicates = append(laq.predicates, ps...)
	return laq
}

// Limit the number of records to be returned by this query.
func (laq *LINENFTActivityQuery) Limit(limit int) *LINENFTActivityQuery {
	laq.ctx.Limit = &limit
	return laq
}

// Offset to start from.
func (laq *LINENFTActivityQuery) Offset(offset int) *LINENFTActivityQuery {
	laq.ctx.Offset = &offset
	return laq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (laq *LINENFTActivityQuery) Unique(unique bool) *LINENFTActivityQuery {
	laq.ctx.Unique = &unique
	return laq
}

// Order specifies how the records should be ordered.
func (laq *LINENFTActivityQuery) Order(o ...linenftactivity.OrderOption) *LINENFTActivityQuery {
	laq.order = append(laq.order, o...)
	return laq
}

// QueryLineNft chains the current query on the "line_nft" edge.
func (laq *LINENFTActivityQuery) QueryLineNft() *LINENFTQuery {
	query := (&LINENFTClient{config: laq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := laq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := laq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(linenftactivity.Table, linenftactivity.FieldID, selector),
			sqlgraph.To(linenft.Table, linenft.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, linenftactivity.LineNftTable, linenftactivity.LineNftColumn),
		)
		fromU = sqlgraph.SetNeighbors(laq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first LINENFTActivity entity from the query.
// Returns a *NotFoundError when no LINENFTActivity was found.
func (laq *LINENFTActivityQuery) First(ctx context.Context) (*LINENFTActivity, error) {
	nodes, err := laq.Limit(1).All(setContextOp(ctx, laq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{linenftactivity.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (laq *LINENFTActivityQuery) FirstX(ctx context.Context) *LINENFTActivity {
	node, err := laq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first LINENFTActivity ID from the query.
// Returns a *NotFoundError when no LINENFTActivity ID was found.
func (laq *LINENFTActivityQuery) FirstID(ctx context.Context) (id uint32, err error) {
	var ids []uint32
	if ids, err = laq.Limit(1).IDs(setContextOp(ctx, laq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{linenftactivity.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (laq *LINENFTActivityQuery) FirstIDX(ctx context.Context) uint32 {
	id, err := laq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single LINENFTActivity entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one LINENFTActivity entity is found.
// Returns a *NotFoundError when no LINENFTActivity entities are found.
func (laq *LINENFTActivityQuery) Only(ctx context.Context) (*LINENFTActivity, error) {
	nodes, err := laq.Limit(2).All(setContextOp(ctx, laq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{linenftactivity.Label}
	default:
		return nil, &NotSingularError{linenftactivity.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (laq *LINENFTActivityQuery) OnlyX(ctx context.Context) *LINENFTActivity {
	node, err := laq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only LINENFTActivity ID in the query.
// Returns a *NotSingularError when more than one LINENFTActivity ID is found.
// Returns a *NotFoundError when no entities are found.
func (laq *LINENFTActivityQuery) OnlyID(ctx context.Context) (id uint32, err error) {
	var ids []uint32
	if ids, err = laq.Limit(2).IDs(setContextOp(ctx, laq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{linenftactivity.Label}
	default:
		err = &NotSingularError{linenftactivity.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (laq *LINENFTActivityQuery) OnlyIDX(ctx context.Context) uint32 {
	id, err := laq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of LINENFTActivities.
func (laq *LINENFTActivityQuery) All(ctx context.Context) ([]*LINENFTActivity, error) {
	ctx = setContextOp(ctx, laq.ctx, "All")
	if err := laq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*LINENFTActivity, *LINENFTActivityQuery]()
	return withInterceptors[[]*LINENFTActivity](ctx, laq, qr, laq.inters)
}

// AllX is like All, but panics if an error occurs.
func (laq *LINENFTActivityQuery) AllX(ctx context.Context) []*LINENFTActivity {
	nodes, err := laq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of LINENFTActivity IDs.
func (laq *LINENFTActivityQuery) IDs(ctx context.Context) (ids []uint32, err error) {
	if laq.ctx.Unique == nil && laq.path != nil {
		laq.Unique(true)
	}
	ctx = setContextOp(ctx, laq.ctx, "IDs")
	if err = laq.Select(linenftactivity.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (laq *LINENFTActivityQuery) IDsX(ctx context.Context) []uint32 {
	ids, err := laq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (laq *LINENFTActivityQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, laq.ctx, "Count")
	if err := laq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, laq, querierCount[*LINENFTActivityQuery](), laq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (laq *LINENFTActivityQuery) CountX(ctx context.Context) int {
	count, err := laq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (laq *LINENFTActivityQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, laq.ctx, "Exist")
	switch _, err := laq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (laq *LINENFTActivityQuery) ExistX(ctx context.Context) bool {
	exist, err := laq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the LINENFTActivityQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (laq *LINENFTActivityQuery) Clone() *LINENFTActivityQuery {
	if laq == nil {
		return nil
	}
	return &LINENFTActivityQuery{
		config:      laq.config,
		ctx:         laq.ctx.Clone(),
		order:       append([]linenftactivity.OrderOption{}, laq.order...),
		inters:      append([]Interceptor{}, laq.inters...),
		predicates:  append([]predicate.LINENFTActivity{}, laq.predicates...),
		withLineNft: laq.withLineNft.Clone(),
		// clone intermediate query.
		sql:  laq.sql.Clone(),
		path: laq.path,
	}
}

// WithLineNft tells the query-builder to eager-load the nodes that are connected to
// the "line_nft" edge. The optional arguments are used to configure the query builder of the edge.
func (laq *LINENFTActivityQuery) WithLineNft(opts ...func(*LINENFTQuery)) *LINENFTActivityQuery {
	query := (&LINENFTClient{config: laq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	laq.withLineNft = query
	return laq
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
//	client.LINENFTActivity.Query().
//		GroupBy(linenftactivity.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (laq *LINENFTActivityQuery) GroupBy(field string, fields ...string) *LINENFTActivityGroupBy {
	laq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &LINENFTActivityGroupBy{build: laq}
	grbuild.flds = &laq.ctx.Fields
	grbuild.label = linenftactivity.Label
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
//	client.LINENFTActivity.Query().
//		Select(linenftactivity.FieldCreatedAt).
//		Scan(ctx, &v)
func (laq *LINENFTActivityQuery) Select(fields ...string) *LINENFTActivitySelect {
	laq.ctx.Fields = append(laq.ctx.Fields, fields...)
	sbuild := &LINENFTActivitySelect{LINENFTActivityQuery: laq}
	sbuild.label = linenftactivity.Label
	sbuild.flds, sbuild.scan = &laq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a LINENFTActivitySelect configured with the given aggregations.
func (laq *LINENFTActivityQuery) Aggregate(fns ...AggregateFunc) *LINENFTActivitySelect {
	return laq.Select().Aggregate(fns...)
}

func (laq *LINENFTActivityQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range laq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, laq); err != nil {
				return err
			}
		}
	}
	for _, f := range laq.ctx.Fields {
		if !linenftactivity.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if laq.path != nil {
		prev, err := laq.path(ctx)
		if err != nil {
			return err
		}
		laq.sql = prev
	}
	return nil
}

func (laq *LINENFTActivityQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*LINENFTActivity, error) {
	var (
		nodes       = []*LINENFTActivity{}
		withFKs     = laq.withFKs
		_spec       = laq.querySpec()
		loadedTypes = [1]bool{
			laq.withLineNft != nil,
		}
	)
	if laq.withLineNft != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, linenftactivity.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*LINENFTActivity).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &LINENFTActivity{config: laq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, laq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := laq.withLineNft; query != nil {
		if err := laq.loadLineNft(ctx, query, nodes, nil,
			func(n *LINENFTActivity, e *LINENFT) { n.Edges.LineNft = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (laq *LINENFTActivityQuery) loadLineNft(ctx context.Context, query *LINENFTQuery, nodes []*LINENFTActivity, init func(*LINENFTActivity), assign func(*LINENFTActivity, *LINENFT)) error {
	ids := make([]uint32, 0, len(nodes))
	nodeids := make(map[uint32][]*LINENFTActivity)
	for i := range nodes {
		if nodes[i].linenft_activities == nil {
			continue
		}
		fk := *nodes[i].linenft_activities
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(linenft.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "linenft_activities" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (laq *LINENFTActivityQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := laq.querySpec()
	_spec.Node.Columns = laq.ctx.Fields
	if len(laq.ctx.Fields) > 0 {
		_spec.Unique = laq.ctx.Unique != nil && *laq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, laq.driver, _spec)
}

func (laq *LINENFTActivityQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(linenftactivity.Table, linenftactivity.Columns, sqlgraph.NewFieldSpec(linenftactivity.FieldID, field.TypeUint32))
	_spec.From = laq.sql
	if unique := laq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if laq.path != nil {
		_spec.Unique = true
	}
	if fields := laq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, linenftactivity.FieldID)
		for i := range fields {
			if fields[i] != linenftactivity.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := laq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := laq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := laq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := laq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (laq *LINENFTActivityQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(laq.driver.Dialect())
	t1 := builder.Table(linenftactivity.Table)
	columns := laq.ctx.Fields
	if len(columns) == 0 {
		columns = linenftactivity.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if laq.sql != nil {
		selector = laq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if laq.ctx.Unique != nil && *laq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range laq.predicates {
		p(selector)
	}
	for _, p := range laq.order {
		p(selector)
	}
	if offset := laq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := laq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// LINENFTActivityGroupBy is the group-by builder for LINENFTActivity entities.
type LINENFTActivityGroupBy struct {
	selector
	build *LINENFTActivityQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (lagb *LINENFTActivityGroupBy) Aggregate(fns ...AggregateFunc) *LINENFTActivityGroupBy {
	lagb.fns = append(lagb.fns, fns...)
	return lagb
}

// Scan applies the selector query and scans the result into the given value.
func (lagb *LINENFTActivityGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, lagb.build.ctx, "GroupBy")
	if err := lagb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*LINENFTActivityQuery, *LINENFTActivityGroupBy](ctx, lagb.build, lagb, lagb.build.inters, v)
}

func (lagb *LINENFTActivityGroupBy) sqlScan(ctx context.Context, root *LINENFTActivityQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(lagb.fns))
	for _, fn := range lagb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*lagb.flds)+len(lagb.fns))
		for _, f := range *lagb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*lagb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := lagb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// LINENFTActivitySelect is the builder for selecting fields of LINENFTActivity entities.
type LINENFTActivitySelect struct {
	*LINENFTActivityQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (las *LINENFTActivitySelect) Aggregate(fns ...AggregateFunc) *LINENFTActivitySelect {
	las.fns = append(las.fns, fns...)
	return las
}

// Scan applies the selector query and scans the result into the given value.
func (las *LINENFTActivitySelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, las.ctx, "Select")
	if err := las.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*LINENFTActivityQuery, *LINENFTActivitySelect](ctx, las.LINENFTActivityQuery, las, las.inters, v)
}

func (las *LINENFTActivitySelect) sqlScan(ctx context.Context, root *LINENFTActivityQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(las.fns))
	for _, fn := range las.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*las.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := las.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
