//nolint:nolintlint,dupl
package admin

import (
	"context"
	"fmt"

	constant "github.com/NpoolPlatform/appuser-manager/pkg/const"
	appcrud "github.com/NpoolPlatform/appuser-manager/pkg/crud/v1/app"
	approlecrud "github.com/NpoolPlatform/appuser-manager/pkg/crud/v1/approle"
	approleusercrud "github.com/NpoolPlatform/appuser-manager/pkg/crud/v1/approleuser"
	appusercrud "github.com/NpoolPlatform/appuser-manager/pkg/crud/v1/appuser"
	appusermw "github.com/NpoolPlatform/appuser-manager/pkg/middleware/appuser"
	npool "github.com/NpoolPlatform/message/npool/appuser/mgr/v1"

	"github.com/google/uuid"
)

func CreateAdminApps(ctx context.Context, in *npool.CreateAdminAppsRequest) (*npool.CreateAdminAppsResponse, error) {
	apps := []*npool.App{}

	genesis, err := appcrud.Get(ctx, &npool.GetAppRequest{
		ID: constant.GenesisAppID,
	})
	if err != nil {
		return nil, fmt.Errorf("fail get genesis app: %v", err)
	}

	if genesis.Info == nil {
		adminApp := npool.App{
			ID:          constant.GenesisAppID,
			CreatedBy:   uuid.UUID{}.String(),
			Name:        constant.GenesisAppName,
			Logo:        "NOT SET",
			Description: "NOT SET",
		}

		resp, err := appcrud.Create(ctx, &npool.CreateAppRequest{
			Info: &adminApp,
		}, true)
		if err != nil {
			return nil, fmt.Errorf("fail create genesis app: %v", err)
		}

		apps = append(apps, resp.Info)
	} else {
		apps = append(apps, genesis.Info)
	}

	church, err := appcrud.Get(ctx, &npool.GetAppRequest{
		ID: constant.ChurchAppID,
	})
	if err != nil {
		return nil, fmt.Errorf("fail get genesis app: %v", err)
	}

	if church.Info == nil {
		adminApp := npool.App{
			ID:          constant.ChurchAppID,
			CreatedBy:   uuid.UUID{}.String(),
			Name:        constant.ChurchAppName,
			Logo:        "NOT SET",
			Description: "NOT SET",
		}

		resp, err := appcrud.Create(ctx, &npool.CreateAppRequest{
			Info: &adminApp,
		}, true)
		if err != nil {
			return nil, fmt.Errorf("fail create church app: %v", err)
		}

		apps = append(apps, resp.Info)
	} else {
		apps = append(apps, church.Info)
	}

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
		return nil, fmt.Errorf("fail get genesis app: %v", err)
	}

	if resp.Info != nil {
		apps = append(apps, resp.Info)
	}

	resp, err = appcrud.Get(ctx, &npool.GetAppRequest{
		ID: constant.ChurchAppID,
	})
	if err != nil {
		return nil, fmt.Errorf("fail get church app: %v", err)
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
		return nil, fmt.Errorf("fail create genesis role: %v", err)
	}

	return &npool.CreateGenesisRoleResponse{
		Info: resp.Info,
	}, nil
}

func GetGenesisRole(ctx context.Context, in *npool.GetGenesisRoleRequest) (*npool.GetGenesisRoleResponse, error) {
	resp, err := approlecrud.GetByAppRole(ctx, &npool.GetAppRoleByAppRoleRequest{
		AppID: constant.GenesisAppID,
		Role:  constant.GenesisRole,
	})
	if err != nil {
		return nil, fmt.Errorf("fail get app role by app role: %v", err)
	}

	return &npool.GetGenesisRoleResponse{
		Info: resp.Info,
	}, nil
}

func CreateGenesisRoleUser(ctx context.Context, in *npool.CreateGenesisRoleUserRequest) (*npool.CreateGenesisRoleUserResponse, error) {
	if in.GetUser().GetAppID() != constant.GenesisAppID && in.GetUser().GetAppID() != constant.ChurchAppID {
		return nil, fmt.Errorf("invalid app id for genesis role user")
	}

	resp, err := approlecrud.GetByAppRole(ctx, &npool.GetAppRoleByAppRoleRequest{
		AppID: constant.GenesisAppID,
		Role:  constant.GenesisRole,
	})
	if err != nil {
		return nil, fmt.Errorf("fail get app role by app role: %v", err)
	}

	resp1, err := approleusercrud.GetUsersByAppRole(ctx, &npool.GetAppRoleUsersByAppRoleRequest{
		AppID:  in.GetUser().GetAppID(),
		RoleID: resp.Info.ID,
	})
	if err != nil {
		return nil, fmt.Errorf("fail get genesis role user: %v", err)
	}
	if len(resp1.Infos) > 0 {
		return nil, fmt.Errorf("genesis user already exist")
	}

	resp2, err := appusercrud.GetByAppAccount(ctx, &npool.GetAppUserByAppAccountRequest{
		AppID:   in.GetUser().GetAppID(),
		Account: in.GetUser().GetEmailAddress(),
	})
	if err != nil {
		return nil, fmt.Errorf("fail get app user: %v", err)
	}

	myUser := resp2.Info

	if myUser == nil {
		resp, err := appusermw.CreateWithSecret(ctx, &npool.CreateAppUserWithSecretRequest{
			User:   in.GetUser(),
			Secret: in.GetSecret(),
		}, false)
		if err != nil {
			return nil, fmt.Errorf("fail create user with secret: %v", err)
		}
		myUser = resp.Info
	}

	resp3, err := approleusercrud.Create(ctx, &npool.CreateAppRoleUserRequest{
		Info: &npool.AppRoleUser{
			AppID:  in.GetUser().GetAppID(),
			RoleID: resp.Info.ID,
			UserID: myUser.ID,
		},
	})
	if err != nil {
		return nil, fmt.Errorf("fail create genesis role user: %v", err)
	}

	return &npool.CreateGenesisRoleUserResponse{
		User:     myUser,
		RoleUser: resp3.Info,
	}, nil
}
