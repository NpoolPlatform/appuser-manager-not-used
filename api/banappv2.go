//go:build !codeanalysis
// +build !codeanalysis

package api

import (
	"context"

	crud "github.com/NpoolPlatform/appuser-manager/pkg/crud/banappv2"
	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/appusermgrv2/banapp"

	"github.com/google/uuid"
)

func checkBanAppInfo(info *npool.BanAppReq) error {
	if _, err := uuid.Parse(info.GetAppID()); err != nil {
		logger.Sugar().Error("AppID is invalid")
		return status.Error(codes.InvalidArgument, "AppID is invalid")
	}
	return nil
}

func banAppRowToObject(row *ent.BanApp) *npool.BanApp {
	return &npool.BanApp{
		ID:      row.ID.String(),
		AppID:   row.AppID.String(),
		Message: row.Message,
	}
}

func (s *BanAppServer) CreateBanAppV2(ctx context.Context, in *npool.CreateBanAppRequest) (*npool.CreateBanAppResponse, error) {
	err := checkBanAppInfo(in.GetInfo())
	if err != nil {
		return &npool.CreateBanAppResponse{}, err
	}

	info, err := crud.Create(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorf("fail create BanApp: %v", err)
		return &npool.CreateBanAppResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateBanAppResponse{
		Info: banAppRowToObject(info),
	}, nil
}

func (s *BanAppServer) CreateBanAppsV2(ctx context.Context, in *npool.CreateBanAppsRequest) (*npool.CreateBanAppsResponse, error) {
	if len(in.GetInfos()) == 0 {
		return &npool.CreateBanAppsResponse{},
			status.Error(codes.InvalidArgument,
				"Batah create resource must more than 1",
			)
	}
	dupAppID := make(map[string]struct{})

	for _, info := range in.GetInfos() {
		err := checkBanAppInfo(info)
		if err != nil {
			return &npool.CreateBanAppsResponse{}, err
		}
		if _, ok := dupAppID[info.GetAppID()]; ok {
			return &npool.CreateBanAppsResponse{},
				status.Errorf(codes.AlreadyExists,
					"AppID: %v duplicate create",
					info.GetAppID(),
				)
		}
		dupAppID[info.GetAppID()] = struct{}{}
	}

	rows, err := crud.CreateBulk(ctx, in.GetInfos())
	if err != nil {
		logger.Sugar().Errorf("fail create BanApps: %v", err)
		return &npool.CreateBanAppsResponse{}, status.Error(codes.Internal, err.Error())
	}

	infos := make([]*npool.BanApp, 0, len(rows))
	for _, val := range rows {
		infos = append(infos, banAppRowToObject(val))
	}

	return &npool.CreateBanAppsResponse{
		Infos: infos,
	}, nil
}

func (s *BanAppServer) UpdateBanAppV2(ctx context.Context, in *npool.UpdateBanAppRequest) (*npool.UpdateBanAppResponse, error) {
	if _, err := uuid.Parse(in.GetInfo().GetID()); err != nil {
		logger.Sugar().Errorf("BanApp id is invalid")
		return &npool.UpdateBanAppResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := crud.Update(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorf("fail update BanApp: %v", err)
		return &npool.UpdateBanAppResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.UpdateBanAppResponse{
		Info: banAppRowToObject(info),
	}, nil
}

func (s *BanAppServer) GetBanAppV2(ctx context.Context, in *npool.GetBanAppRequest) (*npool.GetBanAppResponse, error) {
	id, err := uuid.Parse(in.GetID())
	if err != nil {
		return &npool.GetBanAppResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := crud.Row(ctx, id)
	if err != nil {
		logger.Sugar().Errorf("fail get BanApp: %v", err)
		return &npool.GetBanAppResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetBanAppResponse{
		Info: banAppRowToObject(info),
	}, nil
}

func (s *BanAppServer) GetBanAppOnlyV2(ctx context.Context, in *npool.GetBanAppOnlyRequest) (*npool.GetBanAppOnlyResponse, error) {
	info, err := crud.RowOnly(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorf("fail get BanApps: %v", err)
		return &npool.GetBanAppOnlyResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetBanAppOnlyResponse{
		Info: banAppRowToObject(info),
	}, nil
}

func (s *BanAppServer) GetBanAppsV2(ctx context.Context, in *npool.GetBanAppsRequest) (*npool.GetBanAppsResponse, error) {
	rows, total, err := crud.Rows(ctx, in.GetConds(), int(in.GetOffset()), int(in.GetLimit()))
	if err != nil {
		logger.Sugar().Errorf("fail get BanApps: %v", err)
		return &npool.GetBanAppsResponse{}, status.Error(codes.Internal, err.Error())
	}

	infos := make([]*npool.BanApp, 0, len(rows))
	for _, val := range rows {
		infos = append(infos, banAppRowToObject(val))
	}

	return &npool.GetBanAppsResponse{
		Infos: infos,
		Total: uint32(total),
	}, nil
}

func (s *BanAppServer) ExistBanAppV2(ctx context.Context, in *npool.ExistBanAppRequest) (*npool.ExistBanAppResponse, error) {
	id, err := uuid.Parse(in.GetID())
	if err != nil {
		return &npool.ExistBanAppResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	exist, err := crud.Exist(ctx, id)
	if err != nil {
		logger.Sugar().Errorf("fail check BanApp: %v", err)
		return &npool.ExistBanAppResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistBanAppResponse{
		Info: exist,
	}, nil
}

func (s *BanAppServer) ExistBanAppCondsV2(ctx context.Context, in *npool.ExistBanAppCondsRequest) (*npool.ExistBanAppCondsResponse, error) {
	exist, err := crud.ExistConds(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorf("fail check BanApp: %v", err)
		return &npool.ExistBanAppCondsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistBanAppCondsResponse{
		Info: exist,
	}, nil
}

func (s *BanAppServer) CountBanAppsV2(ctx context.Context, in *npool.CountBanAppsRequest) (*npool.CountBanAppsResponse, error) {
	total, err := crud.Count(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorf("fail count BanApp: %v", err)
		return &npool.CountBanAppsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CountBanAppsResponse{
		Info: total,
	}, nil
}

func (s *BanAppServer) DeleteBanAppV2(ctx context.Context, in *npool.DeleteBanAppRequest) (*npool.DeleteBanAppResponse, error) {
	id, err := uuid.Parse(in.GetID())
	if err != nil {
		return &npool.DeleteBanAppResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := crud.Delete(ctx, id)
	if err != nil {
		logger.Sugar().Errorf("fail delete BanApp: %v", err)
		return &npool.DeleteBanAppResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.DeleteBanAppResponse{
		Info: banAppRowToObject(info),
	}, nil
}
