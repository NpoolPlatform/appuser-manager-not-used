package admin

import (
	"context"

	constant "github.com/NpoolPlatform/appuser-manager/pkg/const"
	appcrud "github.com/NpoolPlatform/appuser-manager/pkg/crud/app"
	approlecrud "github.com/NpoolPlatform/appuser-manager/pkg/crud/approle"
	npool "github.com/NpoolPlatform/message/npool/appusermgr"

	"github.com/google/uuid"

	"golang.org/x/xerrors"
)

func CreateAdminApps(ctx context.Context, in *npool.CreateAdminAppsRequest) (*npool.CreateAdminAppsResponse, error) {
	apps := []*npool.App{}

	adminApp := npool.App{
		ID:          constant.GenesisAppID,
		Name:        constant.GenesisAppName,
		Logo:        "NOT SET",
		Description: "NOT SET",
	}

	resp, err := appcrud.Create(ctx, &npool.CreateAppRequest{
		Info: &adminApp,
	}, true)
	if err != nil {
		return nil, xerrors.Errorf("fail create genesis app: %v", err)
	}

	apps = append(apps, resp.Info)

	adminApp = npool.App{
		ID:          constant.ChurchAppID,
		Name:        constant.ChurchAppName,
		Logo:        "NOT SET",
		Description: "NOT SET",
	}

	resp, err = appcrud.Create(ctx, &npool.CreateAppRequest{
		Info: &adminApp,
	}, true)
	if err != nil {
		return nil, xerrors.Errorf("fail create church app: %v", err)
	}

	apps = append(apps, resp.Info)

	return &npool.CreateAdminAppsResponse{
		Infos: apps,
	}, nil
}

func GetAdminApps(ctx context.Context, in *npool.GetAdminAppsRequest) (*npool.GetAdminAppsResponse, error) {
	apps := []*npool.App{}

	resp, err := appcrud.Get(ctx, &npool.GetAppRequest{
		ID: constant.GenesisAppID,
	})
	if err != nil {
		return nil, xerrors.Errorf("fail get genesis app: %v", err)
	}

	if resp.Info != nil {
		apps = append(apps, resp.Info)
	}

	resp, err = appcrud.Get(ctx, &npool.GetAppRequest{
		ID: constant.ChurchAppID,
	})
	if err != nil {
		return nil, xerrors.Errorf("fail get church app: %v", err)
	}

	if resp.Info != nil {
		apps = append(apps, resp.Info)
	}

	return &npool.GetAdminAppsResponse{
		Infos: apps,
	}, nil
}

func CreateGenesisRole(ctx context.Context, in *npool.CreateGenesisRoleRequest) (*npool.CreateGenesisRoleResponse, error) {
	genesisRole := npool.AppRole{
		AppID:       uuid.UUID{}.String(),
		CreatedBy:   uuid.UUID{}.String(),
		Role:        constant.GenesisRole,
		Description: "NOT SET",
		Default:     false,
	}

	resp, err := approlecrud.Create(ctx, &npool.CreateAppRoleRequest{
		Info: &genesisRole,
	})
	if err != nil {
		return nil, xerrors.Errorf("fail create genesis role: %v", err)
	}

	return &npool.CreateGenesisRoleResponse{
		Info: resp.Info,
	}, nil
}

func GetGenesisRole(ctx context.Context, in *npool.GetGenesisRoleRequest) (*npool.GetGenesisRoleResponse, error) {
	return nil, nil
}

func CreateGenesisRoleUser(ctx context.Context, in *npool.CreateGenesisRoleUserRequest) (*npool.CreateGenesisRoleUserResponse, error) {
	return nil, nil
}
