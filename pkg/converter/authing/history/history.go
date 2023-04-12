package history

import (
	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent"
	npool "github.com/NpoolPlatform/message/npool/appuser/mgr/v2/authing/history"
)

func Ent2Grpc(row *ent.AuthHistory) *npool.History {
	if row == nil {
		return nil
	}

	return &npool.History{
		ID:        row.ID.String(),
		AppID:     row.AppID.String(),
		UserID:    row.UserID.String(),
		Resource:  row.Resource,
		Method:    row.Method,
		Allowed:   row.Allowed,
		CreatedAt: row.CreatedAt,
	}
}

func Ent2GrpcMany(rows []*ent.AuthHistory) []*npool.History {
	infos := []*npool.History{}
	for _, row := range rows {
		infos = append(infos, Ent2Grpc(row))
	}
	return infos
}
