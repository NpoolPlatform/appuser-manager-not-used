//go:build !codeanalysis
// +build !codeanalysis

package api

import (
	"context"
	"fmt"

	crud "github.com/NpoolPlatform/appuser-manager/pkg/crud/appv2"
	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent"
	constant "github.com/NpoolPlatform/appuser-manager/pkg/message/const"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	scodes "go.opentelemetry.io/otel/codes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/appusermgrv2/app"

	"github.com/google/uuid"
)

func checkAppInfo(info *npool.AppReq) error {
	if _, err := uuid.Parse(info.GetCreatedBy()); err != nil {
		logger.Sugar().Error("CreatedBy is invalid")
		return status.Error(codes.InvalidArgument, "CreatedBy is invalid")
	}

	if info.Name == nil {
		logger.Sugar().Error("Name is empty")
		return status.Error(codes.InvalidArgument, "Name is empty")
	}

	if info.GetLogo() == "" {
		logger.Sugar().Error("Logo is empty")
		return status.Error(codes.InvalidArgument, "Logo is empty")
	}

	return nil
}

func appRowToObject(row *ent.App) *npool.App {
	return &npool.App{
		ID:          row.ID.String(),
		CreatedBy:   row.CreatedBy.String(),
		Name:        row.Name,
		Logo:        row.Logo,
		Description: row.Description,
		CreatedAt:   row.CreatedAt,
	}
}

func (s *AppServer) CreateAppV2(ctx context.Context, in *npool.CreateAppRequest) (*npool.CreateAppResponse, error) {
	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateAppV2")
	defer span.End()
	var err error
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()
	span = crud.AppSpanAttributes(span, in.GetInfo())
	err = checkAppInfo(in.GetInfo())
	if err != nil {
		return &npool.CreateAppResponse{}, err
	}
	span.AddEvent("call crud Create")
	info, err := crud.Create(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorf("fail create app: %v", err)
		return &npool.CreateAppResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateAppResponse{
		Info: appRowToObject(info),
	}, nil
}

func (s *AppServer) CreateAppsV2(ctx context.Context, in *npool.CreateAppsRequest) (*npool.CreateAppsResponse, error) {
	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateAppsV2")
	defer span.End()
	var err error
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()
	if len(in.GetInfos()) == 0 {
		return &npool.CreateAppsResponse{},
			status.Error(codes.InvalidArgument,
				"Batah create resource must more than 1",
			)
	}

	dup := make(map[string]struct{})
	for key, info := range in.GetInfos() {
		span.SetAttributes(
			attribute.String("Description"+fmt.Sprintf("%v", key), info.GetDescription()),
			attribute.String("ID"+fmt.Sprintf("%v", key), info.GetID()),
			attribute.String("CreatedBy"+fmt.Sprintf("%v", key), info.GetID()),
			attribute.String("Name"+fmt.Sprintf("%v", key), info.GetCreatedBy()),
			attribute.String("Logo"+fmt.Sprintf("%v", key), info.GetName()),
			attribute.Int("CreatedAt"+fmt.Sprintf("%v", key), int(info.GetCreatedAt())),
		)
		err := checkAppInfo(info)
		if err != nil {
			return &npool.CreateAppsResponse{}, err
		}

		if _, ok := dup[info.GetName()]; ok {
			return &npool.CreateAppsResponse{},
				status.Errorf(codes.AlreadyExists,
					"Name: %v duplicate create",
					info.GetName(),
				)
		}

		dup[info.GetName()] = struct{}{}
	}

	rows, err := crud.CreateBulk(ctx, in.GetInfos())
	if err != nil {
		logger.Sugar().Errorf("fail create Apps: %v", err)
		return &npool.CreateAppsResponse{}, status.Error(codes.Internal, err.Error())
	}

	infos := make([]*npool.App, 0, len(rows))
	for _, val := range rows {
		infos = append(infos, appRowToObject(val))
	}

	return &npool.CreateAppsResponse{
		Infos: infos,
	}, nil
}

func (s *AppServer) UpdateAppV2(ctx context.Context, in *npool.UpdateAppRequest) (*npool.UpdateAppResponse, error) {
	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "UpdateAppV2")
	defer span.End()
	var err error
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
		Info: appRowToObject(info),
	}, nil
}

func (s *AppServer) GetAppV2(ctx context.Context, in *npool.GetAppRequest) (*npool.GetAppResponse, error) {
	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetAppV2")
	defer span.End()
	var err error
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
		Info: appRowToObject(info),
	}, nil
}

func (s *AppServer) GetAppOnlyV2(ctx context.Context, in *npool.GetAppOnlyRequest) (*npool.GetAppOnlyResponse, error) {
	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetAppOnlyV2")
	defer span.End()
	var err error
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
		Info: appRowToObject(info),
	}, nil
}

func (s *AppServer) GetAppsV2(ctx context.Context, in *npool.GetAppsRequest) (*npool.GetAppsResponse, error) {
	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetAppsV2")
	defer span.End()
	var err error
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
		infos = append(infos, appRowToObject(val))
	}

	return &npool.GetAppsResponse{
		Infos: infos,
		Total: uint32(total),
	}, nil
}

func (s *AppServer) ExistAppV2(ctx context.Context, in *npool.ExistAppRequest) (*npool.ExistAppResponse, error) {
	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "ExistAppV2")
	defer span.End()
	var err error
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

func (s *AppServer) ExistAppCondsV2(ctx context.Context, in *npool.ExistAppCondsRequest) (*npool.ExistAppCondsResponse, error) {
	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "ExistAppCondsV2")
	defer span.End()
	var err error
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

func (s *AppServer) CountAppsV2(ctx context.Context, in *npool.CountAppsRequest) (*npool.CountAppsResponse, error) {
	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CountAppsV2")
	defer span.End()
	var err error
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

func (s *AppServer) DeleteAppV2(ctx context.Context, in *npool.DeleteAppRequest) (*npool.DeleteAppResponse, error) {
	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "DeleteAppV2")
	defer span.End()
	var err error
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
		Info: appRowToObject(info),
	}, nil
}
