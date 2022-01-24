package api

import (
	"context"

	npool "github.com/NpoolPlatform/message/npool/appusermgr"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	npool.UnimplementedAppUserManagerServer
}

func Register(server grpc.ServiceRegistrar) {
	npool.RegisterAppUserManagerServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return npool.RegisterAppUserManagerHandlerFromEndpoint(context.Background(), mux, endpoint, opts)
}
