package banapp

import (
	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent"
	npool "github.com/NpoolPlatform/message/npool/appuser/mgr/v2/banapp"
)

func Ent2Grpc(row *ent.BanApp) *npool.BanApp {
	if row == nil {
		return nil
	}

	return &npool.BanApp{
		ID:      row.ID.String(),
		AppID:   row.AppID.String(),
		Message: row.Message,
	}
}

func Ent2GrpcMany(rows []*ent.BanApp) []*npool.BanApp {
	infos := []*npool.BanApp{}
	for _, row := range rows {
		infos = append(infos, Ent2Grpc(row))
	}
	return infos
}
