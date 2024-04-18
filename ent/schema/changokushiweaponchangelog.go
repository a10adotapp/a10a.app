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

// ChangokushiWeaponChangeLog holds the schema definition for the ChangokushiWeaponChangeLog entity.
type ChangokushiWeaponChangeLog struct {
	ent.Schema
}

func (ChangokushiWeaponChangeLog) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Table("changokushi_weapon_change_logs"),
		entsql.WithComments(true),
		schema.Comment("Changokushi weapon change log"),
	}
}

func (ChangokushiWeaponChangeLog) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
		SoftDeleteMixin{},
	}
}

// Fields of the ChangokushiWeaponChangeLog.
func (ChangokushiWeaponChangeLog) Fields() []ent.Field {
	return entfw.Fields(
		field.Uint32("id"),

		field.Uint32("changokushi_weapon_id"),

		field.String("status"),

		field.Int("price"),

		field.Time("published_at").
			SchemaType(map[string]string{
				dialect.MySQL: "datetime",
			}),
	)
}

// Edges of the ChangokushiWeaponChangeLog.
func (ChangokushiWeaponChangeLog) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("changokushi_weapon", ChangokushiWeapon.Type).
			Ref("change_logs").
			Field("changokushi_weapon_id").
			Unique().
			Required(),
	}
}
