package banapp

import (
	"fmt"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/appuser/mgr/v2/banapp"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func validate(info *npool.BanAppReq) error {
	if info.AppID == nil {
		logger.Sugar().Errorw("validate", "AppID", info.AppID)
		return status.Error(codes.InvalidArgument, "AppID is empty")
	}

	if _, err := uuid.Parse(info.GetAppID()); err != nil {
		logger.Sugar().Errorw("validate", "AppID", info.GetAppID(), "error", err)
		return status.Error(codes.InvalidArgument, "AppID is invalid")
	}

	if info.Message == nil || info.GetMessage() == "" {
		logger.Sugar().Errorw("validate", "Message", info.Message, "GetMessage", info.GetMessage())
		return status.Error(codes.InvalidArgument, "Message is empty")
	}

	return nil
}

func validateMany(infos []*npool.BanAppReq) error {
	appIDs := map[string]struct{}{}

	for _, info := range infos {
		if err := validate(info); err != nil {
			return err
		}

		if _, ok := appIDs[info.GetAppID()]; ok {
			return fmt.Errorf("duplicate app id")
		}

		appIDs[info.GetAppID()] = struct{}{}
	}

	return nil
}

func Validate(info *npool.BanAppReq) error {
	return validate(info)
}
