//nolint:dupl
package client

import (
	"context"
	"fmt"
	"time"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/appusermgrv2/appuserextra"

	constant "github.com/NpoolPlatform/appuser-manager/pkg/message/const"
)

func doAppUserExtra(ctx context.Context, fn func(_ctx context.Context, cli npool.AppUserManagerAppUserExtraClient) (cruder.Any, error)) (cruder.Any, error) {
	_ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	conn, err := grpc2.GetGRPCConn(constant.ServiceName, grpc2.GRPCTAG)
	if err != nil {
		return nil, fmt.Errorf("fail get app user extra connection: %v", err)
	}

	defer conn.Close()

	cli := npool.NewAppUserManagerAppUserExtraClient(conn)

	return fn(_ctx, cli)
}

func CreateAppUserExtraV2(ctx context.Context, in *npool.AppUserExtraReq) (*npool.AppUserExtra, error) {
	info, err := doAppUserExtra(ctx, func(_ctx context.Context, cli npool.AppUserManagerAppUserExtraClient) (cruder.Any, error) {
		resp, err := cli.CreateAppUserExtraV2(ctx, &npool.CreateAppUserExtraRequest{
			Info: in,
		})
		if err != nil {
			return nil, fmt.Errorf("fail create app user extra: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail create app user extra: %v", err)
	}
	return info.(*npool.AppUserExtra), nil
}

func CreateAppUserExtrasV2(ctx context.Context, in []*npool.AppUserExtraReq) ([]*npool.AppUserExtra, error) {
	infos, err := doAppUserExtra(ctx, func(_ctx context.Context, cli npool.AppUserManagerAppUserExtraClient) (cruder.Any, error) {
		resp, err := cli.CreateAppUserExtrasV2(ctx, &npool.CreateAppUserExtrasRequest{
			Infos: in,
		})
		if err != nil {
			return nil, fmt.Errorf("fail create app user extras: %v", err)
		}
		return resp.GetInfos(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail create app user extras: %v", err)
	}
	return infos.([]*npool.AppUserExtra), nil
}

func UpdateAppUserExtraV2(ctx context.Context, in *npool.AppUserExtraReq) (*npool.AppUserExtra, error) {
	info, err := doAppUserExtra(ctx, func(_ctx context.Context, cli npool.AppUserManagerAppUserExtraClient) (cruder.Any, error) {
		resp, err := cli.UpdateAppUserExtraV2(ctx, &npool.UpdateAppUserExtraRequest{
			Info: in,
		})
		if err != nil {
			return nil, fmt.Errorf("fail update app user extra: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail update app user extra: %v", err)
	}
	return info.(*npool.AppUserExtra), nil
}

func GetAppUserExtraV2(ctx context.Context, id string) (*npool.AppUserExtra, error) {
	info, err := doAppUserExtra(ctx, func(_ctx context.Context, cli npool.AppUserManagerAppUserExtraClient) (cruder.Any, error) {
		resp, err := cli.GetAppUserExtraV2(ctx, &npool.GetAppUserExtraRequest{
			ID: id,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get app user extra: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail get app user extra: %v", err)
	}
	return info.(*npool.AppUserExtra), nil
}

func GetAppUserExtraOnlyV2(ctx context.Context, conds *npool.Conds) (*npool.AppUserExtra, error) {
	info, err := doAppUserExtra(ctx, func(_ctx context.Context, cli npool.AppUserManagerAppUserExtraClient) (cruder.Any, error) {
		resp, err := cli.GetAppUserExtraOnlyV2(ctx, &npool.GetAppUserExtraOnlyRequest{
			Conds: conds,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get app user extra: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail get app user extra: %v", err)
	}
	return info.(*npool.AppUserExtra), nil
}

func GetAppUserExtrasV2(ctx context.Context, conds *npool.Conds) ([]*npool.AppUserExtra, error) {
	infos, err := doAppUserExtra(ctx, func(_ctx context.Context, cli npool.AppUserManagerAppUserExtraClient) (cruder.Any, error) {
		resp, err := cli.GetAppUserExtrasV2(ctx, &npool.GetAppUserExtrasRequest{
			Conds: conds,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get app user extra: %v", err)
		}
		return resp.GetInfos(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail get app user extra: %v", err)
	}
	return infos.([]*npool.AppUserExtra), nil
}

func ExistAppUserExtraV2(ctx context.Context, id string) (bool, error) {
	infos, err := doAppUserExtra(ctx, func(_ctx context.Context, cli npool.AppUserManagerAppUserExtraClient) (cruder.Any, error) {
		resp, err := cli.ExistAppUserExtraV2(ctx, &npool.ExistAppUserExtraRequest{
			ID: id,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get app user extra: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return false, fmt.Errorf("fail get app user extra: %v", err)
	}
	return infos.(bool), nil
}

func ExistAppUserExtraCondsV2(ctx context.Context, conds *npool.Conds) (bool, error) {
	infos, err := doAppUserExtra(ctx, func(_ctx context.Context, cli npool.AppUserManagerAppUserExtraClient) (cruder.Any, error) {
		resp, err := cli.ExistAppUserExtraCondsV2(ctx, &npool.ExistAppUserExtraCondsRequest{
			Conds: conds,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get app user extra: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return false, fmt.Errorf("fail get app user extra: %v", err)
	}
	return infos.(bool), nil
}

func CountAppUserExtrasV2(ctx context.Context, conds *npool.Conds) (uint32, error) {
	infos, err := doAppUserExtra(ctx, func(_ctx context.Context, cli npool.AppUserManagerAppUserExtraClient) (cruder.Any, error) {
		resp, err := cli.CountAppUserExtrasV2(ctx, &npool.CountAppUserExtrasRequest{
			Conds: conds,
		})
		if err != nil {
			return nil, fmt.Errorf("fail count app user extra: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return 0, fmt.Errorf("fail count app user extra: %v", err)
	}
	return infos.(uint32), nil
}

func DeleteAppUserExtraV2(ctx context.Context, id string) (*npool.AppUserExtra, error) {
	infos, err := doAppUserExtra(ctx, func(_ctx context.Context, cli npool.AppUserManagerAppUserExtraClient) (cruder.Any, error) {
		resp, err := cli.DeleteAppUserExtraV2(ctx, &npool.DeleteAppUserExtraRequest{
			ID: id,
		})
		if err != nil {
			return nil, fmt.Errorf("fail delete app user extra: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail delete app user extra: %v", err)
	}
	return infos.(*npool.AppUserExtra), nil
}
