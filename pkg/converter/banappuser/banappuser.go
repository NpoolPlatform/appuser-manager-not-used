package banappuser

import (
	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent"
	npool "github.com/NpoolPlatform/message/npool/appuser/mgr/v2/banappuser"
)

func Ent2Grpc(row *ent.BanAppUser) *npool.BanAppUser {
	if row == nil {
		return nil
	}

	return &npool.BanAppUser{
		ID:      row.ID.String(),
		AppID:   row.AppID.String(),
		UserID:  row.UserID.String(),
		Message: row.Message,
	}
}

func Ent2GrpcMany(rows []*ent.BanAppUser) []*npool.BanAppUser {
	infos := []*npool.BanAppUser{}
	for _, row := range rows {
		infos = append(infos, Ent2Grpc(row))
	}
	return infos
}
