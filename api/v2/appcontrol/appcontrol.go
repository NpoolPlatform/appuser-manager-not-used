//nolint:dupl
package appcontrol

import (
	"context"

	converter "github.com/NpoolPlatform/appuser-manager/pkg/converter/appcontrol"
	crud "github.com/NpoolPlatform/appuser-manager/pkg/crud/v2/appcontrol"
	constant "github.com/NpoolPlatform/appuser-manager/pkg/message/const"
	tracer "github.com/NpoolPlatform/appuser-manager/pkg/tracer/appcontrol"

	commontracer "github.com/NpoolPlatform/appuser-manager/pkg/tracer"

	"go.opentelemetry.io/otel"
	scodes "go.opentelemetry.io/otel/codes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/appuser/mgr/v2/appcontrol"

	"github.com/google/uuid"
)

func (s *Server) CreateAppControl(ctx context.Context, in *npool.CreateAppControlRequest) (*npool.CreateAppControlResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateAppControl")
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
		logger.Sugar().Errorw("CreateAppControl", "error", err)
		return &npool.CreateAppControlResponse{}, err
	}

	span = commontracer.TraceInvoker(span, "appcontrol", "crud", "Create")

	info, err := crud.Create(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorw("CreateAppControl", "error", err)
		return &npool.CreateAppControlResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateAppControlResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) CreateAppControls(ctx context.Context, in *npool.CreateAppControlsRequest) (*npool.CreateAppControlsResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateAppControls")
	defer span.End()
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	if len(in.GetInfos()) == 0 {
		logger.Sugar().Errorw("CreateAppControls", "error", "Infos is empty")
		return &npool.CreateAppControlsResponse{}, status.Error(codes.InvalidArgument, "Infos is empty")
	}

	span = tracer.TraceMany(span, in.GetInfos())
	span = commontracer.TraceInvoker(span, "appcontrol", "crud", "CreateBulk")

	rows, err := crud.CreateBulk(ctx, in.GetInfos())
	if err != nil {
		logger.Sugar().Errorw("CreateAppControls", "error", err)
		return &npool.CreateAppControlsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateAppControlsResponse{
		Infos: converter.Ent2GrpcMany(rows),
	}, nil
}

func (s *Server) UpdateAppControl(ctx context.Context, in *npool.UpdateAppControlRequest) (*npool.UpdateAppControlResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "UpdateAppControl")
	defer span.End()
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.Trace(span, in.GetInfo())

	if _, err := uuid.Parse(in.GetInfo().GetID()); err != nil {
		logger.Sugar().Errorw("UpdateAppControl", "ID", in.GetInfo().GetID(), "error", err)
		return &npool.UpdateAppControlResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "appcontrol", "crud", "Update")

	info, err := crud.Update(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorw("UpdateAppControl", "ID", in.GetInfo().GetID(), "error", err)
		return &npool.UpdateAppControlResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.UpdateAppControlResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetAppControl(ctx context.Context, in *npool.GetAppControlRequest) (*npool.GetAppControlResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetAppControl")
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
		logger.Sugar().Errorw("GetAppControl", "ID", in.GetID(), "error", err)
		return &npool.GetAppControlResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "appcontrol", "crud", "Row")

	info, err := crud.Row(ctx, id)
	if err != nil {
		logger.Sugar().Errorw("GetAppControl", "ID", in.GetID(), "error", err)
		return &npool.GetAppControlResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetAppControlResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetAppControlOnly(ctx context.Context, in *npool.GetAppControlOnlyRequest) (*npool.GetAppControlOnlyResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetAppControlOnly")
	defer span.End()
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceInvoker(span, "appcontrol", "crud", "RowOnly")

	info, err := crud.RowOnly(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorw("GetAppControlOnly", "error", err)
		return &npool.GetAppControlOnlyResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetAppControlOnlyResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetAppControls(ctx context.Context, in *npool.GetAppControlsRequest) (*npool.GetAppControlsResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetAppControls")
	defer span.End()
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceOffsetLimit(span, int(in.GetOffset()), int(in.GetLimit()))
	span = commontracer.TraceInvoker(span, "appcontrol", "crud", "Rows")

	rows, total, err := crud.Rows(ctx, in.GetConds(), int(in.GetOffset()), int(in.GetLimit()))
	if err != nil {
		logger.Sugar().Errorw("GetAppControls", "error", err)
		return &npool.GetAppControlsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetAppControlsResponse{
		Infos: converter.Ent2GrpcMany(rows),
		Total: uint32(total),
	}, nil
}

func (s *Server) ExistAppControl(ctx context.Context, in *npool.ExistAppControlRequest) (*npool.ExistAppControlResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "ExistAppControl")
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
		logger.Sugar().Errorw("ExistAppControl", "ID", in.GetID(), "error", err)
		return &npool.ExistAppControlResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "appcontrol", "crud", "Exist")

	exist, err := crud.Exist(ctx, id)
	if err != nil {
		logger.Sugar().Errorw("ExistAppControl", "ID", in.GetID(), "error", err)
		return &npool.ExistAppControlResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistAppControlResponse{
		Info: exist,
	}, nil
}

func (s *Server) ExistAppControlConds(ctx context.Context, in *npool.ExistAppControlCondsRequest) (*npool.ExistAppControlCondsResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "ExistAppControlConds")
	defer span.End()
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceInvoker(span, "appcontrol", "crud", "ExistAppControlConds")

	exist, err := crud.ExistConds(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorw("ExistAppControlConds", "error", err)
		return &npool.ExistAppControlCondsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistAppControlCondsResponse{
		Info: exist,
	}, nil
}

func (s *Server) CountAppControls(ctx context.Context, in *npool.CountAppControlsRequest) (*npool.CountAppControlsResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CountAppControls")
	defer span.End()
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceInvoker(span, "appcontrol", "crud", "Count")

	total, err := crud.Count(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorw("Counts", "error", err)
		return &npool.CountAppControlsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CountAppControlsResponse{
		Info: total,
	}, nil
}

func (s *Server) DeleteAppControl(ctx context.Context, in *npool.DeleteAppControlRequest) (*npool.DeleteAppControlResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "DeleteAppControl")
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
		logger.Sugar().Errorw("DeleteAppControl", "ID", in.GetID(), "error", err)
		return &npool.DeleteAppControlResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "appcontrol", "crud", "Delete")

	info, err := crud.Delete(ctx, id)
	if err != nil {
		logger.Sugar().Errorw("DeleteAppControl", "ID", in.GetID(), "error", err)
		return &npool.DeleteAppControlResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.DeleteAppControlResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}
