package appuser

import (
	"context"

	approlecrud "github.com/NpoolPlatform/appuser-manager/pkg/crud/approle"
	approleusercrud "github.com/NpoolPlatform/appuser-manager/pkg/crud/approleuser"
	appusercrud "github.com/NpoolPlatform/appuser-manager/pkg/crud/appuser"
	appusercontrolcrud "github.com/NpoolPlatform/appuser-manager/pkg/crud/appusercontrol"
	appuserextracrud "github.com/NpoolPlatform/appuser-manager/pkg/crud/appuserextra"
	appusersecretcrud "github.com/NpoolPlatform/appuser-manager/pkg/crud/appusersecret"
	banappusercrud "github.com/NpoolPlatform/appuser-manager/pkg/crud/banappuser"
	encrypt "github.com/NpoolPlatform/appuser-manager/pkg/middleware/encrypt"
	npool "github.com/NpoolPlatform/message/npool/appusermgr"

	"golang.org/x/xerrors"
)

func CreateWithSecret(ctx context.Context, in *npool.CreateAppUserWithSecretRequest) (*npool.CreateAppUserWithSecretResponse, error) {
	resp, err := appusercrud.Create(ctx, &npool.CreateAppUserRequest{
		Info: in.GetUser(),
	})
	if err != nil {
		return nil, xerrors.Errorf("fail create app user: %v", err)
	}

	inSecret := in.GetSecret()
	inSecret.UserID = resp.Info.ID

	_, err = appusersecretcrud.Create(ctx, &npool.CreateAppUserSecretRequest{
		Info: inSecret,
	})
	if err != nil {
		// TODO: rollback for secret create error
		return nil, xerrors.Errorf("fail create app user secret: %v", err)
	}

	return &npool.CreateAppUserWithSecretResponse{
		Info: resp.Info,
	}, nil
}

func GetRolesByAppUser(ctx context.Context, in *npool.GetUserRolesByAppUserRequest) (*npool.GetUserRolesByAppUserResponse, error) {
	resp, err := approleusercrud.GetByAppUser(ctx, &npool.GetAppRoleUserByAppUserRequest{
		AppID:  in.GetAppID(),
		UserID: in.GetUserID(),
	})
	if err != nil {
		return nil, xerrors.Errorf("fail get app role user by app user: %v", err)
	}

	roles := []*npool.AppRole{}
	for _, info := range resp.Infos {
		resp1, err := approlecrud.Get(ctx, &npool.GetAppRoleRequest{
			ID: info.RoleID,
		})
		if err != nil {
			return nil, xerrors.Errorf("fail get app role: %v", err)
		}
		roles = append(roles, resp1.Info)
	}

	return &npool.GetUserRolesByAppUserResponse{
		Infos: roles,
	}, nil
}

func expandAppUserInfo(ctx context.Context, user *npool.AppUser) (*npool.AppUserInfo, error) { //nolint
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

	return info, nil
}

func VerifyByAppAccountPassword(ctx context.Context, in *npool.VerifyAppUserByAppAccountPasswordRequest) (*npool.VerifyAppUserByAppAccountPasswordResponse, error) { //nolint
	resp, err := appusercrud.GetByAppAccount(ctx, &npool.GetAppUserByAppAccountRequest{
		AppID:   in.GetAppID(),
		Account: in.GetAccount(),
	})
	if err != nil {
		return nil, xerrors.Errorf("fail get app user by app account: %v", err)
	}
	if resp.Info == nil {
		return nil, xerrors.Errorf("fail get app user by app account")
	}

	resp1, err := appusersecretcrud.GetByAppUser(ctx, &npool.GetAppUserSecretByAppUserRequest{
		AppID:  in.GetAppID(),
		UserID: resp.Info.ID,
	})
	if err != nil {
		return nil, xerrors.Errorf("fail get app user secret by app user: %v", err)
	}
	if resp1.Info == nil {
		return nil, xerrors.Errorf("fail get app user secret by app user")
	}

	err = encrypt.VerifyWithSalt(in.GetPasswordHash(), resp1.Info.PasswordHash, resp1.Info.Salt)
	if err != nil {
		return nil, xerrors.Errorf("invalid account or password: %v", err)
	}

	info, err := expandAppUserInfo(ctx, resp.Info)
	if err != nil {
		return nil, xerrors.Errorf("fail expand app user: %v", err)
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
		return nil, xerrors.Errorf("fail get app user: %v", err)
	}

	info, err := expandAppUserInfo(ctx, resp.Info)
	if err != nil {
		return nil, xerrors.Errorf("fail expand app user: %v", err)
	}

	return &npool.GetAppUserInfoResponse{
		Info: info,
	}, nil
}

func GetAppUserInfosByApp(ctx context.Context, in *npool.GetAppUserInfosByAppRequest) (*npool.GetAppUserInfosByAppResponse, error) {
	resp, err := appusercrud.GetByApp(ctx, &npool.GetAppUsersByAppRequest{
		AppID: in.GetAppID(),
	})
	if err != nil {
		return nil, xerrors.Errorf("fail get app user by app: %v", err)
	}

	infos := []*npool.AppUserInfo{}
	for _, info := range resp.Infos {
		userInfo, err := expandAppUserInfo(ctx, info)
		if err != nil {
			return nil, xerrors.Errorf("fail expand app user: %v", err)
		}
		infos = append(infos, userInfo)
	}

	return &npool.GetAppUserInfosByAppResponse{
		Infos: infos,
	}, nil
}
