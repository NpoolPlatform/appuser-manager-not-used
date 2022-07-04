package appv2

import (
	"context"
	"fmt"
	"time"

	"github.com/NpoolPlatform/appuser-manager/pkg/db"
	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent"
	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent/app"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/appusermgrv2/app"
	"github.com/google/uuid"
)

type App struct {
	*db.Entity
}

func New(ctx context.Context, tx *ent.Tx) (*App, error) {
	e, err := db.NewEntity(ctx, tx)
	if err != nil {
		logger.Sugar().Errorf("fail create entity:  %v", err)
		return nil, err
	}

	return &App{
		Entity: e,
	}, nil
}

func (s *App) Create(ctx context.Context, in *npool.App) (*ent.App, error) {
	var info *ent.App
	var err error

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		c := cli.App.Create()
		if in.ID != nil {
			c.SetID(uuid.MustParse(in.GetID()))
		}
		if in.CreatedBy != nil {
			c.SetCreatedBy(uuid.MustParse(in.GetCreatedBy()))
		}
		if in.Name != nil {
			c.SetName(in.GetName())
		}
		if in.Logo != nil {
			c.SetLogo(in.GetLogo())
		}
		if in.Description != nil {
			c.SetDescription(in.GetDescription())
		}
		info, err = c.Save(_ctx)
		return err
	})
	if err != nil {
		logger.Sugar().Errorf("fail create app:  %v", err)
		return nil, err
	}

	return info, nil
}

func (s *App) CreateBulk(ctx context.Context, in []*npool.App) ([]*ent.App, error) {
	rows := []*ent.App{}
	var err error
	err = db.WithTx(ctx, s.Tx, func(_ctx context.Context) error {
		bulk := make([]*ent.AppCreate, len(in))
		for i, info := range in {
			bulk[i] = s.Tx.App.Create()
			if info.ID != nil {
				bulk[i].SetID(uuid.MustParse(info.GetID()))
			}
			if info.CreatedBy != nil {
				bulk[i].SetCreatedBy(uuid.MustParse(info.GetCreatedBy()))
			}
			if info.Name != nil {
				bulk[i].SetName(info.GetName())
			}
			if info.Logo != nil {
				bulk[i].SetLogo(info.GetLogo())
			}
			if info.Description != nil {
				bulk[i].SetDescription(info.GetDescription())
			}
		}
		rows, err = s.Tx.App.CreateBulk(bulk...).Save(_ctx)
		return err
	})
	if err != nil {
		logger.Sugar().Errorf("fail create apps:  %v", err)
		return nil, err
	}
	return rows, nil
}

func (s *App) Update(ctx context.Context, in *npool.App) (*ent.App, error) {
	var info *ent.App
	var err error

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		u := cli.App.UpdateOneID(uuid.MustParse(in.GetID()))
		if in.Name != nil {
			u.SetName(in.GetName())
		}
		if in.Logo != nil {
			u.SetLogo(in.GetLogo())
		}
		if in.Description != nil {
			u.SetDescription(in.GetDescription())
		}
		info, err = u.Save(_ctx)
		return err
	})
	if err != nil {
		logger.Sugar().Errorf("fail create app:  %v", err)
		return nil, err
	}

	return info, nil
}

func (s *App) Row(ctx context.Context, id uuid.UUID) (*ent.App, error) {
	var info *ent.App
	var err error
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		info, err = s.Tx.App.Query().Where(app.ID(id)).Only(_ctx)
		return err
	})
	if err != nil {
		logger.Sugar().Errorf("fail get app:  %v", err)
		return nil, err
	}

	return info, nil
}

