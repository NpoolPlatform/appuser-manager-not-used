package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/appuser-manager/pkg/db/mixin"

	"github.com/google/uuid"

	sm "github.com/NpoolPlatform/message/npool/appuser/mgr/v2/signmethod"
)

// App holds the schema definition for the App entity.
type App struct {
	ent.Schema
}

func (App) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
	}
}

// Fields of the App.
func (App) Fields() []ent.Field {
	return []ent.Field{
		field.
			UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique(),
		field.
			UUID("created_by", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.UUID{}
			}),
		field.
			String("name").
			Optional().
			Unique(),
		field.
			String("logo").
			Optional().
			Default(""),
		field.
			String("description").
			Optional().
			Default(""),
		field.
			String("signin_verify_type").
			Optional().
			Default(sm.SignMethodType_Email.String()),
	}
}

// Edges of the App.
func (App) Edges() []ent.Edge {
	return nil
}
