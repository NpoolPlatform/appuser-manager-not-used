package appusersecret

import (
	"context"

	npool "github.com/NpoolPlatform/message/npool/appusermgr"

	constant "github.com/NpoolPlatform/appuser-manager/pkg/const"
	db "github.com/NpoolPlatform/appuser-manager/pkg/db"
	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent"
	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent/appusersecret"

	"github.com/google/uuid"

	"golang.org/x/xerrors"
)

func validateAppUserSecret(info *npool.AppUserSecret) error {
	if _, err := uuid.Parse(info.GetAppID()); err != nil {
		return xerrors.Errorf("invalid app id: %v", err)
	}
	if _, err := uuid.Parse(info.GetUserID()); err != nil {
		return xerrors.Errorf("invalid user id: %v", err)
	}
	if info.GetPasswordHash() == "" {
		return xerrors.Errorf("invalid password hash")
	}
	if info.GetSalt() == "" {
		return xerrors.Errorf("invalid password salt")
	}
	return nil
}

func dbRowToAppUserSecret(row *ent.AppUserSecret) *npool.AppUserSecret {
	return &npool.AppUserSecret{
		ID:           row.ID.String(),
		AppID:        row.AppID.String(),
		UserID:       row.UserID.String(),
		PasswordHash: row.PasswordHash,
		Salt:         row.Salt,
		GoogleSecret: row.GoogleSecret,
	}
}

func Create(ctx context.Context, in *npool.CreateAppUserSecretRequest) (*npool.CreateAppUserSecretResponse, error) {
	if err := validateAppUserSecret(in.GetInfo()); err != nil {
		return nil, xerrors.Errorf("invalid parameter: %v", err)
	}

	ctx, cancel := context.WithTimeout(ctx, constant.DBTimeout)
	defer cancel()

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	info, err := cli.
		AppUserSecret.
		Create().
		SetAppID(uuid.MustParse(in.GetInfo().GetAppID())).
		SetUserID(uuid.MustParse(in.GetInfo().GetUserID())).
		SetPasswordHash(in.GetInfo().GetPasswordHash()).
		SetSalt(in.GetInfo().GetSalt()).
		SetGoogleSecret(in.GetInfo().GetGoogleSecret()).
		Save(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail create app user secret: %v", err)
	}

	return &npool.CreateAppUserSecretResponse{
		Info: dbRowToAppUserSecret(info),
	}, nil
}

func Update(ctx context.Context, in *npool.UpdateAppUserSecretRequest) (*npool.UpdateAppUserSecretResponse, error) {
	id, err := uuid.Parse(in.GetInfo().GetID())
	if err != nil {
		return nil, xerrors.Errorf("invalid app user secret id: %v", err)
	}

	if err := validateAppUserSecret(in.GetInfo()); err != nil {
		return nil, xerrors.Errorf("invalid parameter: %v", err)
	}

	ctx, cancel := context.WithTimeout(ctx, constant.DBTimeout)
	defer cancel()

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	info, err := cli.
		AppUserSecret.
		UpdateOneID(id).
		SetPasswordHash(in.GetInfo().GetPasswordHash()).
		SetSalt(in.GetInfo().GetSalt()).
		SetGoogleSecret(in.GetInfo().GetGoogleSecret()).
		Save(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail update app user secret: %v", err)
	}

	return &npool.UpdateAppUserSecretResponse{
		Info: dbRowToAppUserSecret(info),
	}, nil
}

func Get(ctx context.Context, in *npool.GetAppUserSecretRequest) (*npool.GetAppUserSecretResponse, error) {
	id, err := uuid.Parse(in.GetID())
	if err != nil {
		return nil, xerrors.Errorf("invalid app user secret id: %v", err)
	}

	ctx, cancel := context.WithTimeout(ctx, constant.DBTimeout)
	defer cancel()

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	infos, err := cli.
		AppUserSecret.
		Query().
		Where(
			appusersecret.And(
				appusersecret.ID(id),
				appusersecret.DeleteAt(0),
			),
		).
		All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail query app user secret: %v", err)
	}

	var myAppUserSecret *npool.AppUserSecret
	for _, info := range infos {
		myAppUserSecret = dbRowToAppUserSecret(info)
		break
	}

	return &npool.GetAppUserSecretResponse{
		Info: myAppUserSecret,
	}, nil
}

func GetByAppUser(ctx context.Context, in *npool.GetAppUserSecretByAppUserRequest) (*npool.GetAppUserSecretByAppUserResponse, error) {
	appID, err := uuid.Parse(in.GetAppID())
	if err != nil {
		return nil, xerrors.Errorf("invalid app id: %v", err)
	}

	userID, err := uuid.Parse(in.GetUserID())
	if err != nil {
		return nil, xerrors.Errorf("invalid user id: %v", err)
	}

	ctx, cancel := context.WithTimeout(ctx, constant.DBTimeout)
	defer cancel()

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	infos, err := cli.
		AppUserSecret.
		Query().
		Where(
			appusersecret.And(
				appusersecret.AppID(appID),
				appusersecret.UserID(userID),
			),
		).
		All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail query app user secret: %v", err)
	}

	var appUserSecret *npool.AppUserSecret
	for _, info := range infos {
		appUserSecret = dbRowToAppUserSecret(info)
		break
	}

	return &npool.GetAppUserSecretByAppUserResponse{
		Info: appUserSecret,
	}, nil
}