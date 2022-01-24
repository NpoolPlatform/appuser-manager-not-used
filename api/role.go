package api

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	approlecrud "github.com/NpoolPlatform/appuser-manager/pkg/crud/approle"
	npool "github.com/NpoolPlatform/message/npool/appusermgr"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateAppRole(ctx context.Context, in *npool.CreateAppRoleRequest) (*npool.CreateAppRoleResponse, error) {
	resp, err := approlecrud.Create(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("fail create app role: %v", err)
		return &npool.CreateAppRoleResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) GetAppRole(ctx context.Context, in *npool.GetAppRoleRequest) (*npool.GetAppRoleResponse, error) {
	resp, err := approlecrud.Get(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("fail get app role: %v", err)
		return &npool.GetAppRoleResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) GetAppRolesByApp(ctx context.Context, in *npool.GetAppRolesByAppRequest) (*npool.GetAppRolesByAppResponse, error) {
	resp, err := approlecrud.GetByApp(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("fail get app roles by app: %v", err)
		return &npool.GetAppRolesByAppResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) UpdateAppRole(ctx context.Context, in *npool.UpdateAppRoleRequest) (*npool.UpdateAppRoleResponse, error) {
	resp, err := approlecrud.Update(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("fail update app role: %v", err)
		return &npool.UpdateAppRoleResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) CreateAppRoleUser(ctx context.Context, in *npool.CreateAppRoleUserRequest) (*npool.CreateAppRoleUserResponse, error) {
	return nil, nil
}

func (s *Server) GetAppRoleUser(ctx context.Context, in *npool.GetAppRoleUserRequest) (*npool.GetAppRoleUserResponse, error) {
	return nil, nil
}

func (s *Server) GetAppRoleUsersByAppRole(ctx context.Context, in *npool.GetAppRoleUsersByAppRoleRequest) (*npool.GetAppRoleUsersByAppRoleResponse, error) {
	return nil, nil
}

func (s *Server) GetUserRolesByAppUser(ctx context.Context, in *npool.GetUserRolesByAppUserRequest) (*npool.GetUserRolesByAppUserResponse, error) {
	return nil, nil
}

func (s *Server) DeleteAppRoleUser(ctx context.Context, in *npool.DeleteAppRoleUserRequest) (*npool.DeleteAppRoleUserResponse, error) {
	return nil, nil
}
