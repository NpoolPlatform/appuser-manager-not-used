//nolint:nolintlint,dupl
package appuserthirdparty

import (
	"context"

	converter "github.com/NpoolPlatform/appuser-manager/pkg/converter/v2/appuserthirdparty"
	crud "github.com/NpoolPlatform/appuser-manager/pkg/crud/v2/appuserthirdparty"
	constant "github.com/NpoolPlatform/appuser-manager/pkg/message/const"
	tracer "github.com/NpoolPlatform/appuser-manager/pkg/tracer/appuserthirdparty"

	commontracer "github.com/NpoolPlatform/appuser-manager/pkg/tracer"

	"go.opentelemetry.io/otel"
	scodes "go.opentelemetry.io/otel/codes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/appuser/mgr/v2/appuserthirdparty"

	"github.com/google/uuid"
)

func (s *Server) CreateAppUserThirdParty(ctx context.Context,
	in *npool.CreateAppUserThirdPartyRequest) (*npool.CreateAppUserThirdPartyResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateAppUserThirdParty")
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
		logger.Sugar().Errorw("CreateAppUserThirdParty", "error", err)
		return &npool.CreateAppUserThirdPartyResponse{}, err
	}

	span = commontracer.TraceInvoker(span, "appuserthirdparty", "crud", "Create")

	info, err := crud.Create(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorw("CreateAppUserThirdParty", "error", err)
		return &npool.CreateAppUserThirdPartyResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateAppUserThirdPartyResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) CreateAppUserThirdParties(ctx context.Context,
	in *npool.CreateAppUserThirdPartiesRequest) (*npool.CreateAppUserThirdPartiesResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateAppUserThirdParties")
	defer span.End()
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	if len(in.GetInfos()) == 0 {
		logger.Sugar().Errorw("CreateAppUserThirdParties", "error", "Infos is empty")
		return &npool.CreateAppUserThirdPartiesResponse{}, status.Error(codes.InvalidArgument, "Infos is empty")
	}

	span = tracer.TraceMany(span, in.GetInfos())
	span = commontracer.TraceInvoker(span, "appuserthirdparty", "crud", "CreateBulk")

	rows, err := crud.CreateBulk(ctx, in.GetInfos())
	if err != nil {
		logger.Sugar().Errorw("CreateAppUserThirdParties", "error", err)
		return &npool.CreateAppUserThirdPartiesResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateAppUserThirdPartiesResponse{
		Infos: converter.Ent2GrpcMany(rows),
	}, nil
}

func (s *Server) UpdateAppUserThirdParty(ctx context.Context,
	in *npool.UpdateAppUserThirdPartyRequest) (*npool.UpdateAppUserThirdPartyResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "UpdateAppUserThirdParty")
	defer span.End()
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.Trace(span, in.GetInfo())

	if _, err := uuid.Parse(in.GetInfo().GetID()); err != nil {
		logger.Sugar().Errorw("UpdateAppUserThirdParty", "ID", in.GetInfo().GetID(), "error", err)
		return &npool.UpdateAppUserThirdPartyResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "appuserthirdparty", "crud", "Update")

	info, err := crud.Update(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorw("UpdateAppUserThirdParty", "ID", in.GetInfo().GetID(), "error", err)
		return &npool.UpdateAppUserThirdPartyResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.UpdateAppUserThirdPartyResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetAppUserThirdParty(ctx context.Context,
	in *npool.GetAppUserThirdPartyRequest) (*npool.GetAppUserThirdPartyResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetAppUserThirdParty")
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
		logger.Sugar().Errorw("GetAppUserThirdParty", "ID", in.GetID(), "error", err)
		return &npool.GetAppUserThirdPartyResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "appuserthirdparty", "crud", "Row")

	info, err := crud.Row(ctx, id)
	if err != nil {
		logger.Sugar().Errorw("GetAppUserThirdParty", "ID", in.GetID(), "error", err)
		return &npool.GetAppUserThirdPartyResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetAppUserThirdPartyResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetAppUserThirdPartyOnly(ctx context.Context,
	in *npool.GetAppUserThirdPartyOnlyRequest) (*npool.GetAppUserThirdPartyOnlyResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetAppUserThirdPartyOnly")
	defer span.End()
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceInvoker(span, "appuserthirdparty", "crud", "RowOnly")

	info, err := crud.RowOnly(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorw("GetAppUserThirdPartyOnly", "error", err)
		return &npool.GetAppUserThirdPartyOnlyResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetAppUserThirdPartyOnlyResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetAppUserThirdParties(ctx context.Context,
	in *npool.GetAppUserThirdPartiesRequest) (*npool.GetAppUserThirdPartiesResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetAppUserThirdParties")
	defer span.End()
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceOffsetLimit(span, int(in.GetOffset()), int(in.GetLimit()))
	span = commontracer.TraceInvoker(span, "appuserthirdparty", "crud", "Rows")

	rows, total, err := crud.Rows(ctx, in.GetConds(), int(in.GetOffset()), int(in.GetLimit()))
	if err != nil {
		logger.Sugar().Errorw("GetAppUserThirdParties", "error", err)
		return &npool.GetAppUserThirdPartiesResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetAppUserThirdPartiesResponse{
		Infos: converter.Ent2GrpcMany(rows),
		Total: uint32(total),
	}, nil
}

func (s *Server) ExistAppUserThirdParty(ctx context.Context,
	in *npool.ExistAppUserThirdPartyRequest) (*npool.ExistAppUserThirdPartyResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "ExistAppUserThirdParty")
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
		logger.Sugar().Errorw("ExistAppUserThirdParty", "ID", in.GetID(), "error", err)
		return &npool.ExistAppUserThirdPartyResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "appuserthirdparty", "crud", "Exist")

	exist, err := crud.Exist(ctx, id)
	if err != nil {
		logger.Sugar().Errorw("ExistAppUserThirdParty", "ID", in.GetID(), "error", err)
		return &npool.ExistAppUserThirdPartyResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistAppUserThirdPartyResponse{
		Info: exist,
	}, nil
}

func (s *Server) ExistAppUserThirdPartyConds(ctx context.Context,
	in *npool.ExistAppUserThirdPartyCondsRequest) (*npool.ExistAppUserThirdPartyCondsResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "ExistAppUserThirdPartyConds")
	defer span.End()
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceInvoker(span, "appuserthirdparty", "crud", "ExistAppUserThirdPartyConds")

	exist, err := crud.ExistConds(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorw("ExistAppUserThirdPartyConds", "error", err)
		return &npool.ExistAppUserThirdPartyCondsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistAppUserThirdPartyCondsResponse{
		Info: exist,
	}, nil
}

func (s *Server) CountAppUserThirdParties(ctx context.Context,
	in *npool.CountAppUserThirdPartiesRequest) (*npool.CountAppUserThirdPartiesResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CountAppUserThirdParties")
	defer span.End()
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceInvoker(span, "appuserthirdparty", "crud", "Count")

	total, err := crud.Count(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorw("Counts", "error", err)
		return &npool.CountAppUserThirdPartiesResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CountAppUserThirdPartiesResponse{
		Info: total,
	}, nil
}

func (s *Server) DeleteAppUserThirdParty(ctx context.Context,
	in *npool.DeleteAppUserThirdPartyRequest) (*npool.DeleteAppUserThirdPartyResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "DeleteAppUserThirdParty")
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
		logger.Sugar().Errorw("DeleteAppUserThirdParty", "ID", in.GetID(), "error", err)
		return &npool.DeleteAppUserThirdPartyResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "appuserthirdparty", "crud", "Delete")

	info, err := crud.Delete(ctx, id)
	if err != nil {
		logger.Sugar().Errorw("DeleteAppUserThirdParty", "ID", in.GetID(), "error", err)
		return &npool.DeleteAppUserThirdPartyResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.DeleteAppUserThirdPartyResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}
