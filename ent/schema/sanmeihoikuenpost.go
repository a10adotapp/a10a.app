package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"github.com/a10adotapp/entfw/v1"
)

type SanmeiHoikuenPost struct {
	ent.Schema
}

func (SanmeiHoikuenPost) Annotation() []schema.Annotation {
	return []schema.Annotation{
		entsql.Table("sanmei_hoikuen_posts"),
		entsql.WithComments(true),
		schema.Comment("Sanmei Hoikuen Posts"),
	}
}

func (SanmeiHoikuenPost) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
		SoftDeleteMixin{},
	}
}

func (SanmeiHoikuenPost) Fields() []ent.Field {
	return entfw.Fields(
		field.Uint32("id"),

		field.Text("url"),

		field.String("title"),

		field.String("date"),
	)
}

func (SanmeiHoikuenPost) Edges() []ent.Edge {
	return []ent.Edge{}
}
