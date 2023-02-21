package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/appuser-manager/pkg/db/mixin"

	"github.com/google/uuid"

	npool "github.com/NpoolPlatform/message/npool/appuser/mgr/v2/appcontrol"
	rcpt "github.com/NpoolPlatform/message/npool/appuser/mgr/v2/recaptcha"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
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
					basetypes.SignMethod_Mobile.String(),
					basetypes.SignMethod_Email.String(),
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
		field.
			String("create_invitation_code_when").
			Optional().
			Default(npool.CreateInvitationCodeWhen_DefaultWhen.String()),
		field.
			Uint32("max_typed_coupons_per_order").
			Optional().
			Default(1),
	}
}

// Edges of the AppControl.
func (AppControl) Edges() []ent.Edge {
	return nil
}
