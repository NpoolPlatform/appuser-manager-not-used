package api

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	adminmw "github.com/NpoolPlatform/appuser-manager/pkg/middleware/admin"
	npool "github.com/NpoolPlatform/message/npool/appusermgr"

	"golang.org/x/xerrors"
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

func (s *Server) GetGenesisAppRoleUsersByOtherApp(ctx context.Context, in *npool.GetGenesisAppRoleUsersByOtherAppRequest) (*npool.GetGenesisAppRoleUsersByOtherAppResponse, error) {
	role, err := adminmw.GetGenesisRole(ctx, &npool.GetGenesisRoleRequest{})
	if err != nil {
		return nil, xerrors.Errorf("fail get genesis role: %v", err)
	}

	resp, err := s.GetAppRoleUsersByAppRole(ctx, &npool.GetAppRoleUsersByAppRoleRequest{
		AppID:  in.GetAppID(),
		RoleID: role.Info.ID,
	})
	if err != nil {
		logger.Sugar().Errorw("fail get genesis app role user by app: %v", err)
		return &npool.GetGenesisAppRoleUsersByOtherAppResponse{}, status.Error(codes.Internal, err.Error())
	}
	return &npool.GetGenesisAppRoleUsersByOtherAppResponse{
		Infos: resp.Infos,
	}, nil
}

func (s *Server) GetAppRoleUsersByOtherAppRole(ctx context.Context, in *npool.GetAppRoleUsersByOtherAppRoleRequest) (*npool.GetAppRoleUsersByOtherAppRoleResponse, error) {
	resp, err := s.GetAppRoleUsersByAppRole(ctx, &npool.GetAppRoleUsersByAppRoleRequest{
		AppID:  in.GetAppID(),
		RoleID: in.GetRoleID(),
	})
	if err != nil {
		logger.Sugar().Errorw("fail get app role user by other app role: %v", err)
		return &npool.GetAppRoleUsersByOtherAppRoleResponse{}, status.Error(codes.Internal, err.Error())
	}
	return &npool.GetAppRoleUsersByOtherAppRoleResponse{
		Infos: resp.Infos,
	}, nil
}
