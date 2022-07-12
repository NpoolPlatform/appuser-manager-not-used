package approlev2

import (
	"context"
	"fmt"
	"time"

	constant "github.com/NpoolPlatform/appuser-manager/pkg/message/const"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"

	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent/app"

	"github.com/NpoolPlatform/appuser-manager/pkg/db"
	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent"
	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent/approle"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/appusermgrv2/approle"
	"github.com/google/uuid"
)

func AppRoleSpanAttributes(span trace.Span, in *npool.AppRoleReq) trace.Span {
	span.SetAttributes(
		attribute.String("ID", in.GetID()),
		attribute.String("AppID", in.GetAppID()),
		attribute.String("Role", in.GetRole()),
		attribute.String("Description", in.GetDescription()),
		attribute.String("CreatedBy", in.GetCreatedBy()),
		attribute.Bool("Default", in.GetDefault()),
	)
	return span
}

func AppRoleCondsSpanAttributes(span trace.Span, in *npool.Conds) trace.Span {
	span.SetAttributes(
		attribute.String("ID.Op", in.GetID().GetOp()),
		attribute.String("ID.Val", in.GetID().GetValue()),
		attribute.String("AppID.Op", in.GetAppID().GetOp()),
		attribute.String("AppID.Val", in.GetAppID().GetValue()),
		attribute.String("Role.Op", in.GetRole().GetOp()),
		attribute.String("Role.Val", in.GetRole().GetValue()),
		attribute.String("CreatedBy.Op", in.GetCreatedBy().GetOp()),
		attribute.String("CreatedBy.Val", in.GetCreatedBy().GetValue()),
		attribute.String("Default.Op", in.GetDefault().GetOp()),
		attribute.Bool("Default.Val", in.GetDefault().GetValue()),
	)
	return span
}

func Create(ctx context.Context, in *npool.AppRoleReq) (*ent.AppRole, error) {
	var info *ent.AppRole
	var err error
	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "Create")
	defer span.End()
	defer func() {
		if err != nil {
			span.SetStatus(codes.Error, "db operation fail")
			span.RecordError(err)
		}
	}()
	span = AppRoleSpanAttributes(span, in)

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		c := cli.AppRole.Create()
		if in.ID != nil {
			c.SetID(uuid.MustParse(in.GetID()))
		}
		if in.AppID != nil {
			c.SetAppID(uuid.MustParse(in.GetAppID()))
		}
		if in.CreatedBy != nil {
			c.SetCreatedBy(uuid.MustParse(in.GetCreatedBy()))
		}
		if in.Role != nil {
			c.SetRole(in.GetRole())
		}
		if in.Description != nil {
			c.SetDescription(in.GetDescription())
		}
		if in.Default != nil {
			c.SetDefault(in.GetDefault())
		}
		info, err = c.Save(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}

func CreateBulk(ctx context.Context, in []*npool.AppRoleReq) ([]*ent.AppRole, error) {
	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateBulk")
	defer span.End()
	var err error
	defer func() {
		if err != nil {
			span.SetStatus(codes.Error, "db operation fail")
			span.RecordError(err)
		}
	}()
	rows := []*ent.AppRole{}
	for key, info := range in {
		span.SetAttributes(
			attribute.String("ID"+fmt.Sprintf("%v", key), info.GetID()),
			attribute.String("AppID"+fmt.Sprintf("%v", key), info.GetAppID()),
			attribute.String("Role"+fmt.Sprintf("%v", key), info.GetRole()),
			attribute.String("Description"+fmt.Sprintf("%v", key), info.GetDescription()),
			attribute.String("CreatedBy"+fmt.Sprintf("%v", key), info.GetCreatedBy()),
			attribute.Bool("Default"+fmt.Sprintf("%v", key), info.GetDefault()),
		)
		if err != nil {
			return nil, err
		}
	}
	err = db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		bulk := make([]*ent.AppRoleCreate, len(in))
		for i, info := range in {
			bulk[i] = tx.AppRole.Create()
			if info.ID != nil {
				bulk[i].SetID(uuid.MustParse(info.GetID()))
			}
			if info.AppID != nil {
				bulk[i].SetAppID(uuid.MustParse(info.GetAppID()))
			}
			if info.CreatedBy != nil {
				bulk[i].SetCreatedBy(uuid.MustParse(info.GetCreatedBy()))
			}
			if info.Role != nil {
				bulk[i].SetRole(info.GetRole())
			}
			if info.Description != nil {
				bulk[i].SetDescription(info.GetDescription())
			}
			if info.Default != nil {
				bulk[i].SetDefault(info.GetDefault())
			}
		}
		rows, err = tx.AppRole.CreateBulk(bulk...).Save(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}
	return rows, nil
}

