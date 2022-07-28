package approleuser

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	testinit "github.com/NpoolPlatform/appuser-manager/pkg/testinit" //nolint
	npool "github.com/NpoolPlatform/message/npool/appuser/mgr/v1"

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

func assertAppRoleUser(t *testing.T, actual, expected *npool.AppRoleUser) {
	assert.Equal(t, actual.AppID, expected.AppID)
	assert.Equal(t, actual.RoleID, expected.RoleID)
	assert.Equal(t, actual.UserID, expected.UserID)
}

func TestCRUD(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	appRoleUser := npool.AppRoleUser{
		AppID:  uuid.New().String(),
		RoleID: uuid.New().String(),
		UserID: uuid.New().String(),
	}

	resp, err := Create(context.Background(), &npool.CreateAppRoleUserRequest{
		Info: &appRoleUser,
	})

	if assert.Nil(t, err) {
		assert.NotEqual(t, resp.Info.ID, uuid.UUID{})
		assertAppRoleUser(t, resp.Info, &appRoleUser)
	}

	resp1, err := Get(context.Background(), &npool.GetAppRoleUserRequest{
		ID: resp.Info.ID,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, resp1.Info.ID, resp.Info.ID)
		assertAppRoleUser(t, resp1.Info, &appRoleUser)
	}

	resp2, err := GetUsersByAppRole(context.Background(), &npool.GetAppRoleUsersByAppRoleRequest{
		AppID:  appRoleUser.AppID,
		RoleID: appRoleUser.RoleID,
	})
	if assert.Nil(t, err) {
		assert.NotEqual(t, len(resp2.Infos), 0)
	}

	appRoleUser.ID = resp.Info.ID
	resp3, err := Delete(context.Background(), &npool.DeleteAppRoleUserRequest{
		ID: resp.Info.ID,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, resp3.Info.ID, resp.Info.ID)
		assertAppRoleUser(t, resp3.Info, &appRoleUser)
	}
}
