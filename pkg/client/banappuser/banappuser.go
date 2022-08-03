//nolint:nolintlint,dupl
package banappuser

import (
	"context"
	"time"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/appuser/mgr/v2/banappuser"

	constant "github.com/NpoolPlatform/appuser-manager/pkg/message/const"
)

var timeout = 10 * time.Second

type handler func(context.Context, npool.BanAppUserMgrClient) (cruder.Any, error)

func withCRUD(ctx context.Context, handler handler) (cruder.Any, error) {
	_ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	conn, err := grpc2.GetGRPCConn(constant.ServiceName, grpc2.GRPCTAG)
	if err != nil {
		return nil, err
	}

	defer conn.Close()

	cli := npool.NewBanAppUserMgrClient(conn)

	return handler(_ctx, cli)
}

func CreateBanAppUser(ctx context.Context, in *npool.BanAppUserReq) (*npool.BanAppUser, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.BanAppUserMgrClient) (cruder.Any, error) {
		resp, err := cli.CreateBanAppUser(ctx, &npool.CreateBanAppUserRequest{
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
	return info.(*npool.BanAppUser), nil
}

func CreateBanAppUsers(ctx context.Context, in []*npool.BanAppUserReq) ([]*npool.BanAppUser, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.BanAppUserMgrClient) (cruder.Any, error) {
		resp, err := cli.CreateBanAppUsers(ctx, &npool.CreateBanAppUsersRequest{
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
	return infos.([]*npool.BanAppUser), nil
}

func GetBanAppUser(ctx context.Context, id string) (*npool.BanAppUser, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.BanAppUserMgrClient) (cruder.Any, error) {
		resp, err := cli.GetBanAppUser(ctx, &npool.GetBanAppUserRequest{
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
	return info.(*npool.BanAppUser), nil
}

func GetBanAppUserOnly(ctx context.Context, conds *npool.Conds) (*npool.BanAppUser, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.BanAppUserMgrClient) (cruder.Any, error) {
		resp, err := cli.GetBanAppUserOnly(ctx, &npool.GetBanAppUserOnlyRequest{
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
	return info.(*npool.BanAppUser), nil
}

func GetBanAppUsers(ctx context.Context, conds *npool.Conds, limit, offset int32) ([]*npool.BanAppUser, uint32, error) {
	var total uint32
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.BanAppUserMgrClient) (cruder.Any, error) {
		resp, err := cli.GetBanAppUsers(ctx, &npool.GetBanAppUsersRequest{
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
	return infos.([]*npool.BanAppUser), total, nil
}

func ExistBanAppUser(ctx context.Context, id string) (bool, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.BanAppUserMgrClient) (cruder.Any, error) {
		resp, err := cli.ExistBanAppUser(ctx, &npool.ExistBanAppUserRequest{
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

func ExistBanAppUserConds(ctx context.Context, conds *npool.Conds) (bool, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.BanAppUserMgrClient) (cruder.Any, error) {
		resp, err := cli.ExistBanAppUserConds(ctx, &npool.ExistBanAppUserCondsRequest{
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

func CountBanAppUsers(ctx context.Context, conds *npool.Conds) (uint32, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.BanAppUserMgrClient) (cruder.Any, error) {
		resp, err := cli.CountBanAppUsers(ctx, &npool.CountBanAppUsersRequest{
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
