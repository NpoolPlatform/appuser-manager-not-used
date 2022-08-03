package appusercontrol

import (
	"context"
	"fmt"
	"time"

	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent/appusercontrol"

	constant "github.com/NpoolPlatform/appuser-manager/pkg/message/const"
	commontracer "github.com/NpoolPlatform/appuser-manager/pkg/tracer"
	tracer "github.com/NpoolPlatform/appuser-manager/pkg/tracer/appusercontrol"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/codes"

	"github.com/NpoolPlatform/appuser-manager/pkg/db"
	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/appuser/mgr/v2/appusercontrol"
	"github.com/google/uuid"
)

func Create(ctx context.Context, in *npool.AppUserControlReq) (*ent.AppUserControl, error) {
	var info *ent.AppUserControl
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
		c := cli.AppUserControl.Create()
		if in.ID != nil {
			c.SetID(uuid.MustParse(in.GetID()))
		}
		if in.AppID != nil {
			c.SetAppID(uuid.MustParse(in.GetAppID()))
		}
		if in.UserID != nil {
			c.SetUserID(uuid.MustParse(in.GetUserID()))
		}
		if in.SigninVerifyByGoogleAuthentication != nil {
			c.SetSigninVerifyByGoogleAuthentication(in.GetSigninVerifyByGoogleAuthentication())
		}
		if in.GoogleAuthenticationVerified != nil {
			c.SetGoogleAuthenticationVerified(in.GetGoogleAuthenticationVerified())
		}
		info, err = c.Save(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}

func CreateTx(tx *ent.Tx, in *npool.AppUserControlReq) *ent.AppUserControlCreate {
	stm := tx.AppUserControl.Create()
	if in.ID != nil {
		stm.SetID(uuid.MustParse(in.GetID()))
	}
	if in.AppID != nil {
		stm.SetAppID(uuid.MustParse(in.GetAppID()))
	}
	if in.UserID != nil {
		stm.SetUserID(uuid.MustParse(in.GetUserID()))
	}
	if in.SigninVerifyByGoogleAuthentication != nil {
		stm.SetSigninVerifyByGoogleAuthentication(in.GetSigninVerifyByGoogleAuthentication())
	}
	if in.GoogleAuthenticationVerified != nil {
		stm.SetGoogleAuthenticationVerified(in.GetGoogleAuthenticationVerified())
	}

	return stm
}

func CreateBulk(ctx context.Context, in []*npool.AppUserControlReq) ([]*ent.AppUserControl, error) {
	var err error
	rows := []*ent.AppUserControl{}

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateBulk")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(codes.Error, "db operation fail")
			span.RecordError(err)
		}
	}()

	span = tracer.TraceMany(span, in)

	err = db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		bulk := make([]*ent.AppUserControlCreate, len(in))
		for i, info := range in {
			bulk[i] = tx.AppUserControl.Create()
			if info.ID != nil {
				bulk[i].SetID(uuid.MustParse(info.GetID()))
			}
			if info.AppID != nil {
				bulk[i].SetAppID(uuid.MustParse(info.GetAppID()))
			}
			if info.UserID != nil {
				bulk[i].SetUserID(uuid.MustParse(info.GetUserID()))
			}
			if info.SigninVerifyByGoogleAuthentication != nil {
				bulk[i].SetSigninVerifyByGoogleAuthentication(info.GetSigninVerifyByGoogleAuthentication())
			}
			if info.GoogleAuthenticationVerified != nil {
				bulk[i].SetGoogleAuthenticationVerified(info.GetGoogleAuthenticationVerified())
			}
		}
		rows, err = tx.AppUserControl.CreateBulk(bulk...).Save(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return rows, nil
}

func Update(ctx context.Context, in *npool.AppUserControlReq) (*ent.AppUserControl, error) {
	var err error
	var info *ent.AppUserControl

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "Update")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(codes.Error, "db operation fail")
			span.RecordError(err)
		}
	}()

	span = tracer.Trace(span, in)

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		u := cli.AppUserControl.UpdateOneID(uuid.MustParse(in.GetID()))
		if in.SigninVerifyByGoogleAuthentication != nil {
			u.SetSigninVerifyByGoogleAuthentication(in.GetSigninVerifyByGoogleAuthentication())
		}
		if in.GoogleAuthenticationVerified != nil {
			u.SetGoogleAuthenticationVerified(in.GetGoogleAuthenticationVerified())
		}
		info, err = u.Save(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}

