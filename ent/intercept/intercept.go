// Code generated by ent, DO NOT EDIT.

package intercept

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"github.com/a10adotapp/a10a.app/ent"
	"github.com/a10adotapp/a10a.app/ent/changokushiweapon"
	"github.com/a10adotapp/a10a.app/ent/changokushiweaponchangelog"
	"github.com/a10adotapp/a10a.app/ent/predicate"
	"github.com/a10adotapp/a10a.app/ent/sanmeihoikuenpost"
)

// The Query interface represents an operation that queries a graph.
// By using this interface, users can write generic code that manipulates
// query builders of different types.
type Query interface {
	// Type returns the string representation of the query type.
	Type() string
	// Limit the number of records to be returned by this query.
	Limit(int)
	// Offset to start from.
	Offset(int)
	// Unique configures the query builder to filter duplicate records.
	Unique(bool)
	// Order specifies how the records should be ordered.
	Order(...func(*sql.Selector))
	// WhereP appends storage-level predicates to the query builder. Using this method, users
	// can use type-assertion to append predicates that do not depend on any generated package.
	WhereP(...func(*sql.Selector))
}

// The Func type is an adapter that allows ordinary functions to be used as interceptors.
// Unlike traversal functions, interceptors are skipped during graph traversals. Note that the
// implementation of Func is different from the one defined in entgo.io/ent.InterceptFunc.
type Func func(context.Context, Query) error

// Intercept calls f(ctx, q) and then applied the next Querier.
func (f Func) Intercept(next ent.Querier) ent.Querier {
	return ent.QuerierFunc(func(ctx context.Context, q ent.Query) (ent.Value, error) {
		query, err := NewQuery(q)
		if err != nil {
			return nil, err
		}
		if err := f(ctx, query); err != nil {
			return nil, err
		}
		return next.Query(ctx, q)
	})
}

// The TraverseFunc type is an adapter to allow the use of ordinary function as Traverser.
// If f is a function with the appropriate signature, TraverseFunc(f) is a Traverser that calls f.
type TraverseFunc func(context.Context, Query) error

// Intercept is a dummy implementation of Intercept that returns the next Querier in the pipeline.
func (f TraverseFunc) Intercept(next ent.Querier) ent.Querier {
	return next
}

// Traverse calls f(ctx, q).
func (f TraverseFunc) Traverse(ctx context.Context, q ent.Query) error {
	query, err := NewQuery(q)
	if err != nil {
		return err
	}
	return f(ctx, query)
}

// The ChangokushiWeaponFunc type is an adapter to allow the use of ordinary function as a Querier.
type ChangokushiWeaponFunc func(context.Context, *ent.ChangokushiWeaponQuery) (ent.Value, error)

// Query calls f(ctx, q).
func (f ChangokushiWeaponFunc) Query(ctx context.Context, q ent.Query) (ent.Value, error) {
	if q, ok := q.(*ent.ChangokushiWeaponQuery); ok {
		return f(ctx, q)
	}
	return nil, fmt.Errorf("unexpected query type %T. expect *ent.ChangokushiWeaponQuery", q)
}

// The TraverseChangokushiWeapon type is an adapter to allow the use of ordinary function as Traverser.
type TraverseChangokushiWeapon func(context.Context, *ent.ChangokushiWeaponQuery) error

// Intercept is a dummy implementation of Intercept that returns the next Querier in the pipeline.
func (f TraverseChangokushiWeapon) Intercept(next ent.Querier) ent.Querier {
	return next
}

// Traverse calls f(ctx, q).
func (f TraverseChangokushiWeapon) Traverse(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.ChangokushiWeaponQuery); ok {
		return f(ctx, q)
	}
	return fmt.Errorf("unexpected query type %T. expect *ent.ChangokushiWeaponQuery", q)
}

// The ChangokushiWeaponChangeLogFunc type is an adapter to allow the use of ordinary function as a Querier.
type ChangokushiWeaponChangeLogFunc func(context.Context, *ent.ChangokushiWeaponChangeLogQuery) (ent.Value, error)

// Query calls f(ctx, q).
func (f ChangokushiWeaponChangeLogFunc) Query(ctx context.Context, q ent.Query) (ent.Value, error) {
	if q, ok := q.(*ent.ChangokushiWeaponChangeLogQuery); ok {
		return f(ctx, q)
	}
	return nil, fmt.Errorf("unexpected query type %T. expect *ent.ChangokushiWeaponChangeLogQuery", q)
}

