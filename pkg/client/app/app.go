//nolint:dupl
package app

import (
	"context"
	"fmt"
	"time"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/appuser/mgr/v2/app"

	constant "github.com/NpoolPlatform/appuser-manager/pkg/message/const"
)

var timeout = 10 * time.Second

type handler func(context.Context, npool.AppMgrClient) (cruder.Any, error)

func withCRUD(ctx context.Context, handler handler) (cruder.Any, error) {
	_ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	conn, err := grpc2.GetGRPCConn(constant.ServiceName, grpc2.GRPCTAG)
	if err != nil {
		return nil, fmt.Errorf("fail get app connection: %v", err)
	}

	defer conn.Close()

	cli := npool.NewAppMgrClient(conn)

	return handler(_ctx, cli)
}

func CreateApp(ctx context.Context, in *npool.AppReq) (*npool.App, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.AppMgrClient) (cruder.Any, error) {
		resp, err := cli.CreateApp(ctx, &npool.CreateAppRequest{
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

func CreateApps(ctx context.Context, in []*npool.AppReq) ([]*npool.App, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.AppMgrClient) (cruder.Any, error) {
		resp, err := cli.CreateApps(ctx, &npool.CreateAppsRequest{
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

func GetApp(ctx context.Context, id string) (*npool.App, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.AppMgrClient) (cruder.Any, error) {
		resp, err := cli.GetApp(ctx, &npool.GetAppRequest{
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

func GetAppOnly(ctx context.Context, conds *npool.Conds) (*npool.App, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.AppMgrClient) (cruder.Any, error) {
		resp, err := cli.GetAppOnly(ctx, &npool.GetAppOnlyRequest{
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

func GetApps(ctx context.Context, conds *npool.Conds, limit, offset int32) ([]*npool.App, uint32, error) {
	var total uint32
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.AppMgrClient) (cruder.Any, error) {
		resp, err := cli.GetApps(ctx, &npool.GetAppsRequest{
			Conds:  conds,
			Limit:  limit,
			Offset: offset,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get apps: %v", err)
		}
		total = resp.GetTotal()
		return resp.GetInfos(), nil
	})
	if err != nil {
		return nil, 0, fmt.Errorf("fail get apps: %v", err)
	}
	return infos.([]*npool.App), total, nil
}

func ExistApp(ctx context.Context, id string) (bool, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.AppMgrClient) (cruder.Any, error) {
		resp, err := cli.ExistApp(ctx, &npool.ExistAppRequest{
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

func ExistAppConds(ctx context.Context, conds *npool.Conds) (bool, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.AppMgrClient) (cruder.Any, error) {
		resp, err := cli.ExistAppConds(ctx, &npool.ExistAppCondsRequest{
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

func CountApps(ctx context.Context, conds *npool.Conds) (uint32, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.AppMgrClient) (cruder.Any, error) {
		resp, err := cli.CountApps(ctx, &npool.CountAppsRequest{
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
