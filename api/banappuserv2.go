//go:build !codeanalysis
// +build !codeanalysis

package api

import (
	"context"
	"fmt"

	crud "github.com/NpoolPlatform/appuser-manager/pkg/crud/banappuserv2"
	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent"
	constant "github.com/NpoolPlatform/appuser-manager/pkg/message/const"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	scodes "go.opentelemetry.io/otel/codes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/appusermgrv2/banappuser"

	"github.com/google/uuid"
)

func checkBanAppUserInfo(info *npool.BanAppUserReq) error {
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

func banAppUserRowToObject(row *ent.BanAppUser) *npool.BanAppUser {
	return &npool.BanAppUser{
		UserID:  row.UserID.String(),
		Message: row.Message,
		ID:      row.ID.String(),
		AppID:   row.AppID.String(),
	}
}

func (s *BanAppUserServer) CreateBanAppUserV2(ctx context.Context, in *npool.CreateBanAppUserRequest) (*npool.CreateBanAppUserResponse, error) {
	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateBanAppUserV2")
	defer span.End()
	var err error
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()
	span = crud.BanAppUserSpanAttributes(span, in.GetInfo())
	err = checkBanAppUserInfo(in.GetInfo())
	if err != nil {
		return &npool.CreateBanAppUserResponse{}, err
	}
	span.AddEvent("call crud Create")
	info, err := crud.Create(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorf("fail create BanAppUser: %v", err)
		return &npool.CreateBanAppUserResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateBanAppUserResponse{
		Info: banAppUserRowToObject(info),
	}, nil
}

func (s *BanAppUserServer) CreateBanAppUsersV2(ctx context.Context, in *npool.CreateBanAppUsersRequest) (*npool.CreateBanAppUsersResponse, error) {
	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateBanAppUsersV2")
	defer span.End()
	var err error
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()
	if len(in.GetInfos()) == 0 {
		return &npool.CreateBanAppUsersResponse{},
			status.Error(codes.InvalidArgument,
				"Batah create resource must more than 1",
			)
	}
	dupUserID := make(map[string]struct{})

	for key, info := range in.GetInfos() {
		span.SetAttributes(
			attribute.String("UserID"+fmt.Sprintf("%v", key), info.GetUserID()),
			attribute.String("Message"+fmt.Sprintf("%v", key), info.GetMessage()),
			attribute.String("ID"+fmt.Sprintf("%v", key), info.GetID()),
			attribute.String("AppID"+fmt.Sprintf("%v", key), info.GetID()),
		)
		err := checkBanAppUserInfo(info)
		if err != nil {
			return &npool.CreateBanAppUsersResponse{}, err
		}
		if _, ok := dupUserID[info.GetUserID()]; ok {
			return &npool.CreateBanAppUsersResponse{},
				status.Errorf(codes.AlreadyExists,
					"UserID: %v duplicate create",
					info.GetUserID(),
				)
		}
		dupUserID[info.GetUserID()] = struct{}{}
	}
	span.AddEvent("call crud CreateBulk")
	rows, err := crud.CreateBulk(ctx, in.GetInfos())
	if err != nil {
		logger.Sugar().Errorf("fail create BanAppUsers: %v", err)
		return &npool.CreateBanAppUsersResponse{}, status.Error(codes.Internal, err.Error())
	}

	infos := make([]*npool.BanAppUser, 0, len(rows))
	for _, val := range rows {
		infos = append(infos, banAppUserRowToObject(val))
	}

	return &npool.CreateBanAppUsersResponse{
		Infos: infos,
	}, nil
}

func (s *BanAppUserServer) UpdateBanAppUserV2(ctx context.Context, in *npool.UpdateBanAppUserRequest) (*npool.UpdateBanAppUserResponse, error) {
	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "UpdateBanAppUserV2")
	defer span.End()
	var err error
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()
	span = crud.BanAppUserSpanAttributes(span, in.GetInfo())
	if _, err := uuid.Parse(in.GetInfo().GetID()); err != nil {
		logger.Sugar().Errorf("BanAppUser id is invalid")
		return &npool.UpdateBanAppUserResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	span.AddEvent("call crud Update")
	info, err := crud.Update(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorf("fail update BanAppUser: %v", err)
		return &npool.UpdateBanAppUserResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.UpdateBanAppUserResponse{
		Info: banAppUserRowToObject(info),
	}, nil
}

func (s *BanAppUserServer) GetBanAppUserV2(ctx context.Context, in *npool.GetBanAppUserRequest) (*npool.GetBanAppUserResponse, error) {
	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetBanAppUserV2")
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
		return &npool.GetBanAppUserResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	span.AddEvent("call crud Row")
	info, err := crud.Row(ctx, id)
	if err != nil {
		logger.Sugar().Errorf("fail get BanAppUser: %v", err)
		return &npool.GetBanAppUserResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetBanAppUserResponse{
		Info: banAppUserRowToObject(info),
	}, nil
}

func (s *BanAppUserServer) GetBanAppUserOnlyV2(ctx context.Context, in *npool.GetBanAppUserOnlyRequest) (*npool.GetBanAppUserOnlyResponse, error) {
	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetBanAppUserOnlyV2")
	defer span.End()
	var err error
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()
	span = crud.BanAppUserCondsSpanAttributes(span, in.GetConds())
	span.AddEvent("call crud RowOnly")
	info, err := crud.RowOnly(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorf("fail get BanAppUsers: %v", err)
		return &npool.GetBanAppUserOnlyResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetBanAppUserOnlyResponse{
		Info: banAppUserRowToObject(info),
	}, nil
}

func (s *BanAppUserServer) GetBanAppUsersV2(ctx context.Context, in *npool.GetBanAppUsersRequest) (*npool.GetBanAppUsersResponse, error) {
	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetBanAppUsersV2")
	defer span.End()
	var err error
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()
	span = crud.BanAppUserCondsSpanAttributes(span, in.GetConds())
	span.SetAttributes(
		attribute.Int("Offset", int(in.GetOffset())),
		attribute.Int("Limit", int(in.GetLimit())),
	)
	span.AddEvent("call crud Rows")
	rows, total, err := crud.Rows(ctx, in.GetConds(), int(in.GetOffset()), int(in.GetLimit()))
	if err != nil {
		logger.Sugar().Errorf("fail get BanAppUsers: %v", err)
		return &npool.GetBanAppUsersResponse{}, status.Error(codes.Internal, err.Error())
	}

	infos := make([]*npool.BanAppUser, 0, len(rows))
	for _, val := range rows {
		infos = append(infos, banAppUserRowToObject(val))
	}

	return &npool.GetBanAppUsersResponse{
		Infos: infos,
		Total: uint32(total),
	}, nil
}

func (s *BanAppUserServer) ExistBanAppUserV2(ctx context.Context, in *npool.ExistBanAppUserRequest) (*npool.ExistBanAppUserResponse, error) {
	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "ExistBanAppUserV2")
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
		return &npool.ExistBanAppUserResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	span.AddEvent("call crud Exist")
	exist, err := crud.Exist(ctx, id)
	if err != nil {
		logger.Sugar().Errorf("fail check BanAppUser: %v", err)
		return &npool.ExistBanAppUserResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistBanAppUserResponse{
		Info: exist,
	}, nil
}

func (s *BanAppUserServer) ExistBanAppUserCondsV2(ctx context.Context, in *npool.ExistBanAppUserCondsRequest) (*npool.ExistBanAppUserCondsResponse, error) {
	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "ExistBanAppUserCondsV2")
	defer span.End()
	var err error
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()
	span = crud.BanAppUserCondsSpanAttributes(span, in.GetConds())
	span.AddEvent("call crud ExistConds")
	exist, err := crud.ExistConds(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorf("fail check BanAppUser: %v", err)
		return &npool.ExistBanAppUserCondsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistBanAppUserCondsResponse{
		Info: exist,
	}, nil
}

func (s *BanAppUserServer) CountBanAppUsersV2(ctx context.Context, in *npool.CountBanAppUsersRequest) (*npool.CountBanAppUsersResponse, error) {
	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CountBanAppUsersV2")
	defer span.End()
	var err error
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()
	span = crud.BanAppUserCondsSpanAttributes(span, in.GetConds())
	span.AddEvent("call crud Count")
	total, err := crud.Count(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorf("fail count BanAppUser: %v", err)
		return &npool.CountBanAppUsersResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CountBanAppUsersResponse{
		Info: total,
	}, nil
}

func (s *BanAppUserServer) DeleteBanAppUserV2(ctx context.Context, in *npool.DeleteBanAppUserRequest) (*npool.DeleteBanAppUserResponse, error) {
	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "DeleteBanAppUserV2")
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
		return &npool.DeleteBanAppUserResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	span.AddEvent("call crud Delete")
	info, err := crud.Delete(ctx, id)
	if err != nil {
		logger.Sugar().Errorf("fail delete BanAppUser: %v", err)
		return &npool.DeleteBanAppUserResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.DeleteBanAppUserResponse{
		Info: banAppUserRowToObject(info),
	}, nil
}
