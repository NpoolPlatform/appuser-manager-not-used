package appusersecret

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/appuser-manager/pkg/test-init" //nolint
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

func assertAppUserSecret(t *testing.T, actual, expected *npool.AppUserSecret) {
	assert.Equal(t, actual.AppID, expected.AppID)
	assert.Equal(t, actual.UserID, expected.UserID)
	assert.Equal(t, actual.PasswordHash, expected.PasswordHash)
	assert.Equal(t, actual.Salt, expected.Salt)
	assert.Equal(t, actual.GoogleSecret, expected.GoogleSecret)
}

func TestCRUD(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	appUserSecret := npool.AppUserSecret{
		AppID:        uuid.New().String(),
		UserID:       uuid.New().String(),
		PasswordHash: uuid.New().String(),
		Salt:         uuid.New().String(),
		GoogleSecret: uuid.New().String(),
	}

	resp, err := Create(context.Background(), &npool.CreateAppUserSecretRequest{
		Info: &appUserSecret,
	})

	if assert.Nil(t, err) {
		assert.NotEqual(t, resp.Info.ID, uuid.UUID{})
		assertAppUserSecret(t, resp.Info, &appUserSecret)
	}

	resp1, err := Get(context.Background(), &npool.GetAppUserSecretRequest{
		ID: resp.Info.ID,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, resp1.Info.ID, resp.Info.ID)
		assertAppUserSecret(t, resp1.Info, &appUserSecret)
	}

	appUserSecret.ID = resp.Info.ID
	resp2, err := Update(context.Background(), &npool.UpdateAppUserSecretRequest{
		Info: &appUserSecret,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, resp2.Info.ID, resp.Info.ID)
		assertAppUserSecret(t, resp2.Info, &appUserSecret)
	}

	resp3, err := GetByAppUser(context.Background(), &npool.GetAppUserSecretByAppUserRequest{
		AppID:  appUserSecret.AppID,
		UserID: appUserSecret.UserID,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, resp3.Info.ID, resp.Info.ID)
		assertAppUserSecret(t, resp3.Info, &appUserSecret)
	}
}
