package approleuser

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
	npool "github.com/NpoolPlatform/message/npool/appuser/mgr/v2/approleuser"
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

var appRoleDate = npool.AppRoleUser{
	ID:     uuid.NewString(),
	AppID:  uuid.NewString(),
	RoleID: uuid.NewString(),
	UserID: uuid.NewString(),
}

var (
	appRoleInfo = npool.AppRoleUserReq{
		ID:     &appRoleDate.ID,
		AppID:  &appRoleDate.AppID,
		RoleID: &appRoleDate.RoleID,
		UserID: &appRoleDate.UserID,
	}
)

var info *npool.AppRoleUser

func createAppRoleUser(t *testing.T) {
	var err error
	info, err = CreateAppRoleUser(context.Background(), &appRoleInfo)
	if assert.Nil(t, err) {
		assert.Equal(t, info, &appRoleDate)
	}
}

func createAppRoleUsers(t *testing.T) {
	appRoleDates := []npool.AppRoleUser{
		{
			ID:     uuid.NewString(),
			AppID:  uuid.NewString(),
			RoleID: uuid.NewString(),
			UserID: uuid.NewString(),
		},
		{
			ID:     uuid.NewString(),
			AppID:  uuid.NewString(),
			RoleID: uuid.NewString(),
			UserID: uuid.NewString(),
		},
	}

	appRoleUsers := []*npool.AppRoleUserReq{}
	for key := range appRoleDates {
		appRoleUsers = append(appRoleUsers, &npool.AppRoleUserReq{
			ID:     &appRoleDates[key].ID,
			AppID:  &appRoleDates[key].AppID,
			RoleID: &appRoleDates[key].RoleID,
			UserID: &appRoleDates[key].UserID,
		})
	}

	infos, err := CreateAppRoleUsers(context.Background(), appRoleUsers)
	if assert.Nil(t, err) {
		assert.Equal(t, len(infos), 2)
	}
}

func updateAppRoleUser(t *testing.T) {
	var err error
	info, err = UpdateAppRoleUser(context.Background(), &appRoleInfo)
	if assert.Nil(t, err) {
		assert.Equal(t, info, &appRoleDate)
	}
}

func getAppRoleUser(t *testing.T) {
	var err error
	info, err = GetAppRoleUser(context.Background(), info.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, info, &appRoleDate)
	}
}

func getAppRoleUsers(t *testing.T) {
	infos, total, err := GetAppRoleUsers(context.Background(),
		&npool.Conds{
			ID: &val.StringVal{
				Value: info.ID,
				Op:    cruder.EQ,
			},
		}, 1, 0)
	if assert.Nil(t, err) {
		assert.Equal(t, total, uint32(1))
		assert.Equal(t, infos[0], &appRoleDate)
	}
}

func getAppRoleUserOnly(t *testing.T) {
	var err error
	info, err = GetAppRoleUserOnly(context.Background(),
		&npool.Conds{
			ID: &val.StringVal{
				Value: info.ID,
				Op:    cruder.EQ,
			},
		})
	if assert.Nil(t, err) {
		assert.Equal(t, info, &appRoleDate)
	}
}

func countAppRoleUsers(t *testing.T) {
	count, err := CountAppRoleUsers(context.Background(),
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

func existAppRoleUser(t *testing.T) {
	exist, err := ExistAppRoleUser(context.Background(), info.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, exist, true)
	}
}

func existAppRoleUserConds(t *testing.T) {
	exist, err := ExistAppRoleUserConds(context.Background(),
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

func deleteAppRoleUser(t *testing.T) {
	info, err := DeleteAppRoleUser(context.Background(), info.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, info, &appRoleDate)
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

	t.Run("createAppRoleUser", createAppRoleUser)
	t.Run("createAppRoleUsers", createAppRoleUsers)
	t.Run("getAppRoleUser", getAppRoleUser)
	t.Run("getAppRoleUsers", getAppRoleUsers)
	t.Run("getAppRoleUserOnly", getAppRoleUserOnly)
	t.Run("updateAppRoleUser", updateAppRoleUser)
	t.Run("existAppRoleUser", existAppRoleUser)
	t.Run("existAppRoleUserConds", existAppRoleUserConds)
	t.Run("count", countAppRoleUsers)
	t.Run("delete", deleteAppRoleUser)
}
