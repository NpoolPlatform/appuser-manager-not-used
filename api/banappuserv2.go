//go:build !codeanalysis
// +build !codeanalysis

package api

import (
	"context"

	crud "github.com/NpoolPlatform/appuser-manager/pkg/crud/banappuserv2"
	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/appusermgrv2/banappuser"

	"github.com/google/uuid"
)

func checkBanAppUserInfo(info *npool.BanAppUser) error {
	if _, err := uuid.Parse(info.GetAppID()); err != nil {
		logger.Sugar().Error("AppID is invalid")
		return status.Error(codes.InvalidArgument, "AppID is invalid")
	}
	if _, err := uuid.Parse(info.GetUserID()); err != nil {
		logger.Sugar().Error("UserID is invalid")
		return status.Error(codes.InvalidArgument, "UserID is invalid")
	}
	return nil
}

func banAppUserRowToObject(row *ent.BanAppUser) *npool.BanAppUserRes {
	return &npool.BanAppUserRes{
		UserID:  row.UserID.String(),
		Message: row.Message,
		ID:      row.ID.String(),
		AppID:   row.AppID.String(),
	}
}

func (s *Server) CreateBanAppUserV2(ctx context.Context, in *npool.CreateBanAppUserRequest) (*npool.CreateBanAppUserResponse, error) {
	err := checkBanAppUserInfo(in.GetInfo())
	if err != nil {
		return &npool.CreateBanAppUserResponse{}, err
	}

	info, err := crud.Create(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorf("fail create BanAppUser: %v", err)
		return &npool.CreateBanAppUserResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateBanAppUserResponse{
		Info: banAppUserRowToObject(info),
	}, nil
}

func (s *Server) CreateBanAppUsersV2(ctx context.Context, in *npool.CreateBanAppUsersRequest) (*npool.CreateBanAppUsersResponse, error) {
	if len(in.GetInfos()) == 0 {
		return &npool.CreateBanAppUsersResponse{},
			status.Error(codes.InvalidArgument,
				"Batah create resource must more than 1",
			)
	}
	dupUserID := make(map[string]struct{})

	for _, info := range in.GetInfos() {
		err := checkBanAppUserInfo(info)
		if err != nil {
			return &npool.CreateBanAppUsersResponse{}, err
		}
		if _, ok := dupUserID[info.GetUserID()]; ok {
			return &npool.CreateBanAppUsersResponse{},
				status.Errorf(codes.AlreadyExists,
					"UserID: %v duplicate create",
					info.GetUserID(),
				)
		}

	}

	rows, err := crud.CreateBulk(ctx, in.GetInfos())
	if err != nil {
		logger.Sugar().Errorf("fail create BanAppUsers: %v", err)
		return &npool.CreateBanAppUsersResponse{}, status.Error(codes.Internal, err.Error())
	}

	infos := make([]*npool.BanAppUserRes, 0, len(rows))
	for _, val := range rows {
		infos = append(infos, banAppUserRowToObject(val))
	}

	return &npool.CreateBanAppUsersResponse{
		Infos: infos,
	}, nil
}

func (s *Server) UpdateBanAppUserV2(ctx context.Context, in *npool.UpdateBanAppUserRequest) (*npool.UpdateBanAppUserResponse, error) {
	if _, err := uuid.Parse(in.GetInfo().GetID()); err != nil {
		logger.Sugar().Errorf("BanAppUser id is invalid")
		return &npool.UpdateBanAppUserResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := crud.Update(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorf("fail update BanAppUser: %v", err)
		return &npool.UpdateBanAppUserResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.UpdateBanAppUserResponse{
		Info: banAppUserRowToObject(info),
	}, nil
}

func (s *Server) GetBanAppUserV2(ctx context.Context, in *npool.GetBanAppUserRequest) (*npool.GetBanAppUserResponse, error) {
	id, err := uuid.Parse(in.GetID())
	if err != nil {
		return &npool.GetBanAppUserResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := crud.Row(ctx, id)
	if err != nil {
		logger.Sugar().Errorf("fail get BanAppUser: %v", err)
		return &npool.GetBanAppUserResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetBanAppUserResponse{
		Info: banAppUserRowToObject(info),
	}, nil
}

func (s *Server) GetBanAppUserOnlyV2(ctx context.Context, in *npool.GetBanAppUserOnlyRequest) (*npool.GetBanAppUserOnlyResponse, error) {
	info, err := crud.RowOnly(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorf("fail get BanAppUsers: %v", err)
		return &npool.GetBanAppUserOnlyResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetBanAppUserOnlyResponse{
		Info: banAppUserRowToObject(info),
	}, nil
}

func (s *Server) GetBanAppUsersV2(ctx context.Context, in *npool.GetBanAppUsersRequest) (*npool.GetBanAppUsersResponse, error) {
	rows, total, err := crud.Rows(ctx, in.GetConds(), int(in.GetOffset()), int(in.GetLimit()))
	if err != nil {
		logger.Sugar().Errorf("fail get BanAppUsers: %v", err)
		return &npool.GetBanAppUsersResponse{}, status.Error(codes.Internal, err.Error())
	}

	infos := make([]*npool.BanAppUserRes, 0, len(rows))
	for _, val := range rows {
		infos = append(infos, banAppUserRowToObject(val))
	}

	return &npool.GetBanAppUsersResponse{
		Infos: infos,
		Total: uint32(total),
	}, nil
}

func (s *Server) ExistBanAppUserV2(ctx context.Context, in *npool.ExistBanAppUserRequest) (*npool.ExistBanAppUserResponse, error) {
	id, err := uuid.Parse(in.GetID())
	if err != nil {
		return &npool.ExistBanAppUserResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	exist, err := crud.Exist(ctx, id)
	if err != nil {
		logger.Sugar().Errorf("fail check BanAppUser: %v", err)
		return &npool.ExistBanAppUserResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistBanAppUserResponse{
		Info: exist,
	}, nil
}

func (s *Server) ExistBanAppUserCondsV2(ctx context.Context, in *npool.ExistBanAppUserCondsRequest) (*npool.ExistBanAppUserCondsResponse, error) {
	exist, err := crud.ExistConds(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorf("fail check BanAppUser: %v", err)
		return &npool.ExistBanAppUserCondsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistBanAppUserCondsResponse{
		Info: exist,
	}, nil
}

func (s *Server) CountBanAppUsersV2(ctx context.Context, in *npool.CountBanAppUsersRequest) (*npool.CountBanAppUsersResponse, error) {
	total, err := crud.Count(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorf("fail count BanAppUser: %v", err)
		return &npool.CountBanAppUsersResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CountBanAppUsersResponse{
		Info: total,
	}, nil
}

func (s *Server) DeleteBanAppUserV2(ctx context.Context, in *npool.DeleteBanAppUserRequest) (*npool.DeleteBanAppUserResponse, error) {
	id, err := uuid.Parse(in.GetID())
	if err != nil {
		return &npool.DeleteBanAppUserResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := crud.Delete(ctx, id)
	if err != nil {
		logger.Sugar().Errorf("fail delete BanAppUser: %v", err)
		return &npool.DeleteBanAppUserResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.DeleteBanAppUserResponse{
		Info: banAppUserRowToObject(info),
	}, nil
}
