package appuserthirdparty

import (
	"context"
	"fmt"
	"time"

	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent/appuserthirdparty"

	constant "github.com/NpoolPlatform/appuser-manager/pkg/message/const"
	commontracer "github.com/NpoolPlatform/appuser-manager/pkg/tracer"
	tracer "github.com/NpoolPlatform/appuser-manager/pkg/tracer/appuserthirdparty"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/codes"

	"github.com/NpoolPlatform/appuser-manager/pkg/db"
	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/appuser/mgr/v2/appuserthirdparty"
	"github.com/google/uuid"
)

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

	span = tracer.Trace(span, in)

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

func CreateTx(tx *ent.Tx, in *npool.AppUserThirdPartyReq) *ent.AppUserThirdPartyCreate {
	stm := tx.AppUserThirdParty.Create()
	if in.ID != nil {
		stm.SetID(uuid.MustParse(in.GetID()))
	}
	if in.AppID != nil {
		stm.SetAppID(uuid.MustParse(in.GetAppID()))
	}
	if in.UserID != nil {
		stm.SetUserID(uuid.MustParse(in.GetUserID()))
	}
	if in.ThirdPartyUserID != nil {
		stm.SetThirdPartyUserID(in.GetThirdPartyUserID())
	}
	if in.ThirdPartyID != nil {
		stm.SetThirdPartyID(in.GetThirdPartyID())
	}
	if in.ThirdPartyUsername != nil {
		stm.SetThirdPartyUsername(in.GetThirdPartyUsername())
	}
	if in.ThirdPartyUserAvatar != nil {
		stm.SetThirdPartyUserAvatar(in.GetThirdPartyUserAvatar())
	}

	return stm
}

//nolint:nolintlint,gocognit
func CreateBulk(ctx context.Context, in []*npool.AppUserThirdPartyReq) ([]*ent.AppUserThirdParty, error) {
	var err error
	rows := []*ent.AppUserThirdParty{}

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
		bulk := make([]*ent.AppUserThirdPartyCreate, len(in))
		for i, info := range in {
			bulk[i] = tx.AppUserThirdParty.Create()
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
			if info.ThirdPartyUsername != nil {
				bulk[i].SetThirdPartyUsername(info.GetThirdPartyUsername())
			}
			if info.ThirdPartyUserAvatar != nil {
				bulk[i].SetThirdPartyUserAvatar(info.GetThirdPartyUserAvatar())
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
	var err error
	var info *ent.AppUserThirdParty

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

func UpdateTx(tx *ent.Tx, in *npool.AppUserThirdPartyReq) *ent.AppUserThirdPartyUpdate {
	stm := tx.AppUserThirdParty.
		Update().
		Where(
			appuserthirdparty.AppID(uuid.MustParse(in.GetAppID())),
		)
	if in.ThirdPartyUsername != nil {
		stm.SetThirdPartyUsername(in.GetThirdPartyUsername())
	}
	if in.ThirdPartyUserAvatar != nil {
		stm.SetThirdPartyUserAvatar(in.GetThirdPartyUserAvatar())
	}
	return stm
}

func Row(ctx context.Context, id uuid.UUID) (*ent.AppUserThirdParty, error) {
	var info *ent.AppUserThirdParty
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
		info, err = cli.AppUserThirdParty.Query().Where(appuserthirdparty.ID(id)).Only(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}

//nolint:nolintlint,gocyclo
func setQueryConds(conds *npool.Conds, cli *ent.Client) (*ent.AppUserThirdPartyQuery, error) {
	stm := cli.AppUserThirdParty.Query()
	if conds.ID != nil {
		id := uuid.MustParse(conds.GetID().GetValue())
		switch conds.GetID().GetOp() {
		case cruder.EQ:
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
		default:
			return nil, fmt.Errorf("invalid appuserthirdparty field")
		}
	}
	if conds.UserID != nil {
		userID := uuid.MustParse(conds.GetUserID().GetValue())
		switch conds.GetUserID().GetOp() {
		case cruder.EQ:
			stm.Where(appuserthirdparty.UserID(userID))
		default:
			return nil, fmt.Errorf("invalid appuserthirdparty field")
		}
	}
	if conds.ThirdPartyUserID != nil {
		switch conds.GetThirdPartyUserID().GetOp() {
		case cruder.EQ:
			stm.Where(appuserthirdparty.ThirdPartyUserID(conds.GetThirdPartyUserID().GetValue()))
		default:
			return nil, fmt.Errorf("invalid appuserthirdparty field")
		}
	}
	if conds.ThirdPartyID != nil {
		switch conds.GetThirdPartyID().GetOp() {
		case cruder.EQ:
			stm.Where(appuserthirdparty.ThirdPartyID(conds.GetThirdPartyID().GetValue()))
		default:
			return nil, fmt.Errorf("invalid appuserthirdparty field")
		}
	}
	return stm, nil
}

func Rows(ctx context.Context, conds *npool.Conds, offset, limit int) ([]*ent.AppUserThirdParty, int, error) {
	var err error
	rows := []*ent.AppUserThirdParty{}
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
		exist, err = cli.AppUserThirdParty.Query().Where(appuserthirdparty.ID(id)).Exist(_ctx)
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

func Delete(ctx context.Context, id uuid.UUID) (*ent.AppUserThirdParty, error) {
	var info *ent.AppUserThirdParty
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
