package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"

	"github.com/google/uuid"
)

// AppRole holds the schema definition for the AppRole entity.
type AppRole struct {
	ent.Schema
}

// Fields of the AppRole.
func (AppRole) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique(),
		field.UUID("created_by", uuid.UUID{}),
		field.String("role"),
		field.String("description"),
		field.UUID("app_id", uuid.UUID{}),
		field.Bool("default"),
		field.Uint32("create_at").
			DefaultFunc(func() uint32 {
				return uint32(time.Now().Unix())
			}),
		field.Uint32("update_at").
			DefaultFunc(func() uint32 {
				return uint32(time.Now().Unix())
			}).
			UpdateDefault(func() uint32 {
				return uint32(time.Now().Unix())
			}),
		field.Uint32("delete_at").
			DefaultFunc(func() uint32 {
				return 0
			}),
	}
}

// Edges of the AppRole.
func (AppRole) Edges() []ent.Edge {
	return nil
}
