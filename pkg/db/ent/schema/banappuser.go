package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/appuser-manager/pkg/db/mixin"

	"github.com/google/uuid"
)

// BanAppUser holds the schema definition for the BanAppUser entity.
type BanAppUser struct {
	ent.Schema
}

func (BanAppUser) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
	}
}

// Fields of the BanAppUser.
func (BanAppUser) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique(),
		field.UUID("app_id", uuid.UUID{}),
		field.UUID("user_id", uuid.UUID{}),
		field.String("message").
			Default(""),
	}
}

// Edges of the BanAppUser.
func (BanAppUser) Edges() []ent.Edge {
	return nil
}
