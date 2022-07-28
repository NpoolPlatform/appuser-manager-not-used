package appusersecret

import (
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/appuser/mgr/v2/appusersecret"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func validate(info *npool.AppUserSecretReq) error {

	if info.AppID == nil {
		logger.Sugar().Errorw("validate", "AppID", info.AppID)
		return status.Error(codes.InvalidArgument, "AppID is empty")
	}

	if _, err := uuid.Parse(info.GetAppID()); err != nil {
		logger.Sugar().Errorw("validate", "AppID", info.GetAppID(), "error", err)
		return status.Error(codes.InvalidArgument, "AppID is invalid")
	}

	if info.UserID == nil {
		logger.Sugar().Errorw("validate", "UserID", info.UserID)
		return status.Error(codes.InvalidArgument, "UserID is empty")
	}

	if _, err := uuid.Parse(info.GetUserID()); err != nil {
		logger.Sugar().Errorw("validate", "UserID", info.GetUserID(), "error", err)
		return status.Error(codes.InvalidArgument, "UserID is invalid")
	}

	if info.PasswordHash == nil || info.GetPasswordHash() == "" {
		logger.Sugar().Errorw("validate", "PasswordHash", info.PasswordHash, "GetPasswordHash", info.GetPasswordHash())
		return status.Error(codes.InvalidArgument, "PasswordHash is empty")
	}

	if info.Salt == nil || info.GetSalt() == "" {
		logger.Sugar().Errorw("validate", "Salt", info.Salt, "GetSalt", info.GetSalt())
		return status.Error(codes.InvalidArgument, "GetSalt is empty")
	}

	return nil
}
