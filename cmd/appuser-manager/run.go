package main

import (
	"github.com/NpoolPlatform/appuser-manager/api"
	db "github.com/NpoolPlatform/appuser-manager/pkg/db"
	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	apimgrcli "github.com/NpoolPlatform/api-manager/pkg/client"

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
	apimgrcli.RegisterGRPC(server)
	return nil
}

func rpcGatewayRegister(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	err := api.RegisterGateway(mux, endpoint, opts)
	if err != nil {
		return err
	}

	apimgrcli.Register(mux)
	return nil
}
