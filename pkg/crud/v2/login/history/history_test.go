package history

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	testinit "github.com/NpoolPlatform/appuser-manager/pkg/testinit"
	val "github.com/NpoolPlatform/message/npool"
	npool "github.com/NpoolPlatform/message/npool/appuser/mgr/v2/login/history"

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

var entHistory = ent.LoginHistory{
	ID:        uuid.New(),
	AppID:     uuid.New(),
	UserID:    uuid.New(),
	ClientIP:  uuid.New().String(),
	UserAgent: uuid.New().String(),
}

var (
	id     = entHistory.ID.String()
	appID  = entHistory.AppID.String()
	userID = entHistory.UserID.String()

	appInfo = npool.HistoryReq{
		ID:        &id,
		AppID:     &appID,
		UserID:    &userID,
		ClientIP:  &entHistory.ClientIP,
		UserAgent: &entHistory.UserAgent,
	}
)

var info *ent.LoginHistory

func create(t *testing.T) {
	var err error
	info, err = Create(context.Background(), &appInfo)
	if assert.Nil(t, err) {
		if assert.NotEqual(t, info.ID, uuid.UUID{}.String()) {
			entHistory.ID = info.ID
			entHistory.CreatedAt = info.CreatedAt
			entHistory.UpdatedAt = info.UpdatedAt
		}
		assert.Equal(t, info.String(), entHistory.String())
	}
}

func createBulk(t *testing.T) {
	entHistory := []ent.LoginHistory{
		{
			ID:        uuid.New(),
			AppID:     uuid.New(),
			UserID:    uuid.New(),
			ClientIP:  uuid.New().String(),
			UserAgent: uuid.New().String(),
		},
		{
			ID:        uuid.New(),
			AppID:     uuid.New(),
			UserID:    uuid.New(),
			ClientIP:  uuid.New().String(),
			UserAgent: uuid.New().String(),
		},
	}

	apps := []*npool.HistoryReq{}
	for key := range entHistory {
		id := entHistory[key].ID.String()
		appID := entHistory[key].AppID.String()
		userID := entHistory[key].UserID.String()

		apps = append(apps, &npool.HistoryReq{
			ID:        &id,
			AppID:     &appID,
			UserID:    &userID,
			ClientIP:  &entHistory[key].ClientIP,
			UserAgent: &entHistory[key].UserAgent,
		})
	}
	infos, err := CreateBulk(context.Background(), apps)
	if assert.Nil(t, err) {
		assert.Equal(t, len(infos), 2)
		assert.NotEqual(t, infos[0].ID, uuid.UUID{}.String())
		assert.NotEqual(t, infos[1].ID, uuid.UUID{}.String())
	}
}

func update(t *testing.T) {
	var err error
	info, err = Update(context.Background(), &appInfo)
	if assert.Nil(t, err) {
		entHistory.DeletedAt = info.DeletedAt
		entHistory.CreatedAt = info.CreatedAt
		entHistory.UpdatedAt = info.UpdatedAt
		assert.Equal(t, info.String(), entHistory.String())
	}
}

func row(t *testing.T) {
	var err error
	info, err = Row(context.Background(), info.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, info.String(), entHistory.String())
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
		assert.Equal(t, infos[0].String(), entHistory.String())
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
		assert.Equal(t, info.String(), entHistory.String())
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
		entHistory.DeletedAt = info.DeletedAt
		entHistory.CreatedAt = info.CreatedAt
		entHistory.UpdatedAt = info.UpdatedAt
		assert.Equal(t, info.String(), entHistory.String())
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
