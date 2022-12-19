package subscriber

import (
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/appuser/mgr/v2/subscriber"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func validate(info *npool.SubscriberReq) error {
	if info.AppID == nil {
		logger.Sugar().Errorw("validate", "AppID", info.AppID)
		return status.Error(codes.InvalidArgument, "AppID is empty")
	}

	if _, err := uuid.Parse(info.GetAppID()); err != nil {
		logger.Sugar().Errorw("validate", "AppID", info.GetAppID(), "error", err)
		return status.Error(codes.InvalidArgument, "AppID is invalid")
	}

	if info.EmailAddress == nil || info.GetEmailAddress() == "" {
		logger.Sugar().Errorw("validate", "EmailAddress", info.EmailAddress)
		return status.Error(codes.InvalidArgument, "EmailAddress is empty")
	}

	return nil
}

func validateMany(infos []*npool.SubscriberReq) error {
	for _, info := range infos {
		if err := validate(info); err != nil {
			return err
		}
	}
	return nil
}

func Validate(info *npool.SubscriberReq) error {
	return validate(info)
}

func ValidateMany(infos []*npool.SubscriberReq) error {
	return validateMany(infos)
}
