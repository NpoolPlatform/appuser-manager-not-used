package appuserthirdparty

import (
	"github.com/NpoolPlatform/message/npool/appuser/mgr/v2/appuserthirdparty"
	"google.golang.org/grpc"
)

type Server struct {
	appuserthirdparty.UnimplementedManagerServer
}

func Register(server grpc.ServiceRegistrar) {
	appuserthirdparty.RegisterManagerServer(server, &Server{})
}
