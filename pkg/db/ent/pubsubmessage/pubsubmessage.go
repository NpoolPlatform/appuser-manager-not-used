// Code generated by ent, DO NOT EDIT.

package pubsubmessage

import (
	"entgo.io/ent"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the pubsubmessage type in the database.
	Label = "pubsub_message"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldDeletedAt holds the string denoting the deleted_at field in the database.
	FieldDeletedAt = "deleted_at"
	// FieldMessageID holds the string denoting the message_id field in the database.
	FieldMessageID = "message_id"
	// FieldState holds the string denoting the state field in the database.
	FieldState = "state"
	// FieldRespToID holds the string denoting the resp_to_id field in the database.
	FieldRespToID = "resp_to_id"
	// FieldUndoID holds the string denoting the undo_id field in the database.
	FieldUndoID = "undo_id"
	// FieldArguments holds the string denoting the arguments field in the database.
	FieldArguments = "arguments"
	// Table holds the table name of the pubsubmessage in the database.
	Table = "pubsub_messages"
)

// Columns holds all SQL columns for pubsubmessage fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldMessageID,
	FieldState,
	FieldRespToID,
	FieldUndoID,
	FieldArguments,
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
	// DefaultMessageID holds the default value on creation for the "message_id" field.
	DefaultMessageID string
	// DefaultState holds the default value on creation for the "state" field.
	DefaultState string
	// DefaultRespToID holds the default value on creation for the "resp_to_id" field.
	DefaultRespToID func() uuid.UUID
	// DefaultUndoID holds the default value on creation for the "undo_id" field.
	DefaultUndoID func() uuid.UUID
	// DefaultArguments holds the default value on creation for the "arguments" field.
	DefaultArguments string
)
