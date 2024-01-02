// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/a10adotapp/a10a.app/ent/linenft"
	"github.com/a10adotapp/a10a.app/ent/linenftactivity"
	"github.com/a10adotapp/a10a.app/ent/linenftmillionarthursproperty"
)

// LINENFTCreate is the builder for creating a LINENFT entity.
type LINENFTCreate struct {
	config
	mutation *LINENFTMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (lc *LINENFTCreate) SetCreatedAt(t time.Time) *LINENFTCreate {
	lc.mutation.SetCreatedAt(t)
	return lc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (lc *LINENFTCreate) SetNillableCreatedAt(t *time.Time) *LINENFTCreate {
	if t != nil {
		lc.SetCreatedAt(*t)
	}
	return lc
}

// SetUpdatedAt sets the "updated_at" field.
func (lc *LINENFTCreate) SetUpdatedAt(t time.Time) *LINENFTCreate {
	lc.mutation.SetUpdatedAt(t)
	return lc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (lc *LINENFTCreate) SetNillableUpdatedAt(t *time.Time) *LINENFTCreate {
	if t != nil {
		lc.SetUpdatedAt(*t)
	}
	return lc
}

// SetDeletedAt sets the "deleted_at" field.
func (lc *LINENFTCreate) SetDeletedAt(t time.Time) *LINENFTCreate {
	lc.mutation.SetDeletedAt(t)
	return lc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (lc *LINENFTCreate) SetNillableDeletedAt(t *time.Time) *LINENFTCreate {
	if t != nil {
		lc.SetDeletedAt(*t)
	}
	return lc
}

// SetLineNftID sets the "line_nft_id" field.
func (lc *LINENFTCreate) SetLineNftID(u uint32) *LINENFTCreate {
	lc.mutation.SetLineNftID(u)
	return lc
}

// SetContractID sets the "contract_id" field.
func (lc *LINENFTCreate) SetContractID(s string) *LINENFTCreate {
	lc.mutation.SetContractID(s)
	return lc
}

// SetTokenType sets the "token_type" field.
func (lc *LINENFTCreate) SetTokenType(s string) *LINENFTCreate {
	lc.mutation.SetTokenType(s)
	return lc
}

// SetTokenIndex sets the "token_index" field.
func (lc *LINENFTCreate) SetTokenIndex(s string) *LINENFTCreate {
	lc.mutation.SetTokenIndex(s)
	return lc
}

// SetTokenName sets the "token_name" field.
func (lc *LINENFTCreate) SetTokenName(s string) *LINENFTCreate {
	lc.mutation.SetTokenName(s)
	return lc
}

// SetTokenDescription sets the "token_description" field.
func (lc *LINENFTCreate) SetTokenDescription(s string) *LINENFTCreate {
	lc.mutation.SetTokenDescription(s)
	return lc
}

// SetTokenContentURL sets the "token_content_url" field.
func (lc *LINENFTCreate) SetTokenContentURL(s string) *LINENFTCreate {
	lc.mutation.SetTokenContentURL(s)
	return lc
}

// SetID sets the "id" field.
func (lc *LINENFTCreate) SetID(u uint32) *LINENFTCreate {
	lc.mutation.SetID(u)
	return lc
}

// AddActivityIDs adds the "activities" edge to the LINENFTActivity entity by IDs.
func (lc *LINENFTCreate) AddActivityIDs(ids ...uint32) *LINENFTCreate {
	lc.mutation.AddActivityIDs(ids...)
	return lc
}

// AddActivities adds the "activities" edges to the LINENFTActivity entity.
func (lc *LINENFTCreate) AddActivities(l ...*LINENFTActivity) *LINENFTCreate {
	ids := make([]uint32, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return lc.AddActivityIDs(ids...)
}

// SetMillionArthursPropertyID sets the "million_arthurs_property" edge to the LINENFTMillionArthursProperty entity by ID.
func (lc *LINENFTCreate) SetMillionArthursPropertyID(id uint32) *LINENFTCreate {
	lc.mutation.SetMillionArthursPropertyID(id)
	return lc
}

// SetNillableMillionArthursPropertyID sets the "million_arthurs_property" edge to the LINENFTMillionArthursProperty entity by ID if the given value is not nil.
func (lc *LINENFTCreate) SetNillableMillionArthursPropertyID(id *uint32) *LINENFTCreate {
	if id != nil {
		lc = lc.SetMillionArthursPropertyID(*id)
	}
	return lc
}

// SetMillionArthursProperty sets the "million_arthurs_property" edge to the LINENFTMillionArthursProperty entity.
func (lc *LINENFTCreate) SetMillionArthursProperty(l *LINENFTMillionArthursProperty) *LINENFTCreate {
	return lc.SetMillionArthursPropertyID(l.ID)
}

// Mutation returns the LINENFTMutation object of the builder.
func (lc *LINENFTCreate) Mutation() *LINENFTMutation {
	return lc.mutation
}

// Save creates the LINENFT in the database.
func (lc *LINENFTCreate) Save(ctx context.Context) (*LINENFT, error) {
	if err := lc.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, lc.sqlSave, lc.mutation, lc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (lc *LINENFTCreate) SaveX(ctx context.Context) *LINENFT {
	v, err := lc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (lc *LINENFTCreate) Exec(ctx context.Context) error {
	_, err := lc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lc *LINENFTCreate) ExecX(ctx context.Context) {
	if err := lc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (lc *LINENFTCreate) defaults() error {
	if _, ok := lc.mutation.CreatedAt(); !ok {
		if linenft.DefaultCreatedAt == nil {
			return fmt.Errorf("ent: uninitialized linenft.DefaultCreatedAt (forgotten import ent/runtime?)")
		}
		v := linenft.DefaultCreatedAt()
		lc.mutation.SetCreatedAt(v)
	}
	if _, ok := lc.mutation.UpdatedAt(); !ok {
		if linenft.DefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized linenft.DefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := linenft.DefaultUpdatedAt()
		lc.mutation.SetUpdatedAt(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (lc *LINENFTCreate) check() error {
	if _, ok := lc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "LINENFT.created_at"`)}
	}
	if _, ok := lc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "LINENFT.updated_at"`)}
	}
	if _, ok := lc.mutation.LineNftID(); !ok {
		return &ValidationError{Name: "line_nft_id", err: errors.New(`ent: missing required field "LINENFT.line_nft_id"`)}
	}
	if _, ok := lc.mutation.ContractID(); !ok {
		return &ValidationError{Name: "contract_id", err: errors.New(`ent: missing required field "LINENFT.contract_id"`)}
	}
	if _, ok := lc.mutation.TokenType(); !ok {
		return &ValidationError{Name: "token_type", err: errors.New(`ent: missing required field "LINENFT.token_type"`)}
	}
	if _, ok := lc.mutation.TokenIndex(); !ok {
		return &ValidationError{Name: "token_index", err: errors.New(`ent: missing required field "LINENFT.token_index"`)}
	}
	if _, ok := lc.mutation.TokenName(); !ok {
		return &ValidationError{Name: "token_name", err: errors.New(`ent: missing required field "LINENFT.token_name"`)}
	}
	if _, ok := lc.mutation.TokenDescription(); !ok {
		return &ValidationError{Name: "token_description", err: errors.New(`ent: missing required field "LINENFT.token_description"`)}
	}
	if _, ok := lc.mutation.TokenContentURL(); !ok {
		return &ValidationError{Name: "token_content_url", err: errors.New(`ent: missing required field "LINENFT.token_content_url"`)}
	}
	return nil
}

func (lc *LINENFTCreate) sqlSave(ctx context.Context) (*LINENFT, error) {
	if err := lc.check(); err != nil {
		return nil, err
	}
	_node, _spec := lc.createSpec()
	if err := sqlgraph.CreateNode(ctx, lc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = uint32(id)
	}
	lc.mutation.id = &_node.ID
	lc.mutation.done = true
	return _node, nil
}

func (lc *LINENFTCreate) createSpec() (*LINENFT, *sqlgraph.CreateSpec) {
	var (
		_node = &LINENFT{config: lc.config}
		_spec = sqlgraph.NewCreateSpec(linenft.Table, sqlgraph.NewFieldSpec(linenft.FieldID, field.TypeUint32))
	)
	if id, ok := lc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := lc.mutation.CreatedAt(); ok {
		_spec.SetField(linenft.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := lc.mutation.UpdatedAt(); ok {
		_spec.SetField(linenft.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := lc.mutation.DeletedAt(); ok {
		_spec.SetField(linenft.FieldDeletedAt, field.TypeTime, value)
		_node.DeletedAt = &value
	}
	if value, ok := lc.mutation.LineNftID(); ok {
		_spec.SetField(linenft.FieldLineNftID, field.TypeUint32, value)
		_node.LineNftID = value
	}
	if value, ok := lc.mutation.ContractID(); ok {
		_spec.SetField(linenft.FieldContractID, field.TypeString, value)
		_node.ContractID = value
	}
	if value, ok := lc.mutation.TokenType(); ok {
		_spec.SetField(linenft.FieldTokenType, field.TypeString, value)
		_node.TokenType = value
	}
	if value, ok := lc.mutation.TokenIndex(); ok {
		_spec.SetField(linenft.FieldTokenIndex, field.TypeString, value)
		_node.TokenIndex = value
	}
	if value, ok := lc.mutation.TokenName(); ok {
		_spec.SetField(linenft.FieldTokenName, field.TypeString, value)
		_node.TokenName = value
	}
	if value, ok := lc.mutation.TokenDescription(); ok {
		_spec.SetField(linenft.FieldTokenDescription, field.TypeString, value)
		_node.TokenDescription = value
	}
	if value, ok := lc.mutation.TokenContentURL(); ok {
		_spec.SetField(linenft.FieldTokenContentURL, field.TypeString, value)
		_node.TokenContentURL = value
	}
	if nodes := lc.mutation.ActivitiesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   linenft.ActivitiesTable,
			Columns: []string{linenft.ActivitiesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(linenftactivity.FieldID, field.TypeUint32),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := lc.mutation.MillionArthursPropertyIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   linenft.MillionArthursPropertyTable,
			Columns: []string{linenft.MillionArthursPropertyColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(linenftmillionarthursproperty.FieldID, field.TypeUint32),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// LINENFTCreateBulk is the builder for creating many LINENFT entities in bulk.
type LINENFTCreateBulk struct {
	config
	err      error
	builders []*LINENFTCreate
}

// Save creates the LINENFT entities in the database.
func (lcb *LINENFTCreateBulk) Save(ctx context.Context) ([]*LINENFT, error) {
	if lcb.err != nil {
		return nil, lcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(lcb.builders))
	nodes := make([]*LINENFT, len(lcb.builders))
	mutators := make([]Mutator, len(lcb.builders))
	for i := range lcb.builders {
		func(i int, root context.Context) {
			builder := lcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*LINENFTMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, lcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, lcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = uint32(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, lcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (lcb *LINENFTCreateBulk) SaveX(ctx context.Context) []*LINENFT {
	v, err := lcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (lcb *LINENFTCreateBulk) Exec(ctx context.Context) error {
	_, err := lcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lcb *LINENFTCreateBulk) ExecX(ctx context.Context) {
	if err := lcb.Exec(ctx); err != nil {
		panic(err)
	}
}
