// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent/subscriber"
	"github.com/google/uuid"
)

// Subscriber is the model entity for the Subscriber schema.
type Subscriber struct {
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
	// Registered holds the value of the "registered" field.
	Registered bool `json:"registered,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Subscriber) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case subscriber.FieldRegistered:
			values[i] = new(sql.NullBool)
		case subscriber.FieldCreatedAt, subscriber.FieldUpdatedAt, subscriber.FieldDeletedAt:
			values[i] = new(sql.NullInt64)
		case subscriber.FieldEmailAddress:
			values[i] = new(sql.NullString)
		case subscriber.FieldID, subscriber.FieldAppID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Subscriber", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Subscriber fields.
func (s *Subscriber) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case subscriber.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				s.ID = *value
			}
		case subscriber.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				s.CreatedAt = uint32(value.Int64)
			}
		case subscriber.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				s.UpdatedAt = uint32(value.Int64)
			}
		case subscriber.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				s.DeletedAt = uint32(value.Int64)
			}
		case subscriber.FieldAppID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field app_id", values[i])
			} else if value != nil {
				s.AppID = *value
			}
		case subscriber.FieldEmailAddress:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field email_address", values[i])
			} else if value.Valid {
				s.EmailAddress = value.String
			}
		case subscriber.FieldRegistered:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field registered", values[i])
			} else if value.Valid {
				s.Registered = value.Bool
			}
		}
	}
	return nil
}

// Update returns a builder for updating this Subscriber.
// Note that you need to call Subscriber.Unwrap() before calling this method if this Subscriber
// was returned from a transaction, and the transaction was committed or rolled back.
func (s *Subscriber) Update() *SubscriberUpdateOne {
	return (&SubscriberClient{config: s.config}).UpdateOne(s)
}

// Unwrap unwraps the Subscriber entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (s *Subscriber) Unwrap() *Subscriber {
	_tx, ok := s.config.driver.(*txDriver)
	if !ok {
		panic("ent: Subscriber is not a transactional entity")
	}
	s.config.driver = _tx.drv
	return s
}

// String implements the fmt.Stringer.
func (s *Subscriber) String() string {
	var builder strings.Builder
	builder.WriteString("Subscriber(")
	builder.WriteString(fmt.Sprintf("id=%v, ", s.ID))
	builder.WriteString("created_at=")
	builder.WriteString(fmt.Sprintf("%v", s.CreatedAt))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(fmt.Sprintf("%v", s.UpdatedAt))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(fmt.Sprintf("%v", s.DeletedAt))
	builder.WriteString(", ")
	builder.WriteString("app_id=")
	builder.WriteString(fmt.Sprintf("%v", s.AppID))
	builder.WriteString(", ")
	builder.WriteString("email_address=")
	builder.WriteString(s.EmailAddress)
	builder.WriteString(", ")
	builder.WriteString("registered=")
	builder.WriteString(fmt.Sprintf("%v", s.Registered))
	builder.WriteByte(')')
	return builder.String()
}

// Subscribers is a parsable slice of Subscriber.
type Subscribers []*Subscriber

func (s Subscribers) config(cfg config) {
	for _i := range s {
		s[_i].config = cfg
	}
}
