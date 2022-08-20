package history

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
	npool "github.com/NpoolPlatform/message/npool/appuser/mgr/v2/authing/history"
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

var historyDate = npool.History{
	ID:       uuid.New().String(),
	AppID:    uuid.New().String(),
	UserID:   uuid.New().String(),
	Resource: uuid.New().String(),
	Method:   uuid.New().String(),
	Allowed:  true,
}

var (
	historyInfo = npool.HistoryReq{
		ID:       &historyDate.ID,
		AppID:    &historyDate.AppID,
		UserID:   &historyDate.UserID,
		Resource: &historyDate.Resource,
		Method:   &historyDate.Method,
		Allowed:  &historyDate.Allowed,
	}
)

var info *npool.History

func createHistory(t *testing.T) {
	var err error
	info, err = CreateHistory(context.Background(), &historyInfo)
	if assert.Nil(t, err) {
		historyDate.CreatedAt = info.CreatedAt
		assert.Equal(t, info, &historyDate)
	}
}

func createHistorys(t *testing.T) {
	historyDates := []npool.History{
		{
			ID:       uuid.New().String(),
			AppID:    uuid.NewString(),
			UserID:   uuid.NewString(),
			Resource: uuid.New().String(),
			Method:   uuid.New().String(),
			Allowed:  true,
		},
		{
			ID:       uuid.New().String(),
			AppID:    uuid.NewString(),
			UserID:   uuid.NewString(),
			Resource: uuid.New().String(),
			Method:   uuid.New().String(),
			Allowed:  true,
		},
	}

	histories := []*npool.HistoryReq{}
	for key := range historyDates {
		histories = append(histories, &npool.HistoryReq{
			ID:       &historyDates[key].ID,
			AppID:    &historyDates[key].AppID,
			UserID:   &historyDates[key].UserID,
			Resource: &historyDates[key].Resource,
			Method:   &historyDates[key].Method,
			Allowed:  &historyDates[key].Allowed,
		})
	}

	infos, err := CreateHistorys(context.Background(), histories)
	if assert.Nil(t, err) {
		assert.Equal(t, len(infos), 2)
	}
}

func updateHistory(t *testing.T) {
	var err error
	info, err = UpdateHistory(context.Background(), &historyInfo)
	if assert.Nil(t, err) {
		assert.Equal(t, info, &historyDate)
	}
}

func getHistory(t *testing.T) {
	var err error
	info, err = GetHistory(context.Background(), info.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, info, &historyDate)
	}
}

func getHistorys(t *testing.T) {
	infos, total, err := GetHistorys(context.Background(),
		&npool.Conds{
			ID: &val.StringVal{
				Value: info.ID,
				Op:    cruder.EQ,
			},
		}, 0, 1)
	if assert.Nil(t, err) {
		assert.Equal(t, total, uint32(1))
		assert.Equal(t, infos[0], &historyDate)
	}
}

func getHistoryOnly(t *testing.T) {
	var err error
	info, err = GetHistoryOnly(context.Background(),
		&npool.Conds{
			ID: &val.StringVal{
				Value: info.ID,
				Op:    cruder.EQ,
			},
		})
	if assert.Nil(t, err) {
		assert.Equal(t, info, &historyDate)
	}
}

func countHistorys(t *testing.T) {
	count, err := CountHistorys(context.Background(),
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

func existHistory(t *testing.T) {
	exist, err := ExistHistory(context.Background(), info.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, exist, true)
	}
}

func existHistoryConds(t *testing.T) {
	exist, err := ExistHistoryConds(context.Background(),
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

func deleteHistory(t *testing.T) {
	info, err := DeleteHistory(context.Background(), info.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, info, &historyDate)
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

	t.Run("createHistory", createHistory)
	t.Run("createHistorys", createHistorys)
	t.Run("getHistory", getHistory)
	t.Run("getHistorys", getHistorys)
	t.Run("getHistoryOnly", getHistoryOnly)
	t.Run("updateHistory", updateHistory)
	t.Run("existHistory", existHistory)
	t.Run("existHistoryConds", existHistoryConds)
	t.Run("count", countHistorys)
	t.Run("delete", deleteHistory)
}
