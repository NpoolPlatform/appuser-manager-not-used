package appuserthirdparty

import (
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/appuser/mgr/v2/appuserthirdparty"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func validate(info *npool.AppUserThirdPartyReq) error {
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

	if info.ThirdPartyID == nil || info.GetThirdPartyID() == "" {
		logger.Sugar().Errorw("validate", "ThirdPartyID", info.ThirdPartyID, "GetThirdPartyID", info.GetThirdPartyID())
		return status.Error(codes.InvalidArgument, "ThirdPartyID is empty")
	}

	if info.ThirdPartyUserID == nil || info.GetThirdPartyUserID() == "" {
		logger.Sugar().Errorw("validate", "ThirdPartyUserID", info.ThirdPartyUserID, "GetThirdPartyUserID", info.GetThirdPartyUserID())
		return status.Error(codes.InvalidArgument, "ThirdPartyUserID is empty")
	}

	return nil
}

func Validate(info *npool.AppUserThirdPartyReq) error {
	return validate(info)
}
