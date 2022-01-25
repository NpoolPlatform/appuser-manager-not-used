package admin

import (
	"context"

	constant "github.com/NpoolPlatform/appuser-manager/pkg/const"
	appcrud "github.com/NpoolPlatform/appuser-manager/pkg/crud/app"
	npool "github.com/NpoolPlatform/message/npool/appusermgr"

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
	return nil, nil
}

func CreateGenesisRole(ctx context.Context, in *npool.CreateGenesisRoleRequest) (*npool.CreateGenesisRoleResponse, error) {
	return nil, nil
}

func GetGenesisRole(ctx context.Context, in *npool.GetGenesisRoleRequest) (*npool.GetGenesisRoleResponse, error) {
	return nil, nil
}

func CreateGenesisRoleUser(ctx context.Context, in *npool.CreateGenesisRoleUserRequest) (*npool.CreateGenesisRoleUserResponse, error) {
	return nil, nil
}
