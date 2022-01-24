package approle

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/appuser-manager/pkg/test-init" //nolint
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

func assertAppRole(t *testing.T, actual, expected *npool.AppRole) {
	assert.Equal(t, actual.AppID, expected.AppID)
	assert.Equal(t, actual.CreatedBy, expected.CreatedBy)
	assert.Equal(t, actual.Role, expected.Role)
	assert.Equal(t, actual.Description, expected.Description)
	assert.Equal(t, actual.Default, expected.Default)
}

func TestCRUD(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	appRole := npool.AppRole{
		AppID:       uuid.New().String(),
		CreatedBy:   uuid.New().String(),
		Role:        uuid.New().String(),
		Description: uuid.New().String(),
		Default:     false,
	}

	resp, err := Create(context.Background(), &npool.CreateAppRoleRequest{
		Info: &appRole,
	})

	if assert.Nil(t, err) {
		assert.NotEqual(t, resp.Info.ID, uuid.UUID{})
		assertAppRole(t, resp.Info, &appRole)
	}

	resp1, err := Get(context.Background(), &npool.GetAppRoleRequest{
		ID: resp.Info.ID,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, resp1.Info.ID, resp.Info.ID)
		assertAppRole(t, resp1.Info, &appRole)
	}

	appRole.ID = resp.Info.ID
	resp2, err := Update(context.Background(), &npool.UpdateAppRoleRequest{
		Info: &appRole,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, resp2.Info.ID, resp.Info.ID)
		assertAppRole(t, resp2.Info, &appRole)
	}

	resp3, err := GetByApp(context.Background(), &npool.GetAppRolesByAppRequest{
		AppID: appRole.AppID,
	})
	if assert.Nil(t, err) {
		assert.NotEqual(t, len(resp3.Infos), 0)
	}
}
