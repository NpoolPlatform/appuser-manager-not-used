package api

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/appusermgr"

	appcrud "github.com/NpoolPlatform/appuser-manager/pkg/crud/app"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func CreateApp(ctx context.Context, in *npool.CreateAppRequest) (*npool.CreateAppResponse, error) {
	resp, err := appcrud.Create(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("fail create app: %v", err)
		return &npool.CreateAppResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func UpdateApp(ctx context.Context, in *npool.UpdateAppRequest) (*npool.UpdateAppResponse, error) {
	resp, err := appcrud.Update(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("fail update app: %v", err)
		return &npool.UpdateAppResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func GetApp(ctx context.Context, in *npool.GetAppRequest) (*npool.GetAppResponse, error) {
	resp, err := appcrud.Get(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("fail get app: %v", err)
		return &npool.GetAppResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func GetApps(ctx context.Context, in *npool.GetAppsRequest) (*npool.GetAppsResponse, error) {
	resp, err := appcrud.GetAll(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("fail get all apps: %v", err)
		return &npool.GetAppsResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func GetAppsByCreator(ctx context.Context, in *npool.GetAppsByCreatorRequest) (*npool.GetAppsByCreatorResponse, error) {
	resp, err := appcrud.GetByCreator(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("fail get apps by creator: %v", err)
		return &npool.GetAppsByCreatorResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func CreateAppControl(ctx context.Context, in *npool.CreateAppControlRequest) (*npool.CreateAppControlResponse, error) {
	return nil, nil
}

func UpdateAppControl(ctx context.Context, in *npool.UpdateAppControlRequest) (*npool.UpdateAppControlResponse, error) {
	return nil, nil
}

func CreateBanApp(ctx context.Context, in *npool.CreateBanAppRequest) (*npool.CreateBanAppResponse, error) {
	return nil, nil
}

func DeleteBanApp(ctx context.Context, in *npool.DeleteBanAppRequest) (*npool.DeleteBanAppResponse, error) {
	return nil, nil
}

func GetAppInfo(ctx context.Context, in *npool.GetAppInfoRequest) (*npool.GetAppInfoResponse, error) {
	return nil, nil
}

func GetAppInfos(ctx context.Context, in *npool.GetAppInfosRequest) (*npool.GetAppInfosResponse, error) {
	return nil, nil
}

func GetAppInfosByCreator(ctx context.Context, in *npool.GetAppInfosByCreatorRequest) (*npool.GetAppInfosByCreatorResponse, error) {
	return nil, nil
}
