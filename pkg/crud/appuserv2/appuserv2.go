package appuserv2

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
	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent/appuser"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/appusermgrv2/appuser"
	"github.com/google/uuid"
)

func AppUserSpanAttributes(span trace.Span, in *npool.AppUserReq) trace.Span {
	span.SetAttributes(
		attribute.String("PhoneNo", in.GetPhoneNo()),
		attribute.String("ImportFromApp", in.GetImportFromApp()),
		attribute.String("ID", in.GetID()),
		attribute.String("AppID", in.GetAppID()),
		attribute.String("EmailAddress", in.GetEmailAddress()),
	)
	return span
}

func AppUserCondsSpanAttributes(span trace.Span, in *npool.Conds) trace.Span {
	span.SetAttributes(
		attribute.String("PhoneNo.Op", in.GetPhoneNo().GetOp()),
		attribute.String("PhoneNo.Val", in.GetPhoneNo().GetValue()),
		attribute.String("ImportFromApp.Op", in.GetImportFromApp().GetOp()),
		attribute.String("ImportFromApp.Val", in.GetImportFromApp().GetValue()),
		attribute.String("ID.Op", in.GetID().GetOp()),
		attribute.String("ID.Val", in.GetID().GetValue()),
		attribute.String("AppID.Op", in.GetAppID().GetOp()),
		attribute.String("AppID.Val", in.GetAppID().GetValue()),
		attribute.String("EmailAddress.Op", in.GetEmailAddress().GetOp()),
		attribute.String("EmailAddress.Val", in.GetEmailAddress().GetValue()),
	)
	return span
}

func Create(ctx context.Context, in *npool.AppUserReq) (*ent.AppUser, error) {
	var info *ent.AppUser
	var err error
	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "Create")
	defer span.End()
	defer func() {
		if err != nil {
			span.SetStatus(codes.Error, "db operation fail")
			span.RecordError(err)
		}
	}()
	span = AppUserSpanAttributes(span, in)

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
			attribute.String("PhoneNo"+fmt.Sprintf("%v", key), info.GetPhoneNo()),
			attribute.String("ImportFromApp"+fmt.Sprintf("%v", key), info.GetImportFromApp()),
			attribute.String("ID"+fmt.Sprintf("%v", key), info.GetID()),
			attribute.String("AppID"+fmt.Sprintf("%v", key), info.GetAppID()),
			attribute.String("EmailAddress"+fmt.Sprintf("%v", key), info.GetEmailAddress()),
		)
		if err != nil {
			return nil, err
		}
	}
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
	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "Update")
	defer span.End()
	var err error
	defer func() {
		if err != nil {
			span.SetStatus(codes.Error, "db operation fail")
			span.RecordError(err)
		}
	}()
	span = AppUserSpanAttributes(span, in)
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
	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "Rows")
	defer span.End()
	var err error
	defer func() {
		if err != nil {
			span.SetStatus(codes.Error, "db operation fail")
			span.RecordError(err)
		}
	}()
	span = AppUserCondsSpanAttributes(span, conds)
	span.SetAttributes(
		attribute.Int("Offset", offset),
		attribute.Int("Limit", limit),
	)
	rows := []*ent.AppUser{}
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
	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "RowOnly")
	defer span.End()
	var err error
	defer func() {
		if err != nil {
			span.SetStatus(codes.Error, "db operation fail")
			span.RecordError(err)
		}
	}()
	span = AppUserCondsSpanAttributes(span, conds)
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
	span = AppUserCondsSpanAttributes(span, conds)
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
		exist, err = cli.AppUser.Query().Where(appuser.ID(id)).Exist(_ctx)
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
	span = AppUserCondsSpanAttributes(span, conds)
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
