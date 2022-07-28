package appv2

import (
	"context"
	"fmt"
	"time"

	constant "github.com/NpoolPlatform/appuser-manager/pkg/message/const"
	tracer "github.com/NpoolPlatform/appuser-manager/pkg/tracer/app"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"

	"github.com/NpoolPlatform/appuser-manager/pkg/db"
	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent"
	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent/app"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/appuser/mgr/v2/app"
	"github.com/google/uuid"
)

func Create(ctx context.Context, in *npool.AppReq) (*ent.App, error) {
	var info *ent.App
	var err error
	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "Create")
	defer span.End()
	defer func() {
		if err != nil {
			span.SetStatus(codes.Error, "db operation fail")
			span.RecordError(err)
		}
	}()

	span = tracer.Trace(span, in)

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
		return nil, err
	}

	return info, nil
}

func CreateBulk(ctx context.Context, in []*npool.AppReq) ([]*ent.App, error) {
	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateBulk")
	defer span.End()
	var err error
	defer func() {
		if err != nil {
			span.SetStatus(codes.Error, "db operation fail")
			span.RecordError(err)
		}
	}()

	span = tracer.TraceMany(span, in)

	rows := []*ent.App{}

	err = db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		bulk := make([]*ent.AppCreate, len(in))
		for i, info := range in {
			bulk[i] = tx.App.Create()
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
		rows, err = tx.App.CreateBulk(bulk...).Save(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return rows, nil
}

func Update(ctx context.Context, in *npool.AppReq) (*ent.App, error) {
	var info *ent.App
	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "Update")
	defer span.End()
	var err error
	defer func() {
		if err != nil {
			span.SetStatus(codes.Error, "db operation fail")
			span.RecordError(err)
		}
	}()
	span = AppSpanAttributes(span, in)

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
		return nil, err
	}

	return info, nil
}

func Row(ctx context.Context, id uuid.UUID) (*ent.App, error) {
	var info *ent.App
	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "Row")
	defer span.End()
	var err error
	defer func() {
		if err != nil {
			span.SetStatus(codes.Error, "db operation fail")
			span.RecordError(err)
		}
	}()
	span.SetAttributes(
		attribute.String("ID", id.String()),
	)
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		info, err = cli.App.Query().Where(app.ID(id)).Only(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}

//nolint
func setQueryConds(conds *npool.Conds, cli *ent.Client) (*ent.AppQuery, error) {
	stm := cli.App.Query()
	if conds.ID != nil {
		switch conds.GetID().GetOp() {
		case cruder.EQ:
			stm.Where(app.ID(uuid.MustParse(conds.GetID().GetValue())))
		case cruder.IN:
			var ids []uuid.UUID
			for _, val := range conds.GetIDs().GetValue() {
				id, err := uuid.Parse(val)
				if err != nil {
					return nil, err
				}
				ids = append(ids, id)
			}
			stm.Where(app.IDIn(ids...))
		default:
			return nil, fmt.Errorf("invalid app field")
		}
	}
	if conds.CreatedBy != nil {
		createdBy := uuid.MustParse(conds.GetCreatedBy().GetValue())
		switch conds.GetCreatedBy().GetOp() {
		case cruder.EQ:
			stm.Where(app.CreatedBy(createdBy))
		default:
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
			return nil, fmt.Errorf("invalid app field")
		}
	}
	return stm, nil
}

func Rows(ctx context.Context, conds *npool.Conds, offset, limit int) ([]*ent.App, int, error) {
	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "Rows")
	defer span.End()
	var err error
	defer func() {
		if err != nil {
			span.SetStatus(codes.Error, "db operation fail")
			span.RecordError(err)
		}
	}()
	span = AppCondsSpanAttributes(span, conds)
	span.SetAttributes(
		attribute.Int("Offset", offset),
		attribute.Int("Limit", limit),
	)
	rows := []*ent.App{}
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

		rows, err = stm.
			Offset(offset).
			Order(ent.Desc(app.FieldUpdatedAt)).
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

func RowOnly(ctx context.Context, conds *npool.Conds) (*ent.App, error) {
	var info *ent.App
	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "RowOnly")
	defer span.End()
	var err error
	defer func() {
		if err != nil {
			span.SetStatus(codes.Error, "db operation fail")
			span.RecordError(err)
		}
	}()
	span = AppCondsSpanAttributes(span, conds)
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
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
	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "Count")
	defer span.End()
	var err error
	defer func() {
		if err != nil {
			span.SetStatus(codes.Error, "db operation fail")
			span.RecordError(err)
		}
	}()
	span = AppCondsSpanAttributes(span, conds)
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
	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "Exist")
	defer span.End()
	var err error
	defer func() {
		if err != nil {
			span.SetStatus(codes.Error, "db operation fail")
			span.RecordError(err)
		}
	}()
	span.SetAttributes(
		attribute.String("ID", id.String()),
	)
	exist := false

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		exist, err = cli.App.Query().Where(app.ID(id)).Exist(_ctx)
		return err
	})
	if err != nil {
		return false, err
	}

	return exist, nil
}

func ExistConds(ctx context.Context, conds *npool.Conds) (bool, error) {
	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "ExistConds")
	defer span.End()
	var err error
	defer func() {
		if err != nil {
			span.SetStatus(codes.Error, "db operation fail")
			span.RecordError(err)
		}
	}()
	span = AppCondsSpanAttributes(span, conds)
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

func Delete(ctx context.Context, id uuid.UUID) (*ent.App, error) {
	var info *ent.App
	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "Delete")
	defer span.End()
	var err error
	defer func() {
		if err != nil {
			span.SetStatus(codes.Error, "db operation fail")
			span.RecordError(err)
		}
	}()
	span.SetAttributes(
		attribute.String("ID", id.String()),
	)

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		info, err = cli.App.UpdateOneID(id).
			SetDeletedAt(uint32(time.Now().Unix())).
			Save(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}
