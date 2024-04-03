package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/a10adotapp/entfw/v1"
)

// FinschiaItemTokenMillionArthursProperty holds the schema definition for the FinschiaItemTokenMillionArthursProperty entity.
type FinschiaItemTokenMillionArthursProperty struct {
	ent.Schema
}

func (FinschiaItemTokenMillionArthursProperty) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Table("finschia_item_token_million_arthurs_properties"),
		entsql.WithComments(true),
		schema.Comment("Million Arthurs' Properties for Finschia Item Tokens"),
	}
}

func (FinschiaItemTokenMillionArthursProperty) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
		SoftDeleteMixin{},
	}
}

// Fields of the FinschiaItemTokenMillionArthursProperty.
func (FinschiaItemTokenMillionArthursProperty) Fields() []ent.Field {
	return entfw.Fields(
		field.Uint32("id"),

		field.Uint32("finschia_item_token_id"),

		field.String("series").
			Optional().
			Nillable(),

		field.String("gear_category").
			Optional().
			Nillable(),

		field.String("gear_rarity").
			Optional().
			Nillable(),
	)
}

// Edges of the FinschiaItemTokenMillionArthursProperty.
func (FinschiaItemTokenMillionArthursProperty) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("finschia_item_token", FinschiaItemToken.Type).
			Ref("million_arthurs_properties").
			Field("finschia_item_token_id").
			Unique().
			Required(),
	}
}
