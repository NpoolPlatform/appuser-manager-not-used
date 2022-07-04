package appv2

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	"github.com/NpoolPlatform/appuser-manager/pkg/test-init" //nolint
	npool "github.com/NpoolPlatform/message/npool/appusermgrv2/app"

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

var appRes = npool.AppRes{
	ID:          uuid.New().String(),
	CreatedBy:   uuid.New().String(),
	Name:        uuid.New().String(),
	Description: uuid.New().String(),
	Logo:        uuid.New().String(),
}

var appInfo = npool.App{
	ID:          &appRes.ID,
	CreatedBy:   &appRes.CreatedBy,
	Name:        &appRes.Name,
	Description: &appRes.Description,
	Logo:        &appRes.Logo,
}

var info *npool.AppRes

func Create(t *testing.T) {
	schema, err := New(context.Background(), nil)
	assert.Nil(t, err)

	info, err = schema.Create(context.Background(), &appInfo)
	if assert.Nil(t, err) {
		if assert.NotEqual(t, info.ID, uuid.UUID{}.String()) {
			appRes.ID = info.GetID()
			appRes.CreateAt = info.GetCreateAt()
		}
		assert.Equal(t, info, &appRes)
	}
}

func CreateBulk(t *testing.T) {
	appRes := []npool.AppRes{
		{
			ID:          uuid.New().String(),
			CreatedBy:   uuid.New().String(),
			Name:        uuid.New().String(),
			Description: uuid.New().String(),
			Logo:        uuid.New().String(),
		},
		{
			ID:          uuid.New().String(),
			CreatedBy:   uuid.New().String(),
			Name:        uuid.New().String(),
			Description: uuid.New().String(),
			Logo:        uuid.New().String(),
		},
	}
	apps := []*npool.App{}
	for key := range appRes {
		apps = append(apps, &npool.App{
			ID:          &appRes[key].ID,
			CreatedBy:   &appRes[key].CreatedBy,
			Name:        &appRes[key].Name,
			Logo:        &appRes[key].Logo,
			Description: &appRes[key].Description,
			CreateAt:    &appRes[key].CreateAt,
		})
	}
	schema, err := New(context.Background(), nil)
	assert.Nil(t, err)

	infos, err := schema.CreateBulk(context.Background(), apps)
	if assert.Nil(t, err) {
		assert.Equal(t, len(infos), 2)
		assert.NotEqual(t, infos[0].ID, uuid.UUID{}.String())
		assert.NotEqual(t, infos[1].ID, uuid.UUID{}.String())
	}
}

func Update(t *testing.T) {
	schema, err := New(context.Background(), nil)
	assert.Nil(t, err)

	info, err = schema.Update(context.Background(), &appInfo)
	if assert.Nil(t, err) {
		assert.Equal(t, info, &appRes)
	}
}

func Row(t *testing.T) {
	schema, err := New(context.Background(), nil)
	assert.Nil(t, err)

	info, err := schema.Row(context.Background(), uuid.MustParse(info.ID))
	if assert.Nil(t, err) {
		assert.Equal(t, info, &appRes)
	}
}

func Rows(t *testing.T) {
	schema, err := New(context.Background(), nil)
	assert.Nil(t, err)
	infos, total, err := schema.Rows(context.Background(),
		&npool.Conds{
			ID: &npool.IDVal{
				Value: info.GetID(),
				Op:    cruder.EQ,
			},
		}, 0, 0)
	if assert.Nil(t, err) {
		assert.Equal(t, total, 1)
		assert.Equal(t, infos[0], &appRes)
	}
}

func RowOnly(t *testing.T) {
	schema, err := New(context.Background(), nil)
	assert.Nil(t, err)

	info, err = schema.RowOnly(context.Background(),
		&npool.Conds{
			ID: &npool.IDVal{
				Value: info.GetID(),
				Op:    cruder.EQ,
			},
		})
	if assert.Nil(t, err) {
		assert.Equal(t, info, &appRes)
	}
}

func Count(t *testing.T) {
	schema, err := New(context.Background(), nil)
	assert.Nil(t, err)

	count, err := schema.Count(context.Background(),
		&npool.Conds{
			ID: &npool.IDVal{
				Value: info.GetID(),
				Op:    cruder.EQ,
			},
		},
	)
	if assert.Nil(t, err) {
		assert.Equal(t, count, count)
	}
}

func Exist(t *testing.T) {
	schema, err := New(context.Background(), nil)
	assert.Nil(t, err)
	exist, err := schema.Exist(context.Background(), uuid.MustParse(info.ID))
	if assert.Nil(t, err) {
		assert.Equal(t, exist, true)
	}
}

func ExistConds(t *testing.T) {
	schema, err := New(context.Background(), nil)
	assert.Nil(t, err)

	exist, err := schema.ExistConds(context.Background(),
		&npool.Conds{
			ID: &npool.IDVal{
				Value: info.GetID(),
				Op:    cruder.EQ,
			},
		},
	)
	if assert.Nil(t, err) {
		assert.Equal(t, exist, true)
	}
}

func Delete(t *testing.T) {
	schema, err := New(context.Background(), nil)
	assert.Nil(t, err)

	info, err = schema.Delete(context.Background(), uuid.MustParse(info.ID))
	if assert.Nil(t, err) {
		assert.Equal(t, info, &appRes)
	}
}

func TestMainOrder(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	t.Run("create", Create)
	t.Run("createBulk", CreateBulk)
	t.Run("row", Row)
	t.Run("rows", Rows)
	t.Run("rowOnly", RowOnly)
	t.Run("update", Update)
	t.Run("exist", Exist)
	t.Run("existConds", ExistConds)
	t.Run("delete", Delete)
	t.Run("count", Count)
}
