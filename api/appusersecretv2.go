//go:build !codeanalysis
// +build !codeanalysis

package api

import (
	"context"

	crud "github.com/NpoolPlatform/appuser-manager/pkg/crud/appusersecretv2"
	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/appusermgrv2/appusersecret"

	"github.com/google/uuid"
)

func checkAppUserSecretInfo(info *npool.AppUserSecret) error {
	if _, err := uuid.Parse(info.GetAppID()); err != nil {
		logger.Sugar().Error("AppID is invalid")
		return status.Error(codes.InvalidArgument, "AppID is invalid")
	}
	if _, err := uuid.Parse(info.GetUserID()); err != nil {
		logger.Sugar().Error("UserID is invalid")
		return status.Error(codes.InvalidArgument, "UserID is invalid")
	}
	if info.GetPasswordHash() == "" {
		logger.Sugar().Error("PasswordHash is invalid")
		return status.Error(codes.InvalidArgument, "PasswordHash is invalid")
	}
	return nil
}

func appUserSecretRowToObject(row *ent.AppUserSecret) *npool.AppUserSecretRes {
	return &npool.AppUserSecretRes{
		Salt:         row.Salt,
		GoogleSecret: row.GoogleSecret,
		ID:           row.ID.String(),
		AppID:        row.AppID.String(),
		UserID:       row.UserID.String(),
		PasswordHash: row.PasswordHash,
	}
}

func (s *Server) CreateAppUserSecretV2(ctx context.Context, in *npool.CreateAppUserSecretRequest) (*npool.CreateAppUserSecretResponse, error) {
	err := checkAppUserSecretInfo(in.GetInfo())
	if err != nil {
		return &npool.CreateAppUserSecretResponse{}, err
	}

	info, err := crud.Create(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorf("fail create AppUserSecret: %v", err)
		return &npool.CreateAppUserSecretResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateAppUserSecretResponse{
		Info: appUserSecretRowToObject(info),
	}, nil
}

func (s *Server) CreateAppUserSecretsV2(ctx context.Context, in *npool.CreateAppUserSecretsRequest) (*npool.CreateAppUserSecretsResponse, error) {
	if len(in.GetInfos()) == 0 {
		return &npool.CreateAppUserSecretsResponse{},
			status.Error(codes.InvalidArgument,
				"Batah create resource must more than 1",
			)
	}
	for _, info := range in.GetInfos() {
		err := checkAppUserSecretInfo(info)
		if err != nil {
			return &npool.CreateAppUserSecretsResponse{}, err
		}
	}

	rows, err := crud.CreateBulk(ctx, in.GetInfos())
	if err != nil {
		logger.Sugar().Errorf("fail create AppUserSecrets: %v", err)
		return &npool.CreateAppUserSecretsResponse{}, status.Error(codes.Internal, err.Error())
	}

	infos := make([]*npool.AppUserSecretRes, 0, len(rows))
	for _, val := range rows {
		infos = append(infos, appUserSecretRowToObject(val))
	}

	return &npool.CreateAppUserSecretsResponse{
		Infos: infos,
	}, nil
}

func (s *Server) UpdateAppUserSecretV2(ctx context.Context, in *npool.UpdateAppUserSecretRequest) (*npool.UpdateAppUserSecretResponse, error) {
	if _, err := uuid.Parse(in.GetInfo().GetID()); err != nil {
		logger.Sugar().Errorf("AppUserSecret id is invalid")
		return &npool.UpdateAppUserSecretResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := crud.Update(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorf("fail update AppUserSecret: %v", err)
		return &npool.UpdateAppUserSecretResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.UpdateAppUserSecretResponse{
		Info: appUserSecretRowToObject(info),
	}, nil
}

func (s *Server) GetAppUserSecretV2(ctx context.Context, in *npool.GetAppUserSecretRequest) (*npool.GetAppUserSecretResponse, error) {
	id, err := uuid.Parse(in.GetID())
	if err != nil {
		return &npool.GetAppUserSecretResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := crud.Row(ctx, id)
	if err != nil {
		logger.Sugar().Errorf("fail get AppUserSecret: %v", err)
		return &npool.GetAppUserSecretResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetAppUserSecretResponse{
		Info: appUserSecretRowToObject(info),
	}, nil
}

func (s *Server) GetAppUserSecretOnlyV2(ctx context.Context, in *npool.GetAppUserSecretOnlyRequest) (*npool.GetAppUserSecretOnlyResponse, error) {
	info, err := crud.RowOnly(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorf("fail get AppUserSecrets: %v", err)
		return &npool.GetAppUserSecretOnlyResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetAppUserSecretOnlyResponse{
		Info: appUserSecretRowToObject(info),
	}, nil
}

func (s *Server) GetAppUserSecretsV2(ctx context.Context, in *npool.GetAppUserSecretsRequest) (*npool.GetAppUserSecretsResponse, error) {
	rows, total, err := crud.Rows(ctx, in.GetConds(), int(in.GetOffset()), int(in.GetLimit()))
	if err != nil {
		logger.Sugar().Errorf("fail get AppUserSecrets: %v", err)
		return &npool.GetAppUserSecretsResponse{}, status.Error(codes.Internal, err.Error())
	}

	infos := make([]*npool.AppUserSecretRes, 0, len(rows))
	for _, val := range rows {
		infos = append(infos, appUserSecretRowToObject(val))
	}

	return &npool.GetAppUserSecretsResponse{
		Infos: infos,
		Total: uint32(total),
	}, nil
}

func (s *Server) ExistAppUserSecretV2(ctx context.Context, in *npool.ExistAppUserSecretRequest) (*npool.ExistAppUserSecretResponse, error) {
	id, err := uuid.Parse(in.GetID())
	if err != nil {
		return &npool.ExistAppUserSecretResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	exist, err := crud.Exist(ctx, id)
	if err != nil {
		logger.Sugar().Errorf("fail check AppUserSecret: %v", err)
		return &npool.ExistAppUserSecretResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistAppUserSecretResponse{
		Info: exist,
	}, nil
}

func (s *Server) ExistAppUserSecretCondsV2(ctx context.Context, in *npool.ExistAppUserSecretCondsRequest) (*npool.ExistAppUserSecretCondsResponse, error) {
	exist, err := crud.ExistConds(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorf("fail check AppUserSecret: %v", err)
		return &npool.ExistAppUserSecretCondsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistAppUserSecretCondsResponse{
		Info: exist,
	}, nil
}

func (s *Server) CountAppUserSecretsV2(ctx context.Context, in *npool.CountAppUserSecretsRequest) (*npool.CountAppUserSecretsResponse, error) {
	total, err := crud.Count(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorf("fail count AppUserSecret: %v", err)
		return &npool.CountAppUserSecretsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CountAppUserSecretsResponse{
		Info: total,
	}, nil
}

func (s *Server) DeleteAppUserSecretV2(ctx context.Context, in *npool.DeleteAppUserSecretRequest) (*npool.DeleteAppUserSecretResponse, error) {
	id, err := uuid.Parse(in.GetID())
	if err != nil {
		return &npool.DeleteAppUserSecretResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := crud.Delete(ctx, id)
	if err != nil {
		logger.Sugar().Errorf("fail delete AppUserSecret: %v", err)
		return &npool.DeleteAppUserSecretResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.DeleteAppUserSecretResponse{
		Info: appUserSecretRowToObject(info),
	}, nil
}
