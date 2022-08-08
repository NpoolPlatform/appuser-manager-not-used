package api

import (
	"context"

	appusermgrv1 "github.com/NpoolPlatform/appuser-manager/api/v1"
	"github.com/NpoolPlatform/appuser-manager/api/v2/app"
	"github.com/NpoolPlatform/appuser-manager/api/v2/appcontrol"
	"github.com/NpoolPlatform/appuser-manager/api/v2/approle"
	"github.com/NpoolPlatform/appuser-manager/api/v2/approleuser"
	"github.com/NpoolPlatform/appuser-manager/api/v2/appuser"
	"github.com/NpoolPlatform/appuser-manager/api/v2/appusercontrol"
	"github.com/NpoolPlatform/appuser-manager/api/v2/appuserextra"
	"github.com/NpoolPlatform/appuser-manager/api/v2/appusersecret"
	"github.com/NpoolPlatform/appuser-manager/api/v2/appuserthirdparty"
	"github.com/NpoolPlatform/appuser-manager/api/v2/banapp"
	"github.com/NpoolPlatform/appuser-manager/api/v2/banappuser"
	appusermgr "github.com/NpoolPlatform/message/npool/appuser/mgr/v2"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	appusermgr.UnimplementedManagerServer
}

func Register(server grpc.ServiceRegistrar) {
	appusermgr.RegisterManagerServer(server, &Server{})
	appusermgrv1.Register(server)
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
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	if err := appusermgrv1.RegisterGateway(mux, endpoint, opts); err != nil {
		return err
	}
	if err := appusermgr.RegisterManagerHandlerFromEndpoint(context.Background(), mux, endpoint, opts); err != nil {
		return err
	}
	return nil
}
