package appuserextrav2

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
	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent/appuserextra"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/appusermgrv2/appuserextra"
	"github.com/google/uuid"
)

func AppUserExtraSpanAttributes(span trace.Span, in *npool.AppUserExtraReq) trace.Span {
	span.SetAttributes(
		attribute.String("ID", in.GetID()),
		attribute.String("AppID", in.GetAppID()),
		attribute.String("UserID", in.GetUserID()),
		attribute.StringSlice("AddressFields", in.GetAddressFields()),
		attribute.String("Username", in.GetUsername()),
		attribute.Int("Age", int(in.GetAge())),
		attribute.String("Avatar", in.GetAvatar()),
		attribute.Int("Birthday", int(in.GetBirthday())),
		attribute.String("FirstName", in.GetFirstName()),
		attribute.String("Gender", in.GetGender()),
		attribute.String("IDNumber", in.GetIDNumber()),
		attribute.String("LastName", in.GetLastName()),
		attribute.String("Organization", in.GetOrganization()),
		attribute.String("PostalCode", in.GetPostalCode()),
	)
	return span
}

func AppUserExtraCondsSpanAttributes(span trace.Span, in *npool.Conds) trace.Span {
	span.SetAttributes(
		attribute.String("ID.Op", in.GetID().GetOp()),
		attribute.String("ID.Val", in.GetID().GetValue()),
		attribute.String("AppID.Op", in.GetAppID().GetOp()),
		attribute.String("AppID.Val", in.GetAppID().GetValue()),
		attribute.String("UserID.Op", in.GetUserID().GetOp()),
		attribute.String("UserID.Val", in.GetUserID().GetValue()),
		attribute.String("Username.Op", in.GetUsername().GetOp()),
		attribute.String("Username.Val", in.GetUsername().GetValue()),
		attribute.String("Age.Op", in.GetAge().GetOp()),
		attribute.Int("Age.Val", int(in.GetAge().GetValue())),
		attribute.String("Avatar.Op", in.GetAvatar().GetOp()),
		attribute.String("Avatar.Val", in.GetAvatar().GetValue()),
		attribute.String("Birthday.Op", in.GetBirthday().GetOp()),
		attribute.Int("Birthday.Val", int(in.GetBirthday().GetValue())),
		attribute.String("FirstName.Op", in.GetFirstName().GetOp()),
		attribute.String("FirstName.Val", in.GetFirstName().GetValue()),
		attribute.String("Gender.Op", in.GetGender().GetOp()),
		attribute.String("Gender.Val", in.GetGender().GetValue()),
		attribute.String("IDNumber.Op", in.GetIDNumber().GetOp()),
		attribute.String("IDNumber.Val", in.GetIDNumber().GetValue()),
		attribute.String("LastName.Op", in.GetLastName().GetOp()),
		attribute.String("LastName.Val", in.GetLastName().GetValue()),
		attribute.String("Organization.Op", in.GetOrganization().GetOp()),
		attribute.String("Organization.Val", in.GetOrganization().GetValue()),
		attribute.String("PostalCode.Op", in.GetPostalCode().GetOp()),
		attribute.String("PostalCode.Val", in.GetPostalCode().GetValue()),
	)
	return span
}

func Create(ctx context.Context, in *npool.AppUserExtraReq) (*ent.AppUserExtra, error) {
	var info *ent.AppUserExtra
	var err error
	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "Create")
	defer span.End()
	defer func() {
		if err != nil {
			span.SetStatus(codes.Error, "db operation fail")
			span.RecordError(err)
		}
	}()
	span = AppUserExtraSpanAttributes(span, in)

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		c := cli.AppUserExtra.Create()
		if in.FirstName != nil {
			c.SetFirstName(in.GetFirstName())
		}
		if in.AppID != nil {
			c.SetAppID(uuid.MustParse(in.GetAppID()))
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
		if in.ID != nil {
			c.SetID(uuid.MustParse(in.GetID()))
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
		if in.UserID != nil {
			c.SetUserID(uuid.MustParse(in.GetUserID()))
		}
		if in.UserID != nil {
			c.SetAddressFields(in.GetAddressFields())
		}
		info, err = c.Save(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}

func CreateBulk(ctx context.Context, in []*npool.AppUserExtraReq) ([]*ent.AppUserExtra, error) {
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
			attribute.String("ID"+fmt.Sprintf("%v", key), info.GetID()),
			attribute.String("AppID"+fmt.Sprintf("%v", key), info.GetAppID()),
			attribute.String("UserID"+fmt.Sprintf("%v", key), info.GetUserID()),
			attribute.StringSlice("AddressFields"+fmt.Sprintf("%v", key), info.GetAddressFields()),
			attribute.String("Username"+fmt.Sprintf("%v", key), info.GetUsername()),
			attribute.Int("Age"+fmt.Sprintf("%v", key), int(info.GetAge())),
			attribute.String("Avatar"+fmt.Sprintf("%v", key), info.GetAvatar()),
			attribute.Int("Birthday"+fmt.Sprintf("%v", key), int(info.GetBirthday())),
			attribute.String("FirstName"+fmt.Sprintf("%v", key), info.GetFirstName()),
			attribute.String("Gender"+fmt.Sprintf("%v", key), info.GetGender()),
			attribute.String("IDNumber"+fmt.Sprintf("%v", key), info.GetIDNumber()),
			attribute.String("LastName"+fmt.Sprintf("%v", key), info.GetLastName()),
			attribute.String("Organization"+fmt.Sprintf("%v", key), info.GetOrganization()),
			attribute.String("PostalCode"+fmt.Sprintf("%v", key), info.GetPostalCode()),
		)
		if err != nil {
			return nil, err
		}
	}
	rows := []*ent.AppUserExtra{}
	err = db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		bulk := make([]*ent.AppUserExtraCreate, len(in))
		for i, info := range in {
			bulk[i] = tx.AppUserExtra.Create()
			if info.ID != nil {
				bulk[i].SetID(uuid.MustParse(info.GetID()))
			}
			if info.AppID != nil {
				bulk[i].SetAppID(uuid.MustParse(info.GetAppID()))
			}
			if info.Avatar != nil {
				bulk[i].SetAvatar(info.GetAvatar())
			}
			if info.Organization != nil {
				bulk[i].SetOrganization(info.GetOrganization())
			}
			if info.UserID != nil {
				bulk[i].SetUserID(uuid.MustParse(info.GetUserID()))
			}
			if info.LastName != nil {
				bulk[i].SetLastName(info.GetLastName())
			}
			if info.PostalCode != nil {
				bulk[i].SetPostalCode(info.GetPostalCode())
			}
			if info.Age != nil {
				bulk[i].SetAge(info.GetAge())
			}
			if info.FirstName != nil {
				bulk[i].SetFirstName(info.GetFirstName())
			}
			if info.IDNumber != nil {
				bulk[i].SetIDNumber(info.GetIDNumber())
			}
			if info.Gender != nil {
				bulk[i].SetGender(info.GetGender())
			}
			if info.Birthday != nil {
				bulk[i].SetBirthday(info.GetBirthday())
			}
			if info.Username != nil {
				bulk[i].SetUsername(info.GetUsername())
			}
			if info.AddressFields != nil {
				bulk[i].SetAddressFields(info.GetAddressFields())
			}
		}
		rows, err = tx.AppUserExtra.CreateBulk(bulk...).Save(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}
	return rows, nil
}

