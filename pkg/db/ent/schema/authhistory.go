package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/appuser-manager/pkg/db/mixin"

	"github.com/google/uuid"
)

// AuthHistory holds the schema definition for the AuthHistory entity.
type AuthHistory struct {
	ent.Schema
}

func (AuthHistory) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
	}
}

// Fields of the AuthHistory.
func (AuthHistory) Fields() []ent.Field {
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
			String("resource").
			Optional().
			Default(""),
		field.
			String("method").
			Optional().
			Default(""),
		field.
			Bool("allowed").
			Optional().
			Default(false),
	}
}

// Edges of the AuthHistory.
func (AuthHistory) Edges() []ent.Edge {
	return nil
}
