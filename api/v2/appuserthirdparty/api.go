package appuserthirdparty

import (
	"github.com/NpoolPlatform/message/npool/appuser/mgr/v2/banappuser"
	"google.golang.org/grpc"
)

type Server struct {
	banappuser.UnimplementedBanAppUserMgrServer
}

func Register(server grpc.ServiceRegistrar) {
	banappuser.RegisterBanAppUserMgrServer(server, &Server{})
}
