package cmds

import (
	"fmt"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/urfave/cli/v2"
	"github.com/ynwcel/gox/app/gclix/pkg"
)

func init() {
	var (
		output   string = filepath.Base(pkg.MustGetGoMod())
		go_build *cli.Command
	)
	if strings.ToLower(runtime.GOOS) == "windows" {
		output = fmt.Sprintf("%s.ext", output)
	}

	go_build = &cli.Command{
		Name:      "go-build",
		Usage:     "go-build golang project",
		UsageText: fmt.Sprintf("%s go-build [options...]", appName),
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:        "i",
				Aliases:     []string{"install"},
				DefaultText: "false",
			},
			&cli.StringFlag{
				Name:        "o",
				Aliases:     []string{"output"},
				DefaultText: output,
			},
			&cli.StringFlag{
				Name:        "os",
				DefaultText: runtime.GOOS,
			},
			&cli.StringFlag{
				Name:        "arch",
				DefaultText: runtime.GOARCH,
			},
			&cli.StringSliceFlag{
				Name:        "args",
				DefaultText: "[]string",
			},
		},
		Action: buildAction,
	}
	clixApp.Commands = append(clixApp.Commands, go_build)
}

func buildAction(ctx *cli.Context) error {
	return nil
}
