package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// FinschiaItemToken holds the schema definition for the FinschiaItemToken entity.
type FinschiaItemToken struct {
	ent.Schema
}

func (FinschiaItemToken) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Table("finschia_item_tokens"),
		entsql.WithComments(true),
		schema.Comment("Finschia Item Tokens"),
	}
}

func (FinschiaItemToken) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
		SoftDeleteMixin{},
	}
}

// Fields of the FinschiaItemToken.
func (FinschiaItemToken) Fields() []ent.Field {
	return WrapFields(
		field.Uint32("id"),

		field.String("contract_id"),

		field.String("token_type"),

		field.String("name"),
	)
}

// Edges of the FinschiaItemToken.
func (FinschiaItemToken) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("activities", FinschiaItemTokenActivity.Type),
		edge.To("million_arthurs_properties", FinschiaItemTokenMillionArthursProperty.Type),
	}
}
