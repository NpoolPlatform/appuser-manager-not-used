package appuser

import (
	"context"
	"fmt"
	"time"

	npool "github.com/NpoolPlatform/message/npool/appusermgr"

	constant "github.com/NpoolPlatform/appuser-manager/pkg/const"
	"github.com/NpoolPlatform/appuser-manager/pkg/db"
	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent"
	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent/appuser"

	"github.com/google/uuid"
)

func validateAppUser(info *npool.AppUser) error {
	if _, err := uuid.Parse(info.GetAppID()); err != nil {
		return fmt.Errorf("invalid app id: %v", err)
	}
	return nil
}

func dbRowToAppUser(row *ent.AppUser) *npool.AppUser {
	return &npool.AppUser{
		ID:            row.ID.String(),
		AppID:         row.AppID.String(),
		EmailAddress:  row.EmailAddress,
		PhoneNO:       row.PhoneNo,
		ImportFromApp: row.ImportFromApp.String(),
		CreateAt:      row.CreatedAt,
	}
}

func Create(ctx context.Context, in *npool.CreateAppUserRequest) (*npool.CreateAppUserResponse, error) {
	if err := validateAppUser(in.GetInfo()); err != nil {
		return nil, fmt.Errorf("invalid parameter: %v", err)
	}

	importApp, err := uuid.Parse(in.GetInfo().GetImportFromApp())
	if err != nil {
		importApp = uuid.UUID{}
	}

	ctx, cancel := context.WithTimeout(ctx, constant.DBTimeout)
	defer cancel()

	cli, err := db.Client()
	if err != nil {
		return nil, fmt.Errorf("fail get db client: %v", err)
	}
	id := uuid.New()
	if in.GetInfo().GetID() != "" {
		id = uuid.MustParse(in.GetInfo().GetID())
	}
	info, err := cli.
		AppUser.
		Create().
		SetID(id).
		SetAppID(uuid.MustParse(in.GetInfo().GetAppID())).
		SetEmailAddress(in.GetInfo().GetEmailAddress()).
		SetPhoneNo(in.GetInfo().GetPhoneNO()).
		SetImportFromApp(importApp).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("fail create app user: %v", err)
	}

	return &npool.CreateAppUserResponse{
		Info: dbRowToAppUser(info),
	}, nil
}

func CreateRevert(ctx context.Context, in *npool.CreateAppUserRequest) (*npool.CreateAppUserResponse, error) {
	id, err := uuid.Parse(in.GetInfo().GetID())
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(ctx, constant.DBTimeout)
	defer cancel()

	cli, err := db.Client()
	if err != nil {
		return nil, fmt.Errorf("fail get db client: %v", err)
	}
	err = cli.
		AppUser.UpdateOneID(id).SetDeletedAt(uint32(time.Now().Unix())).Exec(ctx)
	if err != nil {
		return nil, fmt.Errorf("fail create app user: %v", err)
	}

	return &npool.CreateAppUserResponse{}, nil
}

func Update(ctx context.Context, in *npool.UpdateAppUserRequest) (*npool.UpdateAppUserResponse, error) {
	id, err := uuid.Parse(in.GetInfo().GetID())
	if err != nil {
		return nil, fmt.Errorf("invalid app user id: %v", err)
	}

	if err := validateAppUser(in.GetInfo()); err != nil {
		return nil, fmt.Errorf("invalid parameter: %v", err)
	}

	ctx, cancel := context.WithTimeout(ctx, constant.DBTimeout)
	defer cancel()

	cli, err := db.Client()
	if err != nil {
		return nil, fmt.Errorf("fail get db client: %v", err)
	}

	info, err := cli.
		AppUser.
		UpdateOneID(id).
		SetEmailAddress(in.GetInfo().GetEmailAddress()).
		SetPhoneNo(in.GetInfo().GetPhoneNO()).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("fail update app user: %v", err)
	}

	return &npool.UpdateAppUserResponse{
		Info: dbRowToAppUser(info),
	}, nil
}

