//go:build !codeanalysis
// +build !codeanalysis

package api

import (
	"context"
	"fmt"

	constant "github.com/NpoolPlatform/appuser-manager/pkg/message/const"
	scodes "go.opentelemetry.io/otel/codes"
	"google.golang.org/grpc/codes"

	crud "github.com/NpoolPlatform/appuser-manager/pkg/crud/approleuserv2"
	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/appusermgrv2/approleuser"
	"github.com/google/uuid"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
)

func checkAppRoleUserInfo(info *npool.AppRoleUserReq) error {
	if _, err := uuid.Parse(info.GetAppID()); err != nil {
		logger.Sugar().Error("AppID is invalid")
		return status.Error(codes.InvalidArgument, "AppID is invalid")
	}
	if _, err := uuid.Parse(info.GetRoleID()); err != nil {
		logger.Sugar().Error("RoleID is invalid")
		return status.Error(codes.InvalidArgument, "RoleID is invalid")
	}
	if _, err := uuid.Parse(info.GetUserID()); err != nil {
		logger.Sugar().Error("UserID is invalid")
		return status.Error(codes.InvalidArgument, "UserID is invalid")
	}
	return nil
}

func appRoleUserRowToObject(row *ent.AppRoleUser) *npool.AppRoleUser {
	return &npool.AppRoleUser{
		AppID:  row.AppID.String(),
		RoleID: row.RoleID.String(),
		UserID: row.UserID.String(),
		ID:     row.ID.String(),
	}
}

