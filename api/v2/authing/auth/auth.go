//nolint:nolintlint,dupl
package auth

import (
	"context"

	converter "github.com/NpoolPlatform/appuser-manager/pkg/converter/v2/authing/auth"
	crud "github.com/NpoolPlatform/appuser-manager/pkg/crud/v2/authing/auth"
	constant "github.com/NpoolPlatform/appuser-manager/pkg/message/const"
	tracer "github.com/NpoolPlatform/appuser-manager/pkg/tracer/authing/auth"

	commontracer "github.com/NpoolPlatform/appuser-manager/pkg/tracer"

	"go.opentelemetry.io/otel"
	scodes "go.opentelemetry.io/otel/codes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/appuser/mgr/v2/authing/auth"

	"github.com/google/uuid"
)

func (s *Server) CreateAuth(ctx context.Context, in *npool.CreateAuthRequest) (*npool.CreateAuthResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateAuth")
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
		logger.Sugar().Errorw("CreateAuth", "error", err)
		return &npool.CreateAuthResponse{}, err
	}

	span = commontracer.TraceInvoker(span, "auth", "crud", "Create")

	info, err := crud.Create(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorw("CreateAuth", "error", err)
		return &npool.CreateAuthResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateAuthResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) CreateAuths(ctx context.Context, in *npool.CreateAuthsRequest) (*npool.CreateAuthsResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateAuths")
	defer span.End()
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	if len(in.GetInfos()) == 0 {
		logger.Sugar().Errorw("CreateAuths", "error", "Infos is empty")
		return &npool.CreateAuthsResponse{}, status.Error(codes.InvalidArgument, "Infos is empty")
	}

	if err := validateMany(in.GetInfos()); err != nil {
		logger.Sugar().Errorw("CreateAuths", "error", err)
		return &npool.CreateAuthsResponse{}, status.Error(codes.InvalidArgument, "Infos is invalid")
	}

	span = tracer.TraceMany(span, in.GetInfos())
	span = commontracer.TraceInvoker(span, "auth", "crud", "CreateBulk")

	rows, err := crud.CreateBulk(ctx, in.GetInfos())
	if err != nil {
		logger.Sugar().Errorw("CreateAuths", "error", err)
		return &npool.CreateAuthsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateAuthsResponse{
		Infos: converter.Ent2GrpcMany(rows),
	}, nil
}

func (s *Server) UpdateAuth(ctx context.Context, in *npool.UpdateAuthRequest) (*npool.UpdateAuthResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "UpdateAuth")
	defer span.End()
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.Trace(span, in.GetInfo())

	if _, err := uuid.Parse(in.GetInfo().GetID()); err != nil {
		logger.Sugar().Errorw("UpdateAuth", "ID", in.GetInfo().GetID(), "error", err)
		return &npool.UpdateAuthResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "auth", "crud", "Update")

	info, err := crud.Update(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorw("UpdateAuth", "ID", in.GetInfo().GetID(), "error", err)
		return &npool.UpdateAuthResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.UpdateAuthResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetAuth(ctx context.Context, in *npool.GetAuthRequest) (*npool.GetAuthResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetAuth")
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
		logger.Sugar().Errorw("GetAuth", "ID", in.GetID(), "error", err)
		return &npool.GetAuthResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "auth", "crud", "Row")

	info, err := crud.Row(ctx, id)
	if err != nil {
		logger.Sugar().Errorw("GetAuth", "ID", in.GetID(), "error", err)
		return &npool.GetAuthResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetAuthResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetAuthOnly(ctx context.Context, in *npool.GetAuthOnlyRequest) (*npool.GetAuthOnlyResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetAuthOnly")
	defer span.End()
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceInvoker(span, "auth", "crud", "RowOnly")

	info, err := crud.RowOnly(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorw("GetAuthOnly", "error", err)
		return &npool.GetAuthOnlyResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetAuthOnlyResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetAuths(ctx context.Context, in *npool.GetAuthsRequest) (*npool.GetAuthsResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetAuths")
	defer span.End()
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceOffsetLimit(span, int(in.GetOffset()), int(in.GetLimit()))
	span = commontracer.TraceInvoker(span, "auth", "crud", "Rows")

	rows, total, err := crud.Rows(ctx, in.GetConds(), int(in.GetOffset()), int(in.GetLimit()))
	if err != nil {
		logger.Sugar().Errorw("GetAuths", "error", err)
		return &npool.GetAuthsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetAuthsResponse{
		Infos: converter.Ent2GrpcMany(rows),
		Total: uint32(total),
	}, nil
}

func (s *Server) ExistAuth(ctx context.Context, in *npool.ExistAuthRequest) (*npool.ExistAuthResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "ExistAuth")
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
		logger.Sugar().Errorw("ExistAuth", "ID", in.GetID(), "error", err)
		return &npool.ExistAuthResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "auth", "crud", "Exist")

	exist, err := crud.Exist(ctx, id)
	if err != nil {
		logger.Sugar().Errorw("ExistAuth", "ID", in.GetID(), "error", err)
		return &npool.ExistAuthResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistAuthResponse{
		Info: exist,
	}, nil
}

func (s *Server) ExistAuthConds(ctx context.Context, in *npool.ExistAuthCondsRequest) (*npool.ExistAuthCondsResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "ExistAuthConds")
	defer span.End()
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceInvoker(span, "auth", "crud", "ExistAuthConds")

	exist, err := crud.ExistConds(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorw("ExistAuthConds", "error", err)
		return &npool.ExistAuthCondsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistAuthCondsResponse{
		Info: exist,
	}, nil
}

func (s *Server) CountAuths(ctx context.Context, in *npool.CountAuthsRequest) (*npool.CountAuthsResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CountAuths")
	defer span.End()
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceInvoker(span, "auth", "crud", "Count")

	total, err := crud.Count(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorw("Counts", "error", err)
		return &npool.CountAuthsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CountAuthsResponse{
		Info: total,
	}, nil
}

func (s *Server) DeleteAuth(ctx context.Context, in *npool.DeleteAuthRequest) (*npool.DeleteAuthResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "DeleteAuth")
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
		logger.Sugar().Errorw("DeleteAuth", "ID", in.GetID(), "error", err)
		return &npool.DeleteAuthResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "auth", "crud", "Delete")

	info, err := crud.Delete(ctx, id)
	if err != nil {
		logger.Sugar().Errorw("DeleteAuth", "ID", in.GetID(), "error", err)
		return &npool.DeleteAuthResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.DeleteAuthResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}
