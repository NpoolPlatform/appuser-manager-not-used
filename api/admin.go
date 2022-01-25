package api

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	adminmw "github.com/NpoolPlatform/appuser-manager/pkg/middleware/admin"
	npool "github.com/NpoolPlatform/message/npool/appusermgr"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateAdminApps(ctx context.Context, in *npool.CreateAdminAppsRequest) (*npool.CreateAdminAppsResponse, error) {
	resp, err := adminmw.CreateAdminApps(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("fail create admin apps: %v", err)
		return &npool.CreateAdminAppsResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) GetAdminApps(ctx context.Context, in *npool.GetAdminAppsRequest) (*npool.GetAdminAppsResponse, error) {
	resp, err := adminmw.GetAdminApps(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("fail get admin apps: %v", err)
		return &npool.GetAdminAppsResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) CreateGenesisRole(ctx context.Context, in *npool.CreateGenesisRoleRequest) (*npool.CreateGenesisRoleResponse, error) {
	resp, err := adminmw.CreateGenesisRole(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("fail create genesis role: %v", err)
		return &npool.CreateGenesisRoleResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) GetGenesisRole(ctx context.Context, in *npool.GetGenesisRoleRequest) (*npool.GetGenesisRoleResponse, error) {
	resp, err := adminmw.GetGenesisRole(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("fail get genesis role: %v", err)
		return &npool.GetGenesisRoleResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) CreateGenesisRoleUser(ctx context.Context, in *npool.CreateGenesisRoleUserRequest) (*npool.CreateGenesisRoleUserResponse, error) {
	resp, err := adminmw.CreateGenesisRoleUser(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("fail create genesis role user: %v", err)
		return &npool.CreateGenesisRoleUserResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}
