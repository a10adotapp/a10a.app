// Code generated by ent, DO NOT EDIT.

package finschiaitemtokenmillionarthursproperty

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/a10adotapp/a10a.app/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uint32) predicate.FinschiaItemTokenMillionArthursProperty {
	return predicate.FinschiaItemTokenMillionArthursProperty(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uint32) predicate.FinschiaItemTokenMillionArthursProperty {
	return predicate.FinschiaItemTokenMillionArthursProperty(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uint32) predicate.FinschiaItemTokenMillionArthursProperty {
	return predicate.FinschiaItemTokenMillionArthursProperty(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uint32) predicate.FinschiaItemTokenMillionArthursProperty {
	return predicate.FinschiaItemTokenMillionArthursProperty(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uint32) predicate.FinschiaItemTokenMillionArthursProperty {
	return predicate.FinschiaItemTokenMillionArthursProperty(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uint32) predicate.FinschiaItemTokenMillionArthursProperty {
	return predicate.FinschiaItemTokenMillionArthursProperty(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uint32) predicate.FinschiaItemTokenMillionArthursProperty {
	return predicate.FinschiaItemTokenMillionArthursProperty(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uint32) predicate.FinschiaItemTokenMillionArthursProperty {
	return predicate.FinschiaItemTokenMillionArthursProperty(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uint32) predicate.FinschiaItemTokenMillionArthursProperty {
	return predicate.FinschiaItemTokenMillionArthursProperty(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.FinschiaItemTokenMillionArthursProperty {
	return predicate.FinschiaItemTokenMillionArthursProperty(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.FinschiaItemTokenMillionArthursProperty {
	return predicate.FinschiaItemTokenMillionArthursProperty(sql.FieldEQ(FieldUpdatedAt, v))
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v time.Time) predicate.FinschiaItemTokenMillionArthursProperty {
	return predicate.FinschiaItemTokenMillionArthursProperty(sql.FieldEQ(FieldDeletedAt, v))
}

// FinschiaItemTokenID applies equality check predicate on the "finschia_item_token_id" field. It's identical to FinschiaItemTokenIDEQ.
func FinschiaItemTokenID(v uint32) predicate.FinschiaItemTokenMillionArthursProperty {
	return predicate.FinschiaItemTokenMillionArthursProperty(sql.FieldEQ(FieldFinschiaItemTokenID, v))
}

// Series applies equality check predicate on the "series" field. It's identical to SeriesEQ.
func Series(v string) predicate.FinschiaItemTokenMillionArthursProperty {
	return predicate.FinschiaItemTokenMillionArthursProperty(sql.FieldEQ(FieldSeries, v))
}

// GearCategory applies equality check predicate on the "gear_category" field. It's identical to GearCategoryEQ.
func GearCategory(v string) predicate.FinschiaItemTokenMillionArthursProperty {
	return predicate.FinschiaItemTokenMillionArthursProperty(sql.FieldEQ(FieldGearCategory, v))
}

// GearRarity applies equality check predicate on the "gear_rarity" field. It's identical to GearRarityEQ.
func GearRarity(v string) predicate.FinschiaItemTokenMillionArthursProperty {
	return predicate.FinschiaItemTokenMillionArthursProperty(sql.FieldEQ(FieldGearRarity, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.FinschiaItemTokenMillionArthursProperty {
	return predicate.FinschiaItemTokenMillionArthursProperty(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.FinschiaItemTokenMillionArthursProperty {
	return predicate.FinschiaItemTokenMillionArthursProperty(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.FinschiaItemTokenMillionArthursProperty {
	return predicate.FinschiaItemTokenMillionArthursProperty(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.FinschiaItemTokenMillionArthursProperty {
	return predicate.FinschiaItemTokenMillionArthursProperty(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.FinschiaItemTokenMillionArthursProperty {
	return predicate.FinschiaItemTokenMillionArthursProperty(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.FinschiaItemTokenMillionArthursProperty {
	return predicate.FinschiaItemTokenMillionArthursProperty(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.FinschiaItemTokenMillionArthursProperty {
	return predicate.FinschiaItemTokenMillionArthursProperty(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.FinschiaItemTokenMillionArthursProperty {
	return predicate.FinschiaItemTokenMillionArthursProperty(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.FinschiaItemTokenMillionArthursProperty {
	return predicate.FinschiaItemTokenMillionArthursProperty(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.FinschiaItemTokenMillionArthursProperty {
	return predicate.FinschiaItemTokenMillionArthursProperty(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.FinschiaItemTokenMillionArthursProperty {
	return predicate.FinschiaItemTokenMillionArthursProperty(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.FinschiaItemTokenMillionArthursProperty {
	return predicate.FinschiaItemTokenMillionArthursProperty(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.FinschiaItemTokenMillionArthursProperty {
	return predicate.FinschiaItemTokenMillionArthursProperty(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.FinschiaItemTokenMillionArthursProperty {
	return predicate.FinschiaItemTokenMillionArthursProperty(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.FinschiaItemTokenMillionArthursProperty {
	return predicate.FinschiaItemTokenMillionArthursProperty(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.FinschiaItemTokenMillionArthursProperty {
	return predicate.FinschiaItemTokenMillionArthursProperty(sql.FieldLTE(FieldUpdatedAt, v))
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v time.Time) predicate.FinschiaItemTokenMillionArthursProperty {
	return predicate.FinschiaItemTokenMillionArthursProperty(sql.FieldEQ(FieldDeletedAt, v))
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v time.Time) predicate.FinschiaItemTokenMillionArthursProperty {
	return predicate.FinschiaItemTokenMillionArthursProperty(sql.FieldNEQ(FieldDeletedAt, v))
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...time.Time) predicate.FinschiaItemTokenMillionArthursProperty {
	return predicate.FinschiaItemTokenMillionArthursProperty(sql.FieldIn(FieldDeletedAt, vs...))
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...time.Time) predicate.FinschiaItemTokenMillionArthursProperty {
	return predicate.FinschiaItemTokenMillionArthursProperty(sql.FieldNotIn(FieldDeletedAt, vs...))
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v time.Time) predicate.FinschiaItemTokenMillionArthursProperty {
	return predicate.FinschiaItemTokenMillionArthursProperty(sql.FieldGT(FieldDeletedAt, v))
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v time.Time) predicate.FinschiaItemTokenMillionArthursProperty {
	return predicate.FinschiaItemTokenMillionArthursProperty(sql.FieldGTE(FieldDeletedAt, v))
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v time.Time) predicate.FinschiaItemTokenMillionArthursProperty {
	return predicate.FinschiaItemTokenMillionArthursProperty(sql.FieldLT(FieldDeletedAt, v))
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v time.Time) predicate.FinschiaItemTokenMillionArthursProperty {
	return predicate.FinschiaItemTokenMillionArthursProperty(sql.FieldLTE(FieldDeletedAt, v))
}

// DeletedAtIsNil applies the IsNil predicate on the "deleted_at" field.
func DeletedAtIsNil() predicate.FinschiaItemTokenMillionArthursProperty {
	return predicate.FinschiaItemTokenMillionArthursProperty(sql.FieldIsNull(FieldDeletedAt))
}

// DeletedAtNotNil applies the NotNil predicate on the "deleted_at" field.
func DeletedAtNotNil() predicate.FinschiaItemTokenMillionArthursProperty {
	return predicate.FinschiaItemTokenMillionArthursProperty(sql.FieldNotNull(FieldDeletedAt))
}

// FinschiaItemTokenIDEQ applies the EQ predicate on the "finschia_item_token_id" field.
func FinschiaItemTokenIDEQ(v uint32) predicate.FinschiaItemTokenMillionArthursProperty {
	return predicate.FinschiaItemTokenMillionArthursProperty(sql.FieldEQ(FieldFinschiaItemTokenID, v))
}

// FinschiaItemTokenIDNEQ applies the NEQ predicate on the "finschia_item_token_id" field.
func FinschiaItemTokenIDNEQ(v uint32) predicate.FinschiaItemTokenMillionArthursProperty {
	return predicate.FinschiaItemTokenMillionArthursProperty(sql.FieldNEQ(FieldFinschiaItemTokenID, v))
}

// FinschiaItemTokenIDIn applies the In predicate on the "finschia_item_token_id" field.
func FinschiaItemTokenIDIn(vs ...uint32) predicate.FinschiaItemTokenMillionArthursProperty {
	return predicate.FinschiaItemTokenMillionArthursProperty(sql.FieldIn(FieldFinschiaItemTokenID, vs...))
}

// FinschiaItemTokenIDNotIn applies the NotIn predicate on the "finschia_item_token_id" field.
func FinschiaItemTokenIDNotIn(vs ...uint32) predicate.FinschiaItemTokenMillionArthursProperty {
	return predicate.FinschiaItemTokenMillionArthursProperty(sql.FieldNotIn(FieldFinschiaItemTokenID, vs...))
}

// SeriesEQ applies the EQ predicate on the "series" field.
func SeriesEQ(v string) predicate.FinschiaItemTokenMillionArthursProperty {
	return predicate.FinschiaItemTokenMillionArthursProperty(sql.FieldEQ(FieldSeries, v))
}

// SeriesNEQ applies the NEQ predicate on the "series" field.
func SeriesNEQ(v string) predicate.FinschiaItemTokenMillionArthursProperty {
	return predicate.FinschiaItemTokenMillionArthursProperty(sql.FieldNEQ(FieldSeries, v))
}

// SeriesIn applies the In predicate on the "series" field.
func SeriesIn(vs ...string) predicate.FinschiaItemTokenMillionArthursProperty {
	return predicate.FinschiaItemTokenMillionArthursProperty(sql.FieldIn(FieldSeries, vs...))
}

// SeriesNotIn applies the NotIn predicate on the "series" field.
func SeriesNotIn(vs ...string) predicate.FinschiaItemTokenMillionArthursProperty {
	return predicate.FinschiaItemTokenMillionArthursProperty(sql.FieldNotIn(FieldSeries, vs...))
}

// SeriesGT applies the GT predicate on the "series" field.
func SeriesGT(v string) predicate.FinschiaItemTokenMillionArthursProperty {
	return predicate.FinschiaItemTokenMillionArthursProperty(sql.FieldGT(FieldSeries, v))
}

// SeriesGTE applies the GTE predicate on the "series" field.
func SeriesGTE(v string) predicate.FinschiaItemTokenMillionArthursProperty {
	return predicate.FinschiaItemTokenMillionArthursProperty(sql.FieldGTE(FieldSeries, v))
}

// SeriesLT applies the LT predicate on the "series" field.
func SeriesLT(v string) predicate.FinschiaItemTokenMillionArthursProperty {
	return predicate.FinschiaItemTokenMillionArthursProperty(sql.FieldLT(FieldSeries, v))
}

// SeriesLTE applies the LTE predicate on the "series" field.
func SeriesLTE(v string) predicate.FinschiaItemTokenMillionArthursProperty {
	return predicate.FinschiaItemTokenMillionArthursProperty(sql.FieldLTE(FieldSeries, v))
}

// SeriesContains applies the Contains predicate on the "series" field.
func SeriesContains(v string) predicate.FinschiaItemTokenMillionArthursProperty {
	return predicate.FinschiaItemTokenMillionArthursProperty(sql.FieldContains(FieldSeries, v))
}

// SeriesHasPrefix applies the HasPrefix predicate on the "series" field.
func SeriesHasPrefix(v string) predicate.FinschiaItemTokenMillionArthursProperty {
	return predicate.FinschiaItemTokenMillionArthursProperty(sql.FieldHasPrefix(FieldSeries, v))
}

// SeriesHasSuffix applies the HasSuffix predicate on the "series" field.
func SeriesHasSuffix(v string) predicate.FinschiaItemTokenMillionArthursProperty {
	return predicate.FinschiaItemTokenMillionArthursProperty(sql.FieldHasSuffix(FieldSeries, v))
}

// SeriesIsNil applies the IsNil predicate on the "series" field.
func SeriesIsNil() predicate.FinschiaItemTokenMillionArthursProperty {
	return predicate.FinschiaItemTokenMillionArthursProperty(sql.FieldIsNull(FieldSeries))
}

// SeriesNotNil applies the NotNil predicate on the "series" field.
func SeriesNotNil() predicate.FinschiaItemTokenMillionArthursProperty {
	return predicate.FinschiaItemTokenMillionArthursProperty(sql.FieldNotNull(FieldSeries))
}

// SeriesEqualFold applies the EqualFold predicate on the "series" field.
func SeriesEqualFold(v string) predicate.FinschiaItemTokenMillionArthursProperty {
	return predicate.FinschiaItemTokenMillionArthursProperty(sql.FieldEqualFold(FieldSeries, v))
}

// SeriesContainsFold applies the ContainsFold predicate on the "series" field.
func SeriesContainsFold(v string) predicate.FinschiaItemTokenMillionArthursProperty {
	return predicate.FinschiaItemTokenMillionArthursProperty(sql.FieldContainsFold(FieldSeries, v))
}

// GearCategoryEQ applies the EQ predicate on the "gear_category" field.
func GearCategoryEQ(v string) predicate.FinschiaItemTokenMillionArthursProperty {
	return predicate.FinschiaItemTokenMillionArthursProperty(sql.FieldEQ(FieldGearCategory, v))
}

// GearCategoryNEQ applies the NEQ predicate on the "gear_category" field.
func GearCategoryNEQ(v string) predicate.FinschiaItemTokenMillionArthursProperty {
	return predicate.FinschiaItemTokenMillionArthursProperty(sql.FieldNEQ(FieldGearCategory, v))
}

// GearCategoryIn applies the In predicate on the "gear_category" field.
func GearCategoryIn(vs ...string) predicate.FinschiaItemTokenMillionArthursProperty {
	return predicate.FinschiaItemTokenMillionArthursProperty(sql.FieldIn(FieldGearCategory, vs...))
}

// GearCategoryNotIn applies the NotIn predicate on the "gear_category" field.
func GearCategoryNotIn(vs ...string) predicate.FinschiaItemTokenMillionArthursProperty {
	return predicate.FinschiaItemTokenMillionArthursProperty(sql.FieldNotIn(FieldGearCategory, vs...))
}

// GearCategoryGT applies the GT predicate on the "gear_category" field.
func GearCategoryGT(v string) predicate.FinschiaItemTokenMillionArthursProperty {
	return predicate.FinschiaItemTokenMillionArthursProperty(sql.FieldGT(FieldGearCategory, v))
}

// GearCategoryGTE applies the GTE predicate on the "gear_category" field.
func GearCategoryGTE(v string) predicate.FinschiaItemTokenMillionArthursProperty {
	return predicate.FinschiaItemTokenMillionArthursProperty(sql.FieldGTE(FieldGearCategory, v))
}

// GearCategoryLT applies the LT predicate on the "gear_category" field.
func GearCategoryLT(v string) predicate.FinschiaItemTokenMillionArthursProperty {
	return predicate.FinschiaItemTokenMillionArthursProperty(sql.FieldLT(FieldGearCategory, v))
}

// GearCategoryLTE applies the LTE predicate on the "gear_category" field.
func GearCategoryLTE(v string) predicate.FinschiaItemTokenMillionArthursProperty {
	return predicate.FinschiaItemTokenMillionArthursProperty(sql.FieldLTE(FieldGearCategory, v))
}

// GearCategoryContains applies the Contains predicate on the "gear_category" field.
func GearCategoryContains(v string) predicate.FinschiaItemTokenMillionArthursProperty {
	return predicate.FinschiaItemTokenMillionArthursProperty(sql.FieldContains(FieldGearCategory, v))
}

// GearCategoryHasPrefix applies the HasPrefix predicate on the "gear_category" field.
func GearCategoryHasPrefix(v string) predicate.FinschiaItemTokenMillionArthursProperty {
	return predicate.FinschiaItemTokenMillionArthursProperty(sql.FieldHasPrefix(FieldGearCategory, v))
}

// GearCategoryHasSuffix applies the HasSuffix predicate on the "gear_category" field.
func GearCategoryHasSuffix(v string) predicate.FinschiaItemTokenMillionArthursProperty {
	return predicate.FinschiaItemTokenMillionArthursProperty(sql.FieldHasSuffix(FieldGearCategory, v))
}

// GearCategoryIsNil applies the IsNil predicate on the "gear_category" field.
func GearCategoryIsNil() predicate.FinschiaItemTokenMillionArthursProperty {
	return predicate.FinschiaItemTokenMillionArthursProperty(sql.FieldIsNull(FieldGearCategory))
}

// GearCategoryNotNil applies the NotNil predicate on the "gear_category" field.
func GearCategoryNotNil() predicate.FinschiaItemTokenMillionArthursProperty {
	return predicate.FinschiaItemTokenMillionArthursProperty(sql.FieldNotNull(FieldGearCategory))
}

// GearCategoryEqualFold applies the EqualFold predicate on the "gear_category" field.
func GearCategoryEqualFold(v string) predicate.FinschiaItemTokenMillionArthursProperty {
	return predicate.FinschiaItemTokenMillionArthursProperty(sql.FieldEqualFold(FieldGearCategory, v))
}

// GearCategoryContainsFold applies the ContainsFold predicate on the "gear_category" field.
func GearCategoryContainsFold(v string) predicate.FinschiaItemTokenMillionArthursProperty {
	return predicate.FinschiaItemTokenMillionArthursProperty(sql.FieldContainsFold(FieldGearCategory, v))
}

// GearRarityEQ applies the EQ predicate on the "gear_rarity" field.
func GearRarityEQ(v string) predicate.FinschiaItemTokenMillionArthursProperty {
	return predicate.FinschiaItemTokenMillionArthursProperty(sql.FieldEQ(FieldGearRarity, v))
}

// GearRarityNEQ applies the NEQ predicate on the "gear_rarity" field.
func GearRarityNEQ(v string) predicate.FinschiaItemTokenMillionArthursProperty {
	return predicate.FinschiaItemTokenMillionArthursProperty(sql.FieldNEQ(FieldGearRarity, v))
}

// GearRarityIn applies the In predicate on the "gear_rarity" field.
func GearRarityIn(vs ...string) predicate.FinschiaItemTokenMillionArthursProperty {
	return predicate.FinschiaItemTokenMillionArthursProperty(sql.FieldIn(FieldGearRarity, vs...))
}

// GearRarityNotIn applies the NotIn predicate on the "gear_rarity" field.
func GearRarityNotIn(vs ...string) predicate.FinschiaItemTokenMillionArthursProperty {
	return predicate.FinschiaItemTokenMillionArthursProperty(sql.FieldNotIn(FieldGearRarity, vs...))
}

// GearRarityGT applies the GT predicate on the "gear_rarity" field.
func GearRarityGT(v string) predicate.FinschiaItemTokenMillionArthursProperty {
	return predicate.FinschiaItemTokenMillionArthursProperty(sql.FieldGT(FieldGearRarity, v))
}

// GearRarityGTE applies the GTE predicate on the "gear_rarity" field.
func GearRarityGTE(v string) predicate.FinschiaItemTokenMillionArthursProperty {
	return predicate.FinschiaItemTokenMillionArthursProperty(sql.FieldGTE(FieldGearRarity, v))
}

// GearRarityLT applies the LT predicate on the "gear_rarity" field.
func GearRarityLT(v string) predicate.FinschiaItemTokenMillionArthursProperty {
	return predicate.FinschiaItemTokenMillionArthursProperty(sql.FieldLT(FieldGearRarity, v))
}

// GearRarityLTE applies the LTE predicate on the "gear_rarity" field.
func GearRarityLTE(v string) predicate.FinschiaItemTokenMillionArthursProperty {
	return predicate.FinschiaItemTokenMillionArthursProperty(sql.FieldLTE(FieldGearRarity, v))
}

// GearRarityContains applies the Contains predicate on the "gear_rarity" field.
func GearRarityContains(v string) predicate.FinschiaItemTokenMillionArthursProperty {
	return predicate.FinschiaItemTokenMillionArthursProperty(sql.FieldContains(FieldGearRarity, v))
}

// GearRarityHasPrefix applies the HasPrefix predicate on the "gear_rarity" field.
func GearRarityHasPrefix(v string) predicate.FinschiaItemTokenMillionArthursProperty {
	return predicate.FinschiaItemTokenMillionArthursProperty(sql.FieldHasPrefix(FieldGearRarity, v))
}

// GearRarityHasSuffix applies the HasSuffix predicate on the "gear_rarity" field.
func GearRarityHasSuffix(v string) predicate.FinschiaItemTokenMillionArthursProperty {
	return predicate.FinschiaItemTokenMillionArthursProperty(sql.FieldHasSuffix(FieldGearRarity, v))
}

// GearRarityIsNil applies the IsNil predicate on the "gear_rarity" field.
func GearRarityIsNil() predicate.FinschiaItemTokenMillionArthursProperty {
	return predicate.FinschiaItemTokenMillionArthursProperty(sql.FieldIsNull(FieldGearRarity))
}

// GearRarityNotNil applies the NotNil predicate on the "gear_rarity" field.
func GearRarityNotNil() predicate.FinschiaItemTokenMillionArthursProperty {
	return predicate.FinschiaItemTokenMillionArthursProperty(sql.FieldNotNull(FieldGearRarity))
}

// GearRarityEqualFold applies the EqualFold predicate on the "gear_rarity" field.
func GearRarityEqualFold(v string) predicate.FinschiaItemTokenMillionArthursProperty {
	return predicate.FinschiaItemTokenMillionArthursProperty(sql.FieldEqualFold(FieldGearRarity, v))
}

// GearRarityContainsFold applies the ContainsFold predicate on the "gear_rarity" field.
func GearRarityContainsFold(v string) predicate.FinschiaItemTokenMillionArthursProperty {
	return predicate.FinschiaItemTokenMillionArthursProperty(sql.FieldContainsFold(FieldGearRarity, v))
}

// HasFinschiaItemToken applies the HasEdge predicate on the "finschia_item_token" edge.
func HasFinschiaItemToken() predicate.FinschiaItemTokenMillionArthursProperty {
	return predicate.FinschiaItemTokenMillionArthursProperty(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, FinschiaItemTokenTable, FinschiaItemTokenColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasFinschiaItemTokenWith applies the HasEdge predicate on the "finschia_item_token" edge with a given conditions (other predicates).
func HasFinschiaItemTokenWith(preds ...predicate.FinschiaItemToken) predicate.FinschiaItemTokenMillionArthursProperty {
	return predicate.FinschiaItemTokenMillionArthursProperty(func(s *sql.Selector) {
		step := newFinschiaItemTokenStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.FinschiaItemTokenMillionArthursProperty) predicate.FinschiaItemTokenMillionArthursProperty {
	return predicate.FinschiaItemTokenMillionArthursProperty(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.FinschiaItemTokenMillionArthursProperty) predicate.FinschiaItemTokenMillionArthursProperty {
	return predicate.FinschiaItemTokenMillionArthursProperty(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.FinschiaItemTokenMillionArthursProperty) predicate.FinschiaItemTokenMillionArthursProperty {
	return predicate.FinschiaItemTokenMillionArthursProperty(sql.NotPredicates(p))
}