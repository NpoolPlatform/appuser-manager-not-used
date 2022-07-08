package appuserextrav2

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
	npool "github.com/NpoolPlatform/message/npool/appusermgrv2/appuserextra"

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

var entAppUserExtra = ent.AppUserExtra{
	UserID:        uuid.New(),
	Age:           uint32(10),
	Birthday:      uint32(10),
	AppID:         uuid.New(),
	Username:      uuid.New().String(),
	PostalCode:    uuid.New().String(),
	Organization:  uuid.New().String(),
	ID:            uuid.New(),
	LastName:      uuid.New().String(),
	Gender:        uuid.New().String(),
	Avatar:        uuid.New().String(),
	FirstName:     uuid.New().String(),
	IDNumber:      uuid.New().String(),
	AddressFields: []string{uuid.New().String(), uuid.New().String()},
}

var (
	id     = entAppUserExtra.ID.String()
	userID = entAppUserExtra.UserID.String()
	appID  = entAppUserExtra.AppID.String()

	appuserextraInfo = npool.AppUserExtraReq{
		FirstName:     &entAppUserExtra.FirstName,
		IDNumber:      &entAppUserExtra.IDNumber,
		ID:            &id,
		LastName:      &entAppUserExtra.LastName,
		PostalCode:    &entAppUserExtra.PostalCode,
		Age:           &entAppUserExtra.Age,
		Organization:  &entAppUserExtra.Organization,
		UserID:        &userID,
		Birthday:      &entAppUserExtra.Birthday,
		AppID:         &appID,
		Username:      &entAppUserExtra.Username,
		Gender:        &entAppUserExtra.Gender,
		Avatar:        &entAppUserExtra.Avatar,
		AddressFields: entAppUserExtra.AddressFields,
	}
)

var info *ent.AppUserExtra

func rowToObject(row *ent.AppUserExtra) *ent.AppUserExtra {
	return &ent.AppUserExtra{
		FirstName:     row.FirstName,
		IDNumber:      row.IDNumber,
		ID:            row.ID,
		LastName:      row.LastName,
		PostalCode:    row.PostalCode,
		Age:           row.Age,
		Organization:  row.Organization,
		UserID:        row.UserID,
		Birthday:      row.Birthday,
		AppID:         row.AppID,
		Username:      row.Username,
		Gender:        row.Gender,
		Avatar:        row.Avatar,
		AddressFields: row.AddressFields,
	}
}

func create(t *testing.T) {
	var err error
	info, err = Create(context.Background(), &appuserextraInfo)
	if assert.Nil(t, err) {
		if assert.NotEqual(t, info.ID, uuid.UUID{}.String()) {
			entAppUserExtra.ID = info.ID
		}
		assert.Equal(t, rowToObject(info), &entAppUserExtra)
	}
}

func createBulk(t *testing.T) {
	entAppUserExtra := []ent.AppUserExtra{
		{
			UserID:        uuid.New(),
			Age:           uint32(10),
			Birthday:      uint32(10),
			AppID:         uuid.New(),
			Username:      uuid.New().String(),
			PostalCode:    uuid.New().String(),
			Organization:  uuid.New().String(),
			ID:            uuid.New(),
			LastName:      uuid.New().String(),
			Gender:        uuid.New().String(),
			Avatar:        uuid.New().String(),
			FirstName:     uuid.New().String(),
			IDNumber:      uuid.New().String(),
			AddressFields: []string{uuid.New().String(), uuid.New().String()},
		},
		{
			UserID:        uuid.New(),
			Age:           uint32(10),
			Birthday:      uint32(10),
			AppID:         uuid.New(),
			Username:      uuid.New().String(),
			PostalCode:    uuid.New().String(),
			Organization:  uuid.New().String(),
			ID:            uuid.New(),
			LastName:      uuid.New().String(),
			Gender:        uuid.New().String(),
			Avatar:        uuid.New().String(),
			FirstName:     uuid.New().String(),
			IDNumber:      uuid.New().String(),
			AddressFields: []string{uuid.New().String(), uuid.New().String()},
		},
	}

	appuserextras := []*npool.AppUserExtraReq{}
	for key := range entAppUserExtra {
		id := entAppUserExtra[key].ID.String()
		userID := entAppUserExtra[key].UserID.String()
		appID := entAppUserExtra[key].AppID.String()

		appuserextras = append(appuserextras, &npool.AppUserExtraReq{
			FirstName:     &entAppUserExtra[key].FirstName,
			IDNumber:      &entAppUserExtra[key].IDNumber,
			ID:            &id,
			LastName:      &entAppUserExtra[key].LastName,
			PostalCode:    &entAppUserExtra[key].PostalCode,
			Age:           &entAppUserExtra[key].Age,
			Organization:  &entAppUserExtra[key].Organization,
			UserID:        &userID,
			Birthday:      &entAppUserExtra[key].Birthday,
			AppID:         &appID,
			Username:      &entAppUserExtra[key].Username,
			Gender:        &entAppUserExtra[key].Gender,
			Avatar:        &entAppUserExtra[key].Avatar,
			AddressFields: entAppUserExtra[key].AddressFields,
		})
	}
	infos, err := CreateBulk(context.Background(), appuserextras)
	if assert.Nil(t, err) {
		assert.Equal(t, len(infos), 2)
		assert.NotEqual(t, infos[0].ID, uuid.UUID{}.String())
		assert.NotEqual(t, infos[1].ID, uuid.UUID{}.String())
	}
}

func update(t *testing.T) {
	var err error
	info, err = Update(context.Background(), &appuserextraInfo)
	if assert.Nil(t, err) {
		assert.Equal(t, rowToObject(info), &entAppUserExtra)
	}
}

func row(t *testing.T) {
	var err error
	info, err = Row(context.Background(), info.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, rowToObject(info), &entAppUserExtra)
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
		assert.Equal(t, rowToObject(infos[0]), &entAppUserExtra)
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
		assert.Equal(t, rowToObject(info), &entAppUserExtra)
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
		assert.Equal(t, rowToObject(info), &entAppUserExtra)
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
