package appusercontrol

import (
	"context"
	"fmt"

	npool "github.com/NpoolPlatform/message/npool/appusermgr"

	constant "github.com/NpoolPlatform/appuser-manager/pkg/const"
	db "github.com/NpoolPlatform/appuser-manager/pkg/db"
	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent"
	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent/appusercontrol"

	"github.com/google/uuid"
)

func validateAppUserControl(info *npool.AppUserControl) error {
	if _, err := uuid.Parse(info.GetAppID()); err != nil {
		return fmt.Errorf("invalid app id: %v", err)
	}
	if _, err := uuid.Parse(info.GetUserID()); err != nil {
		return fmt.Errorf("invalid user id: %v", err)
	}
	return nil
}

func dbRowToAppUserControl(row *ent.AppUserControl) *npool.AppUserControl {
	return &npool.AppUserControl{
		ID:                                 row.ID.String(),
		AppID:                              row.AppID.String(),
		UserID:                             row.UserID.String(),
		SigninVerifyByGoogleAuthentication: row.SigninVerifyByGoogleAuthentication,
		GoogleAuthenticationVerified:       row.GoogleAuthenticationVerified,
	}
}

func Create(ctx context.Context, in *npool.CreateAppUserControlRequest) (*npool.CreateAppUserControlResponse, error) {
	if err := validateAppUserControl(in.GetInfo()); err != nil {
		return nil, fmt.Errorf("invalid parameter: %v", err)
	}

	ctx, cancel := context.WithTimeout(ctx, constant.DBTimeout)
	defer cancel()

	cli, err := db.Client()
	if err != nil {
		return nil, fmt.Errorf("fail get db client: %v", err)
	}

	info, err := cli.
		AppUserControl.
		Create().
		SetAppID(uuid.MustParse(in.GetInfo().GetAppID())).
		SetUserID(uuid.MustParse(in.GetInfo().GetUserID())).
		SetSigninVerifyByGoogleAuthentication(in.GetInfo().GetSigninVerifyByGoogleAuthentication()).
		SetGoogleAuthenticationVerified(in.GetInfo().GetGoogleAuthenticationVerified()).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("fail create app user control: %v", err)
	}

	return &npool.CreateAppUserControlResponse{
		Info: dbRowToAppUserControl(info),
	}, nil
}

func Update(ctx context.Context, in *npool.UpdateAppUserControlRequest) (*npool.UpdateAppUserControlResponse, error) {
	id, err := uuid.Parse(in.GetInfo().GetID())
	if err != nil {
		return nil, fmt.Errorf("invalid app user control id: %v", err)
	}

	if err := validateAppUserControl(in.GetInfo()); err != nil {
		return nil, fmt.Errorf("invalid parameter: %v", err)
	}

	ctx, cancel := context.WithTimeout(ctx, constant.DBTimeout)
	defer cancel()

	cli, err := db.Client()
	if err != nil {
		return nil, fmt.Errorf("fail get db client: %v", err)
	}

	info, err := cli.
		AppUserControl.
		UpdateOneID(id).
		SetSigninVerifyByGoogleAuthentication(in.GetInfo().GetSigninVerifyByGoogleAuthentication()).
		SetGoogleAuthenticationVerified(in.GetInfo().GetGoogleAuthenticationVerified()).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("fail update app user control: %v", err)
	}

	return &npool.UpdateAppUserControlResponse{
		Info: dbRowToAppUserControl(info),
	}, nil
}

func Get(ctx context.Context, in *npool.GetAppUserControlRequest) (*npool.GetAppUserControlResponse, error) {
	id, err := uuid.Parse(in.GetID())
	if err != nil {
		return nil, fmt.Errorf("invalid app user control id: %v", err)
	}

	ctx, cancel := context.WithTimeout(ctx, constant.DBTimeout)
	defer cancel()

	cli, err := db.Client()
	if err != nil {
		return nil, fmt.Errorf("fail get db client: %v", err)
	}

	infos, err := cli.
		AppUserControl.
		Query().
		Where(
			appusercontrol.And(
				appusercontrol.ID(id),
			),
		).
		All(ctx)
	if err != nil {
		return nil, fmt.Errorf("fail query app user control: %v", err)
	}

	var myAppUserControl *npool.AppUserControl
	for _, info := range infos {
		myAppUserControl = dbRowToAppUserControl(info)
		break
	}

	return &npool.GetAppUserControlResponse{
		Info: myAppUserControl,
	}, nil
}

func GetByAppUser(ctx context.Context, in *npool.GetAppUserControlByAppUserRequest) (*npool.GetAppUserControlByAppUserResponse, error) {
	appID, err := uuid.Parse(in.GetAppID())
	if err != nil {
		return nil, fmt.Errorf("invalid app id: %v", err)
	}

	userID, err := uuid.Parse(in.GetUserID())
	if err != nil {
		return nil, fmt.Errorf("invalid user id: %v", err)
	}

	ctx, cancel := context.WithTimeout(ctx, constant.DBTimeout)
	defer cancel()

	cli, err := db.Client()
	if err != nil {
		return nil, fmt.Errorf("fail get db client: %v", err)
	}

	infos, err := cli.
		AppUserControl.
		Query().
		Where(
			appusercontrol.And(
				appusercontrol.AppID(appID),
				appusercontrol.UserID(userID),
			),
		).
		All(ctx)
	if err != nil {
		return nil, fmt.Errorf("fail query app user control: %v", err)
	}

	var appUserControl *npool.AppUserControl
	for _, info := range infos {
		appUserControl = dbRowToAppUserControl(info)
		break
	}

	return &npool.GetAppUserControlByAppUserResponse{
		Info: appUserControl,
	}, nil
}
