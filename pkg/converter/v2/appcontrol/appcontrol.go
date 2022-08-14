package appcontrol

import (
	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent"
	npool "github.com/NpoolPlatform/message/npool/appuser/mgr/v2/appcontrol"
	rcpt "github.com/NpoolPlatform/message/npool/appuser/mgr/v2/recaptcha"
	sm "github.com/NpoolPlatform/message/npool/appuser/mgr/v2/signmethod"
)

func Ent2Grpc(row *ent.AppControl) *npool.AppControl {
	if row == nil {
		return nil
	}

	methods := []sm.SignMethodType{}
	for _, m := range row.SignupMethods {
		methods = append(methods, sm.SignMethodType(sm.SignMethodType_value[m]))
	}

	emethods := []sm.SignMethodType{}
	for _, m := range row.ExternSigninMethods {
		emethods = append(emethods, sm.SignMethodType(sm.SignMethodType_value[m]))
	}

	return &npool.AppControl{
		ID:                 row.ID.String(),
		AppID:              row.AppID.String(),
		SignupMethods:      methods,
		ExtSigninMethods:   emethods,
		RecaptchaMethod:    rcpt.RecaptchaType(rcpt.RecaptchaType_value[row.RecaptchaMethod]),
		KycEnable:          row.KycEnable,
		SigninVerifyEnable: row.SigninVerifyEnable,
		InvitationCodeMust: row.InvitationCodeMust,
	}
}

func Ent2GrpcMany(rows []*ent.AppControl) []*npool.AppControl {
	infos := []*npool.AppControl{}
	for _, row := range rows {
		infos = append(infos, Ent2Grpc(row))
	}
	return infos
}
