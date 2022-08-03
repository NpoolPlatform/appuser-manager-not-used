package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/appuser-manager/pkg/db/mixin"

	"github.com/google/uuid"
)

// AppControl holds the schema definition for the AppControl entity.
type AppControl struct {
	ent.Schema
}

func (AppControl) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
	}
}

// Fields of the AppControl.
func (AppControl) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique(),
		field.UUID("app_id", uuid.UUID{}).
			Unique(),
		field.JSON("signup_methods", []string{}),
		field.JSON("extern_signin_methods", []string{}),
		field.String("recaptcha_method"),
		field.Bool("kyc_enable"),
		field.Bool("signin_verify_enable"),
		field.Bool("invitation_code_must"),
	}
}

// Edges of the AppControl.
func (AppControl) Edges() []ent.Edge {
	return nil
}
