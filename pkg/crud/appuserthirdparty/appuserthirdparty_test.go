package appuserthirdparty

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

func assertAppUser(t *testing.T, actual, expected *npool.AppUserThirdParty) {
	assert.Equal(t, actual.AppID, expected.AppID)
	assert.Equal(t, actual.UserID, expected.UserID)
	assert.Equal(t, actual.ThirdPartyUserID, expected.ThirdPartyUserID)
	assert.Equal(t, actual.ThirdPartyUsername, expected.ThirdPartyUsername)
	assert.Equal(t, actual.ThirdPartyUserAvatar, expected.ThirdPartyUserAvatar)
	assert.Equal(t, actual.ThirdPartyID, expected.ThirdPartyID)
}

func TestCRUD(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	appUser := npool.AppUserThirdParty{
		AppID:                uuid.New().String(),
		UserID:               uuid.New().String(),
		ThirdPartyUserID:     uuid.New().String(),
		ThirdPartyUsername:   uuid.New().String(),
		ThirdPartyUserAvatar: uuid.New().String(),
		ThirdPartyID:         uuid.New().String(),
	}

	resp, err := Create(context.Background(), &npool.CreateAppUserThirdPartyRequest{
		Info: &appUser,
	})

	if assert.Nil(t, err) {
		assert.NotEqual(t, resp.Info.ID, uuid.UUID{})
		assertAppUser(t, resp.Info, &appUser)
	}

	resp1, err := GetByAppUserThirdParty(context.Background(), &npool.GetAppUserThirdPartyByAppThirdPartyIDRequest{
		AppID:            resp.Info.AppID,
		ThirdPartyID:     resp.Info.ThirdPartyID,
		ThirdPartyUserID: resp.Info.ThirdPartyUserID,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, resp1.Info.ID, resp.Info.ID)
		assertAppUser(t, resp1.Info, &appUser)
	}
}
