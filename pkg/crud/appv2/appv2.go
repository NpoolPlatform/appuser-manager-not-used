package appv2

import (
	"context"
	"fmt"
	"time"

	constant "github.com/NpoolPlatform/appuser-manager/pkg/message/const"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"

	"github.com/NpoolPlatform/appuser-manager/pkg/db"
	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent"
	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent/app"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/appusermgrv2/app"
	"github.com/google/uuid"
)

func AppSpanAttributes(span trace.Span, in *npool.AppReq) trace.Span {
	span.SetAttributes(
		attribute.String("Description", in.GetDescription()),
		attribute.String("ID", in.GetID()),
		attribute.String("CreatedBy", in.GetID()),
		attribute.String("Name", in.GetCreatedBy()),
		attribute.String("Logo", in.GetName()),
		attribute.Int("CreatedAt", int(in.GetCreatedAt())),
	)
	return span
}

func AppCondsSpanAttributes(span trace.Span, in *npool.Conds) trace.Span {
	span.SetAttributes(
		attribute.String("Description.Op", in.GetDescription().GetOp()),
		attribute.String("Description.Val", in.GetDescription().GetValue()),
		attribute.String("ID.Op", in.GetID().GetOp()),
		attribute.String("ID.Val", in.GetID().GetValue()),
		attribute.String("CreatedBy.Op", in.GetID().GetOp()),
		attribute.String("CreatedBy.Val", in.GetID().GetValue()),
		attribute.String("Name.Op", in.GetCreatedBy().GetOp()),
		attribute.String("Name.Val", in.GetCreatedBy().GetValue()),
		attribute.String("Logo.Op", in.GetName().GetOp()),
		attribute.String("Logo.Val", in.GetName().GetValue()),
	)
	return span
}

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
	span = AppSpanAttributes(span, in)

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
	for key, info := range in {
		span.SetAttributes(
			attribute.String("Description"+fmt.Sprintf("%v", key), info.GetDescription()),
			attribute.String("ID"+fmt.Sprintf("%v", key), info.GetID()),
			attribute.String("CreatedBy"+fmt.Sprintf("%v", key), info.GetID()),
			attribute.String("Name"+fmt.Sprintf("%v", key), info.GetCreatedBy()),
			attribute.String("Logo"+fmt.Sprintf("%v", key), info.GetName()),
			attribute.Int("CreatedAt"+fmt.Sprintf("%v", key), int(info.GetCreatedAt())),
		)
		if err != nil {
			return nil, err
		}
	}

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
		id := uuid.MustParse(conds.GetID().GetValue())
		switch conds.GetID().GetOp() {
		case cruder.EQ:
			stm.Where(app.ID(id))
		case cruder.IN:
			stm.Where(app.IDIn(id))
		default:
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
