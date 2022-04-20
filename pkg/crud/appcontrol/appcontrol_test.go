package appcontrol

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

func assertAppControl(t *testing.T, actual, expected *npool.AppControl) {
	assert.Equal(t, actual.AppID, expected.AppID)
	assert.Equal(t, actual.SignupMethods, expected.SignupMethods)
	assert.Equal(t, actual.ExternSigninMethods, expected.ExternSigninMethods)
	assert.Equal(t, actual.RecaptchaMethod, expected.RecaptchaMethod)
	assert.Equal(t, actual.KycEnable, expected.KycEnable)
	assert.Equal(t, actual.SigninVerifyEnable, expected.SigninVerifyEnable)
}

func TestCRUD(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	appControl := npool.AppControl{
		AppID:               uuid.New().String(),
		SignupMethods:       []string{"mobile", "email"},
		ExternSigninMethods: []string{"facebook", "wechat"},
		RecaptchaMethod:     "google-recaptcha-v3",
		KycEnable:           false,
		SigninVerifyEnable:  false,
	}

	resp, err := Create(context.Background(), &npool.CreateAppControlRequest{
		Info: &appControl,
	})

	if assert.Nil(t, err) {
		assert.NotEqual(t, resp.Info.ID, uuid.UUID{})
		assertAppControl(t, resp.Info, &appControl)
	}

	resp1, err := Get(context.Background(), &npool.GetAppControlRequest{
		ID: resp.Info.ID,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, resp1.Info.ID, resp.Info.ID)
		assertAppControl(t, resp1.Info, &appControl)
	}

	appControl.ID = resp.Info.ID
	appControl.KycEnable = true
	resp2, err := Update(context.Background(), &npool.UpdateAppControlRequest{
		Info: &appControl,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, resp2.Info.ID, resp.Info.ID)
		assertAppControl(t, resp2.Info, &appControl)
	}

	resp3, err := GetByApp(context.Background(), &npool.GetAppControlByAppRequest{
		AppID: appControl.AppID,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, resp3.Info.ID, resp.Info.ID)
		assertAppControl(t, resp3.Info, &appControl)
	}
}
