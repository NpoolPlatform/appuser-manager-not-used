//nolint:nolintlint,dupl
package appuserthirdparty

import (
	"context"
	"time"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/appuser/mgr/v2/appuserthirdparty"

	constant "github.com/NpoolPlatform/appuser-manager/pkg/message/const"
)

var timeout = 10 * time.Second

type handler func(context.Context, npool.AppUserThirdPartyMgrClient) (cruder.Any, error)

func withCRUD(ctx context.Context, handler handler) (cruder.Any, error) {
	_ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	conn, err := grpc2.GetGRPCConn(constant.ServiceName, grpc2.GRPCTAG)
	if err != nil {
		return nil, err
	}

	defer conn.Close()

	cli := npool.NewAppUserThirdPartyMgrClient(conn)

	return handler(_ctx, cli)
}

func CreateAppUserThirdParty(ctx context.Context, in *npool.AppUserThirdPartyReq) (*npool.AppUserThirdParty, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.AppUserThirdPartyMgrClient) (cruder.Any, error) {
		resp, err := cli.CreateAppUserThirdParty(ctx, &npool.CreateAppUserThirdPartyRequest{
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
	return info.(*npool.AppUserThirdParty), nil
}

func CreateAppUserThirdParties(ctx context.Context, in []*npool.AppUserThirdPartyReq) ([]*npool.AppUserThirdParty, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.AppUserThirdPartyMgrClient) (cruder.Any, error) {
		resp, err := cli.CreateAppUserThirdParties(ctx, &npool.CreateAppUserThirdPartiesRequest{
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
	return infos.([]*npool.AppUserThirdParty), nil
}

func GetAppUserThirdParty(ctx context.Context, id string) (*npool.AppUserThirdParty, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.AppUserThirdPartyMgrClient) (cruder.Any, error) {
		resp, err := cli.GetAppUserThirdParty(ctx, &npool.GetAppUserThirdPartyRequest{
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
	return info.(*npool.AppUserThirdParty), nil
}

func GetAppUserThirdPartyOnly(ctx context.Context, conds *npool.Conds) (*npool.AppUserThirdParty, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.AppUserThirdPartyMgrClient) (cruder.Any, error) {
		resp, err := cli.GetAppUserThirdPartyOnly(ctx, &npool.GetAppUserThirdPartyOnlyRequest{
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
	return info.(*npool.AppUserThirdParty), nil
}

func GetAppUserThirdParties(ctx context.Context, conds *npool.Conds, limit, offset int32) ([]*npool.AppUserThirdParty, uint32, error) {
	var total uint32
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.AppUserThirdPartyMgrClient) (cruder.Any, error) {
		resp, err := cli.GetAppUserThirdParties(ctx, &npool.GetAppUserThirdPartiesRequest{
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
	return infos.([]*npool.AppUserThirdParty), total, nil
}

func ExistAppUserThirdParty(ctx context.Context, id string) (bool, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.AppUserThirdPartyMgrClient) (cruder.Any, error) {
		resp, err := cli.ExistAppUserThirdParty(ctx, &npool.ExistAppUserThirdPartyRequest{
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

func ExistAppUserThirdPartyConds(ctx context.Context, conds *npool.Conds) (bool, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.AppUserThirdPartyMgrClient) (cruder.Any, error) {
		resp, err := cli.ExistAppUserThirdPartyConds(ctx, &npool.ExistAppUserThirdPartyCondsRequest{
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

func CountAppUserThirdParties(ctx context.Context, conds *npool.Conds) (uint32, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.AppUserThirdPartyMgrClient) (cruder.Any, error) {
		resp, err := cli.CountAppUserThirdParties(ctx, &npool.CountAppUserThirdPartiesRequest{
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
