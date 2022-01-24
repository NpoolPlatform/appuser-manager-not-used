package app

import (
	"context"

	npool "github.com/NpoolPlatform/message/npool/appusermgr"

	constant "github.com/NpoolPlatform/appuser-manager/pkg/const"
	db "github.com/NpoolPlatform/appuser-manager/pkg/db"
	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent"

	"github.com/google/uuid"

	"golang.org/x/xerrors"
)

func validateApp(info *npool.App) error {
	if _, err := uuid.Parse(info.GetCreatedBy()); err != nil {
		return xerrors.Errorf("invalid app creator: %v", err)
	}
	if info.GetName() == "" {
		return xerrors.Errorf("invalid app name")
	}
	if info.GetLogo() == "" {
		return xerrors.Errorf("invalid app logo")
	}
	if info.GetDescription() == "" {
		return xerrors.Errorf("invalid app description")
	}
	return nil
}

func dbRowToApp(row *ent.App) *npool.App {
	return &npool.App{
		ID:          row.ID.String(),
		CreatedBy:   row.CreatedBy.String(),
		Name:        row.Name,
		Logo:        row.Logo,
		Description: row.Description,
	}
}

func Create(ctx context.Context, in *npool.CreateAppRequest) (*npool.CreateAppResponse, error) {
	if err := validateApp(in.GetInfo()); err != nil {
		return nil, xerrors.Errorf("invalid parameter: %v", err)
	}

	ctx, cancel := context.WithTimeout(ctx, constant.DBTimeout)
	defer cancel()

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	info, err := cli.
		App.
		Create().
		SetCreatedBy(uuid.MustParse(in.GetInfo().GetCreatedBy())).
		SetName(in.GetInfo().GetName()).
		SetLogo(in.GetInfo().GetLogo()).
		SetDescription(in.GetInfo().GetDescription()).
		Save(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail create app: %v", err)
	}

	return &npool.CreateAppResponse{
		Info: dbRowToApp(info),
	}, nil
}

func Update(ctx context.Context, in *npool.UpdateAppRequest) (*npool.UpdateAppResponse, error) {
	return nil, nil
}
