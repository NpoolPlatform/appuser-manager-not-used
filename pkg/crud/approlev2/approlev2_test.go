package approlev2

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
	npool "github.com/NpoolPlatform/message/npool/appusermgrv2/approle"

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

var entAppRole = ent.AppRole{
	ID:          uuid.New(),
	AppID:       uuid.New(),
	CreatedBy:   uuid.New(),
	Role:        uuid.New().String(),
	Description: uuid.New().String(),
	Default:     true,
}

var (
	id          = entAppRole.ID.String()
	appID       = entAppRole.AppID.String()
	createdBy   = entAppRole.CreatedBy.String()
	appRoleInfo = npool.AppRoleReq{
		ID:          &id,
		AppID:       &appID,
		CreatedBy:   &createdBy,
		Role:        &entAppRole.Role,
		Description: &entAppRole.Description,
		Default:     &entAppRole.Default,
	}
)

var info *ent.AppRole

func rowToObject(row *ent.AppRole) *ent.AppRole {
	return &ent.AppRole{
		ID:          row.ID,
		AppID:       row.AppID,
		CreatedBy:   row.CreatedBy,
		Role:        row.Role,
		Description: row.Description,
		Default:     row.Default,
	}
}

func create(t *testing.T) {
	var err error
	info, err = Create(context.Background(), &appRoleInfo)
	if assert.Nil(t, err) {
		if assert.NotEqual(t, info.ID, uuid.UUID{}.String()) {
			entAppRole.ID = info.ID
		}
		assert.Equal(t, rowToObject(info), &entAppRole)
	}
}

func createBulk(t *testing.T) {
	entAppRole := []ent.AppRole{
		{
			ID:          uuid.New(),
			AppID:       uuid.New(),
			CreatedBy:   uuid.New(),
			Role:        uuid.New().String(),
			Description: uuid.New().String(),
			Default:     true,
		},
		{
			ID:          uuid.New(),
			AppID:       uuid.New(),
			CreatedBy:   uuid.New(),
			Role:        uuid.New().String(),
			Description: uuid.New().String(),
			Default:     true,
		},
	}

	appRoles := []*npool.AppRoleReq{}
	for key := range entAppRole {
		id := entAppRole[key].ID.String()
		createdBy := entAppRole[key].CreatedBy.String()
		appRoles = append(appRoles, &npool.AppRoleReq{
			ID:          &id,
			AppID:       &appID,
			CreatedBy:   &createdBy,
			Role:        &entAppRole[key].Role,
			Description: &entAppRole[key].Description,
			Default:     &entAppRole[key].Default,
		})
	}
	infos, err := CreateBulk(context.Background(), appRoles)
	if assert.Nil(t, err) {
		assert.Equal(t, len(infos), 2)
		assert.NotEqual(t, infos[0].ID, uuid.UUID{}.String())
		assert.NotEqual(t, infos[1].ID, uuid.UUID{}.String())
	}
}

func update(t *testing.T) {
	var err error
	info, err = Update(context.Background(), &appRoleInfo)
	if assert.Nil(t, err) {
		assert.Equal(t, rowToObject(info), &entAppRole)
	}
}

func row(t *testing.T) {
	var err error
	info, err = Row(context.Background(), info.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, rowToObject(info), &entAppRole)
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
		assert.Equal(t, rowToObject(infos[0]), &entAppRole)
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
		assert.Equal(t, rowToObject(info), &entAppRole)
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
		assert.Equal(t, rowToObject(info), &entAppRole)
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
