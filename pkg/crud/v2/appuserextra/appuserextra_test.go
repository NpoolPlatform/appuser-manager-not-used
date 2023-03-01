//nolint:nolintlint,dupl
package appuserextra

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	val "github.com/NpoolPlatform/message/npool"

	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	testinit "github.com/NpoolPlatform/appuser-manager/pkg/testinit"
	npool "github.com/NpoolPlatform/message/npool/appuser/mgr/v2/appuserextra"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"

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

var ret = ent.AppUserExtra{
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
	ActionCredits: decimal.NewFromInt(0),
}

var (
	id     = ret.ID.String()
	userID = ret.UserID.String()
	appID  = ret.AppID.String()

	req = npool.AppUserExtraReq{
		FirstName:     &ret.FirstName,
		IDNumber:      &ret.IDNumber,
		ID:            &id,
		LastName:      &ret.LastName,
		PostalCode:    &ret.PostalCode,
		Age:           &ret.Age,
		Organization:  &ret.Organization,
		UserID:        &userID,
		Birthday:      &ret.Birthday,
		AppID:         &appID,
		Username:      &ret.Username,
		Gender:        &ret.Gender,
		Avatar:        &ret.Avatar,
		AddressFields: ret.AddressFields,
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
		ActionCredits: row.ActionCredits,
	}
}

func create(t *testing.T) {
	var err error
	info, err = Create(context.Background(), &req)
	if assert.Nil(t, err) {
		if assert.NotEqual(t, info.ID, uuid.UUID{}.String()) {
			ret.ID = info.ID
		}
		assert.Equal(t, rowToObject(info).String(), ret.String())
	}
}

func createBulk(t *testing.T) {
	ret := []ent.AppUserExtra{
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
			ActionCredits: decimal.NewFromInt(0),
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
			ActionCredits: decimal.NewFromInt(0),
		},
	}

	appuserextras := []*npool.AppUserExtraReq{}
	for key := range ret {
		id := ret[key].ID.String()
		userID := ret[key].UserID.String()
		appID := ret[key].AppID.String()

		appuserextras = append(appuserextras, &npool.AppUserExtraReq{
			FirstName:     &ret[key].FirstName,
			IDNumber:      &ret[key].IDNumber,
			ID:            &id,
			LastName:      &ret[key].LastName,
			PostalCode:    &ret[key].PostalCode,
			Age:           &ret[key].Age,
			Organization:  &ret[key].Organization,
			UserID:        &userID,
			Birthday:      &ret[key].Birthday,
			AppID:         &appID,
			Username:      &ret[key].Username,
			Gender:        &ret[key].Gender,
			Avatar:        &ret[key].Avatar,
			AddressFields: ret[key].AddressFields,
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

	credits := "123.1"

	ret.ActionCredits = decimal.RequireFromString(credits)
	req.ActionCredits = &credits

	info, err = Update(context.Background(), &req)
	if assert.Nil(t, err) {
		assert.Equal(t, rowToObject(info).String(), ret.String())
	}

	ret.ActionCredits = ret.ActionCredits.Add(ret.ActionCredits)

	info, err = Update(context.Background(), &req)
	if assert.Nil(t, err) {
		assert.Equal(t, rowToObject(info).String(), ret.String())
	}
}

func row(t *testing.T) {
	var err error
	info, err = Row(context.Background(), info.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, rowToObject(info).String(), ret.String())
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
		assert.Equal(t, rowToObject(infos[0]).String(), ret.String())
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
		assert.Equal(t, rowToObject(info).String(), ret.String())
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
		assert.Equal(t, rowToObject(info).String(), ret.String())
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
