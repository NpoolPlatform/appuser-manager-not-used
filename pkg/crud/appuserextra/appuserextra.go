package appuserextra

import (
	"context"
	"fmt"

	npool "github.com/NpoolPlatform/message/npool/appusermgr"

	constant "github.com/NpoolPlatform/appuser-manager/pkg/const"
	db "github.com/NpoolPlatform/appuser-manager/pkg/db"
	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent"
	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent/appuserextra"

	"github.com/google/uuid"
)

func validateAppUserExtra(info *npool.AppUserExtra) error {
	if _, err := uuid.Parse(info.GetAppID()); err != nil {
		return fmt.Errorf("invalid app id: %v", err)
	}
	if _, err := uuid.Parse(info.GetUserID()); err != nil {
		return fmt.Errorf("invalid user id: %v", err)
	}
	return nil
}

func dbRowToAppUserExtra(row *ent.AppUserExtra) *npool.AppUserExtra {
	return &npool.AppUserExtra{
		ID:            row.ID.String(),
		AppID:         row.AppID.String(),
		UserID:        row.UserID.String(),
		Username:      row.Username,
		FirstName:     row.FirstName,
		LastName:      row.LastName,
		AddressFields: row.AddressFields,
		Gender:        row.Gender,
		PostalCode:    row.PostalCode,
		Age:           row.Age,
		Birthday:      row.Birthday,
		Avatar:        row.Avatar,
		Organization:  row.Organization,
		IDNumber:      row.IDNumber,
	}
}

func Create(ctx context.Context, in *npool.CreateAppUserExtraRequest) (*npool.CreateAppUserExtraResponse, error) {
	if err := validateAppUserExtra(in.GetInfo()); err != nil {
		return nil, fmt.Errorf("invalid parameter: %v", err)
	}

	ctx, cancel := context.WithTimeout(ctx, constant.DBTimeout)
	defer cancel()

	cli, err := db.Client()
	if err != nil {
		return nil, fmt.Errorf("fail get db client: %v", err)
	}

	info, err := cli.
		AppUserExtra.
		Create().
		SetAppID(uuid.MustParse(in.GetInfo().GetAppID())).
		SetUserID(uuid.MustParse(in.GetInfo().GetUserID())).
		SetUsername(in.GetInfo().GetUsername()).
		SetFirstName(in.GetInfo().GetFirstName()).
		SetLastName(in.GetInfo().GetLastName()).
		SetAddressFields(in.GetInfo().GetAddressFields()).
		SetGender(in.GetInfo().GetGender()).
		SetPostalCode(in.GetInfo().GetPostalCode()).
		SetAge(in.GetInfo().GetAge()).
		SetBirthday(in.GetInfo().GetBirthday()).
		SetAvatar(in.GetInfo().GetAvatar()).
		SetOrganization(in.GetInfo().GetOrganization()).
		SetIDNumber(in.GetInfo().GetIDNumber()).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("fail create app user extra: %v", err)
	}

	return &npool.CreateAppUserExtraResponse{
		Info: dbRowToAppUserExtra(info),
	}, nil
}

func Update(ctx context.Context, in *npool.UpdateAppUserExtraRequest) (*npool.UpdateAppUserExtraResponse, error) {
	id, err := uuid.Parse(in.GetInfo().GetID())
	if err != nil {
		return nil, fmt.Errorf("invalid app user extra id: %v", err)
	}

	if err := validateAppUserExtra(in.GetInfo()); err != nil {
		return nil, fmt.Errorf("invalid parameter: %v", err)
	}

	ctx, cancel := context.WithTimeout(ctx, constant.DBTimeout)
	defer cancel()

	cli, err := db.Client()
	if err != nil {
		return nil, fmt.Errorf("fail get db client: %v", err)
	}

	info, err := cli.
		AppUserExtra.
		UpdateOneID(id).
		SetUsername(in.GetInfo().GetUsername()).
		SetFirstName(in.GetInfo().GetFirstName()).
		SetLastName(in.GetInfo().GetLastName()).
		SetAddressFields(in.GetInfo().GetAddressFields()).
		SetGender(in.GetInfo().GetGender()).
		SetPostalCode(in.GetInfo().GetPostalCode()).
		SetAge(in.GetInfo().GetAge()).
		SetBirthday(in.GetInfo().GetBirthday()).
		SetAvatar(in.GetInfo().GetAvatar()).
		SetOrganization(in.GetInfo().GetOrganization()).
		SetIDNumber(in.GetInfo().GetIDNumber()).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("fail update app user extra: %v", err)
	}

	return &npool.UpdateAppUserExtraResponse{
		Info: dbRowToAppUserExtra(info),
	}, nil
}

func Get(ctx context.Context, in *npool.GetAppUserExtraRequest) (*npool.GetAppUserExtraResponse, error) {
	id, err := uuid.Parse(in.GetID())
	if err != nil {
		return nil, fmt.Errorf("invalid app user extra id: %v", err)
	}

	ctx, cancel := context.WithTimeout(ctx, constant.DBTimeout)
	defer cancel()

	cli, err := db.Client()
	if err != nil {
		return nil, fmt.Errorf("fail get db client: %v", err)
	}

	infos, err := cli.
		AppUserExtra.
		Query().
		Where(
			appuserextra.And(
				appuserextra.ID(id),
			),
		).
		All(ctx)
	if err != nil {
		return nil, fmt.Errorf("fail query app user extra: %v", err)
	}

	var myAppUserExtra *npool.AppUserExtra
	for _, info := range infos {
		myAppUserExtra = dbRowToAppUserExtra(info)
		break
	}

	return &npool.GetAppUserExtraResponse{
		Info: myAppUserExtra,
	}, nil
}

func GetByAppUser(ctx context.Context, in *npool.GetAppUserExtraByAppUserRequest) (*npool.GetAppUserExtraByAppUserResponse, error) {
	appID, err := uuid.Parse(in.GetAppID())
	if err != nil {
		return nil, fmt.Errorf("invalid app id: %v", err)
	}

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
		AppUserExtra.
		Query().
		Where(
			appuserextra.And(
				appuserextra.AppID(appID),
				appuserextra.UserID(userID),
			),
		).
		All(ctx)
	if err != nil {
		return nil, fmt.Errorf("fail query app user extra: %v", err)
	}

	var appUserExtra *npool.AppUserExtra
	for _, info := range infos {
		appUserExtra = dbRowToAppUserExtra(info)
		break
	}

	return &npool.GetAppUserExtraByAppUserResponse{
		Info: appUserExtra,
	}, nil
}
