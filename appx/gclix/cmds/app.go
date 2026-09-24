package cmds

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/urfave/cli/v2"
	"github.com/ynwcel/gox/xgomod"
	"github.com/ynwcel/gox/xos"
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
	if xos.PathExists(filepath.Join(".", xgomod.FILE_NAME)) {
		clixApp.Commands = append(clixApp.Commands, goRenameGoModCmd)
	}

	clixApp.Commands = append(clixApp.Commands, goBuildCmd)
	clixApp.Commands = append(clixApp.Commands, htmlTestCmd, htmlBuildCmd)
	clixApp.Commands = append(clixApp.Commands, whereIsCmd)
	clixApp.Commands = append(clixApp.Commands, numEncodeCmd, numDecodeCmd)
}

func RunWithVersion(version string) error {
	clixApp.Version = version
	return clixApp.Run(os.Args)
}
