//nolint:nolintlint,dupl
package subscriber

import (
	"context"

	converter "github.com/NpoolPlatform/appuser-manager/pkg/converter/v2/subscriber"
	crud "github.com/NpoolPlatform/appuser-manager/pkg/crud/v2/subscriber"
	constant "github.com/NpoolPlatform/appuser-manager/pkg/message/const"
	tracer "github.com/NpoolPlatform/appuser-manager/pkg/tracer/subscriber"

	commontracer "github.com/NpoolPlatform/appuser-manager/pkg/tracer"

	"go.opentelemetry.io/otel"
	scodes "go.opentelemetry.io/otel/codes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/appuser/mgr/v2/subscriber"

	"github.com/google/uuid"
)

func (s *Server) CreateSubscriber(ctx context.Context, in *npool.CreateSubscriberRequest) (*npool.CreateSubscriberResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateSubscriber")
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
		logger.Sugar().Errorw("CreateSubscriber", "error", err)
		return &npool.CreateSubscriberResponse{}, err
	}

	span = commontracer.TraceInvoker(span, "app", "crud", "Create")

	info, err := crud.Create(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorw("CreateSubscriber", "error", err)
		return &npool.CreateSubscriberResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateSubscriberResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) CreateSubscriberes(ctx context.Context, in *npool.CreateSubscriberesRequest) (*npool.CreateSubscriberesResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateSubscriberes")
	defer span.End()
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	if len(in.GetInfos()) == 0 {
		logger.Sugar().Errorw("CreateSubscriberes", "error", "Infos is empty")
		return &npool.CreateSubscriberesResponse{}, status.Error(codes.InvalidArgument, "Infos is empty")
	}

	if err := validateMany(in.GetInfos()); err != nil {
		logger.Sugar().Errorw("CreateSubscriberes", "error", err)
		return &npool.CreateSubscriberesResponse{}, status.Error(codes.InvalidArgument, "Infos is invalid")
	}

	span = tracer.TraceMany(span, in.GetInfos())
	span = commontracer.TraceInvoker(span, "app", "crud", "CreateBulk")

	rows, err := crud.CreateBulk(ctx, in.GetInfos())
	if err != nil {
		logger.Sugar().Errorw("CreateSubscriberes", "error", err)
		return &npool.CreateSubscriberesResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateSubscriberesResponse{
		Infos: converter.Ent2GrpcMany(rows),
	}, nil
}

func (s *Server) UpdateSubscriber(ctx context.Context, in *npool.UpdateSubscriberRequest) (*npool.UpdateSubscriberResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "UpdateSubscriber")
	defer span.End()
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.Trace(span, in.GetInfo())

	if _, err := uuid.Parse(in.GetInfo().GetID()); err != nil {
		logger.Sugar().Errorw("UpdateSubscriber", "ID", in.GetInfo().GetID(), "error", err)
		return &npool.UpdateSubscriberResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "app", "crud", "Update")

	info, err := crud.Update(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorw("UpdateSubscriber", "ID", in.GetInfo().GetID(), "error", err)
		return &npool.UpdateSubscriberResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.UpdateSubscriberResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetSubscriber(ctx context.Context, in *npool.GetSubscriberRequest) (*npool.GetSubscriberResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetSubscriber")
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
		logger.Sugar().Errorw("GetSubscriber", "ID", in.GetID(), "error", err)
		return &npool.GetSubscriberResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "app", "crud", "Row")

	info, err := crud.Row(ctx, id)
	if err != nil {
		logger.Sugar().Errorw("GetSubscriber", "ID", in.GetID(), "error", err)
		return &npool.GetSubscriberResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetSubscriberResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetSubscriberOnly(ctx context.Context, in *npool.GetSubscriberOnlyRequest) (*npool.GetSubscriberOnlyResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetSubscriberOnly")
	defer span.End()
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceInvoker(span, "app", "crud", "RowOnly")

	info, err := crud.RowOnly(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorw("GetSubscriberOnly", "error", err)
		return &npool.GetSubscriberOnlyResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetSubscriberOnlyResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetSubscriberes(ctx context.Context, in *npool.GetSubscriberesRequest) (*npool.GetSubscriberesResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetSubscriberes")
	defer span.End()
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceOffsetLimit(span, int(in.GetOffset()), int(in.GetLimit()))
	span = commontracer.TraceInvoker(span, "app", "crud", "Rows")

	rows, total, err := crud.Rows(ctx, in.GetConds(), int(in.GetOffset()), int(in.GetLimit()))
	if err != nil {
		logger.Sugar().Errorw("GetSubscriberes", "error", err)
		return &npool.GetSubscriberesResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetSubscriberesResponse{
		Infos: converter.Ent2GrpcMany(rows),
		Total: uint32(total),
	}, nil
}

func (s *Server) ExistSubscriber(ctx context.Context, in *npool.ExistSubscriberRequest) (*npool.ExistSubscriberResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "ExistSubscriber")
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
		logger.Sugar().Errorw("ExistSubscriber", "ID", in.GetID(), "error", err)
		return &npool.ExistSubscriberResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "app", "crud", "Exist")

	exist, err := crud.Exist(ctx, id)
	if err != nil {
		logger.Sugar().Errorw("ExistSubscriber", "ID", in.GetID(), "error", err)
		return &npool.ExistSubscriberResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistSubscriberResponse{
		Info: exist,
	}, nil
}

func (s *Server) ExistSubscriberConds(
	ctx context.Context,
	in *npool.ExistSubscriberCondsRequest,
) (
	*npool.ExistSubscriberCondsResponse,
	error,
) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "ExistSubscriberConds")
	defer span.End()
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceInvoker(span, "app", "crud", "ExistSubscriberConds")

	exist, err := crud.ExistConds(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorw("ExistSubscriberConds", "error", err)
		return &npool.ExistSubscriberCondsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistSubscriberCondsResponse{
		Info: exist,
	}, nil
}

func (s *Server) CountSubscriberes(ctx context.Context, in *npool.CountSubscriberesRequest) (*npool.CountSubscriberesResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CountSubscriberes")
	defer span.End()
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceInvoker(span, "app", "crud", "Count")

	total, err := crud.Count(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorw("Counts", "error", err)
		return &npool.CountSubscriberesResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CountSubscriberesResponse{
		Info: total,
	}, nil
}

func (s *Server) DeleteSubscriber(ctx context.Context, in *npool.DeleteSubscriberRequest) (*npool.DeleteSubscriberResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "DeleteSubscriber")
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
		logger.Sugar().Errorw("DeleteSubscriber", "ID", in.GetID(), "error", err)
		return &npool.DeleteSubscriberResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "app", "crud", "Delete")

	info, err := crud.Delete(ctx, id)
	if err != nil {
		logger.Sugar().Errorw("DeleteSubscriber", "ID", in.GetID(), "error", err)
		return &npool.DeleteSubscriberResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.DeleteSubscriberResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}
