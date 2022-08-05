//nolint:nolintlint,dupl
package appuserextra

import (
	"context"
	"time"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/appuser/mgr/v2/appuserextra"

	constant "github.com/NpoolPlatform/appuser-manager/pkg/message/const"
)

var timeout = 10 * time.Second

type handler func(context.Context, npool.AppUserExtraMgrClient) (cruder.Any, error)

func withCRUD(ctx context.Context, handler handler) (cruder.Any, error) {
	_ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	conn, err := grpc2.GetGRPCConn(constant.ServiceName, grpc2.GRPCTAG)
	if err != nil {
		return nil, err
	}

	defer conn.Close()

	cli := npool.NewAppUserExtraMgrClient(conn)

	return handler(_ctx, cli)
}

func CreateAppUserExtra(ctx context.Context, in *npool.AppUserExtraReq) (*npool.AppUserExtra, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.AppUserExtraMgrClient) (cruder.Any, error) {
		resp, err := cli.CreateAppUserExtra(ctx, &npool.CreateAppUserExtraRequest{
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
	return info.(*npool.AppUserExtra), nil
}

func CreateAppUserExtras(ctx context.Context, in []*npool.AppUserExtraReq) ([]*npool.AppUserExtra, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.AppUserExtraMgrClient) (cruder.Any, error) {
		resp, err := cli.CreateAppUserExtras(ctx, &npool.CreateAppUserExtrasRequest{
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
	return infos.([]*npool.AppUserExtra), nil
}

func UpdateAppUserExtra(ctx context.Context, in *npool.AppUserExtraReq) (*npool.AppUserExtra, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.AppUserExtraMgrClient) (cruder.Any, error) {
		resp, err := cli.UpdateAppUserExtra(ctx, &npool.UpdateAppUserExtraRequest{
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
	return info.(*npool.AppUserExtra), nil
}

func GetAppUserExtra(ctx context.Context, id string) (*npool.AppUserExtra, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.AppUserExtraMgrClient) (cruder.Any, error) {
		resp, err := cli.GetAppUserExtra(ctx, &npool.GetAppUserExtraRequest{
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
	return info.(*npool.AppUserExtra), nil
}

func GetAppUserExtraOnly(ctx context.Context, conds *npool.Conds) (*npool.AppUserExtra, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.AppUserExtraMgrClient) (cruder.Any, error) {
		resp, err := cli.GetAppUserExtraOnly(ctx, &npool.GetAppUserExtraOnlyRequest{
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
	return info.(*npool.AppUserExtra), nil
}

func GetAppUserExtras(ctx context.Context, conds *npool.Conds, limit, offset int32) ([]*npool.AppUserExtra, uint32, error) {
	var total uint32
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.AppUserExtraMgrClient) (cruder.Any, error) {
		resp, err := cli.GetAppUserExtras(ctx, &npool.GetAppUserExtrasRequest{
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
	return infos.([]*npool.AppUserExtra), total, nil
}

func ExistAppUserExtra(ctx context.Context, id string) (bool, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.AppUserExtraMgrClient) (cruder.Any, error) {
		resp, err := cli.ExistAppUserExtra(ctx, &npool.ExistAppUserExtraRequest{
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

func ExistAppUserExtraConds(ctx context.Context, conds *npool.Conds) (bool, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.AppUserExtraMgrClient) (cruder.Any, error) {
		resp, err := cli.ExistAppUserExtraConds(ctx, &npool.ExistAppUserExtraCondsRequest{
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

func CountAppUserExtras(ctx context.Context, conds *npool.Conds) (uint32, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.AppUserExtraMgrClient) (cruder.Any, error) {
		resp, err := cli.CountAppUserExtras(ctx, &npool.CountAppUserExtrasRequest{
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

func DeleteAppUserExtra(ctx context.Context, id string) (*npool.AppUserExtra, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.AppUserExtraMgrClient) (cruder.Any, error) {
		resp, err := cli.DeleteAppUserExtra(ctx, &npool.DeleteAppUserExtraRequest{
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
	return info.(*npool.AppUserExtra), nil
}
