package appusersecret

import (
	"context"
	"fmt"
	"time"

	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent/appusersecret"

	servicename "github.com/NpoolPlatform/appuser-manager/pkg/servicename"
	commontracer "github.com/NpoolPlatform/appuser-manager/pkg/tracer"
	tracer "github.com/NpoolPlatform/appuser-manager/pkg/tracer/appusersecret"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/codes"

	"github.com/NpoolPlatform/appuser-manager/pkg/db"
	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/appuser/mgr/v2/appusersecret"
	"github.com/google/uuid"
)

func Create(ctx context.Context, in *npool.AppUserSecretReq) (*ent.AppUserSecret, error) {
	var info *ent.AppUserSecret
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
		c := CreateSet(cli.AppUserSecret.Create(), in)
		info, err = c.Save(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}

func CreateSet(c *ent.AppUserSecretCreate, in *npool.AppUserSecretReq) *ent.AppUserSecretCreate {
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
	if in.GoogleSecret != nil {
		c.SetGoogleSecret(in.GetGoogleSecret())
	}
	return c
}

//nolint:nolintlint,gocognit
func CreateBulk(ctx context.Context, in []*npool.AppUserSecretReq) ([]*ent.AppUserSecret, error) {
	var err error
	rows := []*ent.AppUserSecret{}

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
		bulk := make([]*ent.AppUserSecretCreate, len(in))
		for i, info := range in {
			bulk[i] = CreateSet(tx.AppUserSecret.Create(), info)
		}
		rows, err = tx.AppUserSecret.CreateBulk(bulk...).Save(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return rows, nil
}

func UpdateSet(info *ent.AppUserSecret, in *npool.AppUserSecretReq) *ent.AppUserSecretUpdateOne {
	u := info.Update()

	if in.PasswordHash != nil {
		u.SetPasswordHash(in.GetPasswordHash())
	}
	if in.Salt != nil {
		u.SetSalt(in.GetSalt())
	}
	if in.GoogleSecret != nil {
		u.SetGoogleSecret(in.GetGoogleSecret())
	}

	return u
}

func Update(ctx context.Context, in *npool.AppUserSecretReq) (*ent.AppUserSecret, error) {
	var err error
	var info *ent.AppUserSecret

	_, span := otel.Tracer(servicename.ServiceDomain).Start(ctx, "Update")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(codes.Error, "db operation fail")
			span.RecordError(err)
		}
	}()

	span = tracer.Trace(span, in)

	err = db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		info, err = tx.AppUserSecret.Query().Where(appusersecret.ID(uuid.MustParse(in.GetID()))).ForUpdate().Only(_ctx)
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

func Row(ctx context.Context, id uuid.UUID) (*ent.AppUserSecret, error) {
	var info *ent.AppUserSecret
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
		info, err = cli.AppUserSecret.Query().Where(appusersecret.ID(id)).Only(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}

//nolint:nolintlint,gocyclo
func SetQueryConds(conds *npool.Conds, cli *ent.Client) (*ent.AppUserSecretQuery, error) {
	stm := cli.AppUserSecret.Query()

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
			stm.Where(appusersecret.ID(id))
		default:
			return nil, fmt.Errorf("invalid appusersecret field")
		}
	}

	if conds.AppID != nil {
		appID, err := uuid.Parse(conds.GetAppID().GetValue())
		if err != nil {
			return nil, err
		}
		switch conds.GetAppID().GetOp() {
		case cruder.EQ:
			stm.Where(appusersecret.AppID(appID))
		default:
			return nil, fmt.Errorf("invalid appusersecret field")
		}
	}

	if conds.UserID != nil {
		userID, err := uuid.Parse(conds.GetUserID().GetValue())
		if err != nil {
			return nil, err
		}

		switch conds.GetUserID().GetOp() {
		case cruder.EQ:
			stm.Where(appusersecret.UserID(userID))
		default:
			return nil, fmt.Errorf("invalid appusersecret field")
		}
	}
	return stm, nil
}

func Rows(ctx context.Context, conds *npool.Conds, offset, limit int) ([]*ent.AppUserSecret, int, error) {
	var err error
	rows := []*ent.AppUserSecret{}
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
		exist, err = cli.AppUserSecret.Query().Where(appusersecret.ID(id)).Exist(_ctx)
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

func Delete(ctx context.Context, id uuid.UUID) (*ent.AppUserSecret, error) {
	var info *ent.AppUserSecret
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
