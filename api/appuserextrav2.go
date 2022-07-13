//go:build !codeanalysis
// +build !codeanalysis

package api

import (
	"context"
	"fmt"

	crud "github.com/NpoolPlatform/appuser-manager/pkg/crud/appuserextrav2"
	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent"
	constant "github.com/NpoolPlatform/appuser-manager/pkg/message/const"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	scodes "go.opentelemetry.io/otel/codes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/appusermgrv2/appuserextra"

	"github.com/google/uuid"
)

func checkAppUserExtraInfo(info *npool.AppUserExtraReq) error {
	if _, err := uuid.Parse(info.GetAppID()); err != nil {
		logger.Sugar().Error("AppID is invalid")
		return status.Error(codes.InvalidArgument, "AppID is invalid")
	}
	if _, err := uuid.Parse(info.GetUserID()); err != nil {
		logger.Sugar().Error("UserID is invalid")
		return status.Error(codes.InvalidArgument, "UserID is invalid")
	}
	return nil
}

func appUserExtraRowToObject(row *ent.AppUserExtra) *npool.AppUserExtra {
	return &npool.AppUserExtra{
		PostalCode:    row.PostalCode,
		Avatar:        row.Avatar,
		Organization:  row.Organization,
		Birthday:      row.Birthday,
		ID:            row.ID.String(),
		AppID:         row.AppID.String(),
		Username:      row.Username,
		Gender:        row.Gender,
		LastName:      row.LastName,
		Age:           row.Age,
		UserID:        row.UserID.String(),
		FirstName:     row.FirstName,
		IDNumber:      row.IDNumber,
		AddressFields: row.AddressFields,
	}
}

