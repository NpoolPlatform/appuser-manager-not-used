package banappuser

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

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	val "github.com/NpoolPlatform/message/npool"

	testinit "github.com/NpoolPlatform/appuser-manager/pkg/testinit"
	npool "github.com/NpoolPlatform/message/npool/appuser/mgr/v2/banappuser"
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

var banAppUserDate = npool.BanAppUser{
	ID:      uuid.NewString(),
	AppID:   uuid.NewString(),
	UserID:  uuid.NewString(),
	Message: uuid.NewString(),
}

var (
	banAppUserInfo = npool.BanAppUserReq{
		ID:      &banAppUserDate.ID,
		AppID:   &banAppUserDate.AppID,
		UserID:  &banAppUserDate.UserID,
		Message: &banAppUserDate.Message,
	}
)

var info *npool.BanAppUser

func createBanAppUser(t *testing.T) {
	var err error
	info, err = CreateBanAppUser(context.Background(), &banAppUserInfo)
	if assert.Nil(t, err) {
		assert.Equal(t, info, &banAppUserDate)
	}
}

func createBanAppUsers(t *testing.T) {
	banAppUserDates := []npool.BanAppUser{
		{
			ID:      uuid.NewString(),
			AppID:   uuid.NewString(),
			UserID:  uuid.NewString(),
			Message: uuid.NewString(),
		},
		{
			ID:      uuid.NewString(),
			AppID:   uuid.NewString(),
			UserID:  uuid.NewString(),
			Message: uuid.NewString(),
		},
	}

	banAppUsers := []*npool.BanAppUserReq{}
	for key := range banAppUserDates {
		banAppUsers = append(banAppUsers, &npool.BanAppUserReq{
			ID:      &banAppUserDates[key].ID,
			AppID:   &banAppUserDates[key].AppID,
			UserID:  &banAppUserDates[key].UserID,
			Message: &banAppUserDates[key].Message,
		})
	}

	infos, err := CreateBanAppUsers(context.Background(), banAppUsers)
	if assert.Nil(t, err) {
		assert.Equal(t, len(infos), 2)
	}
}

func updateBanAppUser(t *testing.T) {
	var err error
	info, err = UpdateBanAppUser(context.Background(), &banAppUserInfo)
	if assert.Nil(t, err) {
		assert.Equal(t, info, &banAppUserDate)
	}
}

func getBanAppUser(t *testing.T) {
	var err error
	info, err = GetBanAppUser(context.Background(), info.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, info, &banAppUserDate)
	}
}

func getBanAppUsers(t *testing.T) {
	infos, total, err := GetBanAppUsers(context.Background(),
		&npool.Conds{
			ID: &val.StringVal{
				Value: info.ID,
				Op:    cruder.EQ,
			},
		}, 1, 0)
	if assert.Nil(t, err) {
		assert.Equal(t, total, uint32(1))
		assert.Equal(t, infos[0], &banAppUserDate)
	}
}

func getBanAppUserOnly(t *testing.T) {
	var err error
	info, err = GetBanAppUserOnly(context.Background(),
		&npool.Conds{
			ID: &val.StringVal{
				Value: info.ID,
				Op:    cruder.EQ,
			},
		})
	if assert.Nil(t, err) {
		assert.Equal(t, info, &banAppUserDate)
	}
}

func countBanAppUsers(t *testing.T) {
	count, err := CountBanAppUsers(context.Background(),
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

func existBanAppUser(t *testing.T) {
	exist, err := ExistBanAppUser(context.Background(), info.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, exist, true)
	}
}

func existBanAppUserConds(t *testing.T) {
	exist, err := ExistBanAppUserConds(context.Background(),
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

func deleteBanAppUser(t *testing.T) {
	info, err := DeleteBanAppUser(context.Background(), info.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, info, &banAppUserDate)
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

	t.Run("createBanAppUser", createBanAppUser)
	t.Run("createBanAppUsers", createBanAppUsers)
	t.Run("getBanAppUser", getBanAppUser)
	t.Run("getBanAppUsers", getBanAppUsers)
	t.Run("getBanAppUserOnly", getBanAppUserOnly)
	t.Run("updateBanAppUser", updateBanAppUser)
	t.Run("existBanAppUser", existBanAppUser)
	t.Run("existBanAppUserConds", existBanAppUserConds)
	t.Run("count", countBanAppUsers)
	t.Run("delete", deleteBanAppUser)
}
