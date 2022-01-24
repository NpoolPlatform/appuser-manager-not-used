package banapp

import (
	"context"

	constant "github.com/NpoolPlatform/appuser-manager/pkg/const"
	npool "github.com/NpoolPlatform/message/npool/appusermgr"

	db "github.com/NpoolPlatform/appuser-manager/pkg/db"
	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent"

	"github.com/google/uuid"

	"golang.org/x/xerrors"
)

func validateBanApp(info *npool.BanApp) error {
	if _, err := uuid.Parse(info.GetAppID()); err != nil {
		return xerrors.Errorf("invalid app id: %v", err)
	}
	return nil
}

func dbRowToBanApp(row *ent.BanApp) *npool.BanApp {
	return &npool.BanApp{
		ID:      row.ID.String(),
		AppID:   row.AppID.String(),
		Message: row.Message,
	}
}

func Create(ctx context.Context, in *npool.CreateBanAppRequest) (*npool.CreateBanAppResponse, error) {
	if err := validateBanApp(in.GetInfo()); err != nil {
		return nil, xerrors.Errorf("invalid parameter: %v", err)
	}

	ctx, cancel := context.WithTimeout(ctx, constant.DBTimeout)
	defer cancel()

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	info, err := cli.
		BanApp.
		Create().
		SetAppID(uuid.MustParse(in.GetInfo().GetAppID())).
		Save(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail create ban app: %v", err)
	}

	return &npool.CreateBanAppResponse{
		Info: dbRowToBanApp(info),
	}, nil
}

func Get(ctx context.Context, in *npool.GetBanAppRequest) (*npool.GetBanAppResponse, error) {
	return nil, nil
}

func GetByApp(ctx context.Context, in *npool.GetBanAppByAppRequest) (*npool.GetBanAppByAppResponse, error) {
	return nil, nil
}

func Delete(ctx context.Context, in *npool.DeleteBanAppRequest) (*npool.DeleteBanAppResponse, error) {
	return nil, nil
}
