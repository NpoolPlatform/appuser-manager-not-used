package banappuserv2

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
	npool "github.com/NpoolPlatform/message/npool/appusermgrv2/banappuser"

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

var entBanAppUser = ent.BanAppUser{
	AppID:   uuid.New(),
	UserID:  uuid.New(),
	Message: uuid.New().String(),
	ID:      uuid.New(),
}

var (
	appID  = entBanAppUser.AppID.String()
	userID = entBanAppUser.UserID.String()
	id     = entBanAppUser.ID.String()

	banappuserInfo = npool.BanAppUserReq{
		AppID:   &appID,
		UserID:  &userID,
		Message: &entBanAppUser.Message,
		ID:      &id,
	}
)

var info *ent.BanAppUser

func rowToObject(row *ent.BanAppUser) *ent.BanAppUser {
	return &ent.BanAppUser{
		AppID:   row.AppID,
		UserID:  row.UserID,
		Message: row.Message,
		ID:      row.ID,
	}
}

func create(t *testing.T) {
	var err error
	info, err = Create(context.Background(), &banappuserInfo)
	if assert.Nil(t, err) {
		if assert.NotEqual(t, info.ID, uuid.UUID{}.String()) {
			entBanAppUser.ID = info.ID
		}
		assert.Equal(t, rowToObject(info), &entBanAppUser)
	}
}

func createBulk(t *testing.T) {
	entBanAppUser := []ent.BanAppUser{
		{
			AppID:   uuid.New(),
			UserID:  uuid.New(),
			Message: uuid.New().String(),
			ID:      uuid.New(),
		},
		{
			AppID:   uuid.New(),
			UserID:  uuid.New(),
			Message: uuid.New().String(),
			ID:      uuid.New(),
		},
	}

	banappusers := []*npool.BanAppUserReq{}
	for key := range entBanAppUser {
		appID := entBanAppUser[key].AppID.String()
		userID := entBanAppUser[key].UserID.String()
		id := entBanAppUser[key].ID.String()

		banappusers = append(banappusers, &npool.BanAppUserReq{
			AppID:   &appID,
			UserID:  &userID,
			Message: &entBanAppUser[key].Message,
			ID:      &id,
		})
	}
	infos, err := CreateBulk(context.Background(), banappusers)
	if assert.Nil(t, err) {
		assert.Equal(t, len(infos), 2)
		assert.NotEqual(t, infos[0].ID, uuid.UUID{}.String())
		assert.NotEqual(t, infos[1].ID, uuid.UUID{}.String())
	}
}

func update(t *testing.T) {
	var err error
	info, err = Update(context.Background(), &banappuserInfo)
	if assert.Nil(t, err) {
		assert.Equal(t, rowToObject(info), &entBanAppUser)
	}
}

func row(t *testing.T) {
	var err error
	info, err = Row(context.Background(), info.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, rowToObject(info), &entBanAppUser)
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
		assert.Equal(t, rowToObject(infos[0]), &entBanAppUser)
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
		assert.Equal(t, rowToObject(info), &entBanAppUser)
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
		assert.Equal(t, rowToObject(info), &entBanAppUser)
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
