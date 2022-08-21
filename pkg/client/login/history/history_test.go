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
	npool "github.com/NpoolPlatform/message/npool/appuser/mgr/v2/login/history"
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
	ID:        uuid.New().String(),
	AppID:     uuid.New().String(),
	UserID:    uuid.New().String(),
	ClientIP:  uuid.New().String(),
	UserAgent: uuid.New().String(),
}

var (
	historyInfo = npool.HistoryReq{
		ID:        &historyDate.ID,
		AppID:     &historyDate.AppID,
		UserID:    &historyDate.UserID,
		ClientIP:  &historyDate.ClientIP,
		UserAgent: &historyDate.UserAgent,
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

func createHistories(t *testing.T) {
	historyDates := []npool.History{
		{
			ID:        uuid.New().String(),
			AppID:     uuid.NewString(),
			UserID:    uuid.NewString(),
			ClientIP:  uuid.New().String(),
			UserAgent: uuid.New().String(),
		},
		{
			ID:        uuid.New().String(),
			AppID:     uuid.NewString(),
			UserID:    uuid.NewString(),
			ClientIP:  uuid.New().String(),
			UserAgent: uuid.New().String(),
		},
	}

	histories := []*npool.HistoryReq{}
	for key := range historyDates {
		histories = append(histories, &npool.HistoryReq{
			ID:        &historyDates[key].ID,
			AppID:     &historyDates[key].AppID,
			UserID:    &historyDates[key].UserID,
			ClientIP:  &historyDates[key].ClientIP,
			UserAgent: &historyDates[key].UserAgent,
		})
	}

	infos, err := CreateHistories(context.Background(), histories)
	if assert.Nil(t, err) {
		assert.Equal(t, len(infos), 2)
	}
}

func updateHistory(t *testing.T) {
	var err error

	location := "HI9hkjjhjk"
	historyInfo.Location = &location
	historyDate.Location = location

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

func getHistories(t *testing.T) {
	infos, total, err := GetHistories(context.Background(),
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

func countHistories(t *testing.T) {
	count, err := CountHistories(context.Background(),
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
	t.Run("createHistories", createHistories)
	t.Run("getHistory", getHistory)
	t.Run("getHistories", getHistories)
	t.Run("getHistoryOnly", getHistoryOnly)
	t.Run("updateHistory", updateHistory)
	t.Run("existHistory", existHistory)
	t.Run("existHistoryConds", existHistoryConds)
	t.Run("count", countHistories)
	t.Run("delete", deleteHistory)
}
