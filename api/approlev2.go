//go:build !codeanalysis
// +build !codeanalysis

package api

import (
	"context"

	crud "github.com/NpoolPlatform/appuser-manager/pkg/crud/approlev2"
	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/appusermgrv2/approle"

	"github.com/google/uuid"
)

func checkAppRoleInfo(info *npool.AppRoleReq) error {
	if _, err := uuid.Parse(info.GetCreatedBy()); err != nil {
		logger.Sugar().Error("CreatedBy is invalid")
		return status.Error(codes.InvalidArgument, "CreatedBy is invalid")
	}

	if info.AppID == nil {
		logger.Sugar().Error("AppID is empty")
		return status.Error(codes.InvalidArgument, "AppID is empty")
	}

	if info.CreatedBy == nil {
		logger.Sugar().Error("CreatedBy is empty")
		return status.Error(codes.InvalidArgument, "CreatedBy is empty")
	}

	if info.Role == nil {
		logger.Sugar().Error("Role is empty")
		return status.Error(codes.InvalidArgument, "Role is empty")
	}

	return nil
}

func appRoleRowToObject(row *ent.AppRole) *npool.AppRole {
	return &npool.AppRole{
		ID:          row.ID.String(),
		AppID:       row.AppID.String(),
		CreatedBy:   row.CreatedBy.String(),
		Role:        row.Role,
		Description: row.Description,
		Default:     row.Default,
	}
}

func (s *AppRoleServer) CreateAppRoleV2(ctx context.Context, in *npool.CreateAppRoleRequest) (*npool.CreateAppRoleResponse, error) {
	err := checkAppRoleInfo(in.GetInfo())
	if err != nil {
		return &npool.CreateAppRoleResponse{}, err
	}

	info, err := crud.Create(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorf("fail create app role: %v", err)
		return &npool.CreateAppRoleResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateAppRoleResponse{
		Info: appRoleRowToObject(info),
	}, nil
}

func (s *AppRoleServer) CreateAppRolesV2(ctx context.Context, in *npool.CreateAppRolesRequest) (*npool.CreateAppRolesResponse, error) {
	if len(in.GetInfos()) == 0 {
		return &npool.CreateAppRolesResponse{},
			status.Error(codes.InvalidArgument,
				"Batah create resource must more than 1",
			)
	}

	rows, err := crud.CreateBulk(ctx, in.GetInfos())
	if err != nil {
		logger.Sugar().Errorf("fail create app roles: %v", err)
		return &npool.CreateAppRolesResponse{}, status.Error(codes.Internal, err.Error())
	}

	infos := make([]*npool.AppRole, 0, len(rows))
	for _, val := range rows {
		infos = append(infos, appRoleRowToObject(val))
	}

	return &npool.CreateAppRolesResponse{
		Infos: infos,
	}, nil
}

func (s *AppRoleServer) UpdateAppRoleV2(ctx context.Context, in *npool.UpdateAppRoleRequest) (*npool.UpdateAppRoleResponse, error) {
	if _, err := uuid.Parse(in.GetInfo().GetID()); err != nil {
		logger.Sugar().Errorf("app role id is invalid")
		return &npool.UpdateAppRoleResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := crud.Update(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorf("fail update app role: %v", err)
		return &npool.UpdateAppRoleResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.UpdateAppRoleResponse{
		Info: appRoleRowToObject(info),
	}, nil
}

func (s *AppRoleServer) GetAppRoleV2(ctx context.Context, in *npool.GetAppRoleRequest) (*npool.GetAppRoleResponse, error) {
	id, err := uuid.Parse(in.GetID())
	if err != nil {
		return &npool.GetAppRoleResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := crud.Row(ctx, id)
	if err != nil {
		logger.Sugar().Errorf("fail get app role: %v", err)
		return &npool.GetAppRoleResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetAppRoleResponse{
		Info: appRoleRowToObject(info),
	}, nil
}

func (s *AppRoleServer) GetAppRoleOnlyV2(ctx context.Context, in *npool.GetAppRoleOnlyRequest) (*npool.GetAppRoleOnlyResponse, error) {
	info, err := crud.RowOnly(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorf("fail get app roles: %v", err)
		return &npool.GetAppRoleOnlyResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetAppRoleOnlyResponse{
		Info: appRoleRowToObject(info),
	}, nil
}

func (s *AppRoleServer) GetAppRolesV2(ctx context.Context, in *npool.GetAppRolesRequest) (*npool.GetAppRolesResponse, error) {
	rows, total, err := crud.Rows(ctx, in.GetConds(), int(in.GetOffset()), int(in.GetLimit()))
	if err != nil {
		logger.Sugar().Errorf("fail get app roles: %v", err)
		return &npool.GetAppRolesResponse{}, status.Error(codes.Internal, err.Error())
	}

	infos := make([]*npool.AppRole, 0, len(rows))
	for _, val := range rows {
		infos = append(infos, appRoleRowToObject(val))
	}

	return &npool.GetAppRolesResponse{
		Infos: infos,
		Total: uint32(total),
	}, nil
}

func (s *AppRoleServer) ExistAppRoleV2(ctx context.Context, in *npool.ExistAppRoleRequest) (*npool.ExistAppRoleResponse, error) {
	id, err := uuid.Parse(in.GetID())
	if err != nil {
		return &npool.ExistAppRoleResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	exist, err := crud.Exist(ctx, id)
	if err != nil {
		logger.Sugar().Errorf("fail check app role: %v", err)
		return &npool.ExistAppRoleResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistAppRoleResponse{
		Info: exist,
	}, nil
}

func (s *AppRoleServer) ExistAppRoleCondsV2(ctx context.Context, in *npool.ExistAppRoleCondsRequest) (*npool.ExistAppRoleCondsResponse, error) {
	exist, err := crud.ExistConds(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorf("fail check app role: %v", err)
		return &npool.ExistAppRoleCondsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistAppRoleCondsResponse{
		Info: exist,
	}, nil
}

func (s *AppRoleServer) CountAppRolesV2(ctx context.Context, in *npool.CountAppRolesRequest) (*npool.CountAppRolesResponse, error) {
	total, err := crud.Count(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorf("fail count app role : %v", err)
		return &npool.CountAppRolesResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CountAppRolesResponse{
		Info: total,
	}, nil
}

func (s *AppRoleServer) DeleteAppRoleV2(ctx context.Context, in *npool.DeleteAppRoleRequest) (*npool.DeleteAppRoleResponse, error) {
	id, err := uuid.Parse(in.GetID())
	if err != nil {
		return &npool.DeleteAppRoleResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := crud.Delete(ctx, id)
	if err != nil {
		logger.Sugar().Errorf("fail delete app role: %v", err)
		return &npool.DeleteAppRoleResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.DeleteAppRoleResponse{
		Info: appRoleRowToObject(info),
	}, nil
}
