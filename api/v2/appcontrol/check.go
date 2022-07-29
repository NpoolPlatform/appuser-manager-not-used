package appcontrol

import (
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/appuser/mgr/v2/appcontrol"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func validate(info *npool.AppControlReq) error {
	if info.AppID == nil {
		logger.Sugar().Errorw("validate", "AppID", info.AppID)
		return status.Error(codes.InvalidArgument, "AppID is empty")
	}

	if _, err := uuid.Parse(info.GetAppID()); err != nil {
		logger.Sugar().Errorw("validate", "AppID", info.GetAppID(), "error", err)
		return status.Error(codes.InvalidArgument, "AppID is invalid")
	}

	if info.ID == nil {
		logger.Sugar().Errorw("validate", "ID", info.ID)
		return status.Error(codes.InvalidArgument, "ID is empty")
	}

	if _, err := uuid.Parse(info.GetID()); err != nil {
		logger.Sugar().Errorw("validate", "ID", info.GetID(), "error", err)
		return status.Error(codes.InvalidArgument, "ID is invalid")
	}

	return nil
}
