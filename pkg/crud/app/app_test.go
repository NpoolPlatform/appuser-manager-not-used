package app

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	testinit "github.com/NpoolPlatform/appuser-manager/pkg/testinit" //nolint
	npool "github.com/NpoolPlatform/message/npool/appusermgr"

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

func assertApp(t *testing.T, actual, expected *npool.App) {
	assert.Equal(t, actual.CreatedBy, expected.CreatedBy)
	assert.Equal(t, actual.Name, expected.Name)
	assert.Equal(t, actual.Logo, expected.Logo)
	assert.Equal(t, actual.Description, expected.Description)
}

func TestCRUD(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	app := npool.App{
		CreatedBy:   uuid.New().String(),
		Name:        fmt.Sprintf("App-%v", uuid.New().String()),
		Logo:        fmt.Sprintf("App-%v", uuid.New().String()),
		Description: fmt.Sprintf("App-%v", uuid.New().String()),
	}

	resp, err := Create(context.Background(), &npool.CreateAppRequest{
		Info: &app,
	}, false)

	if assert.Nil(t, err) {
		assert.NotEqual(t, resp.Info.ID, uuid.UUID{})
		assertApp(t, resp.Info, &app)
	}

	resp1, err := Get(context.Background(), &npool.GetAppRequest{
		ID: resp.Info.ID,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, resp1.Info.ID, resp.Info.ID)
		assertApp(t, resp1.Info, &app)
	}

	app.ID = resp.Info.ID
	app.Name = fmt.Sprintf("App-%v", uuid.New().String())
	resp2, err := Update(context.Background(), &npool.UpdateAppRequest{
		Info: &app,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, resp2.Info.ID, resp.Info.ID)
		assertApp(t, resp2.Info, &app)
	}

	resp3, err := GetAll(context.Background(), &npool.GetAppsRequest{})
	if assert.Nil(t, err) {
		assert.NotEqual(t, len(resp3.Infos), 0)
	}

	resp4, err := GetByCreator(context.Background(), &npool.GetAppsByCreatorRequest{
		UserID: app.CreatedBy,
	})
	if assert.Nil(t, err) {
		assert.NotEqual(t, len(resp4.Infos), 0)
	}
}
