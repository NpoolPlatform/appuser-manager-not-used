package appusersecret

import (
	"github.com/NpoolPlatform/message/npool/appuser/mgr/v2/appusersecret"
	"google.golang.org/grpc"
)

type Server struct {
	appusersecret.UnimplementedAppUserSecretMgrServer
}

func Register(server grpc.ServiceRegistrar) {
	appusersecret.RegisterAppUserSecretMgrServer(server, &Server{})
}
