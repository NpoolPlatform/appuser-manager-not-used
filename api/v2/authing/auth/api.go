package auth

import (
	"github.com/NpoolPlatform/message/npool/appuser/mgr/v2/authing/auth"
	"google.golang.org/grpc"
)

type Server struct {
	auth.UnimplementedManagerServer
}

func Register(server grpc.ServiceRegistrar) {
	auth.RegisterManagerServer(server, &Server{})
}
