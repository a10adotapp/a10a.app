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
	"github.com/a10adotapp/a10a.app/ent/finschiaitemtoken"
	"github.com/a10adotapp/a10a.app/ent/finschiaitemtokenactivity"
	"github.com/a10adotapp/a10a.app/ent/finschiaitemtokenmillionarthursproperty"
	"github.com/a10adotapp/a10a.app/ent/predicate"
)

// FinschiaItemTokenQuery is the builder for querying FinschiaItemToken entities.
type FinschiaItemTokenQuery struct {
	config
	ctx                          *QueryContext
	order                        []finschiaitemtoken.OrderOption
	inters                       []Interceptor
	predicates                   []predicate.FinschiaItemToken
	withActivities               *FinschiaItemTokenActivityQuery
	withMillionArthursProperties *FinschiaItemTokenMillionArthursPropertyQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the FinschiaItemTokenQuery builder.
func (fitq *FinschiaItemTokenQuery) Where(ps ...predicate.FinschiaItemToken) *FinschiaItemTokenQuery {
	fitq.predicates = append(fitq.predicates, ps...)
	return fitq
}

// Limit the number of records to be returned by this query.
func (fitq *FinschiaItemTokenQuery) Limit(limit int) *FinschiaItemTokenQuery {
	fitq.ctx.Limit = &limit
	return fitq
}

// Offset to start from.
func (fitq *FinschiaItemTokenQuery) Offset(offset int) *FinschiaItemTokenQuery {
	fitq.ctx.Offset = &offset
	return fitq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (fitq *FinschiaItemTokenQuery) Unique(unique bool) *FinschiaItemTokenQuery {
	fitq.ctx.Unique = &unique
	return fitq
}

// Order specifies how the records should be ordered.
func (fitq *FinschiaItemTokenQuery) Order(o ...finschiaitemtoken.OrderOption) *FinschiaItemTokenQuery {
	fitq.order = append(fitq.order, o...)
	return fitq
}

// QueryActivities chains the current query on the "activities" edge.
func (fitq *FinschiaItemTokenQuery) QueryActivities() *FinschiaItemTokenActivityQuery {
	query := (&FinschiaItemTokenActivityClient{config: fitq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := fitq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := fitq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(finschiaitemtoken.Table, finschiaitemtoken.FieldID, selector),
			sqlgraph.To(finschiaitemtokenactivity.Table, finschiaitemtokenactivity.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, finschiaitemtoken.ActivitiesTable, finschiaitemtoken.ActivitiesColumn),
		)
		fromU = sqlgraph.SetNeighbors(fitq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryMillionArthursProperties chains the current query on the "million_arthurs_properties" edge.
func (fitq *FinschiaItemTokenQuery) QueryMillionArthursProperties() *FinschiaItemTokenMillionArthursPropertyQuery {
	query := (&FinschiaItemTokenMillionArthursPropertyClient{config: fitq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := fitq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := fitq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(finschiaitemtoken.Table, finschiaitemtoken.FieldID, selector),
			sqlgraph.To(finschiaitemtokenmillionarthursproperty.Table, finschiaitemtokenmillionarthursproperty.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, finschiaitemtoken.MillionArthursPropertiesTable, finschiaitemtoken.MillionArthursPropertiesColumn),
		)
		fromU = sqlgraph.SetNeighbors(fitq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first FinschiaItemToken entity from the query.
// Returns a *NotFoundError when no FinschiaItemToken was found.
func (fitq *FinschiaItemTokenQuery) First(ctx context.Context) (*FinschiaItemToken, error) {
	nodes, err := fitq.Limit(1).All(setContextOp(ctx, fitq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{finschiaitemtoken.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (fitq *FinschiaItemTokenQuery) FirstX(ctx context.Context) *FinschiaItemToken {
	node, err := fitq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first FinschiaItemToken ID from the query.
// Returns a *NotFoundError when no FinschiaItemToken ID was found.
func (fitq *FinschiaItemTokenQuery) FirstID(ctx context.Context) (id uint32, err error) {
	var ids []uint32
	if ids, err = fitq.Limit(1).IDs(setContextOp(ctx, fitq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{finschiaitemtoken.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (fitq *FinschiaItemTokenQuery) FirstIDX(ctx context.Context) uint32 {
	id, err := fitq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single FinschiaItemToken entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one FinschiaItemToken entity is found.
// Returns a *NotFoundError when no FinschiaItemToken entities are found.
func (fitq *FinschiaItemTokenQuery) Only(ctx context.Context) (*FinschiaItemToken, error) {
	nodes, err := fitq.Limit(2).All(setContextOp(ctx, fitq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{finschiaitemtoken.Label}
	default:
		return nil, &NotSingularError{finschiaitemtoken.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (fitq *FinschiaItemTokenQuery) OnlyX(ctx context.Context) *FinschiaItemToken {
	node, err := fitq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only FinschiaItemToken ID in the query.
// Returns a *NotSingularError when more than one FinschiaItemToken ID is found.
// Returns a *NotFoundError when no entities are found.
func (fitq *FinschiaItemTokenQuery) OnlyID(ctx context.Context) (id uint32, err error) {
	var ids []uint32
	if ids, err = fitq.Limit(2).IDs(setContextOp(ctx, fitq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{finschiaitemtoken.Label}
	default:
		err = &NotSingularError{finschiaitemtoken.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (fitq *FinschiaItemTokenQuery) OnlyIDX(ctx context.Context) uint32 {
	id, err := fitq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of FinschiaItemTokens.
func (fitq *FinschiaItemTokenQuery) All(ctx context.Context) ([]*FinschiaItemToken, error) {
	ctx = setContextOp(ctx, fitq.ctx, "All")
	if err := fitq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*FinschiaItemToken, *FinschiaItemTokenQuery]()
	return withInterceptors[[]*FinschiaItemToken](ctx, fitq, qr, fitq.inters)
}

// AllX is like All, but panics if an error occurs.
func (fitq *FinschiaItemTokenQuery) AllX(ctx context.Context) []*FinschiaItemToken {
	nodes, err := fitq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of FinschiaItemToken IDs.
func (fitq *FinschiaItemTokenQuery) IDs(ctx context.Context) (ids []uint32, err error) {
	if fitq.ctx.Unique == nil && fitq.path != nil {
		fitq.Unique(true)
	}
	ctx = setContextOp(ctx, fitq.ctx, "IDs")
	if err = fitq.Select(finschiaitemtoken.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (fitq *FinschiaItemTokenQuery) IDsX(ctx context.Context) []uint32 {
	ids, err := fitq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (fitq *FinschiaItemTokenQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, fitq.ctx, "Count")
	if err := fitq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, fitq, querierCount[*FinschiaItemTokenQuery](), fitq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (fitq *FinschiaItemTokenQuery) CountX(ctx context.Context) int {
	count, err := fitq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (fitq *FinschiaItemTokenQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, fitq.ctx, "Exist")
	switch _, err := fitq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (fitq *FinschiaItemTokenQuery) ExistX(ctx context.Context) bool {
	exist, err := fitq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the FinschiaItemTokenQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (fitq *FinschiaItemTokenQuery) Clone() *FinschiaItemTokenQuery {
	if fitq == nil {
		return nil
	}
	return &FinschiaItemTokenQuery{
		config:                       fitq.config,
		ctx:                          fitq.ctx.Clone(),
		order:                        append([]finschiaitemtoken.OrderOption{}, fitq.order...),
		inters:                       append([]Interceptor{}, fitq.inters...),
		predicates:                   append([]predicate.FinschiaItemToken{}, fitq.predicates...),
		withActivities:               fitq.withActivities.Clone(),
		withMillionArthursProperties: fitq.withMillionArthursProperties.Clone(),
		// clone intermediate query.
		sql:  fitq.sql.Clone(),
		path: fitq.path,
	}
}

// WithActivities tells the query-builder to eager-load the nodes that are connected to
// the "activities" edge. The optional arguments are used to configure the query builder of the edge.
func (fitq *FinschiaItemTokenQuery) WithActivities(opts ...func(*FinschiaItemTokenActivityQuery)) *FinschiaItemTokenQuery {
	query := (&FinschiaItemTokenActivityClient{config: fitq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	fitq.withActivities = query
	return fitq
}

// WithMillionArthursProperties tells the query-builder to eager-load the nodes that are connected to
// the "million_arthurs_properties" edge. The optional arguments are used to configure the query builder of the edge.
func (fitq *FinschiaItemTokenQuery) WithMillionArthursProperties(opts ...func(*FinschiaItemTokenMillionArthursPropertyQuery)) *FinschiaItemTokenQuery {
	query := (&FinschiaItemTokenMillionArthursPropertyClient{config: fitq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	fitq.withMillionArthursProperties = query
	return fitq
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
//	client.FinschiaItemToken.Query().
//		GroupBy(finschiaitemtoken.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (fitq *FinschiaItemTokenQuery) GroupBy(field string, fields ...string) *FinschiaItemTokenGroupBy {
	fitq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &FinschiaItemTokenGroupBy{build: fitq}
	grbuild.flds = &fitq.ctx.Fields
	grbuild.label = finschiaitemtoken.Label
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
//	client.FinschiaItemToken.Query().
//		Select(finschiaitemtoken.FieldCreatedAt).
//		Scan(ctx, &v)
func (fitq *FinschiaItemTokenQuery) Select(fields ...string) *FinschiaItemTokenSelect {
	fitq.ctx.Fields = append(fitq.ctx.Fields, fields...)
	sbuild := &FinschiaItemTokenSelect{FinschiaItemTokenQuery: fitq}
	sbuild.label = finschiaitemtoken.Label
	sbuild.flds, sbuild.scan = &fitq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a FinschiaItemTokenSelect configured with the given aggregations.
func (fitq *FinschiaItemTokenQuery) Aggregate(fns ...AggregateFunc) *FinschiaItemTokenSelect {
	return fitq.Select().Aggregate(fns...)
}

func (fitq *FinschiaItemTokenQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range fitq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, fitq); err != nil {
				return err
			}
		}
	}
	for _, f := range fitq.ctx.Fields {
		if !finschiaitemtoken.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if fitq.path != nil {
		prev, err := fitq.path(ctx)
		if err != nil {
			return err
		}
		fitq.sql = prev
	}
	return nil
}

func (fitq *FinschiaItemTokenQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*FinschiaItemToken, error) {
	var (
		nodes       = []*FinschiaItemToken{}
		_spec       = fitq.querySpec()
		loadedTypes = [2]bool{
			fitq.withActivities != nil,
			fitq.withMillionArthursProperties != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*FinschiaItemToken).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &FinschiaItemToken{config: fitq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, fitq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := fitq.withActivities; query != nil {
		if err := fitq.loadActivities(ctx, query, nodes,
			func(n *FinschiaItemToken) { n.Edges.Activities = []*FinschiaItemTokenActivity{} },
			func(n *FinschiaItemToken, e *FinschiaItemTokenActivity) {
				n.Edges.Activities = append(n.Edges.Activities, e)
			}); err != nil {
			return nil, err
		}
	}
	if query := fitq.withMillionArthursProperties; query != nil {
		if err := fitq.loadMillionArthursProperties(ctx, query, nodes,
			func(n *FinschiaItemToken) {
				n.Edges.MillionArthursProperties = []*FinschiaItemTokenMillionArthursProperty{}
			},
			func(n *FinschiaItemToken, e *FinschiaItemTokenMillionArthursProperty) {
				n.Edges.MillionArthursProperties = append(n.Edges.MillionArthursProperties, e)
			}); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (fitq *FinschiaItemTokenQuery) loadActivities(ctx context.Context, query *FinschiaItemTokenActivityQuery, nodes []*FinschiaItemToken, init func(*FinschiaItemToken), assign func(*FinschiaItemToken, *FinschiaItemTokenActivity)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[uint32]*FinschiaItemToken)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(finschiaitemtokenactivity.FieldFinschiaItemTokenID)
	}
	query.Where(predicate.FinschiaItemTokenActivity(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(finschiaitemtoken.ActivitiesColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.FinschiaItemTokenID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "finschia_item_token_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (fitq *FinschiaItemTokenQuery) loadMillionArthursProperties(ctx context.Context, query *FinschiaItemTokenMillionArthursPropertyQuery, nodes []*FinschiaItemToken, init func(*FinschiaItemToken), assign func(*FinschiaItemToken, *FinschiaItemTokenMillionArthursProperty)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[uint32]*FinschiaItemToken)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(finschiaitemtokenmillionarthursproperty.FieldFinschiaItemTokenID)
	}
	query.Where(predicate.FinschiaItemTokenMillionArthursProperty(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(finschiaitemtoken.MillionArthursPropertiesColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.FinschiaItemTokenID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "finschia_item_token_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (fitq *FinschiaItemTokenQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := fitq.querySpec()
	_spec.Node.Columns = fitq.ctx.Fields
	if len(fitq.ctx.Fields) > 0 {
		_spec.Unique = fitq.ctx.Unique != nil && *fitq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, fitq.driver, _spec)
}

func (fitq *FinschiaItemTokenQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(finschiaitemtoken.Table, finschiaitemtoken.Columns, sqlgraph.NewFieldSpec(finschiaitemtoken.FieldID, field.TypeUint32))
	_spec.From = fitq.sql
	if unique := fitq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if fitq.path != nil {
		_spec.Unique = true
	}
	if fields := fitq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, finschiaitemtoken.FieldID)
		for i := range fields {
			if fields[i] != finschiaitemtoken.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := fitq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := fitq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := fitq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := fitq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (fitq *FinschiaItemTokenQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(fitq.driver.Dialect())
	t1 := builder.Table(finschiaitemtoken.Table)
	columns := fitq.ctx.Fields
	if len(columns) == 0 {
		columns = finschiaitemtoken.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if fitq.sql != nil {
		selector = fitq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if fitq.ctx.Unique != nil && *fitq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range fitq.predicates {
		p(selector)
	}
	for _, p := range fitq.order {
		p(selector)
	}
	if offset := fitq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := fitq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// FinschiaItemTokenGroupBy is the group-by builder for FinschiaItemToken entities.
type FinschiaItemTokenGroupBy struct {
	selector
	build *FinschiaItemTokenQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (fitgb *FinschiaItemTokenGroupBy) Aggregate(fns ...AggregateFunc) *FinschiaItemTokenGroupBy {
	fitgb.fns = append(fitgb.fns, fns...)
	return fitgb
}

// Scan applies the selector query and scans the result into the given value.
func (fitgb *FinschiaItemTokenGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, fitgb.build.ctx, "GroupBy")
	if err := fitgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*FinschiaItemTokenQuery, *FinschiaItemTokenGroupBy](ctx, fitgb.build, fitgb, fitgb.build.inters, v)
}

func (fitgb *FinschiaItemTokenGroupBy) sqlScan(ctx context.Context, root *FinschiaItemTokenQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(fitgb.fns))
	for _, fn := range fitgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*fitgb.flds)+len(fitgb.fns))
		for _, f := range *fitgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*fitgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := fitgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// FinschiaItemTokenSelect is the builder for selecting fields of FinschiaItemToken entities.
type FinschiaItemTokenSelect struct {
	*FinschiaItemTokenQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (fits *FinschiaItemTokenSelect) Aggregate(fns ...AggregateFunc) *FinschiaItemTokenSelect {
	fits.fns = append(fits.fns, fns...)
	return fits
}

// Scan applies the selector query and scans the result into the given value.
func (fits *FinschiaItemTokenSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, fits.ctx, "Select")
	if err := fits.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*FinschiaItemTokenQuery, *FinschiaItemTokenSelect](ctx, fits.FinschiaItemTokenQuery, fits, fits.inters, v)
}

func (fits *FinschiaItemTokenSelect) sqlScan(ctx context.Context, root *FinschiaItemTokenQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(fits.fns))
	for _, fn := range fits.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*fits.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := fits.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
