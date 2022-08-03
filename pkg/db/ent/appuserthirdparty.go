// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent/appuserthirdparty"
	"github.com/google/uuid"
)

// AppUserThirdParty is the model entity for the AppUserThirdParty schema.
type AppUserThirdParty struct {
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
	// UserID holds the value of the "user_id" field.
	UserID uuid.UUID `json:"user_id,omitempty"`
	// ThirdPartyUserID holds the value of the "third_party_user_id" field.
	ThirdPartyUserID string `json:"third_party_user_id,omitempty"`
	// ThirdPartyID holds the value of the "third_party_id" field.
	ThirdPartyID string `json:"third_party_id,omitempty"`
	// ThirdPartyUsername holds the value of the "third_party_username" field.
	ThirdPartyUsername string `json:"third_party_username,omitempty"`
	// ThirdPartyUserAvatar holds the value of the "third_party_user_avatar" field.
	ThirdPartyUserAvatar string `json:"third_party_user_avatar,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*AppUserThirdParty) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case appuserthirdparty.FieldCreatedAt, appuserthirdparty.FieldUpdatedAt, appuserthirdparty.FieldDeletedAt:
			values[i] = new(sql.NullInt64)
		case appuserthirdparty.FieldThirdPartyUserID, appuserthirdparty.FieldThirdPartyID, appuserthirdparty.FieldThirdPartyUsername, appuserthirdparty.FieldThirdPartyUserAvatar:
			values[i] = new(sql.NullString)
		case appuserthirdparty.FieldID, appuserthirdparty.FieldAppID, appuserthirdparty.FieldUserID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type AppUserThirdParty", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the AppUserThirdParty fields.
func (autp *AppUserThirdParty) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case appuserthirdparty.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				autp.ID = *value
			}
		case appuserthirdparty.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				autp.CreatedAt = uint32(value.Int64)
			}
		case appuserthirdparty.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				autp.UpdatedAt = uint32(value.Int64)
			}
		case appuserthirdparty.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				autp.DeletedAt = uint32(value.Int64)
			}
		case appuserthirdparty.FieldAppID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field app_id", values[i])
			} else if value != nil {
				autp.AppID = *value
			}
		case appuserthirdparty.FieldUserID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value != nil {
				autp.UserID = *value
			}
		case appuserthirdparty.FieldThirdPartyUserID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field third_party_user_id", values[i])
			} else if value.Valid {
				autp.ThirdPartyUserID = value.String
			}
		case appuserthirdparty.FieldThirdPartyID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field third_party_id", values[i])
			} else if value.Valid {
				autp.ThirdPartyID = value.String
			}
		case appuserthirdparty.FieldThirdPartyUsername:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field third_party_username", values[i])
			} else if value.Valid {
				autp.ThirdPartyUsername = value.String
			}
		case appuserthirdparty.FieldThirdPartyUserAvatar:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field third_party_user_avatar", values[i])
			} else if value.Valid {
				autp.ThirdPartyUserAvatar = value.String
			}
		}
	}
	return nil
}

// Update returns a builder for updating this AppUserThirdParty.
// Note that you need to call AppUserThirdParty.Unwrap() before calling this method if this AppUserThirdParty
// was returned from a transaction, and the transaction was committed or rolled back.
func (autp *AppUserThirdParty) Update() *AppUserThirdPartyUpdateOne {
	return (&AppUserThirdPartyClient{config: autp.config}).UpdateOne(autp)
}

// Unwrap unwraps the AppUserThirdParty entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (autp *AppUserThirdParty) Unwrap() *AppUserThirdParty {
	_tx, ok := autp.config.driver.(*txDriver)
	if !ok {
		panic("ent: AppUserThirdParty is not a transactional entity")
	}
	autp.config.driver = _tx.drv
	return autp
}

// String implements the fmt.Stringer.
func (autp *AppUserThirdParty) String() string {
	var builder strings.Builder
	builder.WriteString("AppUserThirdParty(")
	builder.WriteString(fmt.Sprintf("id=%v, ", autp.ID))
	builder.WriteString("created_at=")
	builder.WriteString(fmt.Sprintf("%v", autp.CreatedAt))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(fmt.Sprintf("%v", autp.UpdatedAt))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(fmt.Sprintf("%v", autp.DeletedAt))
	builder.WriteString(", ")
	builder.WriteString("app_id=")
	builder.WriteString(fmt.Sprintf("%v", autp.AppID))
	builder.WriteString(", ")
	builder.WriteString("user_id=")
	builder.WriteString(fmt.Sprintf("%v", autp.UserID))
	builder.WriteString(", ")
	builder.WriteString("third_party_user_id=")
	builder.WriteString(autp.ThirdPartyUserID)
	builder.WriteString(", ")
	builder.WriteString("third_party_id=")
	builder.WriteString(autp.ThirdPartyID)
	builder.WriteString(", ")
	builder.WriteString("third_party_username=")
	builder.WriteString(autp.ThirdPartyUsername)
	builder.WriteString(", ")
	builder.WriteString("third_party_user_avatar=")
	builder.WriteString(autp.ThirdPartyUserAvatar)
	builder.WriteByte(')')
	return builder.String()
}

// AppUserThirdParties is a parsable slice of AppUserThirdParty.
type AppUserThirdParties []*AppUserThirdParty

func (autp AppUserThirdParties) config(cfg config) {
	for _i := range autp {
		autp[_i].config = cfg
	}
}
