package appusercontrol

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
	npool "github.com/NpoolPlatform/message/npool/appuser/mgr/v2/appusercontrol"

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

var entAppUserControl = ent.AppUserControl{
	ID:                                 uuid.New(),
	AppID:                              uuid.New(),
	UserID:                             uuid.New(),
	SigninVerifyByGoogleAuthentication: true,
	GoogleAuthenticationVerified:       true,
}

var (
	appID  = entAppUserControl.AppID.String()
	userID = entAppUserControl.UserID.String()
	id     = entAppUserControl.ID.String()

	appusercontrolInfo = npool.AppUserControlReq{
		ID:                                 &id,
		AppID:                              &appID,
		UserID:                             &userID,
		SigninVerifyByGoogleAuthentication: &entAppUserControl.SigninVerifyByGoogleAuthentication,
		GoogleAuthenticationVerified:       &entAppUserControl.GoogleAuthenticationVerified,
	}
)

var info *ent.AppUserControl

func rowToObject(row *ent.AppUserControl) *ent.AppUserControl {
	return &ent.AppUserControl{
		ID:                                 row.ID,
		AppID:                              row.AppID,
		UserID:                             row.UserID,
		SigninVerifyByGoogleAuthentication: row.SigninVerifyByGoogleAuthentication,
		GoogleAuthenticationVerified:       row.GoogleAuthenticationVerified,
	}
}

func create(t *testing.T) {
	var err error
	info, err = Create(context.Background(), &appusercontrolInfo)
	if assert.Nil(t, err) {
		if assert.NotEqual(t, info.ID, uuid.UUID{}.String()) {
			entAppUserControl.ID = info.ID
		}
		assert.Equal(t, rowToObject(info), &entAppUserControl)
	}
}

func createBulk(t *testing.T) {
	entAppUserControl := []ent.AppUserControl{
		{
			ID:                                 uuid.New(),
			AppID:                              uuid.New(),
			UserID:                             uuid.New(),
			SigninVerifyByGoogleAuthentication: true,
			GoogleAuthenticationVerified:       true,
		},
		{
			ID:                                 uuid.New(),
			AppID:                              uuid.New(),
			UserID:                             uuid.New(),
			SigninVerifyByGoogleAuthentication: true,
			GoogleAuthenticationVerified:       true,
		},
	}

	appusercontrols := []*npool.AppUserControlReq{}
	for key := range entAppUserControl {
		appID := entAppUserControl[key].AppID.String()
		userID := entAppUserControl[key].UserID.String()
		id := entAppUserControl[key].ID.String()

		appusercontrols = append(appusercontrols, &npool.AppUserControlReq{
			ID:                                 &id,
			AppID:                              &appID,
			UserID:                             &userID,
			SigninVerifyByGoogleAuthentication: &entAppUserControl[key].SigninVerifyByGoogleAuthentication,
			GoogleAuthenticationVerified:       &entAppUserControl[key].GoogleAuthenticationVerified,
		})
	}
	infos, err := CreateBulk(context.Background(), appusercontrols)
	if assert.Nil(t, err) {
		assert.Equal(t, len(infos), 2)
		assert.NotEqual(t, infos[0].ID, uuid.UUID{}.String())
		assert.NotEqual(t, infos[1].ID, uuid.UUID{}.String())
	}
}

func update(t *testing.T) {
	var err error
	info, err = Update(context.Background(), &appusercontrolInfo)
	if assert.Nil(t, err) {
		assert.Equal(t, rowToObject(info), &entAppUserControl)
	}
}

func row(t *testing.T) {
	var err error
	info, err = Row(context.Background(), info.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, rowToObject(info), &entAppUserControl)
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
		assert.Equal(t, rowToObject(infos[0]), &entAppUserControl)
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
		assert.Equal(t, rowToObject(info), &entAppUserControl)
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
		assert.Equal(t, rowToObject(info), &entAppUserControl)
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