package approle

import (
	"context"
	"fmt"

	npool "github.com/NpoolPlatform/message/npool/appusermgr"

	constant "github.com/NpoolPlatform/appuser-manager/pkg/const"
	db "github.com/NpoolPlatform/appuser-manager/pkg/db"
	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent"
	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent/approle"

	"github.com/google/uuid"
)

func validateAppRole(info *npool.AppRole) error {
	if _, err := uuid.Parse(info.GetAppID()); err != nil {
		return fmt.Errorf("invalid app id: %v", err)
	}
	if _, err := uuid.Parse(info.GetCreatedBy()); err != nil {
		return fmt.Errorf("invalid creator: %v", err)
	}
	return nil
}

func dbRowToAppRole(row *ent.AppRole) *npool.AppRole {
	return &npool.AppRole{
		ID:          row.ID.String(),
		AppID:       row.AppID.String(),
		CreatedBy:   row.CreatedBy.String(),
		Role:        row.Role,
		Description: row.Description,
		Default:     row.Default,
	}
}

func Create(ctx context.Context, in *npool.CreateAppRoleRequest) (*npool.CreateAppRoleResponse, error) {
	if err := validateAppRole(in.GetInfo()); err != nil {
		return nil, fmt.Errorf("invalid parameter: %v", err)
	}

	ctx, cancel := context.WithTimeout(ctx, constant.DBTimeout)
	defer cancel()

	cli, err := db.Client()
	if err != nil {
		return nil, fmt.Errorf("fail get db client: %v", err)
	}

	info, err := cli.
		AppRole.
		Create().
		SetAppID(uuid.MustParse(in.GetInfo().GetAppID())).
		SetCreatedBy(uuid.MustParse(in.GetInfo().GetCreatedBy())).
		SetRole(in.GetInfo().GetRole()).
		SetDescription(in.GetInfo().GetDescription()).
		SetDefault(in.GetInfo().GetDefault()).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("fail create app role: %v", err)
	}

	return &npool.CreateAppRoleResponse{
		Info: dbRowToAppRole(info),
	}, nil
}

func Update(ctx context.Context, in *npool.UpdateAppRoleRequest) (*npool.UpdateAppRoleResponse, error) {
	id, err := uuid.Parse(in.GetInfo().GetID())
	if err != nil {
		return nil, fmt.Errorf("invalid app role id: %v", err)
	}

	if err := validateAppRole(in.GetInfo()); err != nil {
		return nil, fmt.Errorf("invalid parameter: %v", err)
	}

	ctx, cancel := context.WithTimeout(ctx, constant.DBTimeout)
	defer cancel()

	cli, err := db.Client()
	if err != nil {
		return nil, fmt.Errorf("fail get db client: %v", err)
	}

	info, err := cli.
		AppRole.
		UpdateOneID(id).
		SetRole(in.GetInfo().GetRole()).
		SetDescription(in.GetInfo().GetDescription()).
		SetDefault(in.GetInfo().GetDefault()).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("fail update app role: %v", err)
	}

	return &npool.UpdateAppRoleResponse{
		Info: dbRowToAppRole(info),
	}, nil
}

func Get(ctx context.Context, in *npool.GetAppRoleRequest) (*npool.GetAppRoleResponse, error) {
	id, err := uuid.Parse(in.GetID())
	if err != nil {
		return nil, fmt.Errorf("invalid app role id: %v", err)
	}

	ctx, cancel := context.WithTimeout(ctx, constant.DBTimeout)
	defer cancel()

	cli, err := db.Client()
	if err != nil {
		return nil, fmt.Errorf("fail get db client: %v", err)
	}

	infos, err := cli.
		AppRole.
		Query().
		Where(
			approle.And(
				approle.ID(id),
			),
		).
		All(ctx)
	if err != nil {
		return nil, fmt.Errorf("fail query app role: %v", err)
	}

	var myAppRole *npool.AppRole
	for _, info := range infos {
		myAppRole = dbRowToAppRole(info)
		break
	}

	return &npool.GetAppRoleResponse{
		Info: myAppRole,
	}, nil
}

func GetByAppRole(ctx context.Context, in *npool.GetAppRoleByAppRoleRequest) (*npool.GetAppRoleByAppRoleResponse, error) {
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
		AppRole.
		Query().
		Where(
			approle.And(
				approle.AppID(appID),
				approle.Role(in.GetRole()),
			),
		).
		All(ctx)
	if err != nil {
		return nil, fmt.Errorf("fail query app role: %v", err)
	}

	var appRole *npool.AppRole
	for _, info := range infos {
		appRole = dbRowToAppRole(info)
		break
	}

	return &npool.GetAppRoleByAppRoleResponse{
		Info: appRole,
	}, nil
}

func GetByApp(ctx context.Context, in *npool.GetAppRolesByAppRequest) (*npool.GetAppRolesByAppResponse, error) {
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
		AppRole.
		Query().
		Where(
			approle.AppID(appID),
		).
		All(ctx)
	if err != nil {
		return nil, fmt.Errorf("fail query app role: %v", err)
	}

	appRoles := []*npool.AppRole{}
	for _, info := range infos {
		appRoles = append(appRoles, dbRowToAppRole(info))
	}

	return &npool.GetAppRolesByAppResponse{
		Infos: appRoles,
	}, nil
}

func GetAppDefaultRole(ctx context.Context, appID string) (*npool.AppRole, error) {
	myAppID, err := uuid.Parse(appID)
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
		AppRole.
		Query().
		Where(
			approle.AppID(myAppID),
			approle.Default(true),
		).
		All(ctx)
	if err != nil {
		return nil, fmt.Errorf("fail query app role: %v", err)
	}

	var appRole *npool.AppRole
	for _, info := range infos {
		appRole = dbRowToAppRole(info)
		break
	}

	return appRole, nil
}
