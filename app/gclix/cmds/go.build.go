package cmds

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"github.com/urfave/cli/v2"
	"github.com/ynwcel/gox/app/gclix/pkg"
)

var goBuildCmd = &cli.Command{
	Name:      "go-build",
	Usage:     "go-build golang project",
	UsageText: fmt.Sprintf("%s go-build [options...]", appName),
	Action:    goBuildAction,
}

func init() {
	var (
		generate_flag = &cli.BoolFlag{
			Name:        "generate",
			Aliases:     []string{"g"},
			DefaultText: "true",
			Value:       true,
		}
		install_flags = &cli.BoolFlag{
			Name:    "install",
			Aliases: []string{"i"},
		}
		output_flags = &cli.StringFlag{
			Name:    "output",
			Aliases: []string{"o"},
		}
		dist_flags = &cli.StringFlag{
			Name:    "dist",
			Aliases: []string{"d"},
		}
	)
	if pkg.FileExists(GOMOD_FILE) {
		mod_name := pkg.MustGetGoMod()
		output_txt := fmt.Sprintf("%s.%s.%s", filepath.Base(mod_name), runtime.GOOS, runtime.GOARCH)
		if strings.ToLower(runtime.GOOS) == OS_WINDOWS {
			output_txt = fmt.Sprintf("%s.exe", output_txt)
		}
		output_flags.DefaultText = output_txt
	}

	dist_flags.Value = fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH)
	dist_flags.DefaultText = fmt.Sprintf("%s/%s [use `go tool dist list` list all]", runtime.GOOS, runtime.GOARCH)
	goBuildCmd.Flags = append(goBuildCmd.Flags, generate_flag, install_flags, output_flags, dist_flags)
}

func goBuildAction(ctx *cli.Context) error {
	var (
		goGenCmd   *exec.Cmd
		goBuildCmd *exec.Cmd

		generate = ctx.Bool("generate")
		install  = ctx.Bool("install")
		dist     = ctx.String("dist")
		output   = ctx.String("output")

		target_os   = runtime.GOOS
		target_arch = runtime.GOARCH
		mode_name   = filepath.Base(pkg.MustGetGoMod())
	)
	if install && len(output) > 0 {
		return fmt.Errorf("`go install` command not support `--output` flag")
	} else if install {
		goBuildCmd = exec.Command("go", "install")
	} else {
		goBuildCmd = exec.Command("go", "build")
	}

	goBuildCmd.Env = os.Environ()
	if !strings.EqualFold(dist, fmt.Sprintf("%s/%s", target_os, target_arch)) {
		dists := strings.SplitN(dist, "/", 2)
		if len(dists) < 2 {
			return errors.New("--dist must use `os/arch`")
		}
		target_os, target_arch = dists[0], dists[1]
		goBuildCmd.Env = append(goBuildCmd.Env, "CGO_ENABLE=0")
		goBuildCmd.Env = append(goBuildCmd.Env, fmt.Sprintf("GOOS=%s", target_os))
		goBuildCmd.Env = append(goBuildCmd.Env, fmt.Sprintf("GOARCH=%s", target_arch))
	}

	goBuildCmd.Args = append(goBuildCmd.Args, "-ldflags", fmt.Sprintf("-X main.buildVersion=%s", build_version()))
	if !install {
		if len(output) <= 0 {
			output = fmt.Sprintf("%s.%s.%s", mode_name, target_os, target_arch)
			if target_os == OS_WINDOWS {
				output = fmt.Sprintf("%s.exe", output)
			}
		}
		goBuildCmd.Args = append(goBuildCmd.Args, "-o", output)
	}

	if generate {
		goGenCmd = exec.Command("go", "generate", "./...")
		goGenCmd.Env = goBuildCmd.Env[:]
		fmt.Println(goGenCmd)

		if _, err := goGenCmd.CombinedOutput(); err != nil {
			return err
		}
	}

	fmt.Println(goBuildCmd)

	if _, err := goBuildCmd.CombinedOutput(); err != nil {
		return err
	}
	return nil
}

func build_version() string {
	var (
		version      = time.Now().Format("060102.1504")
		git_commitid string
	)
	if v, err := exec.Command("git", "describe", "rev-parse", "--short HEAD").Output(); err == nil {
		git_commitid = string(v)
	} else {
		git_commitid = "no-commit-id"
	}
	return fmt.Sprintf("%s.%s", version, git_commitid)
}
