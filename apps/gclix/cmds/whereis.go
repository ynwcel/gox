package cmds

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/urfave/cli/v2"
	"github.com/ynwcel/gox/xos"
)

var whereIsCmd = &cli.Command{
	Name:   "whereis",
	Usage:  fmt.Sprintf("%s whereis <cmd-name>", appName),
	Action: whereIsAction,
}

func whereIsAction(ctx *cli.Context) error {
	if ctx.NArg() < 1 {
		cli.ShowSubcommandHelp(ctx)
		return nil
	}
	var (
		find_cmd_path = ""
		find_cmd_name = ctx.Args().Get(0)
		env_path      = os.Getenv("PATH")
		find_paths    = []string{}
	)
	if xos.IsWindows() {
		find_paths = strings.Split(env_path, ";")
		find_cmd_name = fmt.Sprintf("%s.exe", strings.TrimRight(find_cmd_name, "."))
	} else {
		find_paths = strings.Split(env_path, ":")
	}
	find_paths = append(find_paths, ".")
	for _, path := range find_paths {
		find_path := filepath.Clean(filepath.Join(path, find_cmd_name))
		if xos.PathIsFile(find_path) {
			find_cmd_path = fmt.Sprintf("%s : %s", find_cmd_name, find_path)
			break
		}
	}
	if len(find_cmd_path) <= 0 {
		find_cmd_path = fmt.Sprintf("%s: <not-found>", find_cmd_name)
	}
	fmt.Println(find_cmd_path)
	return nil
}
