package api

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	constant "github.com/NpoolPlatform/appuser-manager/pkg/const"
	approlecrud "github.com/NpoolPlatform/appuser-manager/pkg/crud/approle"
	approleusercrud "github.com/NpoolPlatform/appuser-manager/pkg/crud/approleuser"
	npool "github.com/NpoolPlatform/message/npool/appusermgr"

	"golang.org/x/xerrors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateAppRole(ctx context.Context, in *npool.CreateAppRoleRequest) (*npool.CreateAppRoleResponse, error) {
	if in.GetInfo().GetRole() == constant.GenesisRole {
		return &npool.CreateAppRoleResponse{}, status.Error(codes.Internal, xerrors.Errorf("permission denied").Error())
	}

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

func (s *Server) GetAppRoleByAppRole(ctx context.Context, in *npool.GetAppRoleByAppRoleRequest) (*npool.GetAppRoleByAppRoleResponse, error) {
	resp, err := approlecrud.GetByAppRole(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("fail get app role by app role: %v", err)
		return &npool.GetAppRoleByAppRoleResponse{}, status.Error(codes.Internal, err.Error())
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
	role, err := approlecrud.Get(ctx, &npool.GetAppRoleRequest{
		ID: in.GetInfo().GetRoleID(),
	})
	if err != nil {
		return &npool.CreateAppRoleUserResponse{}, status.Error(codes.Internal, xerrors.Errorf("fail get role: %v", err).Error())
	}

	if role.Info.Role == constant.GenesisRole {
		return &npool.CreateAppRoleUserResponse{}, status.Error(codes.Internal, xerrors.Errorf("permission denied").Error())
	}

	resp, err := approleusercrud.Create(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("fail create app role user: %v", err)
		return &npool.CreateAppRoleUserResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) GetAppRoleUser(ctx context.Context, in *npool.GetAppRoleUserRequest) (*npool.GetAppRoleUserResponse, error) {
	resp, err := approleusercrud.Get(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("fail get app role user: %v", err)
		return &npool.GetAppRoleUserResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) GetAppRoleUsersByAppRole(ctx context.Context, in *npool.GetAppRoleUsersByAppRoleRequest) (*npool.GetAppRoleUsersByAppRoleResponse, error) {
	resp, err := approleusercrud.GetUsersByAppRole(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("fail get users by app role: %v", err)
		return &npool.GetAppRoleUsersByAppRoleResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) GetUserRolesByAppUser(ctx context.Context, in *npool.GetUserRolesByAppUserRequest) (*npool.GetUserRolesByAppUserResponse, error) {
	return nil, nil
}

func (s *Server) DeleteAppRoleUser(ctx context.Context, in *npool.DeleteAppRoleUserRequest) (*npool.DeleteAppRoleUserResponse, error) {
	resp, err := approleusercrud.Delete(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("fail delete app role user: %v", err)
		return &npool.DeleteAppRoleUserResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}
