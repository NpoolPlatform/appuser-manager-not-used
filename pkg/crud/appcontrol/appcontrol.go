package appcontrol

import (
	"context"

	npool "github.com/NpoolPlatform/message/npool/appusermgr"

	constant "github.com/NpoolPlatform/appuser-manager/pkg/const"
	db "github.com/NpoolPlatform/appuser-manager/pkg/db"
	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent"
	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent/appcontrol"

	"github.com/google/uuid"

	"golang.org/x/xerrors"
)

func validateAppControl(info *npool.AppControl) error {
	if _, err := uuid.Parse(info.GetAppID()); err != nil {
		return xerrors.Errorf("invalid app id: %v", err)
	}
	return nil
}

func dbRowToAppControl(row *ent.AppControl) *npool.AppControl {
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

func Create(ctx context.Context, in *npool.CreateAppControlRequest) (*npool.CreateAppControlResponse, error) {
	if err := validateAppControl(in.GetInfo()); err != nil {
		return nil, xerrors.Errorf("invalid parameter: %v", err)
	}

	ctx, cancel := context.WithTimeout(ctx, constant.DBTimeout)
	defer cancel()

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	info, err := cli.
		AppControl.
		Create().
		SetAppID(uuid.MustParse(in.GetInfo().GetAppID())).
		SetSignupMethods(in.GetInfo().GetSignupMethods()).
		SetExternSigninMethods(in.GetInfo().GetExternSigninMethods()).
		SetRecaptchaMethod(in.GetInfo().GetRecaptchaMethod()).
		SetKycEnable(in.GetInfo().GetKycEnable()).
		SetSigninVerifyEnable(in.GetInfo().GetSigninVerifyEnable()).
		SetInvitationCodeMust(in.GetInfo().GetInvitationCodeMust()).
		Save(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail create app control: %v", err)
	}

	return &npool.CreateAppControlResponse{
		Info: dbRowToAppControl(info),
	}, nil
}

func Update(ctx context.Context, in *npool.UpdateAppControlRequest) (*npool.UpdateAppControlResponse, error) {
	id, err := uuid.Parse(in.GetInfo().GetID())
	if err != nil {
		return nil, xerrors.Errorf("invalid app control id: %v", err)
	}

	if err := validateAppControl(in.GetInfo()); err != nil {
		return nil, xerrors.Errorf("invalid parameter: %v", err)
	}

	ctx, cancel := context.WithTimeout(ctx, constant.DBTimeout)
	defer cancel()

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	info, err := cli.
		AppControl.
		UpdateOneID(id).
		SetSignupMethods(in.GetInfo().GetSignupMethods()).
		SetExternSigninMethods(in.GetInfo().GetExternSigninMethods()).
		SetRecaptchaMethod(in.GetInfo().GetRecaptchaMethod()).
		SetKycEnable(in.GetInfo().GetKycEnable()).
		SetSigninVerifyEnable(in.GetInfo().GetSigninVerifyEnable()).
		SetInvitationCodeMust(in.GetInfo().GetInvitationCodeMust()).
		Save(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail update app control: %v", err)
	}

	return &npool.UpdateAppControlResponse{
		Info: dbRowToAppControl(info),
	}, nil
}

func Get(ctx context.Context, in *npool.GetAppControlRequest) (*npool.GetAppControlResponse, error) {
	id, err := uuid.Parse(in.GetID())
	if err != nil {
		return nil, xerrors.Errorf("invalid app control id: %v", err)
	}

	ctx, cancel := context.WithTimeout(ctx, constant.DBTimeout)
	defer cancel()

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	infos, err := cli.
		AppControl.
		Query().
		Where(
			appcontrol.ID(id),
		).
		All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail query app control: %v", err)
	}

	var myAppControl *npool.AppControl
	for _, info := range infos {
		myAppControl = dbRowToAppControl(info)
		break
	}

	return &npool.GetAppControlResponse{
		Info: myAppControl,
	}, nil
}

func GetByApp(ctx context.Context, in *npool.GetAppControlByAppRequest) (*npool.GetAppControlByAppResponse, error) {
	appID, err := uuid.Parse(in.GetAppID())
	if err != nil {
		return nil, xerrors.Errorf("invalid app id: %v", err)
	}

	ctx, cancel := context.WithTimeout(ctx, constant.DBTimeout)
	defer cancel()

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	infos, err := cli.
		AppControl.
		Query().
		Where(
			appcontrol.AppID(appID),
		).
		All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail query app control: %v", err)
	}

	var myAppControl *npool.AppControl
	for _, info := range infos {
		myAppControl = dbRowToAppControl(info)
		break
	}

	return &npool.GetAppControlByAppResponse{
		Info: myAppControl,
	}, nil
}
