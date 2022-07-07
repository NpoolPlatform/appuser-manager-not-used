//go:build !codeanalysis
// +build !codeanalysis

package api

import (
	"context"

	crud "github.com/NpoolPlatform/appuser-manager/pkg/crud/appv2"
	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/appusermgrv2/app"

	"github.com/google/uuid"
)

func checkAppInfo(info *npool.App) error {
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

func appRowToObject(row *ent.App) *npool.AppRes {
	return &npool.AppRes{
		ID:          row.ID.String(),
		CreatedBy:   row.CreatedBy.String(),
		Name:        row.Name,
		Logo:        row.Logo,
		Description: row.Description,
		CreatedAt:   row.CreatedAt,
	}
}

func (s *AppServer) CreateAppV2(ctx context.Context, in *npool.CreateAppRequest) (*npool.CreateAppResponse, error) {
	err := checkAppInfo(in.GetInfo())
	if err != nil {
		return &npool.CreateAppResponse{}, err
	}

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
	if len(in.GetInfos()) == 0 {
		return &npool.CreateAppsResponse{},
			status.Error(codes.InvalidArgument,
				"Batah create resource must more than 1",
			)
	}

	dup := make(map[string]struct{})
	for _, info := range in.GetInfos() {
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

	infos := make([]*npool.AppRes, 0, len(rows))
	for _, val := range rows {
		infos = append(infos, appRowToObject(val))
	}

	return &npool.CreateAppsResponse{
		Infos: infos,
	}, nil
}

func (s *AppServer) UpdateAppV2(ctx context.Context, in *npool.UpdateAppRequest) (*npool.UpdateAppResponse, error) {
	if _, err := uuid.Parse(in.GetInfo().GetID()); err != nil {
		logger.Sugar().Errorf("app id is invalid")
		return &npool.UpdateAppResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

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
	id, err := uuid.Parse(in.GetID())
	if err != nil {
		return &npool.GetAppResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

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
	rows, total, err := crud.Rows(ctx, in.GetConds(), int(in.GetOffset()), int(in.GetLimit()))
	if err != nil {
		logger.Sugar().Errorf("fail get Apps: %v", err)
		return &npool.GetAppsResponse{}, status.Error(codes.Internal, err.Error())
	}

	infos := make([]*npool.AppRes, 0, len(rows))
	for _, val := range rows {
		infos = append(infos, appRowToObject(val))
	}

	return &npool.GetAppsResponse{
		Infos: infos,
		Total: uint32(total),
	}, nil
}

func (s *AppServer) ExistAppV2(ctx context.Context, in *npool.ExistAppRequest) (*npool.ExistAppResponse, error) {
	id, err := uuid.Parse(in.GetID())
	if err != nil {
		return &npool.ExistAppResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

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
	id, err := uuid.Parse(in.GetID())
	if err != nil {
		return &npool.DeleteAppResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := crud.Delete(ctx, id)
	if err != nil {
		logger.Sugar().Errorf("fail delete App: %v", err)
		return &npool.DeleteAppResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.DeleteAppResponse{
		Info: appRowToObject(info),
	}, nil
}
