package approleuser

import (
	"github.com/NpoolPlatform/message/npool/appuser/mgr/v2/approleuser"
	"google.golang.org/grpc"
)

type Server struct {
	approleuser.UnimplementedManagerServer
}

func Register(server grpc.ServiceRegistrar) {
	approleuser.RegisterManagerServer(server, &Server{})
}
