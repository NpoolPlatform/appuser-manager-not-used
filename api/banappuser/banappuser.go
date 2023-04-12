//nolint:nolintlint,dupl
package banappuser

import (
	"context"

	converter "github.com/NpoolPlatform/appuser-manager/pkg/converter/banappuser"
	crud "github.com/NpoolPlatform/appuser-manager/pkg/crud/banappuser"
	servicename "github.com/NpoolPlatform/appuser-manager/pkg/servicename"
	tracer "github.com/NpoolPlatform/appuser-manager/pkg/tracer/banappuser"

	commontracer "github.com/NpoolPlatform/appuser-manager/pkg/tracer"

	"go.opentelemetry.io/otel"
	scodes "go.opentelemetry.io/otel/codes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/appuser/mgr/v2/banappuser"

	"github.com/google/uuid"
)

func (s *Server) CreateBanAppUser(ctx context.Context, in *npool.CreateBanAppUserRequest) (*npool.CreateBanAppUserResponse, error) {
	var err error

	_, span := otel.Tracer(servicename.ServiceDomain).Start(ctx, "CreateBanAppUser")
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
		logger.Sugar().Errorw("CreateBanAppUser", "error", err)
		return &npool.CreateBanAppUserResponse{}, err
	}

	span = commontracer.TraceInvoker(span, "banappuser", "crud", "Create")

	info, err := crud.Create(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorw("CreateBanAppUser", "error", err)
		return &npool.CreateBanAppUserResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateBanAppUserResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) CreateBanAppUsers(ctx context.Context, in *npool.CreateBanAppUsersRequest) (*npool.CreateBanAppUsersResponse, error) {
	var err error

	_, span := otel.Tracer(servicename.ServiceDomain).Start(ctx, "CreateBanAppUsers")
	defer span.End()
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	if len(in.GetInfos()) == 0 {
		logger.Sugar().Errorw("CreateBanAppUsers", "error", "Infos is empty")
		return &npool.CreateBanAppUsersResponse{}, status.Error(codes.InvalidArgument, "Infos is empty")
	}

	if err := validateMany(in.GetInfos()); err != nil {
		logger.Sugar().Errorw("CreateBanAppUsers", "error", err)
		return &npool.CreateBanAppUsersResponse{}, status.Error(codes.InvalidArgument, "Infos is invalid")
	}

	span = tracer.TraceMany(span, in.GetInfos())
	span = commontracer.TraceInvoker(span, "banappuser", "crud", "CreateBulk")

	rows, err := crud.CreateBulk(ctx, in.GetInfos())
	if err != nil {
		logger.Sugar().Errorw("CreateBanAppUsers", "error", err)
		return &npool.CreateBanAppUsersResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateBanAppUsersResponse{
		Infos: converter.Ent2GrpcMany(rows),
	}, nil
}

func (s *Server) UpdateBanAppUser(ctx context.Context, in *npool.UpdateBanAppUserRequest) (*npool.UpdateBanAppUserResponse, error) {
	var err error

	_, span := otel.Tracer(servicename.ServiceDomain).Start(ctx, "UpdateBanAppUser")
	defer span.End()
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.Trace(span, in.GetInfo())

	if _, err := uuid.Parse(in.GetInfo().GetID()); err != nil {
		logger.Sugar().Errorw("UpdateBanAppUser", "ID", in.GetInfo().GetID(), "error", err)
		return &npool.UpdateBanAppUserResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "banappuser", "crud", "Update")

	info, err := crud.Update(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorw("UpdateBanAppUser", "ID", in.GetInfo().GetID(), "error", err)
		return &npool.UpdateBanAppUserResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.UpdateBanAppUserResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetBanAppUser(ctx context.Context, in *npool.GetBanAppUserRequest) (*npool.GetBanAppUserResponse, error) {
	var err error

	_, span := otel.Tracer(servicename.ServiceDomain).Start(ctx, "GetBanAppUser")
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
		logger.Sugar().Errorw("GetBanAppUser", "ID", in.GetID(), "error", err)
		return &npool.GetBanAppUserResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "banappuser", "crud", "Row")

	info, err := crud.Row(ctx, id)
	if err != nil {
		logger.Sugar().Errorw("GetBanAppUser", "ID", in.GetID(), "error", err)
		return &npool.GetBanAppUserResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetBanAppUserResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetBanAppUserOnly(ctx context.Context, in *npool.GetBanAppUserOnlyRequest) (*npool.GetBanAppUserOnlyResponse, error) {
	var err error

	_, span := otel.Tracer(servicename.ServiceDomain).Start(ctx, "GetBanAppUserOnly")
	defer span.End()
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceInvoker(span, "banappuser", "crud", "RowOnly")

	info, err := crud.RowOnly(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorw("GetBanAppUserOnly", "error", err)
		return &npool.GetBanAppUserOnlyResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetBanAppUserOnlyResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetBanAppUsers(ctx context.Context, in *npool.GetBanAppUsersRequest) (*npool.GetBanAppUsersResponse, error) {
	var err error

	_, span := otel.Tracer(servicename.ServiceDomain).Start(ctx, "GetBanAppUsers")
	defer span.End()
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceOffsetLimit(span, int(in.GetOffset()), int(in.GetLimit()))
	span = commontracer.TraceInvoker(span, "banappuser", "crud", "Rows")

	rows, total, err := crud.Rows(ctx, in.GetConds(), int(in.GetOffset()), int(in.GetLimit()))
	if err != nil {
		logger.Sugar().Errorw("GetBanAppUsers", "error", err)
		return &npool.GetBanAppUsersResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetBanAppUsersResponse{
		Infos: converter.Ent2GrpcMany(rows),
		Total: uint32(total),
	}, nil
}

func (s *Server) ExistBanAppUser(ctx context.Context, in *npool.ExistBanAppUserRequest) (*npool.ExistBanAppUserResponse, error) {
	var err error

	_, span := otel.Tracer(servicename.ServiceDomain).Start(ctx, "ExistBanAppUser")
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
		logger.Sugar().Errorw("ExistBanAppUser", "ID", in.GetID(), "error", err)
		return &npool.ExistBanAppUserResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "banappuser", "crud", "Exist")

	exist, err := crud.Exist(ctx, id)
	if err != nil {
		logger.Sugar().Errorw("ExistBanAppUser", "ID", in.GetID(), "error", err)
		return &npool.ExistBanAppUserResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistBanAppUserResponse{
		Info: exist,
	}, nil
}

func (s *Server) ExistBanAppUserConds(ctx context.Context,
	in *npool.ExistBanAppUserCondsRequest) (*npool.ExistBanAppUserCondsResponse, error) {
	var err error

	_, span := otel.Tracer(servicename.ServiceDomain).Start(ctx, "ExistBanAppUserConds")
	defer span.End()
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceInvoker(span, "banappuser", "crud", "ExistBanAppUserConds")

	exist, err := crud.ExistConds(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorw("ExistBanAppUserConds", "error", err)
		return &npool.ExistBanAppUserCondsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistBanAppUserCondsResponse{
		Info: exist,
	}, nil
}

func (s *Server) CountBanAppUsers(ctx context.Context, in *npool.CountBanAppUsersRequest) (*npool.CountBanAppUsersResponse, error) {
	var err error

	_, span := otel.Tracer(servicename.ServiceDomain).Start(ctx, "CountBanAppUsers")
	defer span.End()
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceInvoker(span, "banappuser", "crud", "Count")

	total, err := crud.Count(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorw("Counts", "error", err)
		return &npool.CountBanAppUsersResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CountBanAppUsersResponse{
		Info: total,
	}, nil
}

func (s *Server) DeleteBanAppUser(ctx context.Context, in *npool.DeleteBanAppUserRequest) (*npool.DeleteBanAppUserResponse, error) {
	var err error

	_, span := otel.Tracer(servicename.ServiceDomain).Start(ctx, "DeleteBanAppUser")
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
		logger.Sugar().Errorw("DeleteBanAppUser", "ID", in.GetID(), "error", err)
		return &npool.DeleteBanAppUserResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "banappuser", "crud", "Delete")

	info, err := crud.Delete(ctx, id)
	if err != nil {
		logger.Sugar().Errorw("DeleteBanAppUser", "ID", in.GetID(), "error", err)
		return &npool.DeleteBanAppUserResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.DeleteBanAppUserResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}
