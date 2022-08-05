//nolint:nolintlint,dupl
package appusercontrol

import (
	"context"
	"time"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/appuser/mgr/v2/appusercontrol"

	constant "github.com/NpoolPlatform/appuser-manager/pkg/message/const"
)

var timeout = 10 * time.Second

type handler func(context.Context, npool.AppUserControlMgrClient) (cruder.Any, error)

func withCRUD(ctx context.Context, handler handler) (cruder.Any, error) {
	_ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	conn, err := grpc2.GetGRPCConn(constant.ServiceName, grpc2.GRPCTAG)
	if err != nil {
		return nil, err
	}

	defer conn.Close()

	cli := npool.NewAppUserControlMgrClient(conn)

	return handler(_ctx, cli)
}

func CreateAppUserControl(ctx context.Context, in *npool.AppUserControlReq) (*npool.AppUserControl, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.AppUserControlMgrClient) (cruder.Any, error) {
		resp, err := cli.CreateAppUserControl(ctx, &npool.CreateAppUserControlRequest{
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
	return info.(*npool.AppUserControl), nil
}

func CreateAppUserControls(ctx context.Context, in []*npool.AppUserControlReq) ([]*npool.AppUserControl, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.AppUserControlMgrClient) (cruder.Any, error) {
		resp, err := cli.CreateAppUserControls(ctx, &npool.CreateAppUserControlsRequest{
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
	return infos.([]*npool.AppUserControl), nil
}

func UpdateAppUserControl(ctx context.Context, in *npool.AppUserControlReq) (*npool.AppUserControl, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.AppUserControlMgrClient) (cruder.Any, error) {
		resp, err := cli.UpdateAppUserControl(ctx, &npool.UpdateAppUserControlRequest{
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
	return info.(*npool.AppUserControl), nil
}

func GetAppUserControl(ctx context.Context, id string) (*npool.AppUserControl, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.AppUserControlMgrClient) (cruder.Any, error) {
		resp, err := cli.GetAppUserControl(ctx, &npool.GetAppUserControlRequest{
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
	return info.(*npool.AppUserControl), nil
}

func GetAppUserControlOnly(ctx context.Context, conds *npool.Conds) (*npool.AppUserControl, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.AppUserControlMgrClient) (cruder.Any, error) {
		resp, err := cli.GetAppUserControlOnly(ctx, &npool.GetAppUserControlOnlyRequest{
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
	return info.(*npool.AppUserControl), nil
}

func GetAppUserControls(ctx context.Context, conds *npool.Conds, limit, offset int32) ([]*npool.AppUserControl, uint32, error) {
	var total uint32
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.AppUserControlMgrClient) (cruder.Any, error) {
		resp, err := cli.GetAppUserControls(ctx, &npool.GetAppUserControlsRequest{
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
	return infos.([]*npool.AppUserControl), total, nil
}

func ExistAppUserControl(ctx context.Context, id string) (bool, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.AppUserControlMgrClient) (cruder.Any, error) {
		resp, err := cli.ExistAppUserControl(ctx, &npool.ExistAppUserControlRequest{
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

func ExistAppUserControlConds(ctx context.Context, conds *npool.Conds) (bool, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.AppUserControlMgrClient) (cruder.Any, error) {
		resp, err := cli.ExistAppUserControlConds(ctx, &npool.ExistAppUserControlCondsRequest{
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

func CountAppUserControls(ctx context.Context, conds *npool.Conds) (uint32, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.AppUserControlMgrClient) (cruder.Any, error) {
		resp, err := cli.CountAppUserControls(ctx, &npool.CountAppUserControlsRequest{
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

func DeleteAppUserControl(ctx context.Context, id string) (*npool.AppUserControl, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.AppUserControlMgrClient) (cruder.Any, error) {
		resp, err := cli.DeleteAppUserControl(ctx, &npool.DeleteAppUserControlRequest{
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
	return info.(*npool.AppUserControl), nil
}
