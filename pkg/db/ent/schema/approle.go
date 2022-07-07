package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/NpoolPlatform/appuser-manager/pkg/db/mixin"

	"github.com/google/uuid"
)

// AppRole holds the schema definition for the AppRole entity.
type AppRole struct {
	ent.Schema
}

func (AppRole) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
	}
}

// Fields of the AppRole.
func (AppRole) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique(),
		field.UUID("created_by", uuid.UUID{}),
		field.String("role"),
		field.String("description").
			Default(""),
		field.UUID("app_id", uuid.UUID{}),
		field.Bool("default"),
	}
}

// Edges of the AppRole.
func (AppRole) Edges() []ent.Edge {
	return nil
}

func (AppRole) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("app_id", "role").Unique(),
	}
}
