package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/a10adotapp/entfw/v1"
)

// ChangokushiWeapon holds the schema definition for the ChangokushiWeapon entity.
type ChangokushiWeapon struct {
	ent.Schema
}

func (ChangokushiWeapon) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Table("changokushi_weapons"),
		entsql.WithComments(true),
		schema.Comment("Changokushi weapons"),
	}
}

func (ChangokushiWeapon) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
		SoftDeleteMixin{},
	}
}

// Fields of the ChangokushiWeapon.
func (ChangokushiWeapon) Fields() []ent.Field {
	return entfw.Fields(
		field.Uint32("id"),

		field.String("uri"),

		field.String("name"),

		field.Int("rank"),

		field.String("type"),

		field.String("skill"),

		field.Int("vitality"),

		field.Int("strength"),

		field.Int("physical_defense"),

		field.Int("magical_defense"),

		field.Int("agility"),
	)
}

// Edges of the ChangokushiWeapon.
func (ChangokushiWeapon) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("change_logs", ChangokushiWeaponChangeLog.Type),
	}
}
