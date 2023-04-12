package approle

import (
	"github.com/NpoolPlatform/message/npool/appuser/mgr/v2/approle"
	"google.golang.org/grpc"
)

type Server struct {
	approle.UnimplementedManagerServer
}

func Register(server grpc.ServiceRegistrar) {
	approle.RegisterManagerServer(server, &Server{})
}
