package appuserv2

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
	npool "github.com/NpoolPlatform/message/npool/appusermgrv2/appuser"

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

var entAppUser = ent.AppUser{
	ID:            uuid.New(),
	AppID:         uuid.New(),
	EmailAddress:  uuid.New().String(),
	PhoneNo:       uuid.New().String(),
	ImportFromApp: uuid.New(),
}

var (
	id            = entAppUser.ID.String()
	appID         = entAppUser.AppID.String()
	importFromApp = entAppUser.ImportFromApp.String()

	appuserInfo = npool.AppUserReq{
		ID:            &id,
		AppID:         &appID,
		EmailAddress:  &entAppUser.EmailAddress,
		PhoneNo:       &entAppUser.PhoneNo,
		ImportFromApp: &importFromApp,
	}
)

var info *ent.AppUser

func rowToObject(row *ent.AppUser) *ent.AppUser {
	return &ent.AppUser{
		ID:            row.ID,
		AppID:         row.AppID,
		EmailAddress:  row.EmailAddress,
		PhoneNo:       row.PhoneNo,
		ImportFromApp: row.ImportFromApp,
	}
}

func create(t *testing.T) {
	var err error
	info, err = Create(context.Background(), &appuserInfo)
	if assert.Nil(t, err) {
		if assert.NotEqual(t, info.ID, uuid.UUID{}.String()) {
			entAppUser.ID = info.ID
		}
		assert.Equal(t, rowToObject(info), &entAppUser)
	}
}

func createBulk(t *testing.T) {
	entAppUser := []ent.AppUser{
		{
			ID:            uuid.New(),
			AppID:         uuid.New(),
			EmailAddress:  uuid.New().String(),
			PhoneNo:       uuid.New().String(),
			ImportFromApp: uuid.New(),
		},
		{
			ID:            uuid.New(),
			AppID:         uuid.New(),
			EmailAddress:  uuid.New().String(),
			PhoneNo:       uuid.New().String(),
			ImportFromApp: uuid.New(),
		},
	}

	appusers := []*npool.AppUserReq{}
	for key := range entAppUser {
		id := entAppUser[key].ID.String()
		appID := entAppUser[key].AppID.String()
		importFromApp := entAppUser[key].ImportFromApp.String()

		appusers = append(appusers, &npool.AppUserReq{
			ID:            &id,
			AppID:         &appID,
			EmailAddress:  &entAppUser[key].EmailAddress,
			PhoneNo:       &entAppUser[key].PhoneNo,
			ImportFromApp: &importFromApp,
		})
	}
	infos, err := CreateBulk(context.Background(), appusers)
	if assert.Nil(t, err) {
		assert.Equal(t, len(infos), 2)
		assert.NotEqual(t, infos[0].ID, uuid.UUID{}.String())
		assert.NotEqual(t, infos[1].ID, uuid.UUID{}.String())
	}
}

func update(t *testing.T) {
	var err error
	info, err = Update(context.Background(), &appuserInfo)
	if assert.Nil(t, err) {
		assert.Equal(t, rowToObject(info), &entAppUser)
	}
}

func row(t *testing.T) {
	var err error
	info, err = Row(context.Background(), info.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, rowToObject(info), &entAppUser)
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
		assert.Equal(t, rowToObject(infos[0]), &entAppUser)
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
		assert.Equal(t, rowToObject(info), &entAppUser)
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
		assert.Equal(t, rowToObject(info), &entAppUser)
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
