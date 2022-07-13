//go:build !codeanalysis
// +build !codeanalysis

package api

import (
	"context"
	"fmt"

	crud "github.com/NpoolPlatform/appuser-manager/pkg/crud/appuserv2"
	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent"
	constant "github.com/NpoolPlatform/appuser-manager/pkg/message/const"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	scodes "go.opentelemetry.io/otel/codes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/appusermgrv2/appuser"

	"github.com/google/uuid"
)

func checkAppUserInfo(info *npool.AppUserReq) error {
	if _, err := uuid.Parse(info.GetAppID()); err != nil {
		logger.Sugar().Error("AppID is invalid")
		return status.Error(codes.InvalidArgument, "AppID is invalid")
	}
	if _, err := uuid.Parse(info.GetImportFromApp()); err != nil {
		logger.Sugar().Error("ImportFromApp is invalid")
		return status.Error(codes.InvalidArgument, "ImportFromApp is invalid")
	}
	return nil
}

func appUserRowToObject(row *ent.AppUser) *npool.AppUser {
	return &npool.AppUser{
		ID:            row.ID.String(),
		AppID:         row.AppID.String(),
		EmailAddress:  row.EmailAddress,
		PhoneNo:       row.PhoneNo,
		ImportFromApp: row.ImportFromApp.String(),
	}
}

