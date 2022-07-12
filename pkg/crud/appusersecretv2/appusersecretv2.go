package appusersecretv2

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
	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent/appusersecret"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/appusermgrv2/appusersecret"
	"github.com/google/uuid"
)

func AppUserSecretSpanAttributes(span trace.Span, in *npool.AppUserSecretReq) trace.Span {
	span.SetAttributes(
		attribute.String("Salt", in.GetSalt()),
		attribute.String("GoogleSecret", in.GetGoogleSecret()),
		attribute.String("ID", in.GetID()),
		attribute.String("AppID", in.GetAppID()),
		attribute.String("UserID", in.GetUserID()),
		attribute.String("PasswordHash", in.GetPasswordHash()),
	)
	return span
}

func AppUserSecretCondsSpanAttributes(span trace.Span, in *npool.Conds) trace.Span {
	span.SetAttributes(
		attribute.String("Salt.Op", in.GetSalt().GetOp()),
		attribute.String("Salt.Val", in.GetSalt().GetValue()),
		attribute.String("GoogleSecret.Op", in.GetGoogleSecret().GetOp()),
		attribute.String("GoogleSecret.Val", in.GetGoogleSecret().GetValue()),
		attribute.String("ID.Op", in.GetID().GetOp()),
		attribute.String("ID.Val", in.GetID().GetValue()),
		attribute.String("AppID.Op", in.GetAppID().GetOp()),
		attribute.String("AppID.Val", in.GetAppID().GetValue()),
		attribute.String("UserID.Op", in.GetUserID().GetOp()),
		attribute.String("UserID.Val", in.GetUserID().GetValue()),
		attribute.String("PasswordHash.Op", in.GetPasswordHash().GetOp()),
		attribute.String("PasswordHash.Val", in.GetPasswordHash().GetValue()),
	)
	return span
}

func Create(ctx context.Context, in *npool.AppUserSecretReq) (*ent.AppUserSecret, error) {
	var info *ent.AppUserSecret
	var err error
	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "Create")
	defer span.End()
	defer func() {
		if err != nil {
			span.SetStatus(codes.Error, "db operation fail")
			span.RecordError(err)
		}
	}()
	span = AppUserSecretSpanAttributes(span, in)

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		c := cli.AppUserSecret.Create()
		if in.GoogleSecret != nil {
			c.SetGoogleSecret(in.GetGoogleSecret())
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
		if in.PasswordHash != nil {
			c.SetPasswordHash(in.GetPasswordHash())
		}
		if in.Salt != nil {
			c.SetSalt(in.GetSalt())
		}
		info, err = c.Save(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}

func CreateBulk(ctx context.Context, in []*npool.AppUserSecretReq) ([]*ent.AppUserSecret, error) {
	rows := []*ent.AppUserSecret{}
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
			attribute.String("Salt"+fmt.Sprintf("%v", key), info.GetSalt()),
			attribute.String("GoogleSecret"+fmt.Sprintf("%v", key), info.GetGoogleSecret()),
			attribute.String("ID"+fmt.Sprintf("%v", key), info.GetID()),
			attribute.String("AppID"+fmt.Sprintf("%v", key), info.GetAppID()),
			attribute.String("UserID"+fmt.Sprintf("%v", key), info.GetUserID()),
			attribute.String("PasswordHash"+fmt.Sprintf("%v", key), info.GetPasswordHash()),
		)
		if err != nil {
			return nil, err
		}
	}
	err = db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		bulk := make([]*ent.AppUserSecretCreate, len(in))
		for i, info := range in {
			bulk[i] = tx.AppUserSecret.Create()
			if info.ID != nil {
				bulk[i].SetID(uuid.MustParse(info.GetID()))
			}
			if info.AppID != nil {
				bulk[i].SetAppID(uuid.MustParse(info.GetAppID()))
			}
			if info.UserID != nil {
				bulk[i].SetUserID(uuid.MustParse(info.GetUserID()))
			}
			if info.PasswordHash != nil {
				bulk[i].SetPasswordHash(info.GetPasswordHash())
			}
			if info.Salt != nil {
				bulk[i].SetSalt(info.GetSalt())
			}
			if info.GoogleSecret != nil {
				bulk[i].SetGoogleSecret(info.GetGoogleSecret())
			}
		}
		rows, err = tx.AppUserSecret.CreateBulk(bulk...).Save(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}
	return rows, nil
}

