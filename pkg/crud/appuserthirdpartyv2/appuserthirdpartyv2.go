package appuserthirdpartyv2

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
	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent/appuserthirdparty"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/appusermgrv2/appuserthirdparty"
	"github.com/google/uuid"
)

func AppUserThirdPartySpanAttributes(span trace.Span, in *npool.AppUserThirdPartyReq) trace.Span {
	span.SetAttributes(
		attribute.String("UserID", in.GetUserID()),
		attribute.String("ThirdPartyUserID", in.GetThirdPartyUserID()),
		attribute.String("ThirdPartyID", in.GetThirdPartyID()),
		attribute.String("ThirdPartyUsername", in.GetThirdPartyUsername()),
		attribute.String("ThirdPartyUserAvatar", in.GetThirdPartyUserAvatar()),
		attribute.String("ID", in.GetID()),
		attribute.String("AppID", in.GetAppID()),
	)
	return span
}

func AppUserThirdPartyCondsSpanAttributes(span trace.Span, in *npool.Conds) trace.Span {
	span.SetAttributes(
		attribute.String("UserID.Op", in.GetUserID().GetOp()),
		attribute.String("UserID.Val", in.GetUserID().GetValue()),
		attribute.String("ThirdPartyUserID.Op", in.GetThirdPartyUserID().GetOp()),
		attribute.String("ThirdPartyUserID.Val", in.GetThirdPartyUserID().GetValue()),
		attribute.String("ThirdPartyID.Op", in.GetThirdPartyID().GetOp()),
		attribute.String("ThirdPartyID.Val", in.GetThirdPartyID().GetValue()),
		attribute.String("ThirdPartyUsername.Op", in.GetThirdPartyUsername().GetOp()),
		attribute.String("ThirdPartyUsername.Val", in.GetThirdPartyUsername().GetValue()),
		attribute.String("ThirdPartyUserAvatar.Op", in.GetThirdPartyUserAvatar().GetOp()),
		attribute.String("ThirdPartyUserAvatar.Val", in.GetThirdPartyUserAvatar().GetValue()),
		attribute.String("ID.Op", in.GetID().GetOp()),
		attribute.String("ID.Val", in.GetID().GetValue()),
		attribute.String("AppID.Op", in.GetAppID().GetOp()),
		attribute.String("AppID.Val", in.GetAppID().GetValue()),
	)
	return span
}

func Create(ctx context.Context, in *npool.AppUserThirdPartyReq) (*ent.AppUserThirdParty, error) {
	var info *ent.AppUserThirdParty
	var err error
	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "Create")
	defer span.End()
	defer func() {
		if err != nil {
			span.SetStatus(codes.Error, "db operation fail")
			span.RecordError(err)
		}
	}()
	span = AppUserThirdPartySpanAttributes(span, in)

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
			attribute.String("ThirdPartyUserID"+fmt.Sprintf("%v", key), info.GetThirdPartyUserID()),
			attribute.String("ThirdPartyID"+fmt.Sprintf("%v", key), info.GetThirdPartyID()),
			attribute.String("ThirdPartyUsername"+fmt.Sprintf("%v", key), info.GetThirdPartyUsername()),
			attribute.String("ThirdPartyUserAvatar"+fmt.Sprintf("%v", key), info.GetThirdPartyUserAvatar()),
			attribute.String("ID"+fmt.Sprintf("%v", key), info.GetID()),
			attribute.String("AppID"+fmt.Sprintf("%v", key), info.GetAppID()),
		)
		if err != nil {
			return nil, err
		}
	}
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
	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "Update")
	defer span.End()
	var err error
	defer func() {
		if err != nil {
			span.SetStatus(codes.Error, "db operation fail")
			span.RecordError(err)
		}
	}()
	span = AppUserThirdPartySpanAttributes(span, in)

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
	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "Rows")
	defer span.End()
	var err error
	defer func() {
		if err != nil {
			span.SetStatus(codes.Error, "db operation fail")
			span.RecordError(err)
		}
	}()
	span = AppUserThirdPartyCondsSpanAttributes(span, conds)
	span.SetAttributes(
		attribute.Int("Offset", offset),
		attribute.Int("Limit", limit),
	)
	rows := []*ent.AppUserThirdParty{}
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
	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "RowOnly")
	defer span.End()
	var err error
	defer func() {
		if err != nil {
			span.SetStatus(codes.Error, "db operation fail")
			span.RecordError(err)
		}
	}()
	span = AppUserThirdPartyCondsSpanAttributes(span, conds)
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
	span = AppUserThirdPartyCondsSpanAttributes(span, conds)
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
		exist, err = cli.AppUserThirdParty.Query().Where(appuserthirdparty.ID(id)).Exist(_ctx)
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
	span = AppUserThirdPartyCondsSpanAttributes(span, conds)
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
