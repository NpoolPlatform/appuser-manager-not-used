//go:build !codeanalysis
// +build !codeanalysis

package app

import (
	"context"
	"fmt"

	converter "github.com/NpoolPlatform/appuser-manager/pkg/converter/app"
	tracer "github.com/NpoolPlatform/appuser-manager/pkg/converter/app"
	crud "github.com/NpoolPlatform/appuser-manager/pkg/crud/v2/app"
	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent"
	constant "github.com/NpoolPlatform/appuser-manager/pkg/message/const"

	commontracer "github.com/NpoolPlatform/appuser-manager/pkg/tracer"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	scodes "go.opentelemetry.io/otel/codes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/appuser/mgr/v2/app"

	"github.com/google/uuid"
)

func (s *Server) CreateApp(ctx context.Context, in *npool.CreateAppRequest) (*npool.CreateAppResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateApp")
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
		logger.Sguar().Errorw("CreateApp", "error", err)
		return &npool.CreateAppResponse{}, err
	}

	span = commontracer.TraceInvoker(span, "app", "crud", "Create")

	info, err := crud.Create(ctx, in.GetInfo())
	if err != nil {
		logger.Sguar().Errorw("CreateApp", "error", err)
		return &npool.CreateAppResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateAppResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) CreateApps(ctx context.Context, in *npool.CreateAppsRequest) (*npool.CreateAppsResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateApps")
	defer span.End()
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	if len(in.GetInfos()) == 0 {
		logger.Sugar().Errorw("CreateApps", "error", "Infos is empty")
		return &npool.CreateAppsResponse{}, status.Error(codes.InvalidArgument, "Infos is empty")
	}

	if err := validateMany(in.GetInfos()); err != nil {
		logger.Sugar().Errorw("CreateApps", "error", err)
		return &npool.CreateAppsResponse{}, status.Error(codes.InvalidArgument, "Infos is invalid")
	}

	span = tracer.TraceMany(span, in.GetInfos())
	span = commontracer.TraceInvoker(span, "app", "crud", "CreateBulk")

	rows, err := crud.CreateBulk(ctx, in.GetInfos())
	if err != nil {
		logger.Sugar().Errorw("CreateApps", "error", err)
		return &npool.CreateAppsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateAppsResponse{
		Infos: converter.Ent2GrpcMany(rows),
	}, nil
}

func (s *Server) UpdateApp(ctx context.Context, in *npool.UpdateAppRequest) (*npool.UpdateAppResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "UpdateApp")
	defer span.End()
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = crud.AppSpanAttributes(span, in.GetInfo())

	if _, err := uuid.Parse(in.GetInfo().GetID()); err != nil {
		logger.Sugar().Errorf("app id is invalid")
		return &npool.UpdateAppResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span.AddEvent("call crud Update")
	info, err := crud.Update(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorf("fail update app: %v", err)
		return &npool.UpdateAppResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.UpdateAppResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetApp(ctx context.Context, in *npool.GetAppRequest) (*npool.GetAppResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetApp")
	defer span.End()
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span.SetAttributes(
		attribute.String("ID", in.GetID()),
	)

	id, err := uuid.Parse(in.GetID())
	if err != nil {
		return &npool.GetAppResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span.AddEvent("call crud Row")
	info, err := crud.Row(ctx, id)
	if err != nil {
		logger.Sugar().Errorf("fail get App: %v", err)
		return &npool.GetAppResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetAppResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetAppOnly(ctx context.Context, in *npool.GetAppOnlyRequest) (*npool.GetAppOnlyResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetAppOnly")
	defer span.End()
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = crud.AppCondsSpanAttributes(span, in.GetConds())

	span.AddEvent("call crud RowOnly")
	info, err := crud.RowOnly(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorf("fail get Apps: %v", err)
		return &npool.GetAppOnlyResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetAppOnlyResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetApps(ctx context.Context, in *npool.GetAppsRequest) (*npool.GetAppsResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetApps")
	defer span.End()
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = crud.AppCondsSpanAttributes(span, in.GetConds())
	span.SetAttributes(
		attribute.Int("Offset", int(in.GetOffset())),
		attribute.Int("Limit", int(in.GetLimit())),
	)

	span.AddEvent("call crud Rows")
	rows, total, err := crud.Rows(ctx, in.GetConds(), int(in.GetOffset()), int(in.GetLimit()))
	if err != nil {
		logger.Sugar().Errorf("fail get Apps: %v", err)
		return &npool.GetAppsResponse{}, status.Error(codes.Internal, err.Error())
	}

	infos := make([]*npool.App, 0, len(rows))
	for _, val := range rows {
		infos = append(infos, converter.Ent2Grpc(val))
	}

	return &npool.GetAppsResponse{
		Infos: infos,
		Total: uint32(total),
	}, nil
}

func (s *Server) ExistApp(ctx context.Context, in *npool.ExistAppRequest) (*npool.ExistAppResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "ExistApp")
	defer span.End()
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span.SetAttributes(
		attribute.String("ID", in.GetID()),
	)

	id, err := uuid.Parse(in.GetID())
	if err != nil {
		return &npool.ExistAppResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span.AddEvent("call crud Exist")
	exist, err := crud.Exist(ctx, id)
	if err != nil {
		logger.Sugar().Errorf("fail check App: %v", err)
		return &npool.ExistAppResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistAppResponse{
		Info: exist,
	}, nil
}

func (s *Server) ExistAppConds(ctx context.Context, in *npool.ExistAppCondsRequest) (*npool.ExistAppCondsResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "ExistAppConds")
	defer span.End()
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = crud.AppCondsSpanAttributes(span, in.GetConds())

	span.AddEvent("call crud ExistConds")
	exist, err := crud.ExistConds(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorf("fail check App: %v", err)
		return &npool.ExistAppCondsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistAppCondsResponse{
		Info: exist,
	}, nil
}

func (s *Server) CountApps(ctx context.Context, in *npool.CountAppsRequest) (*npool.CountAppsResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CountApps")
	defer span.End()
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = crud.AppCondsSpanAttributes(span, in.GetConds())

	span.AddEvent("call crud Count")
	total, err := crud.Count(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorf("fail count Apps: %v", err)
		return &npool.CountAppsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CountAppsResponse{
		Info: total,
	}, nil
}

func (s *Server) DeleteApp(ctx context.Context, in *npool.DeleteAppRequest) (*npool.DeleteAppResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "DeleteApp")
	defer span.End()
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span.SetAttributes(
		attribute.String("ID", in.GetID()),
	)

	id, err := uuid.Parse(in.GetID())
	if err != nil {
		return &npool.DeleteAppResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span.AddEvent("call crud Delete")
	info, err := crud.Delete(ctx, id)
	if err != nil {
		logger.Sugar().Errorf("fail delete App: %v", err)
		return &npool.DeleteAppResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.DeleteAppResponse{
		Info: (info),
	}, nil
}
