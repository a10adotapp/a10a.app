// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/a10adotapp/a10a.app/ent/changokushiweapon"
	"github.com/a10adotapp/a10a.app/ent/changokushiweaponchangelog"
)

// ChangokushiWeaponChangeLogCreate is the builder for creating a ChangokushiWeaponChangeLog entity.
type ChangokushiWeaponChangeLogCreate struct {
	config
	mutation *ChangokushiWeaponChangeLogMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (cwclc *ChangokushiWeaponChangeLogCreate) SetCreatedAt(t time.Time) *ChangokushiWeaponChangeLogCreate {
	cwclc.mutation.SetCreatedAt(t)
	return cwclc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (cwclc *ChangokushiWeaponChangeLogCreate) SetNillableCreatedAt(t *time.Time) *ChangokushiWeaponChangeLogCreate {
	if t != nil {
		cwclc.SetCreatedAt(*t)
	}
	return cwclc
}

// SetUpdatedAt sets the "updated_at" field.
func (cwclc *ChangokushiWeaponChangeLogCreate) SetUpdatedAt(t time.Time) *ChangokushiWeaponChangeLogCreate {
	cwclc.mutation.SetUpdatedAt(t)
	return cwclc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (cwclc *ChangokushiWeaponChangeLogCreate) SetNillableUpdatedAt(t *time.Time) *ChangokushiWeaponChangeLogCreate {
	if t != nil {
		cwclc.SetUpdatedAt(*t)
	}
	return cwclc
}

// SetDeletedAt sets the "deleted_at" field.
func (cwclc *ChangokushiWeaponChangeLogCreate) SetDeletedAt(t time.Time) *ChangokushiWeaponChangeLogCreate {
	cwclc.mutation.SetDeletedAt(t)
	return cwclc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (cwclc *ChangokushiWeaponChangeLogCreate) SetNillableDeletedAt(t *time.Time) *ChangokushiWeaponChangeLogCreate {
	if t != nil {
		cwclc.SetDeletedAt(*t)
	}
	return cwclc
}

// SetChangokushiWeaponID sets the "changokushi_weapon_id" field.
func (cwclc *ChangokushiWeaponChangeLogCreate) SetChangokushiWeaponID(u uint32) *ChangokushiWeaponChangeLogCreate {
	cwclc.mutation.SetChangokushiWeaponID(u)
	return cwclc
}

// SetStatus sets the "status" field.
func (cwclc *ChangokushiWeaponChangeLogCreate) SetStatus(s string) *ChangokushiWeaponChangeLogCreate {
	cwclc.mutation.SetStatus(s)
	return cwclc
}

// SetPrice sets the "price" field.
func (cwclc *ChangokushiWeaponChangeLogCreate) SetPrice(i int) *ChangokushiWeaponChangeLogCreate {
	cwclc.mutation.SetPrice(i)
	return cwclc
}

// SetPublishedAt sets the "published_at" field.
func (cwclc *ChangokushiWeaponChangeLogCreate) SetPublishedAt(t time.Time) *ChangokushiWeaponChangeLogCreate {
	cwclc.mutation.SetPublishedAt(t)
	return cwclc
}

// SetID sets the "id" field.
func (cwclc *ChangokushiWeaponChangeLogCreate) SetID(u uint32) *ChangokushiWeaponChangeLogCreate {
	cwclc.mutation.SetID(u)
	return cwclc
}

// SetChangokushiWeapon sets the "changokushi_weapon" edge to the ChangokushiWeapon entity.
func (cwclc *ChangokushiWeaponChangeLogCreate) SetChangokushiWeapon(c *ChangokushiWeapon) *ChangokushiWeaponChangeLogCreate {
	return cwclc.SetChangokushiWeaponID(c.ID)
}

// Mutation returns the ChangokushiWeaponChangeLogMutation object of the builder.
func (cwclc *ChangokushiWeaponChangeLogCreate) Mutation() *ChangokushiWeaponChangeLogMutation {
	return cwclc.mutation
}

// Save creates the ChangokushiWeaponChangeLog in the database.
func (cwclc *ChangokushiWeaponChangeLogCreate) Save(ctx context.Context) (*ChangokushiWeaponChangeLog, error) {
	if err := cwclc.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, cwclc.sqlSave, cwclc.mutation, cwclc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (cwclc *ChangokushiWeaponChangeLogCreate) SaveX(ctx context.Context) *ChangokushiWeaponChangeLog {
	v, err := cwclc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cwclc *ChangokushiWeaponChangeLogCreate) Exec(ctx context.Context) error {
	_, err := cwclc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cwclc *ChangokushiWeaponChangeLogCreate) ExecX(ctx context.Context) {
	if err := cwclc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cwclc *ChangokushiWeaponChangeLogCreate) defaults() error {
	if _, ok := cwclc.mutation.CreatedAt(); !ok {
		if changokushiweaponchangelog.DefaultCreatedAt == nil {
			return fmt.Errorf("ent: uninitialized changokushiweaponchangelog.DefaultCreatedAt (forgotten import ent/runtime?)")
		}
		v := changokushiweaponchangelog.DefaultCreatedAt()
		cwclc.mutation.SetCreatedAt(v)
	}
	if _, ok := cwclc.mutation.UpdatedAt(); !ok {
		if changokushiweaponchangelog.DefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized changokushiweaponchangelog.DefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := changokushiweaponchangelog.DefaultUpdatedAt()
		cwclc.mutation.SetUpdatedAt(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (cwclc *ChangokushiWeaponChangeLogCreate) check() error {
	if _, ok := cwclc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "ChangokushiWeaponChangeLog.created_at"`)}
	}
	if _, ok := cwclc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "ChangokushiWeaponChangeLog.updated_at"`)}
	}
	if _, ok := cwclc.mutation.ChangokushiWeaponID(); !ok {
		return &ValidationError{Name: "changokushi_weapon_id", err: errors.New(`ent: missing required field "ChangokushiWeaponChangeLog.changokushi_weapon_id"`)}
	}
	if _, ok := cwclc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`ent: missing required field "ChangokushiWeaponChangeLog.status"`)}
	}
	if _, ok := cwclc.mutation.Price(); !ok {
		return &ValidationError{Name: "price", err: errors.New(`ent: missing required field "ChangokushiWeaponChangeLog.price"`)}
	}
	if _, ok := cwclc.mutation.PublishedAt(); !ok {
		return &ValidationError{Name: "published_at", err: errors.New(`ent: missing required field "ChangokushiWeaponChangeLog.published_at"`)}
	}
	if _, ok := cwclc.mutation.ChangokushiWeaponID(); !ok {
		return &ValidationError{Name: "changokushi_weapon", err: errors.New(`ent: missing required edge "ChangokushiWeaponChangeLog.changokushi_weapon"`)}
	}
	return nil
}

func (cwclc *ChangokushiWeaponChangeLogCreate) sqlSave(ctx context.Context) (*ChangokushiWeaponChangeLog, error) {
	if err := cwclc.check(); err != nil {
		return nil, err
	}
	_node, _spec := cwclc.createSpec()
	if err := sqlgraph.CreateNode(ctx, cwclc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = uint32(id)
	}
	cwclc.mutation.id = &_node.ID
	cwclc.mutation.done = true
	return _node, nil
}

func (cwclc *ChangokushiWeaponChangeLogCreate) createSpec() (*ChangokushiWeaponChangeLog, *sqlgraph.CreateSpec) {
	var (
		_node = &ChangokushiWeaponChangeLog{config: cwclc.config}
		_spec = sqlgraph.NewCreateSpec(changokushiweaponchangelog.Table, sqlgraph.NewFieldSpec(changokushiweaponchangelog.FieldID, field.TypeUint32))
	)
	if id, ok := cwclc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := cwclc.mutation.CreatedAt(); ok {
		_spec.SetField(changokushiweaponchangelog.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := cwclc.mutation.UpdatedAt(); ok {
		_spec.SetField(changokushiweaponchangelog.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := cwclc.mutation.DeletedAt(); ok {
		_spec.SetField(changokushiweaponchangelog.FieldDeletedAt, field.TypeTime, value)
		_node.DeletedAt = &value
	}
	if value, ok := cwclc.mutation.Status(); ok {
		_spec.SetField(changokushiweaponchangelog.FieldStatus, field.TypeString, value)
		_node.Status = value
	}
	if value, ok := cwclc.mutation.Price(); ok {
		_spec.SetField(changokushiweaponchangelog.FieldPrice, field.TypeInt, value)
		_node.Price = value
	}
	if value, ok := cwclc.mutation.PublishedAt(); ok {
		_spec.SetField(changokushiweaponchangelog.FieldPublishedAt, field.TypeTime, value)
		_node.PublishedAt = value
	}
	if nodes := cwclc.mutation.ChangokushiWeaponIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   changokushiweaponchangelog.ChangokushiWeaponTable,
			Columns: []string{changokushiweaponchangelog.ChangokushiWeaponColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(changokushiweapon.FieldID, field.TypeUint32),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.ChangokushiWeaponID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// ChangokushiWeaponChangeLogCreateBulk is the builder for creating many ChangokushiWeaponChangeLog entities in bulk.
type ChangokushiWeaponChangeLogCreateBulk struct {
	config
	err      error
	builders []*ChangokushiWeaponChangeLogCreate
}

// Save creates the ChangokushiWeaponChangeLog entities in the database.
func (cwclcb *ChangokushiWeaponChangeLogCreateBulk) Save(ctx context.Context) ([]*ChangokushiWeaponChangeLog, error) {
	if cwclcb.err != nil {
		return nil, cwclcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(cwclcb.builders))
	nodes := make([]*ChangokushiWeaponChangeLog, len(cwclcb.builders))
	mutators := make([]Mutator, len(cwclcb.builders))
	for i := range cwclcb.builders {
		func(i int, root context.Context) {
			builder := cwclcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ChangokushiWeaponChangeLogMutation)
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
					_, err = mutators[i+1].Mutate(root, cwclcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, cwclcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, cwclcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (cwclcb *ChangokushiWeaponChangeLogCreateBulk) SaveX(ctx context.Context) []*ChangokushiWeaponChangeLog {
	v, err := cwclcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cwclcb *ChangokushiWeaponChangeLogCreateBulk) Exec(ctx context.Context) error {
	_, err := cwclcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cwclcb *ChangokushiWeaponChangeLogCreateBulk) ExecX(ctx context.Context) {
	if err := cwclcb.Exec(ctx); err != nil {
		panic(err)
	}
}