func Update(ctx context.Context, in *npool.AppUserSecretReq) (*ent.AppUserSecret, error) {
	var info *ent.AppUserSecret
	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "Update")
	defer span.End()
	var err error
	defer func() {
		if err != nil {
			span.SetStatus(codes.Error, "db operation fail")
			span.RecordError(err)
		}
	}()
	span = AppUserSecretSpanAttributes(span, in)

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		u := cli.AppUserSecret.UpdateOneID(uuid.MustParse(in.GetID()))
		if in.PasswordHash != nil {
			u.SetPasswordHash(in.GetPasswordHash())
		}
		if in.Salt != nil {
			u.SetSalt(in.GetSalt())
		}
		if in.GoogleSecret != nil {
			u.SetGoogleSecret(in.GetGoogleSecret())
		}
		info, err = u.Save(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}

func Row(ctx context.Context, id uuid.UUID) (*ent.AppUserSecret, error) {
	var info *ent.AppUserSecret
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
		info, err = cli.AppUserSecret.Query().Where(appusersecret.ID(id)).Only(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}

//nolint
func setQueryConds(conds *npool.Conds, cli *ent.Client) (*ent.AppUserSecretQuery, error) {
	stm := cli.AppUserSecret.Query()
	if conds.ID != nil {
		id := uuid.MustParse(conds.GetID().GetValue())
		switch conds.GetID().GetOp() {
		case cruder.EQ:
			stm.Where(appusersecret.ID(id))

		case cruder.IN:
			stm.Where(appusersecret.ID(id))

		default:
			return nil, fmt.Errorf("invalid appusersecret field")
		}
	}
	if conds.AppID != nil {
		appID := uuid.MustParse(conds.GetAppID().GetValue())
		switch conds.GetAppID().GetOp() {
		case cruder.EQ:
			stm.Where(appusersecret.AppID(appID))

		case cruder.IN:
			stm.Where(appusersecret.AppID(appID))

		default:
			return nil, fmt.Errorf("invalid appusersecret field")
		}
	}
	if conds.UserID != nil {
		userID := uuid.MustParse(conds.GetUserID().GetValue())
		switch conds.GetUserID().GetOp() {
		case cruder.EQ:
			stm.Where(appusersecret.UserID(userID))

		case cruder.IN:
			stm.Where(appusersecret.UserID(userID))

		default:
			return nil, fmt.Errorf("invalid appusersecret field")
		}
	}
	return stm, nil
}

func Rows(ctx context.Context, conds *npool.Conds, offset, limit int) ([]*ent.AppUserSecret, int, error) {
	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "Rows")
	defer span.End()
	var err error
	defer func() {
		if err != nil {
			span.SetStatus(codes.Error, "db operation fail")
			span.RecordError(err)
		}
	}()
	span = AppUserSecretCondsSpanAttributes(span, conds)
	span.SetAttributes(
		attribute.Int("Offset", offset),
		attribute.Int("Limit", limit),
	)
	rows := []*ent.AppUserSecret{}
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
			Order(ent.Desc(appusersecret.FieldUpdatedAt)).
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

func RowOnly(ctx context.Context, conds *npool.Conds) (*ent.AppUserSecret, error) {
	var info *ent.AppUserSecret
	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "RowOnly")
	defer span.End()
	var err error
	defer func() {
		if err != nil {
			span.SetStatus(codes.Error, "db operation fail")
			span.RecordError(err)
		}
	}()
	span = AppUserSecretCondsSpanAttributes(span, conds)
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
	span = AppUserSecretCondsSpanAttributes(span, conds)
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
		exist, err = cli.AppUserSecret.Query().Where(appusersecret.ID(id)).Exist(_ctx)
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
	span = AppUserSecretCondsSpanAttributes(span, conds)
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

func Delete(ctx context.Context, id uuid.UUID) (*ent.AppUserSecret, error) {
	var info *ent.AppUserSecret
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
		info, err = cli.AppUserSecret.UpdateOneID(id).
			SetDeletedAt(uint32(time.Now().Unix())).
			Save(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}
