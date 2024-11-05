package cmds

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
	"github.com/ynwcel/gox/app/gclix/pkg"
)

const (
	OS_WINDOWS = "windows"
	OS_LINUX   = "linux"
	OS_DARWIN  = "darwin"
	GOMOD_FILE = "./go.mod"
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

func init() {
	if pkg.FileExists(GOMOD_FILE) {
		clixApp.Commands = append(clixApp.Commands, goBuildCmd, goRenameGoModCmd)
	}

	clixApp.Commands = append(clixApp.Commands, whereIsCmd)
}

func RunWithVersion(version string) error {
	clixApp.Version = version
	return clixApp.Run(os.Args)
}
