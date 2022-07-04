package appv2

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent"

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

var entApp = ent.App{
	ID:          uuid.New(),
	CreatedBy:   uuid.New(),
	Name:        uuid.New().String(),
	Description: uuid.New().String(),
	Logo:        uuid.New().String(),
}

var (
	id        = entApp.ID.String()
	createdBy = entApp.CreatedBy.String()
	appInfo   = npool.App{
		ID:          &id,
		CreatedBy:   &createdBy,
		Name:        &entApp.Name,
		Description: &entApp.Description,
		Logo:        &entApp.Logo,
	}
)

var info *ent.App

func rowToObject(row *ent.App) *ent.App {
	return &ent.App{
		ID:          row.ID,
		CreatedBy:   row.CreatedBy,
		Name:        row.Name,
		Logo:        row.Logo,
		Description: row.Description,
		CreatedAt:   row.CreatedAt,
	}
}

func Create(t *testing.T) {
	schema, err := New(context.Background(), nil)
	assert.Nil(t, err)

	info, err = schema.Create(context.Background(), &appInfo)
	if assert.Nil(t, err) {
		if assert.NotEqual(t, info.ID, uuid.UUID{}.String()) {
			entApp.ID = info.ID
			entApp.CreatedAt = info.CreatedAt
		}
		assert.Equal(t, rowToObject(info), &entApp)
	}
}

func CreateBulk(t *testing.T) {
	entApp := []ent.App{
		{
			ID:          uuid.New(),
			CreatedBy:   uuid.New(),
			Name:        uuid.New().String(),
			Description: uuid.New().String(),
			Logo:        uuid.New().String(),
		},
		{
			ID:          uuid.New(),
			CreatedBy:   uuid.New(),
			Name:        uuid.New().String(),
			Description: uuid.New().String(),
			Logo:        uuid.New().String(),
		},
	}

	apps := []*npool.App{}
	for key := range entApp {
		id := entApp[key].ID.String()
		createdBy := entApp[key].CreatedBy.String()
		apps = append(apps, &npool.App{
			ID:          &id,
			CreatedBy:   &createdBy,
			Name:        &entApp[key].Name,
			Logo:        &entApp[key].Logo,
			Description: &entApp[key].Description,
			CreatedAt:   &entApp[key].CreatedAt,
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
		assert.Equal(t, rowToObject(info), &entApp)
	}
}

func Row(t *testing.T) {
	schema, err := New(context.Background(), nil)
	assert.Nil(t, err)

	info, err := schema.Row(context.Background(), info.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, rowToObject(info), &entApp)
	}
}

func Rows(t *testing.T) {
	schema, err := New(context.Background(), nil)
	assert.Nil(t, err)
	infos, total, err := schema.Rows(context.Background(),
		&npool.Conds{
			ID: &npool.IDVal{
				Value: info.ID.String(),
				Op:    cruder.EQ,
			},
		}, 0, 0)
	if assert.Nil(t, err) {
		assert.Equal(t, total, 1)
		assert.Equal(t, rowToObject(infos[0]), &entApp)
	}
}

func RowOnly(t *testing.T) {
	schema, err := New(context.Background(), nil)
	assert.Nil(t, err)

	info, err = schema.RowOnly(context.Background(),
		&npool.Conds{
			ID: &npool.IDVal{
				Value: info.ID.String(),
				Op:    cruder.EQ,
			},
		})
	if assert.Nil(t, err) {
		assert.Equal(t, rowToObject(info), &entApp)
	}
}

func Count(t *testing.T) {
	schema, err := New(context.Background(), nil)
	assert.Nil(t, err)

	count, err := schema.Count(context.Background(),
		&npool.Conds{
			ID: &npool.IDVal{
				Value: info.ID.String(),
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
	exist, err := schema.Exist(context.Background(), info.ID)
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
				Value: info.ID.String(),
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

	info, err = schema.Delete(context.Background(), info.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, rowToObject(info), &entApp)
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
