package appuserextra

import (
	"context"
	"fmt"
	"time"

	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent/appuserextra"

	servicename "github.com/NpoolPlatform/appuser-manager/pkg/servicename"
	commontracer "github.com/NpoolPlatform/appuser-manager/pkg/tracer"
	tracer "github.com/NpoolPlatform/appuser-manager/pkg/tracer/appuserextra"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/codes"

	"github.com/NpoolPlatform/appuser-manager/pkg/db"
	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/appuser/mgr/v2/appuserextra"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

//nolint:nolintlint,gocyclo
func Create(ctx context.Context, in *npool.AppUserExtraReq) (*ent.AppUserExtra, error) {
	var info *ent.AppUserExtra
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
		c := CreateSet(cli.AppUserExtra.Create(), in)
		info, err = c.Save(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}

func CreateSet(c *ent.AppUserExtraCreate, in *npool.AppUserExtraReq) *ent.AppUserExtraCreate { //nolint
	if in.ID != nil {
		c.SetID(uuid.MustParse(in.GetID()))
	}
	if in.AppID != nil {
		c.SetAppID(uuid.MustParse(in.GetAppID()))
	}
	if in.UserID != nil {
		c.SetUserID(uuid.MustParse(in.GetUserID()))
	}
	if in.FirstName != nil {
		c.SetFirstName(in.GetFirstName())
	}
	if in.Organization != nil {
		c.SetOrganization(in.GetOrganization())
	}
	if in.IDNumber != nil {
		c.SetIDNumber(in.GetIDNumber())
	}
	if in.PostalCode != nil {
		c.SetPostalCode(in.GetPostalCode())
	}
	if in.Age != nil {
		c.SetAge(in.GetAge())
	}
	if in.Birthday != nil {
		c.SetBirthday(in.GetBirthday())
	}
	if in.Avatar != nil {
		c.SetAvatar(in.GetAvatar())
	}
	if in.Username != nil {
		c.SetUsername(in.GetUsername())
	}
	if in.LastName != nil {
		c.SetLastName(in.GetLastName())
	}
	if in.Gender != nil {
		c.SetGender(in.GetGender())
	}
	if in.AddressFields != nil {
		c.SetAddressFields(in.GetAddressFields())
	}
	if in.ActionCredits != nil {
		c.SetActionCredits(decimal.RequireFromString(in.GetActionCredits()))
	}
	return c
}

//nolint:nolintlint,gocyclo
func CreateBulk(ctx context.Context, in []*npool.AppUserExtraReq) ([]*ent.AppUserExtra, error) {
	var err error
	rows := []*ent.AppUserExtra{}

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
		bulk := make([]*ent.AppUserExtraCreate, len(in))
		for i, info := range in {
			bulk[i] = CreateSet(tx.AppUserExtra.Create(), info)
		}
		rows, err = tx.AppUserExtra.CreateBulk(bulk...).Save(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return rows, nil
}

func UpdateSet(info *ent.AppUserExtra, in *npool.AppUserExtraReq) *ent.AppUserExtraUpdateOne {
	u := info.Update()

	if in.Username != nil {
		u.SetUsername(in.GetUsername())
	}
	if in.FirstName != nil {
		u.SetFirstName(in.GetFirstName())
	}
	if in.LastName != nil {
		u.SetLastName(in.GetLastName())
	}
	if in.AddressFields != nil {
		u.SetAddressFields(in.GetAddressFields())
	}
	if in.Gender != nil {
		u.SetGender(in.GetGender())
	}
	if in.PostalCode != nil {
		u.SetPostalCode(in.GetPostalCode())
	}
	if in.IDNumber != nil {
		u.SetIDNumber(in.GetIDNumber())
	}
	if in.Organization != nil {
		u.SetOrganization(in.GetOrganization())
	}
	if in.Age != nil {
		u.SetAge(in.GetAge())
	}
	if in.Birthday != nil {
		u.SetBirthday(in.GetBirthday())
	}
	if in.Avatar != nil {
		u.SetAvatar(in.GetAvatar())
	}
	if in.LastName != nil {
		u.SetLastName(in.GetLastName())
	}
	if in.ActionCredits != nil {
		credits := info.ActionCredits.
			Add(decimal.RequireFromString(in.GetActionCredits()))
		u.SetActionCredits(credits)
	}

	return u
}

func Update(ctx context.Context, in *npool.AppUserExtraReq) (*ent.AppUserExtra, error) {
	var err error
	var info *ent.AppUserExtra

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
		info, err = tx.AppUserExtra.Query().Where(appuserextra.ID(uuid.MustParse(in.GetID()))).ForUpdate().Only(_ctx)
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

func Row(ctx context.Context, id uuid.UUID) (*ent.AppUserExtra, error) {
	var info *ent.AppUserExtra
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
		info, err = cli.AppUserExtra.Query().Where(appuserextra.ID(id)).Only(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}

//nolint:nolintlint,gocyclo
func SetQueryConds(conds *npool.Conds, cli *ent.Client) (*ent.AppUserExtraQuery, error) {
	stm := cli.AppUserExtra.Query()

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
			stm.Where(appuserextra.ID(id))
		default:
			return nil, fmt.Errorf("invalid appuserextra field")
		}
	}

	if conds.AppID != nil {
		appID, err := uuid.Parse(conds.GetAppID().GetValue())
		if err != nil {
			return nil, err
		}

		switch conds.GetAppID().GetOp() {
		case cruder.EQ:
			stm.Where(appuserextra.AppID(appID))
		default:
			return nil, fmt.Errorf("invalid appuserextra field")
		}
	}

	if conds.UserID != nil {
		userID, err := uuid.Parse(conds.GetUserID().GetValue())
		if err != nil {
			return nil, err
		}
		switch conds.GetUserID().GetOp() {
		case cruder.EQ:
			stm.Where(appuserextra.UserID(userID))
		default:
			return nil, fmt.Errorf("invalid appuserextra field")
		}
	}

	if conds.IDNumber != nil {
		switch conds.GetIDNumber().GetOp() {
		case cruder.EQ:
			stm.Where(appuserextra.IDNumber(conds.GetIDNumber().GetValue()))
		default:
			return nil, fmt.Errorf("invalid appuserextra field")
		}
	}
	return stm, nil
}

func Rows(ctx context.Context, conds *npool.Conds, offset, limit int) ([]*ent.AppUserExtra, int, error) {
	var err error
	rows := []*ent.AppUserExtra{}
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
			Order(ent.Desc(appuserextra.FieldUpdatedAt)).
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

func RowOnly(ctx context.Context, conds *npool.Conds) (*ent.AppUserExtra, error) {
	var info *ent.AppUserExtra
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
		exist, err = cli.AppUserExtra.Query().Where(appuserextra.ID(id)).Exist(_ctx)
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

func Delete(ctx context.Context, id uuid.UUID) (*ent.AppUserExtra, error) {
	var info *ent.AppUserExtra
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
		info, err = cli.AppUserExtra.UpdateOneID(id).
			SetDeletedAt(uint32(time.Now().Unix())).
			Save(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}
