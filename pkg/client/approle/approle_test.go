package approle

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"google.golang.org/grpc/credentials/insecure"

	"github.com/NpoolPlatform/go-service-framework/pkg/config"

	"bou.ke/monkey"
	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"
	"google.golang.org/grpc"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	val "github.com/NpoolPlatform/message/npool"

	testinit "github.com/NpoolPlatform/appuser-manager/pkg/testinit"
	npool "github.com/NpoolPlatform/message/npool/appuser/mgr/v2/approle"
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

var appRoleDate = npool.AppRole{
	ID:          uuid.NewString(),
	AppID:       uuid.NewString(),
	CreatedBy:   uuid.NewString(),
	Role:        uuid.NewString(),
	Description: uuid.NewString(),
	Default:     true,
}

var (
	appRoleInfo = npool.AppRoleReq{
		ID:          &appRoleDate.ID,
		AppID:       &appRoleDate.AppID,
		CreatedBy:   &appRoleDate.CreatedBy,
		Role:        &appRoleDate.Role,
		Description: &appRoleDate.Description,
		Default:     &appRoleDate.Default,
	}
)

var info *npool.AppRole

func createAppRole(t *testing.T) {
	var err error
	info, err = CreateAppRole(context.Background(), &appRoleInfo)
	if assert.Nil(t, err) {
		assert.Equal(t, info, &appRoleDate)
	}
}

func createAppRoles(t *testing.T) {
	appRoleDates := []npool.AppRole{
		{
			ID:          uuid.NewString(),
			AppID:       uuid.NewString(),
			CreatedBy:   uuid.NewString(),
			Role:        uuid.NewString(),
			Description: uuid.NewString(),
			Default:     true,
		},
		{
			ID:          uuid.NewString(),
			AppID:       uuid.NewString(),
			CreatedBy:   uuid.NewString(),
			Role:        uuid.NewString(),
			Description: uuid.NewString(),
			Default:     true,
		},
	}

	appRoles := []*npool.AppRoleReq{}
	for key := range appRoleDates {
		appRoles = append(appRoles, &npool.AppRoleReq{
			ID:          &appRoleDates[key].ID,
			AppID:       &appRoleDates[key].AppID,
			CreatedBy:   &appRoleDates[key].CreatedBy,
			Role:        &appRoleDates[key].Role,
			Description: &appRoleDates[key].Description,
			Default:     &appRoleDates[key].Default,
		})
	}

	infos, err := CreateAppRoles(context.Background(), appRoles)
	if assert.Nil(t, err) {
		assert.Equal(t, len(infos), 2)
	}
}

func updateAppRole(t *testing.T) {
	var err error
	info, err = UpdateAppRole(context.Background(), &appRoleInfo)
	if assert.Nil(t, err) {
		assert.Equal(t, info, &appRoleDate)
	}
}

func getAppRole(t *testing.T) {
	var err error
	info, err = GetAppRole(context.Background(), info.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, info, &appRoleDate)
	}
}

func getAppRoles(t *testing.T) {
	infos, total, err := GetAppRoles(context.Background(),
		&npool.Conds{
			ID: &val.StringVal{
				Value: info.ID,
				Op:    cruder.EQ,
			},
		}, 0, 1)
	if assert.Nil(t, err) {
		assert.Equal(t, total, uint32(1))
		assert.Equal(t, infos[0], &appRoleDate)
	}
}

func getAppRoleOnly(t *testing.T) {
	var err error
	info, err = GetAppRoleOnly(context.Background(),
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

func countAppRoles(t *testing.T) {
	count, err := CountAppRoles(context.Background(),
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

func existAppRole(t *testing.T) {
	exist, err := ExistAppRole(context.Background(), info.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, exist, true)
	}
}

func existAppRoleConds(t *testing.T) {
	exist, err := ExistAppRoleConds(context.Background(),
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

func deleteAppRole(t *testing.T) {
	info, err := DeleteAppRole(context.Background(), info.ID)
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

	t.Run("createAppRole", createAppRole)
	t.Run("createAppRoles", createAppRoles)
	t.Run("getAppRole", getAppRole)
	t.Run("getAppRoles", getAppRoles)
	t.Run("getAppRoleOnly", getAppRoleOnly)
	t.Run("updateAppRole", updateAppRole)
	t.Run("existAppRole", existAppRole)
	t.Run("existAppRoleConds", existAppRoleConds)
	t.Run("count", countAppRoles)
	t.Run("delete", deleteAppRole)
}
