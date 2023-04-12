// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent/pubsubmessage"
	"github.com/google/uuid"
)

// PubsubMessage is the model entity for the PubsubMessage schema.
type PubsubMessage struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt uint32 `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt uint32 `json:"updated_at,omitempty"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt uint32 `json:"deleted_at,omitempty"`
	// MessageID holds the value of the "message_id" field.
	MessageID string `json:"message_id,omitempty"`
	// State holds the value of the "state" field.
	State string `json:"state,omitempty"`
	// RespToID holds the value of the "resp_to_id" field.
	RespToID uuid.UUID `json:"resp_to_id,omitempty"`
	// UndoID holds the value of the "undo_id" field.
	UndoID uuid.UUID `json:"undo_id,omitempty"`
	// Arguments holds the value of the "arguments" field.
	Arguments string `json:"arguments,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*PubsubMessage) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case pubsubmessage.FieldCreatedAt, pubsubmessage.FieldUpdatedAt, pubsubmessage.FieldDeletedAt:
			values[i] = new(sql.NullInt64)
		case pubsubmessage.FieldMessageID, pubsubmessage.FieldState, pubsubmessage.FieldArguments:
			values[i] = new(sql.NullString)
		case pubsubmessage.FieldID, pubsubmessage.FieldRespToID, pubsubmessage.FieldUndoID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type PubsubMessage", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the PubsubMessage fields.
func (pm *PubsubMessage) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case pubsubmessage.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				pm.ID = *value
			}
		case pubsubmessage.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				pm.CreatedAt = uint32(value.Int64)
			}
		case pubsubmessage.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				pm.UpdatedAt = uint32(value.Int64)
			}
		case pubsubmessage.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				pm.DeletedAt = uint32(value.Int64)
			}
		case pubsubmessage.FieldMessageID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field message_id", values[i])
			} else if value.Valid {
				pm.MessageID = value.String
			}
		case pubsubmessage.FieldState:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field state", values[i])
			} else if value.Valid {
				pm.State = value.String
			}
		case pubsubmessage.FieldRespToID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field resp_to_id", values[i])
			} else if value != nil {
				pm.RespToID = *value
			}
		case pubsubmessage.FieldUndoID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field undo_id", values[i])
			} else if value != nil {
				pm.UndoID = *value
			}
		case pubsubmessage.FieldArguments:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field arguments", values[i])
			} else if value.Valid {
				pm.Arguments = value.String
			}
		}
	}
	return nil
}

// Update returns a builder for updating this PubsubMessage.
// Note that you need to call PubsubMessage.Unwrap() before calling this method if this PubsubMessage
// was returned from a transaction, and the transaction was committed or rolled back.
func (pm *PubsubMessage) Update() *PubsubMessageUpdateOne {
	return (&PubsubMessageClient{config: pm.config}).UpdateOne(pm)
}

// Unwrap unwraps the PubsubMessage entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pm *PubsubMessage) Unwrap() *PubsubMessage {
	_tx, ok := pm.config.driver.(*txDriver)
	if !ok {
		panic("ent: PubsubMessage is not a transactional entity")
	}
	pm.config.driver = _tx.drv
	return pm
}

// String implements the fmt.Stringer.
func (pm *PubsubMessage) String() string {
	var builder strings.Builder
	builder.WriteString("PubsubMessage(")
	builder.WriteString(fmt.Sprintf("id=%v, ", pm.ID))
	builder.WriteString("created_at=")
	builder.WriteString(fmt.Sprintf("%v", pm.CreatedAt))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(fmt.Sprintf("%v", pm.UpdatedAt))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(fmt.Sprintf("%v", pm.DeletedAt))
	builder.WriteString(", ")
	builder.WriteString("message_id=")
	builder.WriteString(pm.MessageID)
	builder.WriteString(", ")
	builder.WriteString("state=")
	builder.WriteString(pm.State)
	builder.WriteString(", ")
	builder.WriteString("resp_to_id=")
	builder.WriteString(fmt.Sprintf("%v", pm.RespToID))
	builder.WriteString(", ")
	builder.WriteString("undo_id=")
	builder.WriteString(fmt.Sprintf("%v", pm.UndoID))
	builder.WriteString(", ")
	builder.WriteString("arguments=")
	builder.WriteString(pm.Arguments)
	builder.WriteByte(')')
	return builder.String()
}

// PubsubMessages is a parsable slice of PubsubMessage.
type PubsubMessages []*PubsubMessage

func (pm PubsubMessages) config(cfg config) {
	for _i := range pm {
		pm[_i].config = cfg
	}
}
