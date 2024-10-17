package clix

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
)

var (
	appName = "ghack"
	clixApp = &cli.App{
		Name:            appName,
		Usage:           "golang dev hack tools",
		UsageText:       fmt.Sprintf("%s <command> [options...]", appName),
		HideHelpCommand: true,
	}
)

func RunWithVersion(version string) error {
	clixApp.Version = version
	return clixApp.Run(os.Args)
}
