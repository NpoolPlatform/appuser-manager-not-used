package appusersecret

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
	npool "github.com/NpoolPlatform/message/npool/appuser/mgr/v2/appusersecret"
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

var appUserSecretDate = npool.AppUserSecret{
	ID:           uuid.NewString(),
	AppID:        uuid.NewString(),
	UserID:       uuid.NewString(),
	PasswordHash: uuid.NewString(),
	Salt:         uuid.NewString(),
	GoogleSecret: uuid.NewString(),
}

var (
	appUserSecretInfo = npool.AppUserSecretReq{
		ID:           &appUserSecretDate.ID,
		AppID:        &appUserSecretDate.AppID,
		UserID:       &appUserSecretDate.UserID,
		PasswordHash: &appUserSecretDate.PasswordHash,
		Salt:         &appUserSecretDate.Salt,
		GoogleSecret: &appUserSecretDate.GoogleSecret,
	}
)

var info *npool.AppUserSecret

func createAppUserSecret(t *testing.T) {
	var err error
	info, err = CreateAppUserSecret(context.Background(), &appUserSecretInfo)
	if assert.Nil(t, err) {
		assert.Equal(t, info, &appUserSecretDate)
	}
}

func createAppUserSecrets(t *testing.T) {
	appUserSecretDates := []npool.AppUserSecret{
		{
			ID:           uuid.NewString(),
			AppID:        uuid.NewString(),
			UserID:       uuid.NewString(),
			PasswordHash: uuid.NewString(),
			Salt:         uuid.NewString(),
			GoogleSecret: uuid.NewString(),
		},
		{
			ID:           uuid.NewString(),
			AppID:        uuid.NewString(),
			UserID:       uuid.NewString(),
			PasswordHash: uuid.NewString(),
			Salt:         uuid.NewString(),
			GoogleSecret: uuid.NewString(),
		},
	}

	appUserSecrets := []*npool.AppUserSecretReq{}
	for key := range appUserSecretDates {
		appUserSecrets = append(appUserSecrets, &npool.AppUserSecretReq{
			ID:           &appUserSecretDates[key].ID,
			AppID:        &appUserSecretDates[key].AppID,
			UserID:       &appUserSecretDates[key].UserID,
			PasswordHash: &appUserSecretDates[key].PasswordHash,
			Salt:         &appUserSecretDates[key].Salt,
			GoogleSecret: &appUserSecretDates[key].GoogleSecret,
		})
	}

	infos, err := CreateAppUserSecrets(context.Background(), appUserSecrets)
	if assert.Nil(t, err) {
		assert.Equal(t, len(infos), 2)
	}
}

func updateAppUserSecret(t *testing.T) {
	var err error
	info, err = UpdateAppUserSecret(context.Background(), &appUserSecretInfo)
	if assert.Nil(t, err) {
		assert.Equal(t, info, &appUserSecretDate)
	}
}

func getAppUserSecret(t *testing.T) {
	var err error
	info, err = GetAppUserSecret(context.Background(), info.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, info, &appUserSecretDate)
	}
}

func getAppUserSecrets(t *testing.T) {
	infos, total, err := GetAppUserSecrets(context.Background(),
		&npool.Conds{
			ID: &val.StringVal{
				Value: info.ID,
				Op:    cruder.EQ,
			},
		}, 1, 0)
	if assert.Nil(t, err) {
		assert.Equal(t, total, uint32(1))
		assert.Equal(t, infos[0], &appUserSecretDate)
	}
}

func getAppUserSecretOnly(t *testing.T) {
	var err error
	info, err = GetAppUserSecretOnly(context.Background(),
		&npool.Conds{
			ID: &val.StringVal{
				Value: info.ID,
				Op:    cruder.EQ,
			},
		})
	if assert.Nil(t, err) {
		assert.Equal(t, info, &appUserSecretDate)
	}
}

func countAppUserSecrets(t *testing.T) {
	count, err := CountAppUserSecrets(context.Background(),
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

func existAppUserSecret(t *testing.T) {
	exist, err := ExistAppUserSecret(context.Background(), info.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, exist, true)
	}
}

func existAppUserSecretConds(t *testing.T) {
	exist, err := ExistAppUserSecretConds(context.Background(),
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

func deleteAppUserSecret(t *testing.T) {
	info, err := DeleteAppUserSecret(context.Background(), info.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, info, &appUserSecretDate)
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

	t.Run("createAppUserSecret", createAppUserSecret)
	t.Run("createAppUserSecrets", createAppUserSecrets)
	t.Run("getAppUserSecret", getAppUserSecret)
	t.Run("getAppUserSecrets", getAppUserSecrets)
	t.Run("getAppUserSecretOnly", getAppUserSecretOnly)
	t.Run("updateAppUserSecret", updateAppUserSecret)
	t.Run("existAppUserSecret", existAppUserSecret)
	t.Run("existAppUserSecretConds", existAppUserSecretConds)
	t.Run("count", countAppUserSecrets)
	t.Run("delete", deleteAppUserSecret)
}
