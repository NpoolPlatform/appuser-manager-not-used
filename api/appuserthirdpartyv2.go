//go:build !codeanalysis
// +build !codeanalysis

package api

import (
	"context"

	crud "github.com/NpoolPlatform/appuser-manager/pkg/crud/appuserthirdpartyv2"
	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/appusermgrv2/appuserthirdparty"

	"github.com/google/uuid"
)

func checkAppUserThirdPartyInfo(info *npool.AppUserThirdParty) error {
	if _, err := uuid.Parse(info.GetAppID()); err != nil {
		logger.Sugar().Error("AppID is invalid")
		return status.Error(codes.InvalidArgument, "AppID is invalid")
	}
	if _, err := uuid.Parse(info.GetUserID()); err != nil {
		logger.Sugar().Error("UserID is invalid")
		return status.Error(codes.InvalidArgument, "UserID is invalid")
	}
	if info.GetThirdPartyUserID() == "" {
		logger.Sugar().Error("ThirdPartyUserID is invalid")
		return status.Error(codes.InvalidArgument, "ThirdPartyUserID is invalid")
	}
	if info.GetThirdPartyID() == "" {
		logger.Sugar().Error("ThirdPartyID is invalid")
		return status.Error(codes.InvalidArgument, "ThirdPartyID is invalid")
	}
	return nil
}

func appUserThirdPartyRowToObject(row *ent.AppUserThirdParty) *npool.AppUserThirdPartyRes {
	return &npool.AppUserThirdPartyRes{
		UserID:               row.UserID.String(),
		ThirdPartyUserID:     row.ThirdPartyUserID,
		ThirdPartyID:         row.ThirdPartyID,
		ThirdPartyUsername:   row.ThirdPartyUsername,
		ThirdPartyUserAvatar: row.ThirdPartyUserAvatar,
		ID:                   row.ID.String(),
		AppID:                row.AppID.String(),
	}
}

