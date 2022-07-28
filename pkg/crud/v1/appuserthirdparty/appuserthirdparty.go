package appuserthirdparty

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent/appuserthirdparty"
	npool "github.com/NpoolPlatform/message/npool/appuser/mgr/v1"

	constant "github.com/NpoolPlatform/appuser-manager/pkg/const"
	db "github.com/NpoolPlatform/appuser-manager/pkg/db"
	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent"
	"github.com/google/uuid"
)

func validateAppUserThirdParty(info *npool.AppUserThirdParty) error {
	if _, err := uuid.Parse(info.GetAppID()); err != nil {
		return fmt.Errorf("invalid app id: %v", err)
	}
	if _, err := uuid.Parse(info.GetUserID()); err != nil {
		return fmt.Errorf("invalid user id: %v", err)
	}
	if info.GetThirdPartyID() == "" {
		return fmt.Errorf("invalid third party id")
	}
	if info.GetThirdPartyUserID() == "" {
		return fmt.Errorf("invalid third party user id")
	}
	return nil
}

func dbRowToAppUserThirdParty(row *ent.AppUserThirdParty) *npool.AppUserThirdParty {
	return &npool.AppUserThirdParty{
		ID:                   row.ID.String(),
		AppID:                row.AppID.String(),
		UserID:               row.UserID.String(),
		ThirdPartyUserID:     row.ThirdPartyUserID,
		ThirdPartyID:         row.ThirdPartyID,
		ThirdPartyUsername:   row.ThirdPartyUsername,
		ThirdPartyUserAvatar: row.ThirdPartyUserAvatar,
	}
}

func Create(ctx context.Context, in *npool.CreateAppUserThirdPartyRequest) (*npool.CreateAppUserThirdPartyResponse, error) {
	if err := validateAppUserThirdParty(in.GetInfo()); err != nil {
		return nil, fmt.Errorf("invalid parameter: %v", err)
	}

	ctx, cancel := context.WithTimeout(ctx, constant.DBTimeout)
	defer cancel()

	cli, err := db.Client()
	if err != nil {
		return nil, fmt.Errorf("fail get db client: %v", err)
	}
	info, err := cli.
		AppUserThirdParty.
		Create().
		SetAppID(uuid.MustParse(in.GetInfo().GetAppID())).
		SetUserID(uuid.MustParse(in.GetInfo().GetUserID())).
		SetThirdPartyUsername(in.GetInfo().GetThirdPartyUsername()).
		SetThirdPartyUserAvatar(in.GetInfo().GetThirdPartyUserAvatar()).
		SetThirdPartyID(in.GetInfo().GetThirdPartyID()).
		SetThirdPartyUserID(in.GetInfo().GetThirdPartyUserID()).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("fail create app user third party: %v", err)
	}

	return &npool.CreateAppUserThirdPartyResponse{
		Info: dbRowToAppUserThirdParty(info),
	}, nil
}

func GetByAppUserThirdParty(ctx context.Context, in *npool.GetAppUserThirdPartyByAppThirdPartyIDRequest) (*npool.GetAppUserThirdPartyByAppThirdPartyIDResponse, error) {
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
		AppUserThirdParty.
		Query().
		Where(
			appuserthirdparty.And(
				appuserthirdparty.AppID(appID),
				appuserthirdparty.ThirdPartyID(in.GetThirdPartyID()),
				appuserthirdparty.ThirdPartyUserID(in.GetThirdPartyUserID()),
			),
		).
		All(ctx)
	if err != nil {
		return nil, fmt.Errorf("fail query app user third party: %v", err)
	}

	var appUserThirdParty *npool.AppUserThirdParty
	for _, info := range infos {
		appUserThirdParty = dbRowToAppUserThirdParty(info)
		break
	}

	return &npool.GetAppUserThirdPartyByAppThirdPartyIDResponse{
		Info: appUserThirdParty,
	}, nil
}
