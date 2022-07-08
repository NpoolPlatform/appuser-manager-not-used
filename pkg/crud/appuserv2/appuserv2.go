package appuserv2

import (
	"context"
	"fmt"
	"time"

	"github.com/NpoolPlatform/appuser-manager/pkg/db"
	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent"
	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent/appuser"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/appusermgrv2/appuser"
	"github.com/google/uuid"
)

func Create(ctx context.Context, in *npool.AppUserReq) (*ent.AppUser, error) {
	var info *ent.AppUser
	var err error

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		c := cli.AppUser.Create()
		if in.EmailAddress != nil {
			c.SetEmailAddress(in.GetEmailAddress())
		}
		if in.PhoneNo != nil {
			c.SetPhoneNo(in.GetPhoneNo())
		}
		if in.ImportFromApp != nil {
			c.SetImportFromApp(uuid.MustParse(in.GetImportFromApp()))
		}
		if in.ID != nil {
			c.SetID(uuid.MustParse(in.GetID()))
		}
		if in.AppID != nil {
			c.SetAppID(uuid.MustParse(in.GetAppID()))
		}
		info, err = c.Save(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}

func CreateBulk(ctx context.Context, in []*npool.AppUserReq) ([]*ent.AppUser, error) {
	rows := []*ent.AppUser{}
	var err error
	err = db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		bulk := make([]*ent.AppUserCreate, len(in))
		for i, info := range in {
			bulk[i] = tx.AppUser.Create()
			if info.PhoneNo != nil {
				bulk[i].SetPhoneNo(info.GetPhoneNo())
			}
			if info.ImportFromApp != nil {
				bulk[i].SetImportFromApp(uuid.MustParse(info.GetImportFromApp()))
			}
			if info.ID != nil {
				bulk[i].SetID(uuid.MustParse(info.GetID()))
			}
			if info.AppID != nil {
				bulk[i].SetAppID(uuid.MustParse(info.GetAppID()))
			}
			if info.EmailAddress != nil {
				bulk[i].SetEmailAddress(info.GetEmailAddress())
			}
		}
		rows, err = tx.AppUser.CreateBulk(bulk...).Save(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}
	return rows, nil
}

func Update(ctx context.Context, in *npool.AppUserReq) (*ent.AppUser, error) {
	var info *ent.AppUser
	var err error

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		u := cli.AppUser.UpdateOneID(uuid.MustParse(in.GetID()))
		if in.EmailAddress != nil {
			u.SetEmailAddress(in.GetEmailAddress())
		}
		if in.PhoneNo != nil {
			u.SetPhoneNo(in.GetPhoneNo())
		}
		info, err = u.Save(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}

func Row(ctx context.Context, id uuid.UUID) (*ent.AppUser, error) {
	var info *ent.AppUser
	var err error
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		info, err = cli.AppUser.Query().Where(appuser.ID(id)).Only(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}

//nolint
func setQueryConds(conds *npool.Conds, cli *ent.Client) (*ent.AppUserQuery, error) {
	stm := cli.AppUser.Query()
	if conds.ID != nil {
		id := uuid.MustParse(conds.GetID().GetValue())
		switch conds.GetID().GetOp() {
		case cruder.EQ:
			stm.Where(appuser.ID(id))

		case cruder.IN:
			stm.Where(appuser.ID(id))

		default:
			return nil, fmt.Errorf("invalid appuser field")
		}
	}
	if conds.AppID != nil {
		appID := uuid.MustParse(conds.GetAppID().GetValue())
		switch conds.GetAppID().GetOp() {
		case cruder.EQ:
			stm.Where(appuser.AppID(appID))

		case cruder.IN:
			stm.Where(appuser.AppID(appID))

		default:
			return nil, fmt.Errorf("invalid appuser field")
		}
	}
	if conds.EmailAddress != nil {
		switch conds.GetEmailAddress().GetOp() {
		case cruder.EQ:
			stm.Where(appuser.EmailAddress(conds.GetEmailAddress().GetValue()))

		case cruder.IN:
			stm.Where(appuser.EmailAddress(conds.GetEmailAddress().GetValue()))

		default:
			return nil, fmt.Errorf("invalid appuser field")
		}
	}
	if conds.PhoneNo != nil {
		switch conds.GetPhoneNo().GetOp() {
		case cruder.EQ:
			stm.Where(appuser.PhoneNo(conds.GetPhoneNo().GetValue()))

		case cruder.IN:
			stm.Where(appuser.PhoneNo(conds.GetPhoneNo().GetValue()))

		default:
			return nil, fmt.Errorf("invalid appuser field")
		}
	}
	if conds.ImportFromApp != nil {
		importFromApp := uuid.MustParse(conds.GetImportFromApp().GetValue())
		switch conds.GetImportFromApp().GetOp() {
		case cruder.EQ:
			stm.Where(appuser.ImportFromApp(importFromApp))

		case cruder.IN:
			stm.Where(appuser.ImportFromApp(importFromApp))

		default:
			return nil, fmt.Errorf("invalid appuser field")
		}
	}
	return stm, nil
}

func Rows(ctx context.Context, conds *npool.Conds, offset, limit int) ([]*ent.AppUser, int, error) {
	rows := []*ent.AppUser{}
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
			Order(ent.Desc(appuser.FieldUpdatedAt)).
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

func RowOnly(ctx context.Context, conds *npool.Conds) (*ent.AppUser, error) {
	var info *ent.AppUser

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
		exist, err = cli.AppUser.Query().Where(appuser.ID(id)).Exist(_ctx)
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

func Delete(ctx context.Context, id uuid.UUID) (*ent.AppUser, error) {
	var info *ent.AppUser
	var err error

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		info, err = cli.AppUser.UpdateOneID(id).
			SetDeletedAt(uint32(time.Now().Unix())).
			Save(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}
