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

// KusogeeeeeeNFTCreate is the builder for creating a KusogeeeeeeNFT entity.
type KusogeeeeeeNFTCreate struct {
	config
	mutation *KusogeeeeeeNFTMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (knc *KusogeeeeeeNFTCreate) SetCreatedAt(t time.Time) *KusogeeeeeeNFTCreate {
	knc.mutation.SetCreatedAt(t)
	return knc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (knc *KusogeeeeeeNFTCreate) SetNillableCreatedAt(t *time.Time) *KusogeeeeeeNFTCreate {
	if t != nil {
		knc.SetCreatedAt(*t)
	}
	return knc
}

// SetUpdatedAt sets the "updated_at" field.
func (knc *KusogeeeeeeNFTCreate) SetUpdatedAt(t time.Time) *KusogeeeeeeNFTCreate {
	knc.mutation.SetUpdatedAt(t)
	return knc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (knc *KusogeeeeeeNFTCreate) SetNillableUpdatedAt(t *time.Time) *KusogeeeeeeNFTCreate {
	if t != nil {
		knc.SetUpdatedAt(*t)
	}
	return knc
}

// SetDeletedAt sets the "deleted_at" field.
func (knc *KusogeeeeeeNFTCreate) SetDeletedAt(t time.Time) *KusogeeeeeeNFTCreate {
	knc.mutation.SetDeletedAt(t)
	return knc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (knc *KusogeeeeeeNFTCreate) SetNillableDeletedAt(t *time.Time) *KusogeeeeeeNFTCreate {
	if t != nil {
		knc.SetDeletedAt(*t)
	}
	return knc
}

// SetURI sets the "uri" field.
func (knc *KusogeeeeeeNFTCreate) SetURI(s string) *KusogeeeeeeNFTCreate {
	knc.mutation.SetURI(s)
	return knc
}

// SetType sets the "type" field.
func (knc *KusogeeeeeeNFTCreate) SetType(s string) *KusogeeeeeeNFTCreate {
	knc.mutation.SetType(s)
	return knc
}

// SetName sets the "name" field.
func (knc *KusogeeeeeeNFTCreate) SetName(s string) *KusogeeeeeeNFTCreate {
	knc.mutation.SetName(s)
	return knc
}

// SetStatus sets the "status" field.
func (knc *KusogeeeeeeNFTCreate) SetStatus(s string) *KusogeeeeeeNFTCreate {
	knc.mutation.SetStatus(s)
	return knc
}

// SetPrice sets the "price" field.
func (knc *KusogeeeeeeNFTCreate) SetPrice(i int) *KusogeeeeeeNFTCreate {
	knc.mutation.SetPrice(i)
	return knc
}

// SetPublishedAt sets the "published_at" field.
func (knc *KusogeeeeeeNFTCreate) SetPublishedAt(t time.Time) *KusogeeeeeeNFTCreate {
	knc.mutation.SetPublishedAt(t)
	return knc
}

// SetWeaponRank sets the "weapon_rank" field.
func (knc *KusogeeeeeeNFTCreate) SetWeaponRank(i int) *KusogeeeeeeNFTCreate {
	knc.mutation.SetWeaponRank(i)
	return knc
}

// SetNillableWeaponRank sets the "weapon_rank" field if the given value is not nil.
func (knc *KusogeeeeeeNFTCreate) SetNillableWeaponRank(i *int) *KusogeeeeeeNFTCreate {
	if i != nil {
		knc.SetWeaponRank(*i)
	}
	return knc
}

// SetWeaponType sets the "weapon_type" field.
func (knc *KusogeeeeeeNFTCreate) SetWeaponType(s string) *KusogeeeeeeNFTCreate {
	knc.mutation.SetWeaponType(s)
	return knc
}

// SetNillableWeaponType sets the "weapon_type" field if the given value is not nil.
func (knc *KusogeeeeeeNFTCreate) SetNillableWeaponType(s *string) *KusogeeeeeeNFTCreate {
	if s != nil {
		knc.SetWeaponType(*s)
	}
	return knc
}

// SetWeaponVitality sets the "weapon_vitality" field.
func (knc *KusogeeeeeeNFTCreate) SetWeaponVitality(i int) *KusogeeeeeeNFTCreate {
	knc.mutation.SetWeaponVitality(i)
	return knc
}

// SetNillableWeaponVitality sets the "weapon_vitality" field if the given value is not nil.
func (knc *KusogeeeeeeNFTCreate) SetNillableWeaponVitality(i *int) *KusogeeeeeeNFTCreate {
	if i != nil {
		knc.SetWeaponVitality(*i)
	}
	return knc
}

// SetWeaponStrength sets the "weapon_strength" field.
func (knc *KusogeeeeeeNFTCreate) SetWeaponStrength(i int) *KusogeeeeeeNFTCreate {
	knc.mutation.SetWeaponStrength(i)
	return knc
}

// SetNillableWeaponStrength sets the "weapon_strength" field if the given value is not nil.
func (knc *KusogeeeeeeNFTCreate) SetNillableWeaponStrength(i *int) *KusogeeeeeeNFTCreate {
	if i != nil {
		knc.SetWeaponStrength(*i)
	}
	return knc
}

// SetWeaponPhysicalDefense sets the "weapon_physical_defense" field.
func (knc *KusogeeeeeeNFTCreate) SetWeaponPhysicalDefense(i int) *KusogeeeeeeNFTCreate {
	knc.mutation.SetWeaponPhysicalDefense(i)
	return knc
}

// SetNillableWeaponPhysicalDefense sets the "weapon_physical_defense" field if the given value is not nil.
func (knc *KusogeeeeeeNFTCreate) SetNillableWeaponPhysicalDefense(i *int) *KusogeeeeeeNFTCreate {
	if i != nil {
		knc.SetWeaponPhysicalDefense(*i)
	}
	return knc
}

// SetWeaponMagicalDefense sets the "weapon_magical_defense" field.
func (knc *KusogeeeeeeNFTCreate) SetWeaponMagicalDefense(i int) *KusogeeeeeeNFTCreate {
	knc.mutation.SetWeaponMagicalDefense(i)
	return knc
}

// SetNillableWeaponMagicalDefense sets the "weapon_magical_defense" field if the given value is not nil.
func (knc *KusogeeeeeeNFTCreate) SetNillableWeaponMagicalDefense(i *int) *KusogeeeeeeNFTCreate {
	if i != nil {
		knc.SetWeaponMagicalDefense(*i)
	}
	return knc
}

// SetWeaponAgility sets the "weapon_agility" field.
func (knc *KusogeeeeeeNFTCreate) SetWeaponAgility(i int) *KusogeeeeeeNFTCreate {
	knc.mutation.SetWeaponAgility(i)
	return knc
}

// SetNillableWeaponAgility sets the "weapon_agility" field if the given value is not nil.
func (knc *KusogeeeeeeNFTCreate) SetNillableWeaponAgility(i *int) *KusogeeeeeeNFTCreate {
	if i != nil {
		knc.SetWeaponAgility(*i)
	}
	return knc
}

// SetCharacterRank sets the "character_rank" field.
func (knc *KusogeeeeeeNFTCreate) SetCharacterRank(s string) *KusogeeeeeeNFTCreate {
	knc.mutation.SetCharacterRank(s)
	return knc
}

// SetNillableCharacterRank sets the "character_rank" field if the given value is not nil.
func (knc *KusogeeeeeeNFTCreate) SetNillableCharacterRank(s *string) *KusogeeeeeeNFTCreate {
	if s != nil {
		knc.SetCharacterRank(*s)
	}
	return knc
}

// SetCharacterTotalSupply sets the "character_total_supply" field.
func (knc *KusogeeeeeeNFTCreate) SetCharacterTotalSupply(i int) *KusogeeeeeeNFTCreate {
	knc.mutation.SetCharacterTotalSupply(i)
	return knc
}

// SetNillableCharacterTotalSupply sets the "character_total_supply" field if the given value is not nil.
func (knc *KusogeeeeeeNFTCreate) SetNillableCharacterTotalSupply(i *int) *KusogeeeeeeNFTCreate {
	if i != nil {
		knc.SetCharacterTotalSupply(*i)
	}
	return knc
}

// SetID sets the "id" field.
func (knc *KusogeeeeeeNFTCreate) SetID(u uint32) *KusogeeeeeeNFTCreate {
	knc.mutation.SetID(u)
	return knc
}

// AddChangeLogIDs adds the "change_logs" edge to the KusogeeeeeeNFTChangeLog entity by IDs.
func (knc *KusogeeeeeeNFTCreate) AddChangeLogIDs(ids ...uint32) *KusogeeeeeeNFTCreate {
	knc.mutation.AddChangeLogIDs(ids...)
	return knc
}

// AddChangeLogs adds the "change_logs" edges to the KusogeeeeeeNFTChangeLog entity.
func (knc *KusogeeeeeeNFTCreate) AddChangeLogs(k ...*KusogeeeeeeNFTChangeLog) *KusogeeeeeeNFTCreate {
	ids := make([]uint32, len(k))
	for i := range k {
		ids[i] = k[i].ID
	}
	return knc.AddChangeLogIDs(ids...)
}

// Mutation returns the KusogeeeeeeNFTMutation object of the builder.
func (knc *KusogeeeeeeNFTCreate) Mutation() *KusogeeeeeeNFTMutation {
	return knc.mutation
}

// Save creates the KusogeeeeeeNFT in the database.
func (knc *KusogeeeeeeNFTCreate) Save(ctx context.Context) (*KusogeeeeeeNFT, error) {
	if err := knc.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, knc.sqlSave, knc.mutation, knc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (knc *KusogeeeeeeNFTCreate) SaveX(ctx context.Context) *KusogeeeeeeNFT {
	v, err := knc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (knc *KusogeeeeeeNFTCreate) Exec(ctx context.Context) error {
	_, err := knc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (knc *KusogeeeeeeNFTCreate) ExecX(ctx context.Context) {
	if err := knc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (knc *KusogeeeeeeNFTCreate) defaults() error {
	if _, ok := knc.mutation.CreatedAt(); !ok {
		if kusogeeeeeenft.DefaultCreatedAt == nil {
			return fmt.Errorf("ent: uninitialized kusogeeeeeenft.DefaultCreatedAt (forgotten import ent/runtime?)")
		}
		v := kusogeeeeeenft.DefaultCreatedAt()
		knc.mutation.SetCreatedAt(v)
	}
	if _, ok := knc.mutation.UpdatedAt(); !ok {
		if kusogeeeeeenft.DefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized kusogeeeeeenft.DefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := kusogeeeeeenft.DefaultUpdatedAt()
		knc.mutation.SetUpdatedAt(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (knc *KusogeeeeeeNFTCreate) check() error {
	if _, ok := knc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "KusogeeeeeeNFT.created_at"`)}
	}
	if _, ok := knc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "KusogeeeeeeNFT.updated_at"`)}
	}
	if _, ok := knc.mutation.URI(); !ok {
		return &ValidationError{Name: "uri", err: errors.New(`ent: missing required field "KusogeeeeeeNFT.uri"`)}
	}
	if _, ok := knc.mutation.GetType(); !ok {
		return &ValidationError{Name: "type", err: errors.New(`ent: missing required field "KusogeeeeeeNFT.type"`)}
	}
	if _, ok := knc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "KusogeeeeeeNFT.name"`)}
	}
	if _, ok := knc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`ent: missing required field "KusogeeeeeeNFT.status"`)}
	}
	if _, ok := knc.mutation.Price(); !ok {
		return &ValidationError{Name: "price", err: errors.New(`ent: missing required field "KusogeeeeeeNFT.price"`)}
	}
	if _, ok := knc.mutation.PublishedAt(); !ok {
		return &ValidationError{Name: "published_at", err: errors.New(`ent: missing required field "KusogeeeeeeNFT.published_at"`)}
	}
	return nil
}

func (knc *KusogeeeeeeNFTCreate) sqlSave(ctx context.Context) (*KusogeeeeeeNFT, error) {
	if err := knc.check(); err != nil {
		return nil, err
	}
	_node, _spec := knc.createSpec()
	if err := sqlgraph.CreateNode(ctx, knc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = uint32(id)
	}
	knc.mutation.id = &_node.ID
	knc.mutation.done = true
	return _node, nil
}

func (knc *KusogeeeeeeNFTCreate) createSpec() (*KusogeeeeeeNFT, *sqlgraph.CreateSpec) {
	var (
		_node = &KusogeeeeeeNFT{config: knc.config}
		_spec = sqlgraph.NewCreateSpec(kusogeeeeeenft.Table, sqlgraph.NewFieldSpec(kusogeeeeeenft.FieldID, field.TypeUint32))
	)
	if id, ok := knc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := knc.mutation.CreatedAt(); ok {
		_spec.SetField(kusogeeeeeenft.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := knc.mutation.UpdatedAt(); ok {
		_spec.SetField(kusogeeeeeenft.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := knc.mutation.DeletedAt(); ok {
		_spec.SetField(kusogeeeeeenft.FieldDeletedAt, field.TypeTime, value)
		_node.DeletedAt = &value
	}
	if value, ok := knc.mutation.URI(); ok {
		_spec.SetField(kusogeeeeeenft.FieldURI, field.TypeString, value)
		_node.URI = value
	}
	if value, ok := knc.mutation.GetType(); ok {
		_spec.SetField(kusogeeeeeenft.FieldType, field.TypeString, value)
		_node.Type = value
	}
	if value, ok := knc.mutation.Name(); ok {
		_spec.SetField(kusogeeeeeenft.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := knc.mutation.Status(); ok {
		_spec.SetField(kusogeeeeeenft.FieldStatus, field.TypeString, value)
		_node.Status = value
	}
	if value, ok := knc.mutation.Price(); ok {
		_spec.SetField(kusogeeeeeenft.FieldPrice, field.TypeInt, value)
		_node.Price = value
	}
	if value, ok := knc.mutation.PublishedAt(); ok {
		_spec.SetField(kusogeeeeeenft.FieldPublishedAt, field.TypeTime, value)
		_node.PublishedAt = value
	}
	if value, ok := knc.mutation.WeaponRank(); ok {
		_spec.SetField(kusogeeeeeenft.FieldWeaponRank, field.TypeInt, value)
		_node.WeaponRank = &value
	}
	if value, ok := knc.mutation.WeaponType(); ok {
		_spec.SetField(kusogeeeeeenft.FieldWeaponType, field.TypeString, value)
		_node.WeaponType = &value
	}
	if value, ok := knc.mutation.WeaponVitality(); ok {
		_spec.SetField(kusogeeeeeenft.FieldWeaponVitality, field.TypeInt, value)
		_node.WeaponVitality = &value
	}
	if value, ok := knc.mutation.WeaponStrength(); ok {
		_spec.SetField(kusogeeeeeenft.FieldWeaponStrength, field.TypeInt, value)
		_node.WeaponStrength = &value
	}
	if value, ok := knc.mutation.WeaponPhysicalDefense(); ok {
		_spec.SetField(kusogeeeeeenft.FieldWeaponPhysicalDefense, field.TypeInt, value)
		_node.WeaponPhysicalDefense = &value
	}
	if value, ok := knc.mutation.WeaponMagicalDefense(); ok {
		_spec.SetField(kusogeeeeeenft.FieldWeaponMagicalDefense, field.TypeInt, value)
		_node.WeaponMagicalDefense = &value
	}
	if value, ok := knc.mutation.WeaponAgility(); ok {
		_spec.SetField(kusogeeeeeenft.FieldWeaponAgility, field.TypeInt, value)
		_node.WeaponAgility = &value
	}
	if value, ok := knc.mutation.CharacterRank(); ok {
		_spec.SetField(kusogeeeeeenft.FieldCharacterRank, field.TypeString, value)
		_node.CharacterRank = &value
	}
	if value, ok := knc.mutation.CharacterTotalSupply(); ok {
		_spec.SetField(kusogeeeeeenft.FieldCharacterTotalSupply, field.TypeInt, value)
		_node.CharacterTotalSupply = &value
	}
	if nodes := knc.mutation.ChangeLogsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   kusogeeeeeenft.ChangeLogsTable,
			Columns: []string{kusogeeeeeenft.ChangeLogsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(kusogeeeeeenftchangelog.FieldID, field.TypeUint32),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// KusogeeeeeeNFTCreateBulk is the builder for creating many KusogeeeeeeNFT entities in bulk.
type KusogeeeeeeNFTCreateBulk struct {
	config
	err      error
	builders []*KusogeeeeeeNFTCreate
}

// Save creates the KusogeeeeeeNFT entities in the database.
func (kncb *KusogeeeeeeNFTCreateBulk) Save(ctx context.Context) ([]*KusogeeeeeeNFT, error) {
	if kncb.err != nil {
		return nil, kncb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(kncb.builders))
	nodes := make([]*KusogeeeeeeNFT, len(kncb.builders))
	mutators := make([]Mutator, len(kncb.builders))
	for i := range kncb.builders {
		func(i int, root context.Context) {
			builder := kncb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*KusogeeeeeeNFTMutation)
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
					_, err = mutators[i+1].Mutate(root, kncb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, kncb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, kncb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (kncb *KusogeeeeeeNFTCreateBulk) SaveX(ctx context.Context) []*KusogeeeeeeNFT {
	v, err := kncb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (kncb *KusogeeeeeeNFTCreateBulk) Exec(ctx context.Context) error {
	_, err := kncb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (kncb *KusogeeeeeeNFTCreateBulk) ExecX(ctx context.Context) {
	if err := kncb.Exec(ctx); err != nil {
		panic(err)
	}
}
