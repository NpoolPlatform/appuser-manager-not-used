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
		methods := []string{}
		for _, m := range info.GetSignupMethods() {
			methods = append(methods, m.String())
		}
		c.SetSignupMethods(methods)
	}
	if info.ExtSigninMethods != nil {
		methods := []string{}
		for _, m := range info.GetExtSigninMethods() {
			methods = append(methods, m.String())
		}
		c.SetExternSigninMethods(methods)
	}
	if info.RecaptchaMethod != nil {
		c.SetRecaptchaMethod(info.GetRecaptchaMethod().String())
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
	if info.CreateInvitationCodeWhen != nil {
		c.SetCreateInvitationCodeWhen(info.GetCreateInvitationCodeWhen().String())
	}
	if info.MaxTypedCouponsPerOrder != nil {
		c.SetMaxTypedCouponsPerOrder(info.GetMaxTypedCouponsPerOrder())
	}
	if info.Maintaining != nil {
		c.SetMaintaining(info.GetMaintaining())
	}
	CommitButtonTargets := []string{}
	if len(info.CommitButtonTargets) > 0 {
		CommitButtonTargets = info.GetCommitButtonTargets()
	}
	c.SetCommitButtonTargets(CommitButtonTargets)
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

func UpdateSet(info *ent.AppControl, in *npool.AppControlReq) *ent.AppControlUpdateOne {
	u := info.Update()

	if in.SignupMethods != nil {
		methods := []string{}
		for _, m := range in.GetSignupMethods() {
			methods = append(methods, m.String())
		}
		u.SetSignupMethods(methods)
	}
	if in.ExtSigninMethods != nil {
		methods := []string{}
		for _, m := range in.GetExtSigninMethods() {
			methods = append(methods, m.String())
		}
		u.SetExternSigninMethods(methods)
	}
	if in.RecaptchaMethod != nil {
		u.SetRecaptchaMethod(in.GetRecaptchaMethod().String())
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
	if in.CreateInvitationCodeWhen != nil {
		u.SetCreateInvitationCodeWhen(in.GetCreateInvitationCodeWhen().String())
	}
	if in.MaxTypedCouponsPerOrder != nil {
		u.SetMaxTypedCouponsPerOrder(in.GetMaxTypedCouponsPerOrder())
	}
	if in.Maintaining != nil {
		u.SetMaintaining(in.GetMaintaining())
	}
	if len(in.GetCommitButtonTargets()) > 0 {
		u.SetCommitButtonTargets(in.GetCommitButtonTargets())
	}
	return u
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

	err = db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		info, err = tx.AppControl.Query().Where(appcontrol.ID(uuid.MustParse(in.GetID()))).ForUpdate().Only(_ctx)
		if err != nil {
			return err
		}
		info, err = UpdateSet(info, in).Save(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return info, nil
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
func SetQueryConds(conds *npool.Conds, cli *ent.Client) (*ent.AppControlQuery, error) {
	stm := cli.AppControl.Query()

	if conds == nil {
		return stm, nil
	}

	if conds.ID != nil {
		id, err := uuid.Parse(conds.GetID().GetValue())
		if err != nil {
			return nil, err
		}
		switch conds.GetID().GetOp() {
		case cruder.EQ:
			stm.Where(appcontrol.ID(id))
		default:
			return nil, fmt.Errorf("invalid appcontrol field")
		}
	}

	if conds.AppID != nil {
		appID, err := uuid.Parse(conds.GetAppID().GetValue())
		if err != nil {
			return nil, err
		}
		switch conds.GetAppID().GetOp() {
		case cruder.EQ:
			stm.Where(appcontrol.AppID(appID))
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
		stm, err := SetQueryConds(conds, cli)
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
		stm, err := SetQueryConds(conds, cli)
		if err != nil {
			return err
		}

		info, err = stm.Only(_ctx)
		if err != nil {
			if ent.IsNotFound(err) {
				return nil
			}
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
		stm, err := SetQueryConds(conds, cli)
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
		stm, err := SetQueryConds(conds, cli)
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
