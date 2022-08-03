//nolint:nolintlint,dupl
package appusercontrol

import (
	"context"

	converter "github.com/NpoolPlatform/appuser-manager/pkg/converter/v2/appusercontrol"
	crud "github.com/NpoolPlatform/appuser-manager/pkg/crud/v2/appusercontrol"
	constant "github.com/NpoolPlatform/appuser-manager/pkg/message/const"
	tracer "github.com/NpoolPlatform/appuser-manager/pkg/tracer/appusercontrol"

	commontracer "github.com/NpoolPlatform/appuser-manager/pkg/tracer"

	"go.opentelemetry.io/otel"
	scodes "go.opentelemetry.io/otel/codes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/appuser/mgr/v2/appusercontrol"

	"github.com/google/uuid"
)

func (s *Server) CreateAppUserControl(ctx context.Context,
	in *npool.CreateAppUserControlRequest) (*npool.CreateAppUserControlResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateAppUserControl")
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
		logger.Sugar().Errorw("CreateAppUserControl", "error", err)
		return &npool.CreateAppUserControlResponse{}, err
	}

	span = commontracer.TraceInvoker(span, "appusercontrol", "crud", "Create")

	info, err := crud.Create(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorw("CreateAppUserControl", "error", err)
		return &npool.CreateAppUserControlResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateAppUserControlResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) CreateAppUserControls(ctx context.Context,
	in *npool.CreateAppUserControlsRequest) (*npool.CreateAppUserControlsResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateAppUserControls")
	defer span.End()
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	if len(in.GetInfos()) == 0 {
		logger.Sugar().Errorw("CreateAppUserControls", "error", "Infos is empty")
		return &npool.CreateAppUserControlsResponse{}, status.Error(codes.InvalidArgument, "Infos is empty")
	}

	span = tracer.TraceMany(span, in.GetInfos())
	span = commontracer.TraceInvoker(span, "appusercontrol", "crud", "CreateBulk")

	rows, err := crud.CreateBulk(ctx, in.GetInfos())
	if err != nil {
		logger.Sugar().Errorw("CreateAppUserControls", "error", err)
		return &npool.CreateAppUserControlsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateAppUserControlsResponse{
		Infos: converter.Ent2GrpcMany(rows),
	}, nil
}

func (s *Server) UpdateAppUserControl(ctx context.Context,
	in *npool.UpdateAppUserControlRequest) (*npool.UpdateAppUserControlResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "UpdateAppUserControl")
	defer span.End()
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.Trace(span, in.GetInfo())

	if _, err := uuid.Parse(in.GetInfo().GetID()); err != nil {
		logger.Sugar().Errorw("UpdateAppUserControl", "ID", in.GetInfo().GetID(), "error", err)
		return &npool.UpdateAppUserControlResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "appusercontrol", "crud", "Update")

	info, err := crud.Update(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorw("UpdateAppUserControl", "ID", in.GetInfo().GetID(), "error", err)
		return &npool.UpdateAppUserControlResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.UpdateAppUserControlResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetAppUserControl(ctx context.Context, in *npool.GetAppUserControlRequest) (*npool.GetAppUserControlResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetAppUserControl")
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
		logger.Sugar().Errorw("GetAppUserControl", "ID", in.GetID(), "error", err)
		return &npool.GetAppUserControlResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "appusercontrol", "crud", "Row")

	info, err := crud.Row(ctx, id)
	if err != nil {
		logger.Sugar().Errorw("GetAppUserControl", "ID", in.GetID(), "error", err)
		return &npool.GetAppUserControlResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetAppUserControlResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetAppUserControlOnly(ctx context.Context,
	in *npool.GetAppUserControlOnlyRequest) (*npool.GetAppUserControlOnlyResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetAppUserControlOnly")
	defer span.End()
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceInvoker(span, "appusercontrol", "crud", "RowOnly")

	info, err := crud.RowOnly(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorw("GetAppUserControlOnly", "error", err)
		return &npool.GetAppUserControlOnlyResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetAppUserControlOnlyResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetAppUserControls(ctx context.Context, in *npool.GetAppUserControlsRequest) (*npool.GetAppUserControlsResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetAppUserControls")
	defer span.End()
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceOffsetLimit(span, int(in.GetOffset()), int(in.GetLimit()))
	span = commontracer.TraceInvoker(span, "appusercontrol", "crud", "Rows")

	rows, total, err := crud.Rows(ctx, in.GetConds(), int(in.GetOffset()), int(in.GetLimit()))
	if err != nil {
		logger.Sugar().Errorw("GetAppUserControls", "error", err)
		return &npool.GetAppUserControlsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetAppUserControlsResponse{
		Infos: converter.Ent2GrpcMany(rows),
		Total: uint32(total),
	}, nil
}

func (s *Server) ExistAppUserControl(ctx context.Context,
	in *npool.ExistAppUserControlRequest) (*npool.ExistAppUserControlResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "ExistAppUserControl")
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
		logger.Sugar().Errorw("ExistAppUserControl", "ID", in.GetID(), "error", err)
		return &npool.ExistAppUserControlResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "appusercontrol", "crud", "Exist")

	exist, err := crud.Exist(ctx, id)
	if err != nil {
		logger.Sugar().Errorw("ExistAppUserControl", "ID", in.GetID(), "error", err)
		return &npool.ExistAppUserControlResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistAppUserControlResponse{
		Info: exist,
	}, nil
}

func (s *Server) ExistAppUserControlConds(ctx context.Context,
	in *npool.ExistAppUserControlCondsRequest) (*npool.ExistAppUserControlCondsResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "ExistAppUserControlConds")
	defer span.End()
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceInvoker(span, "appusercontrol", "crud", "ExistAppUserControlConds")

	exist, err := crud.ExistConds(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorw("ExistAppUserControlConds", "error", err)
		return &npool.ExistAppUserControlCondsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistAppUserControlCondsResponse{
		Info: exist,
	}, nil
}

func (s *Server) CountAppUserControls(ctx context.Context,
	in *npool.CountAppUserControlsRequest) (*npool.CountAppUserControlsResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CountAppUserControls")
	defer span.End()
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceInvoker(span, "appusercontrol", "crud", "Count")

	total, err := crud.Count(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorw("Counts", "error", err)
		return &npool.CountAppUserControlsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CountAppUserControlsResponse{
		Info: total,
	}, nil
}

func (s *Server) DeleteAppUserControl(ctx context.Context,
	in *npool.DeleteAppUserControlRequest) (*npool.DeleteAppUserControlResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "DeleteAppUserControl")
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
		logger.Sugar().Errorw("DeleteAppUserControl", "ID", in.GetID(), "error", err)
		return &npool.DeleteAppUserControlResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "appusercontrol", "crud", "Delete")

	info, err := crud.Delete(ctx, id)
	if err != nil {
		logger.Sugar().Errorw("DeleteAppUserControl", "ID", in.GetID(), "error", err)
		return &npool.DeleteAppUserControlResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.DeleteAppUserControlResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}
