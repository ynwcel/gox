package cmds

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/urfave/cli/v2"
	"github.com/ynwcel/gox/app/gclix/pkg"
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
		cmdName       = ctx.Args().Get(0)
		cmdName_exe   = fmt.Sprintf("%s.exe", cmdName)
		env_path      = os.Getenv("PATH")
		find_paths    = []string{}
	)
	if runtime.GOOS == OS_WINDOWS {
		find_paths = strings.Split(env_path, ";")
	} else {
		find_paths = strings.Split(env_path, ":")
	}
	find_paths = append(find_paths, ".")
	for _, path := range find_paths {
		find_path := filepath.Clean(filepath.Join(path, cmdName))
		find_path_exe := filepath.Clean(filepath.Join(path, cmdName_exe))
		if pkg.FileExists(find_path) {
			find_cmd_path = fmt.Sprintf("%s : %s", cmdName, find_path)
			break
		} else if pkg.FileExists(find_path_exe) {
			find_cmd_path = fmt.Sprintf("%s : %s", cmdName, find_path_exe)
			break
		}
	}
	if len(find_cmd_path) <= 0 {
		find_cmd_path = fmt.Sprintf("%s:", cmdName)
	}
	fmt.Println(find_cmd_path)
	return nil
}
