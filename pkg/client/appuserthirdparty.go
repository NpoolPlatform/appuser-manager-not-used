//nolint:dupl
package client

import (
	"context"
	"fmt"
	"time"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/appusermgrv2/appuserthirdparty"

	constant "github.com/NpoolPlatform/appuser-manager/pkg/message/const"
)

func doAppUserThirdParty(ctx context.Context, fn func(_ctx context.Context, cli npool.AppUserManagerAppUserThirdPartyClient) (cruder.Any, error)) (cruder.Any, error) {
	_ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	conn, err := grpc2.GetGRPCConn(constant.ServiceName, grpc2.GRPCTAG)
	if err != nil {
		return nil, fmt.Errorf("fail get app user third party connection: %v", err)
	}

	defer conn.Close()

	cli := npool.NewAppUserManagerAppUserThirdPartyClient(conn)

	return fn(_ctx, cli)
}

func CreateAppUserThirdPartyV2(ctx context.Context, in *npool.AppUserThirdPartyReq) (*npool.AppUserThirdParty, error) {
	info, err := doAppUserThirdParty(ctx, func(_ctx context.Context, cli npool.AppUserManagerAppUserThirdPartyClient) (cruder.Any, error) {
		resp, err := cli.CreateAppUserThirdPartyV2(ctx, &npool.CreateAppUserThirdPartyRequest{
			Info: in,
		})
		if err != nil {
			return nil, fmt.Errorf("fail create app user third party: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail create app user third party: %v", err)
	}
	return info.(*npool.AppUserThirdParty), nil
}

func CreateAppUserThirdPartysV2(ctx context.Context, in []*npool.AppUserThirdPartyReq) ([]*npool.AppUserThirdParty, error) {
	infos, err := doAppUserThirdParty(ctx, func(_ctx context.Context, cli npool.AppUserManagerAppUserThirdPartyClient) (cruder.Any, error) {
		resp, err := cli.CreateAppUserThirdPartysV2(ctx, &npool.CreateAppUserThirdPartysRequest{
			Infos: in,
		})
		if err != nil {
			return nil, fmt.Errorf("fail create app user third partys: %v", err)
		}
		return resp.GetInfos(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail create app user third partys: %v", err)
	}
	return infos.([]*npool.AppUserThirdParty), nil
}

func UpdateAppUserThirdPartyV2(ctx context.Context, in *npool.AppUserThirdPartyReq) (*npool.AppUserThirdParty, error) {
	info, err := doAppUserThirdParty(ctx, func(_ctx context.Context, cli npool.AppUserManagerAppUserThirdPartyClient) (cruder.Any, error) {
		resp, err := cli.UpdateAppUserThirdPartyV2(ctx, &npool.UpdateAppUserThirdPartyRequest{
			Info: in,
		})
		if err != nil {
			return nil, fmt.Errorf("fail update app user third party: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail update app user third party: %v", err)
	}
	return info.(*npool.AppUserThirdParty), nil
}

func GetAppUserThirdPartyV2(ctx context.Context, id string) (*npool.AppUserThirdParty, error) {
	info, err := doAppUserThirdParty(ctx, func(_ctx context.Context, cli npool.AppUserManagerAppUserThirdPartyClient) (cruder.Any, error) {
		resp, err := cli.GetAppUserThirdPartyV2(ctx, &npool.GetAppUserThirdPartyRequest{
			ID: id,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get app user third party: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail get app user third party: %v", err)
	}
	return info.(*npool.AppUserThirdParty), nil
}

func GetAppUserThirdPartyOnlyV2(ctx context.Context, conds *npool.Conds) (*npool.AppUserThirdParty, error) {
	info, err := doAppUserThirdParty(ctx, func(_ctx context.Context, cli npool.AppUserManagerAppUserThirdPartyClient) (cruder.Any, error) {
		resp, err := cli.GetAppUserThirdPartyOnlyV2(ctx, &npool.GetAppUserThirdPartyOnlyRequest{
			Conds: conds,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get app user third party: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail get app user third party: %v", err)
	}
	return info.(*npool.AppUserThirdParty), nil
}

func GetAppUserThirdPartysV2(ctx context.Context, conds *npool.Conds) ([]*npool.AppUserThirdParty, error) {
	infos, err := doAppUserThirdParty(ctx, func(_ctx context.Context, cli npool.AppUserManagerAppUserThirdPartyClient) (cruder.Any, error) {
		resp, err := cli.GetAppUserThirdPartysV2(ctx, &npool.GetAppUserThirdPartysRequest{
			Conds: conds,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get app user third party: %v", err)
		}
		return resp.GetInfos(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail get app user third party: %v", err)
	}
	return infos.([]*npool.AppUserThirdParty), nil
}

func ExistAppUserThirdPartyV2(ctx context.Context, id string) (bool, error) {
	infos, err := doAppUserThirdParty(ctx, func(_ctx context.Context, cli npool.AppUserManagerAppUserThirdPartyClient) (cruder.Any, error) {
		resp, err := cli.ExistAppUserThirdPartyV2(ctx, &npool.ExistAppUserThirdPartyRequest{
			ID: id,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get app user third party: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return false, fmt.Errorf("fail get app user third party: %v", err)
	}
	return infos.(bool), nil
}

func ExistAppUserThirdPartyCondsV2(ctx context.Context, conds *npool.Conds) (bool, error) {
	infos, err := doAppUserThirdParty(ctx, func(_ctx context.Context, cli npool.AppUserManagerAppUserThirdPartyClient) (cruder.Any, error) {
		resp, err := cli.ExistAppUserThirdPartyCondsV2(ctx, &npool.ExistAppUserThirdPartyCondsRequest{
			Conds: conds,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get app user third party: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return false, fmt.Errorf("fail get app user third party: %v", err)
	}
	return infos.(bool), nil
}

func CountAppUserThirdPartysV2(ctx context.Context, conds *npool.Conds) (uint32, error) {
	infos, err := doAppUserThirdParty(ctx, func(_ctx context.Context, cli npool.AppUserManagerAppUserThirdPartyClient) (cruder.Any, error) {
		resp, err := cli.CountAppUserThirdPartysV2(ctx, &npool.CountAppUserThirdPartysRequest{
			Conds: conds,
		})
		if err != nil {
			return nil, fmt.Errorf("fail count app user third party: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return 0, fmt.Errorf("fail count app user third party: %v", err)
	}
	return infos.(uint32), nil
}

func DeleteAppUserThirdPartyV2(ctx context.Context, id string) (*npool.AppUserThirdParty, error) {
	infos, err := doAppUserThirdParty(ctx, func(_ctx context.Context, cli npool.AppUserManagerAppUserThirdPartyClient) (cruder.Any, error) {
		resp, err := cli.DeleteAppUserThirdPartyV2(ctx, &npool.DeleteAppUserThirdPartyRequest{
			ID: id,
		})
		if err != nil {
			return nil, fmt.Errorf("fail delete app user third party: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail delete app user third party: %v", err)
	}
	return infos.(*npool.AppUserThirdParty), nil
}
