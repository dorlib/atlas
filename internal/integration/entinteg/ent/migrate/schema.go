// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// ActivitiesColumns holds the columns for the "activities" table.
	ActivitiesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
	}
	// ActivitiesTable holds the schema information for the "activities" table.
	ActivitiesTable = &schema.Table{
		Name:        "activities",
		Columns:     ActivitiesColumns,
		PrimaryKey:  []*schema.Column{ActivitiesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// DefaultContainersColumns holds the columns for the "default_containers" table.
	DefaultContainersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "stringdef", Type: field.TypeString, Default: "default"},
		{Name: "int", Type: field.TypeInt, Default: 1},
		{Name: "bool", Type: field.TypeBool, Default: true},
		{Name: "enum", Type: field.TypeEnum, Enums: []string{"1", "2"}, Default: "1"},
		{Name: "float", Type: field.TypeFloat64, Default: 1.5},
	}
	// DefaultContainersTable holds the schema information for the "default_containers" table.
	DefaultContainersTable = &schema.Table{
		Name:        "default_containers",
		Columns:     DefaultContainersColumns,
		PrimaryKey:  []*schema.Column{DefaultContainersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// GroupsColumns holds the columns for the "groups" table.
	GroupsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
	}
	// GroupsTable holds the schema information for the "groups" table.
	GroupsTable = &schema.Table{
		Name:        "groups",
		Columns:     GroupsColumns,
		PrimaryKey:  []*schema.Column{GroupsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "optional", Type: field.TypeString, Nullable: true},
		{Name: "int", Type: field.TypeInt},
		{Name: "uint", Type: field.TypeUint},
		{Name: "time", Type: field.TypeTime},
		{Name: "bool", Type: field.TypeBool},
		{Name: "enum", Type: field.TypeEnum, Enums: []string{"1", "2", "3"}},
		{Name: "named_enum", Type: field.TypeEnum, Enums: []string{"1", "2", "3"}},
		{Name: "uuid", Type: field.TypeUUID, Unique: true},
		{Name: "bytes", Type: field.TypeBytes},
		{Name: "group_id", Type: field.TypeInt, Nullable: true},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "users_groups_group",
				Columns:    []*schema.Column{UsersColumns[11]},
				RefColumns: []*schema.Column{GroupsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "user_time",
				Unique:  false,
				Columns: []*schema.Column{UsersColumns[5]},
			},
		},
	}
	// UserActivitiesColumns holds the columns for the "user_activities" table.
	UserActivitiesColumns = []*schema.Column{
		{Name: "user_id", Type: field.TypeInt},
		{Name: "activity_id", Type: field.TypeInt},
	}
	// UserActivitiesTable holds the schema information for the "user_activities" table.
	UserActivitiesTable = &schema.Table{
		Name:       "user_activities",
		Columns:    UserActivitiesColumns,
		PrimaryKey: []*schema.Column{UserActivitiesColumns[0], UserActivitiesColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "user_activities_user_id",
				Columns:    []*schema.Column{UserActivitiesColumns[0]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "user_activities_activity_id",
				Columns:    []*schema.Column{UserActivitiesColumns[1]},
				RefColumns: []*schema.Column{ActivitiesColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		ActivitiesTable,
		DefaultContainersTable,
		GroupsTable,
		UsersTable,
		UserActivitiesTable,
	}
)

func init() {
	UsersTable.ForeignKeys[0].RefTable = GroupsTable
	UserActivitiesTable.ForeignKeys[0].RefTable = UsersTable
	UserActivitiesTable.ForeignKeys[1].RefTable = ActivitiesTable
}