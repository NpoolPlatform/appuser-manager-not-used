package banapp

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	testinit "github.com/NpoolPlatform/appuser-manager/pkg/testinit" //nolint
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

func assertBanApp(t *testing.T, actual, expected *npool.BanApp) {
	assert.Equal(t, actual.AppID, expected.AppID)
	assert.Equal(t, actual.Message, expected.Message)
}

func TestCRUD(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	banApp := npool.BanApp{
		AppID:   uuid.New().String(),
		Message: "Just want to ban you",
	}

	resp, err := Create(context.Background(), &npool.CreateBanAppRequest{
		Info: &banApp,
	})

	if assert.Nil(t, err) {
		assert.NotEqual(t, resp.Info.ID, uuid.UUID{})
		assertBanApp(t, resp.Info, &banApp)
	}

	resp1, err := Get(context.Background(), &npool.GetBanAppRequest{
		ID: resp.Info.ID,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, resp1.Info.ID, resp.Info.ID)
		assertBanApp(t, resp1.Info, &banApp)
	}

	resp2, err := GetByApp(context.Background(), &npool.GetBanAppByAppRequest{
		AppID: banApp.AppID,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, resp2.Info.ID, resp.Info.ID)
		assertBanApp(t, resp2.Info, &banApp)
	}

	resp3, err := Delete(context.Background(), &npool.DeleteBanAppRequest{
		ID: resp.Info.ID,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, resp3.Info.ID, resp.Info.ID)
		assertBanApp(t, resp3.Info, &banApp)
	}

	resp4, err := GetByApp(context.Background(), &npool.GetBanAppByAppRequest{
		AppID: banApp.AppID,
	})
	if assert.Nil(t, err) {
		assert.Nil(t, resp4.Info)
	}
}
