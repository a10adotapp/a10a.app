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

// LINENFTActivity holds the schema definition for the LINENFTProperty entity.
type LINENFTActivity struct {
	ent.Schema
}

func (LINENFTActivity) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Table("line_nft_activities"),
		entsql.WithComments(true),
		schema.Comment("LINE NFT 売買履歴"),
	}
}

func (LINENFTActivity) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
		SoftDeleteMixin{},
	}
}

// Fields of the LINENFTActivity.
func (LINENFTActivity) Fields() []ent.Field {
	return entfw.Fields(
		field.Uint32("id"),

		field.String("activity_type"),

		field.Uint32("sale_id"),

		field.String("sale_type"),

		field.String("currency_type"),

		field.Float32("price"),

		field.Time("activated_at").
			SchemaType(map[string]string{
				dialect.MySQL: "datetime",
			}),
	)
}

// Edges of the LINENFTActivity.
func (LINENFTActivity) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("line_nft", LINENFT.Type).
			Ref("activities").
			Unique().
			Required(),
	}
}
