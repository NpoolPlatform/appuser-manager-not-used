//go:build !codeanalysis
// +build !codeanalysis

package api

import (
	"context"
	"fmt"

	crud "github.com/NpoolPlatform/appuser-manager/pkg/crud/approlev2"
	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent"
	constant "github.com/NpoolPlatform/appuser-manager/pkg/message/const"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	scodes "go.opentelemetry.io/otel/codes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/appusermgrv2/approle"

	"github.com/google/uuid"
)

func checkAppRoleInfo(info *npool.AppRoleReq) error {
	if _, err := uuid.Parse(info.GetCreatedBy()); err != nil {
		logger.Sugar().Error("CreatedBy is invalid")
		return status.Error(codes.InvalidArgument, "CreatedBy is invalid")
	}

	if info.AppID == nil {
		logger.Sugar().Error("AppID is empty")
		return status.Error(codes.InvalidArgument, "AppID is empty")
	}

	if info.CreatedBy == nil {
		logger.Sugar().Error("CreatedBy is empty")
		return status.Error(codes.InvalidArgument, "CreatedBy is empty")
	}

	if info.Role == nil {
		logger.Sugar().Error("Role is empty")
		return status.Error(codes.InvalidArgument, "Role is empty")
	}

	return nil
}

func appRoleRowToObject(row *ent.AppRole) *npool.AppRole {
	return &npool.AppRole{
		ID:          row.ID.String(),
		AppID:       row.AppID.String(),
		CreatedBy:   row.CreatedBy.String(),
		Role:        row.Role,
		Description: row.Description,
		Default:     row.Default,
	}
}

