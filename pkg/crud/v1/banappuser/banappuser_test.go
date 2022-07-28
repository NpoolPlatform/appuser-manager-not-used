package banappuser

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

func assertBanAppUser(t *testing.T, actual, expected *npool.BanAppUser) {
	assert.Equal(t, actual.AppID, expected.AppID)
	assert.Equal(t, actual.UserID, expected.UserID)
	assert.Equal(t, actual.Message, expected.Message)
}

func TestCRUD(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	banAppUser := npool.BanAppUser{
		AppID:   uuid.New().String(),
		UserID:  uuid.New().String(),
		Message: "Just want to ban you",
	}

	resp, err := Create(context.Background(), &npool.CreateBanAppUserRequest{
		Info: &banAppUser,
	})

	if assert.Nil(t, err) {
		assert.NotEqual(t, resp.Info.ID, uuid.UUID{})
		assertBanAppUser(t, resp.Info, &banAppUser)
	}

	resp1, err := Get(context.Background(), &npool.GetBanAppUserRequest{
		ID: resp.Info.ID,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, resp1.Info.ID, resp.Info.ID)
		assertBanAppUser(t, resp1.Info, &banAppUser)
	}

	resp2, err := GetByAppUser(context.Background(), &npool.GetBanAppUserByAppUserRequest{
		AppID:  banAppUser.AppID,
		UserID: banAppUser.UserID,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, resp2.Info.ID, resp.Info.ID)
		assertBanAppUser(t, resp2.Info, &banAppUser)
	}

	resp3, err := Delete(context.Background(), &npool.DeleteBanAppUserRequest{
		ID: resp.Info.ID,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, resp3.Info.ID, resp.Info.ID)
		assertBanAppUser(t, resp3.Info, &banAppUser)
	}

	resp4, err := GetByAppUser(context.Background(), &npool.GetBanAppUserByAppUserRequest{
		AppID: banAppUser.AppID,
	})
	if assert.Nil(t, err) {
		assert.Nil(t, resp4.Info)
	}
}
