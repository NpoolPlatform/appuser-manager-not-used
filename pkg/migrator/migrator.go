package migrator

import (
	"context"
	"github.com/NpoolPlatform/appuser-manager/pkg/db"
	sm "github.com/NpoolPlatform/message/npool/appuser/mgr/v2/signmethod"
)

func Migrate(ctx context.Context) error {
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
