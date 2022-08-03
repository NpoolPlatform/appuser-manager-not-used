package appusercontrol

import (
	"github.com/NpoolPlatform/message/npool/appuser/mgr/v2/appusercontrol"
	"google.golang.org/grpc"
)

type Server struct {
	appusercontrol.UnimplementedAppUserControlMgrServer
}

func Register(server grpc.ServiceRegistrar) {
	appusercontrol.RegisterAppUserControlMgrServer(server, &Server{})
}
