// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent/appuser"
	"github.com/google/uuid"
)

// AppUser is the model entity for the AppUser schema.
type AppUser struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt uint32 `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt uint32 `json:"updated_at,omitempty"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt uint32 `json:"deleted_at,omitempty"`
	// AppID holds the value of the "app_id" field.
	AppID uuid.UUID `json:"app_id,omitempty"`
	// EmailAddress holds the value of the "email_address" field.
	EmailAddress string `json:"email_address,omitempty"`
	// PhoneNo holds the value of the "phone_no" field.
	PhoneNo string `json:"phone_no,omitempty"`
	// ImportFromApp holds the value of the "import_from_app" field.
	ImportFromApp uuid.UUID `json:"import_from_app,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*AppUser) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case appuser.FieldCreatedAt, appuser.FieldUpdatedAt, appuser.FieldDeletedAt:
			values[i] = new(sql.NullInt64)
		case appuser.FieldEmailAddress, appuser.FieldPhoneNo:
			values[i] = new(sql.NullString)
		case appuser.FieldID, appuser.FieldAppID, appuser.FieldImportFromApp:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type AppUser", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the AppUser fields.
func (au *AppUser) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case appuser.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				au.ID = *value
			}
		case appuser.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				au.CreatedAt = uint32(value.Int64)
			}
		case appuser.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				au.UpdatedAt = uint32(value.Int64)
			}
		case appuser.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				au.DeletedAt = uint32(value.Int64)
			}
		case appuser.FieldAppID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field app_id", values[i])
			} else if value != nil {
				au.AppID = *value
			}
		case appuser.FieldEmailAddress:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field email_address", values[i])
			} else if value.Valid {
				au.EmailAddress = value.String
			}
		case appuser.FieldPhoneNo:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field phone_no", values[i])
			} else if value.Valid {
				au.PhoneNo = value.String
			}
		case appuser.FieldImportFromApp:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field import_from_app", values[i])
			} else if value != nil {
				au.ImportFromApp = *value
			}
		}
	}
	return nil
}

// Update returns a builder for updating this AppUser.
// Note that you need to call AppUser.Unwrap() before calling this method if this AppUser
// was returned from a transaction, and the transaction was committed or rolled back.
func (au *AppUser) Update() *AppUserUpdateOne {
	return (&AppUserClient{config: au.config}).UpdateOne(au)
}

// Unwrap unwraps the AppUser entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (au *AppUser) Unwrap() *AppUser {
	_tx, ok := au.config.driver.(*txDriver)
	if !ok {
		panic("ent: AppUser is not a transactional entity")
	}
	au.config.driver = _tx.drv
	return au
}

// String implements the fmt.Stringer.
func (au *AppUser) String() string {
	var builder strings.Builder
	builder.WriteString("AppUser(")
	builder.WriteString(fmt.Sprintf("id=%v, ", au.ID))
	builder.WriteString("created_at=")
	builder.WriteString(fmt.Sprintf("%v", au.CreatedAt))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(fmt.Sprintf("%v", au.UpdatedAt))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(fmt.Sprintf("%v", au.DeletedAt))
	builder.WriteString(", ")
	builder.WriteString("app_id=")
	builder.WriteString(fmt.Sprintf("%v", au.AppID))
	builder.WriteString(", ")
	builder.WriteString("email_address=")
	builder.WriteString(au.EmailAddress)
	builder.WriteString(", ")
	builder.WriteString("phone_no=")
	builder.WriteString(au.PhoneNo)
	builder.WriteString(", ")
	builder.WriteString("import_from_app=")
	builder.WriteString(fmt.Sprintf("%v", au.ImportFromApp))
	builder.WriteByte(')')
	return builder.String()
}

// AppUsers is a parsable slice of AppUser.
type AppUsers []*AppUser

func (au AppUsers) config(cfg config) {
	for _i := range au {
		au[_i].config = cfg
	}
}
