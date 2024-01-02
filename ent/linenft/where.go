// Code generated by ent, DO NOT EDIT.

package linenft

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/a10adotapp/a10a.app/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uint32) predicate.LINENFT {
	return predicate.LINENFT(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uint32) predicate.LINENFT {
	return predicate.LINENFT(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uint32) predicate.LINENFT {
	return predicate.LINENFT(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uint32) predicate.LINENFT {
	return predicate.LINENFT(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uint32) predicate.LINENFT {
	return predicate.LINENFT(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uint32) predicate.LINENFT {
	return predicate.LINENFT(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uint32) predicate.LINENFT {
	return predicate.LINENFT(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uint32) predicate.LINENFT {
	return predicate.LINENFT(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uint32) predicate.LINENFT {
	return predicate.LINENFT(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.LINENFT {
	return predicate.LINENFT(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.LINENFT {
	return predicate.LINENFT(sql.FieldEQ(FieldUpdatedAt, v))
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v time.Time) predicate.LINENFT {
	return predicate.LINENFT(sql.FieldEQ(FieldDeletedAt, v))
}

// LineNftID applies equality check predicate on the "line_nft_id" field. It's identical to LineNftIDEQ.
func LineNftID(v uint32) predicate.LINENFT {
	return predicate.LINENFT(sql.FieldEQ(FieldLineNftID, v))
}

// ContractID applies equality check predicate on the "contract_id" field. It's identical to ContractIDEQ.
func ContractID(v string) predicate.LINENFT {
	return predicate.LINENFT(sql.FieldEQ(FieldContractID, v))
}

// TokenType applies equality check predicate on the "token_type" field. It's identical to TokenTypeEQ.
func TokenType(v string) predicate.LINENFT {
	return predicate.LINENFT(sql.FieldEQ(FieldTokenType, v))
}

// TokenIndex applies equality check predicate on the "token_index" field. It's identical to TokenIndexEQ.
func TokenIndex(v string) predicate.LINENFT {
	return predicate.LINENFT(sql.FieldEQ(FieldTokenIndex, v))
}

// TokenName applies equality check predicate on the "token_name" field. It's identical to TokenNameEQ.
func TokenName(v string) predicate.LINENFT {
	return predicate.LINENFT(sql.FieldEQ(FieldTokenName, v))
}

// TokenDescription applies equality check predicate on the "token_description" field. It's identical to TokenDescriptionEQ.
func TokenDescription(v string) predicate.LINENFT {
	return predicate.LINENFT(sql.FieldEQ(FieldTokenDescription, v))
}

// TokenContentURL applies equality check predicate on the "token_content_url" field. It's identical to TokenContentURLEQ.
func TokenContentURL(v string) predicate.LINENFT {
	return predicate.LINENFT(sql.FieldEQ(FieldTokenContentURL, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.LINENFT {
	return predicate.LINENFT(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.LINENFT {
	return predicate.LINENFT(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.LINENFT {
	return predicate.LINENFT(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.LINENFT {
	return predicate.LINENFT(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.LINENFT {
	return predicate.LINENFT(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.LINENFT {
	return predicate.LINENFT(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.LINENFT {
	return predicate.LINENFT(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.LINENFT {
	return predicate.LINENFT(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.LINENFT {
	return predicate.LINENFT(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.LINENFT {
	return predicate.LINENFT(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.LINENFT {
	return predicate.LINENFT(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.LINENFT {
	return predicate.LINENFT(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.LINENFT {
	return predicate.LINENFT(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.LINENFT {
	return predicate.LINENFT(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.LINENFT {
	return predicate.LINENFT(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.LINENFT {
	return predicate.LINENFT(sql.FieldLTE(FieldUpdatedAt, v))
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v time.Time) predicate.LINENFT {
	return predicate.LINENFT(sql.FieldEQ(FieldDeletedAt, v))
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v time.Time) predicate.LINENFT {
	return predicate.LINENFT(sql.FieldNEQ(FieldDeletedAt, v))
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...time.Time) predicate.LINENFT {
	return predicate.LINENFT(sql.FieldIn(FieldDeletedAt, vs...))
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...time.Time) predicate.LINENFT {
	return predicate.LINENFT(sql.FieldNotIn(FieldDeletedAt, vs...))
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v time.Time) predicate.LINENFT {
	return predicate.LINENFT(sql.FieldGT(FieldDeletedAt, v))
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v time.Time) predicate.LINENFT {
	return predicate.LINENFT(sql.FieldGTE(FieldDeletedAt, v))
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v time.Time) predicate.LINENFT {
	return predicate.LINENFT(sql.FieldLT(FieldDeletedAt, v))
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v time.Time) predicate.LINENFT {
	return predicate.LINENFT(sql.FieldLTE(FieldDeletedAt, v))
}

// DeletedAtIsNil applies the IsNil predicate on the "deleted_at" field.
func DeletedAtIsNil() predicate.LINENFT {
	return predicate.LINENFT(sql.FieldIsNull(FieldDeletedAt))
}

// DeletedAtNotNil applies the NotNil predicate on the "deleted_at" field.
func DeletedAtNotNil() predicate.LINENFT {
	return predicate.LINENFT(sql.FieldNotNull(FieldDeletedAt))
}

// LineNftIDEQ applies the EQ predicate on the "line_nft_id" field.
func LineNftIDEQ(v uint32) predicate.LINENFT {
	return predicate.LINENFT(sql.FieldEQ(FieldLineNftID, v))
}

// LineNftIDNEQ applies the NEQ predicate on the "line_nft_id" field.
func LineNftIDNEQ(v uint32) predicate.LINENFT {
	return predicate.LINENFT(sql.FieldNEQ(FieldLineNftID, v))
}

// LineNftIDIn applies the In predicate on the "line_nft_id" field.
func LineNftIDIn(vs ...uint32) predicate.LINENFT {
	return predicate.LINENFT(sql.FieldIn(FieldLineNftID, vs...))
}

// LineNftIDNotIn applies the NotIn predicate on the "line_nft_id" field.
func LineNftIDNotIn(vs ...uint32) predicate.LINENFT {
	return predicate.LINENFT(sql.FieldNotIn(FieldLineNftID, vs...))
}

// LineNftIDGT applies the GT predicate on the "line_nft_id" field.
func LineNftIDGT(v uint32) predicate.LINENFT {
	return predicate.LINENFT(sql.FieldGT(FieldLineNftID, v))
}

// LineNftIDGTE applies the GTE predicate on the "line_nft_id" field.
func LineNftIDGTE(v uint32) predicate.LINENFT {
	return predicate.LINENFT(sql.FieldGTE(FieldLineNftID, v))
}

// LineNftIDLT applies the LT predicate on the "line_nft_id" field.
func LineNftIDLT(v uint32) predicate.LINENFT {
	return predicate.LINENFT(sql.FieldLT(FieldLineNftID, v))
}

// LineNftIDLTE applies the LTE predicate on the "line_nft_id" field.
func LineNftIDLTE(v uint32) predicate.LINENFT {
	return predicate.LINENFT(sql.FieldLTE(FieldLineNftID, v))
}

// ContractIDEQ applies the EQ predicate on the "contract_id" field.
func ContractIDEQ(v string) predicate.LINENFT {
	return predicate.LINENFT(sql.FieldEQ(FieldContractID, v))
}

// ContractIDNEQ applies the NEQ predicate on the "contract_id" field.
func ContractIDNEQ(v string) predicate.LINENFT {
	return predicate.LINENFT(sql.FieldNEQ(FieldContractID, v))
}

// ContractIDIn applies the In predicate on the "contract_id" field.
func ContractIDIn(vs ...string) predicate.LINENFT {
	return predicate.LINENFT(sql.FieldIn(FieldContractID, vs...))
}

// ContractIDNotIn applies the NotIn predicate on the "contract_id" field.
func ContractIDNotIn(vs ...string) predicate.LINENFT {
	return predicate.LINENFT(sql.FieldNotIn(FieldContractID, vs...))
}

// ContractIDGT applies the GT predicate on the "contract_id" field.
func ContractIDGT(v string) predicate.LINENFT {
	return predicate.LINENFT(sql.FieldGT(FieldContractID, v))
}

// ContractIDGTE applies the GTE predicate on the "contract_id" field.
func ContractIDGTE(v string) predicate.LINENFT {
	return predicate.LINENFT(sql.FieldGTE(FieldContractID, v))
}

// ContractIDLT applies the LT predicate on the "contract_id" field.
func ContractIDLT(v string) predicate.LINENFT {
	return predicate.LINENFT(sql.FieldLT(FieldContractID, v))
}

// ContractIDLTE applies the LTE predicate on the "contract_id" field.
func ContractIDLTE(v string) predicate.LINENFT {
	return predicate.LINENFT(sql.FieldLTE(FieldContractID, v))
}

// ContractIDContains applies the Contains predicate on the "contract_id" field.
func ContractIDContains(v string) predicate.LINENFT {
	return predicate.LINENFT(sql.FieldContains(FieldContractID, v))
}

// ContractIDHasPrefix applies the HasPrefix predicate on the "contract_id" field.
func ContractIDHasPrefix(v string) predicate.LINENFT {
	return predicate.LINENFT(sql.FieldHasPrefix(FieldContractID, v))
}

// ContractIDHasSuffix applies the HasSuffix predicate on the "contract_id" field.
func ContractIDHasSuffix(v string) predicate.LINENFT {
	return predicate.LINENFT(sql.FieldHasSuffix(FieldContractID, v))
}

// ContractIDEqualFold applies the EqualFold predicate on the "contract_id" field.
func ContractIDEqualFold(v string) predicate.LINENFT {
	return predicate.LINENFT(sql.FieldEqualFold(FieldContractID, v))
}

// ContractIDContainsFold applies the ContainsFold predicate on the "contract_id" field.
func ContractIDContainsFold(v string) predicate.LINENFT {
	return predicate.LINENFT(sql.FieldContainsFold(FieldContractID, v))
}

// TokenTypeEQ applies the EQ predicate on the "token_type" field.
func TokenTypeEQ(v string) predicate.LINENFT {
	return predicate.LINENFT(sql.FieldEQ(FieldTokenType, v))
}

// TokenTypeNEQ applies the NEQ predicate on the "token_type" field.
func TokenTypeNEQ(v string) predicate.LINENFT {
	return predicate.LINENFT(sql.FieldNEQ(FieldTokenType, v))
}

// TokenTypeIn applies the In predicate on the "token_type" field.
func TokenTypeIn(vs ...string) predicate.LINENFT {
	return predicate.LINENFT(sql.FieldIn(FieldTokenType, vs...))
}

// TokenTypeNotIn applies the NotIn predicate on the "token_type" field.
func TokenTypeNotIn(vs ...string) predicate.LINENFT {
	return predicate.LINENFT(sql.FieldNotIn(FieldTokenType, vs...))
}

// TokenTypeGT applies the GT predicate on the "token_type" field.
func TokenTypeGT(v string) predicate.LINENFT {
	return predicate.LINENFT(sql.FieldGT(FieldTokenType, v))
}

// TokenTypeGTE applies the GTE predicate on the "token_type" field.
func TokenTypeGTE(v string) predicate.LINENFT {
	return predicate.LINENFT(sql.FieldGTE(FieldTokenType, v))
}

// TokenTypeLT applies the LT predicate on the "token_type" field.
func TokenTypeLT(v string) predicate.LINENFT {
	return predicate.LINENFT(sql.FieldLT(FieldTokenType, v))
}

// TokenTypeLTE applies the LTE predicate on the "token_type" field.
func TokenTypeLTE(v string) predicate.LINENFT {
	return predicate.LINENFT(sql.FieldLTE(FieldTokenType, v))
}

// TokenTypeContains applies the Contains predicate on the "token_type" field.
func TokenTypeContains(v string) predicate.LINENFT {
	return predicate.LINENFT(sql.FieldContains(FieldTokenType, v))
}

// TokenTypeHasPrefix applies the HasPrefix predicate on the "token_type" field.
func TokenTypeHasPrefix(v string) predicate.LINENFT {
	return predicate.LINENFT(sql.FieldHasPrefix(FieldTokenType, v))
}

// TokenTypeHasSuffix applies the HasSuffix predicate on the "token_type" field.
func TokenTypeHasSuffix(v string) predicate.LINENFT {
	return predicate.LINENFT(sql.FieldHasSuffix(FieldTokenType, v))
}

// TokenTypeEqualFold applies the EqualFold predicate on the "token_type" field.
func TokenTypeEqualFold(v string) predicate.LINENFT {
	return predicate.LINENFT(sql.FieldEqualFold(FieldTokenType, v))
}

// TokenTypeContainsFold applies the ContainsFold predicate on the "token_type" field.
func TokenTypeContainsFold(v string) predicate.LINENFT {
	return predicate.LINENFT(sql.FieldContainsFold(FieldTokenType, v))
}

// TokenIndexEQ applies the EQ predicate on the "token_index" field.
func TokenIndexEQ(v string) predicate.LINENFT {
	return predicate.LINENFT(sql.FieldEQ(FieldTokenIndex, v))
}

// TokenIndexNEQ applies the NEQ predicate on the "token_index" field.
func TokenIndexNEQ(v string) predicate.LINENFT {
	return predicate.LINENFT(sql.FieldNEQ(FieldTokenIndex, v))
}

// TokenIndexIn applies the In predicate on the "token_index" field.
func TokenIndexIn(vs ...string) predicate.LINENFT {
	return predicate.LINENFT(sql.FieldIn(FieldTokenIndex, vs...))
}

// TokenIndexNotIn applies the NotIn predicate on the "token_index" field.
func TokenIndexNotIn(vs ...string) predicate.LINENFT {
	return predicate.LINENFT(sql.FieldNotIn(FieldTokenIndex, vs...))
}

// TokenIndexGT applies the GT predicate on the "token_index" field.
func TokenIndexGT(v string) predicate.LINENFT {
	return predicate.LINENFT(sql.FieldGT(FieldTokenIndex, v))
}

// TokenIndexGTE applies the GTE predicate on the "token_index" field.
func TokenIndexGTE(v string) predicate.LINENFT {
	return predicate.LINENFT(sql.FieldGTE(FieldTokenIndex, v))
}

// TokenIndexLT applies the LT predicate on the "token_index" field.
func TokenIndexLT(v string) predicate.LINENFT {
	return predicate.LINENFT(sql.FieldLT(FieldTokenIndex, v))
}

// TokenIndexLTE applies the LTE predicate on the "token_index" field.
func TokenIndexLTE(v string) predicate.LINENFT {
	return predicate.LINENFT(sql.FieldLTE(FieldTokenIndex, v))
}

// TokenIndexContains applies the Contains predicate on the "token_index" field.
func TokenIndexContains(v string) predicate.LINENFT {
	return predicate.LINENFT(sql.FieldContains(FieldTokenIndex, v))
}

// TokenIndexHasPrefix applies the HasPrefix predicate on the "token_index" field.
func TokenIndexHasPrefix(v string) predicate.LINENFT {
	return predicate.LINENFT(sql.FieldHasPrefix(FieldTokenIndex, v))
}

// TokenIndexHasSuffix applies the HasSuffix predicate on the "token_index" field.
func TokenIndexHasSuffix(v string) predicate.LINENFT {
	return predicate.LINENFT(sql.FieldHasSuffix(FieldTokenIndex, v))
}

// TokenIndexEqualFold applies the EqualFold predicate on the "token_index" field.
func TokenIndexEqualFold(v string) predicate.LINENFT {
	return predicate.LINENFT(sql.FieldEqualFold(FieldTokenIndex, v))
}

// TokenIndexContainsFold applies the ContainsFold predicate on the "token_index" field.
func TokenIndexContainsFold(v string) predicate.LINENFT {
	return predicate.LINENFT(sql.FieldContainsFold(FieldTokenIndex, v))
}

// TokenNameEQ applies the EQ predicate on the "token_name" field.
func TokenNameEQ(v string) predicate.LINENFT {
	return predicate.LINENFT(sql.FieldEQ(FieldTokenName, v))
}

// TokenNameNEQ applies the NEQ predicate on the "token_name" field.
func TokenNameNEQ(v string) predicate.LINENFT {
	return predicate.LINENFT(sql.FieldNEQ(FieldTokenName, v))
}

// TokenNameIn applies the In predicate on the "token_name" field.
func TokenNameIn(vs ...string) predicate.LINENFT {
	return predicate.LINENFT(sql.FieldIn(FieldTokenName, vs...))
}

// TokenNameNotIn applies the NotIn predicate on the "token_name" field.
func TokenNameNotIn(vs ...string) predicate.LINENFT {
	return predicate.LINENFT(sql.FieldNotIn(FieldTokenName, vs...))
}

// TokenNameGT applies the GT predicate on the "token_name" field.
func TokenNameGT(v string) predicate.LINENFT {
	return predicate.LINENFT(sql.FieldGT(FieldTokenName, v))
}

// TokenNameGTE applies the GTE predicate on the "token_name" field.
func TokenNameGTE(v string) predicate.LINENFT {
	return predicate.LINENFT(sql.FieldGTE(FieldTokenName, v))
}

// TokenNameLT applies the LT predicate on the "token_name" field.
func TokenNameLT(v string) predicate.LINENFT {
	return predicate.LINENFT(sql.FieldLT(FieldTokenName, v))
}

// TokenNameLTE applies the LTE predicate on the "token_name" field.
func TokenNameLTE(v string) predicate.LINENFT {
	return predicate.LINENFT(sql.FieldLTE(FieldTokenName, v))
}

// TokenNameContains applies the Contains predicate on the "token_name" field.
func TokenNameContains(v string) predicate.LINENFT {
	return predicate.LINENFT(sql.FieldContains(FieldTokenName, v))
}

// TokenNameHasPrefix applies the HasPrefix predicate on the "token_name" field.
func TokenNameHasPrefix(v string) predicate.LINENFT {
	return predicate.LINENFT(sql.FieldHasPrefix(FieldTokenName, v))
}

// TokenNameHasSuffix applies the HasSuffix predicate on the "token_name" field.
func TokenNameHasSuffix(v string) predicate.LINENFT {
	return predicate.LINENFT(sql.FieldHasSuffix(FieldTokenName, v))
}

// TokenNameEqualFold applies the EqualFold predicate on the "token_name" field.
func TokenNameEqualFold(v string) predicate.LINENFT {
	return predicate.LINENFT(sql.FieldEqualFold(FieldTokenName, v))
}

// TokenNameContainsFold applies the ContainsFold predicate on the "token_name" field.
func TokenNameContainsFold(v string) predicate.LINENFT {
	return predicate.LINENFT(sql.FieldContainsFold(FieldTokenName, v))
}

// TokenDescriptionEQ applies the EQ predicate on the "token_description" field.
func TokenDescriptionEQ(v string) predicate.LINENFT {
	return predicate.LINENFT(sql.FieldEQ(FieldTokenDescription, v))
}

// TokenDescriptionNEQ applies the NEQ predicate on the "token_description" field.
func TokenDescriptionNEQ(v string) predicate.LINENFT {
	return predicate.LINENFT(sql.FieldNEQ(FieldTokenDescription, v))
}

// TokenDescriptionIn applies the In predicate on the "token_description" field.
func TokenDescriptionIn(vs ...string) predicate.LINENFT {
	return predicate.LINENFT(sql.FieldIn(FieldTokenDescription, vs...))
}

// TokenDescriptionNotIn applies the NotIn predicate on the "token_description" field.
func TokenDescriptionNotIn(vs ...string) predicate.LINENFT {
	return predicate.LINENFT(sql.FieldNotIn(FieldTokenDescription, vs...))
}

// TokenDescriptionGT applies the GT predicate on the "token_description" field.
func TokenDescriptionGT(v string) predicate.LINENFT {
	return predicate.LINENFT(sql.FieldGT(FieldTokenDescription, v))
}

// TokenDescriptionGTE applies the GTE predicate on the "token_description" field.
func TokenDescriptionGTE(v string) predicate.LINENFT {
	return predicate.LINENFT(sql.FieldGTE(FieldTokenDescription, v))
}

// TokenDescriptionLT applies the LT predicate on the "token_description" field.
func TokenDescriptionLT(v string) predicate.LINENFT {
	return predicate.LINENFT(sql.FieldLT(FieldTokenDescription, v))
}

// TokenDescriptionLTE applies the LTE predicate on the "token_description" field.
func TokenDescriptionLTE(v string) predicate.LINENFT {
	return predicate.LINENFT(sql.FieldLTE(FieldTokenDescription, v))
}

// TokenDescriptionContains applies the Contains predicate on the "token_description" field.
func TokenDescriptionContains(v string) predicate.LINENFT {
	return predicate.LINENFT(sql.FieldContains(FieldTokenDescription, v))
}

// TokenDescriptionHasPrefix applies the HasPrefix predicate on the "token_description" field.
func TokenDescriptionHasPrefix(v string) predicate.LINENFT {
	return predicate.LINENFT(sql.FieldHasPrefix(FieldTokenDescription, v))
}

// TokenDescriptionHasSuffix applies the HasSuffix predicate on the "token_description" field.
func TokenDescriptionHasSuffix(v string) predicate.LINENFT {
	return predicate.LINENFT(sql.FieldHasSuffix(FieldTokenDescription, v))
}

// TokenDescriptionEqualFold applies the EqualFold predicate on the "token_description" field.
func TokenDescriptionEqualFold(v string) predicate.LINENFT {
	return predicate.LINENFT(sql.FieldEqualFold(FieldTokenDescription, v))
}

// TokenDescriptionContainsFold applies the ContainsFold predicate on the "token_description" field.
func TokenDescriptionContainsFold(v string) predicate.LINENFT {
	return predicate.LINENFT(sql.FieldContainsFold(FieldTokenDescription, v))
}

// TokenContentURLEQ applies the EQ predicate on the "token_content_url" field.
func TokenContentURLEQ(v string) predicate.LINENFT {
	return predicate.LINENFT(sql.FieldEQ(FieldTokenContentURL, v))
}

// TokenContentURLNEQ applies the NEQ predicate on the "token_content_url" field.
func TokenContentURLNEQ(v string) predicate.LINENFT {
	return predicate.LINENFT(sql.FieldNEQ(FieldTokenContentURL, v))
}

// TokenContentURLIn applies the In predicate on the "token_content_url" field.
func TokenContentURLIn(vs ...string) predicate.LINENFT {
	return predicate.LINENFT(sql.FieldIn(FieldTokenContentURL, vs...))
}

// TokenContentURLNotIn applies the NotIn predicate on the "token_content_url" field.
func TokenContentURLNotIn(vs ...string) predicate.LINENFT {
	return predicate.LINENFT(sql.FieldNotIn(FieldTokenContentURL, vs...))
}

// TokenContentURLGT applies the GT predicate on the "token_content_url" field.
func TokenContentURLGT(v string) predicate.LINENFT {
	return predicate.LINENFT(sql.FieldGT(FieldTokenContentURL, v))
}

// TokenContentURLGTE applies the GTE predicate on the "token_content_url" field.
func TokenContentURLGTE(v string) predicate.LINENFT {
	return predicate.LINENFT(sql.FieldGTE(FieldTokenContentURL, v))
}

// TokenContentURLLT applies the LT predicate on the "token_content_url" field.
func TokenContentURLLT(v string) predicate.LINENFT {
	return predicate.LINENFT(sql.FieldLT(FieldTokenContentURL, v))
}

// TokenContentURLLTE applies the LTE predicate on the "token_content_url" field.
func TokenContentURLLTE(v string) predicate.LINENFT {
	return predicate.LINENFT(sql.FieldLTE(FieldTokenContentURL, v))
}

// TokenContentURLContains applies the Contains predicate on the "token_content_url" field.
func TokenContentURLContains(v string) predicate.LINENFT {
	return predicate.LINENFT(sql.FieldContains(FieldTokenContentURL, v))
}

// TokenContentURLHasPrefix applies the HasPrefix predicate on the "token_content_url" field.
func TokenContentURLHasPrefix(v string) predicate.LINENFT {
	return predicate.LINENFT(sql.FieldHasPrefix(FieldTokenContentURL, v))
}

// TokenContentURLHasSuffix applies the HasSuffix predicate on the "token_content_url" field.
func TokenContentURLHasSuffix(v string) predicate.LINENFT {
	return predicate.LINENFT(sql.FieldHasSuffix(FieldTokenContentURL, v))
}

// TokenContentURLEqualFold applies the EqualFold predicate on the "token_content_url" field.
func TokenContentURLEqualFold(v string) predicate.LINENFT {
	return predicate.LINENFT(sql.FieldEqualFold(FieldTokenContentURL, v))
}

// TokenContentURLContainsFold applies the ContainsFold predicate on the "token_content_url" field.
func TokenContentURLContainsFold(v string) predicate.LINENFT {
	return predicate.LINENFT(sql.FieldContainsFold(FieldTokenContentURL, v))
}

// HasActivities applies the HasEdge predicate on the "activities" edge.
func HasActivities() predicate.LINENFT {
	return predicate.LINENFT(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ActivitiesTable, ActivitiesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasActivitiesWith applies the HasEdge predicate on the "activities" edge with a given conditions (other predicates).
func HasActivitiesWith(preds ...predicate.LINENFTActivity) predicate.LINENFT {
	return predicate.LINENFT(func(s *sql.Selector) {
		step := newActivitiesStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasMillionArthursProperty applies the HasEdge predicate on the "million_arthurs_property" edge.
func HasMillionArthursProperty() predicate.LINENFT {
	return predicate.LINENFT(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, MillionArthursPropertyTable, MillionArthursPropertyColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasMillionArthursPropertyWith applies the HasEdge predicate on the "million_arthurs_property" edge with a given conditions (other predicates).
func HasMillionArthursPropertyWith(preds ...predicate.LINENFTMillionArthursProperty) predicate.LINENFT {
	return predicate.LINENFT(func(s *sql.Selector) {
		step := newMillionArthursPropertyStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.LINENFT) predicate.LINENFT {
	return predicate.LINENFT(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.LINENFT) predicate.LINENFT {
	return predicate.LINENFT(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.LINENFT) predicate.LINENFT {
	return predicate.LINENFT(sql.NotPredicates(p))
}
