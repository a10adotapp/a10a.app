// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// ChangokushiWeaponsColumns holds the columns for the "changokushi_weapons" table.
	ChangokushiWeaponsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint32, Increment: true},
		{Name: "created_at", Type: field.TypeTime, SchemaType: map[string]string{"mysql": "datetime"}},
		{Name: "updated_at", Type: field.TypeTime, SchemaType: map[string]string{"mysql": "datetime"}},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true, SchemaType: map[string]string{"mysql": "datetime"}},
		{Name: "uri", Type: field.TypeString},
		{Name: "name", Type: field.TypeString},
		{Name: "rank", Type: field.TypeInt},
		{Name: "type", Type: field.TypeString},
		{Name: "skill", Type: field.TypeString},
		{Name: "vitality", Type: field.TypeInt},
		{Name: "strength", Type: field.TypeInt},
		{Name: "physical_defense", Type: field.TypeInt},
		{Name: "magical_defense", Type: field.TypeInt},
		{Name: "agility", Type: field.TypeInt},
	}
	// ChangokushiWeaponsTable holds the schema information for the "changokushi_weapons" table.
	ChangokushiWeaponsTable = &schema.Table{
		Name:       "changokushi_weapons",
		Comment:    "Changokushi weapons",
		Columns:    ChangokushiWeaponsColumns,
		PrimaryKey: []*schema.Column{ChangokushiWeaponsColumns[0]},
	}
	// ChangokushiWeaponChangeLogsColumns holds the columns for the "changokushi_weapon_change_logs" table.
	ChangokushiWeaponChangeLogsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint32, Increment: true},
		{Name: "created_at", Type: field.TypeTime, SchemaType: map[string]string{"mysql": "datetime"}},
		{Name: "updated_at", Type: field.TypeTime, SchemaType: map[string]string{"mysql": "datetime"}},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true, SchemaType: map[string]string{"mysql": "datetime"}},
		{Name: "status", Type: field.TypeString},
		{Name: "price", Type: field.TypeInt},
		{Name: "published_at", Type: field.TypeTime, SchemaType: map[string]string{"mysql": "datetime"}},
		{Name: "changokushi_weapon_id", Type: field.TypeUint32},
	}
	// ChangokushiWeaponChangeLogsTable holds the schema information for the "changokushi_weapon_change_logs" table.
	ChangokushiWeaponChangeLogsTable = &schema.Table{
		Name:       "changokushi_weapon_change_logs",
		Comment:    "Changokushi weapon change log",
		Columns:    ChangokushiWeaponChangeLogsColumns,
		PrimaryKey: []*schema.Column{ChangokushiWeaponChangeLogsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "changokushi_weapon_change_logs_changokushi_weapons_change_logs",
				Columns:    []*schema.Column{ChangokushiWeaponChangeLogsColumns[7]},
				RefColumns: []*schema.Column{ChangokushiWeaponsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// SanmeiHoikuenPostsColumns holds the columns for the "sanmei_hoikuen_posts" table.
	SanmeiHoikuenPostsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint32, Increment: true},
		{Name: "created_at", Type: field.TypeTime, SchemaType: map[string]string{"mysql": "datetime"}},
		{Name: "updated_at", Type: field.TypeTime, SchemaType: map[string]string{"mysql": "datetime"}},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true, SchemaType: map[string]string{"mysql": "datetime"}},
		{Name: "url", Type: field.TypeString, Size: 2147483647},
		{Name: "title", Type: field.TypeString},
		{Name: "date", Type: field.TypeString},
	}
	// SanmeiHoikuenPostsTable holds the schema information for the "sanmei_hoikuen_posts" table.
	SanmeiHoikuenPostsTable = &schema.Table{
		Name:       "sanmei_hoikuen_posts",
		Columns:    SanmeiHoikuenPostsColumns,
		PrimaryKey: []*schema.Column{SanmeiHoikuenPostsColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		ChangokushiWeaponsTable,
		ChangokushiWeaponChangeLogsTable,
		SanmeiHoikuenPostsTable,
	}
)

func init() {
	ChangokushiWeaponsTable.Annotation = &entsql.Annotation{
		Table: "changokushi_weapons",
	}
	ChangokushiWeaponChangeLogsTable.ForeignKeys[0].RefTable = ChangokushiWeaponsTable
	ChangokushiWeaponChangeLogsTable.Annotation = &entsql.Annotation{
		Table: "changokushi_weapon_change_logs",
	}
}
