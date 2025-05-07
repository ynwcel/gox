package cmds

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gabriel-vasile/mimetype"
	"github.com/gogf/gf/v2/os/gview"
	"github.com/urfave/cli/v2"
	"github.com/ynwcel/gox/app/gclix/pkg"
)

var (
	output_dir   = "./gclix.html.dist"
	htmlBuildCmd = &cli.Command{
		Name:      "html-build",
		Usage:     "html dev build website",
		UsageText: fmt.Sprintf("%s html-build [options...]", appName),
		Action:    htmlBuildAction,
	}
)

func init() {
	curtime := time.Now()
	output_dir = fmt.Sprintf("./%s.%s", output_dir, curtime.Format("060102"))
	version := 1
	for {
		output_dir = fmt.Sprintf("%s.v%d", output_dir, version)
		if !pkg.IsDir(output_dir) {
			break
		}
		version += 1
	}
	output_dir = fmt.Sprintf("./%s", filepath.Clean(output_dir))
	htmlBuildCmd.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:        "output",
			Aliases:     []string{"o"},
			DefaultText: output_dir,
			Value:       output_dir,
			Destination: &output_dir,
		},
	}

}

func htmlBuildAction(ctx *cli.Context) error {
	files, err := pkg.ListFile("./", ".*")
	if err != nil {
		panic(fmt.Errorf("list files error:%w", err))
	}
	if !pkg.IsDir(output_dir) {
		if err = os.MkdirAll(output_dir, os.ModePerm); err != nil {
			panic(fmt.Errorf("list files error:%w", err))
		}
	}
	view := gview.New()
	view.SetPath("./")
	for _, file := range files {
		target := filepath.Clean(filepath.Join(output_dir, file))
		if pkg.IsDir(file) {
			if err := os.MkdirAll(target, os.ModeDir); err != nil {
				panic(err)
			}
			continue
		}
		file_mimetype, err := mimetype.DetectFile(file)
		if err != nil {
			panic(fmt.Errorf("get file<%s> mimie type error:%w", file, err))
		}
		if strings.ToLower(file_mimetype.Extension()) == ".html" || strings.ToLower(file_mimetype.Extension()) == ".htm" {
			if content, err := view.Parse(ctx.Context, file); err != nil {
				panic(fmt.Errorf("parse html file<%s>:%w", file, err))
			} else if _, err := pkg.PutContent(target, content); err != nil {
				panic(fmt.Errorf("parse html file<%s>:%w", file, err))
			}
		} else {
			if _, err := pkg.CopyFile(file, target); err != nil {
				panic(fmt.Errorf("copy res file<%s> failed:%w", file, err))
			}
		}
	}
	return nil
}
