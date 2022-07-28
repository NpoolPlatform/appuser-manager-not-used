//nolint:dupl
package banapp

import (
	"context"

	converter "github.com/NpoolPlatform/appuser-manager/pkg/converter/banapp"
	crud "github.com/NpoolPlatform/appuser-manager/pkg/crud/v2/banapp"
	constant "github.com/NpoolPlatform/appuser-manager/pkg/message/const"
	tracer "github.com/NpoolPlatform/appuser-manager/pkg/tracer/banapp"

	commontracer "github.com/NpoolPlatform/appuser-manager/pkg/tracer"

	"go.opentelemetry.io/otel"
	scodes "go.opentelemetry.io/otel/codes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/appuser/mgr/v2/banapp"

	"github.com/google/uuid"
)

func (s *Server) CreateBanApp(ctx context.Context, in *npool.CreateBanAppRequest) (*npool.CreateBanAppResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateBanApp")
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
		logger.Sugar().Errorw("CreateBanApp", "error", err)
		return &npool.CreateBanAppResponse{}, err
	}

	span = commontracer.TraceInvoker(span, "banapp", "crud", "Create")

	info, err := crud.Create(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorw("CreateBanApp", "error", err)
		return &npool.CreateBanAppResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateBanAppResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) CreateBanApps(ctx context.Context, in *npool.CreateBanAppsRequest) (*npool.CreateBanAppsResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateBanApps")
	defer span.End()
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	if len(in.GetInfos()) == 0 {
		logger.Sugar().Errorw("CreateBanApps", "error", "Infos is empty")
		return &npool.CreateBanAppsResponse{}, status.Error(codes.InvalidArgument, "Infos is empty")
	}

	if err := validateMany(in.GetInfos()); err != nil {
		logger.Sugar().Errorw("CreateBanApps", "error", err)
		return &npool.CreateBanAppsResponse{}, status.Error(codes.InvalidArgument, "Infos is invalid")
	}

	span = tracer.TraceMany(span, in.GetInfos())
	span = commontracer.TraceInvoker(span, "banapp", "crud", "CreateBulk")

	rows, err := crud.CreateBulk(ctx, in.GetInfos())
	if err != nil {
		logger.Sugar().Errorw("CreateBanApps", "error", err)
		return &npool.CreateBanAppsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateBanAppsResponse{
		Infos: converter.Ent2GrpcMany(rows),
	}, nil
}

func (s *Server) UpdateBanApp(ctx context.Context, in *npool.UpdateBanAppRequest) (*npool.UpdateBanAppResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "UpdateBanApp")
	defer span.End()
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.Trace(span, in.GetInfo())

	if _, err := uuid.Parse(in.GetInfo().GetID()); err != nil {
		logger.Sugar().Errorw("UpdateBanApp", "ID", in.GetInfo().GetID(), "error", err)
		return &npool.UpdateBanAppResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "banapp", "crud", "Update")

	info, err := crud.Update(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorw("UpdateBanApp", "ID", in.GetInfo().GetID(), "error", err)
		return &npool.UpdateBanAppResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.UpdateBanAppResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetBanApp(ctx context.Context, in *npool.GetBanAppRequest) (*npool.GetBanAppResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetBanApp")
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
		logger.Sugar().Errorw("GetBanApp", "ID", in.GetID(), "error", err)
		return &npool.GetBanAppResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "banapp", "crud", "Row")

	info, err := crud.Row(ctx, id)
	if err != nil {
		logger.Sugar().Errorw("GetBanApp", "ID", in.GetID(), "error", err)
		return &npool.GetBanAppResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetBanAppResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetBanAppOnly(ctx context.Context, in *npool.GetBanAppOnlyRequest) (*npool.GetBanAppOnlyResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetBanAppOnly")
	defer span.End()
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceInvoker(span, "banapp", "crud", "RowOnly")

	info, err := crud.RowOnly(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorw("GetBanAppOnly", "error", err)
		return &npool.GetBanAppOnlyResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetBanAppOnlyResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetBanApps(ctx context.Context, in *npool.GetBanAppsRequest) (*npool.GetBanAppsResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetBanApps")
	defer span.End()
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceOffsetLimit(span, int(in.GetOffset()), int(in.GetLimit()))
	span = commontracer.TraceInvoker(span, "banapp", "crud", "Rows")

	rows, total, err := crud.Rows(ctx, in.GetConds(), int(in.GetOffset()), int(in.GetLimit()))
	if err != nil {
		logger.Sugar().Errorw("GetBanApps", "error", err)
		return &npool.GetBanAppsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetBanAppsResponse{
		Infos: converter.Ent2GrpcMany(rows),
		Total: uint32(total),
	}, nil
}

func (s *Server) ExistBanApp(ctx context.Context, in *npool.ExistBanAppRequest) (*npool.ExistBanAppResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "ExistBanApp")
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
		logger.Sugar().Errorw("ExistBanApp", "ID", in.GetID(), "error", err)
		return &npool.ExistBanAppResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "banapp", "crud", "Exist")

	exist, err := crud.Exist(ctx, id)
	if err != nil {
		logger.Sugar().Errorw("ExistBanApp", "ID", in.GetID(), "error", err)
		return &npool.ExistBanAppResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistBanAppResponse{
		Info: exist,
	}, nil
}

func (s *Server) ExistBanAppConds(ctx context.Context, in *npool.ExistBanAppCondsRequest) (*npool.ExistBanAppCondsResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "ExistBanAppConds")
	defer span.End()
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceInvoker(span, "banapp", "crud", "ExistBanAppConds")

	exist, err := crud.ExistConds(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorw("ExistBanAppConds", "error", err)
		return &npool.ExistBanAppCondsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistBanAppCondsResponse{
		Info: exist,
	}, nil
}

func (s *Server) CountBanApps(ctx context.Context, in *npool.CountBanAppsRequest) (*npool.CountBanAppsResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CountBanApps")
	defer span.End()
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceInvoker(span, "banapp", "crud", "Count")

	total, err := crud.Count(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorw("Counts", "error", err)
		return &npool.CountBanAppsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CountBanAppsResponse{
		Info: total,
	}, nil
}

func (s *Server) DeleteBanApp(ctx context.Context, in *npool.DeleteBanAppRequest) (*npool.DeleteBanAppResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "DeleteBanApp")
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
		logger.Sugar().Errorw("DeleteBanApp", "ID", in.GetID(), "error", err)
		return &npool.DeleteBanAppResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "banapp", "crud", "Delete")

	info, err := crud.Delete(ctx, id)
	if err != nil {
		logger.Sugar().Errorw("DeleteBanApp", "ID", in.GetID(), "error", err)
		return &npool.DeleteBanAppResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.DeleteBanAppResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}
