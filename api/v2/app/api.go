package app

import (
	"github.com/NpoolPlatform/message/npool/appuser/mgr/v2/app"
	"google.golang.org/grpc"
)

type Server struct {
	app.UnimplementedAppMgrServer
}

func Register(server grpc.ServiceRegistrar) {
	app.RegisterAppMgrServer(server, &Server{})
}
