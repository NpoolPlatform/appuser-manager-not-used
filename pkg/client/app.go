//nolint:dupl
package client

import (
	"context"
	"fmt"
	"time"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/appusermgrv2/app"

	constant "github.com/NpoolPlatform/appuser-manager/pkg/message/const"
)

func doApp(ctx context.Context, fn func(_ctx context.Context, cli npool.AppUserManagerAppClient) (cruder.Any, error)) (cruder.Any, error) {
	_ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	conn, err := grpc2.GetGRPCConn(constant.ServiceName, grpc2.GRPCTAG)
	if err != nil {
		return nil, fmt.Errorf("fail get app connection: %v", err)
	}

	defer conn.Close()

	cli := npool.NewAppUserManagerAppClient(conn)

	return fn(_ctx, cli)
}

func CreateAppV2(ctx context.Context, in *npool.AppReq) (*npool.App, error) {
	info, err := doApp(ctx, func(_ctx context.Context, cli npool.AppUserManagerAppClient) (cruder.Any, error) {
		resp, err := cli.CreateAppV2(ctx, &npool.CreateAppRequest{
			Info: in,
		})
		if err != nil {
			return nil, fmt.Errorf("fail create app: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail create app: %v", err)
	}
	return info.(*npool.App), nil
}

func CreateAppsV2(ctx context.Context, in []*npool.AppReq) ([]*npool.App, error) {
	infos, err := doApp(ctx, func(_ctx context.Context, cli npool.AppUserManagerAppClient) (cruder.Any, error) {
		resp, err := cli.CreateAppsV2(ctx, &npool.CreateAppsRequest{
			Infos: in,
		})
		if err != nil {
			return nil, fmt.Errorf("fail create apps: %v", err)
		}
		return resp.GetInfos(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail create apps: %v", err)
	}
	return infos.([]*npool.App), nil
}

func UpdateAppV2(ctx context.Context, in *npool.AppReq) (*npool.App, error) {
	info, err := doApp(ctx, func(_ctx context.Context, cli npool.AppUserManagerAppClient) (cruder.Any, error) {
		resp, err := cli.UpdateAppV2(ctx, &npool.UpdateAppRequest{
			Info: in,
		})
		if err != nil {
			return nil, fmt.Errorf("fail update app: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail update app: %v", err)
	}
	return info.(*npool.App), nil
}

func GetAppV2(ctx context.Context, id string) (*npool.App, error) {
	info, err := doApp(ctx, func(_ctx context.Context, cli npool.AppUserManagerAppClient) (cruder.Any, error) {
		resp, err := cli.GetAppV2(ctx, &npool.GetAppRequest{
			ID: id,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get app: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail get app: %v", err)
	}
	return info.(*npool.App), nil
}

func GetAppOnlyV2(ctx context.Context, conds *npool.Conds) (*npool.App, error) {
	info, err := doApp(ctx, func(_ctx context.Context, cli npool.AppUserManagerAppClient) (cruder.Any, error) {
		resp, err := cli.GetAppOnlyV2(ctx, &npool.GetAppOnlyRequest{
			Conds: conds,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get app: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail get app: %v", err)
	}
	return info.(*npool.App), nil
}

func GetAppsV2(ctx context.Context, conds *npool.Conds) ([]*npool.App, error) {
	infos, err := doApp(ctx, func(_ctx context.Context, cli npool.AppUserManagerAppClient) (cruder.Any, error) {
		resp, err := cli.GetAppsV2(ctx, &npool.GetAppsRequest{
			Conds: conds,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get apps: %v", err)
		}
		return resp.GetInfos(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail get apps: %v", err)
	}
	return infos.([]*npool.App), nil
}

func ExistAppV2(ctx context.Context, id string) (bool, error) {
	infos, err := doApp(ctx, func(_ctx context.Context, cli npool.AppUserManagerAppClient) (cruder.Any, error) {
		resp, err := cli.ExistAppV2(ctx, &npool.ExistAppRequest{
			ID: id,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get app: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return false, fmt.Errorf("fail get app: %v", err)
	}
	return infos.(bool), nil
}

func ExistAppCondsV2(ctx context.Context, conds *npool.Conds) (bool, error) {
	infos, err := doApp(ctx, func(_ctx context.Context, cli npool.AppUserManagerAppClient) (cruder.Any, error) {
		resp, err := cli.ExistAppCondsV2(ctx, &npool.ExistAppCondsRequest{
			Conds: conds,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get app: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return false, fmt.Errorf("fail get app: %v", err)
	}
	return infos.(bool), nil
}

func CountAppsV2(ctx context.Context, conds *npool.Conds) (uint32, error) {
	infos, err := doApp(ctx, func(_ctx context.Context, cli npool.AppUserManagerAppClient) (cruder.Any, error) {
		resp, err := cli.CountAppsV2(ctx, &npool.CountAppsRequest{
			Conds: conds,
		})
		if err != nil {
			return nil, fmt.Errorf("fail count app: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return 0, fmt.Errorf("fail count app: %v", err)
	}
	return infos.(uint32), nil
}

func DeleteAppV2(ctx context.Context, id string) (*npool.App, error) {
	infos, err := doApp(ctx, func(_ctx context.Context, cli npool.AppUserManagerAppClient) (cruder.Any, error) {
		resp, err := cli.DeleteAppV2(ctx, &npool.DeleteAppRequest{
			ID: id,
		})
		if err != nil {
			return nil, fmt.Errorf("fail delete app: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail delete app: %v", err)
	}
	return infos.(*npool.App), nil
}