//nolint
func (s *App) setQueryConds(conds *npool.Conds, cli *ent.Client) (*ent.AppQuery, error) {
	stm := cli.App.Query()
	if conds.ID != nil {
		id := uuid.MustParse(conds.GetID().GetValue())
		switch conds.GetID().GetOp() {
		case cruder.EQ:
			stm.Where(app.ID(id))
		case cruder.IN:
			stm.Where(app.IDIn(id))
		default:
			logger.Sugar().Errorf("invalid app field")
			return nil, fmt.Errorf("invalid app field")
		}
	}
	if conds.CreatedBy != nil {
		createdBy := uuid.MustParse(conds.GetID().GetValue())
		switch conds.GetCreatedBy().GetOp() {
		case cruder.EQ:
			stm.Where(app.CreatedBy(createdBy))
		case cruder.IN:
			stm.Where(app.CreatedByIn(createdBy))
		default:
			logger.Sugar().Errorf("invalid app field")
			return nil, fmt.Errorf("invalid app field")
		}
	}
	if conds.Name != nil {
		switch conds.GetName().GetOp() {
		case cruder.EQ:
			stm.Where(app.Name(conds.GetName().GetValue()))
		case cruder.IN:
			stm.Where(app.NameIn(conds.GetName().GetValue()))
		default:
			logger.Sugar().Errorf("invalid app field")
			return nil, fmt.Errorf("invalid app field")
		}
	}
	if conds.Logo != nil {
		switch conds.GetLogo().GetOp() {
		case cruder.EQ:
			stm.Where(app.Logo(conds.GetLogo().GetValue()))
		case cruder.IN:
			stm.Where(app.LogoIn(conds.GetLogo().GetValue()))
		default:
			logger.Sugar().Errorf("invalid app field")
			return nil, fmt.Errorf("invalid app field")
		}
	}
	if conds.Description != nil {
		switch conds.GetDescription().GetOp() {
		case cruder.EQ:
			stm.Where(app.Description(conds.GetDescription().GetValue()))
		case cruder.IN:
			stm.Where(app.DescriptionIn(conds.GetDescription().GetValue()))
		default:
			logger.Sugar().Errorf("invalid app field")
			return nil, fmt.Errorf("invalid app field")
		}
	}
	return stm, nil
}

func (s *App) Rows(ctx context.Context, conds *npool.Conds, offset, limit int) ([]*ent.App, int, error) {
	rows := []*ent.App{}
	var total int
	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm, err := s.setQueryConds(conds, cli)
		if err != nil {
			logger.Sugar().Errorf("fail construct stm: %v", err)
			return err
		}
		total, err = stm.Count(_ctx)
		if err != nil {
			logger.Sugar().Errorf("fail count app: %v", err)
			return err
		}

		rows, err = stm.
			Offset(offset).
			Order(ent.Desc(app.FieldUpdatedAt)).
			Limit(limit).
			All(_ctx)
		if err != nil {
			logger.Sugar().Errorf("fail query app: %v", err)
			return err
		}

		return nil
	})
	if err != nil {
		logger.Sugar().Errorf("fail get app: %v", err)
		return nil, 0, err
	}
	return rows, total, nil
}

func (s *App) RowOnly(ctx context.Context, conds *npool.Conds) (*ent.App, error) {
	var info *ent.App

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm, err := s.setQueryConds(conds, cli)
		if err != nil {
			logger.Sugar().Errorf("fail construct stm: %v", err)
			return err
		}

		info, err = stm.Only(_ctx)
		if err != nil {
			logger.Sugar().Errorf("fail query app: %v", err)
			return err
		}

		return nil
	})
	if err != nil {
		logger.Sugar().Errorf("fail get app: %v", err)
		return nil, err
	}

	return info, nil
}

func (s *App) Count(ctx context.Context, conds *npool.Conds) (uint32, error) {
	var err error
	var total int

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm, err := s.setQueryConds(conds, cli)
		if err != nil {
			logger.Sugar().Errorf("fail construct stm: %v", err)
			return err
		}

		total, err = stm.Count(_ctx)
		if err != nil {
			logger.Sugar().Errorf("fail check apps: %v", err)
			return err
		}
		return nil
	})
	if err != nil {
		logger.Sugar().Errorf("fail count apps: %v", err)
		return 0, err
	}

	return uint32(total), nil
}

func (s *App) Exist(ctx context.Context, id uuid.UUID) (bool, error) {
	var err error
	exist := false

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		exist, err = s.Tx.App.Query().Where(app.ID(id)).Exist(_ctx)
		return err
	})
	if err != nil {
		logger.Sugar().Errorf("fail check app: %v", err)
		return false, err
	}

	return exist, nil
}

func (s *App) ExistConds(ctx context.Context, conds *npool.Conds) (bool, error) {
	var err error
	exist := false

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm, err := s.setQueryConds(conds, cli)
		if err != nil {
			logger.Sugar().Errorf("fail construct stm: %v", err)
			return err
		}

		exist, err = stm.Exist(_ctx)
		if err != nil {
			logger.Sugar().Errorf("fail check app: %v", err)
			return err
		}

		return nil
	})
	if err != nil {
		logger.Sugar().Errorf("fail check app: %v", err)
		return false, err
	}

	return exist, nil
}

func (s *App) Delete(ctx context.Context, id uuid.UUID) (*ent.App, error) {
	var info *ent.App
	var err error

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		info, err = cli.App.UpdateOneID(id).
			SetDeletedAt(uint32(time.Now().Unix())).
			Save(_ctx)
		return err
	})
	if err != nil {
		logger.Sugar().Errorf("fail delete app: %v", err)
		return nil, err
	}

	return info, nil
}
