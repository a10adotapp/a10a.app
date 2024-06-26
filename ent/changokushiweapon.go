// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/a10adotapp/a10a.app/ent/changokushiweapon"
)

// Changokushi weapons
type ChangokushiWeapon struct {
	config `json:"-"`
	// ID of the ent.
	ID uint32 `json:"id"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt *time.Time `json:"deleted_at"`
	// URI holds the value of the "uri" field.
	URI string `json:"uri"`
	// Name holds the value of the "name" field.
	Name string `json:"name"`
	// Rank holds the value of the "rank" field.
	Rank int `json:"rank"`
	// Type holds the value of the "type" field.
	Type string `json:"type"`
	// Skill holds the value of the "skill" field.
	Skill string `json:"skill"`
	// Vitality holds the value of the "vitality" field.
	Vitality int `json:"vitality"`
	// Strength holds the value of the "strength" field.
	Strength int `json:"strength"`
	// PhysicalDefense holds the value of the "physical_defense" field.
	PhysicalDefense int `json:"physical_defense"`
	// MagicalDefense holds the value of the "magical_defense" field.
	MagicalDefense int `json:"magical_defense"`
	// Agility holds the value of the "agility" field.
	Agility int `json:"agility"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ChangokushiWeaponQuery when eager-loading is set.
	Edges        ChangokushiWeaponEdges `json:"edges"`
	selectValues sql.SelectValues
}

// ChangokushiWeaponEdges holds the relations/edges for other nodes in the graph.
type ChangokushiWeaponEdges struct {
	// ChangeLogs holds the value of the change_logs edge.
	ChangeLogs []*ChangokushiWeaponChangeLog `json:"change_logs,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// ChangeLogsOrErr returns the ChangeLogs value or an error if the edge
// was not loaded in eager-loading.
func (e ChangokushiWeaponEdges) ChangeLogsOrErr() ([]*ChangokushiWeaponChangeLog, error) {
	if e.loadedTypes[0] {
		return e.ChangeLogs, nil
	}
	return nil, &NotLoadedError{edge: "change_logs"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*ChangokushiWeapon) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case changokushiweapon.FieldID, changokushiweapon.FieldRank, changokushiweapon.FieldVitality, changokushiweapon.FieldStrength, changokushiweapon.FieldPhysicalDefense, changokushiweapon.FieldMagicalDefense, changokushiweapon.FieldAgility:
			values[i] = new(sql.NullInt64)
		case changokushiweapon.FieldURI, changokushiweapon.FieldName, changokushiweapon.FieldType, changokushiweapon.FieldSkill:
			values[i] = new(sql.NullString)
		case changokushiweapon.FieldCreatedAt, changokushiweapon.FieldUpdatedAt, changokushiweapon.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the ChangokushiWeapon fields.
func (cw *ChangokushiWeapon) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case changokushiweapon.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			cw.ID = uint32(value.Int64)
		case changokushiweapon.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				cw.CreatedAt = value.Time
			}
		case changokushiweapon.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				cw.UpdatedAt = value.Time
			}
		case changokushiweapon.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				cw.DeletedAt = new(time.Time)
				*cw.DeletedAt = value.Time
			}
		case changokushiweapon.FieldURI:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field uri", values[i])
			} else if value.Valid {
				cw.URI = value.String
			}
		case changokushiweapon.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				cw.Name = value.String
			}
		case changokushiweapon.FieldRank:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field rank", values[i])
			} else if value.Valid {
				cw.Rank = int(value.Int64)
			}
		case changokushiweapon.FieldType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				cw.Type = value.String
			}
		case changokushiweapon.FieldSkill:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field skill", values[i])
			} else if value.Valid {
				cw.Skill = value.String
			}
		case changokushiweapon.FieldVitality:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field vitality", values[i])
			} else if value.Valid {
				cw.Vitality = int(value.Int64)
			}
		case changokushiweapon.FieldStrength:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field strength", values[i])
			} else if value.Valid {
				cw.Strength = int(value.Int64)
			}
		case changokushiweapon.FieldPhysicalDefense:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field physical_defense", values[i])
			} else if value.Valid {
				cw.PhysicalDefense = int(value.Int64)
			}
		case changokushiweapon.FieldMagicalDefense:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field magical_defense", values[i])
			} else if value.Valid {
				cw.MagicalDefense = int(value.Int64)
			}
		case changokushiweapon.FieldAgility:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field agility", values[i])
			} else if value.Valid {
				cw.Agility = int(value.Int64)
			}
		default:
			cw.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the ChangokushiWeapon.
// This includes values selected through modifiers, order, etc.
func (cw *ChangokushiWeapon) Value(name string) (ent.Value, error) {
	return cw.selectValues.Get(name)
}

// QueryChangeLogs queries the "change_logs" edge of the ChangokushiWeapon entity.
func (cw *ChangokushiWeapon) QueryChangeLogs() *ChangokushiWeaponChangeLogQuery {
	return NewChangokushiWeaponClient(cw.config).QueryChangeLogs(cw)
}

// Update returns a builder for updating this ChangokushiWeapon.
// Note that you need to call ChangokushiWeapon.Unwrap() before calling this method if this ChangokushiWeapon
// was returned from a transaction, and the transaction was committed or rolled back.
func (cw *ChangokushiWeapon) Update() *ChangokushiWeaponUpdateOne {
	return NewChangokushiWeaponClient(cw.config).UpdateOne(cw)
}

// Unwrap unwraps the ChangokushiWeapon entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (cw *ChangokushiWeapon) Unwrap() *ChangokushiWeapon {
	_tx, ok := cw.config.driver.(*txDriver)
	if !ok {
		panic("ent: ChangokushiWeapon is not a transactional entity")
	}
	cw.config.driver = _tx.drv
	return cw
}

// String implements the fmt.Stringer.
func (cw *ChangokushiWeapon) String() string {
	var builder strings.Builder
	builder.WriteString("ChangokushiWeapon(")
	builder.WriteString(fmt.Sprintf("id=%v, ", cw.ID))
	builder.WriteString("created_at=")
	builder.WriteString(cw.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(cw.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	if v := cw.DeletedAt; v != nil {
		builder.WriteString("deleted_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	builder.WriteString("uri=")
	builder.WriteString(cw.URI)
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(cw.Name)
	builder.WriteString(", ")
	builder.WriteString("rank=")
	builder.WriteString(fmt.Sprintf("%v", cw.Rank))
	builder.WriteString(", ")
	builder.WriteString("type=")
	builder.WriteString(cw.Type)
	builder.WriteString(", ")
	builder.WriteString("skill=")
	builder.WriteString(cw.Skill)
	builder.WriteString(", ")
	builder.WriteString("vitality=")
	builder.WriteString(fmt.Sprintf("%v", cw.Vitality))
	builder.WriteString(", ")
	builder.WriteString("strength=")
	builder.WriteString(fmt.Sprintf("%v", cw.Strength))
	builder.WriteString(", ")
	builder.WriteString("physical_defense=")
	builder.WriteString(fmt.Sprintf("%v", cw.PhysicalDefense))
	builder.WriteString(", ")
	builder.WriteString("magical_defense=")
	builder.WriteString(fmt.Sprintf("%v", cw.MagicalDefense))
	builder.WriteString(", ")
	builder.WriteString("agility=")
	builder.WriteString(fmt.Sprintf("%v", cw.Agility))
	builder.WriteByte(')')
	return builder.String()
}

// ChangokushiWeapons is a parsable slice of ChangokushiWeapon.
type ChangokushiWeapons []*ChangokushiWeapon
