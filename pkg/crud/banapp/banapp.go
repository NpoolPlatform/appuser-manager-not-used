package banapp

import (
	"context"
	"time"

	constant "github.com/NpoolPlatform/appuser-manager/pkg/const"
	npool "github.com/NpoolPlatform/message/npool/appusermgr"

	db "github.com/NpoolPlatform/appuser-manager/pkg/db"
	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent"
	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent/banapp"

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
		SetMessage(in.GetInfo().GetMessage()).
		Save(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail create ban app: %v", err)
	}

	return &npool.CreateBanAppResponse{
		Info: dbRowToBanApp(info),
	}, nil
}

func Get(ctx context.Context, in *npool.GetBanAppRequest) (*npool.GetBanAppResponse, error) {
	id, err := uuid.Parse(in.GetID())
	if err != nil {
		return nil, xerrors.Errorf("invalid id: %v", err)
	}

	ctx, cancel := context.WithTimeout(ctx, constant.DBTimeout)
	defer cancel()

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	infos, err := cli.
		BanApp.
		Query().
		Where(
			banapp.And(
				banapp.ID(id),
				banapp.DeleteAt(0),
			),
		).
		All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail query ban app: %v", err)
	}

	var banApp *npool.BanApp
	for _, info := range infos {
		banApp = dbRowToBanApp(info)
		break
	}

	return &npool.GetBanAppResponse{
		Info: banApp,
	}, nil
}

func GetByApp(ctx context.Context, in *npool.GetBanAppByAppRequest) (*npool.GetBanAppByAppResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, constant.DBTimeout)
	defer cancel()

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	appID, err := uuid.Parse(in.GetAppID())
	if err != nil {
		return nil, xerrors.Errorf("invalid app id: %v", err)
	}

	infos, err := cli.
		BanApp.
		Query().
		Where(
			banapp.And(
				banapp.AppID(appID),
				banapp.DeleteAt(0),
			),
		).
		All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail query ban app by app: %v", err)
	}

	var banApp *npool.BanApp
	for _, info := range infos {
		banApp = dbRowToBanApp(info)
		break
	}

	return &npool.GetBanAppByAppResponse{
		Info: banApp,
	}, nil
}

func Delete(ctx context.Context, in *npool.DeleteBanAppRequest) (*npool.DeleteBanAppResponse, error) {
	id, err := uuid.Parse(in.GetID())
	if err != nil {
		return nil, xerrors.Errorf("invalid id: %v", err)
	}

	ctx, cancel := context.WithTimeout(ctx, constant.DBTimeout)
	defer cancel()

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	info, err := cli.
		BanApp.
		UpdateOneID(id).
		SetDeleteAt(uint32(time.Now().Unix())).
		Save(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail delete ban app: %v", err)
	}

	return &npool.DeleteBanAppResponse{
		Info: dbRowToBanApp(info),
	}, nil
}
