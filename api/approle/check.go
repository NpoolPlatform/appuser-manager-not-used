package approle

import (
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/appuser/mgr/v2/approle"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func validate(info *npool.AppRoleReq) error {
	if info.AppID == nil {
		logger.Sugar().Errorw("validate", "AppID", info.AppID)
		return status.Error(codes.InvalidArgument, "AppID is empty")
	}

	if _, err := uuid.Parse(info.GetAppID()); err != nil {
		logger.Sugar().Errorw("validate", "AppID", info.GetAppID(), "error", err)
		return status.Error(codes.InvalidArgument, "AppID is invalid")
	}

	if info.CreatedBy == nil {
		logger.Sugar().Errorw("validate", "CreatedBy", info.CreatedBy)
		return status.Error(codes.InvalidArgument, "CreatedBy is empty")
	}

	if _, err := uuid.Parse(info.GetCreatedBy()); err != nil {
		logger.Sugar().Errorw("validate", "CreatedBy", info.GetCreatedBy(), "error", err)
		return status.Error(codes.InvalidArgument, "CreatedBy is invalid")
	}

	if info.Role == nil || info.GetRole() == "" {
		logger.Sugar().Errorw("validate", "Role", info.Role, "GetRole", info.GetRole())
		return status.Error(codes.InvalidArgument, "Role is empty")
	}

	return nil
}

func Validate(info *npool.AppRoleReq) error {
	return validate(info)
}