func Update(ctx context.Context, in *npool.AppRoleReq) (*ent.AppRole, error) {
	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "Update")
	defer span.End()
	var err error
	defer func() {
		if err != nil {
			span.SetStatus(codes.Error, "db operation fail")
			span.RecordError(err)
		}
	}()
	span = AppRoleSpanAttributes(span, in)
	var info *ent.AppRole
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		u := cli.AppRole.UpdateOneID(uuid.MustParse(in.GetID()))
		if in.Role != nil {
			u.SetRole(in.GetRole())
		}
		if in.Description != nil {
			u.SetDescription(in.GetDescription())
		}
		if in.Default != nil {
			u.SetDefault(in.GetDefault())
		}
		info, err = u.Save(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}

func Row(ctx context.Context, id uuid.UUID) (*ent.AppRole, error) {
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
	var info *ent.AppRole
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		info, err = cli.AppRole.Query().Where(approle.ID(id)).Only(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}

//nolint
func setQueryConds(conds *npool.Conds, cli *ent.Client) (*ent.AppRoleQuery, error) {
	stm := cli.AppRole.Query()
	if conds.ID != nil {
		id := uuid.MustParse(conds.GetID().GetValue())
		switch conds.GetID().GetOp() {
		case cruder.EQ:
			stm.Where(approle.ID(id))
		case cruder.IN:
			stm.Where(approle.IDIn(id))
		default:
			return nil, fmt.Errorf("invalid app role field")
		}
	}
	if conds.AppID != nil {
		switch conds.GetAppID().GetOp() {
		case cruder.EQ:
			stm.Where(approle.AppID(uuid.MustParse(conds.GetAppID().GetValue())))
		case cruder.IN:
			stm.Where(approle.AppIDIn(uuid.MustParse(conds.GetAppID().GetValue())))
		default:
			return nil, fmt.Errorf("invalid app role field")
		}
	}
	if conds.CreatedBy != nil {
		createdBy := uuid.MustParse(conds.GetID().GetValue())
		switch conds.GetCreatedBy().GetOp() {
		case cruder.EQ:
			stm.Where(approle.CreatedBy(createdBy))
		case cruder.IN:
			stm.Where(approle.CreatedByIn(createdBy))
		default:
			return nil, fmt.Errorf("invalid app role field")
		}
	}

	if conds.Role != nil {
		switch conds.GetRole().GetOp() {
		case cruder.EQ:
			stm.Where(approle.Role(conds.GetRole().GetValue()))
		case cruder.IN:
			stm.Where(approle.RoleIn(conds.GetRole().GetValue()))
		default:
			return nil, fmt.Errorf("invalid app role field")
		}
	}
	if conds.Default != nil {
		switch conds.GetDefault().GetOp() {
		case cruder.EQ:
			stm.Where(approle.Default(conds.GetDefault().GetValue()))
		case cruder.IN:
			stm.Where(approle.Default(conds.GetDefault().GetValue()))
		default:
			return nil, fmt.Errorf("invalid app role field")
		}
	}
	return stm, nil
}

func Rows(ctx context.Context, conds *npool.Conds, offset, limit int) ([]*ent.AppRole, int, error) {
	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "Rows")
	defer span.End()
	var err error
	defer func() {
		if err != nil {
			span.SetStatus(codes.Error, "db operation fail")
			span.RecordError(err)
		}
	}()
	span = AppRoleCondsSpanAttributes(span, conds)
	span.SetAttributes(
		attribute.Int("Offset", offset),
		attribute.Int("Limit", limit),
	)
	rows := []*ent.AppRole{}
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

func RowOnly(ctx context.Context, conds *npool.Conds) (*ent.AppRole, error) {
	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "RowOnly")
	defer span.End()
	var err error
	defer func() {
		if err != nil {
			span.SetStatus(codes.Error, "db operation fail")
			span.RecordError(err)
		}
	}()
	span = AppRoleCondsSpanAttributes(span, conds)
	var info *ent.AppRole

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
	span = AppRoleCondsSpanAttributes(span, conds)
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
		exist, err = cli.AppRole.Query().Where(approle.ID(id)).Exist(_ctx)
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
	span = AppRoleCondsSpanAttributes(span, conds)
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

func Delete(ctx context.Context, id uuid.UUID) (*ent.AppRole, error) {
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
	var info *ent.AppRole
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		info, err = cli.AppRole.UpdateOneID(id).
			SetDeletedAt(uint32(time.Now().Unix())).
			Save(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}
