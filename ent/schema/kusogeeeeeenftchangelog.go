package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/a10adotapp/entfw/v1"
)

// KusogeeeeeeNFTChangeLog holds the schema definition for the KusogeeeeeeNFTChangeLog entity.
type KusogeeeeeeNFTChangeLog struct {
	ent.Schema
}

func (KusogeeeeeeNFTChangeLog) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Table("kusogeeeeee_nft_change_logs"),
		entsql.WithComments(true),
		schema.Comment("Kusogeeeeee NFT Change Log"),
	}
}

func (KusogeeeeeeNFTChangeLog) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
		SoftDeleteMixin{},
	}
}

// Fields of the KusogeeeeeeNFTChangeLog.
func (KusogeeeeeeNFTChangeLog) Fields() []ent.Field {
	return entfw.Fields(
		field.Uint32("id"),

		field.Uint32("kusogeeeeee_nft_id"),

		field.String("status"),

		field.Int("price"),
	)
}

// Edges of the KusogeeeeeeNFTChangeLog.
func (KusogeeeeeeNFTChangeLog) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("kusogeeeeee_nft", KusogeeeeeeNFT.Type).
			Ref("change_logs").
			Field("kusogeeeeee_nft_id").
			Unique().
			Required(),
	}
}
