package appuserthirdpartyv2

import (
	"context"
	"fmt"
	"time"

	"github.com/NpoolPlatform/appuser-manager/pkg/db"
	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent"
	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent/appuserthirdparty"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/appusermgrv2/appuserthirdparty"
	"github.com/google/uuid"
)

func Create(ctx context.Context, in *npool.AppUserThirdPartyReq) (*ent.AppUserThirdParty, error) {
	var info *ent.AppUserThirdParty
	var err error

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		c := cli.AppUserThirdParty.Create()
		if in.ID != nil {
			c.SetID(uuid.MustParse(in.GetID()))
		}
		if in.AppID != nil {
			c.SetAppID(uuid.MustParse(in.GetAppID()))
		}
		if in.UserID != nil {
			c.SetUserID(uuid.MustParse(in.GetUserID()))
		}
		if in.ThirdPartyUserID != nil {
			c.SetThirdPartyUserID(in.GetThirdPartyUserID())
		}
		if in.ThirdPartyID != nil {
			c.SetThirdPartyID(in.GetThirdPartyID())
		}
		if in.ThirdPartyUsername != nil {
			c.SetThirdPartyUsername(in.GetThirdPartyUsername())
		}
		if in.ThirdPartyUserAvatar != nil {
			c.SetThirdPartyUserAvatar(in.GetThirdPartyUserAvatar())
		}
		info, err = c.Save(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}

func CreateBulk(ctx context.Context, in []*npool.AppUserThirdPartyReq) ([]*ent.AppUserThirdParty, error) {
	rows := []*ent.AppUserThirdParty{}
	var err error
	err = db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		bulk := make([]*ent.AppUserThirdPartyCreate, len(in))
		for i, info := range in {
			bulk[i] = tx.AppUserThirdParty.Create()
			if info.ThirdPartyUsername != nil {
				bulk[i].SetThirdPartyUsername(info.GetThirdPartyUsername())
			}
			if info.ThirdPartyUserAvatar != nil {
				bulk[i].SetThirdPartyUserAvatar(info.GetThirdPartyUserAvatar())
			}
			if info.ID != nil {
				bulk[i].SetID(uuid.MustParse(info.GetID()))
			}
			if info.AppID != nil {
				bulk[i].SetAppID(uuid.MustParse(info.GetAppID()))
			}
			if info.UserID != nil {
				bulk[i].SetUserID(uuid.MustParse(info.GetUserID()))
			}
			if info.ThirdPartyUserID != nil {
				bulk[i].SetThirdPartyUserID(info.GetThirdPartyUserID())
			}
			if info.ThirdPartyID != nil {
				bulk[i].SetThirdPartyID(info.GetThirdPartyID())
			}
		}
		rows, err = tx.AppUserThirdParty.CreateBulk(bulk...).Save(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}
	return rows, nil
}

func Update(ctx context.Context, in *npool.AppUserThirdPartyReq) (*ent.AppUserThirdParty, error) {
	var info *ent.AppUserThirdParty
	var err error

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		u := cli.AppUserThirdParty.UpdateOneID(uuid.MustParse(in.GetID()))
		if in.ThirdPartyUsername != nil {
			u.SetThirdPartyUsername(in.GetThirdPartyUsername())
		}
		if in.ThirdPartyUserAvatar != nil {
			u.SetThirdPartyUserAvatar(in.GetThirdPartyUserAvatar())
		}
		info, err = u.Save(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}

func Row(ctx context.Context, id uuid.UUID) (*ent.AppUserThirdParty, error) {
	var info *ent.AppUserThirdParty
	var err error
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		info, err = cli.AppUserThirdParty.Query().Where(appuserthirdparty.ID(id)).Only(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}

//nolint
func setQueryConds(conds *npool.Conds, cli *ent.Client) (*ent.AppUserThirdPartyQuery, error) {
	stm := cli.AppUserThirdParty.Query()
	if conds.ID != nil {
		id := uuid.MustParse(conds.GetID().GetValue())
		switch conds.GetID().GetOp() {
		case cruder.EQ:
			stm.Where(appuserthirdparty.ID(id))

		case cruder.IN:
			stm.Where(appuserthirdparty.ID(id))

		default:
			return nil, fmt.Errorf("invalid appuserthirdparty field")
		}
	}
	if conds.AppID != nil {
		appID := uuid.MustParse(conds.GetAppID().GetValue())
		switch conds.GetAppID().GetOp() {
		case cruder.EQ:
			stm.Where(appuserthirdparty.AppID(appID))

		case cruder.IN:
			stm.Where(appuserthirdparty.AppID(appID))

		default:
			return nil, fmt.Errorf("invalid appuserthirdparty field")
		}
	}
	if conds.UserID != nil {
		userID := uuid.MustParse(conds.GetUserID().GetValue())
		switch conds.GetUserID().GetOp() {
		case cruder.EQ:
			stm.Where(appuserthirdparty.UserID(userID))

		case cruder.IN:
			stm.Where(appuserthirdparty.UserID(userID))

		default:
			return nil, fmt.Errorf("invalid appuserthirdparty field")
		}
	}
	if conds.ThirdPartyUserID != nil {
		switch conds.GetThirdPartyUserID().GetOp() {
		case cruder.EQ:
			stm.Where(appuserthirdparty.ThirdPartyUserID(conds.GetThirdPartyUserID().GetValue()))

		case cruder.IN:
			stm.Where(appuserthirdparty.ThirdPartyUserID(conds.GetThirdPartyUserID().GetValue()))

		default:
			return nil, fmt.Errorf("invalid appuserthirdparty field")
		}
	}
	if conds.ThirdPartyID != nil {
		switch conds.GetThirdPartyID().GetOp() {
		case cruder.EQ:
			stm.Where(appuserthirdparty.ThirdPartyID(conds.GetThirdPartyID().GetValue()))

		case cruder.IN:
			stm.Where(appuserthirdparty.ThirdPartyID(conds.GetThirdPartyID().GetValue()))

		default:
			return nil, fmt.Errorf("invalid appuserthirdparty field")
		}
	}
	return stm, nil
}

func Rows(ctx context.Context, conds *npool.Conds, offset, limit int) ([]*ent.AppUserThirdParty, int, error) {
	rows := []*ent.AppUserThirdParty{}
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
			Order(ent.Desc(appuserthirdparty.FieldUpdatedAt)).
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

func RowOnly(ctx context.Context, conds *npool.Conds) (*ent.AppUserThirdParty, error) {
	var info *ent.AppUserThirdParty

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
		exist, err = cli.AppUserThirdParty.Query().Where(appuserthirdparty.ID(id)).Exist(_ctx)
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

func Delete(ctx context.Context, id uuid.UUID) (*ent.AppUserThirdParty, error) {
	var info *ent.AppUserThirdParty
	var err error

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		info, err = cli.AppUserThirdParty.UpdateOneID(id).
			SetDeletedAt(uint32(time.Now().Unix())).
			Save(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}
