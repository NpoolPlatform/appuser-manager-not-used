package banappuser

import (
	"context"
	"fmt"
	"time"

	constant "github.com/NpoolPlatform/appuser-manager/pkg/const"
	npool "github.com/NpoolPlatform/message/npool/appusermgr"

	db "github.com/NpoolPlatform/appuser-manager/pkg/db"
	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent"
	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent/banappuser"

	"github.com/google/uuid"
)

func validateBanAppUser(info *npool.BanAppUser) error {
	if _, err := uuid.Parse(info.GetAppID()); err != nil {
		return fmt.Errorf("invalid app id: %v", err)
	}
	if _, err := uuid.Parse(info.GetUserID()); err != nil {
		return fmt.Errorf("invalid user id: %v", err)
	}
	return nil
}

func dbRowToBanAppUser(row *ent.BanAppUser) *npool.BanAppUser {
	return &npool.BanAppUser{
		ID:      row.ID.String(),
		AppID:   row.AppID.String(),
		UserID:  row.UserID.String(),
		Message: row.Message,
	}
}

func Create(ctx context.Context, in *npool.CreateBanAppUserRequest) (*npool.CreateBanAppUserResponse, error) {
	if err := validateBanAppUser(in.GetInfo()); err != nil {
		return nil, fmt.Errorf("invalid parameter: %v", err)
	}

	ctx, cancel := context.WithTimeout(ctx, constant.DBTimeout)
	defer cancel()

	cli, err := db.Client()
	if err != nil {
		return nil, fmt.Errorf("fail get db client: %v", err)
	}

	info, err := cli.
		BanAppUser.
		Create().
		SetAppID(uuid.MustParse(in.GetInfo().GetAppID())).
		SetUserID(uuid.MustParse(in.GetInfo().GetUserID())).
		SetMessage(in.GetInfo().GetMessage()).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("fail create ban app user: %v", err)
	}

	return &npool.CreateBanAppUserResponse{
		Info: dbRowToBanAppUser(info),
	}, nil
}

func Get(ctx context.Context, in *npool.GetBanAppUserRequest) (*npool.GetBanAppUserResponse, error) {
	id, err := uuid.Parse(in.GetID())
	if err != nil {
		return nil, fmt.Errorf("invalid id: %v", err)
	}

	ctx, cancel := context.WithTimeout(ctx, constant.DBTimeout)
	defer cancel()

	cli, err := db.Client()
	if err != nil {
		return nil, fmt.Errorf("fail get db client: %v", err)
	}

	infos, err := cli.
		BanAppUser.
		Query().
		Where(
			banappuser.And(
				banappuser.ID(id),
			),
		).
		All(ctx)
	if err != nil {
		return nil, fmt.Errorf("fail query ban app user: %v", err)
	}

	var banAppUser *npool.BanAppUser
	for _, info := range infos {
		banAppUser = dbRowToBanAppUser(info)
		break
	}

	return &npool.GetBanAppUserResponse{
		Info: banAppUser,
	}, nil
}

func GetByAppUser(ctx context.Context, in *npool.GetBanAppUserByAppUserRequest) (*npool.GetBanAppUserByAppUserResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, constant.DBTimeout)
	defer cancel()

	cli, err := db.Client()
	if err != nil {
		return nil, fmt.Errorf("fail get db client: %v", err)
	}

	appID, err := uuid.Parse(in.GetAppID())
	if err != nil {
		return nil, fmt.Errorf("invalid app id: %v", err)
	}

	infos, err := cli.
		BanAppUser.
		Query().
		Where(
			banappuser.And(
				banappuser.AppID(appID),
			),
		).
		All(ctx)
	if err != nil {
		return nil, fmt.Errorf("fail query ban app user by app: %v", err)
	}

	var banAppUser *npool.BanAppUser
	for _, info := range infos {
		banAppUser = dbRowToBanAppUser(info)
		break
	}

	return &npool.GetBanAppUserByAppUserResponse{
		Info: banAppUser,
	}, nil
}

func Update(ctx context.Context, in *npool.UpdateBanAppUserRequest) (*npool.UpdateBanAppUserResponse, error) {
	id, err := uuid.Parse(in.GetInfo().GetID())
	if err != nil {
		return nil, fmt.Errorf("invalid id: %v", err)
	}

	if err := validateBanAppUser(in.GetInfo()); err != nil {
		return nil, fmt.Errorf("invalid parameter: %v", err)
	}

	ctx, cancel := context.WithTimeout(ctx, constant.DBTimeout)
	defer cancel()

	cli, err := db.Client()
	if err != nil {
		return nil, fmt.Errorf("fail get db client: %v", err)
	}

	info, err := cli.
		BanAppUser.
		UpdateOneID(id).
		SetMessage(in.GetInfo().GetMessage()).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("fail update ban app user: %v", err)
	}

	return &npool.UpdateBanAppUserResponse{
		Info: dbRowToBanAppUser(info),
	}, nil
}

func Delete(ctx context.Context, in *npool.DeleteBanAppUserRequest) (*npool.DeleteBanAppUserResponse, error) {
	id, err := uuid.Parse(in.GetID())
	if err != nil {
		return nil, fmt.Errorf("invalid id: %v", err)
	}

	ctx, cancel := context.WithTimeout(ctx, constant.DBTimeout)
	defer cancel()

	cli, err := db.Client()
	if err != nil {
		return nil, fmt.Errorf("fail get db client: %v", err)
	}

	info, err := cli.
		BanAppUser.
		UpdateOneID(id).
		SetDeletedAt(uint32(time.Now().Unix())).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("fail delete ban app user: %v", err)
	}

	return &npool.DeleteBanAppUserResponse{
		Info: dbRowToBanAppUser(info),
	}, nil
}