func Get(ctx context.Context, in *npool.GetAppUserRequest) (*npool.GetAppUserResponse, error) {
	id, err := uuid.Parse(in.GetID())
	if err != nil {
		return nil, fmt.Errorf("invalid app user id: %v", err)
	}

	ctx, cancel := context.WithTimeout(ctx, constant.DBTimeout)
	defer cancel()

	cli, err := db.Client()
	if err != nil {
		return nil, fmt.Errorf("fail get db client: %v", err)
	}

	infos, err := cli.
		AppUser.
		Query().
		Where(
			appuser.And(
				appuser.ID(id),
			),
		).
		All(ctx)
	if err != nil {
		return nil, fmt.Errorf("fail query app user: %v", err)
	}

	var myAppUser *npool.AppUser
	for _, info := range infos {
		myAppUser = dbRowToAppUser(info)
		break
	}

	return &npool.GetAppUserResponse{
		Info: myAppUser,
	}, nil
}

func GetByApp(ctx context.Context, in *npool.GetAppUsersByAppRequest) (*npool.GetAppUsersByAppResponse, error) {
	appID, err := uuid.Parse(in.GetAppID())
	if err != nil {
		return nil, fmt.Errorf("invalid app id: %v", err)
	}

	ctx, cancel := context.WithTimeout(ctx, constant.DBTimeout)
	defer cancel()

	cli, err := db.Client()
	if err != nil {
		return nil, fmt.Errorf("fail get db client: %v", err)
	}

	infos, err := cli.
		AppUser.
		Query().
		Where(
			appuser.AppID(appID),
		).
		All(ctx)
	if err != nil {
		return nil, fmt.Errorf("fail query app user: %v", err)
	}

	appUsers := []*npool.AppUser{}
	for _, info := range infos {
		appUsers = append(appUsers, dbRowToAppUser(info))
	}

	return &npool.GetAppUsersByAppResponse{
		Infos: appUsers,
	}, nil
}

func GetByAppAccount(ctx context.Context, in *npool.GetAppUserByAppAccountRequest) (*npool.GetAppUserByAppAccountResponse, error) {
	appID, err := uuid.Parse(in.GetAppID())
	if err != nil {
		return nil, fmt.Errorf("invalid app id: %v", err)
	}

	ctx, cancel := context.WithTimeout(ctx, constant.DBTimeout)
	defer cancel()

	cli, err := db.Client()
	if err != nil {
		return nil, fmt.Errorf("fail get db client: %v", err)
	}

	infos, err := cli.
		AppUser.
		Query().
		Where(
			appuser.And(
				appuser.AppID(appID),
				appuser.Or(
					appuser.EmailAddress(in.GetAccount()),
					appuser.PhoneNo(in.GetAccount()),
				),
			),
		).
		All(ctx)
	if err != nil {
		return nil, fmt.Errorf("fail query app user: %v", err)
	}

	var myAppUser *npool.AppUser
	for _, info := range infos {
		myAppUser = dbRowToAppUser(info)
		break
	}

	return &npool.GetAppUserByAppAccountResponse{
		Info: myAppUser,
	}, nil
}

func GetByAppUser(ctx context.Context, in *npool.GetAppUserByAppUserRequest) (*npool.GetAppUserByAppUserResponse, error) {
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
		AppUser.
		Query().
		Where(
			appuser.And(
				appuser.AppID(appID),
				appuser.ID(userID),
			),
		).
		All(ctx)
	if err != nil {
		return nil, fmt.Errorf("fail query app user: %v", err)
	}

	var myAppUser *npool.AppUser
	for _, info := range infos {
		myAppUser = dbRowToAppUser(info)
		break
	}

	return &npool.GetAppUserByAppUserResponse{
		Info: myAppUser,
	}, nil
}
