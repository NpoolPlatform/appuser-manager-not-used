package appuser

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
	npool "github.com/NpoolPlatform/message/npool/appuser/mgr/v2/appuser"
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

var appDate = npool.AppUser{
	ID:            uuid.NewString(),
	AppID:         uuid.NewString(),
	PhoneNo:       uuid.NewString(),
	EmailAddress:  uuid.NewString(),
	ImportFromApp: uuid.NewString(),
}

var (
	appInfo = npool.AppUserReq{
		ID:            &appDate.ID,
		AppID:         &appDate.AppID,
		PhoneNo:       &appDate.PhoneNo,
		EmailAddress:  &appDate.EmailAddress,
		ImportFromApp: &appDate.ImportFromApp,
	}
)

var info *npool.AppUser

func createAppUser(t *testing.T) {
	var err error
	info, err = CreateAppUser(context.Background(), &appInfo)
	if assert.Nil(t, err) {
		assert.Equal(t, info, &appDate)
	}
}

func createAppUsers(t *testing.T) {
	appDates := []npool.AppUser{
		{
			ID:            uuid.NewString(),
			AppID:         uuid.NewString(),
			PhoneNo:       uuid.NewString(),
			EmailAddress:  uuid.NewString(),
			ImportFromApp: uuid.NewString(),
		},
		{
			ID:            uuid.NewString(),
			AppID:         uuid.NewString(),
			PhoneNo:       uuid.NewString(),
			EmailAddress:  uuid.NewString(),
			ImportFromApp: uuid.NewString(),
		},
	}

	appUsers := []*npool.AppUserReq{}
	for key := range appDates {
		appUsers = append(appUsers, &npool.AppUserReq{
			ID:            &appDates[key].ID,
			AppID:         &appDates[key].AppID,
			PhoneNo:       &appDates[key].PhoneNo,
			EmailAddress:  &appDates[key].EmailAddress,
			ImportFromApp: &appDates[key].ImportFromApp,
		})
	}

	infos, err := CreateAppUsers(context.Background(), appUsers)
	if assert.Nil(t, err) {
		assert.Equal(t, len(infos), 2)
	}
}

func updateAppUser(t *testing.T) {
	var err error
	info, err = UpdateAppUser(context.Background(), &appInfo)
	if assert.Nil(t, err) {
		assert.Equal(t, info, &appDate)
	}
}

func getAppUser(t *testing.T) {
	var err error
	info, err = GetAppUser(context.Background(), info.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, info, &appDate)
	}
}

func getAppUsers(t *testing.T) {
	infos, total, err := GetAppUsers(context.Background(),
		&npool.Conds{
			ID: &val.StringVal{
				Value: info.ID,
				Op:    cruder.EQ,
			},
		}, 0, 1)
	if assert.Nil(t, err) {
		assert.Equal(t, total, uint32(1))
		assert.Equal(t, infos[0], &appDate)
	}
}

func getAppUserOnly(t *testing.T) {
	var err error
	info, err = GetAppUserOnly(context.Background(),
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

func countAppUsers(t *testing.T) {
	count, err := CountAppUsers(context.Background(),
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

func existAppUser(t *testing.T) {
	exist, err := ExistAppUser(context.Background(), info.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, exist, true)
	}
}

func existAppUserConds(t *testing.T) {
	exist, err := ExistAppUserConds(context.Background(),
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

func deleteAppUser(t *testing.T) {
	info, err := DeleteAppUser(context.Background(), info.ID)
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

	t.Run("createAppUser", createAppUser)
	t.Run("createAppUsers", createAppUsers)
	t.Run("getAppUser", getAppUser)
	t.Run("getAppUsers", getAppUsers)
	t.Run("getAppUserOnly", getAppUserOnly)
	t.Run("updateAppUser", updateAppUser)
	t.Run("existAppUser", existAppUser)
	t.Run("existAppUserConds", existAppUserConds)
	t.Run("count", countAppUsers)
	t.Run("delete", deleteAppUser)
}