func UpdateTx(tx *ent.Tx, in *npool.AppUserControlReq) *ent.AppUserControlUpdate {
	stm := tx.AppUserControl.
		Update().
		Where(
			appusercontrol.AppID(uuid.MustParse(in.GetAppID())),
		)
	if in.SigninVerifyByGoogleAuthentication != nil {
		stm.SetSigninVerifyByGoogleAuthentication(in.GetSigninVerifyByGoogleAuthentication())
	}
	if in.GoogleAuthenticationVerified != nil {
		stm.SetGoogleAuthenticationVerified(in.GetGoogleAuthenticationVerified())
	}
	return stm
}

func Row(ctx context.Context, id uuid.UUID) (*ent.AppUserControl, error) {
	var info *ent.AppUserControl
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "Row")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(codes.Error, "db operation fail")
			span.RecordError(err)
		}
	}()

	span = commontracer.TraceID(span, id.String())

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		info, err = cli.AppUserControl.Query().Where(appusercontrol.ID(id)).Only(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}

//nolint:nolintlint,gocyclo
func setQueryConds(conds *npool.Conds, cli *ent.Client) (*ent.AppUserControlQuery, error) {
	stm := cli.AppUserControl.Query()

	if conds.ID != nil {
		id := uuid.MustParse(conds.GetID().GetValue())
		switch conds.GetID().GetOp() {
		case cruder.EQ:
			stm.Where(appusercontrol.ID(id))
		default:
			return nil, fmt.Errorf("invalid appusercontrol field")
		}
	}

	if conds.AppID != nil {
		appID := uuid.MustParse(conds.GetAppID().GetValue())
		switch conds.GetAppID().GetOp() {
		case cruder.EQ:
			stm.Where(appusercontrol.AppID(appID))
		default:
			return nil, fmt.Errorf("invalid appusercontrol field")
		}
	}

	if conds.UserID != nil {
		userID := uuid.MustParse(conds.GetUserID().GetValue())
		switch conds.GetUserID().GetOp() {
		case cruder.EQ:
			stm.Where(appusercontrol.UserID(userID))
		default:
			return nil, fmt.Errorf("invalid appusercontrol field")
		}
	}
	return stm, nil
}

func Rows(ctx context.Context, conds *npool.Conds, offset, limit int) ([]*ent.AppUserControl, int, error) {
	var err error
	rows := []*ent.AppUserControl{}
	var total int

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "Rows")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(codes.Error, "db operation fail")
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, conds)
	span = commontracer.TraceOffsetLimit(span, offset, limit)

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
			Order(ent.Desc(appusercontrol.FieldUpdatedAt)).
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

func RowOnly(ctx context.Context, conds *npool.Conds) (*ent.AppUserControl, error) {
	var info *ent.AppUserControl
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "RowOnly")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(codes.Error, "db operation fail")
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, conds)

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
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "Count")
	defer span.End()
	defer func() {
		if err != nil {
			span.SetStatus(codes.Error, "db operation fail")
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, conds)
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

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "Exist")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(codes.Error, "db operation fail")
			span.RecordError(err)
		}
	}()

	span = commontracer.TraceID(span, id.String())

	exist := false

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		exist, err = cli.AppUserControl.Query().Where(appusercontrol.ID(id)).Exist(_ctx)
		return err
	})
	if err != nil {
		return false, err
	}

	return exist, nil
}

func ExistConds(ctx context.Context, conds *npool.Conds) (bool, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "ExistConds")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(codes.Error, "db operation fail")
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, conds)
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

func Delete(ctx context.Context, id uuid.UUID) (*ent.AppUserControl, error) {
	var info *ent.AppUserControl
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "Delete")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(codes.Error, "db operation fail")
			span.RecordError(err)
		}
	}()

	span = commontracer.TraceID(span, id.String())

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		info, err = cli.AppUserControl.UpdateOneID(id).
			SetDeletedAt(uint32(time.Now().Unix())).
			Save(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}
