package app

import (
	"context"
	"fmt"
	"google.golang.org/grpc/credentials/insecure"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/go-service-framework/pkg/config"

	"bou.ke/monkey"
	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"
	"google.golang.org/grpc"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	val "github.com/NpoolPlatform/message/npool"

	testinit "github.com/NpoolPlatform/appuser-manager/pkg/testinit"
	npool "github.com/NpoolPlatform/message/npool/appuser/mgr/v2/app"
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

var appDate = npool.App{
	ID:          uuid.NewString(),
	CreatedBy:   uuid.NewString(),
	Name:        uuid.New().String(),
	Description: uuid.New().String(),
	Logo:        uuid.New().String(),
}

var (
	appInfo = npool.AppReq{
		ID:          &appDate.ID,
		CreatedBy:   &appDate.CreatedBy,
		Name:        &appDate.Name,
		Description: &appDate.Description,
		Logo:        &appDate.Logo,
	}
)

var info *npool.App

func createApp(t *testing.T) {
	var err error
	info, err = CreateApp(context.Background(), &appInfo)
	if assert.Nil(t, err) {
		appDate.CreatedAt = info.CreatedAt
		assert.Equal(t, info, &appDate)
	}
}

func createApps(t *testing.T) {
	appDates := []npool.App{
		{
			ID:          uuid.New().String(),
			CreatedBy:   uuid.New().String(),
			Name:        uuid.New().String(),
			Description: uuid.New().String(),
			Logo:        uuid.New().String(),
		},
		{
			ID:          uuid.New().String(),
			CreatedBy:   uuid.New().String(),
			Name:        uuid.New().String(),
			Description: uuid.New().String(),
			Logo:        uuid.New().String(),
		},
	}

	apps := []*npool.AppReq{}
	for key := range appDates {
		fmt.Println(appDates[key].Name)
		apps = append(apps, &npool.AppReq{
			ID:          &appDates[key].ID,
			CreatedBy:   &appDates[key].CreatedBy,
			Name:        &appDates[key].Name,
			Logo:        &appDates[key].Logo,
			Description: &appDates[key].Description,
			CreatedAt:   &appDates[key].CreatedAt,
		})
	}

	infos, err := CreateApps(context.Background(), apps)
	if assert.Nil(t, err) {
		assert.Equal(t, len(infos), 2)
	}
}

func updateApp(t *testing.T) {
	var err error
	info, err = UpdateApp(context.Background(), &appInfo)
	if assert.Nil(t, err) {
		assert.Equal(t, info, &appDate)
	}
}

func getApp(t *testing.T) {
	var err error
	info, err = GetApp(context.Background(), info.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, info, &appDate)
	}
}

func getApps(t *testing.T) {
	infos, total, err := GetApps(context.Background(),
		&npool.Conds{
			ID: &val.StringVal{
				Value: info.ID,
				Op:    cruder.EQ,
			},
		}, 1, 0)
	if assert.Nil(t, err) {
		assert.Equal(t, total, uint32(1))
		assert.Equal(t, infos[0], &appDate)
	}
}

func getAppOnly(t *testing.T) {
	var err error
	info, err = GetAppOnly(context.Background(),
		&npool.Conds{
			ID: &val.StringVal{
				Value: info.ID,
				Op:    cruder.EQ,
			},
		})
	if assert.Nil(t, err) {
		assert.Equal(t, info, &appDate)
	}
}

func countApps(t *testing.T) {
	count, err := CountApps(context.Background(),
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

func existApp(t *testing.T) {
	exist, err := ExistApp(context.Background(), info.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, exist, true)
	}
}

func existAppConds(t *testing.T) {
	exist, err := ExistAppConds(context.Background(),
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

func deleteApp(t *testing.T) {
	info, err := DeleteApp(context.Background(), info.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, info, &appDate)
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

	t.Run("createApp", createApp)
	t.Run("createApps", createApps)
	t.Run("getApp", getApp)
	t.Run("getApps", getApps)
	t.Run("getAppOnly", getAppOnly)
	t.Run("updateApp", updateApp)
	t.Run("existApp", existApp)
	t.Run("existAppConds", existAppConds)
	t.Run("count", countApps)
	t.Run("delete", deleteApp)
}
