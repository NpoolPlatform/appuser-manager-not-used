package appcontrol

import (
	"github.com/NpoolPlatform/message/npool/appuser/mgr/v2/appcontrol"
	"google.golang.org/grpc"
)

type Server struct {
	appcontrol.UnimplementedAppControlMgrServer
}

func Register(server grpc.ServiceRegistrar) {
	appcontrol.RegisterAppControlMgrServer(server, &Server{})
}
