package appuserthird

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent/appuserthird"
	npool "github.com/NpoolPlatform/message/npool/appusermgr"

	constant "github.com/NpoolPlatform/appuser-manager/pkg/const"
	db "github.com/NpoolPlatform/appuser-manager/pkg/db"
	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent"
	"github.com/google/uuid"
)

func validateAppUserThird(info *npool.AppUserThird) error {
	if _, err := uuid.Parse(info.GetAppID()); err != nil {
		return fmt.Errorf("invalid app id: %v", err)
	}
	if _, err := uuid.Parse(info.GetUserID()); err != nil {
		return fmt.Errorf("invalid user id: %v", err)
	}
	if info.GetThirdID() == "" {
		return fmt.Errorf("invalid third id")
	}
	if info.GetThirdUserID() == "" {
		return fmt.Errorf("invalid third user id")
	}
	return nil
}

func dbRowToAppUserThird(row *ent.AppUserThird) *npool.AppUserThird {
	return &npool.AppUserThird{
		ID:              row.ID.String(),
		AppID:           row.AppID.String(),
		UserID:          row.UserID.String(),
		ThirdUserID:     row.ThirdUserID,
		ThirdID:         row.ThirdID,
		ThirdUserName:   row.ThirdUserName,
		ThirdUserAvatar: row.ThirdUserAvatar,
		Third:           row.Third,
	}
}

func Create(ctx context.Context, in *npool.CreateAppUserThirdRequest) (*npool.CreateAppUserThirdResponse, error) {
	if err := validateAppUserThird(in.GetInfo()); err != nil {
		return nil, fmt.Errorf("invalid parameter: %v", err)
	}

	ctx, cancel := context.WithTimeout(ctx, constant.DBTimeout)
	defer cancel()

	cli, err := db.Client()
	if err != nil {
		return nil, fmt.Errorf("fail get db client: %v", err)
	}
	info, err := cli.
		AppUserThird.
		Create().
		SetAppID(uuid.MustParse(in.GetInfo().GetAppID())).
		SetUserID(uuid.MustParse(in.GetInfo().GetUserID())).
		SetThird(in.GetInfo().GetThird()).
		SetThirdUserName(in.GetInfo().GetThirdUserName()).
		SetThirdUserAvatar(in.GetInfo().GetThirdUserAvatar()).
		SetThirdID(in.GetInfo().GetThirdID()).
		SetThirdUserID(in.GetInfo().GetThirdUserID()).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("fail create app user third: %v", err)
	}

	return &npool.CreateAppUserThirdResponse{
		Info: dbRowToAppUserThird(info),
	}, nil
}

func GetByAppUserThird(ctx context.Context, in *npool.GetAppUserThirdByAppThirdRequest) (*npool.GetAppUserThirdByAppThirdResponse, error) {
	appID, err := uuid.Parse(in.GetAppID())
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
		AppUserThird.
		Query().
		Where(
			appuserthird.And(
				appuserthird.AppID(appID),
				appuserthird.ThirdID(in.GetThirdID()),
				appuserthird.ThirdUserID(in.GetThirdUserID()),
			),
		).
		All(ctx)
	if err != nil {
		return nil, fmt.Errorf("fail query app user third: %v", err)
	}

	var appUserthird *npool.AppUserThird
	for _, info := range infos {
		appUserthird = dbRowToAppUserThird(info)
		break
	}

	return &npool.GetAppUserThirdByAppThirdResponse{
		Info: appUserthird,
	}, nil
}
