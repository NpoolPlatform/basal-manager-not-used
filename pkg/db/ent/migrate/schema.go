// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// ApIsColumns holds the columns for the "ap_is" table.
	ApIsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "protocol", Type: field.TypeString, Nullable: true, Default: "DefaultProtocol"},
		{Name: "service_name", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "method", Type: field.TypeString, Nullable: true, Default: "DefaultMethod"},
		{Name: "method_name", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "path", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "exported", Type: field.TypeBool, Nullable: true, Default: false},
		{Name: "path_prefix", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "domains", Type: field.TypeJSON, Nullable: true},
		{Name: "depracated", Type: field.TypeBool, Nullable: true, Default: false},
	}
	// ApIsTable holds the schema information for the "ap_is" table.
	ApIsTable = &schema.Table{
		Name:       "ap_is",
		Columns:    ApIsColumns,
		PrimaryKey: []*schema.Column{ApIsColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		ApIsTable,
	}
)

func init() {
}
