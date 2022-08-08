package cmd_utils

import (
	"github.com/urfave/cli/v2"
)

var (
	GrpcPort = &cli.IntFlag{
		Name:  "grpc-port",
		Usage: "TODO: xxx",
		Value: 9090,
	}
	LogLevel = &cli.StringFlag{
		Name:  "log-level",
		Usage: "TODO: xxx",
		Value: "info",
	}
	LogCaller = &cli.BoolFlag{
		Name:  "log-caller",
		Usage: "Show log caller for debugging purpose",
		Value: true,
	}
)
