package history

import (
	"context"
	"fmt"
	"time"

	servicename "github.com/NpoolPlatform/appuser-manager/pkg/servicename"
	commontracer "github.com/NpoolPlatform/appuser-manager/pkg/tracer"
	tracer "github.com/NpoolPlatform/appuser-manager/pkg/tracer/authing/history"

	npool "github.com/NpoolPlatform/message/npool/appuser/mgr/v2/authing/history"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/codes"

	"github.com/NpoolPlatform/appuser-manager/pkg/db"
	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent"
	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent/authhistory"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	"github.com/google/uuid"
)

func CreateSet(c *ent.AuthHistoryCreate, info *npool.HistoryReq) *ent.AuthHistoryCreate {
	if info.ID != nil {
		c.SetID(uuid.MustParse(info.GetID()))
	}
	if info.AppID != nil {
		c.SetAppID(uuid.MustParse(info.GetAppID()))
	}
	if info.UserID != nil {
		c.SetUserID(uuid.MustParse(info.GetUserID()))
	}
	if info.Resource != nil {
		c.SetResource(info.GetResource())
	}
	if info.Method != nil {
		c.SetMethod(info.GetMethod())
	}
	if info.Allowed != nil {
		c.SetAllowed(info.GetAllowed())
	}
	return c
}
func Create(ctx context.Context, in *npool.HistoryReq) (*ent.AuthHistory, error) {
	var info *ent.AuthHistory
	var err error

	_, span := otel.Tracer(servicename.ServiceDomain).Start(ctx, "Create")
	defer span.End()
	defer func() {
		if err != nil {
			span.SetStatus(codes.Error, "db operation fail")
			span.RecordError(err)
		}
	}()

	span = tracer.Trace(span, in)

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		c := CreateSet(cli.AuthHistory.Create(), in)
		info, err = c.Save(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}

func CreateBulk(ctx context.Context, in []*npool.HistoryReq) ([]*ent.AuthHistory, error) {
	var err error
	rows := []*ent.AuthHistory{}

	_, span := otel.Tracer(servicename.ServiceDomain).Start(ctx, "CreateBulk")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(codes.Error, "db operation fail")
			span.RecordError(err)
		}
	}()

	span = tracer.TraceMany(span, in)

	err = db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		bulk := make([]*ent.AuthHistoryCreate, len(in))
		for i, info := range in {
			bulk[i] = CreateSet(tx.AuthHistory.Create(), info)
		}
		rows, err = tx.AuthHistory.CreateBulk(bulk...).Save(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return rows, nil
}

func Update(ctx context.Context, in *npool.HistoryReq) (*ent.AuthHistory, error) {
	var err error
	var info *ent.AuthHistory

	_, span := otel.Tracer(servicename.ServiceDomain).Start(ctx, "Update")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(codes.Error, "db operation fail")
			span.RecordError(err)
		}
	}()

	span = tracer.Trace(span, in)

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		u := UpdateSet(cli.AuthHistory.UpdateOneID(uuid.MustParse(in.GetID())), in)
		info, err = u.Save(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}

func UpdateSet(u *ent.AuthHistoryUpdateOne, in *npool.HistoryReq) *ent.AuthHistoryUpdateOne {
	return u
}

func Row(ctx context.Context, id uuid.UUID) (*ent.AuthHistory, error) {
	var info *ent.AuthHistory
	var err error

	_, span := otel.Tracer(servicename.ServiceDomain).Start(ctx, "Row")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(codes.Error, "db operation fail")
			span.RecordError(err)
		}
	}()

	span = commontracer.TraceID(span, id.String())

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		info, err = cli.AuthHistory.Query().Where(authhistory.ID(id)).Only(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}

//nolint:nolintlint,gocyclo
func SetQueryConds(conds *npool.Conds, cli *ent.Client) (*ent.AuthHistoryQuery, error) {
	stm := cli.AuthHistory.Query()

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
			stm.Where(authhistory.ID(id))
		default:
			return nil, fmt.Errorf("invalid app field")
		}
	}

	if conds.AppID != nil {
		appID, err := uuid.Parse(conds.GetAppID().GetValue())
		if err != nil {
			return nil, err
		}
		switch conds.GetAppID().GetOp() {
		case cruder.EQ:
			stm.Where(authhistory.AppID(appID))
		default:
			return nil, fmt.Errorf("invalid app field")
		}
	}

	if conds.UserID != nil {
		userID, err := uuid.Parse(conds.GetUserID().GetValue())
		if err != nil {
			return nil, err
		}
		switch conds.GetUserID().GetOp() {
		case cruder.EQ:
			stm.Where(authhistory.UserID(userID))
		default:
			return nil, fmt.Errorf("invalid app field")
		}
	}

	if conds.Resource != nil {
		switch conds.GetResource().GetOp() {
		case cruder.EQ:
			stm.Where(authhistory.Resource(conds.GetResource().GetValue()))
		default:
			return nil, fmt.Errorf("invalid app field")
		}
	}

	if conds.Method != nil {
		switch conds.GetMethod().GetOp() {
		case cruder.EQ:
			stm.Where(authhistory.Method(conds.GetMethod().GetValue()))
		default:
			return nil, fmt.Errorf("invalid app field")
		}
	}

	if conds.Allowed != nil {
		switch conds.GetAllowed().GetOp() {
		case cruder.EQ:
			stm.Where(authhistory.Allowed(conds.GetAllowed().GetValue()))
		default:
			return nil, fmt.Errorf("invalid app field")
		}
	}

	return stm, nil
}

func Rows(ctx context.Context, conds *npool.Conds, offset, limit int) ([]*ent.AuthHistory, int, error) {
	var err error
	rows := []*ent.AuthHistory{}
	var total int

	_, span := otel.Tracer(servicename.ServiceDomain).Start(ctx, "Rows")
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
			Order(ent.Desc(authhistory.FieldUpdatedAt)).
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

func RowOnly(ctx context.Context, conds *npool.Conds) (*ent.AuthHistory, error) {
	var info *ent.AuthHistory
	var err error

	_, span := otel.Tracer(servicename.ServiceDomain).Start(ctx, "RowOnly")
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

	_, span := otel.Tracer(servicename.ServiceDomain).Start(ctx, "Count")
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

	_, span := otel.Tracer(servicename.ServiceDomain).Start(ctx, "Exist")
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
		exist, err = cli.AuthHistory.Query().Where(authhistory.ID(id)).Exist(_ctx)
		return err
	})
	if err != nil {
		return false, err
	}

	return exist, nil
}

func ExistConds(ctx context.Context, conds *npool.Conds) (bool, error) {
	var err error

	_, span := otel.Tracer(servicename.ServiceDomain).Start(ctx, "ExistConds")
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

func Delete(ctx context.Context, id uuid.UUID) (*ent.AuthHistory, error) {
	var info *ent.AuthHistory
	var err error

	_, span := otel.Tracer(servicename.ServiceDomain).Start(ctx, "Delete")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(codes.Error, "db operation fail")
			span.RecordError(err)
		}
	}()

	span = commontracer.TraceID(span, id.String())

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		info, err = cli.AuthHistory.UpdateOneID(id).
			SetDeletedAt(uint32(time.Now().Unix())).
			Save(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}