func (s *AppRoleUserServer) CreateAppRoleUserV2(ctx context.Context, in *npool.CreateAppRoleUserRequest) (*npool.CreateAppRoleUserResponse, error) {
	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateAppRoleUserV2")
	defer span.End()
	var err error
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()
	span = crud.AppRoleUserSpanAttributes(span, in.GetInfo())
	err = checkAppRoleUserInfo(in.GetInfo())
	if err != nil {
		return &npool.CreateAppRoleUserResponse{}, err
	}
	span.AddEvent("call crud Create")
	info, err := crud.Create(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorf("fail create app role user: %v", err)
		return &npool.CreateAppRoleUserResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateAppRoleUserResponse{
		Info: appRoleUserRowToObject(info),
	}, nil
}

func (s *AppRoleUserServer) CreateAppRoleUsersV2(ctx context.Context, in *npool.CreateAppRoleUsersRequest) (*npool.CreateAppRoleUsersResponse, error) {
	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateAppRoleUsersV2")
	defer span.End()
	var err error
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()
	if len(in.GetInfos()) == 0 {
		return &npool.CreateAppRoleUsersResponse{},
			status.Error(codes.InvalidArgument,
				"Batah create resource must more than 1",
			)
	}
	for key, info := range in.GetInfos() {
		span.SetAttributes(
			attribute.String("UserID"+fmt.Sprintf("%v", key), info.GetUserID()),
			attribute.String("ID"+fmt.Sprintf("%v", key), info.GetID()),
			attribute.String("AppID"+fmt.Sprintf("%v", key), info.GetAppID()),
			attribute.String("RoleID"+fmt.Sprintf("%v", key), info.GetRoleID()),
		)
		err = checkAppRoleUserInfo(info)
		if err != nil {
			return &npool.CreateAppRoleUsersResponse{}, err
		}
	}
	span.AddEvent("call crud CreateBulk")
	rows, err := crud.CreateBulk(ctx, in.GetInfos())
	if err != nil {
		logger.Sugar().Errorf("fail create app role user: %v", err)
		return &npool.CreateAppRoleUsersResponse{}, status.Error(codes.Internal, err.Error())
	}

	infos := make([]*npool.AppRoleUser, 0, len(rows))
	for _, val := range rows {
		infos = append(infos, appRoleUserRowToObject(val))
	}

	return &npool.CreateAppRoleUsersResponse{
		Infos: infos,
	}, nil
}

func (s *AppRoleUserServer) UpdateAppRoleUserV2(ctx context.Context, in *npool.UpdateAppRoleUserRequest) (*npool.UpdateAppRoleUserResponse, error) {
	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "UpdateAppRoleUserV2")
	defer span.End()
	var err error
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()
	span = crud.AppRoleUserSpanAttributes(span, in.GetInfo())
	if _, err := uuid.Parse(in.GetInfo().GetID()); err != nil {
		logger.Sugar().Errorf("AppRoleUser id is invalid")
		return &npool.UpdateAppRoleUserResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	span.AddEvent("call crud Update")
	info, err := crud.Update(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorf("fail update app role user: %v", err)
		return &npool.UpdateAppRoleUserResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.UpdateAppRoleUserResponse{
		Info: appRoleUserRowToObject(info),
	}, nil
}

func (s *AppRoleUserServer) GetAppRoleUserV2(ctx context.Context, in *npool.GetAppRoleUserRequest) (*npool.GetAppRoleUserResponse, error) {
	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetAppRoleUserV2")
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
		return &npool.GetAppRoleUserResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	span.AddEvent("call crud Row")
	info, err := crud.Row(ctx, id)
	if err != nil {
		logger.Sugar().Errorf("fail get app role user: %v", err)
		return &npool.GetAppRoleUserResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetAppRoleUserResponse{
		Info: appRoleUserRowToObject(info),
	}, nil
}

func (s *AppRoleUserServer) GetAppRoleUserOnlyV2(ctx context.Context, in *npool.GetAppRoleUserOnlyRequest) (*npool.GetAppRoleUserOnlyResponse, error) {
	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetAppRoleUserOnlyV2")
	defer span.End()
	var err error
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()
	span = crud.AppRoleUserCondsSpanAttributes(span, in.Conds)
	span.AddEvent("call crud RowOnly")
	info, err := crud.RowOnly(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorf("fail get app role users: %v", err)
		return &npool.GetAppRoleUserOnlyResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetAppRoleUserOnlyResponse{
		Info: appRoleUserRowToObject(info),
	}, nil
}

func (s *AppRoleUserServer) GetAppRoleUsersV2(ctx context.Context, in *npool.GetAppRoleUsersRequest) (*npool.GetAppRoleUsersResponse, error) {
	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetAppRoleUsersV2")
	defer span.End()
	var err error
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()
	span = crud.AppRoleUserCondsSpanAttributes(span, in.Conds)
	span.SetAttributes(
		attribute.Int("Offset", int(in.GetOffset())),
		attribute.Int("Limit", int(in.GetLimit())),
	)
	span.AddEvent("call crud Rows")
	rows, total, err := crud.Rows(ctx, in.GetConds(), int(in.GetOffset()), int(in.GetLimit()))
	if err != nil {
		logger.Sugar().Errorf("fail get app role users: %v", err)
		return &npool.GetAppRoleUsersResponse{}, status.Error(codes.Internal, err.Error())
	}

	infos := make([]*npool.AppRoleUser, 0, len(rows))
	for _, val := range rows {
		infos = append(infos, appRoleUserRowToObject(val))
	}

	return &npool.GetAppRoleUsersResponse{
		Infos: infos,
		Total: uint32(total),
	}, nil
}

func (s *AppRoleUserServer) ExistAppRoleUserV2(ctx context.Context, in *npool.ExistAppRoleUserRequest) (*npool.ExistAppRoleUserResponse, error) {
	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "ExistAppRoleUserV2")
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
		return &npool.ExistAppRoleUserResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	span.AddEvent("call crud Exist")
	exist, err := crud.Exist(ctx, id)
	if err != nil {
		logger.Sugar().Errorf("fail check app role user: %v", err)
		return &npool.ExistAppRoleUserResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistAppRoleUserResponse{
		Info: exist,
	}, nil
}

func (s *AppRoleUserServer) ExistAppRoleUserCondsV2(ctx context.Context, in *npool.ExistAppRoleUserCondsRequest) (*npool.ExistAppRoleUserCondsResponse, error) {
	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "ExistAppRoleUserCondsV2")
	defer span.End()
	var err error
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()
	span = crud.AppRoleUserCondsSpanAttributes(span, in.Conds)
	span.AddEvent("call crud ExistConds")
	exist, err := crud.ExistConds(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorf("fail check app role user: %v", err)
		return &npool.ExistAppRoleUserCondsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistAppRoleUserCondsResponse{
		Info: exist,
	}, nil
}

func (s *AppRoleUserServer) CountAppRoleUsersV2(ctx context.Context, in *npool.CountAppRoleUsersRequest) (*npool.CountAppRoleUsersResponse, error) {
	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CountAppRoleUsersV2")
	defer span.End()
	var err error
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()
	span = crud.AppRoleUserCondsSpanAttributes(span, in.Conds)
	span.AddEvent("call crud Count")
	total, err := crud.Count(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorf("fail count app role user: %v", err)
		return &npool.CountAppRoleUsersResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CountAppRoleUsersResponse{
		Info: total,
	}, nil
}

func (s *AppRoleUserServer) DeleteAppRoleUserV2(ctx context.Context, in *npool.DeleteAppRoleUserRequest) (*npool.DeleteAppRoleUserResponse, error) {
	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "DeleteAppRoleUserV2")
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
		return &npool.DeleteAppRoleUserResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	span.AddEvent("call crud Delete")
	info, err := crud.Delete(ctx, id)
	if err != nil {
		logger.Sugar().Errorf("fail delete app role user: %v", err)
		return &npool.DeleteAppRoleUserResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.DeleteAppRoleUserResponse{
		Info: appRoleUserRowToObject(info),
	}, nil
}
