package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/a10adotapp/entfw/v1"
)

// LINENFTMillionArthursProperty holds the schema definition for the LINENFTProperty entity.
type LINENFTMillionArthursProperty struct {
	ent.Schema
}

func (LINENFTMillionArthursProperty) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Table("line_nft_million_arthurs_properties"),
		entsql.WithComments(true),
		schema.Comment("LINE NFT 資産性ミリオンアーサーズのプロパティ"),
	}
}

func (LINENFTMillionArthursProperty) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
		SoftDeleteMixin{},
	}
}

// Fields of the LINENFTMillionArthursProperty.
func (LINENFTMillionArthursProperty) Fields() []ent.Field {
	return entfw.Fields(
		field.Uint32("id"),

		field.String("series").
			Optional().
			Nillable(),

		field.String("gear_category").
			Optional().
			Nillable(),

		field.String("gear_rarity").
			Optional().
			Nillable(),

		field.String("omj").
			Optional().
			Nillable(),
	)
}

// Edges of the LINENFTMillionArthursProperty.
func (LINENFTMillionArthursProperty) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("line_nft", LINENFT.Type).
			Ref("million_arthurs_property").
			Unique().
			Required(),
	}
}
