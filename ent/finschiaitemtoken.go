// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/a10adotapp/a10a.app/ent/finschiaitemtoken"
)

// Finschia Item Tokens
type FinschiaItemToken struct {
	config `json:"-"`
	// ID of the ent.
	ID uint32 `json:"id"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt *time.Time `json:"deleted_at"`
	// ContractID holds the value of the "contract_id" field.
	ContractID string `json:"contract_id"`
	// TokenType holds the value of the "token_type" field.
	TokenType string `json:"token_type"`
	// Name holds the value of the "name" field.
	Name string `json:"name"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the FinschiaItemTokenQuery when eager-loading is set.
	Edges        FinschiaItemTokenEdges `json:"edges"`
	selectValues sql.SelectValues
}

// FinschiaItemTokenEdges holds the relations/edges for other nodes in the graph.
type FinschiaItemTokenEdges struct {
	// Activities holds the value of the activities edge.
	Activities []*FinschiaItemTokenActivity `json:"activities,omitempty"`
	// MillionArthursProperties holds the value of the million_arthurs_properties edge.
	MillionArthursProperties []*FinschiaItemTokenMillionArthursProperty `json:"million_arthurs_properties,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// ActivitiesOrErr returns the Activities value or an error if the edge
// was not loaded in eager-loading.
func (e FinschiaItemTokenEdges) ActivitiesOrErr() ([]*FinschiaItemTokenActivity, error) {
	if e.loadedTypes[0] {
		return e.Activities, nil
	}
	return nil, &NotLoadedError{edge: "activities"}
}

// MillionArthursPropertiesOrErr returns the MillionArthursProperties value or an error if the edge
// was not loaded in eager-loading.
func (e FinschiaItemTokenEdges) MillionArthursPropertiesOrErr() ([]*FinschiaItemTokenMillionArthursProperty, error) {
	if e.loadedTypes[1] {
		return e.MillionArthursProperties, nil
	}
	return nil, &NotLoadedError{edge: "million_arthurs_properties"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*FinschiaItemToken) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case finschiaitemtoken.FieldID:
			values[i] = new(sql.NullInt64)
		case finschiaitemtoken.FieldContractID, finschiaitemtoken.FieldTokenType, finschiaitemtoken.FieldName:
			values[i] = new(sql.NullString)
		case finschiaitemtoken.FieldCreatedAt, finschiaitemtoken.FieldUpdatedAt, finschiaitemtoken.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the FinschiaItemToken fields.
func (fit *FinschiaItemToken) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case finschiaitemtoken.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			fit.ID = uint32(value.Int64)
		case finschiaitemtoken.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				fit.CreatedAt = value.Time
			}
		case finschiaitemtoken.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				fit.UpdatedAt = value.Time
			}
		case finschiaitemtoken.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				fit.DeletedAt = new(time.Time)
				*fit.DeletedAt = value.Time
			}
		case finschiaitemtoken.FieldContractID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field contract_id", values[i])
			} else if value.Valid {
				fit.ContractID = value.String
			}
		case finschiaitemtoken.FieldTokenType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field token_type", values[i])
			} else if value.Valid {
				fit.TokenType = value.String
			}
		case finschiaitemtoken.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				fit.Name = value.String
			}
		default:
			fit.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the FinschiaItemToken.
// This includes values selected through modifiers, order, etc.
func (fit *FinschiaItemToken) Value(name string) (ent.Value, error) {
	return fit.selectValues.Get(name)
}

// QueryActivities queries the "activities" edge of the FinschiaItemToken entity.
func (fit *FinschiaItemToken) QueryActivities() *FinschiaItemTokenActivityQuery {
	return NewFinschiaItemTokenClient(fit.config).QueryActivities(fit)
}

// QueryMillionArthursProperties queries the "million_arthurs_properties" edge of the FinschiaItemToken entity.
func (fit *FinschiaItemToken) QueryMillionArthursProperties() *FinschiaItemTokenMillionArthursPropertyQuery {
	return NewFinschiaItemTokenClient(fit.config).QueryMillionArthursProperties(fit)
}

// Update returns a builder for updating this FinschiaItemToken.
// Note that you need to call FinschiaItemToken.Unwrap() before calling this method if this FinschiaItemToken
// was returned from a transaction, and the transaction was committed or rolled back.
func (fit *FinschiaItemToken) Update() *FinschiaItemTokenUpdateOne {
	return NewFinschiaItemTokenClient(fit.config).UpdateOne(fit)
}

// Unwrap unwraps the FinschiaItemToken entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (fit *FinschiaItemToken) Unwrap() *FinschiaItemToken {
	_tx, ok := fit.config.driver.(*txDriver)
	if !ok {
		panic("ent: FinschiaItemToken is not a transactional entity")
	}
	fit.config.driver = _tx.drv
	return fit
}

// String implements the fmt.Stringer.
func (fit *FinschiaItemToken) String() string {
	var builder strings.Builder
	builder.WriteString("FinschiaItemToken(")
	builder.WriteString(fmt.Sprintf("id=%v, ", fit.ID))
	builder.WriteString("created_at=")
	builder.WriteString(fit.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(fit.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	if v := fit.DeletedAt; v != nil {
		builder.WriteString("deleted_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	builder.WriteString("contract_id=")
	builder.WriteString(fit.ContractID)
	builder.WriteString(", ")
	builder.WriteString("token_type=")
	builder.WriteString(fit.TokenType)
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(fit.Name)
	builder.WriteByte(')')
	return builder.String()
}

// FinschiaItemTokens is a parsable slice of FinschiaItemToken.
type FinschiaItemTokens []*FinschiaItemToken
