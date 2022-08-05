//nolint:nolintlint,dupl
package appusersecret

import (
	"context"
	"time"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/appuser/mgr/v2/appusersecret"

	constant "github.com/NpoolPlatform/appuser-manager/pkg/message/const"
)

var timeout = 10 * time.Second

type handler func(context.Context, npool.AppUserSecretMgrClient) (cruder.Any, error)

func withCRUD(ctx context.Context, handler handler) (cruder.Any, error) {
	_ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	conn, err := grpc2.GetGRPCConn(constant.ServiceName, grpc2.GRPCTAG)
	if err != nil {
		return nil, err
	}

	defer conn.Close()

	cli := npool.NewAppUserSecretMgrClient(conn)

	return handler(_ctx, cli)
}

func CreateAppUserSecret(ctx context.Context, in *npool.AppUserSecretReq) (*npool.AppUserSecret, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.AppUserSecretMgrClient) (cruder.Any, error) {
		resp, err := cli.CreateAppUserSecret(ctx, &npool.CreateAppUserSecretRequest{
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
	return info.(*npool.AppUserSecret), nil
}

func CreateAppUserSecrets(ctx context.Context, in []*npool.AppUserSecretReq) ([]*npool.AppUserSecret, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.AppUserSecretMgrClient) (cruder.Any, error) {
		resp, err := cli.CreateAppUserSecrets(ctx, &npool.CreateAppUserSecretsRequest{
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
	return infos.([]*npool.AppUserSecret), nil
}

func UpdateAppUserSecret(ctx context.Context, in *npool.AppUserSecretReq) (*npool.AppUserSecret, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.AppUserSecretMgrClient) (cruder.Any, error) {
		resp, err := cli.UpdateAppUserSecret(ctx, &npool.UpdateAppUserSecretRequest{
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
	return info.(*npool.AppUserSecret), nil
}

func GetAppUserSecret(ctx context.Context, id string) (*npool.AppUserSecret, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.AppUserSecretMgrClient) (cruder.Any, error) {
		resp, err := cli.GetAppUserSecret(ctx, &npool.GetAppUserSecretRequest{
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
	return info.(*npool.AppUserSecret), nil
}

func GetAppUserSecretOnly(ctx context.Context, conds *npool.Conds) (*npool.AppUserSecret, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.AppUserSecretMgrClient) (cruder.Any, error) {
		resp, err := cli.GetAppUserSecretOnly(ctx, &npool.GetAppUserSecretOnlyRequest{
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
	return info.(*npool.AppUserSecret), nil
}

func GetAppUserSecrets(ctx context.Context, conds *npool.Conds, limit, offset int32) ([]*npool.AppUserSecret, uint32, error) {
	var total uint32
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.AppUserSecretMgrClient) (cruder.Any, error) {
		resp, err := cli.GetAppUserSecrets(ctx, &npool.GetAppUserSecretsRequest{
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
	return infos.([]*npool.AppUserSecret), total, nil
}

func ExistAppUserSecret(ctx context.Context, id string) (bool, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.AppUserSecretMgrClient) (cruder.Any, error) {
		resp, err := cli.ExistAppUserSecret(ctx, &npool.ExistAppUserSecretRequest{
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

func ExistAppUserSecretConds(ctx context.Context, conds *npool.Conds) (bool, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.AppUserSecretMgrClient) (cruder.Any, error) {
		resp, err := cli.ExistAppUserSecretConds(ctx, &npool.ExistAppUserSecretCondsRequest{
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

func CountAppUserSecrets(ctx context.Context, conds *npool.Conds) (uint32, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.AppUserSecretMgrClient) (cruder.Any, error) {
		resp, err := cli.CountAppUserSecrets(ctx, &npool.CountAppUserSecretsRequest{
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

func DeleteAppUserSecret(ctx context.Context, id string) (*npool.AppUserSecret, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.AppUserSecretMgrClient) (cruder.Any, error) {
		resp, err := cli.DeleteAppUserSecret(ctx, &npool.DeleteAppUserSecretRequest{
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
	return info.(*npool.AppUserSecret), nil
}
