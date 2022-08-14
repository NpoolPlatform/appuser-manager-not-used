package appusercontrol

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
	npool "github.com/NpoolPlatform/message/npool/appuser/mgr/v2/appusercontrol"
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

var appUserControlDate = npool.AppUserControl{
	ID:                 uuid.NewString(),
	AppID:              uuid.NewString(),
	UserID:             uuid.NewString(),
	GoogleAuthVerified: false,
}

var (
	appUserControlInfo = npool.AppUserControlReq{
		ID:                 &appUserControlDate.ID,
		AppID:              &appUserControlDate.AppID,
		UserID:             &appUserControlDate.UserID,
		GoogleAuthVerified: &appUserControlDate.GoogleAuthVerified,
	}
)

var info *npool.AppUserControl

func createAppUserControl(t *testing.T) {
	var err error
	info, err = CreateAppUserControl(context.Background(), &appUserControlInfo)
	if assert.Nil(t, err) {
		assert.Equal(t, info, &appUserControlDate)
	}
}

func createAppUserControls(t *testing.T) {
	appUserControlDates := []npool.AppUserControl{
		{
			ID:                 uuid.NewString(),
			AppID:              uuid.NewString(),
			UserID:             uuid.NewString(),
			GoogleAuthVerified: false,
		},
		{
			ID:                 uuid.NewString(),
			AppID:              uuid.NewString(),
			UserID:             uuid.NewString(),
			GoogleAuthVerified: false,
		},
	}

	appUserControls := []*npool.AppUserControlReq{}
	for key := range appUserControlDates {
		appUserControls = append(appUserControls, &npool.AppUserControlReq{
			ID:                 &appUserControlDates[key].ID,
			AppID:              &appUserControlDates[key].AppID,
			UserID:             &appUserControlDates[key].UserID,
			GoogleAuthVerified: &appUserControlDates[key].GoogleAuthVerified,
		})
	}

	infos, err := CreateAppUserControls(context.Background(), appUserControls)
	if assert.Nil(t, err) {
		assert.Equal(t, len(infos), 2)
	}
}

func updateAppUserControl(t *testing.T) {
	var err error
	info, err = UpdateAppUserControl(context.Background(), &appUserControlInfo)
	if assert.Nil(t, err) {
		assert.Equal(t, info, &appUserControlDate)
	}
}

func getAppUserControl(t *testing.T) {
	var err error
	info, err = GetAppUserControl(context.Background(), info.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, info, &appUserControlDate)
	}
}

func getAppUserControls(t *testing.T) {
	infos, total, err := GetAppUserControls(context.Background(),
		&npool.Conds{
			ID: &val.StringVal{
				Value: info.ID,
				Op:    cruder.EQ,
			},
		}, 0, 1)
	if assert.Nil(t, err) {
		assert.Equal(t, total, uint32(1))
		assert.Equal(t, infos[0], &appUserControlDate)
	}
}

func getAppUserControlOnly(t *testing.T) {
	var err error
	info, err = GetAppUserControlOnly(context.Background(),
		&npool.Conds{
			ID: &val.StringVal{
				Value: info.ID,
				Op:    cruder.EQ,
			},
		})
	if assert.Nil(t, err) {
		assert.Equal(t, info, &appUserControlDate)
	}
}

func countAppUserControls(t *testing.T) {
	count, err := CountAppUserControls(context.Background(),
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

func existAppUserControl(t *testing.T) {
	exist, err := ExistAppUserControl(context.Background(), info.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, exist, true)
	}
}

func existAppUserControlConds(t *testing.T) {
	exist, err := ExistAppUserControlConds(context.Background(),
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

func deleteAppUserControl(t *testing.T) {
	info, err := DeleteAppUserControl(context.Background(), info.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, info, &appUserControlDate)
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

	t.Run("createAppUserControl", createAppUserControl)
	t.Run("createAppUserControls", createAppUserControls)
	t.Run("getAppUserControl", getAppUserControl)
	t.Run("getAppUserControls", getAppUserControls)
	t.Run("getAppUserControlOnly", getAppUserControlOnly)
	t.Run("updateAppUserControl", updateAppUserControl)
	t.Run("existAppUserControl", existAppUserControl)
	t.Run("existAppUserControlConds", existAppUserControlConds)
	t.Run("count", countAppUserControls)
	t.Run("delete", deleteAppUserControl)
}