func (s *AppRoleServer) CreateAppRoleV2(ctx context.Context, in *npool.CreateAppRoleRequest) (*npool.CreateAppRoleResponse, error) {
	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateAppRoleV2")
	defer span.End()
	var err error
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()
	span = crud.AppRoleSpanAttributes(span, in.GetInfo())
	err = checkAppRoleInfo(in.GetInfo())
	if err != nil {
		return &npool.CreateAppRoleResponse{}, err
	}
	span.AddEvent("call crud Create")
	info, err := crud.Create(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorf("fail create app role: %v", err)
		return &npool.CreateAppRoleResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateAppRoleResponse{
		Info: appRoleRowToObject(info),
	}, nil
}

func (s *AppRoleServer) CreateAppRolesV2(ctx context.Context, in *npool.CreateAppRolesRequest) (*npool.CreateAppRolesResponse, error) {
	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateAppRolesV2")
	defer span.End()
	var err error
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()
	if len(in.GetInfos()) == 0 {
		return &npool.CreateAppRolesResponse{},
			status.Error(codes.InvalidArgument,
				"Batah create resource must more than 1",
			)
	}
	for key, info := range in.GetInfos() {
		span.SetAttributes(
			attribute.String("ID"+fmt.Sprintf("%v", key), info.GetID()),
			attribute.String("AppID"+fmt.Sprintf("%v", key), info.GetAppID()),
			attribute.String("Role"+fmt.Sprintf("%v", key), info.GetRole()),
			attribute.String("Description"+fmt.Sprintf("%v", key), info.GetDescription()),
			attribute.String("CreatedBy"+fmt.Sprintf("%v", key), info.GetCreatedBy()),
			attribute.Bool("Default"+fmt.Sprintf("%v", key), info.GetDefault()),
		)
		err = checkAppRoleInfo(info)
		if err != nil {
			return &npool.CreateAppRolesResponse{}, err
		}
	}
	span.AddEvent("call crud CreateBulk")
	rows, err := crud.CreateBulk(ctx, in.GetInfos())
	if err != nil {
		logger.Sugar().Errorf("fail create app roles: %v", err)
		return &npool.CreateAppRolesResponse{}, status.Error(codes.Internal, err.Error())
	}

	infos := make([]*npool.AppRole, 0, len(rows))
	for _, val := range rows {
		infos = append(infos, appRoleRowToObject(val))
	}

	return &npool.CreateAppRolesResponse{
		Infos: infos,
	}, nil
}

func (s *AppRoleServer) UpdateAppRoleV2(ctx context.Context, in *npool.UpdateAppRoleRequest) (*npool.UpdateAppRoleResponse, error) {
	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "UpdateAppRoleV2")
	defer span.End()
	var err error
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()
	span = crud.AppRoleSpanAttributes(span, in.GetInfo())
	if _, err := uuid.Parse(in.GetInfo().GetID()); err != nil {
		logger.Sugar().Errorf("app role id is invalid")
		return &npool.UpdateAppRoleResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	span.AddEvent("call crud Update")
	info, err := crud.Update(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorf("fail update app role: %v", err)
		return &npool.UpdateAppRoleResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.UpdateAppRoleResponse{
		Info: appRoleRowToObject(info),
	}, nil
}

func (s *AppRoleServer) GetAppRoleV2(ctx context.Context, in *npool.GetAppRoleRequest) (*npool.GetAppRoleResponse, error) {
	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetAppRoleV2")
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
		return &npool.GetAppRoleResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	span.AddEvent("call crud Row")
	info, err := crud.Row(ctx, id)
	if err != nil {
		logger.Sugar().Errorf("fail get app role: %v", err)
		return &npool.GetAppRoleResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetAppRoleResponse{
		Info: appRoleRowToObject(info),
	}, nil
}

func (s *AppRoleServer) GetAppRoleOnlyV2(ctx context.Context, in *npool.GetAppRoleOnlyRequest) (*npool.GetAppRoleOnlyResponse, error) {
	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetAppRoleOnlyV2")
	defer span.End()
	var err error
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()
	span = crud.AppRoleCondsSpanAttributes(span, in.GetConds())
	span.AddEvent("call crud RowOnly")
	info, err := crud.RowOnly(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorf("fail get app roles: %v", err)
		return &npool.GetAppRoleOnlyResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetAppRoleOnlyResponse{
		Info: appRoleRowToObject(info),
	}, nil
}

func (s *AppRoleServer) GetAppRolesV2(ctx context.Context, in *npool.GetAppRolesRequest) (*npool.GetAppRolesResponse, error) {
	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetAppRolesV2")
	defer span.End()
	var err error
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()
	span = crud.AppRoleCondsSpanAttributes(span, in.GetConds())
	span.SetAttributes(
		attribute.Int("Limit", int(in.GetLimit())),
		attribute.Int("Offset", int(in.GetOffset())),
	)
	span.AddEvent("call crud Rows")
	rows, total, err := crud.Rows(ctx, in.GetConds(), int(in.GetOffset()), int(in.GetLimit()))
	if err != nil {
		logger.Sugar().Errorf("fail get app roles: %v", err)
		return &npool.GetAppRolesResponse{}, status.Error(codes.Internal, err.Error())
	}

	infos := make([]*npool.AppRole, 0, len(rows))
	for _, val := range rows {
		infos = append(infos, appRoleRowToObject(val))
	}

	return &npool.GetAppRolesResponse{
		Infos: infos,
		Total: uint32(total),
	}, nil
}

func (s *AppRoleServer) ExistAppRoleV2(ctx context.Context, in *npool.ExistAppRoleRequest) (*npool.ExistAppRoleResponse, error) {
	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "ExistAppRoleV2")
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
		return &npool.ExistAppRoleResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	span.AddEvent("call crud Exist")
	exist, err := crud.Exist(ctx, id)
	if err != nil {
		logger.Sugar().Errorf("fail check app role: %v", err)
		return &npool.ExistAppRoleResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistAppRoleResponse{
		Info: exist,
	}, nil
}

func (s *AppRoleServer) ExistAppRoleCondsV2(ctx context.Context, in *npool.ExistAppRoleCondsRequest) (*npool.ExistAppRoleCondsResponse, error) {
	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "ExistAppRoleCondsV2")
	defer span.End()
	var err error
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()
	span = crud.AppRoleCondsSpanAttributes(span, in.GetConds())
	span.AddEvent("call crud ExistConds")
	exist, err := crud.ExistConds(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorf("fail check app role: %v", err)
		return &npool.ExistAppRoleCondsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistAppRoleCondsResponse{
		Info: exist,
	}, nil
}

func (s *AppRoleServer) CountAppRolesV2(ctx context.Context, in *npool.CountAppRolesRequest) (*npool.CountAppRolesResponse, error) {
	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CountAppRolesV2")
	defer span.End()
	var err error
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()
	span = crud.AppRoleCondsSpanAttributes(span, in.GetConds())
	span.AddEvent("call crud Count")
	total, err := crud.Count(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorf("fail count app role : %v", err)
		return &npool.CountAppRolesResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CountAppRolesResponse{
		Info: total,
	}, nil
}

func (s *AppRoleServer) DeleteAppRoleV2(ctx context.Context, in *npool.DeleteAppRoleRequest) (*npool.DeleteAppRoleResponse, error) {
	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "DeleteAppRoleV2")
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
		return &npool.DeleteAppRoleResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	span.AddEvent("call crud Delete")
	info, err := crud.Delete(ctx, id)
	if err != nil {
		logger.Sugar().Errorf("fail delete app role: %v", err)
		return &npool.DeleteAppRoleResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.DeleteAppRoleResponse{
		Info: appRoleRowToObject(info),
	}, nil
}
