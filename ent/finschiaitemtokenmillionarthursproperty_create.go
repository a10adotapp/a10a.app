// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/a10adotapp/a10a.app/ent/finschiaitemtoken"
	"github.com/a10adotapp/a10a.app/ent/finschiaitemtokenmillionarthursproperty"
)

// FinschiaItemTokenMillionArthursPropertyCreate is the builder for creating a FinschiaItemTokenMillionArthursProperty entity.
type FinschiaItemTokenMillionArthursPropertyCreate struct {
	config
	mutation *FinschiaItemTokenMillionArthursPropertyMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (fitmapc *FinschiaItemTokenMillionArthursPropertyCreate) SetCreatedAt(t time.Time) *FinschiaItemTokenMillionArthursPropertyCreate {
	fitmapc.mutation.SetCreatedAt(t)
	return fitmapc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (fitmapc *FinschiaItemTokenMillionArthursPropertyCreate) SetNillableCreatedAt(t *time.Time) *FinschiaItemTokenMillionArthursPropertyCreate {
	if t != nil {
		fitmapc.SetCreatedAt(*t)
	}
	return fitmapc
}

// SetUpdatedAt sets the "updated_at" field.
func (fitmapc *FinschiaItemTokenMillionArthursPropertyCreate) SetUpdatedAt(t time.Time) *FinschiaItemTokenMillionArthursPropertyCreate {
	fitmapc.mutation.SetUpdatedAt(t)
	return fitmapc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (fitmapc *FinschiaItemTokenMillionArthursPropertyCreate) SetNillableUpdatedAt(t *time.Time) *FinschiaItemTokenMillionArthursPropertyCreate {
	if t != nil {
		fitmapc.SetUpdatedAt(*t)
	}
	return fitmapc
}

// SetDeletedAt sets the "deleted_at" field.
func (fitmapc *FinschiaItemTokenMillionArthursPropertyCreate) SetDeletedAt(t time.Time) *FinschiaItemTokenMillionArthursPropertyCreate {
	fitmapc.mutation.SetDeletedAt(t)
	return fitmapc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (fitmapc *FinschiaItemTokenMillionArthursPropertyCreate) SetNillableDeletedAt(t *time.Time) *FinschiaItemTokenMillionArthursPropertyCreate {
	if t != nil {
		fitmapc.SetDeletedAt(*t)
	}
	return fitmapc
}

// SetFinschiaItemTokenID sets the "finschia_item_token_id" field.
func (fitmapc *FinschiaItemTokenMillionArthursPropertyCreate) SetFinschiaItemTokenID(u uint32) *FinschiaItemTokenMillionArthursPropertyCreate {
	fitmapc.mutation.SetFinschiaItemTokenID(u)
	return fitmapc
}

// SetSeries sets the "series" field.
func (fitmapc *FinschiaItemTokenMillionArthursPropertyCreate) SetSeries(s string) *FinschiaItemTokenMillionArthursPropertyCreate {
	fitmapc.mutation.SetSeries(s)
	return fitmapc
}

// SetNillableSeries sets the "series" field if the given value is not nil.
func (fitmapc *FinschiaItemTokenMillionArthursPropertyCreate) SetNillableSeries(s *string) *FinschiaItemTokenMillionArthursPropertyCreate {
	if s != nil {
		fitmapc.SetSeries(*s)
	}
	return fitmapc
}

// SetGearCategory sets the "gear_category" field.
func (fitmapc *FinschiaItemTokenMillionArthursPropertyCreate) SetGearCategory(s string) *FinschiaItemTokenMillionArthursPropertyCreate {
	fitmapc.mutation.SetGearCategory(s)
	return fitmapc
}

// SetNillableGearCategory sets the "gear_category" field if the given value is not nil.
func (fitmapc *FinschiaItemTokenMillionArthursPropertyCreate) SetNillableGearCategory(s *string) *FinschiaItemTokenMillionArthursPropertyCreate {
	if s != nil {
		fitmapc.SetGearCategory(*s)
	}
	return fitmapc
}

// SetGearRarity sets the "gear_rarity" field.
func (fitmapc *FinschiaItemTokenMillionArthursPropertyCreate) SetGearRarity(s string) *FinschiaItemTokenMillionArthursPropertyCreate {
	fitmapc.mutation.SetGearRarity(s)
	return fitmapc
}

// SetNillableGearRarity sets the "gear_rarity" field if the given value is not nil.
func (fitmapc *FinschiaItemTokenMillionArthursPropertyCreate) SetNillableGearRarity(s *string) *FinschiaItemTokenMillionArthursPropertyCreate {
	if s != nil {
		fitmapc.SetGearRarity(*s)
	}
	return fitmapc
}

// SetID sets the "id" field.
func (fitmapc *FinschiaItemTokenMillionArthursPropertyCreate) SetID(u uint32) *FinschiaItemTokenMillionArthursPropertyCreate {
	fitmapc.mutation.SetID(u)
	return fitmapc
}

// SetFinschiaItemToken sets the "finschia_item_token" edge to the FinschiaItemToken entity.
func (fitmapc *FinschiaItemTokenMillionArthursPropertyCreate) SetFinschiaItemToken(f *FinschiaItemToken) *FinschiaItemTokenMillionArthursPropertyCreate {
	return fitmapc.SetFinschiaItemTokenID(f.ID)
}

// Mutation returns the FinschiaItemTokenMillionArthursPropertyMutation object of the builder.
func (fitmapc *FinschiaItemTokenMillionArthursPropertyCreate) Mutation() *FinschiaItemTokenMillionArthursPropertyMutation {
	return fitmapc.mutation
}

// Save creates the FinschiaItemTokenMillionArthursProperty in the database.
func (fitmapc *FinschiaItemTokenMillionArthursPropertyCreate) Save(ctx context.Context) (*FinschiaItemTokenMillionArthursProperty, error) {
	if err := fitmapc.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, fitmapc.sqlSave, fitmapc.mutation, fitmapc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (fitmapc *FinschiaItemTokenMillionArthursPropertyCreate) SaveX(ctx context.Context) *FinschiaItemTokenMillionArthursProperty {
	v, err := fitmapc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (fitmapc *FinschiaItemTokenMillionArthursPropertyCreate) Exec(ctx context.Context) error {
	_, err := fitmapc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fitmapc *FinschiaItemTokenMillionArthursPropertyCreate) ExecX(ctx context.Context) {
	if err := fitmapc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (fitmapc *FinschiaItemTokenMillionArthursPropertyCreate) defaults() error {
	if _, ok := fitmapc.mutation.CreatedAt(); !ok {
		if finschiaitemtokenmillionarthursproperty.DefaultCreatedAt == nil {
			return fmt.Errorf("ent: uninitialized finschiaitemtokenmillionarthursproperty.DefaultCreatedAt (forgotten import ent/runtime?)")
		}
		v := finschiaitemtokenmillionarthursproperty.DefaultCreatedAt()
		fitmapc.mutation.SetCreatedAt(v)
	}
	if _, ok := fitmapc.mutation.UpdatedAt(); !ok {
		if finschiaitemtokenmillionarthursproperty.DefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized finschiaitemtokenmillionarthursproperty.DefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := finschiaitemtokenmillionarthursproperty.DefaultUpdatedAt()
		fitmapc.mutation.SetUpdatedAt(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (fitmapc *FinschiaItemTokenMillionArthursPropertyCreate) check() error {
	if _, ok := fitmapc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "FinschiaItemTokenMillionArthursProperty.created_at"`)}
	}
	if _, ok := fitmapc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "FinschiaItemTokenMillionArthursProperty.updated_at"`)}
	}
	if _, ok := fitmapc.mutation.FinschiaItemTokenID(); !ok {
		return &ValidationError{Name: "finschia_item_token_id", err: errors.New(`ent: missing required field "FinschiaItemTokenMillionArthursProperty.finschia_item_token_id"`)}
	}
	if _, ok := fitmapc.mutation.FinschiaItemTokenID(); !ok {
		return &ValidationError{Name: "finschia_item_token", err: errors.New(`ent: missing required edge "FinschiaItemTokenMillionArthursProperty.finschia_item_token"`)}
	}
	return nil
}

func (fitmapc *FinschiaItemTokenMillionArthursPropertyCreate) sqlSave(ctx context.Context) (*FinschiaItemTokenMillionArthursProperty, error) {
	if err := fitmapc.check(); err != nil {
		return nil, err
	}
	_node, _spec := fitmapc.createSpec()
	if err := sqlgraph.CreateNode(ctx, fitmapc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = uint32(id)
	}
	fitmapc.mutation.id = &_node.ID
	fitmapc.mutation.done = true
	return _node, nil
}

func (fitmapc *FinschiaItemTokenMillionArthursPropertyCreate) createSpec() (*FinschiaItemTokenMillionArthursProperty, *sqlgraph.CreateSpec) {
	var (
		_node = &FinschiaItemTokenMillionArthursProperty{config: fitmapc.config}
		_spec = sqlgraph.NewCreateSpec(finschiaitemtokenmillionarthursproperty.Table, sqlgraph.NewFieldSpec(finschiaitemtokenmillionarthursproperty.FieldID, field.TypeUint32))
	)
	if id, ok := fitmapc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := fitmapc.mutation.CreatedAt(); ok {
		_spec.SetField(finschiaitemtokenmillionarthursproperty.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := fitmapc.mutation.UpdatedAt(); ok {
		_spec.SetField(finschiaitemtokenmillionarthursproperty.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := fitmapc.mutation.DeletedAt(); ok {
		_spec.SetField(finschiaitemtokenmillionarthursproperty.FieldDeletedAt, field.TypeTime, value)
		_node.DeletedAt = &value
	}
	if value, ok := fitmapc.mutation.Series(); ok {
		_spec.SetField(finschiaitemtokenmillionarthursproperty.FieldSeries, field.TypeString, value)
		_node.Series = &value
	}
	if value, ok := fitmapc.mutation.GearCategory(); ok {
		_spec.SetField(finschiaitemtokenmillionarthursproperty.FieldGearCategory, field.TypeString, value)
		_node.GearCategory = &value
	}
	if value, ok := fitmapc.mutation.GearRarity(); ok {
		_spec.SetField(finschiaitemtokenmillionarthursproperty.FieldGearRarity, field.TypeString, value)
		_node.GearRarity = &value
	}
	if nodes := fitmapc.mutation.FinschiaItemTokenIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   finschiaitemtokenmillionarthursproperty.FinschiaItemTokenTable,
			Columns: []string{finschiaitemtokenmillionarthursproperty.FinschiaItemTokenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(finschiaitemtoken.FieldID, field.TypeUint32),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.FinschiaItemTokenID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// FinschiaItemTokenMillionArthursPropertyCreateBulk is the builder for creating many FinschiaItemTokenMillionArthursProperty entities in bulk.
type FinschiaItemTokenMillionArthursPropertyCreateBulk struct {
	config
	err      error
	builders []*FinschiaItemTokenMillionArthursPropertyCreate
}

// Save creates the FinschiaItemTokenMillionArthursProperty entities in the database.
func (fitmapcb *FinschiaItemTokenMillionArthursPropertyCreateBulk) Save(ctx context.Context) ([]*FinschiaItemTokenMillionArthursProperty, error) {
	if fitmapcb.err != nil {
		return nil, fitmapcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(fitmapcb.builders))
	nodes := make([]*FinschiaItemTokenMillionArthursProperty, len(fitmapcb.builders))
	mutators := make([]Mutator, len(fitmapcb.builders))
	for i := range fitmapcb.builders {
		func(i int, root context.Context) {
			builder := fitmapcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*FinschiaItemTokenMillionArthursPropertyMutation)
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
					_, err = mutators[i+1].Mutate(root, fitmapcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, fitmapcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, fitmapcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (fitmapcb *FinschiaItemTokenMillionArthursPropertyCreateBulk) SaveX(ctx context.Context) []*FinschiaItemTokenMillionArthursProperty {
	v, err := fitmapcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (fitmapcb *FinschiaItemTokenMillionArthursPropertyCreateBulk) Exec(ctx context.Context) error {
	_, err := fitmapcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fitmapcb *FinschiaItemTokenMillionArthursPropertyCreateBulk) ExecX(ctx context.Context) {
	if err := fitmapcb.Exec(ctx); err != nil {
		panic(err)
	}
}