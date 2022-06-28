package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/NpoolPlatform/appuser-manager/pkg/db/mixin"

	"github.com/google/uuid"
)

// AppUserControl holds the schema definition for the AppUserControl entity.
type AppUserControl struct {
	ent.Schema
}

func (AppUserControl) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
	}
}

// Fields of the AppUserControl.
func (AppUserControl) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique(),
		field.UUID("app_id", uuid.UUID{}),
		field.UUID("user_id", uuid.UUID{}),
		field.Bool("signin_verify_by_google_authentication"),
		field.Bool("google_authentication_verified"),
	}
}

// Edges of the AppUserControl.
func (AppUserControl) Edges() []ent.Edge {
	return nil
}

func (AppUserControl) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("app_id", "user_id").Unique(),
	}
}
