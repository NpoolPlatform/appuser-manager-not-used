//nolint:dupl
package appuserextra

import (
	"context"

	converter "github.com/NpoolPlatform/appuser-manager/pkg/converter/appuserextra"
	crud "github.com/NpoolPlatform/appuser-manager/pkg/crud/v2/appuserextra"
	constant "github.com/NpoolPlatform/appuser-manager/pkg/message/const"
	tracer "github.com/NpoolPlatform/appuser-manager/pkg/tracer/appuserextra"

	commontracer "github.com/NpoolPlatform/appuser-manager/pkg/tracer"

	"go.opentelemetry.io/otel"
	scodes "go.opentelemetry.io/otel/codes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/appuser/mgr/v2/appuserextra"

	"github.com/google/uuid"
)

func (s *Server) CreateAppUserExtra(ctx context.Context, in *npool.CreateAppUserExtraRequest) (*npool.CreateAppUserExtraResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateAppUserExtra")
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
		logger.Sugar().Errorw("CreateAppUserExtra", "error", err)
		return &npool.CreateAppUserExtraResponse{}, err
	}

	span = commontracer.TraceInvoker(span, "appuserextra", "crud", "Create")

	info, err := crud.Create(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorw("CreateAppUserExtra", "error", err)
		return &npool.CreateAppUserExtraResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateAppUserExtraResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) CreateAppUserExtras(ctx context.Context, in *npool.CreateAppUserExtrasRequest) (*npool.CreateAppUserExtrasResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateAppUserExtras")
	defer span.End()
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	if len(in.GetInfos()) == 0 {
		logger.Sugar().Errorw("CreateAppUserExtras", "error", "Infos is empty")
		return &npool.CreateAppUserExtrasResponse{}, status.Error(codes.InvalidArgument, "Infos is empty")
	}

	if err := validateMany(in.GetInfos()); err != nil {
		logger.Sugar().Errorw("CreateAppUserExtras", "error", err)
		return &npool.CreateAppUserExtrasResponse{}, status.Error(codes.InvalidArgument, "Infos is invalid")
	}

	span = tracer.TraceMany(span, in.GetInfos())
	span = commontracer.TraceInvoker(span, "appuserextra", "crud", "CreateBulk")

	rows, err := crud.CreateBulk(ctx, in.GetInfos())
	if err != nil {
		logger.Sugar().Errorw("CreateAppUserExtras", "error", err)
		return &npool.CreateAppUserExtrasResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateAppUserExtrasResponse{
		Infos: converter.Ent2GrpcMany(rows),
	}, nil
}

func (s *Server) UpdateAppUserExtra(ctx context.Context, in *npool.UpdateAppUserExtraRequest) (*npool.UpdateAppUserExtraResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "UpdateAppUserExtra")
	defer span.End()
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.Trace(span, in.GetInfo())

	if _, err := uuid.Parse(in.GetInfo().GetID()); err != nil {
		logger.Sugar().Errorw("UpdateAppUserExtra", "ID", in.GetInfo().GetID(), "error", err)
		return &npool.UpdateAppUserExtraResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "appuserextra", "crud", "Update")

	info, err := crud.Update(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorw("UpdateAppUserExtra", "ID", in.GetInfo().GetID(), "error", err)
		return &npool.UpdateAppUserExtraResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.UpdateAppUserExtraResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetAppUserExtra(ctx context.Context, in *npool.GetAppUserExtraRequest) (*npool.GetAppUserExtraResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetAppUserExtra")
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
		logger.Sugar().Errorw("GetAppUserExtra", "ID", in.GetID(), "error", err)
		return &npool.GetAppUserExtraResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "appuserextra", "crud", "Row")

	info, err := crud.Row(ctx, id)
	if err != nil {
		logger.Sugar().Errorw("GetAppUserExtra", "ID", in.GetID(), "error", err)
		return &npool.GetAppUserExtraResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetAppUserExtraResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetAppUserExtraOnly(ctx context.Context, in *npool.GetAppUserExtraOnlyRequest) (*npool.GetAppUserExtraOnlyResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetAppUserExtraOnly")
	defer span.End()
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceInvoker(span, "appuserextra", "crud", "RowOnly")

	info, err := crud.RowOnly(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorw("GetAppUserExtraOnly", "error", err)
		return &npool.GetAppUserExtraOnlyResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetAppUserExtraOnlyResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetAppUserExtras(ctx context.Context, in *npool.GetAppUserExtrasRequest) (*npool.GetAppUserExtrasResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetAppUserExtras")
	defer span.End()
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceOffsetLimit(span, int(in.GetOffset()), int(in.GetLimit()))
	span = commontracer.TraceInvoker(span, "appuserextra", "crud", "Rows")

	rows, total, err := crud.Rows(ctx, in.GetConds(), int(in.GetOffset()), int(in.GetLimit()))
	if err != nil {
		logger.Sugar().Errorw("GetAppUserExtras", "error", err)
		return &npool.GetAppUserExtrasResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetAppUserExtrasResponse{
		Infos: converter.Ent2GrpcMany(rows),
		Total: uint32(total),
	}, nil
}

func (s *Server) ExistAppUserExtra(ctx context.Context, in *npool.ExistAppUserExtraRequest) (*npool.ExistAppUserExtraResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "ExistAppUserExtra")
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
		logger.Sugar().Errorw("ExistAppUserExtra", "ID", in.GetID(), "error", err)
		return &npool.ExistAppUserExtraResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "appuserextra", "crud", "Exist")

	exist, err := crud.Exist(ctx, id)
	if err != nil {
		logger.Sugar().Errorw("ExistAppUserExtra", "ID", in.GetID(), "error", err)
		return &npool.ExistAppUserExtraResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistAppUserExtraResponse{
		Info: exist,
	}, nil
}

func (s *Server) ExistAppUserExtraConds(ctx context.Context, in *npool.ExistAppUserExtraCondsRequest) (*npool.ExistAppUserExtraCondsResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "ExistAppUserExtraConds")
	defer span.End()
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceInvoker(span, "appuserextra", "crud", "ExistAppUserExtraConds")

	exist, err := crud.ExistConds(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorw("ExistAppUserExtraConds", "error", err)
		return &npool.ExistAppUserExtraCondsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistAppUserExtraCondsResponse{
		Info: exist,
	}, nil
}

func (s *Server) CountAppUserExtras(ctx context.Context, in *npool.CountAppUserExtrasRequest) (*npool.CountAppUserExtrasResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CountAppUserExtras")
	defer span.End()
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceInvoker(span, "appuserextra", "crud", "Count")

	total, err := crud.Count(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorw("Counts", "error", err)
		return &npool.CountAppUserExtrasResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CountAppUserExtrasResponse{
		Info: total,
	}, nil
}

func (s *Server) DeleteAppUserExtra(ctx context.Context, in *npool.DeleteAppUserExtraRequest) (*npool.DeleteAppUserExtraResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "DeleteAppUserExtra")
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
		logger.Sugar().Errorw("DeleteAppUserExtra", "ID", in.GetID(), "error", err)
		return &npool.DeleteAppUserExtraResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "appuserextra", "crud", "Delete")

	info, err := crud.Delete(ctx, id)
	if err != nil {
		logger.Sugar().Errorw("DeleteAppUserExtra", "ID", in.GetID(), "error", err)
		return &npool.DeleteAppUserExtraResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.DeleteAppUserExtraResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}
