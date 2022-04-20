package approleuser

import (
	"context"
	"fmt"

	constant "github.com/NpoolPlatform/appuser-manager/pkg/const"
	approlecrud "github.com/NpoolPlatform/appuser-manager/pkg/crud/approle"
	approleusercrud "github.com/NpoolPlatform/appuser-manager/pkg/crud/approleuser"
	npool "github.com/NpoolPlatform/message/npool/appusermgr"
)

func CreateAppRoleUser(ctx context.Context, in *npool.CreateAppRoleUserRequest) (*npool.CreateAppRoleUserResponse, error) {
	role, err := approlecrud.Get(ctx, &npool.GetAppRoleRequest{
		ID: in.GetInfo().GetRoleID(),
	})
	if err != nil {
		return nil, fmt.Errorf("fail get role: %v", err)
	}

	if role.Info.Role == constant.GenesisRole {
		return nil, fmt.Errorf("permission denied")
	}

	resp, err := approleusercrud.GetByAppUser(ctx, &npool.GetAppRoleUserByAppUserRequest{
		AppID:  in.GetInfo().GetAppID(),
		UserID: in.GetInfo().GetUserID(),
	})
	if err != nil {
		return nil, fmt.Errorf("fail get app role user: %v", err)
	}

	for _, info := range resp.Infos {
		if in.GetInfo().GetRoleID() == info.RoleID {
			return nil, fmt.Errorf("app role user exist")
		}
	}

	return approleusercrud.Create(ctx, in)
}
