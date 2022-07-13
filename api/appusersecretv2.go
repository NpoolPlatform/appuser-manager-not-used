//go:build !codeanalysis
// +build !codeanalysis

package api

import (
	"context"
	"fmt"

	crud "github.com/NpoolPlatform/appuser-manager/pkg/crud/appusersecretv2"
	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent"
	constant "github.com/NpoolPlatform/appuser-manager/pkg/message/const"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	scodes "go.opentelemetry.io/otel/codes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/appusermgrv2/appusersecret"

	"github.com/google/uuid"
)

func checkAppUserSecretInfo(info *npool.AppUserSecretReq) error {
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

func appUserSecretRowToObject(row *ent.AppUserSecret) *npool.AppUserSecret {
	return &npool.AppUserSecret{
		Salt:         row.Salt,
		GoogleSecret: row.GoogleSecret,
		ID:           row.ID.String(),
		AppID:        row.AppID.String(),
		UserID:       row.UserID.String(),
		PasswordHash: row.PasswordHash,
	}
}

func (s *AppUserSecretServer) CreateAppUserSecretV2(ctx context.Context, in *npool.CreateAppUserSecretRequest) (*npool.CreateAppUserSecretResponse, error) {
	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateAppUserSecretV2")
	defer span.End()
	var err error
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()
	span = crud.AppUserSecretSpanAttributes(span, in.GetInfo())
	err = checkAppUserSecretInfo(in.GetInfo())
	if err != nil {
		return &npool.CreateAppUserSecretResponse{}, err
	}
	span.AddEvent("call crud Create")
	info, err := crud.Create(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorf("fail create AppUserSecret: %v", err)
		return &npool.CreateAppUserSecretResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateAppUserSecretResponse{
		Info: appUserSecretRowToObject(info),
	}, nil
}

func (s *AppUserSecretServer) CreateAppUserSecretsV2(ctx context.Context, in *npool.CreateAppUserSecretsRequest) (*npool.CreateAppUserSecretsResponse, error) {
	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateAppUserSecretsV2")
	defer span.End()
	var err error
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()
	if len(in.GetInfos()) == 0 {
		return &npool.CreateAppUserSecretsResponse{},
			status.Error(codes.InvalidArgument,
				"Batah create resource must more than 1",
			)
	}
	for key, info := range in.GetInfos() {
		span.SetAttributes(
			attribute.String("Salt"+fmt.Sprintf("%v", key), info.GetSalt()),
			attribute.String("GoogleSecret"+fmt.Sprintf("%v", key), info.GetGoogleSecret()),
			attribute.String("ID"+fmt.Sprintf("%v", key), info.GetID()),
			attribute.String("AppID"+fmt.Sprintf("%v", key), info.GetAppID()),
			attribute.String("UserID"+fmt.Sprintf("%v", key), info.GetUserID()),
			attribute.String("PasswordHash"+fmt.Sprintf("%v", key), info.GetPasswordHash()),
		)
		err = checkAppUserSecretInfo(info)
		if err != nil {
			return &npool.CreateAppUserSecretsResponse{}, err
		}
	}
	span.AddEvent("call crud CreateBulk")
	rows, err := crud.CreateBulk(ctx, in.GetInfos())
	if err != nil {
		logger.Sugar().Errorf("fail create AppUserSecrets: %v", err)
		return &npool.CreateAppUserSecretsResponse{}, status.Error(codes.Internal, err.Error())
	}

	infos := make([]*npool.AppUserSecret, 0, len(rows))
	for _, val := range rows {
		infos = append(infos, appUserSecretRowToObject(val))
	}

	return &npool.CreateAppUserSecretsResponse{
		Infos: infos,
	}, nil
}

func (s *AppUserSecretServer) UpdateAppUserSecretV2(ctx context.Context, in *npool.UpdateAppUserSecretRequest) (*npool.UpdateAppUserSecretResponse, error) {
	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "UpdateAppUserSecretV2")
	defer span.End()
	var err error
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()
	span = crud.AppUserSecretSpanAttributes(span, in.GetInfo())
	if _, err := uuid.Parse(in.GetInfo().GetID()); err != nil {
		logger.Sugar().Errorf("AppUserSecret id is invalid")
		return &npool.UpdateAppUserSecretResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	span.AddEvent("call crud Update")
	info, err := crud.Update(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorf("fail update AppUserSecret: %v", err)
		return &npool.UpdateAppUserSecretResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.UpdateAppUserSecretResponse{
		Info: appUserSecretRowToObject(info),
	}, nil
}

func (s *AppUserSecretServer) GetAppUserSecretV2(ctx context.Context, in *npool.GetAppUserSecretRequest) (*npool.GetAppUserSecretResponse, error) {
	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetAppUserSecretV2")
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
		return &npool.GetAppUserSecretResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	span.AddEvent("call crud Row")
	info, err := crud.Row(ctx, id)
	if err != nil {
		logger.Sugar().Errorf("fail get AppUserSecret: %v", err)
		return &npool.GetAppUserSecretResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetAppUserSecretResponse{
		Info: appUserSecretRowToObject(info),
	}, nil
}

func (s *AppUserSecretServer) GetAppUserSecretOnlyV2(ctx context.Context, in *npool.GetAppUserSecretOnlyRequest) (*npool.GetAppUserSecretOnlyResponse, error) {
	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetAppUserSecretOnlyV2")
	defer span.End()
	var err error
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()
	span = crud.AppUserSecretCondsSpanAttributes(span, in.GetConds())
	span.AddEvent("call crud RowOnly")
	info, err := crud.RowOnly(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorf("fail get AppUserSecrets: %v", err)
		return &npool.GetAppUserSecretOnlyResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetAppUserSecretOnlyResponse{
		Info: appUserSecretRowToObject(info),
	}, nil
}

func (s *AppUserSecretServer) GetAppUserSecretsV2(ctx context.Context, in *npool.GetAppUserSecretsRequest) (*npool.GetAppUserSecretsResponse, error) {
	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetAppUserSecretsV2")
	defer span.End()
	var err error
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()
	span = crud.AppUserSecretCondsSpanAttributes(span, in.GetConds())
	span.SetAttributes(
		attribute.Int("Offset", int(in.GetOffset())),
		attribute.Int("Limit", int(in.GetLimit())),
	)
	span.AddEvent("call crud Rows")
	rows, total, err := crud.Rows(ctx, in.GetConds(), int(in.GetOffset()), int(in.GetLimit()))
	if err != nil {
		logger.Sugar().Errorf("fail get AppUserSecrets: %v", err)
		return &npool.GetAppUserSecretsResponse{}, status.Error(codes.Internal, err.Error())
	}

	infos := make([]*npool.AppUserSecret, 0, len(rows))
	for _, val := range rows {
		infos = append(infos, appUserSecretRowToObject(val))
	}

	return &npool.GetAppUserSecretsResponse{
		Infos: infos,
		Total: uint32(total),
	}, nil
}

func (s *AppUserSecretServer) ExistAppUserSecretV2(ctx context.Context, in *npool.ExistAppUserSecretRequest) (*npool.ExistAppUserSecretResponse, error) {
	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "ExistAppUserSecretV2")
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
		return &npool.ExistAppUserSecretResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	span.AddEvent("call crud Exist")
	exist, err := crud.Exist(ctx, id)
	if err != nil {
		logger.Sugar().Errorf("fail check AppUserSecret: %v", err)
		return &npool.ExistAppUserSecretResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistAppUserSecretResponse{
		Info: exist,
	}, nil
}

func (s *AppUserSecretServer) ExistAppUserSecretCondsV2(ctx context.Context, in *npool.ExistAppUserSecretCondsRequest) (*npool.ExistAppUserSecretCondsResponse, error) {
	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "ExistAppUserSecretCondsV2")
	defer span.End()
	var err error
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()
	span = crud.AppUserSecretCondsSpanAttributes(span, in.GetConds())
	span.AddEvent("call crud ExistConds")
	exist, err := crud.ExistConds(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorf("fail check AppUserSecret: %v", err)
		return &npool.ExistAppUserSecretCondsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistAppUserSecretCondsResponse{
		Info: exist,
	}, nil
}

func (s *AppUserSecretServer) CountAppUserSecretsV2(ctx context.Context, in *npool.CountAppUserSecretsRequest) (*npool.CountAppUserSecretsResponse, error) {
	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CountAppUserSecretsV2")
	defer span.End()
	var err error
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()
	span = crud.AppUserSecretCondsSpanAttributes(span, in.GetConds())
	span.AddEvent("call crud Count")
	total, err := crud.Count(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorf("fail count AppUserSecret: %v", err)
		return &npool.CountAppUserSecretsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CountAppUserSecretsResponse{
		Info: total,
	}, nil
}

func (s *AppUserSecretServer) DeleteAppUserSecretV2(ctx context.Context, in *npool.DeleteAppUserSecretRequest) (*npool.DeleteAppUserSecretResponse, error) {
	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "DeleteAppUserSecretV2")
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
		return &npool.DeleteAppUserSecretResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	span.AddEvent("call crud Delete")
	info, err := crud.Delete(ctx, id)
	if err != nil {
		logger.Sugar().Errorf("fail delete AppUserSecret: %v", err)
		return &npool.DeleteAppUserSecretResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.DeleteAppUserSecretResponse{
		Info: appUserSecretRowToObject(info),
	}, nil
}
