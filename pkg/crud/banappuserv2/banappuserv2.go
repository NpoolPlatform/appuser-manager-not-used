package banappuserv2

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
	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent/banappuser"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/appusermgrv2/banappuser"
	"github.com/google/uuid"
)

func BanAppUserSpanAttributes(span trace.Span, in *npool.BanAppUserReq) trace.Span {
	span.SetAttributes(
		attribute.String("UserID", in.GetUserID()),
		attribute.String("Message", in.GetMessage()),
		attribute.String("ID", in.GetID()),
		attribute.String("AppID", in.GetID()),
	)
	return span
}

func BanAppUserCondsSpanAttributes(span trace.Span, in *npool.Conds) trace.Span {
	span.SetAttributes(
		attribute.String("UserID.Op", in.GetUserID().GetOp()),
		attribute.String("UserID.Val", in.GetUserID().GetValue()),
		attribute.String("Message.Op", in.GetMessage().GetOp()),
		attribute.String("Message.Val", in.GetMessage().GetValue()),
		attribute.String("ID.Op", in.GetID().GetOp()),
		attribute.String("ID.Val", in.GetID().GetValue()),
		attribute.String("AppID.Op", in.GetID().GetOp()),
		attribute.String("AppID.Val", in.GetID().GetValue()),
	)
	return span
}

func Create(ctx context.Context, in *npool.BanAppUserReq) (*ent.BanAppUser, error) {
	var info *ent.BanAppUser
	var err error
	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "Create")
	defer span.End()
	defer func() {
		if err != nil {
			span.SetStatus(codes.Error, "db operation fail")
			span.RecordError(err)
		}
	}()
	span = BanAppUserSpanAttributes(span, in)

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		c := cli.BanAppUser.Create()
		if in.Message != nil {
			c.SetMessage(in.GetMessage())
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
		info, err = c.Save(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}

func CreateBulk(ctx context.Context, in []*npool.BanAppUserReq) ([]*ent.BanAppUser, error) {
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
			attribute.String("UserID"+fmt.Sprintf("%v", key), info.GetUserID()),
			attribute.String("Message"+fmt.Sprintf("%v", key), info.GetMessage()),
			attribute.String("ID"+fmt.Sprintf("%v", key), info.GetID()),
			attribute.String("AppID"+fmt.Sprintf("%v", key), info.GetID()),
		)
		if err != nil {
			return nil, err
		}
	}
	rows := []*ent.BanAppUser{}
	err = db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		bulk := make([]*ent.BanAppUserCreate, len(in))
		for i, info := range in {
			bulk[i] = tx.BanAppUser.Create()
			if info.UserID != nil {
				bulk[i].SetUserID(uuid.MustParse(info.GetUserID()))
			}
			if info.Message != nil {
				bulk[i].SetMessage(info.GetMessage())
			}
			if info.ID != nil {
				bulk[i].SetID(uuid.MustParse(info.GetID()))
			}
			if info.AppID != nil {
				bulk[i].SetAppID(uuid.MustParse(info.GetAppID()))
			}
		}
		rows, err = tx.BanAppUser.CreateBulk(bulk...).Save(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}
	return rows, nil
}

func Update(ctx context.Context, in *npool.BanAppUserReq) (*ent.BanAppUser, error) {
	var info *ent.BanAppUser
	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "Update")
	defer span.End()
	var err error
	defer func() {
		if err != nil {
			span.SetStatus(codes.Error, "db operation fail")
			span.RecordError(err)
		}
	}()
	span = BanAppUserSpanAttributes(span, in)

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		u := cli.BanAppUser.UpdateOneID(uuid.MustParse(in.GetID()))
		if in.Message != nil {
			u.SetMessage(in.GetMessage())
		}
		info, err = u.Save(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}

func Row(ctx context.Context, id uuid.UUID) (*ent.BanAppUser, error) {
	var info *ent.BanAppUser
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
		info, err = cli.BanAppUser.Query().Where(banappuser.ID(id)).Only(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}

//nolint
func setQueryConds(conds *npool.Conds, cli *ent.Client) (*ent.BanAppUserQuery, error) {
	stm := cli.BanAppUser.Query()
	if conds.ID != nil {
		id := uuid.MustParse(conds.GetID().GetValue())
		switch conds.GetID().GetOp() {
		case cruder.EQ:
			stm.Where(banappuser.ID(id))

		case cruder.IN:
			stm.Where(banappuser.ID(id))

		default:
			return nil, fmt.Errorf("invalid banappuser field")
		}
	}
	if conds.AppID != nil {
		appID := uuid.MustParse(conds.GetAppID().GetValue())
		switch conds.GetAppID().GetOp() {
		case cruder.EQ:
			stm.Where(banappuser.AppID(appID))

		case cruder.IN:
			stm.Where(banappuser.AppID(appID))

		default:
			return nil, fmt.Errorf("invalid banappuser field")
		}
	}
	if conds.UserID != nil {
		userID := uuid.MustParse(conds.GetUserID().GetValue())
		switch conds.GetUserID().GetOp() {
		case cruder.EQ:
			stm.Where(banappuser.UserID(userID))

		case cruder.IN:
			stm.Where(banappuser.UserID(userID))

		default:
			return nil, fmt.Errorf("invalid banappuser field")
		}
	}
	return stm, nil
}

func Rows(ctx context.Context, conds *npool.Conds, offset, limit int) ([]*ent.BanAppUser, int, error) {
	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "Rows")
	defer span.End()
	var err error
	defer func() {
		if err != nil {
			span.SetStatus(codes.Error, "db operation fail")
			span.RecordError(err)
		}
	}()
	span = BanAppUserCondsSpanAttributes(span, conds)
	span.SetAttributes(
		attribute.Int("Offset", offset),
		attribute.Int("Limit", limit),
	)
	rows := []*ent.BanAppUser{}
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
			Order(ent.Desc(banappuser.FieldUpdatedAt)).
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

func RowOnly(ctx context.Context, conds *npool.Conds) (*ent.BanAppUser, error) {
	var info *ent.BanAppUser
	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "RowOnly")
	defer span.End()
	var err error
	defer func() {
		if err != nil {
			span.SetStatus(codes.Error, "db operation fail")
			span.RecordError(err)
		}
	}()
	span = BanAppUserCondsSpanAttributes(span, conds)
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
	span = BanAppUserCondsSpanAttributes(span, conds)
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
		exist, err = cli.BanAppUser.Query().Where(banappuser.ID(id)).Exist(_ctx)
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
	span = BanAppUserCondsSpanAttributes(span, conds)
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

func Delete(ctx context.Context, id uuid.UUID) (*ent.BanAppUser, error) {
	var info *ent.BanAppUser
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
		info, err = cli.BanAppUser.UpdateOneID(id).
			SetDeletedAt(uint32(time.Now().Unix())).
			Save(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}
