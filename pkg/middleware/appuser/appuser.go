//nolint:dupl,lll,unparam
package appuser

import (
	"context"
	"fmt"

	appcrud "github.com/NpoolPlatform/appuser-manager/pkg/crud/v1/app"
	appcontrolcrud "github.com/NpoolPlatform/appuser-manager/pkg/crud/v1/appcontrol"
	approlecrud "github.com/NpoolPlatform/appuser-manager/pkg/crud/v1/approle"
	approleusercrud "github.com/NpoolPlatform/appuser-manager/pkg/crud/v1/approleuser"
	appusercrud "github.com/NpoolPlatform/appuser-manager/pkg/crud/v1/appuser"
	appusercontrolcrud "github.com/NpoolPlatform/appuser-manager/pkg/crud/v1/appusercontrol"
	appuserextracrud "github.com/NpoolPlatform/appuser-manager/pkg/crud/v1/appuserextra"
	appusersecretcrud "github.com/NpoolPlatform/appuser-manager/pkg/crud/v1/appusersecret"
	appuserthirdcrud "github.com/NpoolPlatform/appuser-manager/pkg/crud/v1/appuserthirdparty"
	banappcrud "github.com/NpoolPlatform/appuser-manager/pkg/crud/v1/banapp"
	banappusercrud "github.com/NpoolPlatform/appuser-manager/pkg/crud/v1/banappuser"
	encrypt "github.com/NpoolPlatform/appuser-manager/pkg/middleware/encrypt"
	npool "github.com/NpoolPlatform/message/npool/appuser/mgr/v1"
)

func CreateWithSecret(ctx context.Context, in *npool.CreateAppUserWithSecretRequest, setDefaultRole bool) (*npool.CreateAppUserWithSecretResponse, error) {
	resp, err := appusercrud.Create(ctx, &npool.CreateAppUserRequest{
		Info: in.GetUser(),
	})
	if err != nil {
		return nil, fmt.Errorf("fail create app user: %v", err)
	}

	inSecret := in.GetSecret()
	inSecret.UserID = resp.Info.GetID()

	_, err = appusersecretcrud.Create(ctx, &npool.CreateAppUserSecretRequest{
		Info: inSecret,
	})
	if err != nil {
		// TODO: rollback for secret create error
		return nil, fmt.Errorf("fail create app user secret: %v", err)
	}

	if setDefaultRole {
		defaultRole, err := approlecrud.GetAppDefaultRole(ctx, in.GetUser().GetAppID())
		if err != nil {
			return nil, fmt.Errorf("fail get default role: %v", err)
		}
		if defaultRole == nil {
			return nil, fmt.Errorf("fail get default role")
		}

		_, err = approleusercrud.Create(ctx, &npool.CreateAppRoleUserRequest{
			Info: &npool.AppRoleUser{
				AppID:  in.GetUser().GetAppID(),
				RoleID: defaultRole.ID,
				UserID: resp.Info.ID,
			},
		})
		if err != nil {
			// TODO: rollback for role user create error
			return nil, fmt.Errorf("fail create app role user: %v", err)
		}
	}

	return &npool.CreateAppUserWithSecretResponse{
		Info: resp.Info,
	}, nil
}

