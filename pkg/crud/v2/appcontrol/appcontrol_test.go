package appcontrol

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	val "github.com/NpoolPlatform/message/npool"

	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

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

var entAppControl = ent.AppControl{
	ID:                       uuid.New(),
	AppID:                    uuid.New(),
	SignupMethods:            []string{basetypes.SignMethod_Email.String(), basetypes.SignMethod_Mobile.String()},
	ExternSigninMethods:      []string{basetypes.SignMethod_Github.String(), basetypes.SignMethod_Google.String()},
	RecaptchaMethod:          rcpt.RecaptchaType_GoogleRecaptchaV3.String(),
	KycEnable:                false,
	SigninVerifyEnable:       false,
	InvitationCodeMust:       false,
	CreateInvitationCodeWhen: npool.CreateInvitationCodeWhen_DefaultWhen.String(),
	MaxTypedCouponsPerOrder:  1,
	Maintaining:              true,
	CommitButtonTargets:      []string{uuid.NewString()},
}

var (
	id               = entAppControl.ID.String()
	appID            = entAppControl.AppID.String()
	recaptcha        = rcpt.RecaptchaType(rcpt.RecaptchaType_value[entAppControl.RecaptchaMethod])
	signupMethods    = []basetypes.SignMethod{basetypes.SignMethod_Email, basetypes.SignMethod_Mobile}
	extSigninMethods = []basetypes.SignMethod{basetypes.SignMethod_Github, basetypes.SignMethod_Google}

	appcontrolInfo = npool.AppControlReq{
		ID:                  &id,
		AppID:               &appID,
		SignupMethods:       signupMethods,
		ExtSigninMethods:    extSigninMethods,
		RecaptchaMethod:     &recaptcha,
		KycEnable:           &entAppControl.KycEnable,
		SigninVerifyEnable:  &entAppControl.SigninVerifyEnable,
		InvitationCodeMust:  &entAppControl.InvitationCodeMust,
		Maintaining:         &entAppControl.Maintaining,
		CommitButtonTargets: entAppControl.CommitButtonTargets,
	}
)

var info *ent.AppControl

func rowToObject(row *ent.AppControl) *ent.AppControl {
	return &ent.AppControl{
		ID:                       row.ID,
		AppID:                    row.AppID,
		SignupMethods:            row.SignupMethods,
		ExternSigninMethods:      row.ExternSigninMethods,
		RecaptchaMethod:          row.RecaptchaMethod,
		KycEnable:                row.KycEnable,
		SigninVerifyEnable:       row.SigninVerifyEnable,
		InvitationCodeMust:       row.InvitationCodeMust,
		CreateInvitationCodeWhen: row.CreateInvitationCodeWhen,
		MaxTypedCouponsPerOrder:  row.MaxTypedCouponsPerOrder,
		Maintaining:              row.Maintaining,
		CommitButtonTargets:      row.CommitButtonTargets,
	}
}

func create(t *testing.T) {
	var err error
	info, err = Create(context.Background(), &appcontrolInfo)
	if assert.Nil(t, err) {
		if assert.NotEqual(t, info.ID, uuid.UUID{}.String()) {
			entAppControl.ID = info.ID
		}
		assert.Equal(t, rowToObject(info), &entAppControl)
	}
}

