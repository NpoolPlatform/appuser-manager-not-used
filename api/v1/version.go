//go:build !codeanalysis
// +build !codeanalysis

package v1

import (
	"context"

	"github.com/NpoolPlatform/appuser-manager/pkg/version"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) Version(ctx context.Context, in *emptypb.Empty) (*npool.VersionResponse, error) {
	resp, err := version.Version()
	if err != nil {
		logger.Sugar().Errorw("[Version] get service version error: %w", err)
		return &npool.VersionResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}
