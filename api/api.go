package api

import (
	"context"
	"github.com/NpoolPlatform/appuser-manager/api/v2/app"
	appusermgr "github.com/NpoolPlatform/message/npool/appuser/mgr/v2"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	appusermgr.UnimplementedManagerServer
}

func Register(server grpc.ServiceRegistrar) {
	appusermgr.RegisterManagerServer(server, &Server{})
	app.Register(server)
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	if err := appusermgr.RegisterManagerHandlerFromEndpoint(context.Background(), mux, endpoint, opts); err != nil {
		return err
	}
	return nil
}
