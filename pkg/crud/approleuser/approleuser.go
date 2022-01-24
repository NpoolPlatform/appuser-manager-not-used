package approleuser

import (
	"context"
	"time"

	npool "github.com/NpoolPlatform/message/npool/appusermgr"

	constant "github.com/NpoolPlatform/appuser-manager/pkg/const"
	db "github.com/NpoolPlatform/appuser-manager/pkg/db"
	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent"
	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent/approleuser"

	"github.com/google/uuid"

	"golang.org/x/xerrors"
)

func validateAppRoleUser(info *npool.AppRoleUser) error {
	if _, err := uuid.Parse(info.GetAppID()); err != nil {
		return xerrors.Errorf("invalid app id: %v", err)
	}
	if _, err := uuid.Parse(info.GetRoleID()); err != nil {
		return xerrors.Errorf("invalid role id: %v", err)
	}
	if _, err := uuid.Parse(info.GetUserID()); err != nil {
		return xerrors.Errorf("invalid user id: %v", err)
	}
	return nil
}

func dbRowToAppRoleUser(row *ent.AppRoleUser) *npool.AppRoleUser {
	return &npool.AppRoleUser{
		ID:     row.ID.String(),
		AppID:  row.AppID.String(),
		RoleID: row.RoleID.String(),
		UserID: row.UserID.String(),
	}
}

func Create(ctx context.Context, in *npool.CreateAppRoleUserRequest) (*npool.CreateAppRoleUserResponse, error) {
	if err := validateAppRoleUser(in.GetInfo()); err != nil {
		return nil, xerrors.Errorf("invalid parameter: %v", err)
	}

	ctx, cancel := context.WithTimeout(ctx, constant.DBTimeout)
	defer cancel()

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	info, err := cli.
		AppRoleUser.
		Create().
		SetAppID(uuid.MustParse(in.GetInfo().GetAppID())).
		SetRoleID(uuid.MustParse(in.GetInfo().GetRoleID())).
		SetUserID(uuid.MustParse(in.GetInfo().GetUserID())).
		Save(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail create app role user: %v", err)
	}

	return &npool.CreateAppRoleUserResponse{
		Info: dbRowToAppRoleUser(info),
	}, nil
}

func Get(ctx context.Context, in *npool.GetAppRoleUserRequest) (*npool.GetAppRoleUserResponse, error) {
	id, err := uuid.Parse(in.GetID())
	if err != nil {
		return nil, xerrors.Errorf("invalid app role user id: %v", err)
	}

	ctx, cancel := context.WithTimeout(ctx, constant.DBTimeout)
	defer cancel()

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	infos, err := cli.
		AppRoleUser.
		Query().
		Where(
			approleuser.And(
				approleuser.ID(id),
				approleuser.DeleteAt(0),
			),
		).
		All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail query app role user: %v", err)
	}

	var myAppRoleUser *npool.AppRoleUser
	for _, info := range infos {
		myAppRoleUser = dbRowToAppRoleUser(info)
		break
	}

	return &npool.GetAppRoleUserResponse{
		Info: myAppRoleUser,
	}, nil
}

func GetUsersByAppRole(ctx context.Context, in *npool.GetAppRoleUsersByAppRoleRequest) (*npool.GetAppRoleUsersByAppRoleResponse, error) {
	appID, err := uuid.Parse(in.GetAppID())
	if err != nil {
		return nil, xerrors.Errorf("invalid app id: %v", err)
	}

	roleID, err := uuid.Parse(in.GetRoleID())
	if err != nil {
		return nil, xerrors.Errorf("invalid role id: %v", err)
	}

	ctx, cancel := context.WithTimeout(ctx, constant.DBTimeout)
	defer cancel()

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	infos, err := cli.
		AppRoleUser.
		Query().
		Where(
			approleuser.And(
				approleuser.AppID(appID),
				approleuser.RoleID(roleID),
				approleuser.DeleteAt(0),
			),
		).
		All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail query app role user: %v", err)
	}

	appRoleUsers := []*npool.AppRoleUser{}
	for _, info := range infos {
		appRoleUsers = append(appRoleUsers, dbRowToAppRoleUser(info))
	}

	return &npool.GetAppRoleUsersByAppRoleResponse{
		Infos: appRoleUsers,
	}, nil
}

func Delete(ctx context.Context, in *npool.DeleteAppRoleUserRequest) (*npool.DeleteAppRoleUserResponse, error) {
	id, err := uuid.Parse(in.GetID())
	if err != nil {
		return nil, xerrors.Errorf("invalid app role user id: %v", err)
	}

	ctx, cancel := context.WithTimeout(ctx, constant.DBTimeout)
	defer cancel()

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	info, err := cli.
		AppRoleUser.
		UpdateOneID(id).
		SetDeleteAt(uint32(time.Now().Unix())).
		Save(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail update app role user: %v", err)
	}

	return &npool.DeleteAppRoleUserResponse{
		Info: dbRowToAppRoleUser(info),
	}, nil
}
