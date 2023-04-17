package api

import (
	"context"

	"github.com/NpoolPlatform/appuser-manager/api/approle"
	"github.com/NpoolPlatform/appuser-manager/api/approleuser"
	"github.com/NpoolPlatform/appuser-manager/api/appuser"
	"github.com/NpoolPlatform/appuser-manager/api/appusercontrol"
	"github.com/NpoolPlatform/appuser-manager/api/appuserextra"
	"github.com/NpoolPlatform/appuser-manager/api/appusersecret"
	"github.com/NpoolPlatform/appuser-manager/api/appuserthirdparty"
	"github.com/NpoolPlatform/appuser-manager/api/authing/auth"
	authhis "github.com/NpoolPlatform/appuser-manager/api/authing/history"
	"github.com/NpoolPlatform/appuser-manager/api/banapp"
	"github.com/NpoolPlatform/appuser-manager/api/banappuser"
	"github.com/NpoolPlatform/appuser-manager/api/kyc"
	loginhis "github.com/NpoolPlatform/appuser-manager/api/login/history"
	"github.com/NpoolPlatform/appuser-manager/api/subscriber"

	appusermgr "github.com/NpoolPlatform/message/npool/appuser/mgr/v2"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	appusermgr.UnimplementedManagerServer
}

func Register(server grpc.ServiceRegistrar) {
	appusermgr.RegisterManagerServer(server, &Server{})
	subscriber.Register(server)
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
