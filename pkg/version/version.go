package version

import (
	"fmt"

	npool "github.com/NpoolPlatform/message/npool"

	logger "github.com/NpoolPlatform/go-service-framework/pkg/logger" //nolint
	cv "github.com/NpoolPlatform/go-service-framework/pkg/version"    //nolint
)

func Version() (*npool.VersionResponse, error) {
	info, err := cv.GetVersion()
	if err != nil {
		logger.Sugar().Errorf("get service version error: %+w", err)
		return nil, fmt.Errorf("get service version error: %w", err)
	}
	return &npool.VersionResponse{
		Info: info,
	}, nil
}
