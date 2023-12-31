// Code generated by ent, DO NOT EDIT.

package linenftmillionarthursproperty

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the linenftmillionarthursproperty type in the database.
	Label = "linenft_million_arthurs_property"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldDeletedAt holds the string denoting the deleted_at field in the database.
	FieldDeletedAt = "deleted_at"
	// FieldSeries holds the string denoting the series field in the database.
	FieldSeries = "series"
	// FieldGearCategory holds the string denoting the gear_category field in the database.
	FieldGearCategory = "gear_category"
	// FieldGearRarity holds the string denoting the gear_rarity field in the database.
	FieldGearRarity = "gear_rarity"
	// FieldOmj holds the string denoting the omj field in the database.
	FieldOmj = "omj"
	// EdgeLineNft holds the string denoting the line_nft edge name in mutations.
	EdgeLineNft = "line_nft"
	// Table holds the table name of the linenftmillionarthursproperty in the database.
	Table = "line_nft_million_arthurs_properties"
	// LineNftTable is the table that holds the line_nft relation/edge.
	LineNftTable = "line_nft_million_arthurs_properties"
	// LineNftInverseTable is the table name for the LINENFT entity.
	// It exists in this package in order to avoid circular dependency with the "linenft" package.
	LineNftInverseTable = "line_nfts"
	// LineNftColumn is the table column denoting the line_nft relation/edge.
	LineNftColumn = "linenft_million_arthurs_property"
)

// Columns holds all SQL columns for linenftmillionarthursproperty fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldSeries,
	FieldGearCategory,
	FieldGearRarity,
	FieldOmj,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "line_nft_million_arthurs_properties"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"linenft_million_arthurs_property",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
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

// OrderOption defines the ordering options for the LINENFTMillionArthursProperty queries.
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

// ByOmj orders the results by the omj field.
func ByOmj(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldOmj, opts...).ToFunc()
}

// ByLineNftField orders the results by line_nft field.
func ByLineNftField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newLineNftStep(), sql.OrderByField(field, opts...))
	}
}
func newLineNftStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(LineNftInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2O, true, LineNftTable, LineNftColumn),
	)
}
