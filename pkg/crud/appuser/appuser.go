package appuser

import (
	"context"
	"fmt"
	"time"

	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent/appuser"

	servicename "github.com/NpoolPlatform/appuser-manager/pkg/servicename"
	commontracer "github.com/NpoolPlatform/appuser-manager/pkg/tracer"
	tracer "github.com/NpoolPlatform/appuser-manager/pkg/tracer/appuser"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/codes"

	"github.com/NpoolPlatform/appuser-manager/pkg/db"
	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/appuser/mgr/v2/appuser"
	"github.com/google/uuid"
)

func Create(ctx context.Context, in *npool.AppUserReq) (*ent.AppUser, error) {
	var info *ent.AppUser
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
		c := CreateSet(cli.AppUser.Create(), in)
		info, err = c.Save(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}

func CreateSet(c *ent.AppUserCreate, in *npool.AppUserReq) *ent.AppUserCreate {
	if in.ID != nil {
		c.SetID(uuid.MustParse(in.GetID()))
	}
	if in.AppID != nil {
		c.SetAppID(uuid.MustParse(in.GetAppID()))
	}
	if in.EmailAddress != nil {
		c.SetEmailAddress(in.GetEmailAddress())
	}
	if in.PhoneNO != nil {
		c.SetPhoneNo(in.GetPhoneNO())
	}
	if in.ImportFromApp != nil {
		c.SetImportFromApp(uuid.MustParse(in.GetImportFromApp()))
	}
	return c
}

func CreateBulk(ctx context.Context, in []*npool.AppUserReq) ([]*ent.AppUser, error) {
	var err error
	rows := []*ent.AppUser{}

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
		bulk := make([]*ent.AppUserCreate, len(in))
		for i, info := range in {
			bulk[i] = CreateSet(tx.AppUser.Create(), info)
		}
		rows, err = tx.AppUser.CreateBulk(bulk...).Save(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return rows, nil
}

func UpdateSet(u *ent.AppUserUpdateOne, in *npool.AppUserReq) *ent.AppUserUpdateOne {
	if in.EmailAddress != nil {
		u.SetEmailAddress(in.GetEmailAddress())
	}
	if in.PhoneNO != nil {
		u.SetPhoneNo(in.GetPhoneNO())
	}
	return u
}

func Update(ctx context.Context, in *npool.AppUserReq) (*ent.AppUser, error) {
	var err error
	var info *ent.AppUser

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
		u := cli.AppUser.UpdateOneID(uuid.MustParse(in.GetID()))
		info, err = UpdateSet(u, in).Save(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}

func Row(ctx context.Context, id uuid.UUID) (*ent.AppUser, error) {
	var info *ent.AppUser
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
		info, err = cli.AppUser.Query().Where(appuser.ID(id)).Only(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}

//nolint:nolintlint,gocyclo
func SetQueryConds(conds *npool.Conds, cli *ent.Client) (*ent.AppUserQuery, error) {
	stm := cli.AppUser.Query()

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
			stm.Where(appuser.ID(id))
		default:
			return nil, fmt.Errorf("invalid appuser field")
		}
	}

	if conds.IDs != nil {
		if conds.GetIDs().GetOp() == cruder.IN {
			var ids []uuid.UUID
			for _, val := range conds.GetIDs().GetValue() {
				id, err := uuid.Parse(val)
				if err != nil {
					return nil, err
				}
				ids = append(ids, id)
			}
			stm.Where(appuser.IDIn(ids...))
		}
	}

	if conds.AppID != nil {
		appID, err := uuid.Parse(conds.GetAppID().GetValue())
		if err != nil {
			return nil, err
		}
		switch conds.GetAppID().GetOp() {
		case cruder.EQ:
			stm.Where(appuser.AppID(appID))
		default:
			return nil, fmt.Errorf("invalid appuser field")
		}
	}

	if conds.EmailAddress != nil {
		switch conds.GetEmailAddress().GetOp() {
		case cruder.EQ:
			stm.Where(appuser.EmailAddress(conds.GetEmailAddress().GetValue()))
		default:
			return nil, fmt.Errorf("invalid appuser field")
		}
	}

	if conds.PhoneNO != nil {
		switch conds.GetPhoneNO().GetOp() {
		case cruder.EQ:
			stm.Where(appuser.PhoneNo(conds.GetPhoneNO().GetValue()))
		default:
			return nil, fmt.Errorf("invalid appuser field")
		}
	}

	if conds.ImportFromApp != nil {
		importFromApp := uuid.MustParse(conds.GetImportFromApp().GetValue())
		switch conds.GetImportFromApp().GetOp() {
		case cruder.EQ:
			stm.Where(appuser.ImportFromApp(importFromApp))
		default:
			return nil, fmt.Errorf("invalid appuser field")
		}
	}

	return stm, nil
}

func Rows(ctx context.Context, conds *npool.Conds, offset, limit int) ([]*ent.AppUser, int, error) {
	var err error
	rows := []*ent.AppUser{}
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
		exist, err = cli.AppUser.Query().Where(appuser.ID(id)).Exist(_ctx)
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

func Delete(ctx context.Context, id uuid.UUID) (*ent.AppUser, error) {
	var info *ent.AppUser
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
