package appcontrol

import (
	"context"
	"fmt"
	"time"

	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent/appcontrol"

	constant "github.com/NpoolPlatform/appuser-manager/pkg/message/const"
	commontracer "github.com/NpoolPlatform/appuser-manager/pkg/tracer"
	tracer "github.com/NpoolPlatform/appuser-manager/pkg/tracer/appcontrol"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/codes"

	"github.com/NpoolPlatform/appuser-manager/pkg/db"
	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/appuser/mgr/v2/appcontrol"
	"github.com/google/uuid"
)

func Create(ctx context.Context, in *npool.AppControlReq) (*ent.AppControl, error) {
	var info *ent.AppControl
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
		c := CreateSet(cli.AppControl.Create(), in)
		info, err = c.Save(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}

func CreateSet(c *ent.AppControlCreate, info *npool.AppControlReq) *ent.AppControlCreate {
	if info.ID != nil {
		c.SetID(uuid.MustParse(info.GetID()))
	}
	if info.AppID != nil {
		c.SetAppID(uuid.MustParse(info.GetAppID()))
	}
	if info.SignupMethods != nil {
		c.SetSignupMethods(info.GetSignupMethods())
	}
	if info.ExternSigninMethods != nil {
		c.SetExternSigninMethods(info.GetExternSigninMethods())
	}
	if info.RecaptchaMethod != nil {
		c.SetRecaptchaMethod(info.GetRecaptchaMethod())
	}
	if info.KycEnable != nil {
		c.SetKycEnable(info.GetKycEnable())
	}
	if info.SigninVerifyEnable != nil {
		c.SetSigninVerifyEnable(info.GetSigninVerifyEnable())
	}
	if info.InvitationCodeMust != nil {
		c.SetInvitationCodeMust(info.GetInvitationCodeMust())
	}
	return c
}

//nolint:nolintlint,gocognit
func CreateBulk(ctx context.Context, in []*npool.AppControlReq) ([]*ent.AppControl, error) {
	var err error
	rows := []*ent.AppControl{}

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
		bulk := make([]*ent.AppControlCreate, len(in))
		for i, info := range in {
			bulk[i] = CreateSet(tx.AppControl.Create(), info)
		}
		rows, err = tx.AppControl.CreateBulk(bulk...).Save(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return rows, nil
}

func Update(ctx context.Context, in *npool.AppControlReq) (*ent.AppControl, error) {
	var err error
	var info *ent.AppControl

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
		u := cli.AppControl.UpdateOneID(uuid.MustParse(in.GetID()))
		if in.SignupMethods != nil {
			u.SetSignupMethods(in.GetSignupMethods())
		}
		if in.ExternSigninMethods != nil {
			u.SetExternSigninMethods(in.GetExternSigninMethods())
		}
		if in.RecaptchaMethod != nil {
			u.SetRecaptchaMethod(in.GetRecaptchaMethod())
		}
		if in.KycEnable != nil {
			u.SetKycEnable(in.GetKycEnable())
		}
		if in.SigninVerifyEnable != nil {
			u.SetSigninVerifyEnable(in.GetSigninVerifyEnable())
		}
		if in.InvitationCodeMust != nil {
			u.SetInvitationCodeMust(in.GetInvitationCodeMust())
		}
		info, err = u.Save(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}

func UpdateSet(u *ent.AppControlUpdate, in *npool.AppControlReq) *ent.AppControlUpdate {
	if in.SignupMethods != nil {
		u.SetSignupMethods(in.GetSignupMethods())
	}
	if in.ExternSigninMethods != nil {
		u.SetExternSigninMethods(in.GetExternSigninMethods())
	}
	if in.RecaptchaMethod != nil {
		u.SetRecaptchaMethod(in.GetRecaptchaMethod())
	}
	if in.KycEnable != nil {
		u.SetKycEnable(in.GetKycEnable())
	}
	if in.SigninVerifyEnable != nil {
		u.SetSigninVerifyEnable(in.GetSigninVerifyEnable())
	}
	if in.InvitationCodeMust != nil {
		u.SetInvitationCodeMust(in.GetInvitationCodeMust())
	}
	return u
}

func Row(ctx context.Context, id uuid.UUID) (*ent.AppControl, error) {
	var info *ent.AppControl
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
		info, err = cli.AppControl.Query().Where(appcontrol.ID(id)).Only(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}

//nolint:nolintlint,gocyclo
func setQueryConds(conds *npool.Conds, cli *ent.Client) (*ent.AppControlQuery, error) {
	stm := cli.AppControl.Query()

	if conds.ID != nil {
		id := uuid.MustParse(conds.GetID().GetValue())
		switch conds.GetID().GetOp() {
		case cruder.EQ:
			stm.Where(appcontrol.ID(id))
		default:
			return nil, fmt.Errorf("invalid appcontrol field")
		}
	}

	if conds.AppID != nil {
		appID := uuid.MustParse(conds.GetAppID().GetValue())
		switch conds.GetAppID().GetOp() {
		case cruder.EQ:
			stm.Where(appcontrol.AppID(appID))
		default:
			return nil, fmt.Errorf("invalid appcontrol field")
		}
	}

	if conds.ID != nil {
		ID := uuid.MustParse(conds.GetID().GetValue())
		switch conds.GetID().GetOp() {
		case cruder.EQ:
			stm.Where(appcontrol.ID(ID))
		default:
			return nil, fmt.Errorf("invalid appcontrol field")
		}
	}
	return stm, nil
}

func Rows(ctx context.Context, conds *npool.Conds, offset, limit int) ([]*ent.AppControl, int, error) {
	var err error
	rows := []*ent.AppControl{}
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
			Order(ent.Desc(appcontrol.FieldUpdatedAt)).
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

func RowOnly(ctx context.Context, conds *npool.Conds) (*ent.AppControl, error) {
	var info *ent.AppControl
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
		if ent.IsNotFound(err) {
			return nil, nil
		}
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
		exist, err = cli.AppControl.Query().Where(appcontrol.ID(id)).Exist(_ctx)
		return err
	})
	if err != nil {
		return false, err
	}

	return exist, nil
}

//nolint:nolintlint,gocyclo
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

func Delete(ctx context.Context, id uuid.UUID) (*ent.AppControl, error) {
	var info *ent.AppControl
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
		info, err = cli.AppControl.UpdateOneID(id).
			SetDeletedAt(uint32(time.Now().Unix())).
			Save(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}