func (s *Server) CreateAppUserThirdPartyV2(ctx context.Context, in *npool.CreateAppUserThirdPartyRequest) (*npool.CreateAppUserThirdPartyResponse, error) {
	err := checkAppUserThirdPartyInfo(in.GetInfo())
	if err != nil {
		return &npool.CreateAppUserThirdPartyResponse{}, err
	}

	info, err := crud.Create(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorf("fail create AppUserThirdParty: %v", err)
		return &npool.CreateAppUserThirdPartyResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateAppUserThirdPartyResponse{
		Info: appUserThirdPartyRowToObject(info),
	}, nil
}

func (s *Server) CreateAppUserThirdPartysV2(ctx context.Context, in *npool.CreateAppUserThirdPartysRequest) (*npool.CreateAppUserThirdPartysResponse, error) {
	if len(in.GetInfos()) == 0 {
		return &npool.CreateAppUserThirdPartysResponse{},
			status.Error(codes.InvalidArgument,
				"Batah create resource must more than 1",
			)
	}
	dupThirdPartyUserID := make(map[string]struct{})

	for _, info := range in.GetInfos() {
		err := checkAppUserThirdPartyInfo(info)
		if err != nil {
			return &npool.CreateAppUserThirdPartysResponse{}, err
		}
		if _, ok := dupThirdPartyUserID[info.GetThirdPartyUserID()]; ok {
			return &npool.CreateAppUserThirdPartysResponse{},
				status.Errorf(codes.AlreadyExists,
					"ThirdPartyUserID: %v duplicate create",
					info.GetThirdPartyUserID(),
				)
		}

	}

	rows, err := crud.CreateBulk(ctx, in.GetInfos())
	if err != nil {
		logger.Sugar().Errorf("fail create AppUserThirdPartys: %v", err)
		return &npool.CreateAppUserThirdPartysResponse{}, status.Error(codes.Internal, err.Error())
	}

	infos := make([]*npool.AppUserThirdPartyRes, 0, len(rows))
	for _, val := range rows {
		infos = append(infos, appUserThirdPartyRowToObject(val))
	}

	return &npool.CreateAppUserThirdPartysResponse{
		Infos: infos,
	}, nil
}

func (s *Server) UpdateAppUserThirdPartyV2(ctx context.Context, in *npool.UpdateAppUserThirdPartyRequest) (*npool.UpdateAppUserThirdPartyResponse, error) {
	if _, err := uuid.Parse(in.GetInfo().GetID()); err != nil {
		logger.Sugar().Errorf("AppUserThirdParty id is invalid")
		return &npool.UpdateAppUserThirdPartyResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := crud.Update(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorf("fail update AppUserThirdParty: %v", err)
		return &npool.UpdateAppUserThirdPartyResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.UpdateAppUserThirdPartyResponse{
		Info: appUserThirdPartyRowToObject(info),
	}, nil
}

func (s *Server) GetAppUserThirdPartyV2(ctx context.Context, in *npool.GetAppUserThirdPartyRequest) (*npool.GetAppUserThirdPartyResponse, error) {
	id, err := uuid.Parse(in.GetID())
	if err != nil {
		return &npool.GetAppUserThirdPartyResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := crud.Row(ctx, id)
	if err != nil {
		logger.Sugar().Errorf("fail get AppUserThirdParty: %v", err)
		return &npool.GetAppUserThirdPartyResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetAppUserThirdPartyResponse{
		Info: appUserThirdPartyRowToObject(info),
	}, nil
}

func (s *Server) GetAppUserThirdPartyOnlyV2(ctx context.Context, in *npool.GetAppUserThirdPartyOnlyRequest) (*npool.GetAppUserThirdPartyOnlyResponse, error) {
	info, err := crud.RowOnly(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorf("fail get AppUserThirdPartys: %v", err)
		return &npool.GetAppUserThirdPartyOnlyResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetAppUserThirdPartyOnlyResponse{
		Info: appUserThirdPartyRowToObject(info),
	}, nil
}

func (s *Server) GetAppUserThirdPartysV2(ctx context.Context, in *npool.GetAppUserThirdPartysRequest) (*npool.GetAppUserThirdPartysResponse, error) {
	rows, total, err := crud.Rows(ctx, in.GetConds(), int(in.GetOffset()), int(in.GetLimit()))
	if err != nil {
		logger.Sugar().Errorf("fail get AppUserThirdPartys: %v", err)
		return &npool.GetAppUserThirdPartysResponse{}, status.Error(codes.Internal, err.Error())
	}

	infos := make([]*npool.AppUserThirdPartyRes, 0, len(rows))
	for _, val := range rows {
		infos = append(infos, appUserThirdPartyRowToObject(val))
	}

	return &npool.GetAppUserThirdPartysResponse{
		Infos: infos,
		Total: uint32(total),
	}, nil
}

func (s *Server) ExistAppUserThirdPartyV2(ctx context.Context, in *npool.ExistAppUserThirdPartyRequest) (*npool.ExistAppUserThirdPartyResponse, error) {
	id, err := uuid.Parse(in.GetID())
	if err != nil {
		return &npool.ExistAppUserThirdPartyResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	exist, err := crud.Exist(ctx, id)
	if err != nil {
		logger.Sugar().Errorf("fail check AppUserThirdParty: %v", err)
		return &npool.ExistAppUserThirdPartyResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistAppUserThirdPartyResponse{
		Info: exist,
	}, nil
}

func (s *Server) ExistAppUserThirdPartyCondsV2(ctx context.Context, in *npool.ExistAppUserThirdPartyCondsRequest) (*npool.ExistAppUserThirdPartyCondsResponse, error) {
	exist, err := crud.ExistConds(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorf("fail check AppUserThirdParty: %v", err)
		return &npool.ExistAppUserThirdPartyCondsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistAppUserThirdPartyCondsResponse{
		Info: exist,
	}, nil
}

func (s *Server) CountAppUserThirdPartysV2(ctx context.Context, in *npool.CountAppUserThirdPartysRequest) (*npool.CountAppUserThirdPartysResponse, error) {
	total, err := crud.Count(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorf("fail count AppUserThirdParty: %v", err)
		return &npool.CountAppUserThirdPartysResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CountAppUserThirdPartysResponse{
		Info: total,
	}, nil
}

func (s *Server) DeleteAppUserThirdPartyV2(ctx context.Context, in *npool.DeleteAppUserThirdPartyRequest) (*npool.DeleteAppUserThirdPartyResponse, error) {
	id, err := uuid.Parse(in.GetID())
	if err != nil {
		return &npool.DeleteAppUserThirdPartyResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := crud.Delete(ctx, id)
	if err != nil {
		logger.Sugar().Errorf("fail delete AppUserThirdParty: %v", err)
		return &npool.DeleteAppUserThirdPartyResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.DeleteAppUserThirdPartyResponse{
		Info: appUserThirdPartyRowToObject(info),
	}, nil
}
