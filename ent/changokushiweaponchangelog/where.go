// Code generated by ent, DO NOT EDIT.

package changokushiweaponchangelog

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/a10adotapp/a10a.app/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uint32) predicate.ChangokushiWeaponChangeLog {
	return predicate.ChangokushiWeaponChangeLog(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uint32) predicate.ChangokushiWeaponChangeLog {
	return predicate.ChangokushiWeaponChangeLog(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uint32) predicate.ChangokushiWeaponChangeLog {
	return predicate.ChangokushiWeaponChangeLog(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uint32) predicate.ChangokushiWeaponChangeLog {
	return predicate.ChangokushiWeaponChangeLog(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uint32) predicate.ChangokushiWeaponChangeLog {
	return predicate.ChangokushiWeaponChangeLog(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uint32) predicate.ChangokushiWeaponChangeLog {
	return predicate.ChangokushiWeaponChangeLog(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uint32) predicate.ChangokushiWeaponChangeLog {
	return predicate.ChangokushiWeaponChangeLog(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uint32) predicate.ChangokushiWeaponChangeLog {
	return predicate.ChangokushiWeaponChangeLog(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uint32) predicate.ChangokushiWeaponChangeLog {
	return predicate.ChangokushiWeaponChangeLog(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.ChangokushiWeaponChangeLog {
	return predicate.ChangokushiWeaponChangeLog(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.ChangokushiWeaponChangeLog {
	return predicate.ChangokushiWeaponChangeLog(sql.FieldEQ(FieldUpdatedAt, v))
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v time.Time) predicate.ChangokushiWeaponChangeLog {
	return predicate.ChangokushiWeaponChangeLog(sql.FieldEQ(FieldDeletedAt, v))
}

// ChangokushiWeaponID applies equality check predicate on the "changokushi_weapon_id" field. It's identical to ChangokushiWeaponIDEQ.
func ChangokushiWeaponID(v uint32) predicate.ChangokushiWeaponChangeLog {
	return predicate.ChangokushiWeaponChangeLog(sql.FieldEQ(FieldChangokushiWeaponID, v))
}

// Status applies equality check predicate on the "status" field. It's identical to StatusEQ.
func Status(v string) predicate.ChangokushiWeaponChangeLog {
	return predicate.ChangokushiWeaponChangeLog(sql.FieldEQ(FieldStatus, v))
}

// Price applies equality check predicate on the "price" field. It's identical to PriceEQ.
func Price(v int) predicate.ChangokushiWeaponChangeLog {
	return predicate.ChangokushiWeaponChangeLog(sql.FieldEQ(FieldPrice, v))
}

// PublishedAt applies equality check predicate on the "published_at" field. It's identical to PublishedAtEQ.
func PublishedAt(v time.Time) predicate.ChangokushiWeaponChangeLog {
	return predicate.ChangokushiWeaponChangeLog(sql.FieldEQ(FieldPublishedAt, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.ChangokushiWeaponChangeLog {
	return predicate.ChangokushiWeaponChangeLog(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.ChangokushiWeaponChangeLog {
	return predicate.ChangokushiWeaponChangeLog(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.ChangokushiWeaponChangeLog {
	return predicate.ChangokushiWeaponChangeLog(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.ChangokushiWeaponChangeLog {
	return predicate.ChangokushiWeaponChangeLog(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.ChangokushiWeaponChangeLog {
	return predicate.ChangokushiWeaponChangeLog(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.ChangokushiWeaponChangeLog {
	return predicate.ChangokushiWeaponChangeLog(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.ChangokushiWeaponChangeLog {
	return predicate.ChangokushiWeaponChangeLog(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.ChangokushiWeaponChangeLog {
	return predicate.ChangokushiWeaponChangeLog(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.ChangokushiWeaponChangeLog {
	return predicate.ChangokushiWeaponChangeLog(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.ChangokushiWeaponChangeLog {
	return predicate.ChangokushiWeaponChangeLog(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.ChangokushiWeaponChangeLog {
	return predicate.ChangokushiWeaponChangeLog(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.ChangokushiWeaponChangeLog {
	return predicate.ChangokushiWeaponChangeLog(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.ChangokushiWeaponChangeLog {
	return predicate.ChangokushiWeaponChangeLog(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.ChangokushiWeaponChangeLog {
	return predicate.ChangokushiWeaponChangeLog(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.ChangokushiWeaponChangeLog {
	return predicate.ChangokushiWeaponChangeLog(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.ChangokushiWeaponChangeLog {
	return predicate.ChangokushiWeaponChangeLog(sql.FieldLTE(FieldUpdatedAt, v))
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v time.Time) predicate.ChangokushiWeaponChangeLog {
	return predicate.ChangokushiWeaponChangeLog(sql.FieldEQ(FieldDeletedAt, v))
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v time.Time) predicate.ChangokushiWeaponChangeLog {
	return predicate.ChangokushiWeaponChangeLog(sql.FieldNEQ(FieldDeletedAt, v))
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...time.Time) predicate.ChangokushiWeaponChangeLog {
	return predicate.ChangokushiWeaponChangeLog(sql.FieldIn(FieldDeletedAt, vs...))
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...time.Time) predicate.ChangokushiWeaponChangeLog {
	return predicate.ChangokushiWeaponChangeLog(sql.FieldNotIn(FieldDeletedAt, vs...))
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v time.Time) predicate.ChangokushiWeaponChangeLog {
	return predicate.ChangokushiWeaponChangeLog(sql.FieldGT(FieldDeletedAt, v))
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v time.Time) predicate.ChangokushiWeaponChangeLog {
	return predicate.ChangokushiWeaponChangeLog(sql.FieldGTE(FieldDeletedAt, v))
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v time.Time) predicate.ChangokushiWeaponChangeLog {
	return predicate.ChangokushiWeaponChangeLog(sql.FieldLT(FieldDeletedAt, v))
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v time.Time) predicate.ChangokushiWeaponChangeLog {
	return predicate.ChangokushiWeaponChangeLog(sql.FieldLTE(FieldDeletedAt, v))
}

// DeletedAtIsNil applies the IsNil predicate on the "deleted_at" field.
func DeletedAtIsNil() predicate.ChangokushiWeaponChangeLog {
	return predicate.ChangokushiWeaponChangeLog(sql.FieldIsNull(FieldDeletedAt))
}

// DeletedAtNotNil applies the NotNil predicate on the "deleted_at" field.
func DeletedAtNotNil() predicate.ChangokushiWeaponChangeLog {
	return predicate.ChangokushiWeaponChangeLog(sql.FieldNotNull(FieldDeletedAt))
}

// ChangokushiWeaponIDEQ applies the EQ predicate on the "changokushi_weapon_id" field.
func ChangokushiWeaponIDEQ(v uint32) predicate.ChangokushiWeaponChangeLog {
	return predicate.ChangokushiWeaponChangeLog(sql.FieldEQ(FieldChangokushiWeaponID, v))
}

// ChangokushiWeaponIDNEQ applies the NEQ predicate on the "changokushi_weapon_id" field.
func ChangokushiWeaponIDNEQ(v uint32) predicate.ChangokushiWeaponChangeLog {
	return predicate.ChangokushiWeaponChangeLog(sql.FieldNEQ(FieldChangokushiWeaponID, v))
}

// ChangokushiWeaponIDIn applies the In predicate on the "changokushi_weapon_id" field.
func ChangokushiWeaponIDIn(vs ...uint32) predicate.ChangokushiWeaponChangeLog {
	return predicate.ChangokushiWeaponChangeLog(sql.FieldIn(FieldChangokushiWeaponID, vs...))
}

// ChangokushiWeaponIDNotIn applies the NotIn predicate on the "changokushi_weapon_id" field.
func ChangokushiWeaponIDNotIn(vs ...uint32) predicate.ChangokushiWeaponChangeLog {
	return predicate.ChangokushiWeaponChangeLog(sql.FieldNotIn(FieldChangokushiWeaponID, vs...))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v string) predicate.ChangokushiWeaponChangeLog {
	return predicate.ChangokushiWeaponChangeLog(sql.FieldEQ(FieldStatus, v))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v string) predicate.ChangokushiWeaponChangeLog {
	return predicate.ChangokushiWeaponChangeLog(sql.FieldNEQ(FieldStatus, v))
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...string) predicate.ChangokushiWeaponChangeLog {
	return predicate.ChangokushiWeaponChangeLog(sql.FieldIn(FieldStatus, vs...))
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...string) predicate.ChangokushiWeaponChangeLog {
	return predicate.ChangokushiWeaponChangeLog(sql.FieldNotIn(FieldStatus, vs...))
}

// StatusGT applies the GT predicate on the "status" field.
func StatusGT(v string) predicate.ChangokushiWeaponChangeLog {
	return predicate.ChangokushiWeaponChangeLog(sql.FieldGT(FieldStatus, v))
}

// StatusGTE applies the GTE predicate on the "status" field.
func StatusGTE(v string) predicate.ChangokushiWeaponChangeLog {
	return predicate.ChangokushiWeaponChangeLog(sql.FieldGTE(FieldStatus, v))
}

// StatusLT applies the LT predicate on the "status" field.
func StatusLT(v string) predicate.ChangokushiWeaponChangeLog {
	return predicate.ChangokushiWeaponChangeLog(sql.FieldLT(FieldStatus, v))
}

// StatusLTE applies the LTE predicate on the "status" field.
func StatusLTE(v string) predicate.ChangokushiWeaponChangeLog {
	return predicate.ChangokushiWeaponChangeLog(sql.FieldLTE(FieldStatus, v))
}

// StatusContains applies the Contains predicate on the "status" field.
func StatusContains(v string) predicate.ChangokushiWeaponChangeLog {
	return predicate.ChangokushiWeaponChangeLog(sql.FieldContains(FieldStatus, v))
}

// StatusHasPrefix applies the HasPrefix predicate on the "status" field.
func StatusHasPrefix(v string) predicate.ChangokushiWeaponChangeLog {
	return predicate.ChangokushiWeaponChangeLog(sql.FieldHasPrefix(FieldStatus, v))
}

// StatusHasSuffix applies the HasSuffix predicate on the "status" field.
func StatusHasSuffix(v string) predicate.ChangokushiWeaponChangeLog {
	return predicate.ChangokushiWeaponChangeLog(sql.FieldHasSuffix(FieldStatus, v))
}

// StatusEqualFold applies the EqualFold predicate on the "status" field.
func StatusEqualFold(v string) predicate.ChangokushiWeaponChangeLog {
	return predicate.ChangokushiWeaponChangeLog(sql.FieldEqualFold(FieldStatus, v))
}

// StatusContainsFold applies the ContainsFold predicate on the "status" field.
func StatusContainsFold(v string) predicate.ChangokushiWeaponChangeLog {
	return predicate.ChangokushiWeaponChangeLog(sql.FieldContainsFold(FieldStatus, v))
}

// PriceEQ applies the EQ predicate on the "price" field.
func PriceEQ(v int) predicate.ChangokushiWeaponChangeLog {
	return predicate.ChangokushiWeaponChangeLog(sql.FieldEQ(FieldPrice, v))
}

// PriceNEQ applies the NEQ predicate on the "price" field.
func PriceNEQ(v int) predicate.ChangokushiWeaponChangeLog {
	return predicate.ChangokushiWeaponChangeLog(sql.FieldNEQ(FieldPrice, v))
}

// PriceIn applies the In predicate on the "price" field.
func PriceIn(vs ...int) predicate.ChangokushiWeaponChangeLog {
	return predicate.ChangokushiWeaponChangeLog(sql.FieldIn(FieldPrice, vs...))
}

// PriceNotIn applies the NotIn predicate on the "price" field.
func PriceNotIn(vs ...int) predicate.ChangokushiWeaponChangeLog {
	return predicate.ChangokushiWeaponChangeLog(sql.FieldNotIn(FieldPrice, vs...))
}

// PriceGT applies the GT predicate on the "price" field.
func PriceGT(v int) predicate.ChangokushiWeaponChangeLog {
	return predicate.ChangokushiWeaponChangeLog(sql.FieldGT(FieldPrice, v))
}

// PriceGTE applies the GTE predicate on the "price" field.
func PriceGTE(v int) predicate.ChangokushiWeaponChangeLog {
	return predicate.ChangokushiWeaponChangeLog(sql.FieldGTE(FieldPrice, v))
}

// PriceLT applies the LT predicate on the "price" field.
func PriceLT(v int) predicate.ChangokushiWeaponChangeLog {
	return predicate.ChangokushiWeaponChangeLog(sql.FieldLT(FieldPrice, v))
}

// PriceLTE applies the LTE predicate on the "price" field.
func PriceLTE(v int) predicate.ChangokushiWeaponChangeLog {
	return predicate.ChangokushiWeaponChangeLog(sql.FieldLTE(FieldPrice, v))
}

// PublishedAtEQ applies the EQ predicate on the "published_at" field.
func PublishedAtEQ(v time.Time) predicate.ChangokushiWeaponChangeLog {
	return predicate.ChangokushiWeaponChangeLog(sql.FieldEQ(FieldPublishedAt, v))
}

// PublishedAtNEQ applies the NEQ predicate on the "published_at" field.
func PublishedAtNEQ(v time.Time) predicate.ChangokushiWeaponChangeLog {
	return predicate.ChangokushiWeaponChangeLog(sql.FieldNEQ(FieldPublishedAt, v))
}

// PublishedAtIn applies the In predicate on the "published_at" field.
func PublishedAtIn(vs ...time.Time) predicate.ChangokushiWeaponChangeLog {
	return predicate.ChangokushiWeaponChangeLog(sql.FieldIn(FieldPublishedAt, vs...))
}

// PublishedAtNotIn applies the NotIn predicate on the "published_at" field.
func PublishedAtNotIn(vs ...time.Time) predicate.ChangokushiWeaponChangeLog {
	return predicate.ChangokushiWeaponChangeLog(sql.FieldNotIn(FieldPublishedAt, vs...))
}

// PublishedAtGT applies the GT predicate on the "published_at" field.
func PublishedAtGT(v time.Time) predicate.ChangokushiWeaponChangeLog {
	return predicate.ChangokushiWeaponChangeLog(sql.FieldGT(FieldPublishedAt, v))
}

// PublishedAtGTE applies the GTE predicate on the "published_at" field.
func PublishedAtGTE(v time.Time) predicate.ChangokushiWeaponChangeLog {
	return predicate.ChangokushiWeaponChangeLog(sql.FieldGTE(FieldPublishedAt, v))
}

// PublishedAtLT applies the LT predicate on the "published_at" field.
func PublishedAtLT(v time.Time) predicate.ChangokushiWeaponChangeLog {
	return predicate.ChangokushiWeaponChangeLog(sql.FieldLT(FieldPublishedAt, v))
}

// PublishedAtLTE applies the LTE predicate on the "published_at" field.
func PublishedAtLTE(v time.Time) predicate.ChangokushiWeaponChangeLog {
	return predicate.ChangokushiWeaponChangeLog(sql.FieldLTE(FieldPublishedAt, v))
}

// HasChangokushiWeapon applies the HasEdge predicate on the "changokushi_weapon" edge.
func HasChangokushiWeapon() predicate.ChangokushiWeaponChangeLog {
	return predicate.ChangokushiWeaponChangeLog(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ChangokushiWeaponTable, ChangokushiWeaponColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasChangokushiWeaponWith applies the HasEdge predicate on the "changokushi_weapon" edge with a given conditions (other predicates).
func HasChangokushiWeaponWith(preds ...predicate.ChangokushiWeapon) predicate.ChangokushiWeaponChangeLog {
	return predicate.ChangokushiWeaponChangeLog(func(s *sql.Selector) {
		step := newChangokushiWeaponStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.ChangokushiWeaponChangeLog) predicate.ChangokushiWeaponChangeLog {
	return predicate.ChangokushiWeaponChangeLog(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.ChangokushiWeaponChangeLog) predicate.ChangokushiWeaponChangeLog {
	return predicate.ChangokushiWeaponChangeLog(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.ChangokushiWeaponChangeLog) predicate.ChangokushiWeaponChangeLog {
	return predicate.ChangokushiWeaponChangeLog(sql.NotPredicates(p))
}
