package approleuser

import (
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/appuser/mgr/v2/approleuser"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func validate(info *npool.AppRoleUserReq) error {
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

	if info.RoleID == nil {
		logger.Sugar().Errorw("validate", "RoleID", info.RoleID)
		return status.Error(codes.InvalidArgument, "RoleID is empty")
	}

	if _, err := uuid.Parse(info.GetRoleID()); err != nil {
		logger.Sugar().Errorw("validate", "RoleID", info.GetRoleID(), "error", err)
		return status.Error(codes.InvalidArgument, "RoleID is invalid")
	}

	return nil
}

func Validate(info *npool.AppRoleUserReq) error {
	return validate(info)
}
