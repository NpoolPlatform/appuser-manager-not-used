package main

import (
	"github.com/NpoolPlatform/appuser-manager/api"
	"github.com/NpoolPlatform/appuser-manager/pkg/migrator"

	db "github.com/NpoolPlatform/appuser-manager/pkg/db"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	apicli "github.com/NpoolPlatform/basal-middleware/pkg/client/api"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"

	cli "github.com/urfave/cli/v2"

	"google.golang.org/grpc"
)

var runCmd = &cli.Command{
	Name:    "run",
	Aliases: []string{"r"},
	Usage:   "Run the daemon",
	After: func(context *cli.Context) error {
		if err := grpc2.HShutdown(); err != nil {
			logger.Sugar().Warnf("graceful shutdown server")
		}
		grpc2.GShutdown()
		return logger.Sync()
	},
	Action: func(c *cli.Context) error {
		if err := db.Init(); err != nil {
			return err
		}

		if err := migrator.Migrate(c.Context); err != nil {
			return err
		}

		go func() {
			if err := grpc2.RunGRPC(rpcRegister); err != nil {
				logger.Sugar().Errorf("fail to run grpc server: %v", err)
			}
		}()
		return grpc2.RunGRPCGateWay(rpcGatewayRegister)
	},
}

func rpcRegister(server grpc.ServiceRegistrar) error {
	api.Register(server)
	apicli.RegisterGRPC(server)
	return nil
}

func rpcGatewayRegister(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	err := api.RegisterGateway(mux, endpoint, opts)
	if err != nil {
		return err
	}

	_ = apicli.Register(mux)

	return nil
}