func (s *AppUserServer) CreateAppUserV2(ctx context.Context, in *npool.CreateAppUserRequest) (*npool.CreateAppUserResponse, error) {
	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateAppUserV2")
	defer span.End()
	var err error
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()
	span = crud.AppUserSpanAttributes(span, in.GetInfo())
	err = checkAppUserInfo(in.GetInfo())
	if err != nil {
		return &npool.CreateAppUserResponse{}, err
	}
	span.AddEvent("call crud Create")
	info, err := crud.Create(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorf("fail create AppUser: %v", err)
		return &npool.CreateAppUserResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateAppUserResponse{
		Info: appUserRowToObject(info),
	}, nil
}

func (s *AppUserServer) CreateAppUsersV2(ctx context.Context, in *npool.CreateAppUsersRequest) (*npool.CreateAppUsersResponse, error) {
	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateAppUsersV2")
	defer span.End()
	var err error
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()
	if len(in.GetInfos()) == 0 {
		return &npool.CreateAppUsersResponse{},
			status.Error(codes.InvalidArgument,
				"Batah create resource must more than 1",
			)
	}
	dupEmailAddress := make(map[string]struct{})

	dupPhoneNo := make(map[string]struct{})

	for key, info := range in.GetInfos() {
		span.SetAttributes(
			attribute.String("PhoneNo"+fmt.Sprintf("%v", key), info.GetPhoneNo()),
			attribute.String("ImportFromApp"+fmt.Sprintf("%v", key), info.GetImportFromApp()),
			attribute.String("ID"+fmt.Sprintf("%v", key), info.GetID()),
			attribute.String("AppID"+fmt.Sprintf("%v", key), info.GetAppID()),
			attribute.String("EmailAddress"+fmt.Sprintf("%v", key), info.GetEmailAddress()),
		)
		err := checkAppUserInfo(info)
		if err != nil {
			return &npool.CreateAppUsersResponse{}, err
		}
		if _, ok := dupEmailAddress[info.GetEmailAddress()]; ok {
			return &npool.CreateAppUsersResponse{},
				status.Errorf(codes.AlreadyExists,
					"EmailAddress: %v duplicate create",
					info.GetEmailAddress(),
				)
		}
		dupEmailAddress[info.GetEmailAddress()] = struct{}{}
		if _, ok := dupPhoneNo[info.GetPhoneNo()]; ok {
			return &npool.CreateAppUsersResponse{},
				status.Errorf(codes.AlreadyExists,
					"PhoneNo: %v duplicate create",
					info.GetPhoneNo(),
				)
		}
		dupPhoneNo[info.GetPhoneNo()] = struct{}{}
	}
	span.AddEvent("call crud CreateBulk")
	rows, err := crud.CreateBulk(ctx, in.GetInfos())
	if err != nil {
		logger.Sugar().Errorf("fail create AppUsers: %v", err)
		return &npool.CreateAppUsersResponse{}, status.Error(codes.Internal, err.Error())
	}

	infos := make([]*npool.AppUser, 0, len(rows))
	for _, val := range rows {
		infos = append(infos, appUserRowToObject(val))
	}

	return &npool.CreateAppUsersResponse{
		Infos: infos,
	}, nil
}

func (s *AppUserServer) UpdateAppUserV2(ctx context.Context, in *npool.UpdateAppUserRequest) (*npool.UpdateAppUserResponse, error) {
	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "UpdateAppUserV2")
	defer span.End()
	var err error
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()
	span = crud.AppUserSpanAttributes(span, in.GetInfo())
	if _, err := uuid.Parse(in.GetInfo().GetID()); err != nil {
		logger.Sugar().Errorf("AppUser id is invalid")
		return &npool.UpdateAppUserResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	if _, err := uuid.Parse(in.GetInfo().GetAppID()); err != nil {
		logger.Sugar().Errorf("AppUser is invalid")
		return &npool.UpdateAppUserResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	if _, err := uuid.Parse(in.GetInfo().GetImportFromApp()); err != nil {
		logger.Sugar().Errorf("AppUser is invalid")
		return &npool.UpdateAppUserResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	span.AddEvent("call crud Update")
	info, err := crud.Update(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorf("fail update AppUser: %v", err)
		return &npool.UpdateAppUserResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.UpdateAppUserResponse{
		Info: appUserRowToObject(info),
	}, nil
}

func (s *AppUserServer) GetAppUserV2(ctx context.Context, in *npool.GetAppUserRequest) (*npool.GetAppUserResponse, error) {
	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetAppUserV2")
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
		return &npool.GetAppUserResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	span.AddEvent("call crud Row")
	info, err := crud.Row(ctx, id)
	if err != nil {
		logger.Sugar().Errorf("fail get AppUser: %v", err)
		return &npool.GetAppUserResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetAppUserResponse{
		Info: appUserRowToObject(info),
	}, nil
}

func (s *AppUserServer) GetAppUserOnlyV2(ctx context.Context, in *npool.GetAppUserOnlyRequest) (*npool.GetAppUserOnlyResponse, error) {
	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetAppUserOnlyV2")
	defer span.End()
	var err error
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()
	span = crud.AppUserCondsSpanAttributes(span, in.GetConds())
	span.AddEvent("call crud RowOnly")
	info, err := crud.RowOnly(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorf("fail get AppUsers: %v", err)
		return &npool.GetAppUserOnlyResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetAppUserOnlyResponse{
		Info: appUserRowToObject(info),
	}, nil
}

func (s *AppUserServer) GetAppUsersV2(ctx context.Context, in *npool.GetAppUsersRequest) (*npool.GetAppUsersResponse, error) {
	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetAppUsersV2")
	defer span.End()
	var err error
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()
	span = crud.AppUserCondsSpanAttributes(span, in.GetConds())
	span.SetAttributes(
		attribute.Int("Offset", int(in.GetOffset())),
		attribute.Int("Limit", int(in.GetLimit())),
	)
	span.AddEvent("call crud Rows")
	rows, total, err := crud.Rows(ctx, in.GetConds(), int(in.GetOffset()), int(in.GetLimit()))
	if err != nil {
		logger.Sugar().Errorf("fail get AppUsers: %v", err)
		return &npool.GetAppUsersResponse{}, status.Error(codes.Internal, err.Error())
	}

	infos := make([]*npool.AppUser, 0, len(rows))
	for _, val := range rows {
		infos = append(infos, appUserRowToObject(val))
	}

	return &npool.GetAppUsersResponse{
		Infos: infos,
		Total: uint32(total),
	}, nil
}

func (s *AppUserServer) ExistAppUserV2(ctx context.Context, in *npool.ExistAppUserRequest) (*npool.ExistAppUserResponse, error) {
	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "ExistAppUserV2")
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
		return &npool.ExistAppUserResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	span.AddEvent("call crud Exist")
	exist, err := crud.Exist(ctx, id)
	if err != nil {
		logger.Sugar().Errorf("fail check AppUser: %v", err)
		return &npool.ExistAppUserResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistAppUserResponse{
		Info: exist,
	}, nil
}

func (s *AppUserServer) ExistAppUserCondsV2(ctx context.Context, in *npool.ExistAppUserCondsRequest) (*npool.ExistAppUserCondsResponse, error) {
	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "ExistAppUserCondsV2")
	defer span.End()
	var err error
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()
	span = crud.AppUserCondsSpanAttributes(span, in.GetConds())
	span.AddEvent("call crud ExistConds")
	exist, err := crud.ExistConds(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorf("fail check AppUser: %v", err)
		return &npool.ExistAppUserCondsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistAppUserCondsResponse{
		Info: exist,
	}, nil
}

func (s *AppUserServer) CountAppUsersV2(ctx context.Context, in *npool.CountAppUsersRequest) (*npool.CountAppUsersResponse, error) {
	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CountAppUsersV2")
	defer span.End()
	var err error
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()
	span = crud.AppUserCondsSpanAttributes(span, in.GetConds())
	span.AddEvent("call crud Count")
	total, err := crud.Count(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorf("fail count AppUser: %v", err)
		return &npool.CountAppUsersResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CountAppUsersResponse{
		Info: total,
	}, nil
}

func (s *AppUserServer) DeleteAppUserV2(ctx context.Context, in *npool.DeleteAppUserRequest) (*npool.DeleteAppUserResponse, error) {
	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "DeleteAppUserV2")
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
		return &npool.DeleteAppUserResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	span.AddEvent("call crud Delete")
	info, err := crud.Delete(ctx, id)
	if err != nil {
		logger.Sugar().Errorf("fail delete AppUser: %v", err)
		return &npool.DeleteAppUserResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.DeleteAppUserResponse{
		Info: appUserRowToObject(info),
	}, nil
}
