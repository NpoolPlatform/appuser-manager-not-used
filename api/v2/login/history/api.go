package history

import (
	"github.com/NpoolPlatform/message/npool/appuser/mgr/v2/login/history"
	"google.golang.org/grpc"
)

type Server struct {
	history.UnimplementedManagerServer
}

func Register(server grpc.ServiceRegistrar) {
	history.RegisterManagerServer(server, &Server{})
}
