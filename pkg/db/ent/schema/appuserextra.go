package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"

	"github.com/google/uuid"
)

// AppUserExtra holds the schema definition for the AppUserExtra entity.
type AppUserExtra struct {
	ent.Schema
}

// Fields of the AppUserExtra.
func (AppUserExtra) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique(),
		field.UUID("app_id", uuid.UUID{}),
		field.UUID("user_id", uuid.UUID{}),
		field.String("username"),
		field.String("first_name"),
		field.String("last_name"),
		field.JSON("address_fields", []string{}),
		field.String("gender"),
		field.String("postal_code"),
		field.Uint32("age"),
		field.Uint32("birthday"),
		field.String("avatar"),
		field.String("Organization"),
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

// Edges of the AppUserExtra.
func (AppUserExtra) Edges() []ent.Edge {
	return nil
}

func (AppUserExtra) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("app_id", "user_id").Unique(),
	}
}
