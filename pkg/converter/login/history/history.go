package history

import (
	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent"
	npool "github.com/NpoolPlatform/message/npool/appuser/mgr/v2/login/history"
)

func Ent2Grpc(row *ent.LoginHistory) *npool.History {
	if row == nil {
		return nil
	}

	return &npool.History{
		ID:        row.ID.String(),
		AppID:     row.AppID.String(),
		UserID:    row.UserID.String(),
		ClientIP:  row.ClientIP,
		UserAgent: row.UserAgent,
		Location:  row.Location,
		CreatedAt: row.CreatedAt,
	}
}

func Ent2GrpcMany(rows []*ent.LoginHistory) []*npool.History {
	infos := []*npool.History{}
	for _, row := range rows {
		infos = append(infos, Ent2Grpc(row))
	}
	return infos
}
