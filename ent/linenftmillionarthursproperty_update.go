// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/a10adotapp/a10a.app/ent/linenft"
	"github.com/a10adotapp/a10a.app/ent/linenftmillionarthursproperty"
	"github.com/a10adotapp/a10a.app/ent/predicate"
)

// LINENFTMillionArthursPropertyUpdate is the builder for updating LINENFTMillionArthursProperty entities.
type LINENFTMillionArthursPropertyUpdate struct {
	config
	hooks    []Hook
	mutation *LINENFTMillionArthursPropertyMutation
}

// Where appends a list predicates to the LINENFTMillionArthursPropertyUpdate builder.
func (lmapu *LINENFTMillionArthursPropertyUpdate) Where(ps ...predicate.LINENFTMillionArthursProperty) *LINENFTMillionArthursPropertyUpdate {
	lmapu.mutation.Where(ps...)
	return lmapu
}

// SetUpdatedAt sets the "updated_at" field.
func (lmapu *LINENFTMillionArthursPropertyUpdate) SetUpdatedAt(t time.Time) *LINENFTMillionArthursPropertyUpdate {
	lmapu.mutation.SetUpdatedAt(t)
	return lmapu
}

// SetDeletedAt sets the "deleted_at" field.
func (lmapu *LINENFTMillionArthursPropertyUpdate) SetDeletedAt(t time.Time) *LINENFTMillionArthursPropertyUpdate {
	lmapu.mutation.SetDeletedAt(t)
	return lmapu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (lmapu *LINENFTMillionArthursPropertyUpdate) SetNillableDeletedAt(t *time.Time) *LINENFTMillionArthursPropertyUpdate {
	if t != nil {
		lmapu.SetDeletedAt(*t)
	}
	return lmapu
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (lmapu *LINENFTMillionArthursPropertyUpdate) ClearDeletedAt() *LINENFTMillionArthursPropertyUpdate {
	lmapu.mutation.ClearDeletedAt()
	return lmapu
}

// SetSeries sets the "series" field.
func (lmapu *LINENFTMillionArthursPropertyUpdate) SetSeries(s string) *LINENFTMillionArthursPropertyUpdate {
	lmapu.mutation.SetSeries(s)
	return lmapu
}

// SetNillableSeries sets the "series" field if the given value is not nil.
func (lmapu *LINENFTMillionArthursPropertyUpdate) SetNillableSeries(s *string) *LINENFTMillionArthursPropertyUpdate {
	if s != nil {
		lmapu.SetSeries(*s)
	}
	return lmapu
}

// ClearSeries clears the value of the "series" field.
func (lmapu *LINENFTMillionArthursPropertyUpdate) ClearSeries() *LINENFTMillionArthursPropertyUpdate {
	lmapu.mutation.ClearSeries()
	return lmapu
}

// SetGearCategory sets the "gear_category" field.
func (lmapu *LINENFTMillionArthursPropertyUpdate) SetGearCategory(s string) *LINENFTMillionArthursPropertyUpdate {
	lmapu.mutation.SetGearCategory(s)
	return lmapu
}

// SetNillableGearCategory sets the "gear_category" field if the given value is not nil.
func (lmapu *LINENFTMillionArthursPropertyUpdate) SetNillableGearCategory(s *string) *LINENFTMillionArthursPropertyUpdate {
	if s != nil {
		lmapu.SetGearCategory(*s)
	}
	return lmapu
}

// ClearGearCategory clears the value of the "gear_category" field.
func (lmapu *LINENFTMillionArthursPropertyUpdate) ClearGearCategory() *LINENFTMillionArthursPropertyUpdate {
	lmapu.mutation.ClearGearCategory()
	return lmapu
}

// SetGearRarity sets the "gear_rarity" field.
func (lmapu *LINENFTMillionArthursPropertyUpdate) SetGearRarity(s string) *LINENFTMillionArthursPropertyUpdate {
	lmapu.mutation.SetGearRarity(s)
	return lmapu
}

// SetNillableGearRarity sets the "gear_rarity" field if the given value is not nil.
func (lmapu *LINENFTMillionArthursPropertyUpdate) SetNillableGearRarity(s *string) *LINENFTMillionArthursPropertyUpdate {
	if s != nil {
		lmapu.SetGearRarity(*s)
	}
	return lmapu
}

// ClearGearRarity clears the value of the "gear_rarity" field.
func (lmapu *LINENFTMillionArthursPropertyUpdate) ClearGearRarity() *LINENFTMillionArthursPropertyUpdate {
	lmapu.mutation.ClearGearRarity()
	return lmapu
}

// SetOmj sets the "omj" field.
func (lmapu *LINENFTMillionArthursPropertyUpdate) SetOmj(s string) *LINENFTMillionArthursPropertyUpdate {
	lmapu.mutation.SetOmj(s)
	return lmapu
}

// SetNillableOmj sets the "omj" field if the given value is not nil.
func (lmapu *LINENFTMillionArthursPropertyUpdate) SetNillableOmj(s *string) *LINENFTMillionArthursPropertyUpdate {
	if s != nil {
		lmapu.SetOmj(*s)
	}
	return lmapu
}

// ClearOmj clears the value of the "omj" field.
func (lmapu *LINENFTMillionArthursPropertyUpdate) ClearOmj() *LINENFTMillionArthursPropertyUpdate {
	lmapu.mutation.ClearOmj()
	return lmapu
}

// SetLineNftID sets the "line_nft" edge to the LINENFT entity by ID.
func (lmapu *LINENFTMillionArthursPropertyUpdate) SetLineNftID(id uint32) *LINENFTMillionArthursPropertyUpdate {
	lmapu.mutation.SetLineNftID(id)
	return lmapu
}

// SetLineNft sets the "line_nft" edge to the LINENFT entity.
func (lmapu *LINENFTMillionArthursPropertyUpdate) SetLineNft(l *LINENFT) *LINENFTMillionArthursPropertyUpdate {
	return lmapu.SetLineNftID(l.ID)
}

// Mutation returns the LINENFTMillionArthursPropertyMutation object of the builder.
func (lmapu *LINENFTMillionArthursPropertyUpdate) Mutation() *LINENFTMillionArthursPropertyMutation {
	return lmapu.mutation
}

// ClearLineNft clears the "line_nft" edge to the LINENFT entity.
func (lmapu *LINENFTMillionArthursPropertyUpdate) ClearLineNft() *LINENFTMillionArthursPropertyUpdate {
	lmapu.mutation.ClearLineNft()
	return lmapu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (lmapu *LINENFTMillionArthursPropertyUpdate) Save(ctx context.Context) (int, error) {
	if err := lmapu.defaults(); err != nil {
		return 0, err
	}
	return withHooks(ctx, lmapu.sqlSave, lmapu.mutation, lmapu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (lmapu *LINENFTMillionArthursPropertyUpdate) SaveX(ctx context.Context) int {
	affected, err := lmapu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (lmapu *LINENFTMillionArthursPropertyUpdate) Exec(ctx context.Context) error {
	_, err := lmapu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lmapu *LINENFTMillionArthursPropertyUpdate) ExecX(ctx context.Context) {
	if err := lmapu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (lmapu *LINENFTMillionArthursPropertyUpdate) defaults() error {
	if _, ok := lmapu.mutation.UpdatedAt(); !ok {
		if linenftmillionarthursproperty.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized linenftmillionarthursproperty.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := linenftmillionarthursproperty.UpdateDefaultUpdatedAt()
		lmapu.mutation.SetUpdatedAt(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (lmapu *LINENFTMillionArthursPropertyUpdate) check() error {
	if _, ok := lmapu.mutation.LineNftID(); lmapu.mutation.LineNftCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "LINENFTMillionArthursProperty.line_nft"`)
	}
	return nil
}

func (lmapu *LINENFTMillionArthursPropertyUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := lmapu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(linenftmillionarthursproperty.Table, linenftmillionarthursproperty.Columns, sqlgraph.NewFieldSpec(linenftmillionarthursproperty.FieldID, field.TypeUint32))
	if ps := lmapu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := lmapu.mutation.UpdatedAt(); ok {
		_spec.SetField(linenftmillionarthursproperty.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := lmapu.mutation.DeletedAt(); ok {
		_spec.SetField(linenftmillionarthursproperty.FieldDeletedAt, field.TypeTime, value)
	}
	if lmapu.mutation.DeletedAtCleared() {
		_spec.ClearField(linenftmillionarthursproperty.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := lmapu.mutation.Series(); ok {
		_spec.SetField(linenftmillionarthursproperty.FieldSeries, field.TypeString, value)
	}
	if lmapu.mutation.SeriesCleared() {
		_spec.ClearField(linenftmillionarthursproperty.FieldSeries, field.TypeString)
	}
	if value, ok := lmapu.mutation.GearCategory(); ok {
		_spec.SetField(linenftmillionarthursproperty.FieldGearCategory, field.TypeString, value)
	}
	if lmapu.mutation.GearCategoryCleared() {
		_spec.ClearField(linenftmillionarthursproperty.FieldGearCategory, field.TypeString)
	}
	if value, ok := lmapu.mutation.GearRarity(); ok {
		_spec.SetField(linenftmillionarthursproperty.FieldGearRarity, field.TypeString, value)
	}
	if lmapu.mutation.GearRarityCleared() {
		_spec.ClearField(linenftmillionarthursproperty.FieldGearRarity, field.TypeString)
	}
	if value, ok := lmapu.mutation.Omj(); ok {
		_spec.SetField(linenftmillionarthursproperty.FieldOmj, field.TypeString, value)
	}
	if lmapu.mutation.OmjCleared() {
		_spec.ClearField(linenftmillionarthursproperty.FieldOmj, field.TypeString)
	}
	if lmapu.mutation.LineNftCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   linenftmillionarthursproperty.LineNftTable,
			Columns: []string{linenftmillionarthursproperty.LineNftColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(linenft.FieldID, field.TypeUint32),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := lmapu.mutation.LineNftIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   linenftmillionarthursproperty.LineNftTable,
			Columns: []string{linenftmillionarthursproperty.LineNftColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(linenft.FieldID, field.TypeUint32),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, lmapu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{linenftmillionarthursproperty.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	lmapu.mutation.done = true
	return n, nil
}

// LINENFTMillionArthursPropertyUpdateOne is the builder for updating a single LINENFTMillionArthursProperty entity.
type LINENFTMillionArthursPropertyUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *LINENFTMillionArthursPropertyMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (lmapuo *LINENFTMillionArthursPropertyUpdateOne) SetUpdatedAt(t time.Time) *LINENFTMillionArthursPropertyUpdateOne {
	lmapuo.mutation.SetUpdatedAt(t)
	return lmapuo
}

// SetDeletedAt sets the "deleted_at" field.
func (lmapuo *LINENFTMillionArthursPropertyUpdateOne) SetDeletedAt(t time.Time) *LINENFTMillionArthursPropertyUpdateOne {
	lmapuo.mutation.SetDeletedAt(t)
	return lmapuo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (lmapuo *LINENFTMillionArthursPropertyUpdateOne) SetNillableDeletedAt(t *time.Time) *LINENFTMillionArthursPropertyUpdateOne {
	if t != nil {
		lmapuo.SetDeletedAt(*t)
	}
	return lmapuo
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (lmapuo *LINENFTMillionArthursPropertyUpdateOne) ClearDeletedAt() *LINENFTMillionArthursPropertyUpdateOne {
	lmapuo.mutation.ClearDeletedAt()
	return lmapuo
}

// SetSeries sets the "series" field.
func (lmapuo *LINENFTMillionArthursPropertyUpdateOne) SetSeries(s string) *LINENFTMillionArthursPropertyUpdateOne {
	lmapuo.mutation.SetSeries(s)
	return lmapuo
}

// SetNillableSeries sets the "series" field if the given value is not nil.
func (lmapuo *LINENFTMillionArthursPropertyUpdateOne) SetNillableSeries(s *string) *LINENFTMillionArthursPropertyUpdateOne {
	if s != nil {
		lmapuo.SetSeries(*s)
	}
	return lmapuo
}

// ClearSeries clears the value of the "series" field.
func (lmapuo *LINENFTMillionArthursPropertyUpdateOne) ClearSeries() *LINENFTMillionArthursPropertyUpdateOne {
	lmapuo.mutation.ClearSeries()
	return lmapuo
}

// SetGearCategory sets the "gear_category" field.
func (lmapuo *LINENFTMillionArthursPropertyUpdateOne) SetGearCategory(s string) *LINENFTMillionArthursPropertyUpdateOne {
	lmapuo.mutation.SetGearCategory(s)
	return lmapuo
}

// SetNillableGearCategory sets the "gear_category" field if the given value is not nil.
func (lmapuo *LINENFTMillionArthursPropertyUpdateOne) SetNillableGearCategory(s *string) *LINENFTMillionArthursPropertyUpdateOne {
	if s != nil {
		lmapuo.SetGearCategory(*s)
	}
	return lmapuo
}

// ClearGearCategory clears the value of the "gear_category" field.
func (lmapuo *LINENFTMillionArthursPropertyUpdateOne) ClearGearCategory() *LINENFTMillionArthursPropertyUpdateOne {
	lmapuo.mutation.ClearGearCategory()
	return lmapuo
}

// SetGearRarity sets the "gear_rarity" field.
func (lmapuo *LINENFTMillionArthursPropertyUpdateOne) SetGearRarity(s string) *LINENFTMillionArthursPropertyUpdateOne {
	lmapuo.mutation.SetGearRarity(s)
	return lmapuo
}

// SetNillableGearRarity sets the "gear_rarity" field if the given value is not nil.
func (lmapuo *LINENFTMillionArthursPropertyUpdateOne) SetNillableGearRarity(s *string) *LINENFTMillionArthursPropertyUpdateOne {
	if s != nil {
		lmapuo.SetGearRarity(*s)
	}
	return lmapuo
}

// ClearGearRarity clears the value of the "gear_rarity" field.
func (lmapuo *LINENFTMillionArthursPropertyUpdateOne) ClearGearRarity() *LINENFTMillionArthursPropertyUpdateOne {
	lmapuo.mutation.ClearGearRarity()
	return lmapuo
}

// SetOmj sets the "omj" field.
func (lmapuo *LINENFTMillionArthursPropertyUpdateOne) SetOmj(s string) *LINENFTMillionArthursPropertyUpdateOne {
	lmapuo.mutation.SetOmj(s)
	return lmapuo
}

// SetNillableOmj sets the "omj" field if the given value is not nil.
func (lmapuo *LINENFTMillionArthursPropertyUpdateOne) SetNillableOmj(s *string) *LINENFTMillionArthursPropertyUpdateOne {
	if s != nil {
		lmapuo.SetOmj(*s)
	}
	return lmapuo
}

// ClearOmj clears the value of the "omj" field.
func (lmapuo *LINENFTMillionArthursPropertyUpdateOne) ClearOmj() *LINENFTMillionArthursPropertyUpdateOne {
	lmapuo.mutation.ClearOmj()
	return lmapuo
}

// SetLineNftID sets the "line_nft" edge to the LINENFT entity by ID.
func (lmapuo *LINENFTMillionArthursPropertyUpdateOne) SetLineNftID(id uint32) *LINENFTMillionArthursPropertyUpdateOne {
	lmapuo.mutation.SetLineNftID(id)
	return lmapuo
}

// SetLineNft sets the "line_nft" edge to the LINENFT entity.
func (lmapuo *LINENFTMillionArthursPropertyUpdateOne) SetLineNft(l *LINENFT) *LINENFTMillionArthursPropertyUpdateOne {
	return lmapuo.SetLineNftID(l.ID)
}

// Mutation returns the LINENFTMillionArthursPropertyMutation object of the builder.
func (lmapuo *LINENFTMillionArthursPropertyUpdateOne) Mutation() *LINENFTMillionArthursPropertyMutation {
	return lmapuo.mutation
}

// ClearLineNft clears the "line_nft" edge to the LINENFT entity.
func (lmapuo *LINENFTMillionArthursPropertyUpdateOne) ClearLineNft() *LINENFTMillionArthursPropertyUpdateOne {
	lmapuo.mutation.ClearLineNft()
	return lmapuo
}

// Where appends a list predicates to the LINENFTMillionArthursPropertyUpdate builder.
func (lmapuo *LINENFTMillionArthursPropertyUpdateOne) Where(ps ...predicate.LINENFTMillionArthursProperty) *LINENFTMillionArthursPropertyUpdateOne {
	lmapuo.mutation.Where(ps...)
	return lmapuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (lmapuo *LINENFTMillionArthursPropertyUpdateOne) Select(field string, fields ...string) *LINENFTMillionArthursPropertyUpdateOne {
	lmapuo.fields = append([]string{field}, fields...)
	return lmapuo
}

// Save executes the query and returns the updated LINENFTMillionArthursProperty entity.
func (lmapuo *LINENFTMillionArthursPropertyUpdateOne) Save(ctx context.Context) (*LINENFTMillionArthursProperty, error) {
	if err := lmapuo.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, lmapuo.sqlSave, lmapuo.mutation, lmapuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (lmapuo *LINENFTMillionArthursPropertyUpdateOne) SaveX(ctx context.Context) *LINENFTMillionArthursProperty {
	node, err := lmapuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (lmapuo *LINENFTMillionArthursPropertyUpdateOne) Exec(ctx context.Context) error {
	_, err := lmapuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lmapuo *LINENFTMillionArthursPropertyUpdateOne) ExecX(ctx context.Context) {
	if err := lmapuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (lmapuo *LINENFTMillionArthursPropertyUpdateOne) defaults() error {
	if _, ok := lmapuo.mutation.UpdatedAt(); !ok {
		if linenftmillionarthursproperty.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized linenftmillionarthursproperty.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := linenftmillionarthursproperty.UpdateDefaultUpdatedAt()
		lmapuo.mutation.SetUpdatedAt(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (lmapuo *LINENFTMillionArthursPropertyUpdateOne) check() error {
	if _, ok := lmapuo.mutation.LineNftID(); lmapuo.mutation.LineNftCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "LINENFTMillionArthursProperty.line_nft"`)
	}
	return nil
}

func (lmapuo *LINENFTMillionArthursPropertyUpdateOne) sqlSave(ctx context.Context) (_node *LINENFTMillionArthursProperty, err error) {
	if err := lmapuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(linenftmillionarthursproperty.Table, linenftmillionarthursproperty.Columns, sqlgraph.NewFieldSpec(linenftmillionarthursproperty.FieldID, field.TypeUint32))
	id, ok := lmapuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "LINENFTMillionArthursProperty.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := lmapuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, linenftmillionarthursproperty.FieldID)
		for _, f := range fields {
			if !linenftmillionarthursproperty.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != linenftmillionarthursproperty.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := lmapuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := lmapuo.mutation.UpdatedAt(); ok {
		_spec.SetField(linenftmillionarthursproperty.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := lmapuo.mutation.DeletedAt(); ok {
		_spec.SetField(linenftmillionarthursproperty.FieldDeletedAt, field.TypeTime, value)
	}
	if lmapuo.mutation.DeletedAtCleared() {
		_spec.ClearField(linenftmillionarthursproperty.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := lmapuo.mutation.Series(); ok {
		_spec.SetField(linenftmillionarthursproperty.FieldSeries, field.TypeString, value)
	}
	if lmapuo.mutation.SeriesCleared() {
		_spec.ClearField(linenftmillionarthursproperty.FieldSeries, field.TypeString)
	}
	if value, ok := lmapuo.mutation.GearCategory(); ok {
		_spec.SetField(linenftmillionarthursproperty.FieldGearCategory, field.TypeString, value)
	}
	if lmapuo.mutation.GearCategoryCleared() {
		_spec.ClearField(linenftmillionarthursproperty.FieldGearCategory, field.TypeString)
	}
	if value, ok := lmapuo.mutation.GearRarity(); ok {
		_spec.SetField(linenftmillionarthursproperty.FieldGearRarity, field.TypeString, value)
	}
	if lmapuo.mutation.GearRarityCleared() {
		_spec.ClearField(linenftmillionarthursproperty.FieldGearRarity, field.TypeString)
	}
	if value, ok := lmapuo.mutation.Omj(); ok {
		_spec.SetField(linenftmillionarthursproperty.FieldOmj, field.TypeString, value)
	}
	if lmapuo.mutation.OmjCleared() {
		_spec.ClearField(linenftmillionarthursproperty.FieldOmj, field.TypeString)
	}
	if lmapuo.mutation.LineNftCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   linenftmillionarthursproperty.LineNftTable,
			Columns: []string{linenftmillionarthursproperty.LineNftColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(linenft.FieldID, field.TypeUint32),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := lmapuo.mutation.LineNftIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   linenftmillionarthursproperty.LineNftTable,
			Columns: []string{linenftmillionarthursproperty.LineNftColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(linenft.FieldID, field.TypeUint32),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &LINENFTMillionArthursProperty{config: lmapuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, lmapuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{linenftmillionarthursproperty.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	lmapuo.mutation.done = true
	return _node, nil
}
