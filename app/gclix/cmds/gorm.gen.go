package cmds

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/urfave/cli/v2"
)

var gormGenCmd = &cli.Command{
	Name:      "gorm-gen",
	UsageText: fmt.Sprintf("%s gorms [options]", appName),
	Usage:     "gorm gen models",
	HideHelp:  true,
	Action:    gorms_action,
}

func init() {
	var (
		host = &cli.StringFlag{
			Name:        "host",
			Aliases:     []string{"h"},
			DefaultText: "127.0.0.1",
			Value:       "127.0.0.1",
		}
		port = &cli.IntFlag{
			Name:        "port",
			Aliases:     []string{"P"},
			DefaultText: "3306",
			Value:       3306,
		}
		username = &cli.StringFlag{
			Name:        "username",
			Aliases:     []string{"u"},
			DefaultText: "root",
			Value:       "root",
		}
		password = &cli.StringFlag{
			Name:        "password",
			Aliases:     []string{"p"},
			DefaultText: "",
			Value:       "",
		}
		database = &cli.StringFlag{
			Name:        "database",
			Aliases:     []string{"d"},
			DefaultText: "",
			Value:       "",
		}
		driver = &cli.StringFlag{
			Name:        "driver",
			DefaultText: "mysql",
			Value:       "mysql",
		}
		output = &cli.StringFlag{
			Name:        "output",
			Aliases:     []string{"o"},
			DefaultText: "./internal/gorms",
			Value:       "./internal/gorms",
		}
	)
	gormGenCmd.Flags = append(gormGenCmd.Flags, driver, host, port, username, password, database, output)
}

func gorms_action(ctx *cli.Context) error {
	if ctx.NumFlags() <= 0 {
		cli.ShowSubcommandHelp(ctx)
		return nil
	}
	var (
		flag_driver   = ctx.String("driver")
		flag_host     = ctx.String("host")
		flag_port     = ctx.Int("port")
		flag_username = ctx.String("username")
		flag_password = ctx.String("password")
		flag_database = ctx.String("database")
		flag_output   = ctx.String("output")
		model_pkgname = filepath.Base(flag_output)
		dsn           = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", flag_username, flag_password, flag_host, flag_port, flag_database)
		cmd           = exec.Command("go", "run", "gorm.io/gen/tools/gentool@latest")
	)
	if len(flag_database) <= 0 {
		return errors.New("database flag not set")
	}
	cmd.Env = os.Environ()
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin

	cmd.Args = append(cmd.Args, "-db", flag_driver)
	cmd.Args = append(cmd.Args, "-dsn", dsn)
	cmd.Args = append(cmd.Args, "-modelPkgName", model_pkgname)
	cmd.Args = append(cmd.Args, "-outPath", flag_output)
	cmd.Args = append(cmd.Args, "-fieldNullable", "-fieldWithIndexTag", "-fieldWithTypeTag", "-fieldNullable")
	fmt.Println(cmd)
	return nil
}
