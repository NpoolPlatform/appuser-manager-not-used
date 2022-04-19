package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"

	"github.com/google/uuid"
)

// AppUserSecret holds the schema definition for the AppUserSecret entity.
type AppUserSecret struct {
	ent.Schema
}

// Fields of the AppUserSecret.
func (AppUserSecret) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique(),
		field.UUID("app_id", uuid.UUID{}),
		field.UUID("user_id", uuid.UUID{}),
		field.String("password_hash"),
		field.String("salt"),
		field.String("google_secret"),
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

// Edges of the AppUserSecret.
func (AppUserSecret) Edges() []ent.Edge {
	return nil
}

func (AppUserSecret) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("app_id", "user_id", "delete_at").Unique(),
	}
}
