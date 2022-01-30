package api

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/appusermgr"

	appcrud "github.com/NpoolPlatform/appuser-manager/pkg/crud/app"
	appctrlcrud "github.com/NpoolPlatform/appuser-manager/pkg/crud/appcontrol"
	banappcrud "github.com/NpoolPlatform/appuser-manager/pkg/crud/banapp"
	appusermw "github.com/NpoolPlatform/appuser-manager/pkg/middleware/appuser"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateApp(ctx context.Context, in *npool.CreateAppRequest) (*npool.CreateAppResponse, error) {
	resp, err := appcrud.Create(ctx, in, false)
	if err != nil {
		logger.Sugar().Errorw("fail create app: %v", err)
		return &npool.CreateAppResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) UpdateApp(ctx context.Context, in *npool.UpdateAppRequest) (*npool.UpdateAppResponse, error) {
	resp, err := appcrud.Update(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("fail update app: %v", err)
		return &npool.UpdateAppResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) GetApp(ctx context.Context, in *npool.GetAppRequest) (*npool.GetAppResponse, error) {
	resp, err := appcrud.Get(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("fail get app: %v", err)
		return &npool.GetAppResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) GetApps(ctx context.Context, in *npool.GetAppsRequest) (*npool.GetAppsResponse, error) {
	resp, err := appcrud.GetAll(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("fail get all apps: %v", err)
		return &npool.GetAppsResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) GetAppsByCreator(ctx context.Context, in *npool.GetAppsByCreatorRequest) (*npool.GetAppsByCreatorResponse, error) {
	resp, err := appcrud.GetByCreator(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("fail get apps by creator: %v", err)
		return &npool.GetAppsByCreatorResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) CreateAppControl(ctx context.Context, in *npool.CreateAppControlRequest) (*npool.CreateAppControlResponse, error) {
	resp, err := appctrlcrud.Create(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("fail create app control: %v", err)
		return &npool.CreateAppControlResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) UpdateAppControl(ctx context.Context, in *npool.UpdateAppControlRequest) (*npool.UpdateAppControlResponse, error) {
	resp, err := appctrlcrud.Update(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("fail update app control: %v", err)
		return &npool.UpdateAppControlResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) GetAppControl(ctx context.Context, in *npool.GetAppControlRequest) (*npool.GetAppControlResponse, error) {
	resp, err := appctrlcrud.Get(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("fail get app control: %v", err)
		return &npool.GetAppControlResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) GetAppControlByApp(ctx context.Context, in *npool.GetAppControlByAppRequest) (*npool.GetAppControlByAppResponse, error) {
	resp, err := appctrlcrud.GetByApp(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("fail get app control by app: %v", err)
		return &npool.GetAppControlByAppResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) CreateBanApp(ctx context.Context, in *npool.CreateBanAppRequest) (*npool.CreateBanAppResponse, error) {
	resp, err := banappcrud.Create(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("fail create ban app: %v", err)
		return &npool.CreateBanAppResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) GetBanApp(ctx context.Context, in *npool.GetBanAppRequest) (*npool.GetBanAppResponse, error) {
	resp, err := banappcrud.Get(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("fail get ban app: %v", err)
		return &npool.GetBanAppResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) GetBanAppByApp(ctx context.Context, in *npool.GetBanAppByAppRequest) (*npool.GetBanAppByAppResponse, error) {
	resp, err := banappcrud.GetByApp(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("fail get ban app by app: %v", err)
		return &npool.GetBanAppByAppResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) UpdateBanApp(ctx context.Context, in *npool.UpdateBanAppRequest) (*npool.UpdateBanAppResponse, error) {
	resp, err := banappcrud.Update(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("fail update ban app: %v", err)
		return &npool.UpdateBanAppResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) DeleteBanApp(ctx context.Context, in *npool.DeleteBanAppRequest) (*npool.DeleteBanAppResponse, error) {
	resp, err := banappcrud.Delete(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("fail delete ban app: %v", err)
		return &npool.DeleteBanAppResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) GetAppInfo(ctx context.Context, in *npool.GetAppInfoRequest) (*npool.GetAppInfoResponse, error) {
	resp, err := appusermw.GetAppInfo(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("fail get app info: %v", err)
		return &npool.GetAppInfoResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) GetAppInfos(ctx context.Context, in *npool.GetAppInfosRequest) (*npool.GetAppInfosResponse, error) {
	resp, err := appusermw.GetAppInfos(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("fail get app infos: %v", err)
		return &npool.GetAppInfosResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) GetAppInfosByCreator(ctx context.Context, in *npool.GetAppInfosByCreatorRequest) (*npool.GetAppInfosByCreatorResponse, error) {
	resp, err := appusermw.GetAppInfosByCreator(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("fail get app infos by creator: %v", err)
		return &npool.GetAppInfosByCreatorResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}
