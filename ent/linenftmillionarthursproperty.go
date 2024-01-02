// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/a10adotapp/a10a.app/ent/linenft"
	"github.com/a10adotapp/a10a.app/ent/linenftmillionarthursproperty"
)

// LINE NFT 資産性ミリオンアーサーズのプロパティ
type LINENFTMillionArthursProperty struct {
	config `json:"-"`
	// ID of the ent.
	ID uint32 `json:"id"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt *time.Time `json:"deleted_at"`
	// Series holds the value of the "series" field.
	Series *string `json:"series"`
	// GearCategory holds the value of the "gear_category" field.
	GearCategory *string `json:"gear_category"`
	// GearRarity holds the value of the "gear_rarity" field.
	GearRarity *string `json:"gear_rarity"`
	// Omj holds the value of the "omj" field.
	Omj *string `json:"omj"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the LINENFTMillionArthursPropertyQuery when eager-loading is set.
	Edges                            LINENFTMillionArthursPropertyEdges `json:"edges"`
	linenft_million_arthurs_property *uint32
	selectValues                     sql.SelectValues
}

// LINENFTMillionArthursPropertyEdges holds the relations/edges for other nodes in the graph.
type LINENFTMillionArthursPropertyEdges struct {
	// LineNft holds the value of the line_nft edge.
	LineNft *LINENFT `json:"line_nft,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// LineNftOrErr returns the LineNft value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e LINENFTMillionArthursPropertyEdges) LineNftOrErr() (*LINENFT, error) {
	if e.loadedTypes[0] {
		if e.LineNft == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: linenft.Label}
		}
		return e.LineNft, nil
	}
	return nil, &NotLoadedError{edge: "line_nft"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*LINENFTMillionArthursProperty) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case linenftmillionarthursproperty.FieldID:
			values[i] = new(sql.NullInt64)
		case linenftmillionarthursproperty.FieldSeries, linenftmillionarthursproperty.FieldGearCategory, linenftmillionarthursproperty.FieldGearRarity, linenftmillionarthursproperty.FieldOmj:
			values[i] = new(sql.NullString)
		case linenftmillionarthursproperty.FieldCreatedAt, linenftmillionarthursproperty.FieldUpdatedAt, linenftmillionarthursproperty.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		case linenftmillionarthursproperty.ForeignKeys[0]: // linenft_million_arthurs_property
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the LINENFTMillionArthursProperty fields.
func (lmap *LINENFTMillionArthursProperty) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case linenftmillionarthursproperty.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			lmap.ID = uint32(value.Int64)
		case linenftmillionarthursproperty.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				lmap.CreatedAt = value.Time
			}
		case linenftmillionarthursproperty.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				lmap.UpdatedAt = value.Time
			}
		case linenftmillionarthursproperty.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				lmap.DeletedAt = new(time.Time)
				*lmap.DeletedAt = value.Time
			}
		case linenftmillionarthursproperty.FieldSeries:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field series", values[i])
			} else if value.Valid {
				lmap.Series = new(string)
				*lmap.Series = value.String
			}
		case linenftmillionarthursproperty.FieldGearCategory:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field gear_category", values[i])
			} else if value.Valid {
				lmap.GearCategory = new(string)
				*lmap.GearCategory = value.String
			}
		case linenftmillionarthursproperty.FieldGearRarity:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field gear_rarity", values[i])
			} else if value.Valid {
				lmap.GearRarity = new(string)
				*lmap.GearRarity = value.String
			}
		case linenftmillionarthursproperty.FieldOmj:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field omj", values[i])
			} else if value.Valid {
				lmap.Omj = new(string)
				*lmap.Omj = value.String
			}
		case linenftmillionarthursproperty.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field linenft_million_arthurs_property", value)
			} else if value.Valid {
				lmap.linenft_million_arthurs_property = new(uint32)
				*lmap.linenft_million_arthurs_property = uint32(value.Int64)
			}
		default:
			lmap.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the LINENFTMillionArthursProperty.
// This includes values selected through modifiers, order, etc.
func (lmap *LINENFTMillionArthursProperty) Value(name string) (ent.Value, error) {
	return lmap.selectValues.Get(name)
}

// QueryLineNft queries the "line_nft" edge of the LINENFTMillionArthursProperty entity.
func (lmap *LINENFTMillionArthursProperty) QueryLineNft() *LINENFTQuery {
	return NewLINENFTMillionArthursPropertyClient(lmap.config).QueryLineNft(lmap)
}

// Update returns a builder for updating this LINENFTMillionArthursProperty.
// Note that you need to call LINENFTMillionArthursProperty.Unwrap() before calling this method if this LINENFTMillionArthursProperty
// was returned from a transaction, and the transaction was committed or rolled back.
func (lmap *LINENFTMillionArthursProperty) Update() *LINENFTMillionArthursPropertyUpdateOne {
	return NewLINENFTMillionArthursPropertyClient(lmap.config).UpdateOne(lmap)
}

// Unwrap unwraps the LINENFTMillionArthursProperty entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (lmap *LINENFTMillionArthursProperty) Unwrap() *LINENFTMillionArthursProperty {
	_tx, ok := lmap.config.driver.(*txDriver)
	if !ok {
		panic("ent: LINENFTMillionArthursProperty is not a transactional entity")
	}
	lmap.config.driver = _tx.drv
	return lmap
}

// String implements the fmt.Stringer.
func (lmap *LINENFTMillionArthursProperty) String() string {
	var builder strings.Builder
	builder.WriteString("LINENFTMillionArthursProperty(")
	builder.WriteString(fmt.Sprintf("id=%v, ", lmap.ID))
	builder.WriteString("created_at=")
	builder.WriteString(lmap.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(lmap.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	if v := lmap.DeletedAt; v != nil {
		builder.WriteString("deleted_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	if v := lmap.Series; v != nil {
		builder.WriteString("series=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := lmap.GearCategory; v != nil {
		builder.WriteString("gear_category=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := lmap.GearRarity; v != nil {
		builder.WriteString("gear_rarity=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := lmap.Omj; v != nil {
		builder.WriteString("omj=")
		builder.WriteString(*v)
	}
	builder.WriteByte(')')
	return builder.String()
}

// LINENFTMillionArthursProperties is a parsable slice of LINENFTMillionArthursProperty.
type LINENFTMillionArthursProperties []*LINENFTMillionArthursProperty
