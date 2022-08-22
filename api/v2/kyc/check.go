package kyc

import (
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/appuser/mgr/v2/kyc"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func validate(info *npool.KycReq) error {
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

	if info.GetIDNumber() == "" {
		logger.Sugar().Errorw("validate", "IDNumber", info.GetIDNumber())
		return status.Error(codes.InvalidArgument, "IDNumber is empty")
	}

	checkEntityType := false
	for key := range npool.KycEntityType_value {
		if info.EntityType.String() == key {
			checkEntityType = true
		}
	}

	if !checkEntityType {
		logger.Sugar().Errorw("validate", "EntityType", info.GetEntityType())
		return status.Error(codes.InvalidArgument, "EntityType is invalid")
	}

	checkDocumentType := false
	for key := range npool.KycDocumentType_value {
		if info.DocumentType.String() == key {
			checkDocumentType = true
		}
	}

	if !checkDocumentType {
		logger.Sugar().Errorw("validate", "DocumentType", info.GetDocumentType())
		return status.Error(codes.InvalidArgument, "DocumentType is invalid")
	}

	return nil
}

func Validate(info *npool.KycReq) error {
	return validate(info)
}
