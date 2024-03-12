package pclix

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/urfave/cli/v2"
)

type App = cli.App
type Cmd = cli.Command

func NewApp(appName, appVersion string) *App {
	app := cli.NewApp()
	app.Name = appName
	app.Version = appVersion
	app.Usage = fmt.Sprintf("V%s", appVersion)

	app.ExitErrHandler = func(cCtx *cli.Context, err error) {
		if err != nil {
			var (
				cmdName = cCtx.Command.Name
			)
			color.Red("Error:%s", err.Error())
			if appName != cmdName {
				color.Green("please use `%s %s --help` get more", appName, cmdName)
			} else {
				color.Green("please use `%s --help` get more", appName)
			}
			os.Exit(3)
		}
	}
	app.OnUsageError = func(cCtx *cli.Context, err error, isSubcommand bool) error {
		return err
	}
	return app
}

func NewCmd(cmdName string) *Cmd {
	return &cli.Command{
		Name: cmdName,
		OnUsageError: func(cCtx *cli.Context, err error, isSubcommand bool) error {
			return err
		},
	}
}
