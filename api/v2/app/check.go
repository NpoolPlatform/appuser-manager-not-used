package app

import (
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/appuser/mgr/v2/app"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func checkAppInfo(info *npool.AppReq) error {
	if _, err := uuid.Parse(info.GetCreatedBy()); err != nil {
		logger.Sugar().Error("CreatedBy is invalid")
		return status.Error(codes.InvalidArgument, "CreatedBy is invalid")
	}

	if info.Name == nil {
		logger.Sugar().Error("Name is empty")
		return status.Error(codes.InvalidArgument, "Name is empty")
	}

	if info.GetLogo() == "" {
		logger.Sugar().Error("Logo is empty")
		return status.Error(codes.InvalidArgument, "Logo is empty")
	}

	return nil
}
