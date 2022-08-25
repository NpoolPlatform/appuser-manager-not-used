package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/appuser-manager/pkg/db/mixin"

	"github.com/google/uuid"

	npool "github.com/NpoolPlatform/message/npool/appuser/mgr/v2/kyc"
	reviewpb "github.com/NpoolPlatform/message/npool/review/mgr/v2"
)

// Kyc holds the schema definition for the Kyc entity.
type Kyc struct {
	ent.Schema
}

func (Kyc) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
	}
}

// Fields of the Kyc.
func (Kyc) Fields() []ent.Field {
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
			UUID("user_id", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.UUID{}
			}),
		field.
			String("document_type").
			Optional().
			Default(npool.KycDocumentType_DefaultKycDocumentType.String()),
		field.
			String("id_number").
			Optional().
			Default(""),
		field.
			String("front_img").
			Optional().
			Default(""),
		field.
			String("back_img").
			Optional().
			Default(""),
		field.
			String("selfie_img").
			Optional().
			Default(""),
		field.
			String("entity_type").
			Optional().
			Default(npool.KycEntityType_Individual.String()),
		field.
			UUID("review_id", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.UUID{}
			}),
		field.
			String("review_state").
			Optional().
			Default(reviewpb.ReviewState_Wait.String()),
		field.
			String("review_message").
			Optional().
			Default(""),
	}
}

// Edges of the Kyc.
func (Kyc) Edges() []ent.Edge {
	return nil
}
