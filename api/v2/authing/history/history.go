//nolint:nolintlint,dupl
package history

import (
	"context"

	converter "github.com/NpoolPlatform/appuser-manager/pkg/converter/v2/authing/history"
	crud "github.com/NpoolPlatform/appuser-manager/pkg/crud/v2/authing/history"
	constant "github.com/NpoolPlatform/appuser-manager/pkg/message/const"
	tracer "github.com/NpoolPlatform/appuser-manager/pkg/tracer/authing/history"

	commontracer "github.com/NpoolPlatform/appuser-manager/pkg/tracer"

	"go.opentelemetry.io/otel"
	scodes "go.opentelemetry.io/otel/codes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/appuser/mgr/v2/authing/history"

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

func (s *Server) CreateHistorys(ctx context.Context, in *npool.CreateHistorysRequest) (*npool.CreateHistorysResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateHistorys")
	defer span.End()
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	if len(in.GetInfos()) == 0 {
		logger.Sugar().Errorw("CreateHistorys", "error", "Infos is empty")
		return &npool.CreateHistorysResponse{}, status.Error(codes.InvalidArgument, "Infos is empty")
	}

	if err := validateMany(in.GetInfos()); err != nil {
		logger.Sugar().Errorw("CreateHistorys", "error", err)
		return &npool.CreateHistorysResponse{}, status.Error(codes.InvalidArgument, "Infos is invalid")
	}

	span = tracer.TraceMany(span, in.GetInfos())
	span = commontracer.TraceInvoker(span, "history", "crud", "CreateBulk")

	rows, err := crud.CreateBulk(ctx, in.GetInfos())
	if err != nil {
		logger.Sugar().Errorw("CreateHistorys", "error", err)
		return &npool.CreateHistorysResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateHistorysResponse{
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

func (s *Server) GetHistorys(ctx context.Context, in *npool.GetHistorysRequest) (*npool.GetHistorysResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetHistorys")
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
		logger.Sugar().Errorw("GetHistorys", "error", err)
		return &npool.GetHistorysResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetHistorysResponse{
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

func (s *Server) CountHistorys(ctx context.Context, in *npool.CountHistorysRequest) (*npool.CountHistorysResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CountHistorys")
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
		return &npool.CountHistorysResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CountHistorysResponse{
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
