// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/a10adotapp/a10a.app/ent/kusogeeeeeenft"
	"github.com/a10adotapp/a10a.app/ent/kusogeeeeeenftchangelog"
)

// KusogeeeeeeNFTChangeLogCreate is the builder for creating a KusogeeeeeeNFTChangeLog entity.
type KusogeeeeeeNFTChangeLogCreate struct {
	config
	mutation *KusogeeeeeeNFTChangeLogMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (knclc *KusogeeeeeeNFTChangeLogCreate) SetCreatedAt(t time.Time) *KusogeeeeeeNFTChangeLogCreate {
	knclc.mutation.SetCreatedAt(t)
	return knclc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (knclc *KusogeeeeeeNFTChangeLogCreate) SetNillableCreatedAt(t *time.Time) *KusogeeeeeeNFTChangeLogCreate {
	if t != nil {
		knclc.SetCreatedAt(*t)
	}
	return knclc
}

// SetUpdatedAt sets the "updated_at" field.
func (knclc *KusogeeeeeeNFTChangeLogCreate) SetUpdatedAt(t time.Time) *KusogeeeeeeNFTChangeLogCreate {
	knclc.mutation.SetUpdatedAt(t)
	return knclc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (knclc *KusogeeeeeeNFTChangeLogCreate) SetNillableUpdatedAt(t *time.Time) *KusogeeeeeeNFTChangeLogCreate {
	if t != nil {
		knclc.SetUpdatedAt(*t)
	}
	return knclc
}

// SetDeletedAt sets the "deleted_at" field.
func (knclc *KusogeeeeeeNFTChangeLogCreate) SetDeletedAt(t time.Time) *KusogeeeeeeNFTChangeLogCreate {
	knclc.mutation.SetDeletedAt(t)
	return knclc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (knclc *KusogeeeeeeNFTChangeLogCreate) SetNillableDeletedAt(t *time.Time) *KusogeeeeeeNFTChangeLogCreate {
	if t != nil {
		knclc.SetDeletedAt(*t)
	}
	return knclc
}

// SetKusogeeeeeeNftID sets the "kusogeeeeee_nft_id" field.
func (knclc *KusogeeeeeeNFTChangeLogCreate) SetKusogeeeeeeNftID(u uint32) *KusogeeeeeeNFTChangeLogCreate {
	knclc.mutation.SetKusogeeeeeeNftID(u)
	return knclc
}

// SetStatus sets the "status" field.
func (knclc *KusogeeeeeeNFTChangeLogCreate) SetStatus(s string) *KusogeeeeeeNFTChangeLogCreate {
	knclc.mutation.SetStatus(s)
	return knclc
}

// SetPrice sets the "price" field.
func (knclc *KusogeeeeeeNFTChangeLogCreate) SetPrice(i int) *KusogeeeeeeNFTChangeLogCreate {
	knclc.mutation.SetPrice(i)
	return knclc
}

// SetID sets the "id" field.
func (knclc *KusogeeeeeeNFTChangeLogCreate) SetID(u uint32) *KusogeeeeeeNFTChangeLogCreate {
	knclc.mutation.SetID(u)
	return knclc
}

// SetKusogeeeeeeNft sets the "kusogeeeeee_nft" edge to the KusogeeeeeeNFT entity.
func (knclc *KusogeeeeeeNFTChangeLogCreate) SetKusogeeeeeeNft(k *KusogeeeeeeNFT) *KusogeeeeeeNFTChangeLogCreate {
	return knclc.SetKusogeeeeeeNftID(k.ID)
}

// Mutation returns the KusogeeeeeeNFTChangeLogMutation object of the builder.
func (knclc *KusogeeeeeeNFTChangeLogCreate) Mutation() *KusogeeeeeeNFTChangeLogMutation {
	return knclc.mutation
}

// Save creates the KusogeeeeeeNFTChangeLog in the database.
func (knclc *KusogeeeeeeNFTChangeLogCreate) Save(ctx context.Context) (*KusogeeeeeeNFTChangeLog, error) {
	if err := knclc.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, knclc.sqlSave, knclc.mutation, knclc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (knclc *KusogeeeeeeNFTChangeLogCreate) SaveX(ctx context.Context) *KusogeeeeeeNFTChangeLog {
	v, err := knclc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (knclc *KusogeeeeeeNFTChangeLogCreate) Exec(ctx context.Context) error {
	_, err := knclc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (knclc *KusogeeeeeeNFTChangeLogCreate) ExecX(ctx context.Context) {
	if err := knclc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (knclc *KusogeeeeeeNFTChangeLogCreate) defaults() error {
	if _, ok := knclc.mutation.CreatedAt(); !ok {
		if kusogeeeeeenftchangelog.DefaultCreatedAt == nil {
			return fmt.Errorf("ent: uninitialized kusogeeeeeenftchangelog.DefaultCreatedAt (forgotten import ent/runtime?)")
		}
		v := kusogeeeeeenftchangelog.DefaultCreatedAt()
		knclc.mutation.SetCreatedAt(v)
	}
	if _, ok := knclc.mutation.UpdatedAt(); !ok {
		if kusogeeeeeenftchangelog.DefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized kusogeeeeeenftchangelog.DefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := kusogeeeeeenftchangelog.DefaultUpdatedAt()
		knclc.mutation.SetUpdatedAt(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (knclc *KusogeeeeeeNFTChangeLogCreate) check() error {
	if _, ok := knclc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "KusogeeeeeeNFTChangeLog.created_at"`)}
	}
	if _, ok := knclc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "KusogeeeeeeNFTChangeLog.updated_at"`)}
	}
	if _, ok := knclc.mutation.KusogeeeeeeNftID(); !ok {
		return &ValidationError{Name: "kusogeeeeee_nft_id", err: errors.New(`ent: missing required field "KusogeeeeeeNFTChangeLog.kusogeeeeee_nft_id"`)}
	}
	if _, ok := knclc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`ent: missing required field "KusogeeeeeeNFTChangeLog.status"`)}
	}
	if _, ok := knclc.mutation.Price(); !ok {
		return &ValidationError{Name: "price", err: errors.New(`ent: missing required field "KusogeeeeeeNFTChangeLog.price"`)}
	}
	if _, ok := knclc.mutation.KusogeeeeeeNftID(); !ok {
		return &ValidationError{Name: "kusogeeeeee_nft", err: errors.New(`ent: missing required edge "KusogeeeeeeNFTChangeLog.kusogeeeeee_nft"`)}
	}
	return nil
}

func (knclc *KusogeeeeeeNFTChangeLogCreate) sqlSave(ctx context.Context) (*KusogeeeeeeNFTChangeLog, error) {
	if err := knclc.check(); err != nil {
		return nil, err
	}
	_node, _spec := knclc.createSpec()
	if err := sqlgraph.CreateNode(ctx, knclc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = uint32(id)
	}
	knclc.mutation.id = &_node.ID
	knclc.mutation.done = true
	return _node, nil
}

func (knclc *KusogeeeeeeNFTChangeLogCreate) createSpec() (*KusogeeeeeeNFTChangeLog, *sqlgraph.CreateSpec) {
	var (
		_node = &KusogeeeeeeNFTChangeLog{config: knclc.config}
		_spec = sqlgraph.NewCreateSpec(kusogeeeeeenftchangelog.Table, sqlgraph.NewFieldSpec(kusogeeeeeenftchangelog.FieldID, field.TypeUint32))
	)
	if id, ok := knclc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := knclc.mutation.CreatedAt(); ok {
		_spec.SetField(kusogeeeeeenftchangelog.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := knclc.mutation.UpdatedAt(); ok {
		_spec.SetField(kusogeeeeeenftchangelog.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := knclc.mutation.DeletedAt(); ok {
		_spec.SetField(kusogeeeeeenftchangelog.FieldDeletedAt, field.TypeTime, value)
		_node.DeletedAt = &value
	}
	if value, ok := knclc.mutation.Status(); ok {
		_spec.SetField(kusogeeeeeenftchangelog.FieldStatus, field.TypeString, value)
		_node.Status = value
	}
	if value, ok := knclc.mutation.Price(); ok {
		_spec.SetField(kusogeeeeeenftchangelog.FieldPrice, field.TypeInt, value)
		_node.Price = value
	}
	if nodes := knclc.mutation.KusogeeeeeeNftIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   kusogeeeeeenftchangelog.KusogeeeeeeNftTable,
			Columns: []string{kusogeeeeeenftchangelog.KusogeeeeeeNftColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(kusogeeeeeenft.FieldID, field.TypeUint32),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.KusogeeeeeeNftID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// KusogeeeeeeNFTChangeLogCreateBulk is the builder for creating many KusogeeeeeeNFTChangeLog entities in bulk.
type KusogeeeeeeNFTChangeLogCreateBulk struct {
	config
	err      error
	builders []*KusogeeeeeeNFTChangeLogCreate
}

// Save creates the KusogeeeeeeNFTChangeLog entities in the database.
func (knclcb *KusogeeeeeeNFTChangeLogCreateBulk) Save(ctx context.Context) ([]*KusogeeeeeeNFTChangeLog, error) {
	if knclcb.err != nil {
		return nil, knclcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(knclcb.builders))
	nodes := make([]*KusogeeeeeeNFTChangeLog, len(knclcb.builders))
	mutators := make([]Mutator, len(knclcb.builders))
	for i := range knclcb.builders {
		func(i int, root context.Context) {
			builder := knclcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*KusogeeeeeeNFTChangeLogMutation)
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
					_, err = mutators[i+1].Mutate(root, knclcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, knclcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, knclcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (knclcb *KusogeeeeeeNFTChangeLogCreateBulk) SaveX(ctx context.Context) []*KusogeeeeeeNFTChangeLog {
	v, err := knclcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (knclcb *KusogeeeeeeNFTChangeLogCreateBulk) Exec(ctx context.Context) error {
	_, err := knclcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (knclcb *KusogeeeeeeNFTChangeLogCreateBulk) ExecX(ctx context.Context) {
	if err := knclcb.Exec(ctx); err != nil {
		panic(err)
	}
}
