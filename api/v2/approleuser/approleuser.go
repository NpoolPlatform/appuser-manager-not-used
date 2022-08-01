//nolint:nolintlint,dupl
package approleuser

import (
	"context"

	converter "github.com/NpoolPlatform/appuser-manager/pkg/converter/v2/approleuser"
	crud "github.com/NpoolPlatform/appuser-manager/pkg/crud/v2/approleuser"
	constant "github.com/NpoolPlatform/appuser-manager/pkg/message/const"
	tracer "github.com/NpoolPlatform/appuser-manager/pkg/tracer/approleuser"

	commontracer "github.com/NpoolPlatform/appuser-manager/pkg/tracer"

	"go.opentelemetry.io/otel"
	scodes "go.opentelemetry.io/otel/codes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/appuser/mgr/v2/approleuser"

	"github.com/google/uuid"
)

func (s *Server) CreateAppRoleUser(ctx context.Context, in *npool.CreateAppRoleUserRequest) (*npool.CreateAppRoleUserResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateAppRoleUser")
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
		logger.Sugar().Errorw("CreateAppRoleUser", "error", err)
		return &npool.CreateAppRoleUserResponse{}, err
	}

	span = commontracer.TraceInvoker(span, "approleuser", "crud", "Create")

	info, err := crud.Create(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorw("CreateAppRoleUser", "error", err)
		return &npool.CreateAppRoleUserResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateAppRoleUserResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) CreateAppRoleUsers(ctx context.Context, in *npool.CreateAppRoleUsersRequest) (*npool.CreateAppRoleUsersResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateAppRoleUsers")
	defer span.End()
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	if len(in.GetInfos()) == 0 {
		logger.Sugar().Errorw("CreateAppRoleUsers", "error", "Infos is empty")
		return &npool.CreateAppRoleUsersResponse{}, status.Error(codes.InvalidArgument, "Infos is empty")
	}

	span = tracer.TraceMany(span, in.GetInfos())
	span = commontracer.TraceInvoker(span, "approleuser", "crud", "CreateBulk")

	rows, err := crud.CreateBulk(ctx, in.GetInfos())
	if err != nil {
		logger.Sugar().Errorw("CreateAppRoleUsers", "error", err)
		return &npool.CreateAppRoleUsersResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateAppRoleUsersResponse{
		Infos: converter.Ent2GrpcMany(rows),
	}, nil
}

func (s *Server) UpdateAppRoleUser(ctx context.Context, in *npool.UpdateAppRoleUserRequest) (*npool.UpdateAppRoleUserResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "UpdateAppRoleUser")
	defer span.End()
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.Trace(span, in.GetInfo())

	if _, err := uuid.Parse(in.GetInfo().GetID()); err != nil {
		logger.Sugar().Errorw("UpdateAppRoleUser", "ID", in.GetInfo().GetID(), "error", err)
		return &npool.UpdateAppRoleUserResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "approleuser", "crud", "Update")

	info, err := crud.Update(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorw("UpdateAppRoleUser", "ID", in.GetInfo().GetID(), "error", err)
		return &npool.UpdateAppRoleUserResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.UpdateAppRoleUserResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetAppRoleUser(ctx context.Context, in *npool.GetAppRoleUserRequest) (*npool.GetAppRoleUserResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetAppRoleUser")
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
		logger.Sugar().Errorw("GetAppRoleUser", "ID", in.GetID(), "error", err)
		return &npool.GetAppRoleUserResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "approleuser", "crud", "Row")

	info, err := crud.Row(ctx, id)
	if err != nil {
		logger.Sugar().Errorw("GetAppRoleUser", "ID", in.GetID(), "error", err)
		return &npool.GetAppRoleUserResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetAppRoleUserResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetAppRoleUserOnly(ctx context.Context, in *npool.GetAppRoleUserOnlyRequest) (*npool.GetAppRoleUserOnlyResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetAppRoleUserOnly")
	defer span.End()
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceInvoker(span, "approleuser", "crud", "RowOnly")

	info, err := crud.RowOnly(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorw("GetAppRoleUserOnly", "error", err)
		return &npool.GetAppRoleUserOnlyResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetAppRoleUserOnlyResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetAppRoleUsers(ctx context.Context, in *npool.GetAppRoleUsersRequest) (*npool.GetAppRoleUsersResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetAppRoleUsers")
	defer span.End()
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceOffsetLimit(span, int(in.GetOffset()), int(in.GetLimit()))
	span = commontracer.TraceInvoker(span, "approleuser", "crud", "Rows")

	rows, total, err := crud.Rows(ctx, in.GetConds(), int(in.GetOffset()), int(in.GetLimit()))
	if err != nil {
		logger.Sugar().Errorw("GetAppRoleUsers", "error", err)
		return &npool.GetAppRoleUsersResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetAppRoleUsersResponse{
		Infos: converter.Ent2GrpcMany(rows),
		Total: uint32(total),
	}, nil
}

func (s *Server) ExistAppRoleUser(ctx context.Context, in *npool.ExistAppRoleUserRequest) (*npool.ExistAppRoleUserResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "ExistAppRoleUser")
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
		logger.Sugar().Errorw("ExistAppRoleUser", "ID", in.GetID(), "error", err)
		return &npool.ExistAppRoleUserResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "approleuser", "crud", "Exist")

	exist, err := crud.Exist(ctx, id)
	if err != nil {
		logger.Sugar().Errorw("ExistAppRoleUser", "ID", in.GetID(), "error", err)
		return &npool.ExistAppRoleUserResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistAppRoleUserResponse{
		Info: exist,
	}, nil
}

func (s *Server) ExistAppRoleUserConds(ctx context.Context,
	in *npool.ExistAppRoleUserCondsRequest) (*npool.ExistAppRoleUserCondsResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "ExistAppRoleUserConds")
	defer span.End()
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceInvoker(span, "approleuser", "crud", "ExistAppRoleUserConds")

	exist, err := crud.ExistConds(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorw("ExistAppRoleUserConds", "error", err)
		return &npool.ExistAppRoleUserCondsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistAppRoleUserCondsResponse{
		Info: exist,
	}, nil
}

func (s *Server) CountAppRoleUsers(ctx context.Context, in *npool.CountAppRoleUsersRequest) (*npool.CountAppRoleUsersResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CountAppRoleUsers")
	defer span.End()
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceInvoker(span, "approleuser", "crud", "Count")

	total, err := crud.Count(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorw("Counts", "error", err)
		return &npool.CountAppRoleUsersResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CountAppRoleUsersResponse{
		Info: total,
	}, nil
}

func (s *Server) DeleteAppRoleUser(ctx context.Context, in *npool.DeleteAppRoleUserRequest) (*npool.DeleteAppRoleUserResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "DeleteAppRoleUser")
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
		logger.Sugar().Errorw("DeleteAppRoleUser", "ID", in.GetID(), "error", err)
		return &npool.DeleteAppRoleUserResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "approleuser", "crud", "Delete")

	info, err := crud.Delete(ctx, id)
	if err != nil {
		logger.Sugar().Errorw("DeleteAppRoleUser", "ID", in.GetID(), "error", err)
		return &npool.DeleteAppRoleUserResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.DeleteAppRoleUserResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}
