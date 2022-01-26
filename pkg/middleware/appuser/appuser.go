package appuser

import (
	"context"

	appusercrud "github.com/NpoolPlatform/appuser-manager/pkg/crud/appuser"
	appusersecretcrud "github.com/NpoolPlatform/appuser-manager/pkg/crud/appusersecret"
	npool "github.com/NpoolPlatform/message/npool/appusermgr"

	"golang.org/x/xerrors"
)

func CreateWithSecret(ctx context.Context, in *npool.CreateAppUserWithSecretRequest) (*npool.CreateAppUserWithSecretResponse, error) {
	resp, err := appusercrud.Create(ctx, &npool.CreateAppUserRequest{
		Info: in.GetUser(),
	})
	if err != nil {
		return nil, xerrors.Errorf("fail create app user: %v", err)
	}

	inSecret := in.GetSecret()
	inSecret.UserID = resp.Info.ID

	_, err = appusersecretcrud.Create(ctx, &npool.CreateAppUserSecretRequest{
		Info: inSecret,
	})
	if err != nil {
		// TODO: rollback for secret create error
		return nil, xerrors.Errorf("fail create app user secret: %v", err)
	}

	return &npool.CreateAppUserWithSecretResponse{
		Info: resp.Info,
	}, nil
}
