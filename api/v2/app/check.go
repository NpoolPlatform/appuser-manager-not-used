package app

import (
	"fmt"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/appuser/mgr/v2/app"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func validate(info *npool.AppReq) error {
	if info.CreatedBy == nil {
		logger.Sugar().Errorw("validate", "CreatedBy", info.CreatedBy)
		return status.Error(codes.InvalidArgument, "CreatedBy is empty")
	}

	if _, err := uuid.Parse(info.GetCreatedBy()); err != nil {
		logger.Sugar().Errorw("validate", "CreatedBy", info.GetCreatedBy(), "error", err)
		return status.Error(codes.InvalidArgument, "CreatedBy is invalid")
	}

	if info.Name == nil || info.GetName() == "" {
		logger.Sugar().Errorw("validate", "Name", info.Name, "GetName", info.GetName())
		return status.Error(codes.InvalidArgument, "Name is empty")
	}

	if info.Logo == nil || info.GetLogo() == "" {
		logger.Sugar().Errorw("validate", "Logo", info.Logo, "GetLogo", info.GetLogo())
		return status.Error(codes.InvalidArgument, "Logo is empty")
	}

	return nil
}

func validateMany(infos []*npool.AppReq) error {
	names := map[string]struct{}{}

	for _, info := range infos {
		if err := validate(info); err != nil {
			return err
		}

		if _, ok := names[info.GetName()]; ok {
			return fmt.Errorf("duplicate app name")
		}

		names[info.GetName()] = struct{}{}
	}

	return nil
}
