// Code generated by ent, DO NOT EDIT.

package approle

import (
	"entgo.io/ent"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the approle type in the database.
	Label = "app_role"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldDeletedAt holds the string denoting the deleted_at field in the database.
	FieldDeletedAt = "deleted_at"
	// FieldCreatedBy holds the string denoting the created_by field in the database.
	FieldCreatedBy = "created_by"
	// FieldRole holds the string denoting the role field in the database.
	FieldRole = "role"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldAppID holds the string denoting the app_id field in the database.
	FieldAppID = "app_id"
	// FieldDefault holds the string denoting the default field in the database.
	FieldDefault = "default"
	// FieldGenesis holds the string denoting the genesis field in the database.
	FieldGenesis = "genesis"
	// Table holds the table name of the approle in the database.
	Table = "app_roles"
)

// Columns holds all SQL columns for approle fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldCreatedBy,
	FieldRole,
	FieldDescription,
	FieldAppID,
	FieldDefault,
	FieldGenesis,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

// Note that the variables below are initialized by the runtime
// package on the initialization of the application. Therefore,
// it should be imported in the main as follows:
//
//	import _ "github.com/NpoolPlatform/appuser-manager/pkg/db/ent/runtime"
//
var (
	Hooks  [1]ent.Hook
	Policy ent.Policy
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() uint32
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() uint32
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() uint32
	// DefaultDeletedAt holds the default value on creation for the "deleted_at" field.
	DefaultDeletedAt func() uint32
	// DefaultCreatedBy holds the default value on creation for the "created_by" field.
	DefaultCreatedBy func() uuid.UUID
	// DefaultRole holds the default value on creation for the "role" field.
	DefaultRole string
	// DefaultDescription holds the default value on creation for the "description" field.
	DefaultDescription string
	// DefaultAppID holds the default value on creation for the "app_id" field.
	DefaultAppID func() uuid.UUID
	// DefaultDefault holds the default value on creation for the "default" field.
	DefaultDefault bool
	// DefaultGenesis holds the default value on creation for the "genesis" field.
	DefaultGenesis bool
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)
