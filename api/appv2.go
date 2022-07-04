//go:build !codeanalysis
// +build !codeanalysis

package api

import (
	"context"
	"fmt"

	crud "github.com/NpoolPlatform/appuser-manager/pkg/crud/appv2"
	entapp "github.com/NpoolPlatform/appuser-manager/pkg/db/ent/app"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/appusermgrv2/app"

	"github.com/google/uuid"
)

func checkInfo(info *npool.App) error {
	if _, err := uuid.Parse(info.GetCreatedBy()); err != nil {
		logger.Sugar().Error("create by is invalid")
		return fmt.Errorf("create by is invalid")
	}
	if info.Name == nil {
		logger.Sugar().Error("name is empty")
		return fmt.Errorf("name empty")
	}
	if info.GetLogo() == "" {
		logger.Sugar().Error("logo is empty")
		return fmt.Errorf("logo empty")
	}
	return nil
}

func (s *AppService) CreateAppV2(ctx context.Context, in *npool.CreateAppRequest) (*npool.CreateAppResponse, error) {
	err := checkInfo(in.GetInfo())
	if err != nil {
		return &npool.CreateAppResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	schema, err := crud.New(ctx, nil)
	if err != nil {
		logger.Sugar().Errorf("fail create schema entity: %v", err)
		return &npool.CreateAppResponse{}, status.Error(codes.Internal, err.Error())
	}

	info, err := schema.Create(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorf("fail create app: %v", err)
		return &npool.CreateAppResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateAppResponse{
		Info: info,
	}, nil
}

func (s *Server) CreateAppsV2(ctx context.Context, in *npool.CreateAppsRequest) (*npool.CreateAppsResponse, error) {
	for _, info := range in.GetInfos() {
		err := checkInfo(info)
		if err != nil {
			return &npool.CreateAppsResponse{}, status.Error(codes.InvalidArgument, err.Error())
		}
	}

	schema, err := crud.New(ctx, nil)
	if err != nil {
		logger.Sugar().Errorf("fail create schema entity: %v", err)
		return &npool.CreateAppsResponse{}, status.Error(codes.Internal, err.Error())
	}

	infos, err := schema.CreateBulk(ctx, in.GetInfos())
	if err != nil {
		logger.Sugar().Errorf("fail create Apps: %v", err)
		return &npool.CreateAppsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateAppsResponse{
		Infos: infos,
	}, nil
}

func (s *Server) UpdateAppV2(ctx context.Context, in *npool.UpdateAppRequest) (*npool.UpdateAppResponse, error) {
	if _, err := uuid.Parse(in.GetInfo().GetID()); err != nil {
		logger.Sugar().Errorf("app id is invalid")
		return &npool.UpdateAppResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	schema, err := crud.New(ctx, nil)
	if err != nil {
		logger.Sugar().Errorf("fail create schema entity: %v", err)
		return &npool.UpdateAppResponse{}, status.Error(codes.Internal, err.Error())
	}

	info, err := schema.Update(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorf("fail update app: %v", err)
		return &npool.UpdateAppResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.UpdateAppResponse{
		Info: info,
	}, nil
}

func AppFieldsToFields(fields cruder.FilterFields) (cruder.Fields, error) {
	newFields := cruder.NewFields()
	for k, v := range fields {
		switch k {
		case entapp.FieldID:
			fallthrough //nolint
		case entapp.FieldCreatedBy:
			fallthrough
		case entapp.FieldName:
			fallthrough //nolint
		case entapp.FieldDescription:
			fallthrough //nolint
		case entapp.FieldLogo:
			newFields.WithField(k, v.GetNullValue())
			newFields.WithField(k, v.GetNumberValue())
			newFields.WithField(k, v.GetBoolValue())
		default:
			return nil, fmt.Errorf("invalid App field")
		}
	}
	return newFields, nil
}

func (s *AppService) GetAppV2(ctx context.Context, in *npool.GetAppRequest) (*npool.GetAppResponse, error) {
	id, err := uuid.Parse(in.GetID())
	if err != nil {
		return &npool.GetAppResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	schema, err := crud.New(ctx, nil)
	if err != nil {
		logger.Sugar().Errorf("fail create schema entity: %v", err)
		return &npool.GetAppResponse{}, status.Error(codes.Internal, err.Error())
	}

	info, err := schema.Row(ctx, id)
	if err != nil {
		logger.Sugar().Errorf("fail get App: %v", err)
		return &npool.GetAppResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetAppResponse{
		Info: info,
	}, nil
}

func (s *Server) GetAppOnlyV2(ctx context.Context, in *npool.GetAppOnlyRequest) (*npool.GetAppOnlyResponse, error) {
	schema, err := crud.New(ctx, nil)
	if err != nil {
		logger.Sugar().Errorf("fail create schema entity: %v", err)
		return &npool.GetAppOnlyResponse{}, status.Error(codes.Internal, err.Error())
	}

	info, err := schema.RowOnly(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorf("fail get Apps: %v", err)
		return &npool.GetAppOnlyResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetAppOnlyResponse{
		Info: info,
	}, nil
}

func (s *AppService) GetAppsV2(ctx context.Context, in *npool.GetAppsRequest) (*npool.GetAppsResponse, error) {
	schema, err := crud.New(ctx, nil)
	if err != nil {
		logger.Sugar().Errorf("fail create schema entity: %v", err)
		return &npool.GetAppsResponse{}, status.Error(codes.Internal, err.Error())
	}

	infos, total, err := schema.Rows(ctx, in.GetConds(), int(in.GetOffset()), int(in.GetLimit()))
	if err != nil {
		logger.Sugar().Errorf("fail get Apps: %v", err)
		return &npool.GetAppsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetAppsResponse{
		Infos: infos,
		Total: uint32(total),
	}, nil
}

func (s *Server) ExistAppV2(ctx context.Context, in *npool.ExistAppRequest) (*npool.ExistAppResponse, error) {
	id, err := uuid.Parse(in.GetID())
	if err != nil {
		return &npool.ExistAppResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	schema, err := crud.New(ctx, nil)
	if err != nil {
		logger.Sugar().Errorf("fail create schema entity: %v", err)
		return &npool.ExistAppResponse{}, status.Error(codes.Internal, err.Error())
	}

	exist, err := schema.Exist(ctx, id)
	if err != nil {
		logger.Sugar().Errorf("fail check App: %v", err)
		return &npool.ExistAppResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistAppResponse{
		Result: exist,
	}, nil
}

func (s *Server) ExistAppCondsV2(ctx context.Context, in *npool.ExistAppCondsRequest) (*npool.ExistAppCondsResponse, error) {
	schema, err := crud.New(ctx, nil)
	if err != nil {
		logger.Sugar().Errorf("fail create schema entity: %v", err)
		return &npool.ExistAppCondsResponse{}, status.Error(codes.Internal, err.Error())
	}

	exist, err := schema.ExistConds(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorf("fail check App: %v", err)
		return &npool.ExistAppCondsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistAppCondsResponse{
		Result: exist,
	}, nil
}

func (s *Server) CountAppsV2(ctx context.Context, in *npool.CountAppsRequest) (*npool.CountAppsResponse, error) {
	schema, err := crud.New(ctx, nil)
	if err != nil {
		logger.Sugar().Errorf("fail create schema entity: %v", err)
		return &npool.CountAppsResponse{}, status.Error(codes.Internal, err.Error())
	}

	total, err := schema.Count(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorf("fail count Apps: %v", err)
		return &npool.CountAppsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CountAppsResponse{
		Result: total,
	}, nil
}

func (s *Server) DeleteAppV2(ctx context.Context, in *npool.DeleteAppRequest) (*npool.DeleteAppResponse, error) {
	id, err := uuid.Parse(in.GetID())
	if err != nil {
		return &npool.DeleteAppResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	schema, err := crud.New(ctx, nil)
	if err != nil {
		logger.Sugar().Errorf("fail create schema entity: %v", err)
		return &npool.DeleteAppResponse{}, status.Error(codes.Internal, err.Error())
	}

	info, err := schema.Delete(ctx, id)
	if err != nil {
		logger.Sugar().Errorf("fail delete App: %v", err)
		return &npool.DeleteAppResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.DeleteAppResponse{
		Info: info,
	}, nil
}
