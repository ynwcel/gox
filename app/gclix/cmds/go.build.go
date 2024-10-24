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

func init() {
	var (
		go_build = &cli.Command{
			Name:      "go-build",
			Usage:     "go-build golang project",
			UsageText: fmt.Sprintf("%s go-build [options...]", appName),
			Action:    buildAction,
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
	if mod_name, err := pkg.GetGoModName(); err == nil {
		output_txt := fmt.Sprintf("%s.%s.%s", filepath.Base(mod_name), runtime.GOOS, runtime.GOARCH)
		if strings.ToLower(runtime.GOOS) == "windows" {
			output_txt = fmt.Sprintf("%s.exe", output_txt)
		}
		output_flags.DefaultText = output_txt

		dist_flags.Value = fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH)
		dist_flags.DefaultText = fmt.Sprintf("%s/%s [use `go tool dist list` list all]", runtime.GOOS, runtime.GOARCH)
	}
	go_build.Flags = append(go_build.Flags, output_flags, dist_flags)
	clixApp.Commands = append(clixApp.Commands, go_build)
}

func buildAction(ctx *cli.Context) error {
	var (
		gocmd = exec.Command("go", "build")

		dist   = ctx.String("dist")
		output = ctx.String("output")

		target_os   = runtime.GOOS
		target_arch = runtime.GOARCH
		mode_name   = filepath.Base(pkg.MustGetGoMod())
	)
	gocmd.Env = os.Environ()
	if !strings.EqualFold(dist, fmt.Sprintf("%s/%s", target_os, target_arch)) {
		dists := strings.SplitN(dist, "/", 2)
		if len(dists) < 2 {
			return errors.New("--dist must use `os/arch`")
		}
		target_os, target_arch = dists[0], dists[1]
		gocmd.Env = append(gocmd.Env, "CGO_ENABLE=0")
		gocmd.Env = append(gocmd.Env, fmt.Sprintf("GOOS=%s", target_os))
		gocmd.Env = append(gocmd.Env, fmt.Sprintf("GOARCH=%s", target_arch))
	}

	gocmd.Args = append(gocmd.Args, "-ldflags", fmt.Sprintf("-X main.buildVersion=%s", build_version()))
	if len(output) <= 0 {
		output = fmt.Sprintf("%s.%s.%s", mode_name, target_os, target_arch)
		if target_os == "windows" {
			output = fmt.Sprintf("%s.exe", output)
		}
	}
	gocmd.Args = append(gocmd.Args, "-o", output)

	fmt.Println(gocmd)

	if _, err := gocmd.CombinedOutput(); err != nil {
		return err
	}
	return nil
}

func build_version() string {
	var (
		version      = time.Now().Format("20060102")
		git_commitid string
	)
	if v, err := exec.Command("git", "describe", "rev-parse", "--short HEAD").Output(); err == nil {
		git_commitid = string(v)
	} else {
		git_commitid = "nocommitid"
	}
	return fmt.Sprintf("%s.%s", version, git_commitid)
}
