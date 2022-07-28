package approle

import (
	"context"
	"fmt"

	constant "github.com/NpoolPlatform/appuser-manager/pkg/const"
	approlecrud "github.com/NpoolPlatform/appuser-manager/pkg/crud/v1/approle"
	npool "github.com/NpoolPlatform/message/npool/appuser/mgr/v1"
)

func Create(ctx context.Context, in *npool.CreateAppRoleRequest) (*npool.CreateAppRoleResponse, error) {
	if in.GetInfo().GetRole() == constant.GenesisRole {
		return nil, fmt.Errorf("permission denied")
	}

	if in.GetUserID() != in.GetInfo().GetCreatedBy() {
		return nil, fmt.Errorf("permission denied")
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
		return nil, fmt.Errorf("fail create app role: %v", err)
	}

	return &npool.CreateAppRoleForOtherAppResponse{
		Info: resp.Info,
	}, nil
}
