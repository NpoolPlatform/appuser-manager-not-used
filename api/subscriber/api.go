package subscriber

import (
	"github.com/NpoolPlatform/message/npool/appuser/mgr/v2/subscriber"
	"google.golang.org/grpc"
)

type Server struct {
	subscriber.UnimplementedManagerServer
}

func Register(server grpc.ServiceRegistrar) {
	subscriber.RegisterManagerServer(server, &Server{})
}
