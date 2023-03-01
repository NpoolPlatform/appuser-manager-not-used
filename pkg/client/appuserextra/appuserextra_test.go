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
	"google.golang.org/grpc/credentials/insecure"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	val "github.com/NpoolPlatform/message/npool"

	testinit "github.com/NpoolPlatform/appuser-manager/pkg/testinit"
	npool "github.com/NpoolPlatform/message/npool/appuser/mgr/v2/appuserextra"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"

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

var ret = npool.AppUserExtra{
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
	ActionCredits: decimal.NewFromInt(0).String(),
}

var (
	req = npool.AppUserExtraReq{
		ID:            &ret.ID,
		AppID:         &ret.AppID,
		UserID:        &ret.UserID,
		FirstName:     &ret.FirstName,
		Birthday:      &ret.Birthday,
		LastName:      &ret.LastName,
		Gender:        &ret.Gender,
		Avatar:        &ret.Avatar,
		Username:      &ret.Username,
		PostalCode:    &ret.PostalCode,
		Age:           &ret.Age,
		Organization:  &ret.Organization,
		IDNumber:      &ret.IDNumber,
		AddressFields: ret.AddressFields,
	}
)

var info *npool.AppUserExtra

func createAppUserExtra(t *testing.T) {
	var err error
	info, err = CreateAppUserExtra(context.Background(), &req)
	if assert.Nil(t, err) {
		assert.Equal(t, info, &ret)
	}
}

func createAppUserExtras(t *testing.T) {
	rets := []npool.AppUserExtra{
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
			ActionCredits: decimal.NewFromInt(0).String(),
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
			ActionCredits: decimal.NewFromInt(0).String(),
		},
	}

	appUserExtras := []*npool.AppUserExtraReq{}
	for key := range rets {
		appUserExtras = append(appUserExtras, &npool.AppUserExtraReq{
			ID:            &rets[key].ID,
			AppID:         &rets[key].AppID,
			UserID:        &rets[key].UserID,
			FirstName:     &rets[key].FirstName,
			Birthday:      &rets[key].Birthday,
			LastName:      &rets[key].LastName,
			Gender:        &rets[key].Gender,
			Avatar:        &rets[key].Avatar,
			Username:      &rets[key].Username,
			PostalCode:    &rets[key].PostalCode,
			Age:           &rets[key].Age,
			Organization:  &rets[key].Organization,
			IDNumber:      &rets[key].IDNumber,
			AddressFields: rets[key].AddressFields,
		})
	}

	infos, err := CreateAppUserExtras(context.Background(), appUserExtras)
	if assert.Nil(t, err) {
		assert.Equal(t, len(infos), 2)
	}
}

func updateAppUserExtra(t *testing.T) {
	var err error

	credits := "123.1"

	ret.ActionCredits = credits
	req.ActionCredits = &credits

	info, err = UpdateAppUserExtra(context.Background(), &req)
	if assert.Nil(t, err) {
		assert.Equal(t, info, &ret)
	}

	ret.ActionCredits = "246.2"

	info, err = UpdateAppUserExtra(context.Background(), &req)
	if assert.Nil(t, err) {
		assert.Equal(t, info, &ret)
	}
}

func getAppUserExtra(t *testing.T) {
	var err error
	info, err = GetAppUserExtra(context.Background(), info.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, info, &ret)
	}
}

func getAppUserExtras(t *testing.T) {
	infos, total, err := GetAppUserExtras(context.Background(),
		&npool.Conds{
			ID: &val.StringVal{
				Value: info.ID,
				Op:    cruder.EQ,
			},
		}, 0, 1)
	if assert.Nil(t, err) {
		assert.Equal(t, total, uint32(1))
		assert.Equal(t, infos[0], &ret)
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
		assert.Equal(t, info, &ret)
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
		assert.Equal(t, info, &ret)
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
