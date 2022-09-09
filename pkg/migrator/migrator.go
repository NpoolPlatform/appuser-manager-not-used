package migrator

import (
	"context"

	constant "github.com/NpoolPlatform/appuser-manager/pkg/const"
	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent/appcontrol"

	"github.com/NpoolPlatform/appuser-manager/pkg/db"
	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent"

	"github.com/NpoolPlatform/message/npool/appuser/mgr/v2/recaptcha"
)

func Migrate(ctx context.Context) error {
	return UpdateRecaptchaMethod(ctx)
}

func UpdateRecaptchaMethod(ctx context.Context) error {
	err := db.WithTx(ctx, func(ctx context.Context, tx *ent.Tx) error {
		infos, err := tx.AppControl.
			Query().
			Where(
				appcontrol.RecaptchaMethod(constant.RecaptchaGoogleV3),
			).
			ForUpdate().
			All(ctx)
		if err != nil {
			return err
		}
		for _, val := range infos {
			_, err = val.Update().
				SetRecaptchaMethod(recaptcha.RecaptchaType_GoogleRecaptchaV3.String()).
				Save(ctx)
			if err != nil {
				return err
			}
		}
		return err
	})
	return err
}
