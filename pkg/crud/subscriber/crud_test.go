package subscriber

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
	npool "github.com/NpoolPlatform/message/npool/appuser/mgr/v2/subscriber"

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

var ret = ent.Subscriber{
	ID:           uuid.New(),
	AppID:        uuid.New(),
	EmailAddress: uuid.New().String(),
}

var (
	id    = ret.ID.String()
	appID = ret.AppID.String()
	req   = npool.SubscriberReq{
		ID:           &id,
		AppID:        &appID,
		EmailAddress: &ret.EmailAddress,
	}
)

var info *ent.Subscriber

func rowToObject(row *ent.Subscriber) *ent.Subscriber {
	return &ent.Subscriber{
		ID:           row.ID,
		AppID:        row.AppID,
		EmailAddress: row.EmailAddress,
		Registered:   row.Registered,
		CreatedAt:    row.CreatedAt,
		UpdatedAt:    row.UpdatedAt,
	}
}

func create(t *testing.T) {
	var err error
	info, err = Create(context.Background(), &req)
	if assert.Nil(t, err) {
		if assert.NotEqual(t, info.ID, uuid.UUID{}.String()) {
			ret.ID = info.ID
			ret.CreatedAt = info.CreatedAt
			ret.UpdatedAt = info.UpdatedAt
		}
		assert.Equal(t, rowToObject(info), &ret)
	}
}

func createBulk(t *testing.T) {
	ret := []ent.Subscriber{
		{
			ID:           uuid.New(),
			AppID:        uuid.New(),
			EmailAddress: uuid.New().String(),
		},
		{
			ID:           uuid.New(),
			AppID:        uuid.New(),
			EmailAddress: uuid.New().String(),
		},
	}

	apps := []*npool.SubscriberReq{}
	for key := range ret {
		id := ret[key].ID.String()
		appID := ret[key].AppID.String()
		apps = append(apps, &npool.SubscriberReq{
			ID:           &id,
			AppID:        &appID,
			EmailAddress: &ret[key].EmailAddress,
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

	registered := true
	req.Registered = &registered
	ret.Registered = registered

	info, err = Update(context.Background(), &req)
	if assert.Nil(t, err) {
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, rowToObject(info), &ret)
	}
}

func row(t *testing.T) {
	var err error
	info, err = Row(context.Background(), info.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, rowToObject(info), &ret)
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
		assert.Equal(t, rowToObject(infos[0]), &ret)
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
		assert.Equal(t, rowToObject(info), &ret)
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
		assert.Equal(t, rowToObject(info), &ret)
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
