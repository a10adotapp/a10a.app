package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"github.com/a10adotapp/entfw/v1"
)

type TimeMixin struct {
	mixin.Schema
}

const (
	FieldCreatedAt string = "created_at"
	FieldUpdatedAt string = "updated_at"
)

func (TimeMixin) Fields() []ent.Field {
	return entfw.Fields(
		field.Time(FieldCreatedAt).
			SchemaType(map[string]string{
				dialect.MySQL: "datetime",
			}).
			Immutable().
			Default(time.Now),

		field.Time(FieldUpdatedAt).
			SchemaType(map[string]string{
				dialect.MySQL: "datetime",
			}).
			Default(time.Now).
			UpdateDefault(time.Now),
	)
}
