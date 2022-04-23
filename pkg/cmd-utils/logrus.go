package cmd_utils

import (
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

func SetupLogrus(cliCtx *cli.Context) {
	// Log level
	logLevelStr := cliCtx.String(LogLevel.Name)
	logLevel, err := logrus.ParseLevel(logLevelStr)
	if err != nil {
		logrus.SetReportCaller(true)
		logrus.Fatal(err)
	}
	logrus.SetLevel(logLevel)

	// Log caller
	logCaller := cliCtx.Bool(LogCaller.Name)
	logrus.SetReportCaller(logCaller)
}
