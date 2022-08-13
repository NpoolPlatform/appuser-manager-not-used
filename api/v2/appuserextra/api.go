package appuserextra

import (
	"github.com/NpoolPlatform/message/npool/appuser/mgr/v2/appuserextra"
	"google.golang.org/grpc"
)

type Server struct {
	appuserextra.UnimplementedManagerServer
}

func Register(server grpc.ServiceRegistrar) {
	appuserextra.RegisterManagerServer(server, &Server{})
}
