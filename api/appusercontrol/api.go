package appusercontrol

import (
	"github.com/NpoolPlatform/message/npool/appuser/mgr/v2/appusercontrol"
	"google.golang.org/grpc"
)

type Server struct {
	appusercontrol.UnimplementedManagerServer
}

func Register(server grpc.ServiceRegistrar) {
	appusercontrol.RegisterManagerServer(server, &Server{})
}
