package subscriber

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
	npool "github.com/NpoolPlatform/message/npool/appuser/mgr/v2/subscriber"
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

var ret = npool.Subscriber{
	ID:           uuid.NewString(),
	AppID:        uuid.NewString(),
	EmailAddress: uuid.NewString(),
}

var (
	req = npool.SubscriberReq{
		ID:           &ret.ID,
		AppID:        &ret.AppID,
		EmailAddress: &ret.EmailAddress,
	}
)

var info *npool.Subscriber

func createSubscriber(t *testing.T) {
	var err error
	info, err = CreateSubscriber(context.Background(), &req)
	if assert.Nil(t, err) {
		ret.CreatedAt = info.CreatedAt
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, info, &ret)
	}
}

func createSubscriberes(t *testing.T) {
	rets := []npool.Subscriber{
		{
			ID:           uuid.New().String(),
			AppID:        uuid.New().String(),
			EmailAddress: uuid.New().String(),
		},
		{
			ID:           uuid.New().String(),
			AppID:        uuid.New().String(),
			EmailAddress: uuid.New().String(),
		},
	}

	apps := []*npool.SubscriberReq{}
	for key := range rets {
		apps = append(apps, &npool.SubscriberReq{
			ID:           &rets[key].ID,
			AppID:        &rets[key].AppID,
			EmailAddress: &rets[key].EmailAddress,
		})
	}

	infos, err := CreateSubscriberes(context.Background(), apps)
	if assert.Nil(t, err) {
		assert.Equal(t, len(infos), 2)
	}
}

func updateSubscriber(t *testing.T) {
	var err error

	registered := true
	req.Registered = &registered
	ret.Registered = registered

	info, err = UpdateSubscriber(context.Background(), &req)
	if assert.Nil(t, err) {
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, info, &ret)
	}
}

func getSubscriber(t *testing.T) {
	var err error
	info, err = GetSubscriber(context.Background(), info.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, info, &ret)
	}
}

func getSubscriberes(t *testing.T) {
	infos, total, err := GetSubscriberes(context.Background(),
		&npool.Conds{
			ID: &val.StringVal{
				Value: info.ID,
				Op:    cruder.EQ,
			},
		}, 0, 1)
	if assert.Nil(t, err) {
		assert.Equal(t, total, uint32(1))
		assert.Equal(t, infos[0], &ret)
	}
}

func getSubscriberOnly(t *testing.T) {
	var err error
	info, err = GetSubscriberOnly(context.Background(),
		&npool.Conds{
			ID: &val.StringVal{
				Value: info.ID,
				Op:    cruder.EQ,
			},
		})
	if assert.Nil(t, err) {
		assert.Equal(t, info, &ret)
	}
}

func countSubscriberes(t *testing.T) {
	count, err := CountSubscriberes(context.Background(),
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

func existSubscriber(t *testing.T) {
	exist, err := ExistSubscriber(context.Background(), info.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, exist, true)
	}
}

func existSubscriberConds(t *testing.T) {
	exist, err := ExistSubscriberConds(context.Background(),
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

func deleteSubscriber(t *testing.T) {
	info, err := DeleteSubscriber(context.Background(), info.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, info, &ret)
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

	t.Run("createSubscriber", createSubscriber)
	t.Run("createSubscriberes", createSubscriberes)
	t.Run("getSubscriber", getSubscriber)
	t.Run("getSubscriberes", getSubscriberes)
	t.Run("getSubscriberOnly", getSubscriberOnly)
	t.Run("updateSubscriber", updateSubscriber)
	t.Run("existSubscriber", existSubscriber)
	t.Run("existSubscriberConds", existSubscriberConds)
	t.Run("count", countSubscriberes)
	t.Run("delete", deleteSubscriber)
}
