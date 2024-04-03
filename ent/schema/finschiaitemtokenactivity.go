package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/a10adotapp/entfw/v1"
)

// FinschiaItemTokenActivity holds the schema definition for the FinschiaItemTokenActivity entity.
type FinschiaItemTokenActivity struct {
	ent.Schema
}

func (FinschiaItemTokenActivity) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Table("finschia_item_token_activities"),
		entsql.WithComments(true),
		schema.Comment("Activity Records for Finschia Item Tokens"),
	}
}

func (FinschiaItemTokenActivity) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
		SoftDeleteMixin{},
	}
}

// Fields of the FinschiaItemTokenActivity.
func (FinschiaItemTokenActivity) Fields() []ent.Field {
	return entfw.Fields(
		field.Uint32("id"),

		field.String("transaction_hash"),

		field.Uint32("finschia_item_token_id"),

		field.String("activity_type"),

		field.Time("activated_at").
			SchemaType(map[string]string{
				dialect.MySQL: "datetime",
			}),
	)
}

// Edges of the FinschiaItemTokenActivity.
func (FinschiaItemTokenActivity) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("finschia_item_token", FinschiaItemToken.Type).
			Ref("activities").
			Field("finschia_item_token_id").
			Unique().
			Required(),
	}
}
