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

	if in.GetInfo().GetUserID() != in.GetInfo().GetCreatedBy() {
		return nil, xerrors.Errorf("permission denied")
	}

	return approlecrud.Create(ctx, in)
}
