package approle

import (
	"github.com/NpoolPlatform/message/npool/appuser/mgr/v2/approle"
	"google.golang.org/grpc"
)

type Server struct {
	approle.UnimplementedAppRoleMgrServer
}

func Register(server grpc.ServiceRegistrar) {
	approle.RegisterAppRoleMgrServer(server, &Server{})
}
