package kyc

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/go-service-framework/pkg/config"

	"bou.ke/monkey"
	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	val "github.com/NpoolPlatform/message/npool"

	testinit "github.com/NpoolPlatform/appuser-manager/pkg/testinit"
	npool "github.com/NpoolPlatform/message/npool/appuser/mgr/v2/kyc"
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

var kycData = npool.Kyc{
	ID:           uuid.NewString(),
	AppID:        uuid.NewString(),
	UserID:       uuid.NewString(),
	DocumentType: npool.KycDocumentType_IDCard,
	IDNumber:     uuid.NewString(),
	FrontImg:     uuid.NewString(),
	BackImg:      uuid.NewString(),
	SelfieImg:    uuid.NewString(),
	EntityType:   npool.KycEntityType_Individual,
	State:        npool.KycState_Reviewing,
	ReviewID:     uuid.NewString(),
}

var (
	kycInfo = npool.KycReq{
		ID:           &kycData.ID,
		AppID:        &kycData.AppID,
		UserID:       &kycData.UserID,
		DocumentType: &kycData.DocumentType,
		IDNumber:     &kycData.IDNumber,
		FrontImg:     &kycData.FrontImg,
		BackImg:      &kycData.BackImg,
		SelfieImg:    &kycData.SelfieImg,
		EntityType:   &kycData.EntityType,
		State:        &kycData.State,
		ReviewID:     &kycData.ReviewID,
	}
)

var info *npool.Kyc

func createKyc(t *testing.T) {
	var err error
	info, err = CreateKyc(context.Background(), &kycInfo)
	if assert.Nil(t, err) {
		kycData.CreatedAt = info.CreatedAt
		kycData.UpdatedAt = info.UpdatedAt
		assert.Equal(t, info, &kycData)
	}
}

func createKycs(t *testing.T) {
	kycDatas := []npool.Kyc{
		{
			ID:           uuid.NewString(),
			AppID:        uuid.NewString(),
			UserID:       uuid.NewString(),
			DocumentType: npool.KycDocumentType_IDCard,
			IDNumber:     uuid.NewString(),
			FrontImg:     uuid.NewString(),
			BackImg:      uuid.NewString(),
			SelfieImg:    uuid.NewString(),
			EntityType:   npool.KycEntityType_Individual,
			State:        npool.KycState_Reviewing,
			ReviewID:     uuid.NewString(),
		},
		{
			ID:           uuid.NewString(),
			AppID:        uuid.NewString(),
			UserID:       uuid.NewString(),
			DocumentType: npool.KycDocumentType_IDCard,
			IDNumber:     uuid.NewString(),
			FrontImg:     uuid.NewString(),
			BackImg:      uuid.NewString(),
			SelfieImg:    uuid.NewString(),
			EntityType:   npool.KycEntityType_Individual,
			State:        npool.KycState_Reviewing,
			ReviewID:     uuid.NewString(),
		},
	}

	Kycs := []*npool.KycReq{}
	for key := range kycDatas {
		Kycs = append(Kycs, &npool.KycReq{
			ID:           &kycDatas[key].ID,
			AppID:        &kycDatas[key].AppID,
			UserID:       &kycDatas[key].UserID,
			DocumentType: &kycDatas[key].DocumentType,
			IDNumber:     &kycDatas[key].IDNumber,
			FrontImg:     &kycDatas[key].FrontImg,
			BackImg:      &kycDatas[key].BackImg,
			SelfieImg:    &kycDatas[key].SelfieImg,
			EntityType:   &kycDatas[key].EntityType,
			State:        &kycDatas[key].State,
			ReviewID:     &kycDatas[key].ReviewID,
		})
	}

	infos, err := CreateKycs(context.Background(), Kycs)
	if assert.Nil(t, err) {
		assert.Equal(t, len(infos), 2)
	}
}

func updateKyc(t *testing.T) {
	var err error
	info, err = UpdateKyc(context.Background(), &kycInfo)
	if assert.Nil(t, err) {
		kycData.UpdatedAt = info.UpdatedAt
		assert.Equal(t, info, &kycData)
	}
}

func getKyc(t *testing.T) {
	var err error
	info, err = GetKyc(context.Background(), info.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, info, &kycData)
	}
}

func getKycs(t *testing.T) {
	infos, total, err := GetKycs(context.Background(),
		&npool.Conds{
			ID: &val.StringVal{
				Value: info.ID,
				Op:    cruder.EQ,
			},
		}, 0, 1)
	if assert.Nil(t, err) {
		assert.Equal(t, total, uint32(1))
		assert.Equal(t, infos[0], &kycData)
	}
}

func getKycOnly(t *testing.T) {
	var err error
	info, err = GetKycOnly(context.Background(),
		&npool.Conds{
			ID: &val.StringVal{
				Value: info.ID,
				Op:    cruder.EQ,
			},
		})
	if assert.Nil(t, err) {
		assert.Equal(t, info, &kycData)
	}
}

func countKycs(t *testing.T) {
	count, err := CountKycs(context.Background(),
		&npool.Conds{
			ID: &val.StringVal{
				Value: info.ID,
				Op:    cruder.EQ,
			},
		},
	)
	if assert.Nil(t, err) {
		assert.Equal(t, count, uint32(1))
	}
}

func existKyc(t *testing.T) {
	exist, err := ExistKyc(context.Background(), info.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, exist, true)
	}
}

func existKycConds(t *testing.T) {
	exist, err := ExistKycConds(context.Background(),
		&npool.Conds{
			ID: &val.StringVal{
				Value: info.ID,
				Op:    cruder.EQ,
			},
		},
	)
	if assert.Nil(t, err) {
		assert.Equal(t, exist, true)
	}
}

func deleteKyc(t *testing.T) {
	info, err := DeleteKyc(context.Background(), info.ID)
	if assert.Nil(t, err) {
		kycData.UpdatedAt = info.UpdatedAt
		assert.Equal(t, info, &kycData)
	}
}

func TestMainOrder(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	gport := config.GetIntValueWithNameSpace("", config.KeyGRPCPort)

	monkey.Patch(grpc2.GetGRPCConn, func(service string, tags ...string) (*grpc.ClientConn, error) {
		return grpc.Dial(fmt.Sprintf("localhost:%v", gport), grpc.WithTransportCredentials(insecure.NewCredentials()))
	})

	t.Run("createKyc", createKyc)
	t.Run("createKycs", createKycs)
	t.Run("getKyc", getKyc)
	t.Run("getKycs", getKycs)
	t.Run("getKycOnly", getKycOnly)
	t.Run("updateKyc", updateKyc)
	t.Run("existKyc", existKyc)
	t.Run("existKycConds", existKycConds)
	t.Run("count", countKycs)
	t.Run("delete", deleteKyc)
}
