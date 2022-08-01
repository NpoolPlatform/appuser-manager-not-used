package appcontrol

import (
	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent"
	npool "github.com/NpoolPlatform/message/npool/appuser/mgr/v2/appcontrol"
)

func Ent2Grpc(row *ent.AppControl) *npool.AppControl {
	if row == nil {
		return nil
	}

	return &npool.AppControl{
		ID:                  row.ID.String(),
		AppID:               row.AppID.String(),
		SignupMethods:       row.SignupMethods,
		ExternSigninMethods: row.ExternSigninMethods,
		RecaptchaMethod:     row.RecaptchaMethod,
		KycEnable:           row.KycEnable,
		SigninVerifyEnable:  row.SigninVerifyEnable,
		InvitationCodeMust:  row.InvitationCodeMust,
	}
}

func Ent2GrpcMany(rows []*ent.AppControl) []*npool.AppControl {
	infos := []*npool.AppControl{}
	for _, row := range rows {
		infos = append(infos, Ent2Grpc(row))
	}
	return infos
}
