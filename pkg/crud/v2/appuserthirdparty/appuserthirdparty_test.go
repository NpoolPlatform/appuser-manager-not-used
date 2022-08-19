package appuserthirdparty

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
	npool "github.com/NpoolPlatform/message/npool/appuser/mgr/v2/appuserthirdparty"

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

var entAppUserThirdParty = ent.AppUserThirdParty{
	ID:                 uuid.New(),
	AppID:              uuid.New(),
	UserID:             uuid.New(),
	ThirdPartyUserID:   uuid.New().String(),
	ThirdPartyID:       uuid.New().String(),
	ThirdPartyUsername: uuid.New().String(),
	ThirdPartyAvatar:   uuid.New().String(),
}

var (
	userID = entAppUserThirdParty.UserID.String()
	id     = entAppUserThirdParty.ID.String()
	appID  = entAppUserThirdParty.AppID.String()

	appuserthirdpartyInfo = npool.AppUserThirdPartyReq{
		UserID:             &userID,
		ThirdPartyUserID:   &entAppUserThirdParty.ThirdPartyUserID,
		ThirdPartyID:       &entAppUserThirdParty.ThirdPartyID,
		ThirdPartyUsername: &entAppUserThirdParty.ThirdPartyUsername,
		ThirdPartyAvatar:   &entAppUserThirdParty.ThirdPartyAvatar,
		ID:                 &id,
		AppID:              &appID,
	}
)

var info *ent.AppUserThirdParty

func rowToObject(row *ent.AppUserThirdParty) *ent.AppUserThirdParty {
	return &ent.AppUserThirdParty{
		UserID:             row.UserID,
		ThirdPartyUserID:   row.ThirdPartyUserID,
		ThirdPartyID:       row.ThirdPartyID,
		ThirdPartyUsername: row.ThirdPartyUsername,
		ThirdPartyAvatar:   row.ThirdPartyAvatar,
		ID:                 row.ID,
		AppID:              row.AppID,
	}
}

func create(t *testing.T) {
	var err error
	info, err = Create(context.Background(), &appuserthirdpartyInfo)
	if assert.Nil(t, err) {
		if assert.NotEqual(t, info.ID, uuid.UUID{}.String()) {
			entAppUserThirdParty.ID = info.ID
		}
		assert.Equal(t, rowToObject(info), &entAppUserThirdParty)
	}
}

func createBulk(t *testing.T) {
	entAppUserThirdParty := []ent.AppUserThirdParty{
		{
			ID:                 uuid.New(),
			AppID:              uuid.New(),
			UserID:             uuid.New(),
			ThirdPartyUserID:   uuid.New().String(),
			ThirdPartyID:       uuid.New().String(),
			ThirdPartyUsername: uuid.New().String(),
			ThirdPartyAvatar:   uuid.New().String(),
		},
		{
			ID:                 uuid.New(),
			AppID:              uuid.New(),
			UserID:             uuid.New(),
			ThirdPartyUserID:   uuid.New().String(),
			ThirdPartyID:       uuid.New().String(),
			ThirdPartyUsername: uuid.New().String(),
			ThirdPartyAvatar:   uuid.New().String(),
		},
	}

	appuserthirdpartys := []*npool.AppUserThirdPartyReq{}
	for key := range entAppUserThirdParty {
		userID := entAppUserThirdParty[key].UserID.String()
		id := entAppUserThirdParty[key].ID.String()
		appID := entAppUserThirdParty[key].AppID.String()

		appuserthirdpartys = append(appuserthirdpartys, &npool.AppUserThirdPartyReq{
			UserID:             &userID,
			ThirdPartyUserID:   &entAppUserThirdParty[key].ThirdPartyUserID,
			ThirdPartyID:       &entAppUserThirdParty[key].ThirdPartyID,
			ThirdPartyUsername: &entAppUserThirdParty[key].ThirdPartyUsername,
			ThirdPartyAvatar:   &entAppUserThirdParty[key].ThirdPartyAvatar,
			ID:                 &id,
			AppID:              &appID,
		})
	}
	infos, err := CreateBulk(context.Background(), appuserthirdpartys)
	if assert.Nil(t, err) {
		assert.Equal(t, len(infos), 2)
		assert.NotEqual(t, infos[0].ID, uuid.UUID{}.String())
		assert.NotEqual(t, infos[1].ID, uuid.UUID{}.String())
	}
}

func update(t *testing.T) {
	var err error
	info, err = Update(context.Background(), &appuserthirdpartyInfo)
	if assert.Nil(t, err) {
		assert.Equal(t, rowToObject(info), &entAppUserThirdParty)
	}
}

func row(t *testing.T) {
	var err error
	info, err = Row(context.Background(), info.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, rowToObject(info), &entAppUserThirdParty)
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
		assert.Equal(t, rowToObject(infos[0]), &entAppUserThirdParty)
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
		assert.Equal(t, rowToObject(info), &entAppUserThirdParty)
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
		assert.Equal(t, rowToObject(info), &entAppUserThirdParty)
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
