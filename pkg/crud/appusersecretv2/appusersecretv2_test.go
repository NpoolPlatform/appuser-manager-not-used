package appusersecretv2

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
	npool "github.com/NpoolPlatform/message/npool/appusermgrv2/appusersecret"

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

var entAppUserSecret = ent.AppUserSecret{
	GoogleSecret: uuid.New().String(),
	ID:           uuid.New(),
	AppID:        uuid.New(),
	UserID:       uuid.New(),
	PasswordHash: uuid.New().String(),
	Salt:         uuid.New().String(),
}

var (
	appID  = entAppUserSecret.AppID.String()
	userID = entAppUserSecret.UserID.String()
	id     = entAppUserSecret.ID.String()

	appusersecretInfo = npool.AppUserSecretReq{
		AppID:        &appID,
		UserID:       &userID,
		PasswordHash: &entAppUserSecret.PasswordHash,
		Salt:         &entAppUserSecret.Salt,
		GoogleSecret: &entAppUserSecret.GoogleSecret,
		ID:           &id,
	}
)

var info *ent.AppUserSecret

func rowToObject(row *ent.AppUserSecret) *ent.AppUserSecret {
	return &ent.AppUserSecret{
		AppID:        row.AppID,
		UserID:       row.UserID,
		PasswordHash: row.PasswordHash,
		Salt:         row.Salt,
		GoogleSecret: row.GoogleSecret,
		ID:           row.ID,
	}
}

func create(t *testing.T) {
	var err error
	info, err = Create(context.Background(), &appusersecretInfo)
	if assert.Nil(t, err) {
		if assert.NotEqual(t, info.ID, uuid.UUID{}.String()) {
			entAppUserSecret.ID = info.ID
		}
		assert.Equal(t, rowToObject(info), &entAppUserSecret)
	}
}

func createBulk(t *testing.T) {
	entAppUserSecret := []ent.AppUserSecret{
		{
			GoogleSecret: uuid.New().String(),
			ID:           uuid.New(),
			AppID:        uuid.New(),
			UserID:       uuid.New(),
			PasswordHash: uuid.New().String(),
			Salt:         uuid.New().String(),
		},
		{
			GoogleSecret: uuid.New().String(),
			ID:           uuid.New(),
			AppID:        uuid.New(),
			UserID:       uuid.New(),
			PasswordHash: uuid.New().String(),
			Salt:         uuid.New().String(),
		},
	}

	appusersecrets := []*npool.AppUserSecretReq{}
	for key := range entAppUserSecret {
		appID := entAppUserSecret[key].AppID.String()
		userID := entAppUserSecret[key].UserID.String()
		id := entAppUserSecret[key].ID.String()

		appusersecrets = append(appusersecrets, &npool.AppUserSecretReq{
			AppID:        &appID,
			UserID:       &userID,
			PasswordHash: &entAppUserSecret[key].PasswordHash,
			Salt:         &entAppUserSecret[key].Salt,
			GoogleSecret: &entAppUserSecret[key].GoogleSecret,
			ID:           &id,
		})
	}
	infos, err := CreateBulk(context.Background(), appusersecrets)
	if assert.Nil(t, err) {
		assert.Equal(t, len(infos), 2)
		assert.NotEqual(t, infos[0].ID, uuid.UUID{}.String())
		assert.NotEqual(t, infos[1].ID, uuid.UUID{}.String())
	}
}

func update(t *testing.T) {
	var err error
	info, err = Update(context.Background(), &appusersecretInfo)
	if assert.Nil(t, err) {
		assert.Equal(t, rowToObject(info), &entAppUserSecret)
	}
}

func row(t *testing.T) {
	var err error
	info, err = Row(context.Background(), info.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, rowToObject(info), &entAppUserSecret)
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
		assert.Equal(t, rowToObject(infos[0]), &entAppUserSecret)
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
		assert.Equal(t, rowToObject(info), &entAppUserSecret)
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
		assert.Equal(t, rowToObject(info), &entAppUserSecret)
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
