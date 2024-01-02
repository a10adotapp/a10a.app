// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/a10adotapp/a10a.app/ent/linenft"
	"github.com/a10adotapp/a10a.app/ent/predicate"
)

// LINENFTDelete is the builder for deleting a LINENFT entity.
type LINENFTDelete struct {
	config
	hooks    []Hook
	mutation *LINENFTMutation
}

// Where appends a list predicates to the LINENFTDelete builder.
func (ld *LINENFTDelete) Where(ps ...predicate.LINENFT) *LINENFTDelete {
	ld.mutation.Where(ps...)
	return ld
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ld *LINENFTDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, ld.sqlExec, ld.mutation, ld.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (ld *LINENFTDelete) ExecX(ctx context.Context) int {
	n, err := ld.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ld *LINENFTDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(linenft.Table, sqlgraph.NewFieldSpec(linenft.FieldID, field.TypeUint32))
	if ps := ld.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, ld.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	ld.mutation.done = true
	return affected, err
}

// LINENFTDeleteOne is the builder for deleting a single LINENFT entity.
type LINENFTDeleteOne struct {
	ld *LINENFTDelete
}

// Where appends a list predicates to the LINENFTDelete builder.
func (ldo *LINENFTDeleteOne) Where(ps ...predicate.LINENFT) *LINENFTDeleteOne {
	ldo.ld.mutation.Where(ps...)
	return ldo
}

// Exec executes the deletion query.
func (ldo *LINENFTDeleteOne) Exec(ctx context.Context) error {
	n, err := ldo.ld.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{linenft.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ldo *LINENFTDeleteOne) ExecX(ctx context.Context) {
	if err := ldo.Exec(ctx); err != nil {
		panic(err)
	}
}