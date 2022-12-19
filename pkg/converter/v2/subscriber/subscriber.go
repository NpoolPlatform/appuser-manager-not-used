package subscriber

import (
	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent"
	npool "github.com/NpoolPlatform/message/npool/appuser/mgr/v2/subscriber"
)

func Ent2Grpc(row *ent.Subscriber) *npool.Subscriber {
	if row == nil {
		return nil
	}

	return &npool.Subscriber{
		ID:           row.ID.String(),
		AppID:        row.AppID.String(),
		EmailAddress: row.EmailAddress,
		Registered:   row.Registered,
		CreatedAt:    row.CreatedAt,
		UpdatedAt:    row.UpdatedAt,
	}
}

func Ent2GrpcMany(rows []*ent.Subscriber) []*npool.Subscriber {
	infos := []*npool.Subscriber{}
	for _, row := range rows {
		infos = append(infos, Ent2Grpc(row))
	}
	return infos
}
