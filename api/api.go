package api

import (
	"context"

	"github.com/NpoolPlatform/message/npool/appusermgrv2/appuserextra"
	"github.com/NpoolPlatform/message/npool/appusermgrv2/appusersecret"
	"github.com/NpoolPlatform/message/npool/appusermgrv2/appuserthirdparty"
	"github.com/NpoolPlatform/message/npool/appusermgrv2/banapp"
	"github.com/NpoolPlatform/message/npool/appusermgrv2/banappuser"

	npool "github.com/NpoolPlatform/message/npool/appusermgr"
	"github.com/NpoolPlatform/message/npool/appusermgrv2/app"
	"github.com/NpoolPlatform/message/npool/appusermgrv2/approle"
	"github.com/NpoolPlatform/message/npool/appusermgrv2/approleuser"
	"github.com/NpoolPlatform/message/npool/appusermgrv2/appuser"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	npool.UnimplementedAppUserManagerServer
	app.UnimplementedAppUserManagerAppServer
	approle.UnimplementedAppUserManagerAppRoleServer
	approleuser.UnimplementedAppUserManagerAppRoleUserServer
	appuser.UnimplementedAppUserManagerAppUserServer
	appuserextra.UnimplementedAppUserManagerAppUserExtraServer
	appusersecret.UnimplementedAppUserManagerAppUserSecretServer
	appuserthirdparty.UnimplementedAppUserManagerAppUserThirdPartyServer
	banapp.UnimplementedAppUserManagerBanAppServer
	banappuser.UnimplementedAppUserManagerBanAppUserServer
}

func Register(server grpc.ServiceRegistrar) {
	npool.RegisterAppUserManagerServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return npool.RegisterAppUserManagerHandlerFromEndpoint(context.Background(), mux, endpoint, opts)
}
