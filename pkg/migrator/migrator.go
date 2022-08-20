package migrator

import (
	"context"
	"encoding/json"

	"github.com/NpoolPlatform/go-service-framework/pkg/config"

	"github.com/NpoolPlatform/appuser-manager/pkg/db"
	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent"
	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent/approle"

	constant1 "github.com/NpoolPlatform/appuser-manager/pkg/const"
	constant2 "github.com/NpoolPlatform/appuser-manager/pkg/message/const"
	approlepb "github.com/NpoolPlatform/message/npool/appuser/mgr/v2/approle"
	sm "github.com/NpoolPlatform/message/npool/appuser/mgr/v2/signmethod"

	"github.com/google/uuid"
)

func Migrate(ctx context.Context) error {
	err := updateAdminRoleAppID(ctx)
	if err != nil {
		return err
	}

	return setSigninVerifyTypeVal(ctx)
}

func setSigninVerifyTypeVal(ctx context.Context) error {
	cli, err := db.Client()
	if err != nil {
		return err
	}
	appUserControls, err := cli.AppUserControl.Query().All(ctx)
	if err != nil {
		return err
	}

	for _, val := range appUserControls {
		signinVerifyType := sm.SignMethodType_Email.String()
		if val.SigninVerifyByGoogleAuthentication {
			signinVerifyType = sm.SignMethodType_Google.String()
		}

		if _, err = cli.AppUserControl.
			UpdateOneID(val.ID).
			SetSigninVerifyType(signinVerifyType).
			Save(ctx); err != nil {
			return err
		}
	}
	return nil
}

func updateAdminRoleAppID(ctx context.Context) error {
	hostname := config.GetStringValueWithNameSpace(constant2.ServiceName, config.KeyHostname)
	genesisRoleStr := config.GetStringValueWithNameSpace(hostname, constant1.KeyGenesisRole)

	appRoles := []*approlepb.AppRole{}

	err := json.Unmarshal([]byte(genesisRoleStr), &appRoles)
	if err != nil {
		return err
	}

	cli, err := db.Client()
	if err != nil {
		return err
	}

	for _, val := range appRoles {
		role, err := cli.
			AppRole.
			Query().
			Where(
				approle.Role(val.Role),
			).
			Only(ctx)
		if err != nil {
			if ent.IsNotFound(err) {
				continue
			}
			return err
		}

		if _, err = role.
			Update().
			SetAppID(
				uuid.MustParse(val.GetAppID()),
			).
			Save(ctx); err != nil {
			return err
		}
	}
	return nil
}
