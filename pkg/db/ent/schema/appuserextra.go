package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/NpoolPlatform/appuser-manager/pkg/db/mixin"

	"github.com/google/uuid"
)

// AppUserExtra holds the schema definition for the AppUserExtra entity.
type AppUserExtra struct {
	ent.Schema
}

func (AppUserExtra) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
	}
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
		field.String("organization"),
		field.String("id_number"),
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
