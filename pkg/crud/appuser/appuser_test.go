package appuser

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	testinit "github.com/NpoolPlatform/appuser-manager/pkg/test-init" //nolint
	npool "github.com/NpoolPlatform/message/npool/appusermgr"

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

func assertAppUser(t *testing.T, actual, expected *npool.AppUser) {
	assert.Equal(t, actual.AppID, expected.AppID)
	assert.Equal(t, actual.EmailAddress, expected.EmailAddress)
	assert.Equal(t, actual.PhoneNO, expected.PhoneNO)
	assert.Equal(t, actual.ImportFromApp, expected.ImportFromApp)
}

func TestCRUD(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	appUser := npool.AppUser{
		AppID:         uuid.New().String(),
		EmailAddress:  uuid.New().String(),
		PhoneNO:       uuid.New().String(),
		ImportFromApp: uuid.New().String(),
	}

	resp, err := Create(context.Background(), &npool.CreateAppUserRequest{
		Info: &appUser,
	})

	if assert.Nil(t, err) {
		assert.NotEqual(t, resp.Info.ID, uuid.UUID{})
		assertAppUser(t, resp.Info, &appUser)
	}

	resp1, err := Get(context.Background(), &npool.GetAppUserRequest{
		ID: resp.Info.ID,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, resp1.Info.ID, resp.Info.ID)
		assertAppUser(t, resp1.Info, &appUser)
	}

	appUser.ID = resp.Info.ID
	resp2, err := Update(context.Background(), &npool.UpdateAppUserRequest{
		Info: &appUser,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, resp2.Info.ID, resp.Info.ID)
		assertAppUser(t, resp2.Info, &appUser)
	}

	resp3, err := GetByApp(context.Background(), &npool.GetAppUsersByAppRequest{
		AppID: appUser.AppID,
	})
	if assert.Nil(t, err) {
		assert.NotEqual(t, len(resp3.Infos), 0)
	}
}
