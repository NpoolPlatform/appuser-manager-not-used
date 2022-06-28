package appusersecret

import (
	"context"
	"fmt"
	"time"

	npool "github.com/NpoolPlatform/message/npool/appusermgr"

	constant "github.com/NpoolPlatform/appuser-manager/pkg/const"
	db "github.com/NpoolPlatform/appuser-manager/pkg/db"
	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent"
	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent/appusersecret"
	encrypt "github.com/NpoolPlatform/appuser-manager/pkg/middleware/encrypt"

	"github.com/google/uuid"
)

func validateAppUserSecret(info *npool.AppUserSecret) error {
	if _, err := uuid.Parse(info.GetAppID()); err != nil {
		return fmt.Errorf("invalid app id: %v", err)
	}
	if _, err := uuid.Parse(info.GetUserID()); err != nil {
		return fmt.Errorf("invalid user id: %v", err)
	}
	if info.GetPasswordHash() == "" {
		return fmt.Errorf("invalid password hash")
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
		return nil, fmt.Errorf("invalid parameter: %v", err)
	}

	ctx, cancel := context.WithTimeout(ctx, constant.DBTimeout)
	defer cancel()

	cli, err := db.Client()
	if err != nil {
		return nil, fmt.Errorf("fail get db client: %v", err)
	}

	salt := encrypt.Salt()
	password, err := encrypt.EncryptWithSalt(in.GetInfo().GetPasswordHash(), salt)
	if err != nil {
		return nil, fmt.Errorf("fail get encrypted password: %v", err)
	}

	info, err := cli.
		AppUserSecret.
		Create().
		SetAppID(uuid.MustParse(in.GetInfo().GetAppID())).
		SetUserID(uuid.MustParse(in.GetInfo().GetUserID())).
		SetPasswordHash(password).
		SetSalt(salt).
		SetGoogleSecret(in.GetInfo().GetGoogleSecret()).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("fail create app user secret: %v", err)
	}

	return &npool.CreateAppUserSecretResponse{
		Info: dbRowToAppUserSecret(info),
	}, nil
}

func CreateRevert(ctx context.Context, in *npool.CreateAppUserSecretRequest) (*npool.CreateAppUserSecretResponse, error) {
	if err := validateAppUserSecret(in.GetInfo()); err != nil {
		return nil, fmt.Errorf("invalid parameter: %v", err)
	}

	ctx, cancel := context.WithTimeout(ctx, constant.DBTimeout)
	defer cancel()

	cli, err := db.Client()
	if err != nil {
		return nil, fmt.Errorf("fail get db client: %v", err)
	}
	_, err = cli.
		AppUserSecret.
		Update().
		SetDeletedAt(uint32(time.Now().Unix())).
		Where(
			appusersecret.AppID(uuid.MustParse(in.GetInfo().GetAppID())),
			appusersecret.UserID(uuid.MustParse(in.GetInfo().GetUserID())),
		).
		Save(ctx)

	if err != nil {
		return nil, fmt.Errorf("fail delete app user secret: %v", err)
	}

	return &npool.CreateAppUserSecretResponse{}, nil
}

func Update(ctx context.Context, in *npool.UpdateAppUserSecretRequest) (*npool.UpdateAppUserSecretResponse, error) {
	id, err := uuid.Parse(in.GetInfo().GetID())
	if err != nil {
		return nil, fmt.Errorf("invalid app user secret id: %v", err)
	}

	if err := validateAppUserSecret(in.GetInfo()); err != nil {
		return nil, fmt.Errorf("invalid parameter: %v", err)
	}

	ctx, cancel := context.WithTimeout(ctx, constant.DBTimeout)
	defer cancel()

	cli, err := db.Client()
	if err != nil {
		return nil, fmt.Errorf("fail get db client: %v", err)
	}

	salt := in.GetInfo().GetSalt()
	password := in.GetInfo().GetPasswordHash()

	if salt == "" {
		salt = encrypt.Salt()
		password, err = encrypt.EncryptWithSalt(in.GetInfo().GetPasswordHash(), salt)
		if err != nil {
			return nil, fmt.Errorf("fail get encrypted password: %v", err)
		}
	}

	info, err := cli.
		AppUserSecret.
		UpdateOneID(id).
		SetPasswordHash(password).
		SetSalt(salt).
		SetGoogleSecret(in.GetInfo().GetGoogleSecret()).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("fail update app user secret: %v", err)
	}

	return &npool.UpdateAppUserSecretResponse{
		Info: dbRowToAppUserSecret(info),
	}, nil
}

func Get(ctx context.Context, in *npool.GetAppUserSecretRequest) (*npool.GetAppUserSecretResponse, error) {
	id, err := uuid.Parse(in.GetID())
	if err != nil {
		return nil, fmt.Errorf("invalid app user secret id: %v", err)
	}

	ctx, cancel := context.WithTimeout(ctx, constant.DBTimeout)
	defer cancel()

	cli, err := db.Client()
	if err != nil {
		return nil, fmt.Errorf("fail get db client: %v", err)
	}

	infos, err := cli.
		AppUserSecret.
		Query().
		Where(
			appusersecret.And(
				appusersecret.ID(id),
			),
		).
		All(ctx)
	if err != nil {
		return nil, fmt.Errorf("fail query app user secret: %v", err)
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
		return nil, fmt.Errorf("fail query app user secret: %v", err)
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
