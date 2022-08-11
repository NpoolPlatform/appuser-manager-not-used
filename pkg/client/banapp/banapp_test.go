package banapp

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
	npool "github.com/NpoolPlatform/message/npool/appuser/mgr/v2/banapp"
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

var BanAppDate = npool.BanApp{
	ID:      uuid.NewString(),
	AppID:   uuid.NewString(),
	Message: uuid.NewString(),
}

var (
	BanAppInfo = npool.BanAppReq{
		ID:      &BanAppDate.ID,
		AppID:   &BanAppDate.AppID,
		Message: &BanAppDate.Message,
	}
)

var info *npool.BanApp

func createBanApp(t *testing.T) {
	var err error
	info, err = CreateBanApp(context.Background(), &BanAppInfo)
	if assert.Nil(t, err) {
		assert.Equal(t, info, &BanAppDate)
	}
}

func createBanApps(t *testing.T) {
	BanAppDates := []npool.BanApp{
		{
			ID:      uuid.NewString(),
			AppID:   uuid.NewString(),
			Message: uuid.NewString(),
		},
		{
			ID:      uuid.NewString(),
			AppID:   uuid.NewString(),
			Message: uuid.NewString(),
		},
	}

	BanApps := []*npool.BanAppReq{}
	for key := range BanAppDates {
		BanApps = append(BanApps, &npool.BanAppReq{
			ID:      &BanAppDates[key].ID,
			AppID:   &BanAppDates[key].AppID,
			Message: &BanAppDates[key].Message,
		})
	}

	infos, err := CreateBanApps(context.Background(), BanApps)
	if assert.Nil(t, err) {
		assert.Equal(t, len(infos), 2)
	}
}

func updateBanApp(t *testing.T) {
	var err error
	info, err = UpdateBanApp(context.Background(), &BanAppInfo)
	if assert.Nil(t, err) {
		assert.Equal(t, info, &BanAppDate)
	}
}

func getBanApp(t *testing.T) {
	var err error
	info, err = GetBanApp(context.Background(), info.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, info, &BanAppDate)
	}
}

func getBanApps(t *testing.T) {
	infos, total, err := GetBanApps(context.Background(),
		&npool.Conds{
			ID: &val.StringVal{
				Value: info.ID,
				Op:    cruder.EQ,
			},
		}, 0, 1)
	if assert.Nil(t, err) {
		assert.Equal(t, total, uint32(1))
		assert.Equal(t, infos[0], &BanAppDate)
	}
}

func getBanAppOnly(t *testing.T) {
	var err error
	info, err = GetBanAppOnly(context.Background(),
		&npool.Conds{
			ID: &val.StringVal{
				Value: info.ID,
				Op:    cruder.EQ,
			},
		})
	if assert.Nil(t, err) {
		assert.Equal(t, info, &BanAppDate)
	}
}

func countBanApps(t *testing.T) {
	count, err := CountBanApps(context.Background(),
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

func existBanApp(t *testing.T) {
	exist, err := ExistBanApp(context.Background(), info.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, exist, true)
	}
}

func existBanAppConds(t *testing.T) {
	exist, err := ExistBanAppConds(context.Background(),
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

func deleteBanApp(t *testing.T) {
	info, err := DeleteBanApp(context.Background(), info.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, info, &BanAppDate)
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

	t.Run("createBanApp", createBanApp)
	t.Run("createBanApps", createBanApps)
	t.Run("getBanApp", getBanApp)
	t.Run("getBanApps", getBanApps)
	t.Run("getBanAppOnly", getBanAppOnly)
	t.Run("updateBanApp", updateBanApp)
	t.Run("existBanApp", existBanApp)
	t.Run("existBanAppConds", existBanAppConds)
	t.Run("count", countBanApps)
	t.Run("delete", deleteBanApp)
}
