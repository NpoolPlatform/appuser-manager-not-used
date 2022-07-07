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
}

type AppServer struct {
	app.UnimplementedAppUserManagerAppServer
}

type AppRoleServer struct {
	approle.UnimplementedAppUserManagerAppRoleServer
}

type AppRoleUserServer struct {
	approleuser.UnimplementedAppUserManagerAppRoleUserServer
}

type AppUserServer struct {
	appuser.UnimplementedAppUserManagerAppUserServer
}

type AppUserExtraServer struct {
	appuserextra.UnimplementedAppUserManagerAppUserExtraServer
}

type AppUserSecretServer struct {
	appusersecret.UnimplementedAppUserManagerAppUserSecretServer
}

type AppUserThirdPartyServer struct {
	appuserthirdparty.UnimplementedAppUserManagerAppUserThirdPartyServer
}

type BanAppServer struct {
	banapp.UnimplementedAppUserManagerBanAppServer
}

type BanAppUserServer struct {
	banappuser.UnimplementedAppUserManagerBanAppUserServer
}

func Register(server grpc.ServiceRegistrar) {
	npool.RegisterAppUserManagerServer(server, &Server{})
	app.RegisterAppUserManagerAppServer(server, &AppServer{})
	approle.RegisterAppUserManagerAppRoleServer(server, &AppRoleServer{})
	approleuser.RegisterAppUserManagerAppRoleUserServer(server, &AppRoleUserServer{})
	appuser.RegisterAppUserManagerAppUserServer(server, &AppUserServer{})
	appuserextra.RegisterAppUserManagerAppUserExtraServer(server, &AppUserExtraServer{})
	appusersecret.RegisterAppUserManagerAppUserSecretServer(server, &AppUserSecretServer{})
	banapp.RegisterAppUserManagerBanAppServer(server, &BanAppServer{})
	banappuser.RegisterAppUserManagerBanAppUserServer(server, &BanAppUserServer{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return npool.RegisterAppUserManagerHandlerFromEndpoint(context.Background(), mux, endpoint, opts)
}
