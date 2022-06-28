package appv2

import (
	"context"
	"fmt"
	constant "github.com/NpoolPlatform/appuser-manager/pkg/const"
	"github.com/NpoolPlatform/appuser-manager/pkg/db"
	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent"
	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent/app"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/appusermgrv2/app"
	"github.com/google/uuid"
	"time"
)

type App struct {
	*db.Entity
}

func New(ctx context.Context, tx *ent.Tx) (*App, error) {
	e, err := db.NewEntity(ctx, tx)
	if err != nil {
		return nil, fmt.Errorf("fail create entity: %v", err)
	}

	return &App{
		Entity: e,
	}, nil
}

func (s *App) rowToObject(row *ent.App) *npool.App {
	return &npool.App{
		ID:          row.ID.String(),
		CreatedBy:   row.CreatedBy.String(),
		Name:        row.Name,
		Logo:        row.Logo,
		Description: row.Description,
		CreateAt:    row.CreatedAt,
	}
}

func (s *App) Create(ctx context.Context, in *npool.App) (*npool.App, error) {
	var info *ent.App
	var err error

	err = db.WithTx(ctx, s.Tx, func(_ctx context.Context) error {
		id := uuid.New()
		if in.GetID() != "" {
			id = uuid.MustParse(in.GetID())
		}
		info, err = s.Tx.App.Create().
			SetID(id).
			SetCreatedBy(uuid.MustParse(in.GetCreatedBy())).
			SetName(in.GetName()).
			SetLogo(in.GetLogo()).
			SetDescription(in.GetDescription()).
			Save(_ctx)
		return err
	})
	if err != nil {
		return nil, fmt.Errorf("fail create app: %v", err)
	}

	return s.rowToObject(info), nil
}

func (s *App) CreateBulk(ctx context.Context, in []*npool.App) ([]*npool.App, error) {
	rows := []*ent.App{}
	var err error

	err = db.WithTx(ctx, s.Tx, func(_ctx context.Context) error {
		bulk := make([]*ent.AppCreate, len(in))
		for i, info := range in {
			id := uuid.New()
			if info.GetID() != "" {
				id = uuid.MustParse(info.GetID())
			}
			bulk[i] = s.Tx.App.Create().
				SetID(id).
				SetCreatedBy(uuid.MustParse(info.GetCreatedBy())).
				SetName(info.GetName()).
				SetLogo(info.GetLogo()).
				SetDescription(info.GetDescription())
		}
		rows, err = s.Tx.App.CreateBulk(bulk...).Save(_ctx)
		return err
	})
	if err != nil {
		return nil, fmt.Errorf("fail create bulk app: %v", err)
	}

	infos := []*npool.App{}
	for _, row := range rows {
		infos = append(infos, s.rowToObject(row))
	}

	return infos, nil
}

func (s *App) Update(ctx context.Context, in *npool.App) (*npool.App, error) {
	var info *ent.App
	var err error

	err = db.WithTx(ctx, s.Tx, func(_ctx context.Context) error {
		info, err = s.Tx.App.UpdateOneID(uuid.MustParse(in.GetID())).
			SetName(in.GetName()).
			SetLogo(in.GetLogo()).
			SetDescription(in.GetDescription()).
			Save(_ctx)
		return err
	})
	if err != nil {
		return nil, fmt.Errorf("fail update app: %v", err)
	}

	return s.rowToObject(info), nil
}

