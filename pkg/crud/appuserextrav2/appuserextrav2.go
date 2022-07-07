package appuserextrav2

import (
	"context"
	"fmt"
	"time"

	"github.com/NpoolPlatform/appuser-manager/pkg/db"
	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent"
	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent/appuserextra"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/appusermgrv2/appuserextra"
	"github.com/google/uuid"
)

func Create(ctx context.Context, in *npool.AppUserExtra) (*ent.AppUserExtra, error) {
	var info *ent.AppUserExtra
	var err error

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		c := cli.AppUserExtra.Create()
		if in.FirstName != nil {
			c.SetFirstName(in.GetFirstName())
		}
		if in.AppID != nil {
			c.SetAppID(uuid.MustParse(in.GetAppID()))
		}
		if in.Organization != nil {
			c.SetOrganization(in.GetOrganization())
		}
		if in.IDNumber != nil {
			c.SetIDNumber(in.GetIDNumber())
		}
		if in.PostalCode != nil {
			c.SetPostalCode(in.GetPostalCode())
		}
		if in.Age != nil {
			c.SetAge(in.GetAge())
		}
		if in.Birthday != nil {
			c.SetBirthday(in.GetBirthday())
		}
		if in.ID != nil {
			c.SetID(uuid.MustParse(in.GetID()))
		}
		if in.Avatar != nil {
			c.SetAvatar(in.GetAvatar())
		}
		if in.Username != nil {
			c.SetUsername(in.GetUsername())
		}
		if in.LastName != nil {
			c.SetLastName(in.GetLastName())
		}
		if in.Gender != nil {
			c.SetGender(in.GetGender())
		}
		if in.UserID != nil {
			c.SetUserID(uuid.MustParse(in.GetUserID()))
		}
		if in.UserID != nil {
			c.SetAddressFields(in.GetAddressFields())
		}
		info, err = c.Save(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}

func CreateBulk(ctx context.Context, in []*npool.AppUserExtra) ([]*ent.AppUserExtra, error) {
	rows := []*ent.AppUserExtra{}
	var err error
	err = db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		bulk := make([]*ent.AppUserExtraCreate, len(in))
		for i, info := range in {
			bulk[i] = tx.AppUserExtra.Create()
			if info.ID != nil {
				bulk[i].SetID(uuid.MustParse(info.GetID()))
			}
			if info.AppID != nil {
				bulk[i].SetAppID(uuid.MustParse(info.GetAppID()))
			}
			if info.Avatar != nil {
				bulk[i].SetAvatar(info.GetAvatar())
			}
			if info.Organization != nil {
				bulk[i].SetOrganization(info.GetOrganization())
			}
			if info.UserID != nil {
				bulk[i].SetUserID(uuid.MustParse(info.GetUserID()))
			}
			if info.LastName != nil {
				bulk[i].SetLastName(info.GetLastName())
			}
			if info.PostalCode != nil {
				bulk[i].SetPostalCode(info.GetPostalCode())
			}
			if info.Age != nil {
				bulk[i].SetAge(info.GetAge())
			}
			if info.FirstName != nil {
				bulk[i].SetFirstName(info.GetFirstName())
			}
			if info.IDNumber != nil {
				bulk[i].SetIDNumber(info.GetIDNumber())
			}
			if info.Gender != nil {
				bulk[i].SetGender(info.GetGender())
			}
			if info.Birthday != nil {
				bulk[i].SetBirthday(info.GetBirthday())
			}
			if info.Username != nil {
				bulk[i].SetUsername(info.GetUsername())
			}
			if info.AddressFields != nil {
				bulk[i].SetAddressFields(info.GetAddressFields())
			}
		}
		rows, err = tx.AppUserExtra.CreateBulk(bulk...).Save(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}
	return rows, nil
}

func Update(ctx context.Context, in *npool.AppUserExtra) (*ent.AppUserExtra, error) {
	var info *ent.AppUserExtra
	var err error

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		u := cli.AppUserExtra.UpdateOneID(uuid.MustParse(in.GetID()))
		if in.Username != nil {
			u.SetUsername(in.GetUsername())
		}
		if in.FirstName != nil {
			u.SetFirstName(in.GetFirstName())
		}
		if in.LastName != nil {
			u.SetLastName(in.GetLastName())
		}
		if in.AddressFields != nil {
			u.SetAddressFields(in.GetAddressFields())
		}
		info, err = u.Save(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}

func Row(ctx context.Context, id uuid.UUID) (*ent.AppUserExtra, error) {
	var info *ent.AppUserExtra
	var err error
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		info, err = cli.AppUserExtra.Query().Where(appuserextra.ID(id)).Only(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}

//nolint
func setQueryConds(conds *npool.Conds, cli *ent.Client) (*ent.AppUserExtraQuery, error) {
	stm := cli.AppUserExtra.Query()
	if conds.ID != nil {
		id := uuid.MustParse(conds.GetID().GetValue())
		switch conds.GetID().GetOp() {
		case cruder.EQ:
			stm.Where(appuserextra.ID(id))

		case cruder.IN:
			stm.Where(appuserextra.ID(id))

		default:
			return nil, fmt.Errorf("invalid appuserextra field")
		}
	}
	if conds.AppID != nil {
		appID := uuid.MustParse(conds.GetAppID().GetValue())
		switch conds.GetAppID().GetOp() {
		case cruder.EQ:
			stm.Where(appuserextra.AppID(appID))

		case cruder.IN:
			stm.Where(appuserextra.AppID(appID))

		default:
			return nil, fmt.Errorf("invalid appuserextra field")
		}
	}
	if conds.UserID != nil {
		userID := uuid.MustParse(conds.GetUserID().GetValue())
		switch conds.GetUserID().GetOp() {
		case cruder.EQ:
			stm.Where(appuserextra.UserID(userID))

		case cruder.IN:
			stm.Where(appuserextra.UserID(userID))

		default:
			return nil, fmt.Errorf("invalid appuserextra field")
		}
	}
	if conds.IDNumber != nil {
		switch conds.GetIDNumber().GetOp() {
		case cruder.EQ:
			stm.Where(appuserextra.IDNumber(conds.GetIDNumber().GetValue()))

		case cruder.IN:
			stm.Where(appuserextra.IDNumber(conds.GetIDNumber().GetValue()))

		default:
			return nil, fmt.Errorf("invalid appuserextra field")
		}
	}
	return stm, nil
}

func Rows(ctx context.Context, conds *npool.Conds, offset, limit int) ([]*ent.AppUserExtra, int, error) {
	rows := []*ent.AppUserExtra{}
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
			Order(ent.Desc(appuserextra.FieldUpdatedAt)).
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

func RowOnly(ctx context.Context, conds *npool.Conds) (*ent.AppUserExtra, error) {
	var info *ent.AppUserExtra

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
		exist, err = cli.AppUserExtra.Query().Where(appuserextra.ID(id)).Exist(_ctx)
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

func Delete(ctx context.Context, id uuid.UUID) (*ent.AppUserExtra, error) {
	var info *ent.AppUserExtra
	var err error

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		info, err = cli.AppUserExtra.UpdateOneID(id).
			SetDeletedAt(uint32(time.Now().Unix())).
			Save(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}
