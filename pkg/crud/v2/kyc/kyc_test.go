package kyc

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
	npool "github.com/NpoolPlatform/message/npool/appuser/mgr/v2/kyc"
	reviewpb "github.com/NpoolPlatform/message/npool/review/mgr/v2"

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

var entKyc = ent.Kyc{
	ID:           uuid.New(),
	AppID:        uuid.New(),
	UserID:       uuid.New(),
	DocumentType: npool.KycDocumentType_IDCard.String(),
	IDNumber:     uuid.NewString(),
	FrontImg:     uuid.NewString(),
	BackImg:      uuid.NewString(),
	SelfieImg:    uuid.NewString(),
	EntityType:   npool.KycEntityType_Individual.String(),
	ReviewState:  reviewpb.ReviewState_Wait.String(),
}

var (
	id           = entKyc.ID.String()
	appID        = entKyc.AppID.String()
	userID       = entKyc.UserID.String()
	documentType = npool.KycDocumentType_IDCard
	entityType   = npool.KycEntityType_Individual
	reviewState  = reviewpb.ReviewState_Wait
	kycInfo      = npool.KycReq{
		ID:           &id,
		AppID:        &appID,
		UserID:       &userID,
		DocumentType: &documentType,
		IDNumber:     &entKyc.IDNumber,
		FrontImg:     &entKyc.FrontImg,
		BackImg:      &entKyc.BackImg,
		SelfieImg:    &entKyc.SelfieImg,
		EntityType:   &entityType,
		ReviewState:  &reviewState,
	}
)

var info *ent.Kyc

func create(t *testing.T) {
	var err error
	info, err = Create(context.Background(), &kycInfo)
	if assert.Nil(t, err) {
		if assert.NotEqual(t, info.ID, uuid.UUID{}.String()) {
			entKyc.ID = info.ID
			entKyc.CreatedAt = info.CreatedAt
			entKyc.UpdatedAt = info.UpdatedAt
		}
		assert.Equal(t, info.String(), entKyc.String())
	}
}

func createBulk(t *testing.T) {
	entKyc := []ent.Kyc{
		{
			ID:           uuid.New(),
			AppID:        uuid.New(),
			UserID:       uuid.New(),
			DocumentType: npool.KycDocumentType_IDCard.String(),
			IDNumber:     uuid.NewString(),
			FrontImg:     uuid.NewString(),
			BackImg:      uuid.NewString(),
			SelfieImg:    uuid.NewString(),
		},
		{
			ID:           uuid.New(),
			AppID:        uuid.New(),
			UserID:       uuid.New(),
			DocumentType: npool.KycDocumentType_IDCard.String(),
			IDNumber:     uuid.NewString(),
			FrontImg:     uuid.NewString(),
			BackImg:      uuid.NewString(),
			SelfieImg:    uuid.NewString(),
		},
	}

	apps := []*npool.KycReq{}
	for key := range entKyc {
		id := entKyc[key].ID.String()
		appID := entKyc[key].AppID.String()
		userID := entKyc[key].UserID.String()
		documentType = npool.KycDocumentType_IDCard

		apps = append(apps, &npool.KycReq{
			ID:           &id,
			AppID:        &appID,
			UserID:       &userID,
			DocumentType: &documentType,
			IDNumber:     &entKyc[key].IDNumber,
			FrontImg:     &entKyc[key].FrontImg,
			BackImg:      &entKyc[key].BackImg,
			SelfieImg:    &entKyc[key].SelfieImg,
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
	info, err = Update(context.Background(), &kycInfo)
	if assert.Nil(t, err) {
		entKyc.UpdatedAt = info.UpdatedAt
		assert.Equal(t, info.String(), entKyc.String())
	}
}

func row(t *testing.T) {
	var err error
	info, err = Row(context.Background(), info.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, info.String(), entKyc.String())
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
		assert.Equal(t, infos[0].String(), entKyc.String())
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
		assert.Equal(t, info.String(), entKyc.String())
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
		entKyc.DeletedAt = info.DeletedAt
		entKyc.UpdatedAt = info.UpdatedAt
		assert.Equal(t, info.String(), entKyc.String())
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
