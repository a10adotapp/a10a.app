package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// LINENFT holds the schema definition for the LINENFT entity.
type LINENFT struct {
	ent.Schema
}

func (LINENFT) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Table("line_nfts"),
		entsql.WithComments(true),
		schema.Comment("LINE NFT"),
	}
}

func (LINENFT) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
		SoftDeleteMixin{},
	}
}

// Fields of the LINENFT.
func (LINENFT) Fields() []ent.Field {
	return WrapFields(
		field.Uint32("id"),

		field.Uint32("line_nft_id"),

		field.String("contract_id"),

		field.String("token_type"),

		field.String("token_index"),

		field.String("token_name"),

		field.Text("token_description"),

		field.String("token_content_url"),
	)
}

// Edges of the LINENFT.
func (LINENFT) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("activities", LINENFTActivity.Type),

		edge.To("million_arthurs_properties", LINENFTMillionArthursProperty.Type).
			Unique(),
	}
}
