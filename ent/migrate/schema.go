// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// DestiniesColumns holds the columns for the "destinies" table.
	DestiniesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "price", Type: field.TypeFloat64},
		{Name: "meta", Type: field.TypeString, Size: 160},
		{Name: "description", Type: field.TypeString, Nullable: true, Size: 2147483647},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
	}
	// DestiniesTable holds the schema information for the "destinies" table.
	DestiniesTable = &schema.Table{
		Name:       "destinies",
		Columns:    DestiniesColumns,
		PrimaryKey: []*schema.Column{DestiniesColumns[0]},
	}
	// DestinyPicturesColumns holds the columns for the "destiny_pictures" table.
	DestinyPicturesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "picture", Type: field.TypeString},
		{Name: "path", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "destiny_id", Type: field.TypeInt},
	}
	// DestinyPicturesTable holds the schema information for the "destiny_pictures" table.
	DestinyPicturesTable = &schema.Table{
		Name:       "destiny_pictures",
		Columns:    DestinyPicturesColumns,
		PrimaryKey: []*schema.Column{DestinyPicturesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "destiny_pictures_destinies_destiny_pictures",
				Columns:    []*schema.Column{DestinyPicturesColumns[5]},
				RefColumns: []*schema.Column{DestiniesColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// TestimoniesColumns holds the columns for the "testimonies" table.
	TestimoniesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "testimony", Type: field.TypeString, Size: 2147483647},
		{Name: "name", Type: field.TypeString},
		{Name: "picture", Type: field.TypeString, Unique: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "destiny_id", Type: field.TypeInt},
	}
	// TestimoniesTable holds the schema information for the "testimonies" table.
	TestimoniesTable = &schema.Table{
		Name:       "testimonies",
		Columns:    TestimoniesColumns,
		PrimaryKey: []*schema.Column{TestimoniesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "testimonies_destinies_testimonies",
				Columns:    []*schema.Column{TestimoniesColumns[6]},
				RefColumns: []*schema.Column{DestiniesColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "email", Type: field.TypeString, Unique: true},
		{Name: "password", Type: field.TypeString, Nullable: true},
		{Name: "rand_security", Type: field.TypeString, Nullable: true},
		{Name: "password_token", Type: field.TypeString, Nullable: true, Size: 36},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		DestiniesTable,
		DestinyPicturesTable,
		TestimoniesTable,
		UsersTable,
	}
)

func init() {
	DestinyPicturesTable.ForeignKeys[0].RefTable = DestiniesTable
	TestimoniesTable.ForeignKeys[0].RefTable = DestiniesTable
}
