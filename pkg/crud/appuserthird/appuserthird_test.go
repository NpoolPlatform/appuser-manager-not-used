package appuserthird

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	constant "github.com/NpoolPlatform/appuser-manager/pkg/const"
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

func assertAppUser(t *testing.T, actual, expected *npool.AppUserThird) {
	assert.Equal(t, actual.AppID, expected.AppID)
	assert.Equal(t, actual.UserID, expected.UserID)
	assert.Equal(t, actual.ThirdUserId, expected.ThirdUserId)
	assert.Equal(t, actual.Third, constant.ThirdGithub)
	assert.Equal(t, actual.ThirdUserName, expected.ThirdUserName)
	assert.Equal(t, actual.ThirdUserPicture, expected.ThirdUserPicture)
	assert.Equal(t, actual.ThirdExtra, expected.ThirdExtra)
	assert.Equal(t, actual.ThirdId, expected.ThirdId)
}

func TestCRUD(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	appUser := npool.AppUserThird{
		AppID:            uuid.New().String(),
		UserID:           uuid.New().String(),
		ThirdUserId:      uuid.New().String(),
		Third:            constant.ThirdGithub,
		ThirdUserName:    uuid.New().String(),
		ThirdUserPicture: uuid.New().String(),
		ThirdExtra:       uuid.New().String(),
		ThirdId:          uuid.New().String(),
	}

	resp, err := Create(context.Background(), &npool.CreateAppUserThirdRequest{
		Info: &appUser,
	})

	if assert.Nil(t, err) {
		assert.NotEqual(t, resp.Info.ID, uuid.UUID{})
		assertAppUser(t, resp.Info, &appUser)
	}

	resp1, err := GetByAppUserThird(context.Background(), &npool.GetAppUserThirdByAppThirdRequest{
		AppID:       resp.Info.AppID,
		ThirdID:     resp.Info.ThirdId,
		ThirdUserID: resp.Info.ThirdUserId,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, resp1.Info.ID, resp.Info.ID)
		assertAppUser(t, resp1.Info, &appUser)
	}
}
