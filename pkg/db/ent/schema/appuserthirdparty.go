package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// AppUserThirdParty holds the schema definition for the AppUserThirdParty entity.
type AppUserThirdParty struct {
	ent.Schema
}

func (AppUserThirdParty) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

// Fields of the AppUserThirdParty.
func (AppUserThirdParty) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique(),
		field.UUID("app_id", uuid.UUID{}),
		field.UUID("user_id", uuid.UUID{}),
		field.String("third_party_user_id"),
		field.String("third_party_id"),
		field.String("third_party_user_name"),
		field.String("third_party_user_avatar").MaxLen(1024),
	}
}

// Edges of the AppUserThirdParty.
func (AppUserThirdParty) Edges() []ent.Edge {
	return nil
}

func (AppUserThirdParty) Indexes() []ent.Index {
	return []ent.Index{}
}
