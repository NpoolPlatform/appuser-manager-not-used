package appcontrol

import (
	"github.com/NpoolPlatform/message/npool/appuser/mgr/v2/appcontrol"
	"google.golang.org/grpc"
)

type Server struct {
	appcontrol.UnimplementedManagerServer
}

func Register(server grpc.ServiceRegistrar) {
	appcontrol.RegisterManagerServer(server, &Server{})
}