func CreateWithSecretRevert(ctx context.Context, in *npool.CreateAppUserWithSecretRequest, setDefaultRole bool) (*npool.CreateAppUserWithSecretResponse, error) {
	resp, err := appusercrud.CreateRevert(ctx, &npool.CreateAppUserRequest{
		Info: in.GetUser(),
	})
	if err != nil {
		return nil, fmt.Errorf("fail create app user: %v", err)
	}

	inSecret := in.GetSecret()
	inSecret.UserID = in.GetUser().GetID()

	_, err = appusersecretcrud.CreateRevert(ctx, &npool.CreateAppUserSecretRequest{
		Info: inSecret,
	})
	if err != nil {
		// TODO: rollback for secret create error
		return nil, fmt.Errorf("fail create app user secret: %v", err)
	}

	if setDefaultRole {
		defaultRole, err := approlecrud.GetAppDefaultRole(ctx, in.GetUser().GetAppID())
		if err != nil {
			return nil, fmt.Errorf("fail get default role: %v", err)
		}
		if defaultRole == nil {
			return nil, fmt.Errorf("fail get default role")
		}

		_, err = approleusercrud.CreateRevert(ctx, &npool.CreateAppRoleUserRequest{
			Info: &npool.AppRoleUser{
				AppID:  in.GetUser().GetAppID(),
				RoleID: defaultRole.ID,
				UserID: in.GetUser().GetID(),
			},
		})
		if err != nil {
			// TODO: rollback for role user create error
			return nil, fmt.Errorf("fail create app role user: %v", err)
		}
	}
	return &npool.CreateAppUserWithSecretResponse{
		Info: resp.Info,
	}, nil
}

func CreateWithThirdParty(ctx context.Context, in *npool.CreateAppUserWithThirdPartyRequest, setDefaultRole bool) (*npool.CreateAppUserWithThirdPartyResponse, error) {
	resp, err := appusercrud.Create(ctx, &npool.CreateAppUserRequest{
		Info: in.GetUser(),
	})
	if err != nil {
		return nil, fmt.Errorf("fail create app user: %v", err)
	}

	inThird := in.GetThirdParty()
	inThird.UserID = resp.Info.ID

	_, err = appuserthirdcrud.Create(ctx, &npool.CreateAppUserThirdPartyRequest{
		Info: inThird,
	})
	if err != nil {
		return nil, fmt.Errorf("fail create app user third: %v", err)
	}

	if setDefaultRole {
		defaultRole, err := approlecrud.GetAppDefaultRole(ctx, in.GetUser().GetAppID())
		if err != nil {
			return nil, fmt.Errorf("fail get default role: %v", err)
		}
		if defaultRole == nil {
			return nil, fmt.Errorf("fail get default role")
		}

		_, err = approleusercrud.CreateRevert(ctx, &npool.CreateAppRoleUserRequest{
			Info: &npool.AppRoleUser{
				AppID:  in.GetUser().GetAppID(),
				RoleID: defaultRole.ID,
				UserID: in.GetUser().GetID(),
			},
		})
		if err != nil {
			// TODO: rollback for role user create error
			return nil, fmt.Errorf("fail create app role user: %v", err)
		}
	}
	return &npool.CreateAppUserWithThirdPartyResponse{
		Info: resp.Info,
	}, nil
}

