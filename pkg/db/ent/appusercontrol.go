// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent/appusercontrol"
	"github.com/google/uuid"
)

// AppUserControl is the model entity for the AppUserControl schema.
type AppUserControl struct {
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
	// SigninVerifyByGoogleAuthentication holds the value of the "signin_verify_by_google_authentication" field.
	SigninVerifyByGoogleAuthentication bool `json:"signin_verify_by_google_authentication,omitempty"`
	// GoogleAuthenticationVerified holds the value of the "google_authentication_verified" field.
	GoogleAuthenticationVerified bool `json:"google_authentication_verified,omitempty"`
	// SigninVerifyType holds the value of the "signin_verify_type" field.
	SigninVerifyType string `json:"signin_verify_type,omitempty"`
	// Kol holds the value of the "kol" field.
	Kol bool `json:"kol,omitempty"`
	// KolConfirmed holds the value of the "kol_confirmed" field.
	KolConfirmed bool `json:"kol_confirmed,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*AppUserControl) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case appusercontrol.FieldSigninVerifyByGoogleAuthentication, appusercontrol.FieldGoogleAuthenticationVerified, appusercontrol.FieldKol, appusercontrol.FieldKolConfirmed:
			values[i] = new(sql.NullBool)
		case appusercontrol.FieldCreatedAt, appusercontrol.FieldUpdatedAt, appusercontrol.FieldDeletedAt:
			values[i] = new(sql.NullInt64)
		case appusercontrol.FieldSigninVerifyType:
			values[i] = new(sql.NullString)
		case appusercontrol.FieldID, appusercontrol.FieldAppID, appusercontrol.FieldUserID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type AppUserControl", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the AppUserControl fields.
func (auc *AppUserControl) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case appusercontrol.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				auc.ID = *value
			}
		case appusercontrol.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				auc.CreatedAt = uint32(value.Int64)
			}
		case appusercontrol.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				auc.UpdatedAt = uint32(value.Int64)
			}
		case appusercontrol.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				auc.DeletedAt = uint32(value.Int64)
			}
		case appusercontrol.FieldAppID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field app_id", values[i])
			} else if value != nil {
				auc.AppID = *value
			}
		case appusercontrol.FieldUserID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value != nil {
				auc.UserID = *value
			}
		case appusercontrol.FieldSigninVerifyByGoogleAuthentication:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field signin_verify_by_google_authentication", values[i])
			} else if value.Valid {
				auc.SigninVerifyByGoogleAuthentication = value.Bool
			}
		case appusercontrol.FieldGoogleAuthenticationVerified:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field google_authentication_verified", values[i])
			} else if value.Valid {
				auc.GoogleAuthenticationVerified = value.Bool
			}
		case appusercontrol.FieldSigninVerifyType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field signin_verify_type", values[i])
			} else if value.Valid {
				auc.SigninVerifyType = value.String
			}
		case appusercontrol.FieldKol:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field kol", values[i])
			} else if value.Valid {
				auc.Kol = value.Bool
			}
		case appusercontrol.FieldKolConfirmed:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field kol_confirmed", values[i])
			} else if value.Valid {
				auc.KolConfirmed = value.Bool
			}
		}
	}
	return nil
}

// Update returns a builder for updating this AppUserControl.
// Note that you need to call AppUserControl.Unwrap() before calling this method if this AppUserControl
// was returned from a transaction, and the transaction was committed or rolled back.
func (auc *AppUserControl) Update() *AppUserControlUpdateOne {
	return (&AppUserControlClient{config: auc.config}).UpdateOne(auc)
}

// Unwrap unwraps the AppUserControl entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (auc *AppUserControl) Unwrap() *AppUserControl {
	_tx, ok := auc.config.driver.(*txDriver)
	if !ok {
		panic("ent: AppUserControl is not a transactional entity")
	}
	auc.config.driver = _tx.drv
	return auc
}

// String implements the fmt.Stringer.
func (auc *AppUserControl) String() string {
	var builder strings.Builder
	builder.WriteString("AppUserControl(")
	builder.WriteString(fmt.Sprintf("id=%v, ", auc.ID))
	builder.WriteString("created_at=")
	builder.WriteString(fmt.Sprintf("%v", auc.CreatedAt))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(fmt.Sprintf("%v", auc.UpdatedAt))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(fmt.Sprintf("%v", auc.DeletedAt))
	builder.WriteString(", ")
	builder.WriteString("app_id=")
	builder.WriteString(fmt.Sprintf("%v", auc.AppID))
	builder.WriteString(", ")
	builder.WriteString("user_id=")
	builder.WriteString(fmt.Sprintf("%v", auc.UserID))
	builder.WriteString(", ")
	builder.WriteString("signin_verify_by_google_authentication=")
	builder.WriteString(fmt.Sprintf("%v", auc.SigninVerifyByGoogleAuthentication))
	builder.WriteString(", ")
	builder.WriteString("google_authentication_verified=")
	builder.WriteString(fmt.Sprintf("%v", auc.GoogleAuthenticationVerified))
	builder.WriteString(", ")
	builder.WriteString("signin_verify_type=")
	builder.WriteString(auc.SigninVerifyType)
	builder.WriteString(", ")
	builder.WriteString("kol=")
	builder.WriteString(fmt.Sprintf("%v", auc.Kol))
	builder.WriteString(", ")
	builder.WriteString("kol_confirmed=")
	builder.WriteString(fmt.Sprintf("%v", auc.KolConfirmed))
	builder.WriteByte(')')
	return builder.String()
}

// AppUserControls is a parsable slice of AppUserControl.
type AppUserControls []*AppUserControl

func (auc AppUserControls) config(cfg config) {
	for _i := range auc {
		auc[_i].config = cfg
	}
}
