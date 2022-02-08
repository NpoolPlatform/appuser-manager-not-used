package api

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	approlecrud "github.com/NpoolPlatform/appuser-manager/pkg/crud/approle"
	approleusercrud "github.com/NpoolPlatform/appuser-manager/pkg/crud/approleuser"
	approlemw "github.com/NpoolPlatform/appuser-manager/pkg/middleware/approle"
	approleusermw "github.com/NpoolPlatform/appuser-manager/pkg/middleware/approleuser"
	npool "github.com/NpoolPlatform/message/npool/appusermgr"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateAppRole(ctx context.Context, in *npool.CreateAppRoleRequest) (*npool.CreateAppRoleResponse, error) {
	resp, err := approlemw.Create(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("fail create app role: %v", err)
		return &npool.CreateAppRoleResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) CreateAppRoleForOtherApp(ctx context.Context, in *npool.CreateAppRoleForOtherAppRequest) (*npool.CreateAppRoleForOtherAppResponse, error) {
	resp, err := approlemw.CreateForOtherApp(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("fail create app role for other app: %v", err)
		return &npool.CreateAppRoleForOtherAppResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) GetAppRole(ctx context.Context, in *npool.GetAppRoleRequest) (*npool.GetAppRoleResponse, error) {
	resp, err := approlecrud.Get(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("fail get app role: %v", err)
		return &npool.GetAppRoleResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) GetAppRoleByAppRole(ctx context.Context, in *npool.GetAppRoleByAppRoleRequest) (*npool.GetAppRoleByAppRoleResponse, error) {
	resp, err := approlecrud.GetByAppRole(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("fail get app role by app role: %v", err)
		return &npool.GetAppRoleByAppRoleResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) GetAppRolesByApp(ctx context.Context, in *npool.GetAppRolesByAppRequest) (*npool.GetAppRolesByAppResponse, error) {
	resp, err := approlecrud.GetByApp(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("fail get app roles by app: %v", err)
		return &npool.GetAppRolesByAppResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) GetAppRolesByOtherApp(ctx context.Context, in *npool.GetAppRolesByOtherAppRequest) (*npool.GetAppRolesByOtherAppResponse, error) {
	resp, err := approlecrud.GetByApp(ctx, &npool.GetAppRolesByAppRequest{
		AppID: in.GetTargetAppID(),
	})
	if err != nil {
		logger.Sugar().Errorf("fail get app role: %v", err)
		return &npool.GetAppRolesByOtherAppResponse{}, status.Error(codes.Internal, err.Error())
	}
	return &npool.GetAppRolesByOtherAppResponse{
		Infos: resp.Infos,
	}, nil
}

func (s *Server) UpdateAppRole(ctx context.Context, in *npool.UpdateAppRoleRequest) (*npool.UpdateAppRoleResponse, error) {
	resp, err := approlecrud.Update(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("fail update app role: %v", err)
		return &npool.UpdateAppRoleResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) CreateAppRoleUser(ctx context.Context, in *npool.CreateAppRoleUserRequest) (*npool.CreateAppRoleUserResponse, error) {
	resp, err := approleusermw.CreateAppRoleUser(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("fail create app role user: %v", err)
		return &npool.CreateAppRoleUserResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) CreateAppRoleUserForOtherAppUser(ctx context.Context, in *npool.CreateAppRoleUserForOtherAppUserRequest) (*npool.CreateAppRoleUserForOtherAppUserResponse, error) {
	info := in.GetInfo()
	info.AppID = in.GetTargetAppID()
	info.UserID = in.GetTargetUserID()
	resp, err := s.CreateAppRoleUser(ctx, &npool.CreateAppRoleUserRequest{
		Info: info,
	})
	if err != nil {
		logger.Sugar().Errorf("fail create app role user for other app: %v", err)
		return &npool.CreateAppRoleUserForOtherAppUserResponse{}, status.Error(codes.Internal, err.Error())
	}
	return &npool.CreateAppRoleUserForOtherAppUserResponse{
		Info: resp.Info,
	}, nil
}

func (s *Server) GetAppRoleUser(ctx context.Context, in *npool.GetAppRoleUserRequest) (*npool.GetAppRoleUserResponse, error) {
	resp, err := approleusercrud.Get(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("fail get app role user: %v", err)
		return &npool.GetAppRoleUserResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) GetAppRoleUserByAppUser(ctx context.Context, in *npool.GetAppRoleUserByAppUserRequest) (*npool.GetAppRoleUserByAppUserResponse, error) {
	resp, err := approleusercrud.GetByAppUser(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("fail get app role user by app user: %v", err)
		return &npool.GetAppRoleUserByAppUserResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) GetAppRoleUsersByAppRole(ctx context.Context, in *npool.GetAppRoleUsersByAppRoleRequest) (*npool.GetAppRoleUsersByAppRoleResponse, error) {
	resp, err := approleusercrud.GetUsersByAppRole(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("fail get users by app role: %v", err)
		return &npool.GetAppRoleUsersByAppRoleResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) GetAppRoleUsersByApp(ctx context.Context, in *npool.GetAppRoleUsersByAppRequest) (*npool.GetAppRoleUsersByAppResponse, error) {
	resp, err := approleusercrud.GetUsersByApp(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("fail get users by app: %v", err)
		return &npool.GetAppRoleUsersByAppResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) GetAppRoleUsersByOtherApp(ctx context.Context, in *npool.GetAppRoleUsersByOtherAppRequest) (*npool.GetAppRoleUsersByOtherAppResponse, error) {
	resp, err := approleusercrud.GetUsersByApp(ctx, &npool.GetAppRoleUsersByAppRequest{
		AppID: in.GetTargetAppID(),
	})
	if err != nil {
		logger.Sugar().Errorf("fail get users by app: %v", err)
		return &npool.GetAppRoleUsersByOtherAppResponse{}, status.Error(codes.Internal, err.Error())
	}
	return &npool.GetAppRoleUsersByOtherAppResponse{
		Infos: resp.Infos,
	}, nil
}

func (s *Server) DeleteAppRoleUser(ctx context.Context, in *npool.DeleteAppRoleUserRequest) (*npool.DeleteAppRoleUserResponse, error) {
	resp, err := approleusercrud.Delete(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("fail delete app role user: %v", err)
		return &npool.DeleteAppRoleUserResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}
