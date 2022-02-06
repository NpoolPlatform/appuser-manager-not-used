package approle

import (
	"context"

	constant "github.com/NpoolPlatform/appuser-manager/pkg/const"
	approlecrud "github.com/NpoolPlatform/appuser-manager/pkg/crud/approle"
	npool "github.com/NpoolPlatform/message/npool/appusermgr"

	"golang.org/x/xerrors"
)

func Create(ctx context.Context, in *npool.CreateAppRoleRequest) (*npool.CreateAppRoleResponse, error) {
	if in.GetInfo().GetRole() == constant.GenesisRole {
		return nil, xerrors.Errorf("permission denied")
	}

	if in.GetUserID() != in.GetInfo().GetCreatedBy() {
		return nil, xerrors.Errorf("permission denied")
	}

	return approlecrud.Create(ctx, in)
}

func CreateForOtherApp(ctx context.Context, in *npool.CreateAppRoleForOtherAppRequest) (*npool.CreateAppRoleForOtherAppResponse, error) {
	info := in.GetInfo()
	info.AppID = in.GetTargetAppID()

	resp, err := Create(ctx, &npool.CreateAppRoleRequest{
		UserID: info.GetCreatedBy(),
		Info:   info,
	})
	if err != nil {
		return nil, xerrors.Errorf("fail create app role: %v", err)
	}

	return &npool.CreateAppRoleForOtherAppResponse{
		Info: resp.Info,
	}, nil
}
