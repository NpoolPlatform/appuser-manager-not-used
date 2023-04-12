//nolint:nolintlint,dupl
package kyc

import (
	"context"

	converter "github.com/NpoolPlatform/appuser-manager/pkg/converter/kyc"

	crud "github.com/NpoolPlatform/appuser-manager/pkg/crud/kyc"
	servicename "github.com/NpoolPlatform/appuser-manager/pkg/servicename"
	tracer "github.com/NpoolPlatform/appuser-manager/pkg/tracer/kyc"

	commontracer "github.com/NpoolPlatform/appuser-manager/pkg/tracer"

	"go.opentelemetry.io/otel"
	scodes "go.opentelemetry.io/otel/codes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/appuser/mgr/v2/kyc"

	"github.com/google/uuid"
)

func (s *Server) CreateKyc(ctx context.Context, in *npool.CreateKycRequest) (*npool.CreateKycResponse, error) {
	var err error

	_, span := otel.Tracer(servicename.ServiceDomain).Start(ctx, "CreateKyc")
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
		logger.Sugar().Errorw("CreateKyc", "error", err)
		return &npool.CreateKycResponse{}, err
	}

	span = commontracer.TraceInvoker(span, "kyc", "crud", "Create")

	info, err := crud.Create(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorw("CreateKyc", "error", err)
		return &npool.CreateKycResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateKycResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) CreateKycs(ctx context.Context, in *npool.CreateKycsRequest) (*npool.CreateKycsResponse, error) {
	var err error

	_, span := otel.Tracer(servicename.ServiceDomain).Start(ctx, "CreateKycs")
	defer span.End()
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	if len(in.GetInfos()) == 0 {
		logger.Sugar().Errorw("CreateKycs", "error", "Infos is empty")
		return &npool.CreateKycsResponse{}, status.Error(codes.InvalidArgument, "Infos is empty")
	}

	span = tracer.TraceMany(span, in.GetInfos())
	span = commontracer.TraceInvoker(span, "kyc", "crud", "CreateBulk")

	rows, err := crud.CreateBulk(ctx, in.GetInfos())
	if err != nil {
		logger.Sugar().Errorw("CreateKycs", "error", err)
		return &npool.CreateKycsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateKycsResponse{
		Infos: converter.Ent2GrpcMany(rows),
	}, nil
}

func (s *Server) UpdateKyc(ctx context.Context, in *npool.UpdateKycRequest) (*npool.UpdateKycResponse, error) {
	var err error

	_, span := otel.Tracer(servicename.ServiceDomain).Start(ctx, "UpdateKyc")
	defer span.End()
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.Trace(span, in.GetInfo())

	if _, err := uuid.Parse(in.GetInfo().GetID()); err != nil {
		logger.Sugar().Errorw("UpdateKyc", "ID", in.GetInfo().GetID(), "error", err)
		return &npool.UpdateKycResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "kyc", "crud", "Update")

	info, err := crud.Update(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorw("UpdateKyc", "ID", in.GetInfo().GetID(), "error", err)
		return &npool.UpdateKycResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.UpdateKycResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetKyc(ctx context.Context, in *npool.GetKycRequest) (*npool.GetKycResponse, error) {
	var err error

	_, span := otel.Tracer(servicename.ServiceDomain).Start(ctx, "GetKyc")
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
		logger.Sugar().Errorw("GetKyc", "ID", in.GetID(), "error", err)
		return &npool.GetKycResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "kyc", "crud", "Row")

	info, err := crud.Row(ctx, id)
	if err != nil {
		logger.Sugar().Errorw("GetKyc", "ID", in.GetID(), "error", err)
		return &npool.GetKycResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetKycResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetKycOnly(ctx context.Context, in *npool.GetKycOnlyRequest) (*npool.GetKycOnlyResponse, error) {
	var err error

	_, span := otel.Tracer(servicename.ServiceDomain).Start(ctx, "GetKycOnly")
	defer span.End()
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceInvoker(span, "kyc", "crud", "RowOnly")

	info, err := crud.RowOnly(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorw("GetKycOnly", "error", err)
		return &npool.GetKycOnlyResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetKycOnlyResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetKycs(ctx context.Context, in *npool.GetKycsRequest) (*npool.GetKycsResponse, error) {
	var err error

	_, span := otel.Tracer(servicename.ServiceDomain).Start(ctx, "GetKycs")
	defer span.End()
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceOffsetLimit(span, int(in.GetOffset()), int(in.GetLimit()))
	span = commontracer.TraceInvoker(span, "kyc", "crud", "Rows")

	rows, total, err := crud.Rows(ctx, in.GetConds(), int(in.GetOffset()), int(in.GetLimit()))
	if err != nil {
		logger.Sugar().Errorw("GetKycs", "error", err)
		return &npool.GetKycsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetKycsResponse{
		Infos: converter.Ent2GrpcMany(rows),
		Total: uint32(total),
	}, nil
}

func (s *Server) ExistKyc(ctx context.Context, in *npool.ExistKycRequest) (*npool.ExistKycResponse, error) {
	var err error

	_, span := otel.Tracer(servicename.ServiceDomain).Start(ctx, "ExistKyc")
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
		logger.Sugar().Errorw("ExistKyc", "ID", in.GetID(), "error", err)
		return &npool.ExistKycResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "kyc", "crud", "Exist")

	exist, err := crud.Exist(ctx, id)
	if err != nil {
		logger.Sugar().Errorw("ExistKyc", "ID", in.GetID(), "error", err)
		return &npool.ExistKycResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistKycResponse{
		Info: exist,
	}, nil
}

func (s *Server) ExistKycConds(ctx context.Context, in *npool.ExistKycCondsRequest) (*npool.ExistKycCondsResponse, error) {
	var err error

	_, span := otel.Tracer(servicename.ServiceDomain).Start(ctx, "ExistKycConds")
	defer span.End()
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceInvoker(span, "kyc", "crud", "ExistKycConds")

	exist, err := crud.ExistConds(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorw("ExistKycConds", "error", err)
		return &npool.ExistKycCondsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistKycCondsResponse{
		Info: exist,
	}, nil
}

func (s *Server) CountKycs(ctx context.Context, in *npool.CountKycsRequest) (*npool.CountKycsResponse, error) {
	var err error

	_, span := otel.Tracer(servicename.ServiceDomain).Start(ctx, "CountKycs")
	defer span.End()
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceInvoker(span, "kyc", "crud", "Count")

	total, err := crud.Count(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorw("Counts", "error", err)
		return &npool.CountKycsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CountKycsResponse{
		Info: total,
	}, nil
}

func (s *Server) DeleteKyc(ctx context.Context, in *npool.DeleteKycRequest) (*npool.DeleteKycResponse, error) {
	var err error

	_, span := otel.Tracer(servicename.ServiceDomain).Start(ctx, "DeleteKyc")
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
		logger.Sugar().Errorw("DeleteKyc", "ID", in.GetID(), "error", err)
		return &npool.DeleteKycResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "kyc", "crud", "Delete")

	info, err := crud.Delete(ctx, id)
	if err != nil {
		logger.Sugar().Errorw("DeleteKyc", "ID", in.GetID(), "error", err)
		return &npool.DeleteKycResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.DeleteKycResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}
