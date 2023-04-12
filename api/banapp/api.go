package banapp

import (
	"github.com/NpoolPlatform/message/npool/appuser/mgr/v2/banapp"
	"google.golang.org/grpc"
)

type Server struct {
	banapp.UnimplementedManagerServer
}

func Register(server grpc.ServiceRegistrar) {
	banapp.RegisterManagerServer(server, &Server{})
}
