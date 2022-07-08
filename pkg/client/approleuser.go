//nolint:dupl
package client

import (
	"context"
	"fmt"
	"time"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/appusermgrv2/approleuser"

	constant "github.com/NpoolPlatform/appuser-manager/pkg/message/const"
)

func doAppRoleUser(ctx context.Context, fn func(_ctx context.Context, cli npool.AppUserManagerAppRoleUserClient) (cruder.Any, error)) (cruder.Any, error) {
	_ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	conn, err := grpc2.GetGRPCConn(constant.ServiceName, grpc2.GRPCTAG)
	if err != nil {
		return nil, fmt.Errorf("fail get app role user connection: %v", err)
	}

	defer conn.Close()

	cli := npool.NewAppUserManagerAppRoleUserClient(conn)

	return fn(_ctx, cli)
}

func CreateAppRoleUserV2(ctx context.Context, in *npool.AppRoleUserReq) (*npool.AppRoleUser, error) {
	info, err := doAppRoleUser(ctx, func(_ctx context.Context, cli npool.AppUserManagerAppRoleUserClient) (cruder.Any, error) {
		resp, err := cli.CreateAppRoleUserV2(ctx, &npool.CreateAppRoleUserRequest{
			Info: in,
		})
		if err != nil {
			return nil, fmt.Errorf("fail create app role user: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail create app role user: %v", err)
	}
	return info.(*npool.AppRoleUser), nil
}

func CreateAppRoleUsersV2(ctx context.Context, in []*npool.AppRoleUserReq) ([]*npool.AppRoleUser, error) {
	infos, err := doAppRoleUser(ctx, func(_ctx context.Context, cli npool.AppUserManagerAppRoleUserClient) (cruder.Any, error) {
		resp, err := cli.CreateAppRoleUsersV2(ctx, &npool.CreateAppRoleUsersRequest{
			Infos: in,
		})
		if err != nil {
			return nil, fmt.Errorf("fail create app role users: %v", err)
		}
		return resp.GetInfos(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail create app role users: %v", err)
	}
	return infos.([]*npool.AppRoleUser), nil
}

func UpdateAppRoleUserV2(ctx context.Context, in *npool.AppRoleUserReq) (*npool.AppRoleUser, error) {
	info, err := doAppRoleUser(ctx, func(_ctx context.Context, cli npool.AppUserManagerAppRoleUserClient) (cruder.Any, error) {
		resp, err := cli.UpdateAppRoleUserV2(ctx, &npool.UpdateAppRoleUserRequest{
			Info: in,
		})
		if err != nil {
			return nil, fmt.Errorf("fail update app role user: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail update app role user: %v", err)
	}
	return info.(*npool.AppRoleUser), nil
}

func GetAppRoleUserV2(ctx context.Context, id string) (*npool.AppRoleUser, error) {
	info, err := doAppRoleUser(ctx, func(_ctx context.Context, cli npool.AppUserManagerAppRoleUserClient) (cruder.Any, error) {
		resp, err := cli.GetAppRoleUserV2(ctx, &npool.GetAppRoleUserRequest{
			ID: id,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get app role user: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail get app role user: %v", err)
	}
	return info.(*npool.AppRoleUser), nil
}

func GetAppRoleUserOnlyV2(ctx context.Context, conds *npool.Conds) (*npool.AppRoleUser, error) {
	info, err := doAppRoleUser(ctx, func(_ctx context.Context, cli npool.AppUserManagerAppRoleUserClient) (cruder.Any, error) {
		resp, err := cli.GetAppRoleUserOnlyV2(ctx, &npool.GetAppRoleUserOnlyRequest{
			Conds: conds,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get app role user: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail get app role user: %v", err)
	}
	return info.(*npool.AppRoleUser), nil
}

func GetAppRoleUsersV2(ctx context.Context, conds *npool.Conds) ([]*npool.AppRoleUser, error) {
	infos, err := doAppRoleUser(ctx, func(_ctx context.Context, cli npool.AppUserManagerAppRoleUserClient) (cruder.Any, error) {
		resp, err := cli.GetAppRoleUsersV2(ctx, &npool.GetAppRoleUsersRequest{
			Conds: conds,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get app role user: %v", err)
		}
		return resp.GetInfos(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail get app role user: %v", err)
	}
	return infos.([]*npool.AppRoleUser), nil
}

func ExistAppRoleUserV2(ctx context.Context, id string) (bool, error) {
	infos, err := doAppRoleUser(ctx, func(_ctx context.Context, cli npool.AppUserManagerAppRoleUserClient) (cruder.Any, error) {
		resp, err := cli.ExistAppRoleUserV2(ctx, &npool.ExistAppRoleUserRequest{
			ID: id,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get app role user: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return false, fmt.Errorf("fail get app role user: %v", err)
	}
	return infos.(bool), nil
}

func ExistAppRoleUserCondsV2(ctx context.Context, conds *npool.Conds) (bool, error) {
	infos, err := doAppRoleUser(ctx, func(_ctx context.Context, cli npool.AppUserManagerAppRoleUserClient) (cruder.Any, error) {
		resp, err := cli.ExistAppRoleUserCondsV2(ctx, &npool.ExistAppRoleUserCondsRequest{
			Conds: conds,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get app role user: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return false, fmt.Errorf("fail get app role user: %v", err)
	}
	return infos.(bool), nil
}

func CountAppRoleUsersV2(ctx context.Context, conds *npool.Conds) (uint32, error) {
	infos, err := doAppRoleUser(ctx, func(_ctx context.Context, cli npool.AppUserManagerAppRoleUserClient) (cruder.Any, error) {
		resp, err := cli.CountAppRoleUsersV2(ctx, &npool.CountAppRoleUsersRequest{
			Conds: conds,
		})
		if err != nil {
			return nil, fmt.Errorf("fail count app role user: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return 0, fmt.Errorf("fail count app role user: %v", err)
	}
	return infos.(uint32), nil
}

func DeleteAppRoleUserV2(ctx context.Context, id string) (*npool.AppRoleUser, error) {
	infos, err := doAppRoleUser(ctx, func(_ctx context.Context, cli npool.AppUserManagerAppRoleUserClient) (cruder.Any, error) {
		resp, err := cli.DeleteAppRoleUserV2(ctx, &npool.DeleteAppRoleUserRequest{
			ID: id,
		})
		if err != nil {
			return nil, fmt.Errorf("fail delete app role user: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail delete app role user: %v", err)
	}
	return infos.(*npool.AppRoleUser), nil
}