func GetRolesByAppUser(ctx context.Context, in *npool.GetUserRolesByAppUserRequest) (*npool.GetUserRolesByAppUserResponse, error) {
	resp, err := approleusercrud.GetByAppUser(ctx, &npool.GetAppRoleUserByAppUserRequest{
		AppID:  in.GetAppID(),
		UserID: in.GetUserID(),
	})
	if err != nil {
		return nil, fmt.Errorf("fail get app role user by app user: %v", err)
	}

	roles := []*npool.AppRole{}
	for _, info := range resp.Infos {
		resp1, err := approlecrud.Get(ctx, &npool.GetAppRoleRequest{
			ID: info.RoleID,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get app role: %v", err)
		}
		if resp1.Info == nil {
			continue
		}

		roles = append(roles, resp1.Info)
	}

	return &npool.GetUserRolesByAppUserResponse{
		Infos: roles,
	}, nil
}

func expandAppUserInfo(ctx context.Context, user *npool.AppUser) (*npool.AppUserInfo, error) {
	info := &npool.AppUserInfo{
		User: user,
	}

	resp1, err := appuserextracrud.GetByAppUser(ctx, &npool.GetAppUserExtraByAppUserRequest{
		AppID:  user.AppID,
		UserID: user.ID,
	})
	if err == nil && resp1.Info != nil {
		info.Extra = resp1.Info
	}

	resp2, err := appusercontrolcrud.GetByAppUser(ctx, &npool.GetAppUserControlByAppUserRequest{
		AppID:  user.AppID,
		UserID: user.ID,
	})
	if err == nil && resp2.Info != nil {
		info.Ctrl = resp2.Info
	}

	resp3, err := banappusercrud.GetByAppUser(ctx, &npool.GetBanAppUserByAppUserRequest{
		AppID:  user.AppID,
		UserID: user.ID,
	})
	if err == nil && resp3.Info != nil {
		info.Ban = resp3.Info
	}

	resp4, err := GetRolesByAppUser(ctx, &npool.GetUserRolesByAppUserRequest{
		AppID:  user.AppID,
		UserID: user.ID,
	})
	if err == nil && resp4.Infos != nil {
		info.Roles = resp4.Infos
	}

	resp5, err := appusersecretcrud.GetByAppUser(ctx, &npool.GetAppUserSecretByAppUserRequest{
		AppID:  user.AppID,
		UserID: user.ID,
	})
	if err == nil && resp5.Info != nil {
		secretMap := &npool.AppUserSecretMap{}
		if resp5.Info.GoogleSecret != "" {
			secretMap.HasGoogleSecret = true
		}
		info.SecretMap = secretMap
	}

	return info, nil
}

func VerifyByAppAccountPassword(ctx context.Context, in *npool.VerifyAppUserByAppAccountPasswordRequest) (*npool.VerifyAppUserByAppAccountPasswordResponse, error) {
	resp, err := appusercrud.GetByAppAccount(ctx, &npool.GetAppUserByAppAccountRequest{
		AppID:   in.GetAppID(),
		Account: in.GetAccount(),
	})
	if err != nil {
		return nil, fmt.Errorf("fail get app user by app account: %v", err)
	}
	if resp.Info == nil {
		return nil, fmt.Errorf("fail get app user by app account")
	}

	resp1, err := appusersecretcrud.GetByAppUser(ctx, &npool.GetAppUserSecretByAppUserRequest{
		AppID:  in.GetAppID(),
		UserID: resp.Info.ID,
	})
	if err != nil {
		return nil, fmt.Errorf("fail get app user secret by app user: %v", err)
	}
	if resp1.Info == nil {
		return nil, fmt.Errorf("fail get app user secret by app user")
	}

	err = encrypt.VerifyWithSalt(in.GetPasswordHash(), resp1.Info.PasswordHash, resp1.Info.Salt)
	if err != nil {
		return nil, fmt.Errorf("invalid account or password: %v", err)
	}

	info, err := expandAppUserInfo(ctx, resp.Info)
	if err != nil {
		return nil, fmt.Errorf("fail expand app user: %v", err)
	}

	return &npool.VerifyAppUserByAppAccountPasswordResponse{
		Info: info,
	}, nil
}

func GetAppUserInfo(ctx context.Context, in *npool.GetAppUserInfoRequest) (*npool.GetAppUserInfoResponse, error) {
	resp, err := appusercrud.Get(ctx, &npool.GetAppUserRequest{
		ID: in.GetID(),
	})
	if err != nil {
		return nil, fmt.Errorf("fail get app user: %v", err)
	}
	if resp.Info == nil {
		return nil, fmt.Errorf("fail get app user")
	}

	info, err := expandAppUserInfo(ctx, resp.Info)
	if err != nil {
		return nil, fmt.Errorf("fail expand app user: %v", err)
	}

	return &npool.GetAppUserInfoResponse{
		Info: info,
	}, nil
}

func GetAppUserInfoByAppUser(ctx context.Context, in *npool.GetAppUserInfoByAppUserRequest) (*npool.GetAppUserInfoByAppUserResponse, error) {
	resp, err := appusercrud.GetByAppUser(ctx, &npool.GetAppUserByAppUserRequest{
		AppID:  in.GetAppID(),
		UserID: in.GetUserID(),
	})
	if err != nil {
		return nil, fmt.Errorf("fail get app user: %v", err)
	}
	if resp.Info == nil {
		return nil, fmt.Errorf("fail get app user")
	}

	info, err := expandAppUserInfo(ctx, resp.Info)
	if err != nil {
		return nil, fmt.Errorf("fail expand app user: %v", err)
	}

	return &npool.GetAppUserInfoByAppUserResponse{
		Info: info,
	}, nil
}

func GetAppUserInfosByApp(ctx context.Context, in *npool.GetAppUserInfosByAppRequest) (*npool.GetAppUserInfosByAppResponse, error) {
	resp, err := appusercrud.GetByApp(ctx, &npool.GetAppUsersByAppRequest{
		AppID: in.GetAppID(),
	})
	if err != nil {
		return nil, fmt.Errorf("fail get app user by app: %v", err)
	}

	infos := []*npool.AppUserInfo{}
	for _, info := range resp.Infos {
		userInfo, err := expandAppUserInfo(ctx, info)
		if err != nil {
			return nil, fmt.Errorf("fail expand app user: %v", err)
		}
		infos = append(infos, userInfo)
	}

	return &npool.GetAppUserInfosByAppResponse{
		Infos: infos,
	}, nil
}

func expandAppInfo(ctx context.Context, app *npool.App) (*npool.AppInfo, error) {
	info := npool.AppInfo{
		App: app,
	}

	resp, err := appcontrolcrud.GetByApp(ctx, &npool.GetAppControlByAppRequest{
		AppID: app.ID,
	})
	if err == nil {
		info.Ctrl = resp.Info
	}

	resp1, err := banappcrud.GetByApp(ctx, &npool.GetBanAppByAppRequest{
		AppID: app.ID,
	})
	if err == nil {
		info.Ban = resp1.Info
	}

	return &info, nil
}

func GetAppInfo(ctx context.Context, in *npool.GetAppInfoRequest) (*npool.GetAppInfoResponse, error) {
	resp, err := appcrud.Get(ctx, &npool.GetAppRequest{
		ID: in.GetID(),
	})
	if err != nil {
		return nil, fmt.Errorf("fail get app: %v", err)
	}
	if resp.Info == nil {
		return nil, fmt.Errorf("fail get app")
	}

	info, err := expandAppInfo(ctx, resp.Info)
	if err != nil {
		return nil, fmt.Errorf("fail expand app info: %v", err)
	}

	return &npool.GetAppInfoResponse{
		Info: info,
	}, nil
}

func GetAppInfos(ctx context.Context, in *npool.GetAppInfosRequest) (*npool.GetAppInfosResponse, error) {
	resp, err := appcrud.GetAll(ctx, &npool.GetAppsRequest{})
	if err != nil {
		return nil, fmt.Errorf("fail get apps: %v", err)
	}

	infos := []*npool.AppInfo{}
	for _, info := range resp.Infos {
		appInfo, err := expandAppInfo(ctx, info)
		if err != nil {
			return nil, fmt.Errorf("fail expand app info: %v", err)
		}
		infos = append(infos, appInfo)
	}

	return &npool.GetAppInfosResponse{
		Infos: infos,
	}, nil
}

func GetAppInfosByCreator(ctx context.Context, in *npool.GetAppInfosByCreatorRequest) (*npool.GetAppInfosByCreatorResponse, error) {
	resp, err := appcrud.GetByCreator(ctx, &npool.GetAppsByCreatorRequest{
		UserID: in.GetUserID(),
	})
	if err != nil {
		return nil, fmt.Errorf("fail get apps by creator: %v", err)
	}

	infos := []*npool.AppInfo{}
	for _, info := range resp.Infos {
		appInfo, err := expandAppInfo(ctx, info)
		if err != nil {
			return nil, fmt.Errorf("fail expand app info: %v", err)
		}
		infos = append(infos, appInfo)
	}

	return &npool.GetAppInfosByCreatorResponse{
		Infos: infos,
	}, nil
}
