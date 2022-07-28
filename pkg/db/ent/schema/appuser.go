package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/appuser-manager/pkg/db/mixin"

	"github.com/google/uuid"
)

// AppUser holds the schema definition for the AppUser entity.
type AppUser struct {
	ent.Schema
}

func (AppUser) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
	}
}

// Fields of the AppUser.
func (AppUser) Fields() []ent.Field {
	return []ent.Field{
		field.
			UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique(),
		field.
			UUID("app_id", uuid.UUID{}),
		field.
			String("email_address").
			Default(""),
		field.
			String("phone_no").
			Default(""),
		field.
			UUID("import_from_app", uuid.UUID{}),
	}
}
