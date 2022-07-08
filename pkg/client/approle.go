//nolint:dupl
package client

import (
	"context"
	"fmt"
	"time"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/appusermgrv2/approle"

	constant "github.com/NpoolPlatform/appuser-manager/pkg/message/const"
)

func doAppRole(ctx context.Context, fn func(_ctx context.Context, cli npool.AppUserManagerAppRoleClient) (cruder.Any, error)) (cruder.Any, error) {
	_ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	conn, err := grpc2.GetGRPCConn(constant.ServiceName, grpc2.GRPCTAG)
	if err != nil {
		return nil, fmt.Errorf("fail get app role connection: %v", err)
	}

	defer conn.Close()

	cli := npool.NewAppUserManagerAppRoleClient(conn)

	return fn(_ctx, cli)
}

func CreateAppRoleV2(ctx context.Context, in *npool.AppRoleReq) (*npool.AppRole, error) {
	info, err := doAppRole(ctx, func(_ctx context.Context, cli npool.AppUserManagerAppRoleClient) (cruder.Any, error) {
		resp, err := cli.CreateAppRoleV2(ctx, &npool.CreateAppRoleRequest{
			Info: in,
		})
		if err != nil {
			return nil, fmt.Errorf("fail create app role: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail create app role: %v", err)
	}
	return info.(*npool.AppRole), nil
}

func CreateAppRolesV2(ctx context.Context, in []*npool.AppRoleReq) ([]*npool.AppRole, error) {
	infos, err := doAppRole(ctx, func(_ctx context.Context, cli npool.AppUserManagerAppRoleClient) (cruder.Any, error) {
		resp, err := cli.CreateAppRolesV2(ctx, &npool.CreateAppRolesRequest{
			Infos: in,
		})
		if err != nil {
			return nil, fmt.Errorf("fail create app roles: %v", err)
		}
		return resp.GetInfos(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail create app roles: %v", err)
	}
	return infos.([]*npool.AppRole), nil
}

func UpdateAppRoleV2(ctx context.Context, in *npool.AppRoleReq) (*npool.AppRole, error) {
	info, err := doAppRole(ctx, func(_ctx context.Context, cli npool.AppUserManagerAppRoleClient) (cruder.Any, error) {
		resp, err := cli.UpdateAppRoleV2(ctx, &npool.UpdateAppRoleRequest{
			Info: in,
		})
		if err != nil {
			return nil, fmt.Errorf("fail update app role: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail update app role: %v", err)
	}
	return info.(*npool.AppRole), nil
}

func GetAppRoleV2(ctx context.Context, id string) (*npool.AppRole, error) {
	info, err := doAppRole(ctx, func(_ctx context.Context, cli npool.AppUserManagerAppRoleClient) (cruder.Any, error) {
		resp, err := cli.GetAppRoleV2(ctx, &npool.GetAppRoleRequest{
			ID: id,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get app role: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail get app role: %v", err)
	}
	return info.(*npool.AppRole), nil
}

func GetAppRoleOnlyV2(ctx context.Context, conds *npool.Conds) (*npool.AppRole, error) {
	info, err := doAppRole(ctx, func(_ctx context.Context, cli npool.AppUserManagerAppRoleClient) (cruder.Any, error) {
		resp, err := cli.GetAppRoleOnlyV2(ctx, &npool.GetAppRoleOnlyRequest{
			Conds: conds,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get app role: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail get app role: %v", err)
	}
	return info.(*npool.AppRole), nil
}

func GetAppRolesV2(ctx context.Context, conds *npool.Conds) ([]*npool.AppRole, error) {
	infos, err := doAppRole(ctx, func(_ctx context.Context, cli npool.AppUserManagerAppRoleClient) (cruder.Any, error) {
		resp, err := cli.GetAppRolesV2(ctx, &npool.GetAppRolesRequest{
			Conds: conds,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get app role: %v", err)
		}
		return resp.GetInfos(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail get app role: %v", err)
	}
	return infos.([]*npool.AppRole), nil
}

func ExistAppRoleV2(ctx context.Context, id string) (bool, error) {
	infos, err := doAppRole(ctx, func(_ctx context.Context, cli npool.AppUserManagerAppRoleClient) (cruder.Any, error) {
		resp, err := cli.ExistAppRoleV2(ctx, &npool.ExistAppRoleRequest{
			ID: id,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get app role: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return false, fmt.Errorf("fail get app role: %v", err)
	}
	return infos.(bool), nil
}

func ExistAppRoleCondsV2(ctx context.Context, conds *npool.Conds) (bool, error) {
	infos, err := doAppRole(ctx, func(_ctx context.Context, cli npool.AppUserManagerAppRoleClient) (cruder.Any, error) {
		resp, err := cli.ExistAppRoleCondsV2(ctx, &npool.ExistAppRoleCondsRequest{
			Conds: conds,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get app role: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return false, fmt.Errorf("fail get app role: %v", err)
	}
	return infos.(bool), nil
}

func CountAppRolesV2(ctx context.Context, conds *npool.Conds) (uint32, error) {
	infos, err := doAppRole(ctx, func(_ctx context.Context, cli npool.AppUserManagerAppRoleClient) (cruder.Any, error) {
		resp, err := cli.CountAppRolesV2(ctx, &npool.CountAppRolesRequest{
			Conds: conds,
		})
		if err != nil {
			return nil, fmt.Errorf("fail count app role: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return 0, fmt.Errorf("fail count app role: %v", err)
	}
	return infos.(uint32), nil
}

func DeleteAppRoleV2(ctx context.Context, id string) (*npool.AppRole, error) {
	infos, err := doAppRole(ctx, func(_ctx context.Context, cli npool.AppUserManagerAppRoleClient) (cruder.Any, error) {
		resp, err := cli.DeleteAppRoleV2(ctx, &npool.DeleteAppRoleRequest{
			ID: id,
		})
		if err != nil {
			return nil, fmt.Errorf("fail delete app role: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail delete app role: %v", err)
	}
	return infos.(*npool.AppRole), nil
}
