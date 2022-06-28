package api

import (
	"context"

	npool "github.com/NpoolPlatform/message/npool/appusermgr"
	"github.com/NpoolPlatform/message/npool/appusermgrv2/app"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	npool.UnimplementedAppUserManagerServer
}

type AppService struct {
	app.UnimplementedAppUserManagerAppServer
}

func Register(server grpc.ServiceRegistrar) {
	npool.RegisterAppUserManagerServer(server, &Server{})
	app.RegisterAppUserManagerAppServer(server, &AppService{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return npool.RegisterAppUserManagerHandlerFromEndpoint(context.Background(), mux, endpoint, opts)
}
