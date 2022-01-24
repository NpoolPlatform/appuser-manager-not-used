// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent/genesisuser"
	"github.com/google/uuid"
)

// GenesisUser is the model entity for the GenesisUser schema.
type GenesisUser struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// UserID holds the value of the "user_id" field.
	UserID uuid.UUID `json:"user_id,omitempty"`
	// CreateAt holds the value of the "create_at" field.
	CreateAt uint32 `json:"create_at,omitempty"`
	// UpdateAt holds the value of the "update_at" field.
	UpdateAt uint32 `json:"update_at,omitempty"`
	// DeleteAt holds the value of the "delete_at" field.
	DeleteAt uint32 `json:"delete_at,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*GenesisUser) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case genesisuser.FieldCreateAt, genesisuser.FieldUpdateAt, genesisuser.FieldDeleteAt:
			values[i] = new(sql.NullInt64)
		case genesisuser.FieldID, genesisuser.FieldUserID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type GenesisUser", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the GenesisUser fields.
func (gu *GenesisUser) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case genesisuser.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				gu.ID = *value
			}
		case genesisuser.FieldUserID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value != nil {
				gu.UserID = *value
			}
		case genesisuser.FieldCreateAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field create_at", values[i])
			} else if value.Valid {
				gu.CreateAt = uint32(value.Int64)
			}
		case genesisuser.FieldUpdateAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field update_at", values[i])
			} else if value.Valid {
				gu.UpdateAt = uint32(value.Int64)
			}
		case genesisuser.FieldDeleteAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field delete_at", values[i])
			} else if value.Valid {
				gu.DeleteAt = uint32(value.Int64)
			}
		}
	}
	return nil
}

// Update returns a builder for updating this GenesisUser.
// Note that you need to call GenesisUser.Unwrap() before calling this method if this GenesisUser
// was returned from a transaction, and the transaction was committed or rolled back.
func (gu *GenesisUser) Update() *GenesisUserUpdateOne {
	return (&GenesisUserClient{config: gu.config}).UpdateOne(gu)
}

// Unwrap unwraps the GenesisUser entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (gu *GenesisUser) Unwrap() *GenesisUser {
	tx, ok := gu.config.driver.(*txDriver)
	if !ok {
		panic("ent: GenesisUser is not a transactional entity")
	}
	gu.config.driver = tx.drv
	return gu
}

// String implements the fmt.Stringer.
func (gu *GenesisUser) String() string {
	var builder strings.Builder
	builder.WriteString("GenesisUser(")
	builder.WriteString(fmt.Sprintf("id=%v", gu.ID))
	builder.WriteString(", user_id=")
	builder.WriteString(fmt.Sprintf("%v", gu.UserID))
	builder.WriteString(", create_at=")
	builder.WriteString(fmt.Sprintf("%v", gu.CreateAt))
	builder.WriteString(", update_at=")
	builder.WriteString(fmt.Sprintf("%v", gu.UpdateAt))
	builder.WriteString(", delete_at=")
	builder.WriteString(fmt.Sprintf("%v", gu.DeleteAt))
	builder.WriteByte(')')
	return builder.String()
}

// GenesisUsers is a parsable slice of GenesisUser.
type GenesisUsers []*GenesisUser

func (gu GenesisUsers) config(cfg config) {
	for _i := range gu {
		gu[_i].config = cfg
	}
}
