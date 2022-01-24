package api

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/appusermgr"

	appusercrud "github.com/NpoolPlatform/appuser-manager/pkg/crud/appuser"
	appusercontrolcrud "github.com/NpoolPlatform/appuser-manager/pkg/crud/appusercontrol"
	appuserextracrud "github.com/NpoolPlatform/appuser-manager/pkg/crud/appuserextra"
	appusersecretcrud "github.com/NpoolPlatform/appuser-manager/pkg/crud/appusersecret"
	banappusercrud "github.com/NpoolPlatform/appuser-manager/pkg/crud/banappuser"

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
	resp, err := appusercrud.Get(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("fail get app user: %v", err)
		return &npool.GetAppUserResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) GetAppUsersByApp(ctx context.Context, in *npool.GetAppUsersByAppRequest) (*npool.GetAppUsersByAppResponse, error) {
	resp, err := appusercrud.GetByApp(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("fail get app user by app: %v", err)
		return &npool.GetAppUsersByAppResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) UpdateAppUser(ctx context.Context, in *npool.UpdateAppUserRequest) (*npool.UpdateAppUserResponse, error) {
	resp, err := appusercrud.Update(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("fail update app user: %v", err)
		return &npool.UpdateAppUserResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) CreateAppUserSecret(ctx context.Context, in *npool.CreateAppUserSecretRequest) (*npool.CreateAppUserSecretResponse, error) {
	resp, err := appusersecretcrud.Create(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("fail create app user secret: %v", err)
		return &npool.CreateAppUserSecretResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) GetAppUserSecret(ctx context.Context, in *npool.GetAppUserSecretRequest) (*npool.GetAppUserSecretResponse, error) {
	resp, err := appusersecretcrud.Get(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("fail get app user secret: %v", err)
		return &npool.GetAppUserSecretResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) GetAppUserSecretByAppUser(ctx context.Context, in *npool.GetAppUserSecretByAppUserRequest) (*npool.GetAppUserSecretByAppUserResponse, error) {
	resp, err := appusersecretcrud.GetByAppUser(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("fail get app user secret by app user: %v", err)
		return &npool.GetAppUserSecretByAppUserResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) UpdateAppUserSecret(ctx context.Context, in *npool.UpdateAppUserSecretRequest) (*npool.UpdateAppUserSecretResponse, error) {
	resp, err := appusersecretcrud.Update(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("fail update app user secret: %v", err)
		return &npool.UpdateAppUserSecretResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) CreateAppUserExtra(ctx context.Context, in *npool.CreateAppUserExtraRequest) (*npool.CreateAppUserExtraResponse, error) {
	resp, err := appuserextracrud.Create(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("fail create app user extra: %v", err)
		return &npool.CreateAppUserExtraResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) GetAppUserExtra(ctx context.Context, in *npool.GetAppUserExtraRequest) (*npool.GetAppUserExtraResponse, error) {
	resp, err := appuserextracrud.Get(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("fail get app user extra: %v", err)
		return &npool.GetAppUserExtraResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) GetAppUserExtraByAppUser(ctx context.Context, in *npool.GetAppUserExtraByAppUserRequest) (*npool.GetAppUserExtraByAppUserResponse, error) {
	resp, err := appuserextracrud.GetByAppUser(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("fail get app user extra by app user: %v", err)
		return &npool.GetAppUserExtraByAppUserResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) UpdateAppUserExtra(ctx context.Context, in *npool.UpdateAppUserExtraRequest) (*npool.UpdateAppUserExtraResponse, error) {
	resp, err := appuserextracrud.Update(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("fail update app user extra: %v", err)
		return &npool.UpdateAppUserExtraResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) CreateBanAppUser(ctx context.Context, in *npool.CreateBanAppUserRequest) (*npool.CreateBanAppUserResponse, error) {
	resp, err := banappusercrud.Create(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("fail create ban app user: %v", err)
		return &npool.CreateBanAppUserResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) GetBanAppUser(ctx context.Context, in *npool.GetBanAppUserRequest) (*npool.GetBanAppUserResponse, error) {
	resp, err := banappusercrud.Get(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("fail get ban app user: %v", err)
		return &npool.GetBanAppUserResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) GetBanAppUserByAppUser(ctx context.Context, in *npool.GetBanAppUserByAppUserRequest) (*npool.GetBanAppUserByAppUserResponse, error) {
	resp, err := banappusercrud.GetByAppUser(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("fail get ban app user by app user: %v", err)
		return &npool.GetBanAppUserByAppUserResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) DeleteBanAppUser(ctx context.Context, in *npool.DeleteBanAppUserRequest) (*npool.DeleteBanAppUserResponse, error) {
	resp, err := banappusercrud.Delete(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("fail delete ban app user: %v", err)
		return &npool.DeleteBanAppUserResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) CreateAppUserControl(ctx context.Context, in *npool.CreateAppUserControlRequest) (*npool.CreateAppUserControlResponse, error) {
	resp, err := appusercontrolcrud.Create(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("fail create app user control: %v", err)
		return &npool.CreateAppUserControlResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) GetAppUserControl(ctx context.Context, in *npool.GetAppUserControlRequest) (*npool.GetAppUserControlResponse, error) {
	resp, err := appusercontrolcrud.Get(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("fail get app user control: %v", err)
		return &npool.GetAppUserControlResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) GetAppUserControlByAppUser(ctx context.Context, in *npool.GetAppUserControlByAppUserRequest) (*npool.GetAppUserControlByAppUserResponse, error) {
	resp, err := appusercontrolcrud.GetByAppUser(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("fail get app user control by app user: %v", err)
		return &npool.GetAppUserControlByAppUserResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) UpdateAppUserControl(ctx context.Context, in *npool.UpdateAppUserControlRequest) (*npool.UpdateAppUserControlResponse, error) {
	resp, err := appusercontrolcrud.Update(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("fail update app user control: %v", err)
		return &npool.UpdateAppUserControlResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
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
