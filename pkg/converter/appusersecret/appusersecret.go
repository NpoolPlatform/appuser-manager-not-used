package appusersecret

import (
	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent"
	npool "github.com/NpoolPlatform/message/npool/appuser/mgr/v2/appusersecret"
)

func Ent2Grpc(row *ent.AppUserSecret) *npool.AppUserSecret {
	if row == nil {
		return nil
	}

	return &npool.AppUserSecret{
		Salt:         row.Salt,
		GoogleSecret: row.GoogleSecret,
		ID:           row.ID.String(),
		AppID:        row.AppID.String(),
		UserID:       row.UserID.String(),
		PasswordHash: row.PasswordHash,
	}
}

func Ent2GrpcMany(rows []*ent.AppUserSecret) []*npool.AppUserSecret {
	infos := []*npool.AppUserSecret{}
	for _, row := range rows {
		infos = append(infos, Ent2Grpc(row))
	}
	return infos
}
