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

	testinit "github.com/NpoolPlatform/appuser-manager/pkg/testinit" //nolint
	npool "github.com/NpoolPlatform/message/npool/appuser/mgr/v2/appcontrol"

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

var entAppControl = ent.AppControl{
	ID:                  uuid.New(),
	AppID:               uuid.New(),
	RecaptchaMethod:     uuid.New().String(),
	SignupMethods:       []string{uuid.New().String()},
	ExternSigninMethods: []string{uuid.New().String()},
	KycEnable:           false,
	SigninVerifyEnable:  false,
	InvitationCodeMust:  false,
}

var (
	id    = entAppControl.ID.String()
	appID = entAppControl.AppID.String()

	appcontrolInfo = npool.AppControlReq{
		ID:                  &id,
		AppID:               &appID,
		RecaptchaMethod:     &entAppControl.RecaptchaMethod,
		KycEnable:           &entAppControl.KycEnable,
		SignupMethods:       entAppControl.SignupMethods,
		ExternSigninMethods: entAppControl.ExternSigninMethods,
		SigninVerifyEnable:  &entAppControl.SigninVerifyEnable,
		InvitationCodeMust:  &entAppControl.InvitationCodeMust,
	}
)

var info *ent.AppControl

func rowToObject(row *ent.AppControl) *ent.AppControl {
	return &ent.AppControl{
		ID:                  row.ID,
		AppID:               row.AppID,
		SignupMethods:       row.SignupMethods,
		ExternSigninMethods: row.ExternSigninMethods,
		RecaptchaMethod:     row.RecaptchaMethod,
		KycEnable:           row.KycEnable,
		SigninVerifyEnable:  row.SigninVerifyEnable,
		InvitationCodeMust:  row.InvitationCodeMust,
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
			ID:                  uuid.New(),
			AppID:               uuid.New(),
			RecaptchaMethod:     uuid.New().String(),
			SignupMethods:       []string{uuid.New().String()},
			ExternSigninMethods: []string{uuid.New().String()},
			KycEnable:           false,
			SigninVerifyEnable:  false,
			InvitationCodeMust:  false,
		},
		{
			ID:                  uuid.New(),
			AppID:               uuid.New(),
			RecaptchaMethod:     uuid.New().String(),
			SignupMethods:       []string{uuid.New().String()},
			ExternSigninMethods: []string{uuid.New().String()},
			KycEnable:           false,
			SigninVerifyEnable:  false,
			InvitationCodeMust:  false,
		},
	}
	appcontrols := []*npool.AppControlReq{}
	for key := range entAppControl {
		id := entAppControl[key].ID.String()
		appID := entAppControl[key].AppID.String()

		appcontrols = append(appcontrols, &npool.AppControlReq{
			ID:                  &id,
			AppID:               &appID,
			RecaptchaMethod:     &entAppControl[key].RecaptchaMethod,
			SignupMethods:       entAppControl[key].SignupMethods,
			ExternSigninMethods: entAppControl[key].ExternSigninMethods,
			KycEnable:           &entAppControl[key].KycEnable,
			SigninVerifyEnable:  &entAppControl[key].SigninVerifyEnable,
			InvitationCodeMust:  &entAppControl[key].InvitationCodeMust,
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

func delete(t *testing.T) {
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
	t.Run("delete", delete)
	t.Run("count", count)
}
