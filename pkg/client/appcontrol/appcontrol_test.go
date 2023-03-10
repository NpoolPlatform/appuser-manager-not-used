package appcontrol

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
	"google.golang.org/grpc/credentials/insecure"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	val "github.com/NpoolPlatform/message/npool"

	testinit "github.com/NpoolPlatform/appuser-manager/pkg/testinit"
	npool "github.com/NpoolPlatform/message/npool/appuser/mgr/v2/appcontrol"
	"github.com/google/uuid"

	"github.com/stretchr/testify/assert"

	rcpt "github.com/NpoolPlatform/message/npool/appuser/mgr/v2/recaptcha"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
)

func init() {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	if err := testinit.Init(); err != nil {
		fmt.Printf("cannot init test stub: %v\n", err)
	}
}

var appControlDate = npool.AppControl{
	ID:                       uuid.NewString(),
	AppID:                    uuid.NewString(),
	SignupMethods:            []basetypes.SignMethod{basetypes.SignMethod_Email, basetypes.SignMethod_Mobile},
	ExtSigninMethods:         []basetypes.SignMethod{basetypes.SignMethod_Google, basetypes.SignMethod_Github},
	RecaptchaMethod:          rcpt.RecaptchaType_GoogleRecaptchaV3,
	KycEnable:                false,
	SigninVerifyEnable:       false,
	InvitationCodeMust:       false,
	CreateInvitationCodeWhen: npool.CreateInvitationCodeWhen_DefaultWhen,
	MaxTypedCouponsPerOrder:  1,
	UnderMaintenance:         true,
	CommitButtons:            []string{uuid.NewString()},
}

var (
	appControlInfo = npool.AppControlReq{
		ID:                 &appControlDate.ID,
		AppID:              &appControlDate.AppID,
		SignupMethods:      appControlDate.SignupMethods,
		ExtSigninMethods:   appControlDate.ExtSigninMethods,
		RecaptchaMethod:    &appControlDate.RecaptchaMethod,
		KycEnable:          &appControlDate.KycEnable,
		SigninVerifyEnable: &appControlDate.SigninVerifyEnable,
		InvitationCodeMust: &appControlDate.InvitationCodeMust,
		UnderMaintenance:   &appControlDate.UnderMaintenance,
		CommitButtons:      appControlDate.CommitButtons,
	}
)

var info *npool.AppControl

func createAppControl(t *testing.T) {
	var err error
	info, err = CreateAppControl(context.Background(), &appControlInfo)
	if assert.Nil(t, err) {
		assert.Equal(t, info, &appControlDate)
	}
}

func createAppControls(t *testing.T) {
	appControlDates := []npool.AppControl{
		{
			ID:                       uuid.NewString(),
			AppID:                    uuid.NewString(),
			RecaptchaMethod:          rcpt.RecaptchaType_GoogleRecaptchaV3,
			SignupMethods:            []basetypes.SignMethod{basetypes.SignMethod_Email, basetypes.SignMethod_Mobile},
			ExtSigninMethods:         []basetypes.SignMethod{basetypes.SignMethod_Google, basetypes.SignMethod_Github},
			KycEnable:                false,
			SigninVerifyEnable:       false,
			InvitationCodeMust:       false,
			CreateInvitationCodeWhen: npool.CreateInvitationCodeWhen_DefaultWhen,
			MaxTypedCouponsPerOrder:  1,
			UnderMaintenance:         true,
			CommitButtons:            []string{uuid.NewString()},
		},
		{
			ID:                       uuid.NewString(),
			AppID:                    uuid.NewString(),
			RecaptchaMethod:          rcpt.RecaptchaType_GoogleRecaptchaV3,
			SignupMethods:            []basetypes.SignMethod{basetypes.SignMethod_Email, basetypes.SignMethod_Mobile},
			ExtSigninMethods:         []basetypes.SignMethod{basetypes.SignMethod_Google, basetypes.SignMethod_Github},
			KycEnable:                false,
			SigninVerifyEnable:       false,
			InvitationCodeMust:       false,
			CreateInvitationCodeWhen: npool.CreateInvitationCodeWhen_DefaultWhen,
			MaxTypedCouponsPerOrder:  1,
			UnderMaintenance:         true,
			CommitButtons:            []string{uuid.NewString()},
		},
	}

	appControls := []*npool.AppControlReq{}
	for key := range appControlDates {
		appControls = append(appControls, &npool.AppControlReq{
			ID:                 &appControlDates[key].ID,
			AppID:              &appControlDates[key].AppID,
			SignupMethods:      appControlDates[key].SignupMethods,
			ExtSigninMethods:   appControlDates[key].ExtSigninMethods,
			RecaptchaMethod:    &appControlDates[key].RecaptchaMethod,
			KycEnable:          &appControlDates[key].KycEnable,
			SigninVerifyEnable: &appControlDates[key].SigninVerifyEnable,
			InvitationCodeMust: &appControlDates[key].InvitationCodeMust,
			UnderMaintenance:   &appControlDates[key].UnderMaintenance,
			CommitButtons:      appControlDates[key].CommitButtons,
		})
	}

	infos, err := CreateAppControls(context.Background(), appControls)
	if assert.Nil(t, err) {
		assert.Equal(t, len(infos), 2)
	}
}

func updateAppControl(t *testing.T) {
	var err error
	info, err = UpdateAppControl(context.Background(), &appControlInfo)
	if assert.Nil(t, err) {
		assert.Equal(t, info, &appControlDate)
	}
}

func getAppControl(t *testing.T) {
	var err error
	info, err = GetAppControl(context.Background(), info.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, info, &appControlDate)
	}
}

func getAppControls(t *testing.T) {
	infos, total, err := GetAppControls(context.Background(),
		&npool.Conds{
			ID: &val.StringVal{
				Value: info.ID,
				Op:    cruder.EQ,
			},
		}, 0, 1)
	if assert.Nil(t, err) {
		assert.Equal(t, total, uint32(1))
		assert.Equal(t, infos[0], &appControlDate)
	}
}

func getAppControlOnly(t *testing.T) {
	var err error
	info, err = GetAppControlOnly(context.Background(),
		&npool.Conds{
			ID: &val.StringVal{
				Value: info.ID,
				Op:    cruder.EQ,
			},
		})
	if assert.Nil(t, err) {
		assert.Equal(t, info, &appControlDate)
	}
}

func countAppControls(t *testing.T) {
	count, err := CountAppControls(context.Background(),
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

func existAppControl(t *testing.T) {
	exist, err := ExistAppControl(context.Background(), info.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, exist, true)
	}
}

func existAppControlConds(t *testing.T) {
	exist, err := ExistAppControlConds(context.Background(),
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

func deleteAppControl(t *testing.T) {
	info, err := DeleteAppControl(context.Background(), info.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, info, &appControlDate)
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

	t.Run("createAppControl", createAppControl)
	t.Run("createAppControls", createAppControls)
	t.Run("getAppControl", getAppControl)
	t.Run("getAppControls", getAppControls)
	t.Run("getAppControlOnly", getAppControlOnly)
	t.Run("updateAppControl", updateAppControl)
	t.Run("existAppControl", existAppControl)
	t.Run("existAppControlConds", existAppControlConds)
	t.Run("count", countAppControls)
	t.Run("delete", deleteAppControl)
}
