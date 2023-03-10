package appcontrol

import (
	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent"
	npool "github.com/NpoolPlatform/message/npool/appuser/mgr/v2/appcontrol"
	rcpt "github.com/NpoolPlatform/message/npool/appuser/mgr/v2/recaptcha"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
)

func Ent2Grpc(row *ent.AppControl) *npool.AppControl {
	if row == nil {
		return nil
	}

	methods := []basetypes.SignMethod{}
	for _, m := range row.SignupMethods {
		methods = append(methods, basetypes.SignMethod(basetypes.SignMethod_value[m]))
	}

	emethods := []basetypes.SignMethod{}
	for _, m := range row.ExternSigninMethods {
		emethods = append(emethods, basetypes.SignMethod(basetypes.SignMethod_value[m]))
	}

	return &npool.AppControl{
		ID:                       row.ID.String(),
		AppID:                    row.AppID.String(),
		SignupMethods:            methods,
		ExtSigninMethods:         emethods,
		RecaptchaMethod:          rcpt.RecaptchaType(rcpt.RecaptchaType_value[row.RecaptchaMethod]),
		KycEnable:                row.KycEnable,
		SigninVerifyEnable:       row.SigninVerifyEnable,
		InvitationCodeMust:       row.InvitationCodeMust,
		CreateInvitationCodeWhen: npool.CreateInvitationCodeWhen(npool.CreateInvitationCodeWhen_value[row.CreateInvitationCodeWhen]),
		MaxTypedCouponsPerOrder:  row.MaxTypedCouponsPerOrder,
		UnderMaintenance:         row.UnderMaintenance,
		CommitButtons:            row.CommitButtons,
	}
}

func Ent2GrpcMany(rows []*ent.AppControl) []*npool.AppControl {
	infos := []*npool.AppControl{}
	for _, row := range rows {
		infos = append(infos, Ent2Grpc(row))
	}
	return infos
}