func createBulk(t *testing.T) {
	entAppControl := []ent.AppControl{
		{
			ID:                       uuid.New(),
			AppID:                    uuid.New(),
			RecaptchaMethod:          rcpt.RecaptchaType_GoogleRecaptchaV3.String(),
			SignupMethods:            []string{basetypes.SignMethod_Email.String(), basetypes.SignMethod_Mobile.String()},
			ExternSigninMethods:      []string{basetypes.SignMethod_Github.String(), basetypes.SignMethod_Google.String()},
			KycEnable:                false,
			SigninVerifyEnable:       false,
			InvitationCodeMust:       false,
			CreateInvitationCodeWhen: npool.CreateInvitationCodeWhen_DefaultWhen.String(),
			MaxTypedCouponsPerOrder:  1,
			Maintaining:              true,
			CommitButtonTargets:      []string{uuid.NewString()},
		},
		{
			ID:                       uuid.New(),
			AppID:                    uuid.New(),
			RecaptchaMethod:          rcpt.RecaptchaType_GoogleRecaptchaV3.String(),
			SignupMethods:            []string{basetypes.SignMethod_Email.String(), basetypes.SignMethod_Mobile.String()},
			ExternSigninMethods:      []string{basetypes.SignMethod_Github.String(), basetypes.SignMethod_Google.String()},
			KycEnable:                false,
			SigninVerifyEnable:       false,
			InvitationCodeMust:       false,
			CreateInvitationCodeWhen: npool.CreateInvitationCodeWhen_DefaultWhen.String(),
			MaxTypedCouponsPerOrder:  1,
			Maintaining:              true,
			CommitButtonTargets:      []string{uuid.NewString()},
		},
	}
	appcontrols := []*npool.AppControlReq{}
	for key := range entAppControl {
		id := entAppControl[key].ID.String()
		appID := entAppControl[key].AppID.String()
		recaptcha := rcpt.RecaptchaType(rcpt.RecaptchaType_value[entAppControl[key].RecaptchaMethod])

		appcontrols = append(appcontrols, &npool.AppControlReq{
			ID:                  &id,
			AppID:               &appID,
			SignupMethods:       []basetypes.SignMethod{basetypes.SignMethod_Email, basetypes.SignMethod_Mobile},
			ExtSigninMethods:    []basetypes.SignMethod{basetypes.SignMethod_Github, basetypes.SignMethod_Google},
			RecaptchaMethod:     &recaptcha,
			KycEnable:           &entAppControl[key].KycEnable,
			SigninVerifyEnable:  &entAppControl[key].SigninVerifyEnable,
			InvitationCodeMust:  &entAppControl[key].InvitationCodeMust,
			Maintaining:         &entAppControl[key].Maintaining,
			CommitButtonTargets: entAppControl[key].CommitButtonTargets,
		})
	}
	infos, err := CreateBulk(context.Background(), appcontrols)
	if assert.Nil(t, err) {
		assert.Equal(t, len(infos), 2)
		assert.NotEqual(t, infos[0].ID, uuid.UUID{}.String())
		assert.NotEqual(t, infos[1].ID, uuid.UUID{}.String())
	}
}

func update(t *testing.T) {
	var err error
	info, err = Update(context.Background(), &appcontrolInfo)
	if assert.Nil(t, err) {
		assert.Equal(t, rowToObject(info), &entAppControl)
	}
}

func row(t *testing.T) {
	var err error
	info, err = Row(context.Background(), info.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, rowToObject(info), &entAppControl)
	}
}

func rows(t *testing.T) {
	infos, total, err := Rows(context.Background(),
		&npool.Conds{
			ID: &val.StringVal{
				Value: info.ID.String(),
				Op:    cruder.EQ,
			},
		}, 0, 0)
	if assert.Nil(t, err) {
		assert.Equal(t, total, 1)
		assert.Equal(t, rowToObject(infos[0]), &entAppControl)
	}
}

func rowOnly(t *testing.T) {
	var err error
	info, err = RowOnly(context.Background(),
		&npool.Conds{
			ID: &val.StringVal{
				Value: info.ID.String(),
				Op:    cruder.EQ,
			},
		})
	if assert.Nil(t, err) {
		assert.Equal(t, rowToObject(info), &entAppControl)
	}
}

func count(t *testing.T) {
	count, err := Count(context.Background(),
		&npool.Conds{
			ID: &val.StringVal{
				Value: info.ID.String(),
				Op:    cruder.EQ,
			},
		},
	)
	if assert.Nil(t, err) {
		assert.Equal(t, count, count)
	}
}

func exist(t *testing.T) {
	exist, err := Exist(context.Background(), info.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, exist, true)
	}
}

func existConds(t *testing.T) {
	exist, err := ExistConds(context.Background(),
		&npool.Conds{
			ID: &val.StringVal{
				Value: info.ID.String(),
				Op:    cruder.EQ,
			},
		},
	)
	if assert.Nil(t, err) {
		assert.Equal(t, exist, true)
	}
}

func deleteT(t *testing.T) {
	info, err := Delete(context.Background(), info.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, rowToObject(info), &entAppControl)
	}
}

func TestMainOrder(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	t.Run("create", create)
	t.Run("createBulk", createBulk)
	t.Run("row", row)
	t.Run("rows", rows)
	t.Run("rowOnly", rowOnly)
	t.Run("update", update)
	t.Run("exist", exist)
	t.Run("existConds", existConds)
	t.Run("delete", deleteT)
	t.Run("count", count)
}
