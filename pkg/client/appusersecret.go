//nolint:dupl
package client

import (
	"context"
	"fmt"
	"time"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/appusermgrv2/appusersecret"

	constant "github.com/NpoolPlatform/appuser-manager/pkg/message/const"
)

func doAppUserSecret(ctx context.Context, fn func(_ctx context.Context, cli npool.AppUserManagerAppUserSecretClient) (cruder.Any, error)) (cruder.Any, error) {
	_ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	conn, err := grpc2.GetGRPCConn(constant.ServiceName, grpc2.GRPCTAG)
	if err != nil {
		return nil, fmt.Errorf("fail get app user secret connection: %v", err)
	}

	defer conn.Close()

	cli := npool.NewAppUserManagerAppUserSecretClient(conn)

	return fn(_ctx, cli)
}

func CreateAppUserSecretV2(ctx context.Context, in *npool.AppUserSecretReq) (*npool.AppUserSecret, error) {
	info, err := doAppUserSecret(ctx, func(_ctx context.Context, cli npool.AppUserManagerAppUserSecretClient) (cruder.Any, error) {
		resp, err := cli.CreateAppUserSecretV2(ctx, &npool.CreateAppUserSecretRequest{
			Info: in,
		})
		if err != nil {
			return nil, fmt.Errorf("fail create app user secret: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail create app user secret: %v", err)
	}
	return info.(*npool.AppUserSecret), nil
}

func CreateAppUserSecretsV2(ctx context.Context, in []*npool.AppUserSecretReq) ([]*npool.AppUserSecret, error) {
	infos, err := doAppUserSecret(ctx, func(_ctx context.Context, cli npool.AppUserManagerAppUserSecretClient) (cruder.Any, error) {
		resp, err := cli.CreateAppUserSecretsV2(ctx, &npool.CreateAppUserSecretsRequest{
			Infos: in,
		})
		if err != nil {
			return nil, fmt.Errorf("fail create app user secrets: %v", err)
		}
		return resp.GetInfos(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail create app user secrets: %v", err)
	}
	return infos.([]*npool.AppUserSecret), nil
}

func UpdateAppUserSecretV2(ctx context.Context, in *npool.AppUserSecretReq) (*npool.AppUserSecret, error) {
	info, err := doAppUserSecret(ctx, func(_ctx context.Context, cli npool.AppUserManagerAppUserSecretClient) (cruder.Any, error) {
		resp, err := cli.UpdateAppUserSecretV2(ctx, &npool.UpdateAppUserSecretRequest{
			Info: in,
		})
		if err != nil {
			return nil, fmt.Errorf("fail update app user secret: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail update app user secret: %v", err)
	}
	return info.(*npool.AppUserSecret), nil
}

func GetAppUserSecretV2(ctx context.Context, id string) (*npool.AppUserSecret, error) {
	info, err := doAppUserSecret(ctx, func(_ctx context.Context, cli npool.AppUserManagerAppUserSecretClient) (cruder.Any, error) {
		resp, err := cli.GetAppUserSecretV2(ctx, &npool.GetAppUserSecretRequest{
			ID: id,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get app user secret: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail get app user secret: %v", err)
	}
	return info.(*npool.AppUserSecret), nil
}

func GetAppUserSecretOnlyV2(ctx context.Context, conds *npool.Conds) (*npool.AppUserSecret, error) {
	info, err := doAppUserSecret(ctx, func(_ctx context.Context, cli npool.AppUserManagerAppUserSecretClient) (cruder.Any, error) {
		resp, err := cli.GetAppUserSecretOnlyV2(ctx, &npool.GetAppUserSecretOnlyRequest{
			Conds: conds,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get app user secret: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail get app user secret: %v", err)
	}
	return info.(*npool.AppUserSecret), nil
}

func GetAppUserSecretsV2(ctx context.Context, conds *npool.Conds) ([]*npool.AppUserSecret, error) {
	infos, err := doAppUserSecret(ctx, func(_ctx context.Context, cli npool.AppUserManagerAppUserSecretClient) (cruder.Any, error) {
		resp, err := cli.GetAppUserSecretsV2(ctx, &npool.GetAppUserSecretsRequest{
			Conds: conds,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get app user secret: %v", err)
		}
		return resp.GetInfos(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail get app user secret: %v", err)
	}
	return infos.([]*npool.AppUserSecret), nil
}

func ExistAppUserSecretV2(ctx context.Context, id string) (bool, error) {
	infos, err := doAppUserSecret(ctx, func(_ctx context.Context, cli npool.AppUserManagerAppUserSecretClient) (cruder.Any, error) {
		resp, err := cli.ExistAppUserSecretV2(ctx, &npool.ExistAppUserSecretRequest{
			ID: id,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get app user secret: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return false, fmt.Errorf("fail get app user secret: %v", err)
	}
	return infos.(bool), nil
}

func ExistAppUserSecretCondsV2(ctx context.Context, conds *npool.Conds) (bool, error) {
	infos, err := doAppUserSecret(ctx, func(_ctx context.Context, cli npool.AppUserManagerAppUserSecretClient) (cruder.Any, error) {
		resp, err := cli.ExistAppUserSecretCondsV2(ctx, &npool.ExistAppUserSecretCondsRequest{
			Conds: conds,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get app user secret: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return false, fmt.Errorf("fail get app user secret: %v", err)
	}
	return infos.(bool), nil
}

func CountAppUserSecretsV2(ctx context.Context, conds *npool.Conds) (uint32, error) {
	infos, err := doAppUserSecret(ctx, func(_ctx context.Context, cli npool.AppUserManagerAppUserSecretClient) (cruder.Any, error) {
		resp, err := cli.CountAppUserSecretsV2(ctx, &npool.CountAppUserSecretsRequest{
			Conds: conds,
		})
		if err != nil {
			return nil, fmt.Errorf("fail count app user secret: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return 0, fmt.Errorf("fail count app user secret: %v", err)
	}
	return infos.(uint32), nil
}

func DeleteAppUserSecretV2(ctx context.Context, id string) (*npool.AppUserSecret, error) {
	infos, err := doAppUserSecret(ctx, func(_ctx context.Context, cli npool.AppUserManagerAppUserSecretClient) (cruder.Any, error) {
		resp, err := cli.DeleteAppUserSecretV2(ctx, &npool.DeleteAppUserSecretRequest{
			ID: id,
		})
		if err != nil {
			return nil, fmt.Errorf("fail delete app user secret: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail delete app user secret: %v", err)
	}
	return infos.(*npool.AppUserSecret), nil
}
