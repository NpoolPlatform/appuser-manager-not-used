package history

import (
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/appuser/mgr/v2/login/history"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func validate(info *npool.HistoryReq) error {
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

	if info.ClientIP == nil || info.GetClientIP() == "" {
		logger.Sugar().Errorw("validate", "ClientIP", info.ClientIP, "GetClientIP", info.GetClientIP())
		return status.Error(codes.InvalidArgument, "ClientIP is empty")
	}

	if info.UserAgent == nil || info.GetUserAgent() == "" {
		logger.Sugar().Errorw("validate", "UserAgent", info.UserAgent, "GetUserAgent", info.GetUserAgent())
		return status.Error(codes.InvalidArgument, "UserAgent is empty")
	}

	return nil
}

func validateMany(infos []*npool.HistoryReq) error {
	for _, info := range infos {
		if err := validate(info); err != nil {
			return err
		}
	}
	return nil
}

func Validate(info *npool.HistoryReq) error {
	return validate(info)
}

func ValidateMany(infos []*npool.HistoryReq) error {
	return validateMany(infos)
}
