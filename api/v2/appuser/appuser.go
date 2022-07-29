//nolint:dupl
package appuser

import (
	"context"

	converter "github.com/NpoolPlatform/appuser-manager/pkg/converter/appuser"
	crud "github.com/NpoolPlatform/appuser-manager/pkg/crud/v2/appuser"
	constant "github.com/NpoolPlatform/appuser-manager/pkg/message/const"
	tracer "github.com/NpoolPlatform/appuser-manager/pkg/tracer/appuser"

	commontracer "github.com/NpoolPlatform/appuser-manager/pkg/tracer"

	"go.opentelemetry.io/otel"
	scodes "go.opentelemetry.io/otel/codes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/appuser/mgr/v2/appuser"

	"github.com/google/uuid"
)

func (s *Server) CreateAppUser(ctx context.Context, in *npool.CreateAppUserRequest) (*npool.CreateAppUserResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateAppUser")
	defer span.End()
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.Trace(span, in.GetInfo())

	err = validate(in.GetInfo())
	if err != nil {
		logger.Sugar().Errorw("CreateAppUser", "error", err)
		return &npool.CreateAppUserResponse{}, err
	}

	span = commontracer.TraceInvoker(span, "appuser", "crud", "Create")

	info, err := crud.Create(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorw("CreateAppUser", "error", err)
		return &npool.CreateAppUserResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateAppUserResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) CreateAppUsers(ctx context.Context, in *npool.CreateAppUsersRequest) (*npool.CreateAppUsersResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateAppUsers")
	defer span.End()
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	if len(in.GetInfos()) == 0 {
		logger.Sugar().Errorw("CreateAppUsers", "error", "Infos is empty")
		return &npool.CreateAppUsersResponse{}, status.Error(codes.InvalidArgument, "Infos is empty")
	}

	span = tracer.TraceMany(span, in.GetInfos())
	span = commontracer.TraceInvoker(span, "appuser", "crud", "CreateBulk")

	rows, err := crud.CreateBulk(ctx, in.GetInfos())
	if err != nil {
		logger.Sugar().Errorw("CreateAppUsers", "error", err)
		return &npool.CreateAppUsersResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateAppUsersResponse{
		Infos: converter.Ent2GrpcMany(rows),
	}, nil
}

func (s *Server) UpdateAppUser(ctx context.Context, in *npool.UpdateAppUserRequest) (*npool.UpdateAppUserResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "UpdateAppUser")
	defer span.End()
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.Trace(span, in.GetInfo())

	if _, err := uuid.Parse(in.GetInfo().GetID()); err != nil {
		logger.Sugar().Errorw("UpdateAppUser", "ID", in.GetInfo().GetID(), "error", err)
		return &npool.UpdateAppUserResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "appuser", "crud", "Update")

	info, err := crud.Update(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorw("UpdateAppUser", "ID", in.GetInfo().GetID(), "error", err)
		return &npool.UpdateAppUserResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.UpdateAppUserResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetAppUser(ctx context.Context, in *npool.GetAppUserRequest) (*npool.GetAppUserResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetAppUser")
	defer span.End()
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = commontracer.TraceID(span, in.GetID())

	id, err := uuid.Parse(in.GetID())
	if err != nil {
		logger.Sugar().Errorw("GetAppUser", "ID", in.GetID(), "error", err)
		return &npool.GetAppUserResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "appuser", "crud", "Row")

	info, err := crud.Row(ctx, id)
	if err != nil {
		logger.Sugar().Errorw("GetAppUser", "ID", in.GetID(), "error", err)
		return &npool.GetAppUserResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetAppUserResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetAppUserOnly(ctx context.Context, in *npool.GetAppUserOnlyRequest) (*npool.GetAppUserOnlyResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetAppUserOnly")
	defer span.End()
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceInvoker(span, "appuser", "crud", "RowOnly")

	info, err := crud.RowOnly(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorw("GetAppUserOnly", "error", err)
		return &npool.GetAppUserOnlyResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetAppUserOnlyResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetAppUsers(ctx context.Context, in *npool.GetAppUsersRequest) (*npool.GetAppUsersResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetAppUsers")
	defer span.End()
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceOffsetLimit(span, int(in.GetOffset()), int(in.GetLimit()))
	span = commontracer.TraceInvoker(span, "appuser", "crud", "Rows")

	rows, total, err := crud.Rows(ctx, in.GetConds(), int(in.GetOffset()), int(in.GetLimit()))
	if err != nil {
		logger.Sugar().Errorw("GetAppUsers", "error", err)
		return &npool.GetAppUsersResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetAppUsersResponse{
		Infos: converter.Ent2GrpcMany(rows),
		Total: uint32(total),
	}, nil
}

func (s *Server) ExistAppUser(ctx context.Context, in *npool.ExistAppUserRequest) (*npool.ExistAppUserResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "ExistAppUser")
	defer span.End()
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = commontracer.TraceID(span, in.GetID())

	id, err := uuid.Parse(in.GetID())
	if err != nil {
		logger.Sugar().Errorw("ExistAppUser", "ID", in.GetID(), "error", err)
		return &npool.ExistAppUserResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "appuser", "crud", "Exist")

	exist, err := crud.Exist(ctx, id)
	if err != nil {
		logger.Sugar().Errorw("ExistAppUser", "ID", in.GetID(), "error", err)
		return &npool.ExistAppUserResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistAppUserResponse{
		Info: exist,
	}, nil
}

func (s *Server) ExistAppUserConds(ctx context.Context, in *npool.ExistAppUserCondsRequest) (*npool.ExistAppUserCondsResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "ExistAppUserConds")
	defer span.End()
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceInvoker(span, "appuser", "crud", "ExistAppUserConds")

	exist, err := crud.ExistConds(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorw("ExistAppUserConds", "error", err)
		return &npool.ExistAppUserCondsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistAppUserCondsResponse{
		Info: exist,
	}, nil
}

func (s *Server) CountAppUsers(ctx context.Context, in *npool.CountAppUsersRequest) (*npool.CountAppUsersResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CountAppUsers")
	defer span.End()
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceInvoker(span, "appuser", "crud", "Count")

	total, err := crud.Count(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorw("Counts", "error", err)
		return &npool.CountAppUsersResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CountAppUsersResponse{
		Info: total,
	}, nil
}

func (s *Server) DeleteAppUser(ctx context.Context, in *npool.DeleteAppUserRequest) (*npool.DeleteAppUserResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "DeleteAppUser")
	defer span.End()
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = commontracer.TraceID(span, in.GetID())

	id, err := uuid.Parse(in.GetID())
	if err != nil {
		logger.Sugar().Errorw("DeleteAppUser", "ID", in.GetID(), "error", err)
		return &npool.DeleteAppUserResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "appuser", "crud", "Delete")

	info, err := crud.Delete(ctx, id)
	if err != nil {
		logger.Sugar().Errorw("DeleteAppUser", "ID", in.GetID(), "error", err)
		return &npool.DeleteAppUserResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.DeleteAppUserResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}
