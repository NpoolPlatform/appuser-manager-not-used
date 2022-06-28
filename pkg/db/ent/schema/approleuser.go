package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/appuser-manager/pkg/db/mixin"

	"github.com/google/uuid"
)

// AppRoleUser holds the schema definition for the AppRoleUser entity.
type AppRoleUser struct {
	ent.Schema
}

func (AppRoleUser) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
	}
}

// Fields of the AppRoleUser.
func (AppRoleUser) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique(),
		field.UUID("app_id", uuid.UUID{}),
		field.UUID("role_id", uuid.UUID{}),
		field.UUID("user_id", uuid.UUID{}),
	}
}

// Edges of the AppRoleUser.
func (AppRoleUser) Edges() []ent.Edge {
	return nil
}
