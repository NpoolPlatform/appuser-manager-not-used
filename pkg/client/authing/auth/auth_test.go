package auth

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
	npool "github.com/NpoolPlatform/message/npool/appuser/mgr/v2/authing/auth"
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

var authDate = npool.Auth{
	ID:       uuid.New().String(),
	AppID:    uuid.New().String(),
	RoleID:   uuid.New().String(),
	UserID:   uuid.New().String(),
	Resource: uuid.New().String(),
	Method:   uuid.New().String(),
}

var (
	authInfo = npool.AuthReq{
		ID:       &authDate.ID,
		AppID:    &authDate.AppID,
		RoleID:   &authDate.RoleID,
		UserID:   &authDate.UserID,
		Resource: &authDate.Resource,
		Method:   &authDate.Method,
	}
)

var info *npool.Auth

func createAuth(t *testing.T) {
	var err error
	info, err = CreateAuth(context.Background(), &authInfo)
	if assert.Nil(t, err) {
		authDate.CreatedAt = info.CreatedAt
		assert.Equal(t, info, &authDate)
	}
}

func createAuths(t *testing.T) {
	authDates := []npool.Auth{
		{
			ID:       uuid.New().String(),
			AppID:    uuid.New().String(),
			RoleID:   uuid.New().String(),
			UserID:   uuid.New().String(),
			Resource: uuid.New().String(),
			Method:   uuid.New().String(),
		},
		{
			ID:       uuid.New().String(),
			AppID:    uuid.New().String(),
			RoleID:   uuid.New().String(),
			UserID:   uuid.New().String(),
			Resource: uuid.New().String(),
			Method:   uuid.New().String(),
		},
	}

	auths := []*npool.AuthReq{}
	for key := range authDates {
		auths = append(auths, &npool.AuthReq{
			ID:       &authDates[key].ID,
			AppID:    &authDates[key].AppID,
			RoleID:   &authDates[key].RoleID,
			UserID:   &authDates[key].UserID,
			Resource: &authDates[key].Resource,
			Method:   &authDates[key].Method,
		})
	}

	infos, err := CreateAuths(context.Background(), auths)
	if assert.Nil(t, err) {
		assert.Equal(t, len(infos), 2)
	}
}

func updateAuth(t *testing.T) {
	var err error
	info, err = UpdateAuth(context.Background(), &authInfo)
	if assert.Nil(t, err) {
		assert.Equal(t, info, &authDate)
	}
}

func getAuth(t *testing.T) {
	var err error
	info, err = GetAuth(context.Background(), info.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, info, &authDate)
	}
}

func getAuths(t *testing.T) {
	infos, total, err := GetAuths(context.Background(),
		&npool.Conds{
			ID: &val.StringVal{
				Value: info.ID,
				Op:    cruder.EQ,
			},
		}, 0, 1)
	if assert.Nil(t, err) {
		assert.Equal(t, total, uint32(1))
		assert.Equal(t, infos[0], &authDate)
	}
}

func getAuthOnly(t *testing.T) {
	var err error
	info, err = GetAuthOnly(context.Background(),
		&npool.Conds{
			ID: &val.StringVal{
				Value: info.ID,
				Op:    cruder.EQ,
			},
		})
	if assert.Nil(t, err) {
		assert.Equal(t, info, &authDate)
	}
}

func countAuths(t *testing.T) {
	count, err := CountAuths(context.Background(),
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

func existAuth(t *testing.T) {
	exist, err := ExistAuth(context.Background(), info.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, exist, true)
	}
}

func existAuthConds(t *testing.T) {
	exist, err := ExistAuthConds(context.Background(),
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

func deleteAuth(t *testing.T) {
	info, err := DeleteAuth(context.Background(), info.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, info, &authDate)
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

	t.Run("createAuth", createAuth)
	t.Run("createAuths", createAuths)
	t.Run("getAuth", getAuth)
	t.Run("getAuths", getAuths)
	t.Run("getAuthOnly", getAuthOnly)
	t.Run("updateAuth", updateAuth)
	t.Run("existAuth", existAuth)
	t.Run("existAuthConds", existAuthConds)
	t.Run("count", countAuths)
	t.Run("delete", deleteAuth)
}
