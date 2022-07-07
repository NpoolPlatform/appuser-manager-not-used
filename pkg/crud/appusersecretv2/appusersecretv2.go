package appusersecretv2

import (
	"context"
	"fmt"
	"time"

	"github.com/NpoolPlatform/appuser-manager/pkg/db"
	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent"
	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent/appusersecret"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/appusermgrv2/appusersecret"
	"github.com/google/uuid"
)

func Create(ctx context.Context, in *npool.AppUserSecret) (*ent.AppUserSecret, error) {
	var info *ent.AppUserSecret
	var err error

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		c := cli.AppUserSecret.Create()
		if in.GoogleSecret != nil {
			c.SetGoogleSecret(in.GetGoogleSecret())
		}
		if in.ID != nil {
			c.SetID(uuid.MustParse(in.GetID()))
		}
		if in.AppID != nil {
			c.SetAppID(uuid.MustParse(in.GetAppID()))
		}
		if in.UserID != nil {
			c.SetUserID(uuid.MustParse(in.GetUserID()))
		}
		if in.PasswordHash != nil {
			c.SetPasswordHash(in.GetPasswordHash())
		}
		if in.Salt != nil {
			c.SetSalt(in.GetSalt())
		}
		info, err = c.Save(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}

func CreateBulk(ctx context.Context, in []*npool.AppUserSecret) ([]*ent.AppUserSecret, error) {
	rows := []*ent.AppUserSecret{}
	var err error
	err = db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		bulk := make([]*ent.AppUserSecretCreate, len(in))
		for i, info := range in {
			bulk[i] = tx.AppUserSecret.Create()
			if info.ID != nil {
				bulk[i].SetID(uuid.MustParse(info.GetID()))
			}
			if info.AppID != nil {
				bulk[i].SetAppID(uuid.MustParse(info.GetAppID()))
			}
			if info.UserID != nil {
				bulk[i].SetUserID(uuid.MustParse(info.GetUserID()))
			}
			if info.PasswordHash != nil {
				bulk[i].SetPasswordHash(info.GetPasswordHash())
			}
			if info.Salt != nil {
				bulk[i].SetSalt(info.GetSalt())
			}
			if info.GoogleSecret != nil {
				bulk[i].SetGoogleSecret(info.GetGoogleSecret())
			}
		}
		rows, err = tx.AppUserSecret.CreateBulk(bulk...).Save(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}
	return rows, nil
}

func Update(ctx context.Context, in *npool.AppUserSecret) (*ent.AppUserSecret, error) {
	var info *ent.AppUserSecret
	var err error

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		u := cli.AppUserSecret.UpdateOneID(uuid.MustParse(in.GetID()))
		if in.PasswordHash != nil {
			u.SetPasswordHash(in.GetPasswordHash())
		}
		if in.Salt != nil {
			u.SetSalt(in.GetSalt())
		}
		if in.GoogleSecret != nil {
			u.SetGoogleSecret(in.GetGoogleSecret())
		}
		info, err = u.Save(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}

func Row(ctx context.Context, id uuid.UUID) (*ent.AppUserSecret, error) {
	var info *ent.AppUserSecret
	var err error
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		info, err = cli.AppUserSecret.Query().Where(appusersecret.ID(id)).Only(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}

//nolint
func setQueryConds(conds *npool.Conds, cli *ent.Client) (*ent.AppUserSecretQuery, error) {
	stm := cli.AppUserSecret.Query()
	if conds.ID != nil {
		id := uuid.MustParse(conds.GetID().GetValue())
		switch conds.GetID().GetOp() {
		case cruder.EQ:
			stm.Where(appusersecret.ID(id))

		case cruder.IN:
			stm.Where(appusersecret.ID(id))

		default:
			return nil, fmt.Errorf("invalid appusersecret field")
		}
	}
	if conds.AppID != nil {
		appID := uuid.MustParse(conds.GetAppID().GetValue())
		switch conds.GetAppID().GetOp() {
		case cruder.EQ:
			stm.Where(appusersecret.AppID(appID))

		case cruder.IN:
			stm.Where(appusersecret.AppID(appID))

		default:
			return nil, fmt.Errorf("invalid appusersecret field")
		}
	}
	if conds.UserID != nil {
		userID := uuid.MustParse(conds.GetUserID().GetValue())
		switch conds.GetUserID().GetOp() {
		case cruder.EQ:
			stm.Where(appusersecret.UserID(userID))

		case cruder.IN:
			stm.Where(appusersecret.UserID(userID))

		default:
			return nil, fmt.Errorf("invalid appusersecret field")
		}
	}
	return stm, nil
}

func Rows(ctx context.Context, conds *npool.Conds, offset, limit int) ([]*ent.AppUserSecret, int, error) {
	rows := []*ent.AppUserSecret{}
	var total int
	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm, err := setQueryConds(conds, cli)
		if err != nil {
			return err
		}
		total, err = stm.Count(_ctx)
		if err != nil {
			return err
		}

		rows, err = stm.
			Offset(offset).
			Order(ent.Desc(appusersecret.FieldUpdatedAt)).
			Limit(limit).
			All(_ctx)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, 0, err
	}
	return rows, total, nil
}

func RowOnly(ctx context.Context, conds *npool.Conds) (*ent.AppUserSecret, error) {
	var info *ent.AppUserSecret

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm, err := setQueryConds(conds, cli)
		if err != nil {
			return err
		}

		info, err = stm.Only(_ctx)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}

func Count(ctx context.Context, conds *npool.Conds) (uint32, error) {
	var err error
	var total int

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm, err := setQueryConds(conds, cli)
		if err != nil {
			return err
		}

		total, err = stm.Count(_ctx)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return 0, err
	}

	return uint32(total), nil
}

func Exist(ctx context.Context, id uuid.UUID) (bool, error) {
	var err error
	exist := false

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		exist, err = cli.AppUserSecret.Query().Where(appusersecret.ID(id)).Exist(_ctx)
		return err
	})
	if err != nil {
		return false, err
	}

	return exist, nil
}

func ExistConds(ctx context.Context, conds *npool.Conds) (bool, error) {
	var err error
	exist := false

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm, err := setQueryConds(conds, cli)
		if err != nil {
			return err
		}

		exist, err = stm.Exist(_ctx)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return false, err
	}

	return exist, nil
}

func Delete(ctx context.Context, id uuid.UUID) (*ent.AppUserSecret, error) {
	var info *ent.AppUserSecret
	var err error

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		info, err = cli.AppUserSecret.UpdateOneID(id).
			SetDeletedAt(uint32(time.Now().Unix())).
			Save(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}