func Update(ctx context.Context, in *npool.AppUserExtraReq) (*ent.AppUserExtra, error) {
	var info *ent.AppUserExtra
	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "Update")
	defer span.End()
	var err error
	defer func() {
		if err != nil {
			span.SetStatus(codes.Error, "db operation fail")
			span.RecordError(err)
		}
	}()
	span = AppUserExtraSpanAttributes(span, in)

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		u := cli.AppUserExtra.UpdateOneID(uuid.MustParse(in.GetID()))
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
		info, err = u.Save(_ctx)
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
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		info, err = cli.AppUserExtra.Query().Where(appuserextra.ID(id)).Only(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}

//nolint
func setQueryConds(conds *npool.Conds, cli *ent.Client) (*ent.AppUserExtraQuery, error) {
	stm := cli.AppUserExtra.Query()
	if conds.ID != nil {
		id := uuid.MustParse(conds.GetID().GetValue())
		switch conds.GetID().GetOp() {
		case cruder.EQ:
			stm.Where(appuserextra.ID(id))

		case cruder.IN:
			stm.Where(appuserextra.ID(id))

		default:
			return nil, fmt.Errorf("invalid appuserextra field")
		}
	}
	if conds.AppID != nil {
		appID := uuid.MustParse(conds.GetAppID().GetValue())
		switch conds.GetAppID().GetOp() {
		case cruder.EQ:
			stm.Where(appuserextra.AppID(appID))

		case cruder.IN:
			stm.Where(appuserextra.AppID(appID))

		default:
			return nil, fmt.Errorf("invalid appuserextra field")
		}
	}
	if conds.UserID != nil {
		userID := uuid.MustParse(conds.GetUserID().GetValue())
		switch conds.GetUserID().GetOp() {
		case cruder.EQ:
			stm.Where(appuserextra.UserID(userID))

		case cruder.IN:
			stm.Where(appuserextra.UserID(userID))

		default:
			return nil, fmt.Errorf("invalid appuserextra field")
		}
	}
	if conds.IDNumber != nil {
		switch conds.GetIDNumber().GetOp() {
		case cruder.EQ:
			stm.Where(appuserextra.IDNumber(conds.GetIDNumber().GetValue()))

		case cruder.IN:
			stm.Where(appuserextra.IDNumber(conds.GetIDNumber().GetValue()))

		default:
			return nil, fmt.Errorf("invalid appuserextra field")
		}
	}
	return stm, nil
}

func Rows(ctx context.Context, conds *npool.Conds, offset, limit int) ([]*ent.AppUserExtra, int, error) {
	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "Rows")
	defer span.End()
	var err error
	defer func() {
		if err != nil {
			span.SetStatus(codes.Error, "db operation fail")
			span.RecordError(err)
		}
	}()
	span = AppUserExtraCondsSpanAttributes(span, conds)
	span.SetAttributes(
		attribute.Int("Offset", offset),
		attribute.Int("Limit", limit),
	)
	rows := []*ent.AppUserExtra{}
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
	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "RowOnly")
	defer span.End()
	var err error
	defer func() {
		if err != nil {
			span.SetStatus(codes.Error, "db operation fail")
			span.RecordError(err)
		}
	}()
	span = AppUserExtraCondsSpanAttributes(span, conds)
	var info *ent.AppUserExtra
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
	span = AppUserExtraCondsSpanAttributes(span, conds)
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
		exist, err = cli.AppUserExtra.Query().Where(appuserextra.ID(id)).Exist(_ctx)
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
	span = AppUserExtraCondsSpanAttributes(span, conds)
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

func Delete(ctx context.Context, id uuid.UUID) (*ent.AppUserExtra, error) {
	var info *ent.AppUserExtra
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
