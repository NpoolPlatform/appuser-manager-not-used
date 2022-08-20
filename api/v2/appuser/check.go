package appuser

import (
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/appuser/mgr/v2/appuser"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func validate(info *npool.AppUserReq) error {
	if info.AppID == nil {
		logger.Sugar().Errorw("validate", "AppID", info.AppID)
		return status.Error(codes.InvalidArgument, "AppID is empty")
	}

	if _, err := uuid.Parse(info.GetAppID()); err != nil {
		logger.Sugar().Errorw("validate", "AppID", info.GetAppID(), "error", err)
		return status.Error(codes.InvalidArgument, "AppID is invalid")
	}

	if info.ImportFromApp != nil {
		if _, err := uuid.Parse(info.GetImportFromApp()); err != nil {
			logger.Sugar().Errorw("validate", "ImportFromApp", info.GetImportFromApp(), "error", err)
			return status.Error(codes.InvalidArgument, "ImportFromApp is invalid")
		}
	}

	if info.EmailAddress == nil && info.PhoneNO == nil {
		logger.Sugar().Errorw("validate", "EmailAddress", info.EmailAddress, "PhoneNO", info.PhoneNO)
		return status.Error(codes.InvalidArgument, "EmailAddress and PhoneNO are empty")
	}

	if info.GetEmailAddress() == "" && info.GetPhoneNO() == "" {
		logger.Sugar().Errorw("validate", "EmailAddress", info.EmailAddress, "PhoneNO", info.PhoneNO)
		return status.Error(codes.InvalidArgument, "EmailAddress and PhoneNO are invalid")
	}

	return nil
}

func Validate(info *npool.AppUserReq) error {
	return validate(info)
}
