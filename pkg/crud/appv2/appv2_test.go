package appv2

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	testinit "github.com/NpoolPlatform/appuser-manager/pkg/test-init" //nolint
	val "github.com/NpoolPlatform/message/npool"
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
	appInfo   = npool.AppReq{
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

func create(t *testing.T) {
	var err error
	info, err = Create(context.Background(), &appInfo)
	if assert.Nil(t, err) {
		if assert.NotEqual(t, info.ID, uuid.UUID{}.String()) {
			entApp.ID = info.ID
			entApp.CreatedAt = info.CreatedAt
		}
		assert.Equal(t, rowToObject(info), &entApp)
	}
}

func createBulk(t *testing.T) {
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

	apps := []*npool.AppReq{}
	for key := range entApp {
		id := entApp[key].ID.String()
		createdBy := entApp[key].CreatedBy.String()
		apps = append(apps, &npool.AppReq{
			ID:          &id,
			CreatedBy:   &createdBy,
			Name:        &entApp[key].Name,
			Logo:        &entApp[key].Logo,
			Description: &entApp[key].Description,
			CreatedAt:   &entApp[key].CreatedAt,
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
		assert.Equal(t, rowToObject(info), &entApp)
	}
}

func row(t *testing.T) {
	var err error
	info, err = Row(context.Background(), info.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, rowToObject(info), &entApp)
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
		assert.Equal(t, rowToObject(infos[0]), &entApp)
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
		assert.Equal(t, rowToObject(info), &entApp)
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
		assert.Equal(t, rowToObject(info), &entApp)
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
