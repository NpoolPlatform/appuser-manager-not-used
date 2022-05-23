// Code generated by entc, DO NOT EDIT.

package appuserthird

import (
	"entgo.io/ent"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the appuserthird type in the database.
	Label = "app_user_third"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreateAt holds the string denoting the create_at field in the database.
	FieldCreateAt = "create_at"
	// FieldUpdateAt holds the string denoting the update_at field in the database.
	FieldUpdateAt = "update_at"
	// FieldDeleteAt holds the string denoting the delete_at field in the database.
	FieldDeleteAt = "delete_at"
	// FieldAppID holds the string denoting the app_id field in the database.
	FieldAppID = "app_id"
	// FieldUserID holds the string denoting the user_id field in the database.
	FieldUserID = "user_id"
	// FieldThirdUserID holds the string denoting the third_user_id field in the database.
	FieldThirdUserID = "third_user_id"
	// FieldThird holds the string denoting the third field in the database.
	FieldThird = "third"
	// FieldThirdID holds the string denoting the third_id field in the database.
	FieldThirdID = "third_id"
	// FieldThirdUserName holds the string denoting the third_user_name field in the database.
	FieldThirdUserName = "third_user_name"
	// FieldThirdUserAvatar holds the string denoting the third_user_avatar field in the database.
	FieldThirdUserAvatar = "third_user_avatar"
	// Table holds the table name of the appuserthird in the database.
	Table = "app_user_thirds"
)

// Columns holds all SQL columns for appuserthird fields.
var Columns = []string{
	FieldID,
	FieldCreateAt,
	FieldUpdateAt,
	FieldDeleteAt,
	FieldAppID,
	FieldUserID,
	FieldThirdUserID,
	FieldThird,
	FieldThirdID,
	FieldThirdUserName,
	FieldThirdUserAvatar,
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
	// DefaultCreateAt holds the default value on creation for the "create_at" field.
	DefaultCreateAt func() uint32
	// DefaultUpdateAt holds the default value on creation for the "update_at" field.
	DefaultUpdateAt func() uint32
	// UpdateDefaultUpdateAt holds the default value on update for the "update_at" field.
	UpdateDefaultUpdateAt func() uint32
	// DefaultDeleteAt holds the default value on creation for the "delete_at" field.
	DefaultDeleteAt func() uint32
	// ThirdUserAvatarValidator is a validator for the "third_user_avatar" field. It is called by the builders before save.
	ThirdUserAvatarValidator func(string) error
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)
