package appuserextra

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/go-service-framework/pkg/config"

	"bou.ke/monkey"
	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"
	"google.golang.org/grpc"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	val "github.com/NpoolPlatform/message/npool"

	testinit "github.com/NpoolPlatform/appuser-manager/pkg/testinit"
	npool "github.com/NpoolPlatform/message/npool/appuser/mgr/v2/appuserextra"
	"github.com/google/uuid"

	"github.com/stretchr/testify/assert"
)

func init() {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	if err := testinit.Init(); err != nil {
		fmt.Printf("cannot init test stub: %v\n", err)
	}
}

var appUserExtraDate = npool.AppUserExtra{
	ID:            uuid.NewString(),
	AppID:         uuid.NewString(),
	UserID:        uuid.NewString(),
	FirstName:     uuid.NewString(),
	Birthday:      0,
	LastName:      uuid.NewString(),
	Gender:        uuid.NewString(),
	Avatar:        uuid.NewString(),
	Username:      uuid.NewString(),
	PostalCode:    uuid.NewString(),
	Age:           0,
	Organization:  uuid.NewString(),
	IDNumber:      uuid.NewString(),
	AddressFields: []string{uuid.NewString()},
}

var (
	appUserExtraInfo = npool.AppUserExtraReq{
		ID:            &appUserExtraDate.ID,
		AppID:         &appUserExtraDate.AppID,
		UserID:        &appUserExtraDate.UserID,
		FirstName:     &appUserExtraDate.FirstName,
		Birthday:      &appUserExtraDate.Birthday,
		LastName:      &appUserExtraDate.LastName,
		Gender:        &appUserExtraDate.Gender,
		Avatar:        &appUserExtraDate.Avatar,
		Username:      &appUserExtraDate.Username,
		PostalCode:    &appUserExtraDate.PostalCode,
		Age:           &appUserExtraDate.Age,
		Organization:  &appUserExtraDate.Organization,
		IDNumber:      &appUserExtraDate.IDNumber,
		AddressFields: appUserExtraDate.AddressFields,
	}
)

var info *npool.AppUserExtra

func createAppUserExtra(t *testing.T) {
	var err error
	info, err = CreateAppUserExtra(context.Background(), &appUserExtraInfo)
	if assert.Nil(t, err) {
		assert.Equal(t, info, &appUserExtraDate)
	}
}

func createAppUserExtras(t *testing.T) {
	appUserExtraDates := []npool.AppUserExtra{
		{
			ID:            uuid.NewString(),
			AppID:         uuid.NewString(),
			UserID:        uuid.NewString(),
			FirstName:     uuid.NewString(),
			Birthday:      0,
			LastName:      uuid.NewString(),
			Gender:        uuid.NewString(),
			Avatar:        uuid.NewString(),
			Username:      uuid.NewString(),
			PostalCode:    uuid.NewString(),
			Age:           0,
			Organization:  uuid.NewString(),
			IDNumber:      uuid.NewString(),
			AddressFields: []string{uuid.NewString()},
		},
		{
			ID:            uuid.NewString(),
			AppID:         uuid.NewString(),
			UserID:        uuid.NewString(),
			FirstName:     uuid.NewString(),
			Birthday:      0,
			LastName:      uuid.NewString(),
			Gender:        uuid.NewString(),
			Avatar:        uuid.NewString(),
			Username:      uuid.NewString(),
			PostalCode:    uuid.NewString(),
			Age:           0,
			Organization:  uuid.NewString(),
			IDNumber:      uuid.NewString(),
			AddressFields: []string{uuid.NewString()},
		},
	}

	appUserExtras := []*npool.AppUserExtraReq{}
	for key := range appUserExtraDates {
		appUserExtras = append(appUserExtras, &npool.AppUserExtraReq{
			ID:            &appUserExtraDates[key].ID,
			AppID:         &appUserExtraDates[key].AppID,
			UserID:        &appUserExtraDates[key].UserID,
			FirstName:     &appUserExtraDates[key].FirstName,
			Birthday:      &appUserExtraDates[key].Birthday,
			LastName:      &appUserExtraDates[key].LastName,
			Gender:        &appUserExtraDates[key].Gender,
			Avatar:        &appUserExtraDates[key].Avatar,
			Username:      &appUserExtraDates[key].Username,
			PostalCode:    &appUserExtraDates[key].PostalCode,
			Age:           &appUserExtraDates[key].Age,
			Organization:  &appUserExtraDates[key].Organization,
			IDNumber:      &appUserExtraDates[key].IDNumber,
			AddressFields: appUserExtraDates[key].AddressFields,
		})
	}

	infos, err := CreateAppUserExtras(context.Background(), appUserExtras)
	if assert.Nil(t, err) {
		assert.Equal(t, len(infos), 2)
	}
}

