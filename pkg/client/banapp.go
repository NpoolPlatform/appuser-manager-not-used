//nolint:dupl
package client

import (
	"context"
	"fmt"
	"time"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/appusermgrv2/banapp"

	constant "github.com/NpoolPlatform/appuser-manager/pkg/message/const"
)

func doBanApp(ctx context.Context, fn func(_ctx context.Context, cli npool.AppUserManagerBanAppClient) (cruder.Any, error)) (cruder.Any, error) {
	_ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	conn, err := grpc2.GetGRPCConn(constant.ServiceName, grpc2.GRPCTAG)
	if err != nil {
		return nil, fmt.Errorf("fail get ban app connection: %v", err)
	}

	defer conn.Close()

	cli := npool.NewAppUserManagerBanAppClient(conn)

	return fn(_ctx, cli)
}

func CreateBanAppV2(ctx context.Context, in *npool.BanAppReq) (*npool.BanApp, error) {
	info, err := doBanApp(ctx, func(_ctx context.Context, cli npool.AppUserManagerBanAppClient) (cruder.Any, error) {
		resp, err := cli.CreateBanAppV2(ctx, &npool.CreateBanAppRequest{
			Info: in,
		})
		if err != nil {
			return nil, fmt.Errorf("fail create ban app: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail create ban app: %v", err)
	}
	return info.(*npool.BanApp), nil
}

func CreateBanAppsV2(ctx context.Context, in []*npool.BanAppReq) ([]*npool.BanApp, error) {
	infos, err := doBanApp(ctx, func(_ctx context.Context, cli npool.AppUserManagerBanAppClient) (cruder.Any, error) {
		resp, err := cli.CreateBanAppsV2(ctx, &npool.CreateBanAppsRequest{
			Infos: in,
		})
		if err != nil {
			return nil, fmt.Errorf("fail create ban apps: %v", err)
		}
		return resp.GetInfos(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail create ban apps: %v", err)
	}
	return infos.([]*npool.BanApp), nil
}

func UpdateBanAppV2(ctx context.Context, in *npool.BanAppReq) (*npool.BanApp, error) {
	info, err := doBanApp(ctx, func(_ctx context.Context, cli npool.AppUserManagerBanAppClient) (cruder.Any, error) {
		resp, err := cli.UpdateBanAppV2(ctx, &npool.UpdateBanAppRequest{
			Info: in,
		})
		if err != nil {
			return nil, fmt.Errorf("fail update ban app: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail update ban app: %v", err)
	}
	return info.(*npool.BanApp), nil
}

func GetBanAppV2(ctx context.Context, id string) (*npool.BanApp, error) {
	info, err := doBanApp(ctx, func(_ctx context.Context, cli npool.AppUserManagerBanAppClient) (cruder.Any, error) {
		resp, err := cli.GetBanAppV2(ctx, &npool.GetBanAppRequest{
			ID: id,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get ban app: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail get ban app: %v", err)
	}
	return info.(*npool.BanApp), nil
}

func GetBanAppOnlyV2(ctx context.Context, conds *npool.Conds) (*npool.BanApp, error) {
	info, err := doBanApp(ctx, func(_ctx context.Context, cli npool.AppUserManagerBanAppClient) (cruder.Any, error) {
		resp, err := cli.GetBanAppOnlyV2(ctx, &npool.GetBanAppOnlyRequest{
			Conds: conds,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get ban app: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail get ban app: %v", err)
	}
	return info.(*npool.BanApp), nil
}

func GetBanAppsV2(ctx context.Context, conds *npool.Conds) ([]*npool.BanApp, error) {
	infos, err := doBanApp(ctx, func(_ctx context.Context, cli npool.AppUserManagerBanAppClient) (cruder.Any, error) {
		resp, err := cli.GetBanAppsV2(ctx, &npool.GetBanAppsRequest{
			Conds: conds,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get ban app: %v", err)
		}
		return resp.GetInfos(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail get ban app: %v", err)
	}
	return infos.([]*npool.BanApp), nil
}

func ExistBanAppV2(ctx context.Context, id string) (bool, error) {
	infos, err := doBanApp(ctx, func(_ctx context.Context, cli npool.AppUserManagerBanAppClient) (cruder.Any, error) {
		resp, err := cli.ExistBanAppV2(ctx, &npool.ExistBanAppRequest{
			ID: id,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get ban app: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return false, fmt.Errorf("fail get ban app: %v", err)
	}
	return infos.(bool), nil
}

func ExistBanAppCondsV2(ctx context.Context, conds *npool.Conds) (bool, error) {
	infos, err := doBanApp(ctx, func(_ctx context.Context, cli npool.AppUserManagerBanAppClient) (cruder.Any, error) {
		resp, err := cli.ExistBanAppCondsV2(ctx, &npool.ExistBanAppCondsRequest{
			Conds: conds,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get ban app: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return false, fmt.Errorf("fail get ban app: %v", err)
	}
	return infos.(bool), nil
}

func CountBanAppsV2(ctx context.Context, conds *npool.Conds) (uint32, error) {
	infos, err := doBanApp(ctx, func(_ctx context.Context, cli npool.AppUserManagerBanAppClient) (cruder.Any, error) {
		resp, err := cli.CountBanAppsV2(ctx, &npool.CountBanAppsRequest{
			Conds: conds,
		})
		if err != nil {
			return nil, fmt.Errorf("fail count ban app: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return 0, fmt.Errorf("fail count ban app: %v", err)
	}
	return infos.(uint32), nil
}

func DeleteBanAppV2(ctx context.Context, id string) (*npool.BanApp, error) {
	infos, err := doBanApp(ctx, func(_ctx context.Context, cli npool.AppUserManagerBanAppClient) (cruder.Any, error) {
		resp, err := cli.DeleteBanAppV2(ctx, &npool.DeleteBanAppRequest{
			ID: id,
		})
		if err != nil {
			return nil, fmt.Errorf("fail delete ban app: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail delete ban app: %v", err)
	}
	return infos.(*npool.BanApp), nil
}
