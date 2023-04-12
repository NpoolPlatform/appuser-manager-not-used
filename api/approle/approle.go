//nolint:nolintlint,dupl
package approle

import (
	"context"

	converter "github.com/NpoolPlatform/appuser-manager/pkg/converter/approle"
	crud "github.com/NpoolPlatform/appuser-manager/pkg/crud/approle"
	servicename "github.com/NpoolPlatform/appuser-manager/pkg/servicename"
	tracer "github.com/NpoolPlatform/appuser-manager/pkg/tracer/approle"

	commontracer "github.com/NpoolPlatform/appuser-manager/pkg/tracer"

	"go.opentelemetry.io/otel"
	scodes "go.opentelemetry.io/otel/codes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/appuser/mgr/v2/approle"

	"github.com/google/uuid"
)

func (s *Server) CreateAppRole(ctx context.Context, in *npool.CreateAppRoleRequest) (*npool.CreateAppRoleResponse, error) {
	var err error

	_, span := otel.Tracer(servicename.ServiceDomain).Start(ctx, "CreateAppRole")
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
		logger.Sugar().Errorw("CreateAppRole", "error", err)
		return &npool.CreateAppRoleResponse{}, err
	}

	span = commontracer.TraceInvoker(span, "approle", "crud", "Create")

	info, err := crud.Create(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorw("CreateAppRole", "error", err)
		return &npool.CreateAppRoleResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateAppRoleResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) CreateAppRoles(ctx context.Context, in *npool.CreateAppRolesRequest) (*npool.CreateAppRolesResponse, error) {
	var err error

	_, span := otel.Tracer(servicename.ServiceDomain).Start(ctx, "CreateAppRoles")
	defer span.End()
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	if len(in.GetInfos()) == 0 {
		logger.Sugar().Errorw("CreateAppRoles", "error", "Infos is empty")
		return &npool.CreateAppRolesResponse{}, status.Error(codes.InvalidArgument, "Infos is empty")
	}

	span = tracer.TraceMany(span, in.GetInfos())
	span = commontracer.TraceInvoker(span, "approle", "crud", "CreateBulk")

	rows, err := crud.CreateBulk(ctx, in.GetInfos())
	if err != nil {
		logger.Sugar().Errorw("CreateAppRoles", "error", err)
		return &npool.CreateAppRolesResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateAppRolesResponse{
		Infos: converter.Ent2GrpcMany(rows),
	}, nil
}

func (s *Server) UpdateAppRole(ctx context.Context, in *npool.UpdateAppRoleRequest) (*npool.UpdateAppRoleResponse, error) {
	var err error

	_, span := otel.Tracer(servicename.ServiceDomain).Start(ctx, "UpdateAppRole")
	defer span.End()
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.Trace(span, in.GetInfo())

	if _, err := uuid.Parse(in.GetInfo().GetID()); err != nil {
		logger.Sugar().Errorw("UpdateAppRole", "ID", in.GetInfo().GetID(), "error", err)
		return &npool.UpdateAppRoleResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "approle", "crud", "Update")

	info, err := crud.Update(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorw("UpdateAppRole", "ID", in.GetInfo().GetID(), "error", err)
		return &npool.UpdateAppRoleResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.UpdateAppRoleResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetAppRole(ctx context.Context, in *npool.GetAppRoleRequest) (*npool.GetAppRoleResponse, error) {
	var err error

	_, span := otel.Tracer(servicename.ServiceDomain).Start(ctx, "GetAppRole")
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
		logger.Sugar().Errorw("GetAppRole", "ID", in.GetID(), "error", err)
		return &npool.GetAppRoleResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "approle", "crud", "Row")

	info, err := crud.Row(ctx, id)
	if err != nil {
		logger.Sugar().Errorw("GetAppRole", "ID", in.GetID(), "error", err)
		return &npool.GetAppRoleResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetAppRoleResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetAppRoleOnly(ctx context.Context, in *npool.GetAppRoleOnlyRequest) (*npool.GetAppRoleOnlyResponse, error) {
	var err error

	_, span := otel.Tracer(servicename.ServiceDomain).Start(ctx, "GetAppRoleOnly")
	defer span.End()
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceInvoker(span, "approle", "crud", "RowOnly")

	info, err := crud.RowOnly(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorw("GetAppRoleOnly", "error", err)
		return &npool.GetAppRoleOnlyResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetAppRoleOnlyResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetAppRoles(ctx context.Context, in *npool.GetAppRolesRequest) (*npool.GetAppRolesResponse, error) {
	var err error

	_, span := otel.Tracer(servicename.ServiceDomain).Start(ctx, "GetAppRoles")
	defer span.End()
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceOffsetLimit(span, int(in.GetOffset()), int(in.GetLimit()))
	span = commontracer.TraceInvoker(span, "approle", "crud", "Rows")

	rows, total, err := crud.Rows(ctx, in.GetConds(), int(in.GetOffset()), int(in.GetLimit()))
	if err != nil {
		logger.Sugar().Errorw("GetAppRoles", "error", err)
		return &npool.GetAppRolesResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetAppRolesResponse{
		Infos: converter.Ent2GrpcMany(rows),
		Total: uint32(total),
	}, nil
}

func (s *Server) ExistAppRole(ctx context.Context, in *npool.ExistAppRoleRequest) (*npool.ExistAppRoleResponse, error) {
	var err error

	_, span := otel.Tracer(servicename.ServiceDomain).Start(ctx, "ExistAppRole")
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
		logger.Sugar().Errorw("ExistAppRole", "ID", in.GetID(), "error", err)
		return &npool.ExistAppRoleResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "approle", "crud", "Exist")

	exist, err := crud.Exist(ctx, id)
	if err != nil {
		logger.Sugar().Errorw("ExistAppRole", "ID", in.GetID(), "error", err)
		return &npool.ExistAppRoleResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistAppRoleResponse{
		Info: exist,
	}, nil
}

func (s *Server) ExistAppRoleConds(ctx context.Context, in *npool.ExistAppRoleCondsRequest) (*npool.ExistAppRoleCondsResponse, error) {
	var err error

	_, span := otel.Tracer(servicename.ServiceDomain).Start(ctx, "ExistAppRoleConds")
	defer span.End()
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceInvoker(span, "approle", "crud", "ExistAppRoleConds")

	exist, err := crud.ExistConds(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorw("ExistAppRoleConds", "error", err)
		return &npool.ExistAppRoleCondsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistAppRoleCondsResponse{
		Info: exist,
	}, nil
}

func (s *Server) CountAppRoles(ctx context.Context, in *npool.CountAppRolesRequest) (*npool.CountAppRolesResponse, error) {
	var err error

	_, span := otel.Tracer(servicename.ServiceDomain).Start(ctx, "CountAppRoles")
	defer span.End()
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceInvoker(span, "approle", "crud", "Count")

	total, err := crud.Count(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorw("Counts", "error", err)
		return &npool.CountAppRolesResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CountAppRolesResponse{
		Info: total,
	}, nil
}

func (s *Server) DeleteAppRole(ctx context.Context, in *npool.DeleteAppRoleRequest) (*npool.DeleteAppRoleResponse, error) {
	var err error

	_, span := otel.Tracer(servicename.ServiceDomain).Start(ctx, "DeleteAppRole")
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
		logger.Sugar().Errorw("DeleteAppRole", "ID", in.GetID(), "error", err)
		return &npool.DeleteAppRoleResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "approle", "crud", "Delete")

	info, err := crud.Delete(ctx, id)
	if err != nil {
		logger.Sugar().Errorw("DeleteAppRole", "ID", in.GetID(), "error", err)
		return &npool.DeleteAppRoleResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.DeleteAppRoleResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}
