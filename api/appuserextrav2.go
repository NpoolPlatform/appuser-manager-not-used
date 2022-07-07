//go:build !codeanalysis
// +build !codeanalysis

package api

import (
	"context"

	crud "github.com/NpoolPlatform/appuser-manager/pkg/crud/appuserextrav2"
	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/appusermgrv2/appuserextra"

	"github.com/google/uuid"
)

func checkAppUserExtraInfo(info *npool.AppUserExtra) error {
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

func appUserExtraRowToObject(row *ent.AppUserExtra) *npool.AppUserExtraRes {
	return &npool.AppUserExtraRes{
		PostalCode:    row.PostalCode,
		Avatar:        row.Avatar,
		Organization:  row.Organization,
		Birthday:      row.Birthday,
		ID:            row.ID.String(),
		AppID:         row.AppID.String(),
		Username:      row.Username,
		Gender:        row.Gender,
		LastName:      row.LastName,
		Age:           row.Age,
		UserID:        row.UserID.String(),
		FirstName:     row.FirstName,
		IDNumber:      row.IDNumber,
		AddressFields: row.AddressFields,
	}
}

func (s *AppUserExtraServer) CreateAppUserExtraV2(ctx context.Context, in *npool.CreateAppUserExtraRequest) (*npool.CreateAppUserExtraResponse, error) {
	err := checkAppUserExtraInfo(in.GetInfo())
	if err != nil {
		return &npool.CreateAppUserExtraResponse{}, err
	}

	info, err := crud.Create(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorf("fail create AppUserExtra: %v", err)
		return &npool.CreateAppUserExtraResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateAppUserExtraResponse{
		Info: appUserExtraRowToObject(info),
	}, nil
}

func (s *AppUserExtraServer) CreateAppUserExtrasV2(ctx context.Context, in *npool.CreateAppUserExtrasRequest) (*npool.CreateAppUserExtrasResponse, error) {
	if len(in.GetInfos()) == 0 {
		return &npool.CreateAppUserExtrasResponse{},
			status.Error(codes.InvalidArgument,
				"Batah create resource must more than 1",
			)
	}
	dupIDNumber := make(map[string]struct{})

	for _, info := range in.GetInfos() {
		err := checkAppUserExtraInfo(info)
		if err != nil {
			return &npool.CreateAppUserExtrasResponse{}, err
		}
		if _, ok := dupIDNumber[info.GetIDNumber()]; ok {
			return &npool.CreateAppUserExtrasResponse{},
				status.Errorf(codes.AlreadyExists,
					"IDNumber: %v duplicate create",
					info.GetIDNumber(),
				)
		}

		dupIDNumber[info.GetIDNumber()] = struct{}{}
	}

	rows, err := crud.CreateBulk(ctx, in.GetInfos())
	if err != nil {
		logger.Sugar().Errorf("fail create AppUserExtras: %v", err)
		return &npool.CreateAppUserExtrasResponse{}, status.Error(codes.Internal, err.Error())
	}

	infos := make([]*npool.AppUserExtraRes, 0, len(rows))
	for _, val := range rows {
		infos = append(infos, appUserExtraRowToObject(val))
	}

	return &npool.CreateAppUserExtrasResponse{
		Infos: infos,
	}, nil
}

func (s *AppUserExtraServer) UpdateAppUserExtraV2(ctx context.Context, in *npool.UpdateAppUserExtraRequest) (*npool.UpdateAppUserExtraResponse, error) {
	if _, err := uuid.Parse(in.GetInfo().GetID()); err != nil {
		logger.Sugar().Errorf("AppUserExtra id is invalid")
		return &npool.UpdateAppUserExtraResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := crud.Update(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorf("fail update AppUserExtra: %v", err)
		return &npool.UpdateAppUserExtraResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.UpdateAppUserExtraResponse{
		Info: appUserExtraRowToObject(info),
	}, nil
}

func (s *AppUserExtraServer) GetAppUserExtraV2(ctx context.Context, in *npool.GetAppUserExtraRequest) (*npool.GetAppUserExtraResponse, error) {
	id, err := uuid.Parse(in.GetID())
	if err != nil {
		return &npool.GetAppUserExtraResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := crud.Row(ctx, id)
	if err != nil {
		logger.Sugar().Errorf("fail get AppUserExtra: %v", err)
		return &npool.GetAppUserExtraResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetAppUserExtraResponse{
		Info: appUserExtraRowToObject(info),
	}, nil
}

func (s *AppUserExtraServer) GetAppUserExtraOnlyV2(ctx context.Context, in *npool.GetAppUserExtraOnlyRequest) (*npool.GetAppUserExtraOnlyResponse, error) {
	info, err := crud.RowOnly(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorf("fail get AppUserExtras: %v", err)
		return &npool.GetAppUserExtraOnlyResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetAppUserExtraOnlyResponse{
		Info: appUserExtraRowToObject(info),
	}, nil
}

func (s *AppUserExtraServer) GetAppUserExtrasV2(ctx context.Context, in *npool.GetAppUserExtrasRequest) (*npool.GetAppUserExtrasResponse, error) {
	rows, total, err := crud.Rows(ctx, in.GetConds(), int(in.GetOffset()), int(in.GetLimit()))
	if err != nil {
		logger.Sugar().Errorf("fail get AppUserExtras: %v", err)
		return &npool.GetAppUserExtrasResponse{}, status.Error(codes.Internal, err.Error())
	}

	infos := make([]*npool.AppUserExtraRes, 0, len(rows))
	for _, val := range rows {
		infos = append(infos, appUserExtraRowToObject(val))
	}

	return &npool.GetAppUserExtrasResponse{
		Infos: infos,
		Total: uint32(total),
	}, nil
}

func (s *AppUserExtraServer) ExistAppUserExtraV2(ctx context.Context, in *npool.ExistAppUserExtraRequest) (*npool.ExistAppUserExtraResponse, error) {
	id, err := uuid.Parse(in.GetID())
	if err != nil {
		return &npool.ExistAppUserExtraResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	exist, err := crud.Exist(ctx, id)
	if err != nil {
		logger.Sugar().Errorf("fail check AppUserExtra: %v", err)
		return &npool.ExistAppUserExtraResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistAppUserExtraResponse{
		Info: exist,
	}, nil
}

func (s *AppUserExtraServer) ExistAppUserExtraCondsV2(ctx context.Context, in *npool.ExistAppUserExtraCondsRequest) (*npool.ExistAppUserExtraCondsResponse, error) {
	exist, err := crud.ExistConds(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorf("fail check AppUserExtra: %v", err)
		return &npool.ExistAppUserExtraCondsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistAppUserExtraCondsResponse{
		Info: exist,
	}, nil
}

func (s *AppUserExtraServer) CountAppUserExtrasV2(ctx context.Context, in *npool.CountAppUserExtrasRequest) (*npool.CountAppUserExtrasResponse, error) {
	total, err := crud.Count(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorf("fail count AppUserExtra: %v", err)
		return &npool.CountAppUserExtrasResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CountAppUserExtrasResponse{
		Info: total,
	}, nil
}

func (s *AppUserExtraServer) DeleteAppUserExtraV2(ctx context.Context, in *npool.DeleteAppUserExtraRequest) (*npool.DeleteAppUserExtraResponse, error) {
	id, err := uuid.Parse(in.GetID())
	if err != nil {
		return &npool.DeleteAppUserExtraResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := crud.Delete(ctx, id)
	if err != nil {
		logger.Sugar().Errorf("fail delete AppUserExtra: %v", err)
		return &npool.DeleteAppUserExtraResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.DeleteAppUserExtraResponse{
		Info: appUserExtraRowToObject(info),
	}, nil
}
