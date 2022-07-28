package appuserextra

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"

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

func assertAppUserExtra(t *testing.T, actual, expected *npool.AppUserExtra) {
	assert.Equal(t, actual.AppID, expected.AppID)
	assert.Equal(t, actual.UserID, expected.UserID)
	assert.Equal(t, actual.Username, expected.Username)
	assert.Equal(t, actual.FirstName, expected.FirstName)
	assert.Equal(t, actual.LastName, expected.LastName)
	assert.Equal(t, actual.AddressFields, expected.AddressFields)
	assert.Equal(t, actual.Gender, expected.Gender)
	assert.Equal(t, actual.PostalCode, expected.PostalCode)
	assert.Equal(t, actual.Age, expected.Age)
	assert.Equal(t, actual.Birthday, expected.Birthday)
	assert.Equal(t, actual.Avatar, expected.Avatar)
	assert.Equal(t, actual.Organization, expected.Organization)
}

func TestCRUD(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	appUserExtra := npool.AppUserExtra{
		AppID:         uuid.New().String(),
		UserID:        uuid.New().String(),
		Username:      uuid.New().String(),
		FirstName:     uuid.New().String(),
		LastName:      uuid.New().String(),
		AddressFields: []string{uuid.New().String(), uuid.New().String()},
		Gender:        uuid.New().String(),
		PostalCode:    uuid.New().String(),
		Age:           100,
		Birthday:      uint32(time.Now().Unix()),
		Avatar:        uuid.New().String(),
		Organization:  uuid.New().String(),
	}

	resp, err := Create(context.Background(), &npool.CreateAppUserExtraRequest{
		Info: &appUserExtra,
	})

	if assert.Nil(t, err) {
		assert.NotEqual(t, resp.Info.ID, uuid.UUID{})
		assertAppUserExtra(t, resp.Info, &appUserExtra)
	}

	resp1, err := Get(context.Background(), &npool.GetAppUserExtraRequest{
		ID: resp.Info.ID,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, resp1.Info.ID, resp.Info.ID)
		assertAppUserExtra(t, resp1.Info, &appUserExtra)
	}

	appUserExtra.ID = resp.Info.ID
	resp2, err := Update(context.Background(), &npool.UpdateAppUserExtraRequest{
		Info: &appUserExtra,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, resp2.Info.ID, resp.Info.ID)
		assertAppUserExtra(t, resp2.Info, &appUserExtra)
	}

	resp3, err := GetByAppUser(context.Background(), &npool.GetAppUserExtraByAppUserRequest{
		AppID:  appUserExtra.AppID,
		UserID: appUserExtra.UserID,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, resp3.Info.ID, resp.Info.ID)
		assertAppUserExtra(t, resp3.Info, &appUserExtra)
	}
}
