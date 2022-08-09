package appuserthirdparty

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
	npool "github.com/NpoolPlatform/message/npool/appuser/mgr/v2/appuserthirdparty"
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

var appUserThirdPartyDate = npool.AppUserThirdParty{
	ID:                   uuid.NewString(),
	AppID:                uuid.NewString(),
	UserID:               uuid.NewString(),
	ThirdPartyID:         uuid.NewString(),
	ThirdPartyUserID:     uuid.NewString(),
	ThirdPartyUsername:   uuid.NewString(),
	ThirdPartyUserAvatar: uuid.NewString(),
}

var (
	appUserThirdPartyInfo = npool.AppUserThirdPartyReq{
		ID:                   &appUserThirdPartyDate.ID,
		AppID:                &appUserThirdPartyDate.AppID,
		UserID:               &appUserThirdPartyDate.UserID,
		ThirdPartyID:         &appUserThirdPartyDate.ThirdPartyID,
		ThirdPartyUserID:     &appUserThirdPartyDate.ThirdPartyUserID,
		ThirdPartyUsername:   &appUserThirdPartyDate.ThirdPartyUsername,
		ThirdPartyUserAvatar: &appUserThirdPartyDate.ThirdPartyUserAvatar,
	}
)

var info *npool.AppUserThirdParty

func createAppUserThirdParty(t *testing.T) {
	var err error
	info, err = CreateAppUserThirdParty(context.Background(), &appUserThirdPartyInfo)
	if assert.Nil(t, err) {
		assert.Equal(t, info, &appUserThirdPartyDate)
	}
}

func createAppUserThirdParties(t *testing.T) {
	appUserThirdPartyDates := []npool.AppUserThirdParty{
		{
			ID:                   uuid.NewString(),
			AppID:                uuid.NewString(),
			UserID:               uuid.NewString(),
			ThirdPartyID:         uuid.NewString(),
			ThirdPartyUserID:     uuid.NewString(),
			ThirdPartyUsername:   uuid.NewString(),
			ThirdPartyUserAvatar: uuid.NewString(),
		},
		{
			ID:                   uuid.NewString(),
			AppID:                uuid.NewString(),
			UserID:               uuid.NewString(),
			ThirdPartyID:         uuid.NewString(),
			ThirdPartyUserID:     uuid.NewString(),
			ThirdPartyUsername:   uuid.NewString(),
			ThirdPartyUserAvatar: uuid.NewString(),
		},
	}

	appUserThirdParties := []*npool.AppUserThirdPartyReq{}
	for key := range appUserThirdPartyDates {
		appUserThirdParties = append(appUserThirdParties, &npool.AppUserThirdPartyReq{
			ID:                   &appUserThirdPartyDates[key].ID,
			AppID:                &appUserThirdPartyDates[key].AppID,
			UserID:               &appUserThirdPartyDates[key].UserID,
			ThirdPartyID:         &appUserThirdPartyDates[key].ThirdPartyID,
			ThirdPartyUserID:     &appUserThirdPartyDates[key].ThirdPartyUserID,
			ThirdPartyUsername:   &appUserThirdPartyDates[key].ThirdPartyUsername,
			ThirdPartyUserAvatar: &appUserThirdPartyDates[key].ThirdPartyUserAvatar,
		})
	}

	infos, err := CreateAppUserThirdParties(context.Background(), appUserThirdParties)
	if assert.Nil(t, err) {
		assert.Equal(t, len(infos), 2)
	}
}

func updateAppUserThirdParty(t *testing.T) {
	var err error
	info, err = UpdateAppUserThirdParty(context.Background(), &appUserThirdPartyInfo)
	if assert.Nil(t, err) {
		assert.Equal(t, info, &appUserThirdPartyDate)
	}
}

func getAppUserThirdParty(t *testing.T) {
	var err error
	info, err = GetAppUserThirdParty(context.Background(), info.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, info, &appUserThirdPartyDate)
	}
}

func getAppUserThirdParties(t *testing.T) {
	infos, total, err := GetAppUserThirdParties(context.Background(),
		&npool.Conds{
			ID: &val.StringVal{
				Value: info.ID,
				Op:    cruder.EQ,
			},
		}, 1, 0)
	if assert.Nil(t, err) {
		assert.Equal(t, total, uint32(1))
		assert.Equal(t, infos[0], &appUserThirdPartyDate)
	}
}

func getAppUserThirdPartyOnly(t *testing.T) {
	var err error
	info, err = GetAppUserThirdPartyOnly(context.Background(),
		&npool.Conds{
			ID: &val.StringVal{
				Value: info.ID,
				Op:    cruder.EQ,
			},
		})
	if assert.Nil(t, err) {
		assert.Equal(t, info, &appUserThirdPartyDate)
	}
}

func countAppUserThirdParties(t *testing.T) {
	count, err := CountAppUserThirdParties(context.Background(),
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

func existAppUserThirdParty(t *testing.T) {
	exist, err := ExistAppUserThirdParty(context.Background(), info.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, exist, true)
	}
}

func existAppUserThirdPartyConds(t *testing.T) {
	exist, err := ExistAppUserThirdPartyConds(context.Background(),
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

func deleteAppUserThirdParty(t *testing.T) {
	info, err := DeleteAppUserThirdParty(context.Background(), info.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, info, &appUserThirdPartyDate)
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

	t.Run("createAppUserThirdParty", createAppUserThirdParty)
	t.Run("createAppUserThirdParties", createAppUserThirdParties)
	t.Run("getAppUserThirdParty", getAppUserThirdParty)
	t.Run("getAppUserThirdParties", getAppUserThirdParties)
	t.Run("getAppUserThirdPartyOnly", getAppUserThirdPartyOnly)
	t.Run("updateAppUserThirdParty", updateAppUserThirdParty)
	t.Run("existAppUserThirdParty", existAppUserThirdParty)
	t.Run("existAppUserThirdPartyConds", existAppUserThirdPartyConds)
	t.Run("count", countAppUserThirdParties)
	t.Run("delete", deleteAppUserThirdParty)
}