// The TraverseChangokushiWeaponChangeLog type is an adapter to allow the use of ordinary function as Traverser.
type TraverseChangokushiWeaponChangeLog func(context.Context, *ent.ChangokushiWeaponChangeLogQuery) error

// Intercept is a dummy implementation of Intercept that returns the next Querier in the pipeline.
func (f TraverseChangokushiWeaponChangeLog) Intercept(next ent.Querier) ent.Querier {
	return next
}

// Traverse calls f(ctx, q).
func (f TraverseChangokushiWeaponChangeLog) Traverse(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.ChangokushiWeaponChangeLogQuery); ok {
		return f(ctx, q)
	}
	return fmt.Errorf("unexpected query type %T. expect *ent.ChangokushiWeaponChangeLogQuery", q)
}

// The SanmeiHoikuenPostFunc type is an adapter to allow the use of ordinary function as a Querier.
type SanmeiHoikuenPostFunc func(context.Context, *ent.SanmeiHoikuenPostQuery) (ent.Value, error)

// Query calls f(ctx, q).
func (f SanmeiHoikuenPostFunc) Query(ctx context.Context, q ent.Query) (ent.Value, error) {
	if q, ok := q.(*ent.SanmeiHoikuenPostQuery); ok {
		return f(ctx, q)
	}
	return nil, fmt.Errorf("unexpected query type %T. expect *ent.SanmeiHoikuenPostQuery", q)
}

// The TraverseSanmeiHoikuenPost type is an adapter to allow the use of ordinary function as Traverser.
type TraverseSanmeiHoikuenPost func(context.Context, *ent.SanmeiHoikuenPostQuery) error

// Intercept is a dummy implementation of Intercept that returns the next Querier in the pipeline.
func (f TraverseSanmeiHoikuenPost) Intercept(next ent.Querier) ent.Querier {
	return next
}

// Traverse calls f(ctx, q).
func (f TraverseSanmeiHoikuenPost) Traverse(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.SanmeiHoikuenPostQuery); ok {
		return f(ctx, q)
	}
	return fmt.Errorf("unexpected query type %T. expect *ent.SanmeiHoikuenPostQuery", q)
}

// NewQuery returns the generic Query interface for the given typed query.
func NewQuery(q ent.Query) (Query, error) {
	switch q := q.(type) {
	case *ent.ChangokushiWeaponQuery:
		return &query[*ent.ChangokushiWeaponQuery, predicate.ChangokushiWeapon, changokushiweapon.OrderOption]{typ: ent.TypeChangokushiWeapon, tq: q}, nil
	case *ent.ChangokushiWeaponChangeLogQuery:
		return &query[*ent.ChangokushiWeaponChangeLogQuery, predicate.ChangokushiWeaponChangeLog, changokushiweaponchangelog.OrderOption]{typ: ent.TypeChangokushiWeaponChangeLog, tq: q}, nil
	case *ent.SanmeiHoikuenPostQuery:
		return &query[*ent.SanmeiHoikuenPostQuery, predicate.SanmeiHoikuenPost, sanmeihoikuenpost.OrderOption]{typ: ent.TypeSanmeiHoikuenPost, tq: q}, nil
	default:
		return nil, fmt.Errorf("unknown query type %T", q)
	}
}

type query[T any, P ~func(*sql.Selector), R ~func(*sql.Selector)] struct {
	typ string
	tq  interface {
		Limit(int) T
		Offset(int) T
		Unique(bool) T
		Order(...R) T
		Where(...P) T
	}
}

func (q query[T, P, R]) Type() string {
	return q.typ
}

func (q query[T, P, R]) Limit(limit int) {
	q.tq.Limit(limit)
}

func (q query[T, P, R]) Offset(offset int) {
	q.tq.Offset(offset)
}

func (q query[T, P, R]) Unique(unique bool) {
	q.tq.Unique(unique)
}

func (q query[T, P, R]) Order(orders ...func(*sql.Selector)) {
	rs := make([]R, len(orders))
	for i := range orders {
		rs[i] = orders[i]
	}
	q.tq.Order(rs...)
}

func (q query[T, P, R]) WhereP(ps ...func(*sql.Selector)) {
	p := make([]P, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	q.tq.Where(p...)
}
