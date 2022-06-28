//go:build !codeanalysis
// +build !codeanalysis

package api

import (
	"context"
	"fmt"
	constant "github.com/NpoolPlatform/appuser-manager/pkg/const"

	crud "github.com/NpoolPlatform/appuser-manager/pkg/crud/appv2"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	"github.com/NpoolPlatform/message/npool/appusermgrv2/app"

	"github.com/google/uuid"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func checkInfo(info *app.App) error {
	if _, err := uuid.Parse(info.GetCreatedBy()); err != nil {
		logger.Sugar().Error("create by is invalid")
		return fmt.Errorf("create by is invalid")
	}
	if info.GetName() == "" {
		logger.Sugar().Error("name is empty")
		return fmt.Errorf("name empty")
	}
	if info.GetLogo() == "" {
		logger.Sugar().Error("logo is empty")
		return fmt.Errorf("logo empty")
	}
	return nil
}

func (s *AppService) CreateAppV2(ctx context.Context, in *app.CreateAppRequest) (*app.CreateAppResponse, error) {
	err := checkInfo(in.GetInfo())
	if err != nil {
		return &app.CreateAppResponse{}, status.Error(codes.Internal, err.Error())
	}

	schema, err := crud.New(ctx, nil)
	if err != nil {
		logger.Sugar().Errorf("fail create schema entity: %v", err)
		return &app.CreateAppResponse{}, status.Error(codes.Internal, err.Error())
	}

	info, err := schema.Create(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorf("fail create app: %v", err)
		return &app.CreateAppResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &app.CreateAppResponse{
		Info: info,
	}, nil
}

func (s *AppService) CreateAppsV2(ctx context.Context, in *app.CreateAppsRequest) (*app.CreateAppsResponse, error) {
	for _, info := range in.GetInfos() {
		err := checkInfo(info)
		if err != nil {
			return &app.CreateAppsResponse{}, status.Error(codes.Internal, err.Error())
		}
	}

	schema, err := crud.New(ctx, nil)
	if err != nil {
		logger.Sugar().Errorf("fail create schema entity: %v", err)
		return &app.CreateAppsResponse{}, status.Error(codes.Internal, err.Error())
	}

	infos, err := schema.CreateBulk(ctx, in.GetInfos())
	if err != nil {
		logger.Sugar().Errorf("fail create Apps: %v", err)
		return &app.CreateAppsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &app.CreateAppsResponse{
		Infos: infos,
	}, nil
}

func (s *AppService) UpdateAppV2(ctx context.Context, in *app.UpdateAppRequest) (*app.UpdateAppResponse, error) {
	if _, err := uuid.Parse(in.GetInfo().GetID()); err != nil {
		logger.Sugar().Errorf("app id is invalid")
		return &app.UpdateAppResponse{}, status.Error(codes.Internal, err.Error())
	}

	schema, err := crud.New(ctx, nil)
	if err != nil {
		logger.Sugar().Errorf("fail create schema entity: %v", err)
		return &app.UpdateAppResponse{}, status.Error(codes.Internal, err.Error())
	}

	info, err := schema.Update(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorf("fail update app: %v", err)
		return &app.UpdateAppResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &app.UpdateAppResponse{
		Info: info,
	}, nil
}

func AppFieldsToFields(fields cruder.FilterFields) (cruder.Fields, error) {
	newFields := cruder.NewFields()

	for k, v := range fields {
		switch k {
		case constant.FieldID:
			fallthrough //nolint
		case constant.AppFieldCreatedBy:
			newFields.WithField(k, v.GetStringValue())
		case constant.AppFieldName:
			fallthrough //nolint
		case constant.AppFieldDescription:
			fallthrough //nolint
		default:
			return nil, fmt.Errorf("invalid App field")
		}
	}
	return newFields, nil
}

func (s *AppService) UpdateAppFieldsV2(ctx context.Context, in *app.UpdateAppFieldsRequest) (*app.UpdateAppFieldsResponse, error) {
	id, err := uuid.Parse(in.GetID())
	if err != nil {
		logger.Sugar().Errorf("invalid App id: %v", err)
		return &app.UpdateAppFieldsResponse{}, status.Error(codes.Internal, err.Error())
	}

	fields, err := AppFieldsToFields(in.GetFields())
	if err != nil {
		logger.Sugar().Errorf("invalid App fields: %v", err)
		return &app.UpdateAppFieldsResponse{}, status.Error(codes.Internal, err.Error())
	}

	if len(fields) == 0 {
		logger.Sugar().Errorf("empty App fields: %v", err)
		return &app.UpdateAppFieldsResponse{}, status.Error(codes.Internal, "empty App fields")
	}

	schema, err := crud.New(ctx, nil)
	if err != nil {
		logger.Sugar().Errorf("fail create schema entity: %v", err)
		return &app.UpdateAppFieldsResponse{}, status.Error(codes.Internal, err.Error())
	}

	info, err := schema.UpdateFields(ctx, id, fields)
	if err != nil {
		logger.Sugar().Errorf("fail update App: %v", err)
		return &app.UpdateAppFieldsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &app.UpdateAppFieldsResponse{
		Info: info,
	}, nil
}

func appCondsToConds(conds cruder.FilterConds) (cruder.Conds, error) {
	newConds := cruder.NewConds()

	for k, v := range conds {
		switch v.Op {
		case cruder.EQ:
		case cruder.GT:
		case cruder.LT:
		case cruder.LIKE:
		case cruder.IN:
		default:
			return nil, fmt.Errorf("invalid filter condition op")
		}
		switch k {
		case constant.FieldID:
			fallthrough //nolint
		case constant.AppFieldCreatedBy:
			fallthrough
		case constant.AppFieldName:
			fallthrough //nolint
		case constant.AppFieldLogo:
			fallthrough //nolint
		case constant.AppFieldDescription:
			fallthrough //nolint
		default:
			return nil, fmt.Errorf("invalid App field")
		}
	}

	return newConds, nil
}

func (s *AppService) GetAppV2(ctx context.Context, in *app.GetAppRequest) (*app.GetAppResponse, error) {
	id, err := uuid.Parse(in.GetID())
	if err != nil {
		return &app.GetAppResponse{}, fmt.Errorf("invalid App id: %v", err)
	}

	schema, err := crud.New(ctx, nil)
	if err != nil {
		logger.Sugar().Errorf("fail create schema entity: %v", err)
		return &app.GetAppResponse{}, status.Error(codes.Internal, err.Error())
	}

	info, err := schema.Row(ctx, id)
	if err != nil {
		logger.Sugar().Errorf("fail get App: %v", err)
		return &app.GetAppResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &app.GetAppResponse{
		Info: info,
	}, nil
}

func (s *AppService) GetAppOnlyV2(ctx context.Context, in *app.GetAppOnlyRequest) (*app.GetAppOnlyResponse, error) {
	conds, err := appCondsToConds(in.GetConds())
	if err != nil {
		logger.Sugar().Errorf("invalid App fields: %v", err)
		return &app.GetAppOnlyResponse{}, status.Error(codes.Internal, err.Error())
	}

	schema, err := crud.New(ctx, nil)
	if err != nil {
		logger.Sugar().Errorf("fail create schema entity: %v", err)
		return &app.GetAppOnlyResponse{}, status.Error(codes.Internal, err.Error())
	}

	info, err := schema.RowOnly(ctx, conds)
	if err != nil {
		logger.Sugar().Errorf("fail get Apps: %v", err)
		return &app.GetAppOnlyResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &app.GetAppOnlyResponse{
		Info: info,
	}, nil
}

func (s *AppService) GetAppsV2(ctx context.Context, in *app.GetAppsRequest) (*app.GetAppsResponse, error) {
	conds, err := appCondsToConds(in.GetConds())
	if err != nil {
		logger.Sugar().Errorf("invalid stock fields: %v", err)
		return &app.GetAppsResponse{}, status.Error(codes.Internal, err.Error())
	}
	schema, err := crud.New(ctx, nil)
	if err != nil {
		logger.Sugar().Errorf("fail create schema entity: %v", err)
		return &app.GetAppsResponse{}, status.Error(codes.Internal, err.Error())
	}

	infos, total, err := schema.Rows(ctx, conds, int(in.GetOffset()), int(in.GetLimit()))
	if err != nil {
		logger.Sugar().Errorf("fail get Apps: %v", err)
		return &app.GetAppsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &app.GetAppsResponse{
		Infos: infos,
		Total: uint32(total),
	}, nil
}

func (s *AppService) ExistAppV2(ctx context.Context, in *app.ExistAppRequest) (*app.ExistAppResponse, error) {
	id, err := uuid.Parse(in.GetID())
	if err != nil {
		return &app.ExistAppResponse{}, fmt.Errorf("invalid App id: %v", err)
	}

	schema, err := crud.New(ctx, nil)
	if err != nil {
		logger.Sugar().Errorf("fail create schema entity: %v", err)
		return &app.ExistAppResponse{}, status.Error(codes.Internal, err.Error())
	}

	exist, err := schema.Exist(ctx, id)
	if err != nil {
		logger.Sugar().Errorf("fail check App: %v", err)
		return &app.ExistAppResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &app.ExistAppResponse{
		Result: exist,
	}, nil
}

func (s *AppService) ExistAppCondsV2(ctx context.Context, in *app.ExistAppCondsRequest) (*app.ExistAppCondsResponse, error) {
	conds, err := appCondsToConds(in.GetConds())
	if err != nil {
		logger.Sugar().Errorf("invalid App fields: %v", err)
		return &app.ExistAppCondsResponse{}, status.Error(codes.Internal, err.Error())
	}

	if len(conds) == 0 {
		logger.Sugar().Errorf("empty App fields: %v", err)
		return &app.ExistAppCondsResponse{}, status.Error(codes.Internal, "empty App fields")
	}

	schema, err := crud.New(ctx, nil)
	if err != nil {
		logger.Sugar().Errorf("fail create schema entity: %v", err)
		return &app.ExistAppCondsResponse{}, status.Error(codes.Internal, err.Error())
	}

	exist, err := schema.ExistConds(ctx, conds)
	if err != nil {
		logger.Sugar().Errorf("fail check App: %v", err)
		return &app.ExistAppCondsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &app.ExistAppCondsResponse{
		Result: exist,
	}, nil
}

func (s *AppService) CountAppsV2(ctx context.Context, in *app.CountAppsRequest) (*app.CountAppsResponse, error) {
	conds, err := appCondsToConds(in.GetConds())
	if err != nil {
		logger.Sugar().Errorf("invalid App fields: %v", err)
		return &app.CountAppsResponse{}, status.Error(codes.Internal, err.Error())
	}

	if len(conds) == 0 {
		logger.Sugar().Errorf("empty App fields: %v", err)
		return &app.CountAppsResponse{}, status.Error(codes.Internal, "empty App fields")
	}

	schema, err := crud.New(ctx, nil)
	if err != nil {
		logger.Sugar().Errorf("fail create schema entity: %v", err)
		return &app.CountAppsResponse{}, status.Error(codes.Internal, err.Error())
	}

	total, err := schema.Count(ctx, conds)
	if err != nil {
		logger.Sugar().Errorf("fail count Apps: %v", err)
		return &app.CountAppsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &app.CountAppsResponse{
		Result: total,
	}, nil
}

func (s *AppService) DeleteAppV2(ctx context.Context, in *app.DeleteAppRequest) (*app.DeleteAppResponse, error) {
	id, err := uuid.Parse(in.GetID())
	if err != nil {
		return &app.DeleteAppResponse{}, fmt.Errorf("invalid App id: %v", err)
	}

	schema, err := crud.New(ctx, nil)
	if err != nil {
		logger.Sugar().Errorf("fail create schema entity: %v", err)
		return &app.DeleteAppResponse{}, status.Error(codes.Internal, err.Error())
	}

	info, err := schema.Delete(ctx, id)
	if err != nil {
		logger.Sugar().Errorf("fail delete App: %v", err)
		return &app.DeleteAppResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &app.DeleteAppResponse{
		Info: info,
	}, nil
}
