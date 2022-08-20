package appuser

import (
	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent"
	npool "github.com/NpoolPlatform/message/npool/appuser/mgr/v2/appuser"
)

func Ent2Grpc(row *ent.AppUser) *npool.AppUser {
	if row == nil {
		return nil
	}

	return &npool.AppUser{
		ID:            row.ID.String(),
		AppID:         row.AppID.String(),
		EmailAddress:  row.EmailAddress,
		PhoneNO:       row.PhoneNo,
		ImportFromApp: row.ImportFromApp.String(),
	}
}

func Ent2GrpcMany(rows []*ent.AppUser) []*npool.AppUser {
	infos := []*npool.AppUser{}
	for _, row := range rows {
		infos = append(infos, Ent2Grpc(row))
	}
	return infos
}
