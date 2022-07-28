package api

import (
	appusermgrv1 "github.com/NpoolPlatform/appuser-manager/api/v1"
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
	appusermgrv1.Register(server)
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	if err := appusermgrv1.RegisterGateway(mux, endpoint, opts); err != nil {
		return err
	}
	return nil
}
