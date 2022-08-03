package approleuser

import (
	"github.com/NpoolPlatform/message/npool/appuser/mgr/v2/approleuser"
	"google.golang.org/grpc"
)

type Server struct {
	approleuser.UnimplementedAppRoleUserMgrServer
}

func Register(server grpc.ServiceRegistrar) {
	approleuser.RegisterAppRoleUserMgrServer(server, &Server{})
}
