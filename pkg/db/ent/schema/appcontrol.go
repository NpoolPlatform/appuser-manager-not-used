package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"

	"github.com/google/uuid"
)

// AppControl holds the schema definition for the AppControl entity.
type AppControl struct {
	ent.Schema
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
		field.Uint32("create_at").
			DefaultFunc(func() uint32 {
				return uint32(time.Now().Unix())
			}),
		field.Uint32("update_at").
			DefaultFunc(func() uint32 {
				return uint32(time.Now().Unix())
			}).
			UpdateDefault(func() uint32 {
				return uint32(time.Now().Unix())
			}),
		field.Uint32("delete_at").
			DefaultFunc(func() uint32 {
				return 0
			}),
	}
}

// Edges of the AppControl.
func (AppControl) Edges() []ent.Edge {
	return nil
}
