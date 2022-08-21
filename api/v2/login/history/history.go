//nolint:nolintlint,dupl
package history

import (
	"context"

	converter "github.com/NpoolPlatform/appuser-manager/pkg/converter/v2/login/history"
	crud "github.com/NpoolPlatform/appuser-manager/pkg/crud/v2/login/history"
	constant "github.com/NpoolPlatform/appuser-manager/pkg/message/const"
	tracer "github.com/NpoolPlatform/appuser-manager/pkg/tracer/login/history"

	commontracer "github.com/NpoolPlatform/appuser-manager/pkg/tracer"

	"go.opentelemetry.io/otel"
	scodes "go.opentelemetry.io/otel/codes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/appuser/mgr/v2/login/history"

	"github.com/google/uuid"
)

func (s *Server) CreateHistory(ctx context.Context, in *npool.CreateHistoryRequest) (*npool.CreateHistoryResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateHistory")
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
		logger.Sugar().Errorw("CreateHistory", "error", err)
		return &npool.CreateHistoryResponse{}, err
	}

	span = commontracer.TraceInvoker(span, "history", "crud", "Create")

	info, err := crud.Create(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorw("CreateHistory", "error", err)
		return &npool.CreateHistoryResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateHistoryResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) CreateHistories(ctx context.Context, in *npool.CreateHistoriesRequest) (*npool.CreateHistoriesResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateHistories")
	defer span.End()
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	if len(in.GetInfos()) == 0 {
		logger.Sugar().Errorw("CreateHistories", "error", "Infos is empty")
		return &npool.CreateHistoriesResponse{}, status.Error(codes.InvalidArgument, "Infos is empty")
	}

	if err := validateMany(in.GetInfos()); err != nil {
		logger.Sugar().Errorw("CreateHistories", "error", err)
		return &npool.CreateHistoriesResponse{}, status.Error(codes.InvalidArgument, "Infos is invalid")
	}

	span = tracer.TraceMany(span, in.GetInfos())
	span = commontracer.TraceInvoker(span, "history", "crud", "CreateBulk")

	rows, err := crud.CreateBulk(ctx, in.GetInfos())
	if err != nil {
		logger.Sugar().Errorw("CreateHistories", "error", err)
		return &npool.CreateHistoriesResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateHistoriesResponse{
		Infos: converter.Ent2GrpcMany(rows),
	}, nil
}

func (s *Server) UpdateHistory(ctx context.Context, in *npool.UpdateHistoryRequest) (*npool.UpdateHistoryResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "UpdateHistory")
	defer span.End()
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.Trace(span, in.GetInfo())

	if _, err := uuid.Parse(in.GetInfo().GetID()); err != nil {
		logger.Sugar().Errorw("UpdateHistory", "ID", in.GetInfo().GetID(), "error", err)
		return &npool.UpdateHistoryResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "history", "crud", "Update")

	info, err := crud.Update(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorw("UpdateHistory", "ID", in.GetInfo().GetID(), "error", err)
		return &npool.UpdateHistoryResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.UpdateHistoryResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetHistory(ctx context.Context, in *npool.GetHistoryRequest) (*npool.GetHistoryResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetHistory")
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
		logger.Sugar().Errorw("GetHistory", "ID", in.GetID(), "error", err)
		return &npool.GetHistoryResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "history", "crud", "Row")

	info, err := crud.Row(ctx, id)
	if err != nil {
		logger.Sugar().Errorw("GetHistory", "ID", in.GetID(), "error", err)
		return &npool.GetHistoryResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetHistoryResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetHistoryOnly(ctx context.Context, in *npool.GetHistoryOnlyRequest) (*npool.GetHistoryOnlyResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetHistoryOnly")
	defer span.End()
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceInvoker(span, "history", "crud", "RowOnly")

	info, err := crud.RowOnly(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorw("GetHistoryOnly", "error", err)
		return &npool.GetHistoryOnlyResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetHistoryOnlyResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetHistories(ctx context.Context, in *npool.GetHistoriesRequest) (*npool.GetHistoriesResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetHistories")
	defer span.End()
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceOffsetLimit(span, int(in.GetOffset()), int(in.GetLimit()))
	span = commontracer.TraceInvoker(span, "history", "crud", "Rows")

	rows, total, err := crud.Rows(ctx, in.GetConds(), int(in.GetOffset()), int(in.GetLimit()))
	if err != nil {
		logger.Sugar().Errorw("GetHistories", "error", err)
		return &npool.GetHistoriesResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetHistoriesResponse{
		Infos: converter.Ent2GrpcMany(rows),
		Total: uint32(total),
	}, nil
}

func (s *Server) ExistHistory(ctx context.Context, in *npool.ExistHistoryRequest) (*npool.ExistHistoryResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "ExistHistory")
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
		logger.Sugar().Errorw("ExistHistory", "ID", in.GetID(), "error", err)
		return &npool.ExistHistoryResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "history", "crud", "Exist")

	exist, err := crud.Exist(ctx, id)
	if err != nil {
		logger.Sugar().Errorw("ExistHistory", "ID", in.GetID(), "error", err)
		return &npool.ExistHistoryResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistHistoryResponse{
		Info: exist,
	}, nil
}

func (s *Server) ExistHistoryConds(ctx context.Context, in *npool.ExistHistoryCondsRequest) (*npool.ExistHistoryCondsResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "ExistHistoryConds")
	defer span.End()
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceInvoker(span, "history", "crud", "ExistHistoryConds")

	exist, err := crud.ExistConds(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorw("ExistHistoryConds", "error", err)
		return &npool.ExistHistoryCondsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistHistoryCondsResponse{
		Info: exist,
	}, nil
}

func (s *Server) CountHistories(ctx context.Context, in *npool.CountHistoriesRequest) (*npool.CountHistoriesResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CountHistories")
	defer span.End()
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceInvoker(span, "history", "crud", "Count")

	total, err := crud.Count(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorw("Counts", "error", err)
		return &npool.CountHistoriesResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CountHistoriesResponse{
		Info: total,
	}, nil
}

func (s *Server) DeleteHistory(ctx context.Context, in *npool.DeleteHistoryRequest) (*npool.DeleteHistoryResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "DeleteHistory")
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
		logger.Sugar().Errorw("DeleteHistory", "ID", in.GetID(), "error", err)
		return &npool.DeleteHistoryResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "history", "crud", "Delete")

	info, err := crud.Delete(ctx, id)
	if err != nil {
		logger.Sugar().Errorw("DeleteHistory", "ID", in.GetID(), "error", err)
		return &npool.DeleteHistoryResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.DeleteHistoryResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}
