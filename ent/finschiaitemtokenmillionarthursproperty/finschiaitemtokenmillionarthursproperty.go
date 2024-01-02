// Code generated by ent, DO NOT EDIT.

package finschiaitemtokenmillionarthursproperty

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the finschiaitemtokenmillionarthursproperty type in the database.
	Label = "finschia_item_token_million_arthurs_property"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldDeletedAt holds the string denoting the deleted_at field in the database.
	FieldDeletedAt = "deleted_at"
	// FieldFinschiaItemTokenID holds the string denoting the finschia_item_token_id field in the database.
	FieldFinschiaItemTokenID = "finschia_item_token_id"
	// FieldSeries holds the string denoting the series field in the database.
	FieldSeries = "series"
	// FieldGearCategory holds the string denoting the gear_category field in the database.
	FieldGearCategory = "gear_category"
	// FieldGearRarity holds the string denoting the gear_rarity field in the database.
	FieldGearRarity = "gear_rarity"
	// EdgeFinschiaItemToken holds the string denoting the finschia_item_token edge name in mutations.
	EdgeFinschiaItemToken = "finschia_item_token"
	// Table holds the table name of the finschiaitemtokenmillionarthursproperty in the database.
	Table = "finschia_item_token_million_arthurs_properties"
	// FinschiaItemTokenTable is the table that holds the finschia_item_token relation/edge.
	FinschiaItemTokenTable = "finschia_item_token_million_arthurs_properties"
	// FinschiaItemTokenInverseTable is the table name for the FinschiaItemToken entity.
	// It exists in this package in order to avoid circular dependency with the "finschiaitemtoken" package.
	FinschiaItemTokenInverseTable = "finschia_item_tokens"
	// FinschiaItemTokenColumn is the table column denoting the finschia_item_token relation/edge.
	FinschiaItemTokenColumn = "finschia_item_token_id"
)

// Columns holds all SQL columns for finschiaitemtokenmillionarthursproperty fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldFinschiaItemTokenID,
	FieldSeries,
	FieldGearCategory,
	FieldGearRarity,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

// Note that the variables below are initialized by the runtime
// package on the initialization of the application. Therefore,
// it should be imported in the main as follows:
//
//	import _ "github.com/a10adotapp/a10a.app/ent/runtime"
var (
	Hooks        [1]ent.Hook
	Interceptors [1]ent.Interceptor
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
)

// OrderOption defines the ordering options for the FinschiaItemTokenMillionArthursProperty queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByDeletedAt orders the results by the deleted_at field.
func ByDeletedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDeletedAt, opts...).ToFunc()
}

// ByFinschiaItemTokenID orders the results by the finschia_item_token_id field.
func ByFinschiaItemTokenID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldFinschiaItemTokenID, opts...).ToFunc()
}

// BySeries orders the results by the series field.
func BySeries(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSeries, opts...).ToFunc()
}

// ByGearCategory orders the results by the gear_category field.
func ByGearCategory(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldGearCategory, opts...).ToFunc()
}

// ByGearRarity orders the results by the gear_rarity field.
func ByGearRarity(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldGearRarity, opts...).ToFunc()
}

// ByFinschiaItemTokenField orders the results by finschia_item_token field.
func ByFinschiaItemTokenField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newFinschiaItemTokenStep(), sql.OrderByField(field, opts...))
	}
}
func newFinschiaItemTokenStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(FinschiaItemTokenInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, FinschiaItemTokenTable, FinschiaItemTokenColumn),
	)
}
