package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/appuser-manager/pkg/db/mixin"

	"github.com/google/uuid"
)

// LoginHistory holds the schema definition for the LoginHistory entity.
type LoginHistory struct {
	ent.Schema
}

func (LoginHistory) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
	}
}

// Fields of the LoginHistory.
func (LoginHistory) Fields() []ent.Field {
	return []ent.Field{
		field.
			UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique(),
		field.
			UUID("app_id", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.UUID{}
			}),
		field.
			UUID("user_id", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.UUID{}
			}),
		field.
			String("client_ip").
			Optional().
			Default(""),
		field.
			String("user_agent").
			Optional().
			Default(""),
		field.
			String("location").
			Optional().
			Default(""),
	}
}

// Edges of the LoginHistory.
func (LoginHistory) Edges() []ent.Edge {
	return nil
}