func (s *App) UpdateFields(ctx context.Context, id uuid.UUID, fields cruder.Fields) (*npool.App, error) {
	var info *ent.App
	var err error

	err = db.WithTx(ctx, s.Tx, func(_ctx context.Context) error {
		stm := s.Tx.App.UpdateOneID(id)
		for k, v := range fields {
			switch k {
			case constant.AppFieldName:
				name, err := cruder.AnyTypeString(v)
				if err != nil {
					return fmt.Errorf("invalid value type: %v", err)
				}
				stm = stm.SetName(name)
			case constant.AppFieldDescription:
				description, err := cruder.AnyTypeString(v)
				if err != nil {
					return fmt.Errorf("invalid value type: %v", err)
				}
				stm = stm.SetDescription(description)
			case constant.AppFieldLogo:
				logo, err := cruder.AnyTypeString(v)
				if err != nil {
					return fmt.Errorf("invalid value type: %v", err)
				}
				stm = stm.SetLogo(logo)
			default:
				return fmt.Errorf("invalid app field")
			}
		}

		info, err = stm.Save(_ctx)
		if err != nil {
			return fmt.Errorf("fail update app fields: %v", err)
		}

		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail update app: %v", err)
	}

	return s.rowToObject(info), nil
}

func (s *App) Row(ctx context.Context, id uuid.UUID) (*npool.App, error) {
	var info *ent.App
	var err error
	err = db.WithTx(ctx, s.Tx, func(_ctx context.Context) error {
		info, err = s.Tx.App.Query().Where(app.ID(id)).Only(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return s.rowToObject(info), nil
}

func (s *App) queryFromConds(conds cruder.Conds) (*ent.AppQuery, error) {
	stm := s.Tx.App.Query()
	for k, v := range conds {
		switch k {
		case constant.FieldID:
			switch v.Op {
			case cruder.EQ:
				id, err := cruder.AnyTypeUUID(v.Val)
				if err != nil {
					return nil, fmt.Errorf("invalid id: %v", err)
				}
				stm = stm.Where(app.ID(id))
			case cruder.IN:
				id, err := cruder.AnyTypeUUIDs(v.Val)
				if err != nil {
					return nil, fmt.Errorf("invalid ids: %v", err)
				}
				stm = stm.Where(app.IDIn(id...))
			}

		case constant.AppFieldCreatedBy:
			value, err := cruder.AnyTypeUUID(v.Val)
			if err != nil {
				return nil, fmt.Errorf("invalid value type: %v", err)
			}
			stm = stm.Where(app.CreatedBy(value))
		default:
			return nil, fmt.Errorf("invalid app field")
		}
	}
	return stm, nil
}

func (s *App) Rows(ctx context.Context, conds cruder.Conds, offset, limit int) ([]*npool.App, int, error) {
	rows := []*ent.App{}
	var total int

	err := db.WithTx(ctx, s.Tx, func(_ctx context.Context) error {
		stm, err := s.queryFromConds(conds)
		if err != nil {
			return fmt.Errorf("fail construct stm: %v", err)
		}

		total, err = stm.Count(_ctx)
		if err != nil {
			return fmt.Errorf("fail count app: %v", err)
		}

		rows, err = stm.
			Offset(offset).
			Order(ent.Desc(app.FieldUpdatedAt)).
			Limit(limit).
			All(_ctx)
		if err != nil {
			return fmt.Errorf("fail query app: %v", err)
		}

		return nil
	})
	if err != nil {
		return nil, 0, fmt.Errorf("fail get app: %v", err)
	}

	infos := []*npool.App{}
	for _, row := range rows {
		infos = append(infos, s.rowToObject(row))
	}
	return infos, total, nil
}

func (s *App) RowOnly(ctx context.Context, conds cruder.Conds) (*npool.App, error) {
	var info *ent.App

	err := db.WithTx(ctx, s.Tx, func(_ctx context.Context) error {
		stm, err := s.queryFromConds(conds)
		if err != nil {
			return fmt.Errorf("fail construct stm: %v", err)
		}

		info, err = stm.Only(_ctx)
		if err != nil {
			return fmt.Errorf("fail query app: %v", err)
		}

		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail get app: %v", err)
	}

	return s.rowToObject(info), nil
}

func (s *App) Count(ctx context.Context, conds cruder.Conds) (uint32, error) {
	var err error
	var total int

	err = db.WithTx(ctx, s.Tx, func(_ctx context.Context) error {
		stm, err := s.queryFromConds(conds)
		if err != nil {
			return fmt.Errorf("fail construct stm: %v", err)
		}

		total, err = stm.Count(_ctx)
		if err != nil {
			return fmt.Errorf("fail check apps: %v", err)
		}

		return nil
	})
	if err != nil {
		return 0, fmt.Errorf("fail count apps: %v", err)
	}

	return uint32(total), nil
}

func (s *App) Exist(ctx context.Context, id uuid.UUID) (bool, error) {
	var err error
	exist := false

	err = db.WithTx(ctx, s.Tx, func(_ctx context.Context) error {
		exist, err = s.Tx.App.Query().Where(app.ID(id)).Exist(_ctx)
		return err
	})
	if err != nil {
		return false, fmt.Errorf("fail check app: %v", err)
	}

	return exist, nil
}

func (s *App) ExistConds(ctx context.Context, conds cruder.Conds) (bool, error) {
	var err error
	exist := false

	err = db.WithTx(ctx, s.Tx, func(_ctx context.Context) error {
		stm, err := s.queryFromConds(conds)
		if err != nil {
			return fmt.Errorf("fail construct stm: %v", err)
		}

		exist, err = stm.Exist(_ctx)
		if err != nil {
			return fmt.Errorf("fail check app: %v", err)
		}

		return nil
	})
	if err != nil {
		return false, fmt.Errorf("fail check app: %v", err)
	}

	return exist, nil
}

func (s *App) Delete(ctx context.Context, id uuid.UUID) (*npool.App, error) {
	var info *ent.App
	var err error

	err = db.WithTx(ctx, s.Tx, func(_ctx context.Context) error {
		info, err = s.Tx.App.UpdateOneID(id).
			SetDeletedAt(uint32(time.Now().Unix())).
			Save(_ctx)
		return err
	})
	if err != nil {
		return nil, fmt.Errorf("fail delete app: %v", err)
	}

	return s.rowToObject(info), nil
}
