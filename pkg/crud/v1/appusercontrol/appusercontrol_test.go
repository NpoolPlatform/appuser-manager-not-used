package appusercontrol

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	testinit "github.com/NpoolPlatform/appuser-manager/pkg/testinit" //nolint
	npool "github.com/NpoolPlatform/message/npool/appuser/mgr/v1"

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

func assertAppUserControl(t *testing.T, actual, expected *npool.AppUserControl) {
	assert.Equal(t, actual.AppID, expected.AppID)
	assert.Equal(t, actual.UserID, expected.UserID)
}

func TestCRUD(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	appUserControl := npool.AppUserControl{
		AppID:  uuid.New().String(),
		UserID: uuid.New().String(),
	}

	resp, err := Create(context.Background(), &npool.CreateAppUserControlRequest{
		Info: &appUserControl,
	})

	if assert.Nil(t, err) {
		assert.NotEqual(t, resp.Info.ID, uuid.UUID{})
		assertAppUserControl(t, resp.Info, &appUserControl)
	}

	resp1, err := Get(context.Background(), &npool.GetAppUserControlRequest{
		ID: resp.Info.ID,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, resp1.Info.ID, resp.Info.ID)
		assertAppUserControl(t, resp1.Info, &appUserControl)
	}

	appUserControl.ID = resp.Info.ID
	resp2, err := Update(context.Background(), &npool.UpdateAppUserControlRequest{
		Info: &appUserControl,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, resp2.Info.ID, resp.Info.ID)
		assertAppUserControl(t, resp2.Info, &appUserControl)
	}

	resp3, err := GetByAppUser(context.Background(), &npool.GetAppUserControlByAppUserRequest{
		AppID:  appUserControl.AppID,
		UserID: appUserControl.UserID,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, resp3.Info.ID, resp.Info.ID)
		assertAppUserControl(t, resp3.Info, &appUserControl)
	}
}
