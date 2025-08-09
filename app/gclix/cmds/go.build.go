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
		output_flags.DefaultText = build_name(runtime.GOOS, runtime.GOARCH)
	}

	dist_flags.DefaultText = fmt.Sprintf("%s/%s [use `go tool dist list` list all]", runtime.GOOS, runtime.GOARCH)
	dist_flags.Value = fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH)
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

		cur_datetime = time.Now()
		cur_date     = cur_datetime.Format("060102")
		cur_time     = cur_datetime.Format("1504")
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

	goBuildCmd.Args = append(goBuildCmd.Args, "-ldflags", fmt.Sprintf("-X main.buildVersion=%s.%s.%s", build_git_commitid(), cur_date, cur_time))

	if !install {
		if len(output) <= 0 {
			output = build_name(target_os, target_arch)
		} else if strings.EqualFold(target_os, OS_WINDOWS) && !strings.Contains(output, ".exe") {
			output = fmt.Sprintf("%s.exe", output)
		}

		var (
			output_has_exe  = strings.Contains(strings.ToLower(output), ".exe")
			output_ext_name = filepath.Ext(output)
		)
		if output_has_exe {
			output = strings.TrimSuffix(output, output_ext_name)
		}
		if !strings.Contains(output, target_os) {
			output = fmt.Sprintf("%s.%s", output, target_os)
		}
		if !strings.Contains(output, target_arch) {
			output = fmt.Sprintf("%s.%s", output, target_arch)
		}
		if !strings.Contains(output, cur_date) {
			output = fmt.Sprintf("%s.%s.%s", output, cur_date, cur_time)
		}
		if output_has_exe {
			output = output + output_ext_name
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

func build_name(os, arch string) string {
	name := fmt.Sprintf("%s.%s.%s.%s", filepath.Base(pkg.MustGetGoMod()), os, arch, time.Now().Format("060102.1504"))
	if strings.EqualFold(os, OS_WINDOWS) {
		name = fmt.Sprintf("%s.exe", name)
	}
	return name
}

func build_git_commitid() string {
	var (
		git_commitid string
	)
	if v, err := exec.Command("git", "rev-parse", "--short", "HEAD").Output(); err == nil {
		git_commitid = strings.TrimSpace(string(v))
	}
	return git_commitid
}
