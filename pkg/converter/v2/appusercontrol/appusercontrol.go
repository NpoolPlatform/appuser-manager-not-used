package appusercontrol

import (
	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent"
	npool "github.com/NpoolPlatform/message/npool/appuser/mgr/v2/appusercontrol"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
)

func Ent2Grpc(row *ent.AppUserControl) *npool.AppUserControl {
	if row == nil {
		return nil
	}

	return &npool.AppUserControl{
		ID:                 row.ID.String(),
		AppID:              row.AppID.String(),
		UserID:             row.UserID.String(),
		GoogleAuthVerified: row.GoogleAuthenticationVerified,
		SigninVerifyType:   basetypes.SignMethod(basetypes.SignMethod_value[row.SigninVerifyType]),
		Kol:                row.Kol,
		KolConfirmed:       row.KolConfirmed,
	}
}

func Ent2GrpcMany(rows []*ent.AppUserControl) []*npool.AppUserControl {
	infos := []*npool.AppUserControl{}
	for _, row := range rows {
		infos = append(infos, Ent2Grpc(row))
	}
	return infos
}
