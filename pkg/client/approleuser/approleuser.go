//nolint:nolintlint,dupl
package approleuser

import (
	"context"
	"time"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/appuser/mgr/v2/approleuser"

	constant "github.com/NpoolPlatform/appuser-manager/pkg/message/const"
)

var timeout = 10 * time.Second

type handler func(context.Context, npool.AppRoleUserMgrClient) (cruder.Any, error)

func withCRUD(ctx context.Context, handler handler) (cruder.Any, error) {
	_ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	conn, err := grpc2.GetGRPCConn(constant.ServiceName, grpc2.GRPCTAG)
	if err != nil {
		return nil, err
	}

	defer conn.Close()

	cli := npool.NewAppRoleUserMgrClient(conn)

	return handler(_ctx, cli)
}

func CreateAppRoleUser(ctx context.Context, in *npool.AppRoleUserReq) (*npool.AppRoleUser, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.AppRoleUserMgrClient) (cruder.Any, error) {
		resp, err := cli.CreateAppRoleUser(ctx, &npool.CreateAppRoleUserRequest{
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
	return info.(*npool.AppRoleUser), nil
}

func CreateAppRoleUsers(ctx context.Context, in []*npool.AppRoleUserReq) ([]*npool.AppRoleUser, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.AppRoleUserMgrClient) (cruder.Any, error) {
		resp, err := cli.CreateAppRoleUsers(ctx, &npool.CreateAppRoleUsersRequest{
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
	return infos.([]*npool.AppRoleUser), nil
}

func GetAppRoleUser(ctx context.Context, id string) (*npool.AppRoleUser, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.AppRoleUserMgrClient) (cruder.Any, error) {
		resp, err := cli.GetAppRoleUser(ctx, &npool.GetAppRoleUserRequest{
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
	return info.(*npool.AppRoleUser), nil
}

func GetAppRoleUserOnly(ctx context.Context, conds *npool.Conds) (*npool.AppRoleUser, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.AppRoleUserMgrClient) (cruder.Any, error) {
		resp, err := cli.GetAppRoleUserOnly(ctx, &npool.GetAppRoleUserOnlyRequest{
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
	return info.(*npool.AppRoleUser), nil
}

func GetAppRoleUsers(ctx context.Context, conds *npool.Conds, limit, offset int32) ([]*npool.AppRoleUser, uint32, error) {
	var total uint32
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.AppRoleUserMgrClient) (cruder.Any, error) {
		resp, err := cli.GetAppRoleUsers(ctx, &npool.GetAppRoleUsersRequest{
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
	return infos.([]*npool.AppRoleUser), total, nil
}

func ExistAppRoleUser(ctx context.Context, id string) (bool, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.AppRoleUserMgrClient) (cruder.Any, error) {
		resp, err := cli.ExistAppRoleUser(ctx, &npool.ExistAppRoleUserRequest{
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

func ExistAppRoleUserConds(ctx context.Context, conds *npool.Conds) (bool, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.AppRoleUserMgrClient) (cruder.Any, error) {
		resp, err := cli.ExistAppRoleUserConds(ctx, &npool.ExistAppRoleUserCondsRequest{
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

func CountAppRoleUsers(ctx context.Context, conds *npool.Conds) (uint32, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.AppRoleUserMgrClient) (cruder.Any, error) {
		resp, err := cli.CountAppRoleUsers(ctx, &npool.CountAppRoleUsersRequest{
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
