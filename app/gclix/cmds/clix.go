package cmds

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
)

var (
	appName = "gclix"
	clixApp = &cli.App{
		Name:            appName,
		Usage:           "dev tools for golang",
		UsageText:       fmt.Sprintf("%s <command> [options...]", appName),
		HideHelpCommand: true,
	}
)

func RunWithVersion(version string) error {
	clixApp.Version = version
	return clixApp.Run(os.Args)
}
