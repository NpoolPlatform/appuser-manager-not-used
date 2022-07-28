package v1

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/appuser/mgr/v1"

	appusercrud "github.com/NpoolPlatform/appuser-manager/pkg/crud/v1/appuser"
	appusercontrolcrud "github.com/NpoolPlatform/appuser-manager/pkg/crud/v1/appusercontrol"
	appuserextracrud "github.com/NpoolPlatform/appuser-manager/pkg/crud/v1/appuserextra"
	appusersecretcrud "github.com/NpoolPlatform/appuser-manager/pkg/crud/v1/appusersecret"
	appuserthirdpartycrud "github.com/NpoolPlatform/appuser-manager/pkg/crud/v1/appuserthirdparty"
	banappusercrud "github.com/NpoolPlatform/appuser-manager/pkg/crud/v1/banappuser"
	appusermw "github.com/NpoolPlatform/appuser-manager/pkg/middleware/appuser"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateAppUser(ctx context.Context, in *npool.CreateAppUserRequest) (*npool.CreateAppUserResponse, error) {
	resp, err := appusercrud.Create(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("fail create app user: %v", err)
		return &npool.CreateAppUserResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) GetAppUser(ctx context.Context, in *npool.GetAppUserRequest) (*npool.GetAppUserResponse, error) {
	resp, err := appusercrud.Get(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("fail get app user: %v", err)
		return &npool.GetAppUserResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) GetAppUserByAppUser(ctx context.Context, in *npool.GetAppUserByAppUserRequest) (*npool.GetAppUserByAppUserResponse, error) {
	resp, err := appusercrud.GetByAppUser(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("fail get app user by app user: %v", err)
		return &npool.GetAppUserByAppUserResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) GetAppUsersByApp(ctx context.Context, in *npool.GetAppUsersByAppRequest) (*npool.GetAppUsersByAppResponse, error) {
	resp, err := appusercrud.GetByApp(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("fail get app user by app: %v", err)
		return &npool.GetAppUsersByAppResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) UpdateAppUser(ctx context.Context, in *npool.UpdateAppUserRequest) (*npool.UpdateAppUserResponse, error) {
	resp, err := appusercrud.Update(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("fail update app user: %v", err)
		return &npool.UpdateAppUserResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) CreateAppUserSecret(ctx context.Context, in *npool.CreateAppUserSecretRequest) (*npool.CreateAppUserSecretResponse, error) {
	resp, err := appusersecretcrud.Create(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("fail create app user secret: %v", err)
		return &npool.CreateAppUserSecretResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) GetAppUserSecret(ctx context.Context, in *npool.GetAppUserSecretRequest) (*npool.GetAppUserSecretResponse, error) {
	resp, err := appusersecretcrud.Get(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("fail get app user secret: %v", err)
		return &npool.GetAppUserSecretResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) GetAppUserSecretByAppUser(ctx context.Context, in *npool.GetAppUserSecretByAppUserRequest) (*npool.GetAppUserSecretByAppUserResponse, error) {
	resp, err := appusersecretcrud.GetByAppUser(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("fail get app user secret by app user: %v", err)
		return &npool.GetAppUserSecretByAppUserResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) UpdateAppUserSecret(ctx context.Context, in *npool.UpdateAppUserSecretRequest) (*npool.UpdateAppUserSecretResponse, error) {
	resp, err := appusersecretcrud.Update(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("fail update app user secret: %v", err)
		return &npool.UpdateAppUserSecretResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) CreateAppUserExtra(ctx context.Context, in *npool.CreateAppUserExtraRequest) (*npool.CreateAppUserExtraResponse, error) {
	resp, err := appuserextracrud.Create(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("fail create app user extra: %v", err)
		return &npool.CreateAppUserExtraResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) GetAppUserExtra(ctx context.Context, in *npool.GetAppUserExtraRequest) (*npool.GetAppUserExtraResponse, error) {
	resp, err := appuserextracrud.Get(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("fail get app user extra: %v", err)
		return &npool.GetAppUserExtraResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) GetAppUserExtraByAppUser(ctx context.Context, in *npool.GetAppUserExtraByAppUserRequest) (*npool.GetAppUserExtraByAppUserResponse, error) {
	resp, err := appuserextracrud.GetByAppUser(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("fail get app user extra by app user: %v", err)
		return &npool.GetAppUserExtraByAppUserResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) UpdateAppUserExtra(ctx context.Context, in *npool.UpdateAppUserExtraRequest) (*npool.UpdateAppUserExtraResponse, error) {
	resp, err := appuserextracrud.Update(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("fail update app user extra: %v", err)
		return &npool.UpdateAppUserExtraResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) CreateBanAppUser(ctx context.Context, in *npool.CreateBanAppUserRequest) (*npool.CreateBanAppUserResponse, error) {
	resp, err := banappusercrud.Create(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("fail create ban app user: %v", err)
		return &npool.CreateBanAppUserResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) GetBanAppUser(ctx context.Context, in *npool.GetBanAppUserRequest) (*npool.GetBanAppUserResponse, error) {
	resp, err := banappusercrud.Get(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("fail get ban app user: %v", err)
		return &npool.GetBanAppUserResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) GetBanAppUserByAppUser(ctx context.Context, in *npool.GetBanAppUserByAppUserRequest) (*npool.GetBanAppUserByAppUserResponse, error) {
	resp, err := banappusercrud.GetByAppUser(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("fail get ban app user by app user: %v", err)
		return &npool.GetBanAppUserByAppUserResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) DeleteBanAppUser(ctx context.Context, in *npool.DeleteBanAppUserRequest) (*npool.DeleteBanAppUserResponse, error) {
	resp, err := banappusercrud.Delete(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("fail delete ban app user: %v", err)
		return &npool.DeleteBanAppUserResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) CreateAppUserControl(ctx context.Context, in *npool.CreateAppUserControlRequest) (*npool.CreateAppUserControlResponse, error) {
	resp, err := appusercontrolcrud.Create(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("fail create app user control: %v", err)
		return &npool.CreateAppUserControlResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) GetAppUserControl(ctx context.Context, in *npool.GetAppUserControlRequest) (*npool.GetAppUserControlResponse, error) {
	resp, err := appusercontrolcrud.Get(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("fail get app user control: %v", err)
		return &npool.GetAppUserControlResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) GetAppUserControlByAppUser(ctx context.Context, in *npool.GetAppUserControlByAppUserRequest) (*npool.GetAppUserControlByAppUserResponse, error) {
	resp, err := appusercontrolcrud.GetByAppUser(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("fail get app user control by app user: %v", err)
		return &npool.GetAppUserControlByAppUserResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) UpdateAppUserControl(ctx context.Context, in *npool.UpdateAppUserControlRequest) (*npool.UpdateAppUserControlResponse, error) {
	resp, err := appusercontrolcrud.Update(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("fail update app user control: %v", err)
		return &npool.UpdateAppUserControlResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) GetAppUserInfoByAppUser(ctx context.Context, in *npool.GetAppUserInfoByAppUserRequest) (*npool.GetAppUserInfoByAppUserResponse, error) {
	resp, err := appusermw.GetAppUserInfoByAppUser(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("fail get app user info by app user: %v", err)
		return &npool.GetAppUserInfoByAppUserResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) GetAppUserInfo(ctx context.Context, in *npool.GetAppUserInfoRequest) (*npool.GetAppUserInfoResponse, error) {
	resp, err := appusermw.GetAppUserInfo(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("fail get app user info: %v", err)
		return &npool.GetAppUserInfoResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) GetAppUserInfosByApp(ctx context.Context, in *npool.GetAppUserInfosByAppRequest) (*npool.GetAppUserInfosByAppResponse, error) {
	resp, err := appusermw.GetAppUserInfosByApp(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("fail get app user infos by app: %v", err)
		return &npool.GetAppUserInfosByAppResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) GetAppUserInfosByOtherApp(ctx context.Context, in *npool.GetAppUserInfosByOtherAppRequest) (*npool.GetAppUserInfosByOtherAppResponse, error) {
	resp, err := appusermw.GetAppUserInfosByApp(ctx, &npool.GetAppUserInfosByAppRequest{
		AppID: in.GetTargetAppID(),
	})
	if err != nil {
		logger.Sugar().Errorf("fail get app user infos by app: %v", err)
		return &npool.GetAppUserInfosByOtherAppResponse{}, status.Error(codes.Internal, err.Error())
	}
	return &npool.GetAppUserInfosByOtherAppResponse{
		Infos: resp.Infos,
	}, nil
}

func (s *Server) CreateAppUserWithSecret(ctx context.Context, in *npool.CreateAppUserWithSecretRequest) (*npool.CreateAppUserWithSecretResponse, error) {
	resp, err := appusermw.CreateWithSecret(ctx, in, true)
	if err != nil {
		logger.Sugar().Errorf("fail create app user with secret: %v", err)
		return &npool.CreateAppUserWithSecretResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) CreateAppUserWithThirdParty(ctx context.Context, in *npool.CreateAppUserWithThirdPartyRequest) (*npool.CreateAppUserWithThirdPartyResponse, error) {
	resp, err := appusermw.CreateWithThirdParty(ctx, in, true)
	if err != nil {
		logger.Sugar().Errorf("fail create app user with third: %v", err)
		return &npool.CreateAppUserWithThirdPartyResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) GetAppUserThirdPartyByAppThirdPartyID(ctx context.Context, in *npool.GetAppUserThirdPartyByAppThirdPartyIDRequest) (*npool.GetAppUserThirdPartyByAppThirdPartyIDResponse, error) {
	resp, err := appuserthirdpartycrud.GetByAppUserThirdParty(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("fail create app user ThirdParty: %v", err)
		return &npool.GetAppUserThirdPartyByAppThirdPartyIDResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) CreateAppUserWithSecretRevert(ctx context.Context, in *npool.CreateAppUserWithSecretRequest) (*npool.CreateAppUserWithSecretResponse, error) {
	resp, err := appusermw.CreateWithSecretRevert(ctx, in, true)
	if err != nil {
		logger.Sugar().Errorf("fail create app user with secret: %v", err)
		return &npool.CreateAppUserWithSecretResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) GetAppUserByAppAccount(ctx context.Context, in *npool.GetAppUserByAppAccountRequest) (*npool.GetAppUserByAppAccountResponse, error) {
	resp, err := appusercrud.GetByAppAccount(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("fail get app user by app account: %v", err)
		return &npool.GetAppUserByAppAccountResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) VerifyAppUserByAppAccountPassword(ctx context.Context, in *npool.VerifyAppUserByAppAccountPasswordRequest) (*npool.VerifyAppUserByAppAccountPasswordResponse, error) {
	resp, err := appusermw.VerifyByAppAccountPassword(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("fail verify app user by app account password: %v", err)
		return &npool.VerifyAppUserByAppAccountPasswordResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) GetUserRolesByAppUser(ctx context.Context, in *npool.GetUserRolesByAppUserRequest) (*npool.GetUserRolesByAppUserResponse, error) {
	resp, err := appusermw.GetRolesByAppUser(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("fail get user roles by app user: %v", err)
		return &npool.GetUserRolesByAppUserResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}
