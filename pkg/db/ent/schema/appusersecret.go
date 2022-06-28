package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/NpoolPlatform/appuser-manager/pkg/db/mixin"

	"github.com/google/uuid"
)

// AppUserSecret holds the schema definition for the AppUserSecret entity.
type AppUserSecret struct {
	ent.Schema
}

func (AppUserSecret) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
	}
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
	}
}

// Edges of the AppUserSecret.
func (AppUserSecret) Edges() []ent.Edge {
	return nil
}

func (AppUserSecret) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("app_id", "user_id").Unique(),
	}
}
