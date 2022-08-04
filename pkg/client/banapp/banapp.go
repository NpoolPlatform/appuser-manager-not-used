//nolint:nolintlint,dupl
package banapp

import (
	"context"
	"time"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/appuser/mgr/v2/banapp"

	constant "github.com/NpoolPlatform/appuser-manager/pkg/message/const"
)

var timeout = 10 * time.Second

type handler func(context.Context, npool.BanAppMgrClient) (cruder.Any, error)

func withCRUD(ctx context.Context, handler handler) (cruder.Any, error) {
	_ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	conn, err := grpc2.GetGRPCConn(constant.ServiceName, grpc2.GRPCTAG)
	if err != nil {
		return nil, err
	}

	defer conn.Close()

	cli := npool.NewBanAppMgrClient(conn)

	return handler(_ctx, cli)
}

func CreateBanApp(ctx context.Context, in *npool.BanAppReq) (*npool.BanApp, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.BanAppMgrClient) (cruder.Any, error) {
		resp, err := cli.CreateBanApp(ctx, &npool.CreateBanAppRequest{
			Info: in,
		})
		if err != nil {
			return nil, err
		}
		return resp.Info, nil
	})
	if err != nil {
		return nil, err
	}
	return info.(*npool.BanApp), nil
}

func CreateBanApps(ctx context.Context, in []*npool.BanAppReq) ([]*npool.BanApp, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.BanAppMgrClient) (cruder.Any, error) {
		resp, err := cli.CreateBanApps(ctx, &npool.CreateBanAppsRequest{
			Infos: in,
		})
		if err != nil {
			return nil, err
		}
		return resp.Infos, nil
	})
	if err != nil {
		return nil, err
	}
	return infos.([]*npool.BanApp), nil
}

func GetBanApp(ctx context.Context, id string) (*npool.BanApp, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.BanAppMgrClient) (cruder.Any, error) {
		resp, err := cli.GetBanApp(ctx, &npool.GetBanAppRequest{
			ID: id,
		})
		if err != nil {
			return nil, err
		}
		return resp.Info, nil
	})
	if err != nil {
		return nil, err
	}
	return info.(*npool.BanApp), nil
}

func GetBanAppOnly(ctx context.Context, conds *npool.Conds) (*npool.BanApp, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.BanAppMgrClient) (cruder.Any, error) {
		resp, err := cli.GetBanAppOnly(ctx, &npool.GetBanAppOnlyRequest{
			Conds: conds,
		})
		if err != nil {
			return nil, err
		}
		return resp.Info, nil
	})
	if err != nil {
		return nil, err
	}
	return info.(*npool.BanApp), nil
}

func GetBanApps(ctx context.Context, conds *npool.Conds, limit, offset int32) ([]*npool.BanApp, uint32, error) {
	var total uint32
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.BanAppMgrClient) (cruder.Any, error) {
		resp, err := cli.GetBanApps(ctx, &npool.GetBanAppsRequest{
			Conds:  conds,
			Limit:  limit,
			Offset: offset,
		})
		if err != nil {
			return nil, err
		}
		total = resp.GetTotal()
		return resp.Infos, nil
	})
	if err != nil {
		return nil, 0, err
	}
	return infos.([]*npool.BanApp), total, nil
}

func ExistBanApp(ctx context.Context, id string) (bool, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.BanAppMgrClient) (cruder.Any, error) {
		resp, err := cli.ExistBanApp(ctx, &npool.ExistBanAppRequest{
			ID: id,
		})
		if err != nil {
			return nil, err
		}
		return resp.Info, nil
	})
	if err != nil {
		return false, err
	}
	return infos.(bool), nil
}

func ExistBanAppConds(ctx context.Context, conds *npool.Conds) (bool, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.BanAppMgrClient) (cruder.Any, error) {
		resp, err := cli.ExistBanAppConds(ctx, &npool.ExistBanAppCondsRequest{
			Conds: conds,
		})
		if err != nil {
			return nil, err
		}
		return resp.Info, nil
	})
	if err != nil {
		return false, err
	}
	return infos.(bool), nil
}

func CountBanApps(ctx context.Context, conds *npool.Conds) (uint32, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.BanAppMgrClient) (cruder.Any, error) {
		resp, err := cli.CountBanApps(ctx, &npool.CountBanAppsRequest{
			Conds: conds,
		})
		if err != nil {
			return nil, err
		}
		return resp.Info, nil
	})
	if err != nil {
		return 0, err
	}
	return infos.(uint32), nil
}
