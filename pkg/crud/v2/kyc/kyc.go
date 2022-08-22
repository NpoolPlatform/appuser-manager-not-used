package kyc

import (
	"context"
	"fmt"
	"time"

	constant "github.com/NpoolPlatform/appuser-manager/pkg/message/const"
	commontracer "github.com/NpoolPlatform/appuser-manager/pkg/tracer"
	tracer "github.com/NpoolPlatform/appuser-manager/pkg/tracer/kyc"

	npool "github.com/NpoolPlatform/message/npool/appuser/mgr/v2/kyc"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/codes"

	"github.com/NpoolPlatform/appuser-manager/pkg/db"
	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent"
	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent/kyc"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	"github.com/google/uuid"
)

func CreateSet(c *ent.KycCreate, info *npool.KycReq) *ent.KycCreate {
	if info.ID != nil {
		c.SetID(uuid.MustParse(info.GetID()))
	}
	if info.AppID != nil {
		c.SetAppID(uuid.MustParse(info.GetAppID()))
	}
	if info.UserID != nil {
		c.SetUserID(uuid.MustParse(info.GetUserID()))
	}
	if info.DocumentType != nil {
		c.SetDocumentType(info.GetDocumentType().String())
	}
	if info.IDNumber != nil {
		c.SetIDNumber(info.GetIDNumber())
	}
	if info.FrontImg != nil {
		c.SetFrontImg(info.GetFrontImg())
	}
	if info.BackImg != nil {
		c.SetBackImg(info.GetBackImg())
	}
	if info.SelfieImg != nil {
		c.SetSelfieImg(info.GetSelfieImg())
	}
	if info.EntityType != nil {
		c.SetEntityType(info.GetEntityType().String())
	}
	if info.ReviewID != nil {
		c.SetReviewID(uuid.MustParse(info.GetReviewID()))
	}
	return c
}
func Create(ctx context.Context, in *npool.KycReq) (*ent.Kyc, error) {
	var info *ent.Kyc
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
		c := CreateSet(cli.Kyc.Create(), in)
		info, err = c.Save(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}

func CreateBulk(ctx context.Context, in []*npool.KycReq) ([]*ent.Kyc, error) {
	var err error
	rows := []*ent.Kyc{}

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
		bulk := make([]*ent.KycCreate, len(in))
		for i, info := range in {
			bulk[i] = CreateSet(tx.Kyc.Create(), info)
		}
		rows, err = tx.Kyc.CreateBulk(bulk...).Save(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return rows, nil
}

func Update(ctx context.Context, in *npool.KycReq) (*ent.Kyc, error) {
	var err error
	var info *ent.Kyc

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
		info, err = tx.Kyc.Query().Where(kyc.ID(uuid.MustParse(in.GetID()))).ForUpdate().Only(_ctx)
		if err != nil {
			return err
		}
		u := UpdateSet(info, in)
		info, err = u.Save(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}

func UpdateSet(info *ent.Kyc, in *npool.KycReq) *ent.KycUpdateOne {
	u := info.Update()

	if in.DocumentType != nil {
		u.SetDocumentType(in.GetDocumentType().String())
	}
	if in.IDNumber != nil {
		u.SetIDNumber(in.GetIDNumber())
	}
	if in.FrontImg != nil {
		u.SetFrontImg(in.GetFrontImg())
	}
	if in.BackImg != nil {
		u.SetBackImg(in.GetBackImg())
	}
	if in.SelfieImg != nil {
		u.SetSelfieImg(in.GetSelfieImg())
	}
	if in.EntityType != nil {
		u.SetEntityType(in.GetEntityType().String())
	}
	if in.ReviewID != nil {
		u.SetReviewID(uuid.MustParse(in.GetReviewID()))
	}
	return u
}

func Row(ctx context.Context, id uuid.UUID) (*ent.Kyc, error) {
	var info *ent.Kyc
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
		info, err = cli.Kyc.Query().Where(kyc.ID(id)).Only(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}

//nolint:nolintlint,gocyclo
func setQueryConds(conds *npool.Conds, cli *ent.Client) (*ent.KycQuery, error) {
	stm := cli.Kyc.Query()

	if conds.ID != nil {
		switch conds.GetID().GetOp() {
		case cruder.EQ:
			stm.Where(kyc.ID(uuid.MustParse(conds.GetID().GetValue())))
		default:
			return nil, fmt.Errorf("invalid kyc field")
		}
	}

	if conds.AppID != nil {
		appID := uuid.MustParse(conds.GetAppID().GetValue())
		switch conds.GetAppID().GetOp() {
		case cruder.EQ:
			stm.Where(kyc.AppID(appID))
		default:
			return nil, fmt.Errorf("invalid kyc field")
		}
	}

	if conds.UserID != nil {
		userID := uuid.MustParse(conds.GetUserID().GetValue())
		switch conds.GetUserID().GetOp() {
		case cruder.EQ:
			stm.Where(kyc.UserID(userID))
		default:
			return nil, fmt.Errorf("invalid kyc field")
		}
	}

	if conds.ReviewID != nil {
		reviewID := uuid.MustParse(conds.GetReviewID().GetValue())
		switch conds.GetReviewID().GetOp() {
		case cruder.EQ:
			stm.Where(kyc.ReviewID(reviewID))
		default:
			return nil, fmt.Errorf("invalid kyc field")
		}
	}
	return stm, nil
}

func Rows(ctx context.Context, conds *npool.Conds, offset, limit int) ([]*ent.Kyc, int, error) {
	var err error
	rows := []*ent.Kyc{}
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
			Order(ent.Desc(kyc.FieldUpdatedAt)).
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

func RowOnly(ctx context.Context, conds *npool.Conds) (*ent.Kyc, error) {
	var info *ent.Kyc
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
		exist, err = cli.Kyc.Query().Where(kyc.ID(id)).Exist(_ctx)
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

func Delete(ctx context.Context, id uuid.UUID) (*ent.Kyc, error) {
	var info *ent.Kyc
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
		info, err = cli.Kyc.UpdateOneID(id).
			SetDeletedAt(uint32(time.Now().Unix())).
			Save(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}
