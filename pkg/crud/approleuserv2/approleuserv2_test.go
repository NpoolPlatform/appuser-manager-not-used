package approleuserv2

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	val "github.com/NpoolPlatform/message/npool"

	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	testinit "github.com/NpoolPlatform/appuser-manager/pkg/test-init" //nolint
	npool "github.com/NpoolPlatform/message/npool/appusermgrv2/approleuser"

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

var entAppRoleUser = ent.AppRoleUser{
	AppID:  uuid.New(),
	RoleID: uuid.New(),
	UserID: uuid.New(),
	ID:     uuid.New(),
}

var (
	id     = entAppRoleUser.ID.String()
	appID  = entAppRoleUser.AppID.String()
	roleID = entAppRoleUser.RoleID.String()
	userID = entAppRoleUser.UserID.String()

	approleuserInfo = npool.AppRoleUserReq{
		ID:     &id,
		AppID:  &appID,
		RoleID: &roleID,
		UserID: &userID,
	}
)

var info *ent.AppRoleUser

func rowToObject(row *ent.AppRoleUser) *ent.AppRoleUser {
	return &ent.AppRoleUser{
		ID:     row.ID,
		AppID:  row.AppID,
		RoleID: row.RoleID,
		UserID: row.UserID,
	}
}

func create(t *testing.T) {
	var err error
	info, err = Create(context.Background(), &approleuserInfo)
	if assert.Nil(t, err) {
		if assert.NotEqual(t, info.ID, uuid.UUID{}.String()) {
			entAppRoleUser.ID = info.ID
		}
		assert.Equal(t, rowToObject(info), &entAppRoleUser)
	}
}

func createBulk(t *testing.T) {
	entAppRoleUser := []ent.AppRoleUser{
		{
			AppID:  uuid.New(),
			RoleID: uuid.New(),
			UserID: uuid.New(),
			ID:     uuid.New(),
		},
		{
			AppID:  uuid.New(),
			RoleID: uuid.New(),
			UserID: uuid.New(),
			ID:     uuid.New(),
		},
	}
	approleusers := []*npool.AppRoleUserReq{}
	for key := range entAppRoleUser {
		id := entAppRoleUser[key].ID.String()
		appID := entAppRoleUser[key].AppID.String()
		roleID := entAppRoleUser[key].RoleID.String()
		userID := entAppRoleUser[key].UserID.String()

		approleusers = append(approleusers, &npool.AppRoleUserReq{
			ID:     &id,
			AppID:  &appID,
			RoleID: &roleID,
			UserID: &userID,
		})
	}
	infos, err := CreateBulk(context.Background(), approleusers)
	if assert.Nil(t, err) {
		assert.Equal(t, len(infos), 2)
		assert.NotEqual(t, infos[0].ID, uuid.UUID{}.String())
		assert.NotEqual(t, infos[1].ID, uuid.UUID{}.String())
	}
}

func update(t *testing.T) {
	var err error
	info, err = Update(context.Background(), &approleuserInfo)
	if assert.Nil(t, err) {
		assert.Equal(t, rowToObject(info), &entAppRoleUser)
	}
}

func row(t *testing.T) {
	var err error
	info, err = Row(context.Background(), info.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, rowToObject(info), &entAppRoleUser)
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
		assert.Equal(t, rowToObject(infos[0]), &entAppRoleUser)
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
		assert.Equal(t, rowToObject(info), &entAppRoleUser)
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
		assert.Equal(t, rowToObject(info), &entAppRoleUser)
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
