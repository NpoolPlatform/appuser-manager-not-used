//nolint:nolintlint,dupl
package appuser

import (
	"context"
	"time"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/appuser/mgr/v2/appuser"

	constant "github.com/NpoolPlatform/appuser-manager/pkg/message/const"
)

var timeout = 10 * time.Second

type handler func(context.Context, npool.AppUserMgrClient) (cruder.Any, error)

func withCRUD(ctx context.Context, handler handler) (cruder.Any, error) {
	_ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	conn, err := grpc2.GetGRPCConn(constant.ServiceName, grpc2.GRPCTAG)
	if err != nil {
		return nil, err
	}

	defer conn.Close()

	cli := npool.NewAppUserMgrClient(conn)

	return handler(_ctx, cli)
}

func CreateAppUser(ctx context.Context, in *npool.AppUserReq) (*npool.AppUser, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.AppUserMgrClient) (cruder.Any, error) {
		resp, err := cli.CreateAppUser(ctx, &npool.CreateAppUserRequest{
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
	return info.(*npool.AppUser), nil
}

func CreateAppUsers(ctx context.Context, in []*npool.AppUserReq) ([]*npool.AppUser, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.AppUserMgrClient) (cruder.Any, error) {
		resp, err := cli.CreateAppUsers(ctx, &npool.CreateAppUsersRequest{
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
	return infos.([]*npool.AppUser), nil
}

func UpdateAppUser(ctx context.Context, in *npool.AppUserReq) (*npool.AppUser, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.AppUserMgrClient) (cruder.Any, error) {
		resp, err := cli.UpdateAppUser(ctx, &npool.UpdateAppUserRequest{
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
	return info.(*npool.AppUser), nil
}

func GetAppUser(ctx context.Context, id string) (*npool.AppUser, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.AppUserMgrClient) (cruder.Any, error) {
		resp, err := cli.GetAppUser(ctx, &npool.GetAppUserRequest{
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
	return info.(*npool.AppUser), nil
}

func GetAppUserOnly(ctx context.Context, conds *npool.Conds) (*npool.AppUser, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.AppUserMgrClient) (cruder.Any, error) {
		resp, err := cli.GetAppUserOnly(ctx, &npool.GetAppUserOnlyRequest{
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
	return info.(*npool.AppUser), nil
}

func GetAppUsers(ctx context.Context, conds *npool.Conds, limit, offset int32) ([]*npool.AppUser, uint32, error) {
	var total uint32
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.AppUserMgrClient) (cruder.Any, error) {
		resp, err := cli.GetAppUsers(ctx, &npool.GetAppUsersRequest{
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
	return infos.([]*npool.AppUser), total, nil
}

func ExistAppUser(ctx context.Context, id string) (bool, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.AppUserMgrClient) (cruder.Any, error) {
		resp, err := cli.ExistAppUser(ctx, &npool.ExistAppUserRequest{
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

func ExistAppUserConds(ctx context.Context, conds *npool.Conds) (bool, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.AppUserMgrClient) (cruder.Any, error) {
		resp, err := cli.ExistAppUserConds(ctx, &npool.ExistAppUserCondsRequest{
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

func CountAppUsers(ctx context.Context, conds *npool.Conds) (uint32, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.AppUserMgrClient) (cruder.Any, error) {
		resp, err := cli.CountAppUsers(ctx, &npool.CountAppUsersRequest{
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

func DeleteAppUser(ctx context.Context, id string) (*npool.AppUser, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.AppUserMgrClient) (cruder.Any, error) {
		resp, err := cli.DeleteAppUser(ctx, &npool.DeleteAppUserRequest{
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
	return info.(*npool.AppUser), nil
}
