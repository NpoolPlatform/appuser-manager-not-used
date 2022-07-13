//go:build !codeanalysis
// +build !codeanalysis

package api

import (
	"context"
	"fmt"

	crud "github.com/NpoolPlatform/appuser-manager/pkg/crud/appuserthirdpartyv2"
	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent"
	constant "github.com/NpoolPlatform/appuser-manager/pkg/message/const"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	scodes "go.opentelemetry.io/otel/codes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/appusermgrv2/appuserthirdparty"

	"github.com/google/uuid"
)

func checkAppUserThirdPartyInfo(info *npool.AppUserThirdPartyReq) error {
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

func appUserThirdPartyRowToObject(row *ent.AppUserThirdParty) *npool.AppUserThirdParty {
	return &npool.AppUserThirdParty{
		UserID:               row.UserID.String(),
		ThirdPartyUserID:     row.ThirdPartyUserID,
		ThirdPartyID:         row.ThirdPartyID,
		ThirdPartyUsername:   row.ThirdPartyUsername,
		ThirdPartyUserAvatar: row.ThirdPartyUserAvatar,
		ID:                   row.ID.String(),
		AppID:                row.AppID.String(),
	}
}

func (s *AppUserThirdPartyServer) CreateAppUserThirdPartyV2(ctx context.Context, in *npool.CreateAppUserThirdPartyRequest) (*npool.CreateAppUserThirdPartyResponse, error) {
	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateAppUserThirdPartyV2")
	defer span.End()
	var err error
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()
	span = crud.AppUserThirdPartySpanAttributes(span, in.GetInfo())
	err = checkAppUserThirdPartyInfo(in.GetInfo())
	if err != nil {
		return &npool.CreateAppUserThirdPartyResponse{}, err
	}
	span.AddEvent("call crud Create")
	info, err := crud.Create(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorf("fail create AppUserThirdParty: %v", err)
		return &npool.CreateAppUserThirdPartyResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateAppUserThirdPartyResponse{
		Info: appUserThirdPartyRowToObject(info),
	}, nil
}

func (s *AppUserThirdPartyServer) CreateAppUserThirdPartysV2(ctx context.Context, in *npool.CreateAppUserThirdPartysRequest) (*npool.CreateAppUserThirdPartysResponse, error) {
	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateAppUserThirdPartysV2")
	defer span.End()
	var err error
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()
	if len(in.GetInfos()) == 0 {
		return &npool.CreateAppUserThirdPartysResponse{},
			status.Error(codes.InvalidArgument,
				"Batah create resource must more than 1",
			)
	}
	dupThirdPartyUserID := make(map[string]struct{})

	for key, info := range in.GetInfos() {
		span.SetAttributes(
			attribute.String("UserID"+fmt.Sprintf("%v", key), info.GetUserID()),
			attribute.String("ThirdPartyUserID"+fmt.Sprintf("%v", key), info.GetThirdPartyUserID()),
			attribute.String("ThirdPartyID"+fmt.Sprintf("%v", key), info.GetThirdPartyID()),
			attribute.String("ThirdPartyUsername"+fmt.Sprintf("%v", key), info.GetThirdPartyUsername()),
			attribute.String("ThirdPartyUserAvatar"+fmt.Sprintf("%v", key), info.GetThirdPartyUserAvatar()),
			attribute.String("ID"+fmt.Sprintf("%v", key), info.GetID()),
			attribute.String("AppID"+fmt.Sprintf("%v", key), info.GetAppID()),
		)
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

		dupThirdPartyUserID[info.GetThirdPartyUserID()] = struct{}{}
	}
	span.AddEvent("call crud CreateBulk")
	rows, err := crud.CreateBulk(ctx, in.GetInfos())
	if err != nil {
		logger.Sugar().Errorf("fail create AppUserThirdPartys: %v", err)
		return &npool.CreateAppUserThirdPartysResponse{}, status.Error(codes.Internal, err.Error())
	}

	infos := make([]*npool.AppUserThirdParty, 0, len(rows))
	for _, val := range rows {
		infos = append(infos, appUserThirdPartyRowToObject(val))
	}

	return &npool.CreateAppUserThirdPartysResponse{
		Infos: infos,
	}, nil
}

func (s *AppUserThirdPartyServer) UpdateAppUserThirdPartyV2(ctx context.Context, in *npool.UpdateAppUserThirdPartyRequest) (*npool.UpdateAppUserThirdPartyResponse, error) {
	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "UpdateAppUserThirdPartyV2")
	var err error
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()
	span = crud.AppUserThirdPartySpanAttributes(span, in.GetInfo())
	if _, err := uuid.Parse(in.GetInfo().GetID()); err != nil {
		logger.Sugar().Errorf("AppUserThirdParty id is invalid")
		return &npool.UpdateAppUserThirdPartyResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	span.AddEvent("call crud Update")
	info, err := crud.Update(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorf("fail update AppUserThirdParty: %v", err)
		return &npool.UpdateAppUserThirdPartyResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.UpdateAppUserThirdPartyResponse{
		Info: appUserThirdPartyRowToObject(info),
	}, nil
}

func (s *AppUserThirdPartyServer) GetAppUserThirdPartyV2(ctx context.Context, in *npool.GetAppUserThirdPartyRequest) (*npool.GetAppUserThirdPartyResponse, error) {
	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetAppUserThirdPartyV2")
	defer span.End()
	var err error
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()
	span.SetAttributes(
		attribute.String("ID", in.GetID()),
	)
	id, err := uuid.Parse(in.GetID())
	if err != nil {
		return &npool.GetAppUserThirdPartyResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	span.AddEvent("call crud Row")
	info, err := crud.Row(ctx, id)
	if err != nil {
		logger.Sugar().Errorf("fail get AppUserThirdParty: %v", err)
		return &npool.GetAppUserThirdPartyResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetAppUserThirdPartyResponse{
		Info: appUserThirdPartyRowToObject(info),
	}, nil
}

func (s *AppUserThirdPartyServer) GetAppUserThirdPartyOnlyV2(ctx context.Context, in *npool.GetAppUserThirdPartyOnlyRequest) (*npool.GetAppUserThirdPartyOnlyResponse, error) {
	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetAppUserThirdPartyOnlyV2")
	defer span.End()
	span = crud.AppUserThirdPartyCondsSpanAttributes(span, in.GetConds())
	span.AddEvent("call crud RowOnly")
	info, err := crud.RowOnly(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorf("fail get AppUserThirdPartys: %v", err)
		return &npool.GetAppUserThirdPartyOnlyResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetAppUserThirdPartyOnlyResponse{
		Info: appUserThirdPartyRowToObject(info),
	}, nil
}

func (s *AppUserThirdPartyServer) GetAppUserThirdPartysV2(ctx context.Context, in *npool.GetAppUserThirdPartysRequest) (*npool.GetAppUserThirdPartysResponse, error) {
	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetAppUserThirdPartysV2")
	defer span.End()
	var err error
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()
	span = crud.AppUserThirdPartyCondsSpanAttributes(span, in.GetConds())
	span.SetAttributes(
		attribute.Int("Offset", int(in.GetOffset())),
		attribute.Int("Limit", int(in.GetLimit())),
	)
	span.AddEvent("call crud Rows")
	rows, total, err := crud.Rows(ctx, in.GetConds(), int(in.GetOffset()), int(in.GetLimit()))
	if err != nil {
		logger.Sugar().Errorf("fail get AppUserThirdPartys: %v", err)
		return &npool.GetAppUserThirdPartysResponse{}, status.Error(codes.Internal, err.Error())
	}

	infos := make([]*npool.AppUserThirdParty, 0, len(rows))
	for _, val := range rows {
		infos = append(infos, appUserThirdPartyRowToObject(val))
	}

	return &npool.GetAppUserThirdPartysResponse{
		Infos: infos,
		Total: uint32(total),
	}, nil
}

func (s *AppUserThirdPartyServer) ExistAppUserThirdPartyV2(ctx context.Context, in *npool.ExistAppUserThirdPartyRequest) (*npool.ExistAppUserThirdPartyResponse, error) {
	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "ExistAppUserThirdPartyV2")
	defer span.End()
	var err error
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()
	span.SetAttributes(
		attribute.String("ID", in.GetID()),
	)
	id, err := uuid.Parse(in.GetID())
	if err != nil {
		return &npool.ExistAppUserThirdPartyResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	span.AddEvent("call crud Exist")
	exist, err := crud.Exist(ctx, id)
	if err != nil {
		logger.Sugar().Errorf("fail check AppUserThirdParty: %v", err)
		return &npool.ExistAppUserThirdPartyResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistAppUserThirdPartyResponse{
		Info: exist,
	}, nil
}

func (s *AppUserThirdPartyServer) ExistAppUserThirdPartyCondsV2(ctx context.Context, in *npool.ExistAppUserThirdPartyCondsRequest) (*npool.ExistAppUserThirdPartyCondsResponse, error) {
	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "ExistAppUserThirdPartyCondsV2")
	defer span.End()
	var err error
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()
	span = crud.AppUserThirdPartyCondsSpanAttributes(span, in.GetConds())
	span.AddEvent("call crud ExistConds")
	exist, err := crud.ExistConds(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorf("fail check AppUserThirdParty: %v", err)
		return &npool.ExistAppUserThirdPartyCondsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistAppUserThirdPartyCondsResponse{
		Info: exist,
	}, nil
}

func (s *AppUserThirdPartyServer) CountAppUserThirdPartysV2(ctx context.Context, in *npool.CountAppUserThirdPartysRequest) (*npool.CountAppUserThirdPartysResponse, error) {
	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CountAppUserThirdPartysV2")
	defer span.End()
	var err error
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()
	span = crud.AppUserThirdPartyCondsSpanAttributes(span, in.GetConds())
	span.AddEvent("call crud Count")
	total, err := crud.Count(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorf("fail count AppUserThirdParty: %v", err)
		return &npool.CountAppUserThirdPartysResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CountAppUserThirdPartysResponse{
		Info: total,
	}, nil
}

func (s *AppUserThirdPartyServer) DeleteAppUserThirdPartyV2(ctx context.Context, in *npool.DeleteAppUserThirdPartyRequest) (*npool.DeleteAppUserThirdPartyResponse, error) {
	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "DeleteAppUserThirdPartyV2")
	defer span.End()
	var err error
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()
	span.SetAttributes(
		attribute.String("ID", in.GetID()),
	)
	id, err := uuid.Parse(in.GetID())
	if err != nil {
		return &npool.DeleteAppUserThirdPartyResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	span.AddEvent("call crud Delete")
	info, err := crud.Delete(ctx, id)
	if err != nil {
		logger.Sugar().Errorf("fail delete AppUserThirdParty: %v", err)
		return &npool.DeleteAppUserThirdPartyResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.DeleteAppUserThirdPartyResponse{
		Info: appUserThirdPartyRowToObject(info),
	}, nil
}
