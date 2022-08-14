// Code generated by ent, DO NOT EDIT.

package app

import (
	"entgo.io/ent"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the app type in the database.
	Label = "app"
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
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldLogo holds the string denoting the logo field in the database.
	FieldLogo = "logo"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldSigninVerifyType holds the string denoting the signin_verify_type field in the database.
	FieldSigninVerifyType = "signin_verify_type"
	// Table holds the table name of the app in the database.
	Table = "apps"
)

// Columns holds all SQL columns for app fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldCreatedBy,
	FieldName,
	FieldLogo,
	FieldDescription,
	FieldSigninVerifyType,
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
	// DefaultLogo holds the default value on creation for the "logo" field.
	DefaultLogo string
	// DefaultDescription holds the default value on creation for the "description" field.
	DefaultDescription string
	// DefaultSigninVerifyType holds the default value on creation for the "signin_verify_type" field.
	DefaultSigninVerifyType string
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)
