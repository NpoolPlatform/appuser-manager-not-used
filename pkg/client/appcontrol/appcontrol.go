//nolint:nolintlint,dupl
package appcontrol

import (
	"context"
	"time"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/appuser/mgr/v2/appcontrol"

	constant "github.com/NpoolPlatform/appuser-manager/pkg/message/const"
)

var timeout = 10 * time.Second

type handler func(context.Context, npool.AppControlMgrClient) (cruder.Any, error)

func withCRUD(ctx context.Context, handler handler) (cruder.Any, error) {
	_ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	conn, err := grpc2.GetGRPCConn(constant.ServiceName, grpc2.GRPCTAG)
	if err != nil {
		return nil, err
	}

	defer conn.Close()

	cli := npool.NewAppControlMgrClient(conn)

	return handler(_ctx, cli)
}

func CreateAppControl(ctx context.Context, in *npool.AppControlReq) (*npool.AppControl, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.AppControlMgrClient) (cruder.Any, error) {
		resp, err := cli.CreateAppControl(ctx, &npool.CreateAppControlRequest{
			Info: in,
		})
		if err != nil {
			return nil, err
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, err
	}
	return info.(*npool.AppControl), nil
}

func CreateAppControls(ctx context.Context, in []*npool.AppControlReq) ([]*npool.AppControl, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.AppControlMgrClient) (cruder.Any, error) {
		resp, err := cli.CreateAppControls(ctx, &npool.CreateAppControlsRequest{
			Infos: in,
		})
		if err != nil {
			return nil, err
		}
		return resp.GetInfos(), nil
	})
	if err != nil {
		return nil, err
	}
	return infos.([]*npool.AppControl), nil
}

func GetAppControl(ctx context.Context, id string) (*npool.AppControl, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.AppControlMgrClient) (cruder.Any, error) {
		resp, err := cli.GetAppControl(ctx, &npool.GetAppControlRequest{
			ID: id,
		})
		if err != nil {
			return nil, err
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, err
	}
	return info.(*npool.AppControl), nil
}

func GetAppControlOnly(ctx context.Context, conds *npool.Conds) (*npool.AppControl, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.AppControlMgrClient) (cruder.Any, error) {
		resp, err := cli.GetAppControlOnly(ctx, &npool.GetAppControlOnlyRequest{
			Conds: conds,
		})
		if err != nil {
			return nil, err
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, err
	}
	return info.(*npool.AppControl), nil
}

func GetAppControls(ctx context.Context, conds *npool.Conds, limit, offset int32) ([]*npool.AppControl, uint32, error) {
	var total uint32
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.AppControlMgrClient) (cruder.Any, error) {
		resp, err := cli.GetAppControls(ctx, &npool.GetAppControlsRequest{
			Conds:  conds,
			Limit:  limit,
			Offset: offset,
		})
		if err != nil {
			return nil, err
		}
		total = resp.GetTotal()
		return resp.GetInfos(), nil
	})
	if err != nil {
		return nil, 0, err
	}
	return infos.([]*npool.AppControl), total, nil
}

func ExistAppControl(ctx context.Context, id string) (bool, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.AppControlMgrClient) (cruder.Any, error) {
		resp, err := cli.ExistAppControl(ctx, &npool.ExistAppControlRequest{
			ID: id,
		})
		if err != nil {
			return nil, err
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return false, err
	}
	return infos.(bool), nil
}

func ExistAppControlConds(ctx context.Context, conds *npool.Conds) (bool, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.AppControlMgrClient) (cruder.Any, error) {
		resp, err := cli.ExistAppControlConds(ctx, &npool.ExistAppControlCondsRequest{
			Conds: conds,
		})
		if err != nil {
			return nil, err
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return false, err
	}
	return infos.(bool), nil
}

func CountAppControls(ctx context.Context, conds *npool.Conds) (uint32, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.AppControlMgrClient) (cruder.Any, error) {
		resp, err := cli.CountAppControls(ctx, &npool.CountAppControlsRequest{
			Conds: conds,
		})
		if err != nil {
			return nil, err
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return 0, err
	}
	return infos.(uint32), nil
}
