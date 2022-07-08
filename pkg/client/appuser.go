//nolint:dupl
package client

import (
	"context"
	"fmt"
	"time"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/appusermgrv2/appuser"

	constant "github.com/NpoolPlatform/appuser-manager/pkg/message/const"
)

func doAppUser(ctx context.Context, fn func(_ctx context.Context, cli npool.AppUserManagerAppUserClient) (cruder.Any, error)) (cruder.Any, error) {
	_ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	conn, err := grpc2.GetGRPCConn(constant.ServiceName, grpc2.GRPCTAG)
	if err != nil {
		return nil, fmt.Errorf("fail get app user connection: %v", err)
	}

	defer conn.Close()

	cli := npool.NewAppUserManagerAppUserClient(conn)

	return fn(_ctx, cli)
}

func CreateAppUserV2(ctx context.Context, in *npool.AppUserReq) (*npool.AppUser, error) {
	info, err := doAppUser(ctx, func(_ctx context.Context, cli npool.AppUserManagerAppUserClient) (cruder.Any, error) {
		resp, err := cli.CreateAppUserV2(ctx, &npool.CreateAppUserRequest{
			Info: in,
		})
		if err != nil {
			return nil, fmt.Errorf("fail create app user: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail create app user: %v", err)
	}
	return info.(*npool.AppUser), nil
}

func CreateAppUsersV2(ctx context.Context, in []*npool.AppUserReq) ([]*npool.AppUser, error) {
	infos, err := doAppUser(ctx, func(_ctx context.Context, cli npool.AppUserManagerAppUserClient) (cruder.Any, error) {
		resp, err := cli.CreateAppUsersV2(ctx, &npool.CreateAppUsersRequest{
			Infos: in,
		})
		if err != nil {
			return nil, fmt.Errorf("fail create app users: %v", err)
		}
		return resp.GetInfos(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail create app users: %v", err)
	}
	return infos.([]*npool.AppUser), nil
}

func UpdateAppUserV2(ctx context.Context, in *npool.AppUserReq) (*npool.AppUser, error) {
	info, err := doAppUser(ctx, func(_ctx context.Context, cli npool.AppUserManagerAppUserClient) (cruder.Any, error) {
		resp, err := cli.UpdateAppUserV2(ctx, &npool.UpdateAppUserRequest{
			Info: in,
		})
		if err != nil {
			return nil, fmt.Errorf("fail update app user: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail update app user: %v", err)
	}
	return info.(*npool.AppUser), nil
}

func GetAppUserV2(ctx context.Context, id string) (*npool.AppUser, error) {
	info, err := doAppUser(ctx, func(_ctx context.Context, cli npool.AppUserManagerAppUserClient) (cruder.Any, error) {
		resp, err := cli.GetAppUserV2(ctx, &npool.GetAppUserRequest{
			ID: id,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get app user: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail get app user: %v", err)
	}
	return info.(*npool.AppUser), nil
}

func GetAppUserOnlyV2(ctx context.Context, conds *npool.Conds) (*npool.AppUser, error) {
	info, err := doAppUser(ctx, func(_ctx context.Context, cli npool.AppUserManagerAppUserClient) (cruder.Any, error) {
		resp, err := cli.GetAppUserOnlyV2(ctx, &npool.GetAppUserOnlyRequest{
			Conds: conds,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get app user: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail get app user: %v", err)
	}
	return info.(*npool.AppUser), nil
}

func GetAppUsersV2(ctx context.Context, conds *npool.Conds) ([]*npool.AppUser, error) {
	infos, err := doAppUser(ctx, func(_ctx context.Context, cli npool.AppUserManagerAppUserClient) (cruder.Any, error) {
		resp, err := cli.GetAppUsersV2(ctx, &npool.GetAppUsersRequest{
			Conds: conds,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get app user: %v", err)
		}
		return resp.GetInfos(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail get app user: %v", err)
	}
	return infos.([]*npool.AppUser), nil
}

func ExistAppUserV2(ctx context.Context, id string) (bool, error) {
	infos, err := doAppUser(ctx, func(_ctx context.Context, cli npool.AppUserManagerAppUserClient) (cruder.Any, error) {
		resp, err := cli.ExistAppUserV2(ctx, &npool.ExistAppUserRequest{
			ID: id,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get app user: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return false, fmt.Errorf("fail get app user: %v", err)
	}
	return infos.(bool), nil
}

func ExistAppUserCondsV2(ctx context.Context, conds *npool.Conds) (bool, error) {
	infos, err := doAppUser(ctx, func(_ctx context.Context, cli npool.AppUserManagerAppUserClient) (cruder.Any, error) {
		resp, err := cli.ExistAppUserCondsV2(ctx, &npool.ExistAppUserCondsRequest{
			Conds: conds,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get app user: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return false, fmt.Errorf("fail get app user: %v", err)
	}
	return infos.(bool), nil
}

func CountAppUsersV2(ctx context.Context, conds *npool.Conds) (uint32, error) {
	infos, err := doAppUser(ctx, func(_ctx context.Context, cli npool.AppUserManagerAppUserClient) (cruder.Any, error) {
		resp, err := cli.CountAppUsersV2(ctx, &npool.CountAppUsersRequest{
			Conds: conds,
		})
		if err != nil {
			return nil, fmt.Errorf("fail count app user: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return 0, fmt.Errorf("fail count app user: %v", err)
	}
	return infos.(uint32), nil
}

func DeleteAppUserV2(ctx context.Context, id string) (*npool.AppUser, error) {
	infos, err := doAppUser(ctx, func(_ctx context.Context, cli npool.AppUserManagerAppUserClient) (cruder.Any, error) {
		resp, err := cli.DeleteAppUserV2(ctx, &npool.DeleteAppUserRequest{
			ID: id,
		})
		if err != nil {
			return nil, fmt.Errorf("fail delete app user: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail delete app user: %v", err)
	}
	return infos.(*npool.AppUser), nil
}
