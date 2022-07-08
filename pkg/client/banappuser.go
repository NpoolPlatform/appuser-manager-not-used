//nolint:dupl
package client

import (
	"context"
	"fmt"
	"time"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/appusermgrv2/banappuser"

	constant "github.com/NpoolPlatform/appuser-manager/pkg/message/const"
)

func doBanAppUser(ctx context.Context, fn func(_ctx context.Context, cli npool.AppUserManagerBanAppUserClient) (cruder.Any, error)) (cruder.Any, error) {
	_ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	conn, err := grpc2.GetGRPCConn(constant.ServiceName, grpc2.GRPCTAG)
	if err != nil {
		return nil, fmt.Errorf("fail get ban app user connection: %v", err)
	}

	defer conn.Close()

	cli := npool.NewAppUserManagerBanAppUserClient(conn)

	return fn(_ctx, cli)
}

func CreateBanAppUserV2(ctx context.Context, in *npool.BanAppUserReq) (*npool.BanAppUser, error) {
	info, err := doBanAppUser(ctx, func(_ctx context.Context, cli npool.AppUserManagerBanAppUserClient) (cruder.Any, error) {
		resp, err := cli.CreateBanAppUserV2(ctx, &npool.CreateBanAppUserRequest{
			Info: in,
		})
		if err != nil {
			return nil, fmt.Errorf("fail create ban app user: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail create ban app user: %v", err)
	}
	return info.(*npool.BanAppUser), nil
}

func CreateBanAppUsersV2(ctx context.Context, in []*npool.BanAppUserReq) ([]*npool.BanAppUser, error) {
	infos, err := doBanAppUser(ctx, func(_ctx context.Context, cli npool.AppUserManagerBanAppUserClient) (cruder.Any, error) {
		resp, err := cli.CreateBanAppUsersV2(ctx, &npool.CreateBanAppUsersRequest{
			Infos: in,
		})
		if err != nil {
			return nil, fmt.Errorf("fail create ban app users: %v", err)
		}
		return resp.GetInfos(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail create ban app users: %v", err)
	}
	return infos.([]*npool.BanAppUser), nil
}

func UpdateBanAppUserV2(ctx context.Context, in *npool.BanAppUserReq) (*npool.BanAppUser, error) {
	info, err := doBanAppUser(ctx, func(_ctx context.Context, cli npool.AppUserManagerBanAppUserClient) (cruder.Any, error) {
		resp, err := cli.UpdateBanAppUserV2(ctx, &npool.UpdateBanAppUserRequest{
			Info: in,
		})
		if err != nil {
			return nil, fmt.Errorf("fail update ban app user: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail update ban app user: %v", err)
	}
	return info.(*npool.BanAppUser), nil
}

func GetBanAppUserV2(ctx context.Context, id string) (*npool.BanAppUser, error) {
	info, err := doBanAppUser(ctx, func(_ctx context.Context, cli npool.AppUserManagerBanAppUserClient) (cruder.Any, error) {
		resp, err := cli.GetBanAppUserV2(ctx, &npool.GetBanAppUserRequest{
			ID: id,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get ban app user: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail get ban app user: %v", err)
	}
	return info.(*npool.BanAppUser), nil
}

func GetBanAppUserOnlyV2(ctx context.Context, conds *npool.Conds) (*npool.BanAppUser, error) {
	info, err := doBanAppUser(ctx, func(_ctx context.Context, cli npool.AppUserManagerBanAppUserClient) (cruder.Any, error) {
		resp, err := cli.GetBanAppUserOnlyV2(ctx, &npool.GetBanAppUserOnlyRequest{
			Conds: conds,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get ban app user: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail get ban app user: %v", err)
	}
	return info.(*npool.BanAppUser), nil
}

func GetBanAppUsersV2(ctx context.Context, conds *npool.Conds) ([]*npool.BanAppUser, error) {
	infos, err := doBanAppUser(ctx, func(_ctx context.Context, cli npool.AppUserManagerBanAppUserClient) (cruder.Any, error) {
		resp, err := cli.GetBanAppUsersV2(ctx, &npool.GetBanAppUsersRequest{
			Conds: conds,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get ban app user: %v", err)
		}
		return resp.GetInfos(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail get ban app user: %v", err)
	}
	return infos.([]*npool.BanAppUser), nil
}

func ExistBanAppUserV2(ctx context.Context, id string) (bool, error) {
	infos, err := doBanAppUser(ctx, func(_ctx context.Context, cli npool.AppUserManagerBanAppUserClient) (cruder.Any, error) {
		resp, err := cli.ExistBanAppUserV2(ctx, &npool.ExistBanAppUserRequest{
			ID: id,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get ban app user: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return false, fmt.Errorf("fail get ban app user: %v", err)
	}
	return infos.(bool), nil
}

func ExistBanAppUserCondsV2(ctx context.Context, conds *npool.Conds) (bool, error) {
	infos, err := doBanAppUser(ctx, func(_ctx context.Context, cli npool.AppUserManagerBanAppUserClient) (cruder.Any, error) {
		resp, err := cli.ExistBanAppUserCondsV2(ctx, &npool.ExistBanAppUserCondsRequest{
			Conds: conds,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get ban app user: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return false, fmt.Errorf("fail get ban app user: %v", err)
	}
	return infos.(bool), nil
}

func CountBanAppUsersV2(ctx context.Context, conds *npool.Conds) (uint32, error) {
	infos, err := doBanAppUser(ctx, func(_ctx context.Context, cli npool.AppUserManagerBanAppUserClient) (cruder.Any, error) {
		resp, err := cli.CountBanAppUsersV2(ctx, &npool.CountBanAppUsersRequest{
			Conds: conds,
		})
		if err != nil {
			return nil, fmt.Errorf("fail count ban app user: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return 0, fmt.Errorf("fail count ban app user: %v", err)
	}
	return infos.(uint32), nil
}

func DeleteBanAppUserV2(ctx context.Context, id string) (*npool.BanAppUser, error) {
	infos, err := doBanAppUser(ctx, func(_ctx context.Context, cli npool.AppUserManagerBanAppUserClient) (cruder.Any, error) {
		resp, err := cli.DeleteBanAppUserV2(ctx, &npool.DeleteBanAppUserRequest{
			ID: id,
		})
		if err != nil {
			return nil, fmt.Errorf("fail delete ban app user: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail delete ban app user: %v", err)
	}
	return infos.(*npool.BanAppUser), nil
}
