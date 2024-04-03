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

// KusogeeeeeeNFT holds the schema definition for the KusogeeeeeeNFT entity.
type KusogeeeeeeNFT struct {
	ent.Schema
}

func (KusogeeeeeeNFT) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Table("kusogeeeeee_nfts"),
		entsql.WithComments(true),
		schema.Comment("Kusogeeeeee NFT"),
	}
}

func (KusogeeeeeeNFT) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
		SoftDeleteMixin{},
	}
}

// Fields of the KusogeeeeeeNFT.
func (KusogeeeeeeNFT) Fields() []ent.Field {
	return entfw.Fields(
		field.Uint32("id"),

		field.String("uri"),

		field.String("type"),

		field.String("name"),

		field.String("status"),

		field.Int("price"),

		field.Time("published_at").
			SchemaType(map[string]string{
				dialect.MySQL: "datetime",
			}),

		field.Int("weapon_rank").
			Optional().
			Nillable(),

		field.String("weapon_type").
			Optional().
			Nillable(),

		field.Int("weapon_vitality").
			Optional().
			Nillable(),

		field.Int("weapon_strength").
			Optional().
			Nillable(),

		field.Int("weapon_physical_defense").
			Optional().
			Nillable(),

		field.Int("weapon_magical_defense").
			Optional().
			Nillable(),

		field.Int("weapon_agility").
			Optional().
			Nillable(),

		field.String("character_rank").
			Optional().
			Nillable(),

		field.Int("character_total_supply").
			Optional().
			Nillable(),
	)
}

// Edges of the KusogeeeeeeNFT.
func (KusogeeeeeeNFT) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("change_logs", KusogeeeeeeNFTChangeLog.Type),
	}
}
