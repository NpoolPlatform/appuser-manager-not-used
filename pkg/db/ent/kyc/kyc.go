// Code generated by ent, DO NOT EDIT.

package kyc

import (
	"entgo.io/ent"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the kyc type in the database.
	Label = "kyc"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldDeletedAt holds the string denoting the deleted_at field in the database.
	FieldDeletedAt = "deleted_at"
	// FieldAppID holds the string denoting the app_id field in the database.
	FieldAppID = "app_id"
	// FieldUserID holds the string denoting the user_id field in the database.
	FieldUserID = "user_id"
	// FieldDocumentType holds the string denoting the document_type field in the database.
	FieldDocumentType = "document_type"
	// FieldIDNumber holds the string denoting the id_number field in the database.
	FieldIDNumber = "id_number"
	// FieldFrontImg holds the string denoting the front_img field in the database.
	FieldFrontImg = "front_img"
	// FieldBackImg holds the string denoting the back_img field in the database.
	FieldBackImg = "back_img"
	// FieldSelfieImg holds the string denoting the selfie_img field in the database.
	FieldSelfieImg = "selfie_img"
	// FieldEntityType holds the string denoting the entity_type field in the database.
	FieldEntityType = "entity_type"
	// FieldReviewID holds the string denoting the review_id field in the database.
	FieldReviewID = "review_id"
	// FieldState holds the string denoting the state field in the database.
	FieldState = "state"
	// Table holds the table name of the kyc in the database.
	Table = "kycs"
)

// Columns holds all SQL columns for kyc fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldAppID,
	FieldUserID,
	FieldDocumentType,
	FieldIDNumber,
	FieldFrontImg,
	FieldBackImg,
	FieldSelfieImg,
	FieldEntityType,
	FieldReviewID,
	FieldState,
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
	// DefaultAppID holds the default value on creation for the "app_id" field.
	DefaultAppID func() uuid.UUID
	// DefaultUserID holds the default value on creation for the "user_id" field.
	DefaultUserID func() uuid.UUID
	// DefaultDocumentType holds the default value on creation for the "document_type" field.
	DefaultDocumentType string
	// DefaultIDNumber holds the default value on creation for the "id_number" field.
	DefaultIDNumber string
	// DefaultFrontImg holds the default value on creation for the "front_img" field.
	DefaultFrontImg string
	// DefaultBackImg holds the default value on creation for the "back_img" field.
	DefaultBackImg string
	// DefaultSelfieImg holds the default value on creation for the "selfie_img" field.
	DefaultSelfieImg string
	// DefaultEntityType holds the default value on creation for the "entity_type" field.
	DefaultEntityType string
	// DefaultReviewID holds the default value on creation for the "review_id" field.
	DefaultReviewID func() uuid.UUID
	// DefaultState holds the default value on creation for the "state" field.
	DefaultState string
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)
