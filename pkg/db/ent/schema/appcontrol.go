package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/appuser-manager/pkg/db/mixin"

	"github.com/google/uuid"

	rcpt "github.com/NpoolPlatform/message/npool/appuser/mgr/v2/recaptcha"
	sm "github.com/NpoolPlatform/message/npool/appuser/mgr/v2/signmethod"
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
			JSON("signup_methods", []string{}).
			Optional().
			Default(func() []string {
				return []string{
					sm.SignMethodType_Mobile.String(),
					sm.SignMethodType_Email.String(),
				}
			}),
		field.
			JSON("extern_signin_methods", []string{}).
			Optional().
			Default(func() []string {
				return []string{}
			}),
		field.
			String("recaptcha_method").
			Optional().
			Default(rcpt.RecaptchaType_GoogleRecaptchaV3.String()),
		field.
			Bool("kyc_enable").
			Optional().
			Default(false),
		field.
			Bool("signin_verify_enable").
			Optional().
			Default(false),
		field.
			Bool("invitation_code_must").
			Optional().
			Default(false),
	}
}

// Edges of the AppControl.
func (AppControl) Edges() []ent.Edge {
	return nil
}
