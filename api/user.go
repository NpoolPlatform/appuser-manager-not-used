package api

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/appusermgr"

	appusercrud "github.com/NpoolPlatform/appuser-manager/pkg/crud/appuser"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateAppUser(ctx context.Context, in *npool.CreateAppUserRequest) (*npool.CreateAppUserResponse, error) {
	resp, err := appusercrud.Create(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("fail create app user: %v", err)
		return &npool.CreateAppUserResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) GetAppUser(ctx context.Context, in *npool.GetAppUserRequest) (*npool.GetAppUserResponse, error) {
	return nil, nil
}

func (s *Server) GetAppUsersByApp(ctx context.Context, in *npool.GetAppUsersByAppRequest) (*npool.GetAppUsersByAppResponse, error) {
	return nil, nil
}

func (s *Server) UpdateAppUser(ctx context.Context, in *npool.UpdateAppUserRequest) (*npool.UpdateAppUserResponse, error) {
	return nil, nil
}

func (s *Server) CreateAppUserSecret(ctx context.Context, in *npool.CreateAppUserSecretRequest) (*npool.CreateAppUserSecretResponse, error) {
	return nil, nil
}

func (s *Server) UpdateAppUserSecret(ctx context.Context, in *npool.UpdateAppUserSecretRequest) (*npool.UpdateAppUserSecretResponse, error) {
	return nil, nil
}

func (s *Server) CreateAppUserExtra(ctx context.Context, in *npool.CreateAppUserExtraRequest) (*npool.CreateAppUserExtraResponse, error) {
	return nil, nil
}

func (s *Server) UpdateAppUserExtra(ctx context.Context, in *npool.UpdateAppUserExtraRequest) (*npool.UpdateAppUserExtraResponse, error) {
	return nil, nil
}

func (s *Server) CreateBanAppUser(ctx context.Context, in *npool.CreateBanAppUserRequest) (*npool.CreateBanAppUserResponse, error) {
	return nil, nil
}

func (s *Server) DeleteBanAppUser(ctx context.Context, in *npool.DeleteBanAppUserRequest) (*npool.DeleteBanAppUserResponse, error) {
	return nil, nil
}

func (s *Server) CreateAppUserControl(ctx context.Context, in *npool.CreateAppUserControlRequest) (*npool.CreateAppUserControlResponse, error) {
	return nil, nil
}

func (s *Server) UpdateAppUserControl(ctx context.Context, in *npool.UpdateAppUserControlRequest) (*npool.UpdateAppUserControlResponse, error) {
	return nil, nil
}

func (s *Server) GetAppUserInfo(ctx context.Context, in *npool.GetAppUserInfoRequest) (*npool.GetAppUserInfoResponse, error) {
	return nil, nil
}

func (s *Server) GetAppUserInfos(ctx context.Context, in *npool.GetAppUserInfosRequest) (*npool.GetAppUserInfosResponse, error) {
	return nil, nil
}

func (s *Server) GetAppUserInfosByApp(ctx context.Context, in *npool.GetAppUserInfosByAppRequest) (*npool.GetAppUserInfosByAppResponse, error) {
	return nil, nil
}
