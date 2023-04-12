package banappuser

import (
	"github.com/NpoolPlatform/message/npool/appuser/mgr/v2/banappuser"
	"google.golang.org/grpc"
)

type Server struct {
	banappuser.UnimplementedManagerServer
}

func Register(server grpc.ServiceRegistrar) {
	banappuser.RegisterManagerServer(server, &Server{})
}