func (s *AppUserExtraServer) CreateAppUserExtraV2(ctx context.Context, in *npool.CreateAppUserExtraRequest) (*npool.CreateAppUserExtraResponse, error) {
	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateAppUserExtraV2")
	defer span.End()
	var err error
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()
	span = crud.AppUserExtraSpanAttributes(span, in.GetInfo())
	err = checkAppUserExtraInfo(in.GetInfo())
	if err != nil {
		return &npool.CreateAppUserExtraResponse{}, err
	}
	span.AddEvent("call crud Create")
	info, err := crud.Create(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorf("fail create AppUserExtra: %v", err)
		return &npool.CreateAppUserExtraResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateAppUserExtraResponse{
		Info: appUserExtraRowToObject(info),
	}, nil
}

func (s *AppUserExtraServer) CreateAppUserExtrasV2(ctx context.Context, in *npool.CreateAppUserExtrasRequest) (*npool.CreateAppUserExtrasResponse, error) {
	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateAppUserExtrasV2")
	defer span.End()
	var err error
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()
	if len(in.GetInfos()) == 0 {
		return &npool.CreateAppUserExtrasResponse{},
			status.Error(codes.InvalidArgument,
				"Batah create resource must more than 1",
			)
	}
	dupIDNumber := make(map[string]struct{})

	for key, info := range in.GetInfos() {
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
		err := checkAppUserExtraInfo(info)
		if err != nil {
			return &npool.CreateAppUserExtrasResponse{}, err
		}
		if _, ok := dupIDNumber[info.GetIDNumber()]; ok {
			return &npool.CreateAppUserExtrasResponse{},
				status.Errorf(codes.AlreadyExists,
					"IDNumber: %v duplicate create",
					info.GetIDNumber(),
				)
		}

		dupIDNumber[info.GetIDNumber()] = struct{}{}
	}

	rows, err := crud.CreateBulk(ctx, in.GetInfos())
	if err != nil {
		logger.Sugar().Errorf("fail create AppUserExtras: %v", err)
		return &npool.CreateAppUserExtrasResponse{}, status.Error(codes.Internal, err.Error())
	}

	infos := make([]*npool.AppUserExtra, 0, len(rows))
	for _, val := range rows {
		infos = append(infos, appUserExtraRowToObject(val))
	}

	return &npool.CreateAppUserExtrasResponse{
		Infos: infos,
	}, nil
}

func (s *AppUserExtraServer) UpdateAppUserExtraV2(ctx context.Context, in *npool.UpdateAppUserExtraRequest) (*npool.UpdateAppUserExtraResponse, error) {
	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "UpdateAppUserExtraV2")
	defer span.End()
	var err error
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()
	span = crud.AppUserExtraSpanAttributes(span, in.GetInfo())
	if _, err := uuid.Parse(in.GetInfo().GetID()); err != nil {
		logger.Sugar().Errorf("AppUserExtra id is invalid")
		return &npool.UpdateAppUserExtraResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	span.AddEvent("call crud Create")
	info, err := crud.Update(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorf("fail update AppUserExtra: %v", err)
		return &npool.UpdateAppUserExtraResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.UpdateAppUserExtraResponse{
		Info: appUserExtraRowToObject(info),
	}, nil
}

func (s *AppUserExtraServer) GetAppUserExtraV2(ctx context.Context, in *npool.GetAppUserExtraRequest) (*npool.GetAppUserExtraResponse, error) {
	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetAppUserExtraV2")
	defer span.End()
	var err error
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()
	span.SetAttributes(
		attribute.String("ID", in.GetID()),
	)
	id, err := uuid.Parse(in.GetID())
	if err != nil {
		return &npool.GetAppUserExtraResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	span.AddEvent("call crud Row")
	info, err := crud.Row(ctx, id)
	if err != nil {
		logger.Sugar().Errorf("fail get AppUserExtra: %v", err)
		return &npool.GetAppUserExtraResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetAppUserExtraResponse{
		Info: appUserExtraRowToObject(info),
	}, nil
}

func (s *AppUserExtraServer) GetAppUserExtraOnlyV2(ctx context.Context, in *npool.GetAppUserExtraOnlyRequest) (*npool.GetAppUserExtraOnlyResponse, error) {
	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetAppUserExtraOnlyV2")
	defer span.End()
	var err error
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()
	span = crud.AppUserExtraCondsSpanAttributes(span, in.GetConds())
	span.AddEvent("call crud RowOnly")
	info, err := crud.RowOnly(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorf("fail get AppUserExtras: %v", err)
		return &npool.GetAppUserExtraOnlyResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetAppUserExtraOnlyResponse{
		Info: appUserExtraRowToObject(info),
	}, nil
}

func (s *AppUserExtraServer) GetAppUserExtrasV2(ctx context.Context, in *npool.GetAppUserExtrasRequest) (*npool.GetAppUserExtrasResponse, error) {
	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetAppUserExtrasV2")
	defer span.End()
	var err error
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()
	span = crud.AppUserExtraCondsSpanAttributes(span, in.GetConds())
	span.SetAttributes(
		attribute.Int("Offset", int(in.GetOffset())),
		attribute.Int("Limit", int(in.GetLimit())),
	)
	span.AddEvent("call crud Rows")
	rows, total, err := crud.Rows(ctx, in.GetConds(), int(in.GetOffset()), int(in.GetLimit()))
	if err != nil {
		logger.Sugar().Errorf("fail get AppUserExtras: %v", err)
		return &npool.GetAppUserExtrasResponse{}, status.Error(codes.Internal, err.Error())
	}

	infos := make([]*npool.AppUserExtra, 0, len(rows))
	for _, val := range rows {
		infos = append(infos, appUserExtraRowToObject(val))
	}

	return &npool.GetAppUserExtrasResponse{
		Infos: infos,
		Total: uint32(total),
	}, nil
}

func (s *AppUserExtraServer) ExistAppUserExtraV2(ctx context.Context, in *npool.ExistAppUserExtraRequest) (*npool.ExistAppUserExtraResponse, error) {
	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "ExistAppUserExtraV2")
	defer span.End()
	var err error
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()
	span.SetAttributes(
		attribute.String("ID", in.GetID()),
	)
	id, err := uuid.Parse(in.GetID())
	if err != nil {
		return &npool.ExistAppUserExtraResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	span.AddEvent("call crud Exist")
	exist, err := crud.Exist(ctx, id)
	if err != nil {
		logger.Sugar().Errorf("fail check AppUserExtra: %v", err)
		return &npool.ExistAppUserExtraResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistAppUserExtraResponse{
		Info: exist,
	}, nil
}

func (s *AppUserExtraServer) ExistAppUserExtraCondsV2(ctx context.Context, in *npool.ExistAppUserExtraCondsRequest) (*npool.ExistAppUserExtraCondsResponse, error) {
	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "ExistAppUserExtraCondsV2")
	defer span.End()
	var err error
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()
	span = crud.AppUserExtraCondsSpanAttributes(span, in.GetConds())
	span.AddEvent("call crud ExistConds")
	exist, err := crud.ExistConds(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorf("fail check AppUserExtra: %v", err)
		return &npool.ExistAppUserExtraCondsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistAppUserExtraCondsResponse{
		Info: exist,
	}, nil
}

func (s *AppUserExtraServer) CountAppUserExtrasV2(ctx context.Context, in *npool.CountAppUserExtrasRequest) (*npool.CountAppUserExtrasResponse, error) {
	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CountAppUserExtrasV2")
	defer span.End()
	var err error
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()
	span = crud.AppUserExtraCondsSpanAttributes(span, in.GetConds())
	span.AddEvent("call crud Count")
	total, err := crud.Count(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorf("fail count AppUserExtra: %v", err)
		return &npool.CountAppUserExtrasResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CountAppUserExtrasResponse{
		Info: total,
	}, nil
}

func (s *AppUserExtraServer) DeleteAppUserExtraV2(ctx context.Context, in *npool.DeleteAppUserExtraRequest) (*npool.DeleteAppUserExtraResponse, error) {
	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "DeleteAppUserExtraV2")
	defer span.End()
	var err error
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()
	span.SetAttributes(
		attribute.String("ID", in.GetID()),
	)
	id, err := uuid.Parse(in.GetID())
	if err != nil {
		return &npool.DeleteAppUserExtraResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	span.AddEvent("call crud Delete")
	info, err := crud.Delete(ctx, id)
	if err != nil {
		logger.Sugar().Errorf("fail delete AppUserExtra: %v", err)
		return &npool.DeleteAppUserExtraResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.DeleteAppUserExtraResponse{
		Info: appUserExtraRowToObject(info),
	}, nil
}