func updateAppUserExtra(t *testing.T) {
	var err error
	info, err = UpdateAppUserExtra(context.Background(), &appUserExtraInfo)
	if assert.Nil(t, err) {
		assert.Equal(t, info, &appUserExtraDate)
	}
}

func getAppUserExtra(t *testing.T) {
	var err error
	info, err = GetAppUserExtra(context.Background(), info.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, info, &appUserExtraDate)
	}
}

func getAppUserExtras(t *testing.T) {
	infos, total, err := GetAppUserExtras(context.Background(),
		&npool.Conds{
			ID: &val.StringVal{
				Value: info.ID,
				Op:    cruder.EQ,
			},
		}, 1, 0)
	if assert.Nil(t, err) {
		assert.Equal(t, total, uint32(1))
		assert.Equal(t, infos[0], &appUserExtraDate)
	}
}

func getAppUserExtraOnly(t *testing.T) {
	var err error
	info, err = GetAppUserExtraOnly(context.Background(),
		&npool.Conds{
			ID: &val.StringVal{
				Value: info.ID,
				Op:    cruder.EQ,
			},
		})
	if assert.Nil(t, err) {
		assert.Equal(t, info, &appUserExtraDate)
	}
}

func countAppUserExtras(t *testing.T) {
	count, err := CountAppUserExtras(context.Background(),
		&npool.Conds{
			ID: &val.StringVal{
				Value: info.ID,
				Op:    cruder.EQ,
			},
		},
	)
	if assert.Nil(t, err) {
		assert.Equal(t, count, uint32(1))
	}
}

func existAppUserExtra(t *testing.T) {
	exist, err := ExistAppUserExtra(context.Background(), info.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, exist, true)
	}
}

func existAppUserExtraConds(t *testing.T) {
	exist, err := ExistAppUserExtraConds(context.Background(),
		&npool.Conds{
			ID: &val.StringVal{
				Value: info.ID,
				Op:    cruder.EQ,
			},
		},
	)
	if assert.Nil(t, err) {
		assert.Equal(t, exist, true)
	}
}

func deleteAppUserExtra(t *testing.T) {
	info, err := DeleteAppUserExtra(context.Background(), info.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, info, &appUserExtraDate)
	}
}

func TestMainOrder(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	gport := config.GetIntValueWithNameSpace("", config.KeyGRPCPort)

	monkey.Patch(grpc2.GetGRPCConn, func(service string, tags ...string) (*grpc.ClientConn, error) {
		return grpc.Dial(fmt.Sprintf("localhost:%v", gport), grpc.WithTransportCredentials(insecure.NewCredentials()))
	})

	t.Run("createAppUserExtra", createAppUserExtra)
	t.Run("createAppUserExtras", createAppUserExtras)
	t.Run("getAppUserExtra", getAppUserExtra)
	t.Run("getAppUserExtras", getAppUserExtras)
	t.Run("getAppUserExtraOnly", getAppUserExtraOnly)
	t.Run("updateAppUserExtra", updateAppUserExtra)
	t.Run("existAppUserExtra", existAppUserExtra)
	t.Run("existAppUserExtraConds", existAppUserExtraConds)
	t.Run("count", countAppUserExtras)
	t.Run("delete", deleteAppUserExtra)
}
