package api

import (
	"context"
	"github.com/NpoolPlatform/appuser-manager/api/v2/kyc"

	"github.com/NpoolPlatform/appuser-manager/api/v2/app"
	"github.com/NpoolPlatform/appuser-manager/api/v2/appcontrol"
	"github.com/NpoolPlatform/appuser-manager/api/v2/approle"
	"github.com/NpoolPlatform/appuser-manager/api/v2/approleuser"
	"github.com/NpoolPlatform/appuser-manager/api/v2/appuser"
	"github.com/NpoolPlatform/appuser-manager/api/v2/appusercontrol"
	"github.com/NpoolPlatform/appuser-manager/api/v2/appuserextra"
	"github.com/NpoolPlatform/appuser-manager/api/v2/appusersecret"
	"github.com/NpoolPlatform/appuser-manager/api/v2/appuserthirdparty"
	"github.com/NpoolPlatform/appuser-manager/api/v2/authing/auth"
	authhis "github.com/NpoolPlatform/appuser-manager/api/v2/authing/history"
	"github.com/NpoolPlatform/appuser-manager/api/v2/banapp"
	"github.com/NpoolPlatform/appuser-manager/api/v2/banappuser"
	loginhis "github.com/NpoolPlatform/appuser-manager/api/v2/login/history"

	appusermgr "github.com/NpoolPlatform/message/npool/appuser/mgr/v2"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	appusermgr.UnimplementedManagerServer
}

func Register(server grpc.ServiceRegistrar) {
	appusermgr.RegisterManagerServer(server, &Server{})
	app.Register(server)
	appcontrol.Register(server)
	approle.Register(server)
	approleuser.Register(server)
	appuser.Register(server)
	appusercontrol.Register(server)
	appuserextra.Register(server)
	appusersecret.Register(server)
	appuserthirdparty.Register(server)
	banapp.Register(server)
	banappuser.Register(server)
	auth.Register(server)
	authhis.Register(server)
	loginhis.Register(server)
	kyc.Register(server)
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	if err := appusermgr.RegisterManagerHandlerFromEndpoint(context.Background(), mux, endpoint, opts); err != nil {
		return err
	}
	return nil
}
