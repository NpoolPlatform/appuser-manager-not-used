package app

import (
	"context"
	"fmt"

	npool "github.com/NpoolPlatform/message/npool/appuser/mgr/v1"

	constant "github.com/NpoolPlatform/appuser-manager/pkg/const"
	db "github.com/NpoolPlatform/appuser-manager/pkg/db"
	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent"
	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent/app"

	"github.com/google/uuid"
)

func validateApp(info *npool.App) error {
	if _, err := uuid.Parse(info.GetCreatedBy()); err != nil {
		return fmt.Errorf("invalid app creator: %v", err)
	}
	if info.GetName() == "" {
		return fmt.Errorf("invalid app name")
	}
	if info.GetLogo() == "" {
		return fmt.Errorf("invalid app logo")
	}
	if info.GetDescription() == "" {
		return fmt.Errorf("invalid app description")
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

func Create(ctx context.Context, in *npool.CreateAppRequest, withID bool) (*npool.CreateAppResponse, error) {
	if err := validateApp(in.GetInfo()); err != nil {
		return nil, fmt.Errorf("invalid parameter: %v", err)
	}

	if withID {
		if _, err := uuid.Parse(in.GetInfo().GetID()); err != nil {
			return nil, fmt.Errorf("need id but invalid id: %v", err)
		}
	}

	ctx, cancel := context.WithTimeout(ctx, constant.DBTimeout)
	defer cancel()

	cli, err := db.Client()
	if err != nil {
		return nil, fmt.Errorf("fail get db client: %v", err)
	}

	rc := cli.
		App.
		Create().
		SetCreatedBy(uuid.MustParse(in.GetInfo().GetCreatedBy())).
		SetName(in.GetInfo().GetName()).
		SetLogo(in.GetInfo().GetLogo()).
		SetDescription(in.GetInfo().GetDescription())
	if withID {
		rc = rc.SetID(uuid.MustParse(in.GetInfo().GetID()))
	}

	info, err := rc.Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("fail create app: %v", err)
	}

	return &npool.CreateAppResponse{
		Info: dbRowToApp(info),
	}, nil
}

func Update(ctx context.Context, in *npool.UpdateAppRequest) (*npool.UpdateAppResponse, error) {
	id, err := uuid.Parse(in.GetInfo().GetID())
	if err != nil {
		return nil, fmt.Errorf("invalid app id: %v", err)
	}

	if err := validateApp(in.GetInfo()); err != nil {
		return nil, fmt.Errorf("invalid parameter: %v", err)
	}

	ctx, cancel := context.WithTimeout(ctx, constant.DBTimeout)
	defer cancel()

	cli, err := db.Client()
	if err != nil {
		return nil, fmt.Errorf("fail get db client: %v", err)
	}

	info, err := cli.
		App.
		UpdateOneID(id).
		SetName(in.GetInfo().GetName()).
		SetLogo(in.GetInfo().GetLogo()).
		SetDescription(in.GetInfo().GetDescription()).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("fail update app: %v", err)
	}

	return &npool.UpdateAppResponse{
		Info: dbRowToApp(info),
	}, nil
}

func Get(ctx context.Context, in *npool.GetAppRequest) (*npool.GetAppResponse, error) {
	id, err := uuid.Parse(in.GetID())
	if err != nil {
		return nil, fmt.Errorf("invalid app id: %v", err)
	}

	ctx, cancel := context.WithTimeout(ctx, constant.DBTimeout)
	defer cancel()

	cli, err := db.Client()
	if err != nil {
		return nil, fmt.Errorf("fail get db client: %v", err)
	}

	infos, err := cli.
		App.
		Query().
		Where(
			app.ID(id),
		).
		All(ctx)
	if err != nil {
		return nil, fmt.Errorf("fail query app: %v", err)
	}

	var myApp *npool.App
	for _, info := range infos {
		myApp = dbRowToApp(info)
		break
	}

	return &npool.GetAppResponse{
		Info: myApp,
	}, nil
}

func GetAll(ctx context.Context, in *npool.GetAppsRequest) (*npool.GetAppsResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, constant.DBTimeout)
	defer cancel()

	cli, err := db.Client()
	if err != nil {
		return nil, fmt.Errorf("fail get db client: %v", err)
	}

	infos, err := cli.
		App.
		Query().
		All(ctx)
	if err != nil {
		return nil, fmt.Errorf("fail query app: %v", err)
	}

	apps := []*npool.App{}
	for _, info := range infos {
		apps = append(apps, dbRowToApp(info))
	}

	return &npool.GetAppsResponse{
		Infos: apps,
	}, nil
}

func GetByCreator(ctx context.Context, in *npool.GetAppsByCreatorRequest) (*npool.GetAppsByCreatorResponse, error) {
	userID, err := uuid.Parse(in.GetUserID())
	if err != nil {
		return nil, fmt.Errorf("invalid user id: %v", err)
	}

	ctx, cancel := context.WithTimeout(ctx, constant.DBTimeout)
	defer cancel()

	cli, err := db.Client()
	if err != nil {
		return nil, fmt.Errorf("fail get db client: %v", err)
	}

	infos, err := cli.
		App.
		Query().
		Where(
			app.CreatedBy(userID),
		).
		All(ctx)
	if err != nil {
		return nil, fmt.Errorf("fail query app: %v", err)
	}

	apps := []*npool.App{}
	for _, info := range infos {
		apps = append(apps, dbRowToApp(info))
	}

	return &npool.GetAppsByCreatorResponse{
		Infos: apps,
	}, nil
}
