package appuser

import (
	"github.com/NpoolPlatform/message/npool/appuser/mgr/v2/appuser"
	"google.golang.org/grpc"
)

type Server struct {
	appuser.UnimplementedManagerServer
}

func Register(server grpc.ServiceRegistrar) {
	appuser.RegisterManagerServer(server, &Server{})
}
