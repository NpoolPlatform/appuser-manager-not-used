//nolint:nolintlint,dupl
package appusersecret

import (
	"context"

	converter "github.com/NpoolPlatform/appuser-manager/pkg/converter/appusersecret"
	crud "github.com/NpoolPlatform/appuser-manager/pkg/crud/appusersecret"
	servicename "github.com/NpoolPlatform/appuser-manager/pkg/servicename"
	tracer "github.com/NpoolPlatform/appuser-manager/pkg/tracer/appusersecret"

	commontracer "github.com/NpoolPlatform/appuser-manager/pkg/tracer"

	"go.opentelemetry.io/otel"
	scodes "go.opentelemetry.io/otel/codes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/appuser/mgr/v2/appusersecret"

	"github.com/google/uuid"
)

func (s *Server) CreateAppUserSecret(ctx context.Context,
	in *npool.CreateAppUserSecretRequest) (*npool.CreateAppUserSecretResponse, error) {
	var err error

	_, span := otel.Tracer(servicename.ServiceDomain).Start(ctx, "CreateAppUserSecret")
	defer span.End()
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.Trace(span, in.GetInfo())

	err = validate(in.GetInfo())
	if err != nil {
		logger.Sugar().Errorw("CreateAppUserSecret", "error", err)
		return &npool.CreateAppUserSecretResponse{}, err
	}

	span = commontracer.TraceInvoker(span, "appusersecret", "crud", "Create")

	info, err := crud.Create(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorw("CreateAppUserSecret", "error", err)
		return &npool.CreateAppUserSecretResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateAppUserSecretResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) CreateAppUserSecrets(ctx context.Context,
	in *npool.CreateAppUserSecretsRequest) (*npool.CreateAppUserSecretsResponse, error) {
	var err error

	_, span := otel.Tracer(servicename.ServiceDomain).Start(ctx, "CreateAppUserSecrets")
	defer span.End()
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	if len(in.GetInfos()) == 0 {
		logger.Sugar().Errorw("CreateAppUserSecrets", "error", "Infos is empty")
		return &npool.CreateAppUserSecretsResponse{}, status.Error(codes.InvalidArgument, "Infos is empty")
	}

	span = tracer.TraceMany(span, in.GetInfos())
	span = commontracer.TraceInvoker(span, "appusersecret", "crud", "CreateBulk")

	rows, err := crud.CreateBulk(ctx, in.GetInfos())
	if err != nil {
		logger.Sugar().Errorw("CreateAppUserSecrets", "error", err)
		return &npool.CreateAppUserSecretsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateAppUserSecretsResponse{
		Infos: converter.Ent2GrpcMany(rows),
	}, nil
}

func (s *Server) UpdateAppUserSecret(ctx context.Context,
	in *npool.UpdateAppUserSecretRequest) (*npool.UpdateAppUserSecretResponse, error) {
	var err error

	_, span := otel.Tracer(servicename.ServiceDomain).Start(ctx, "UpdateAppUserSecret")
	defer span.End()
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.Trace(span, in.GetInfo())

	if _, err := uuid.Parse(in.GetInfo().GetID()); err != nil {
		logger.Sugar().Errorw("UpdateAppUserSecret", "ID", in.GetInfo().GetID(), "error", err)
		return &npool.UpdateAppUserSecretResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "appusersecret", "crud", "Update")

	info, err := crud.Update(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorw("UpdateAppUserSecret", "ID", in.GetInfo().GetID(), "error", err)
		return &npool.UpdateAppUserSecretResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.UpdateAppUserSecretResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetAppUserSecret(ctx context.Context, in *npool.GetAppUserSecretRequest) (*npool.GetAppUserSecretResponse, error) {
	var err error

	_, span := otel.Tracer(servicename.ServiceDomain).Start(ctx, "GetAppUserSecret")
	defer span.End()
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = commontracer.TraceID(span, in.GetID())

	id, err := uuid.Parse(in.GetID())
	if err != nil {
		logger.Sugar().Errorw("GetAppUserSecret", "ID", in.GetID(), "error", err)
		return &npool.GetAppUserSecretResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "appusersecret", "crud", "Row")

	info, err := crud.Row(ctx, id)
	if err != nil {
		logger.Sugar().Errorw("GetAppUserSecret", "ID", in.GetID(), "error", err)
		return &npool.GetAppUserSecretResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetAppUserSecretResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetAppUserSecretOnly(ctx context.Context,
	in *npool.GetAppUserSecretOnlyRequest) (*npool.GetAppUserSecretOnlyResponse, error) {
	var err error

	_, span := otel.Tracer(servicename.ServiceDomain).Start(ctx, "GetAppUserSecretOnly")
	defer span.End()
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceInvoker(span, "appusersecret", "crud", "RowOnly")

	info, err := crud.RowOnly(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorw("GetAppUserSecretOnly", "error", err)
		return &npool.GetAppUserSecretOnlyResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetAppUserSecretOnlyResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetAppUserSecrets(ctx context.Context, in *npool.GetAppUserSecretsRequest) (*npool.GetAppUserSecretsResponse, error) {
	var err error

	_, span := otel.Tracer(servicename.ServiceDomain).Start(ctx, "GetAppUserSecrets")
	defer span.End()
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceOffsetLimit(span, int(in.GetOffset()), int(in.GetLimit()))
	span = commontracer.TraceInvoker(span, "appusersecret", "crud", "Rows")

	rows, total, err := crud.Rows(ctx, in.GetConds(), int(in.GetOffset()), int(in.GetLimit()))
	if err != nil {
		logger.Sugar().Errorw("GetAppUserSecrets", "error", err)
		return &npool.GetAppUserSecretsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetAppUserSecretsResponse{
		Infos: converter.Ent2GrpcMany(rows),
		Total: uint32(total),
	}, nil
}

func (s *Server) ExistAppUserSecret(ctx context.Context, in *npool.ExistAppUserSecretRequest) (*npool.ExistAppUserSecretResponse, error) {
	var err error

	_, span := otel.Tracer(servicename.ServiceDomain).Start(ctx, "ExistAppUserSecret")
	defer span.End()
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = commontracer.TraceID(span, in.GetID())

	id, err := uuid.Parse(in.GetID())
	if err != nil {
		logger.Sugar().Errorw("ExistAppUserSecret", "ID", in.GetID(), "error", err)
		return &npool.ExistAppUserSecretResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "appusersecret", "crud", "Exist")

	exist, err := crud.Exist(ctx, id)
	if err != nil {
		logger.Sugar().Errorw("ExistAppUserSecret", "ID", in.GetID(), "error", err)
		return &npool.ExistAppUserSecretResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistAppUserSecretResponse{
		Info: exist,
	}, nil
}

func (s *Server) ExistAppUserSecretConds(ctx context.Context,
	in *npool.ExistAppUserSecretCondsRequest) (*npool.ExistAppUserSecretCondsResponse, error) {
	var err error

	_, span := otel.Tracer(servicename.ServiceDomain).Start(ctx, "ExistAppUserSecretConds")
	defer span.End()
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceInvoker(span, "appusersecret", "crud", "ExistAppUserSecretConds")

	exist, err := crud.ExistConds(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorw("ExistAppUserSecretConds", "error", err)
		return &npool.ExistAppUserSecretCondsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistAppUserSecretCondsResponse{
		Info: exist,
	}, nil
}

func (s *Server) CountAppUserSecrets(ctx context.Context,
	in *npool.CountAppUserSecretsRequest) (*npool.CountAppUserSecretsResponse, error) {
	var err error

	_, span := otel.Tracer(servicename.ServiceDomain).Start(ctx, "CountAppUserSecrets")
	defer span.End()
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceInvoker(span, "appusersecret", "crud", "Count")

	total, err := crud.Count(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorw("Counts", "error", err)
		return &npool.CountAppUserSecretsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CountAppUserSecretsResponse{
		Info: total,
	}, nil
}

func (s *Server) DeleteAppUserSecret(ctx context.Context,
	in *npool.DeleteAppUserSecretRequest) (*npool.DeleteAppUserSecretResponse, error) {
	var err error

	_, span := otel.Tracer(servicename.ServiceDomain).Start(ctx, "DeleteAppUserSecret")
	defer span.End()
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = commontracer.TraceID(span, in.GetID())

	id, err := uuid.Parse(in.GetID())
	if err != nil {
		logger.Sugar().Errorw("DeleteAppUserSecret", "ID", in.GetID(), "error", err)
		return &npool.DeleteAppUserSecretResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "appusersecret", "crud", "Delete")

	info, err := crud.Delete(ctx, id)
	if err != nil {
		logger.Sugar().Errorw("DeleteAppUserSecret", "ID", in.GetID(), "error", err)
		return &npool.DeleteAppUserSecretResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.DeleteAppUserSecretResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}
