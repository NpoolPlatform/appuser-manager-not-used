package kyc

import (
	"github.com/NpoolPlatform/message/npool/appuser/mgr/v2/kyc"
	"google.golang.org/grpc"
)

type Server struct {
	kyc.UnimplementedManagerServer
}

func Register(server grpc.ServiceRegistrar) {
	kyc.RegisterManagerServer(server, &Server{})
}
