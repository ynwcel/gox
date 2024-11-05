package cmds

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/urfave/cli/v2"
	"github.com/ynwcel/gox/app/gclix/pkg"
)

var (
	goRenameGoMod = &cli.Command{
		Name:      "go-rename-gomod",
		Usage:     "rename go mode name",
		UsageText: fmt.Sprintf("%s go-rename-gomod <new mod name>", appName),
		Action:    goRenameGoModAction,
	}
)

func goRenameGoModAction(ctx *cli.Context) error {
	if ctx.NArg() < 1 {
		cli.ShowSubcommandHelp(ctx)
		return nil
	}
	var (
		new_mod_name       = ctx.Args().First()
		new_mod_name_bytes []byte
		rewrite_files      = []string{}
		old_mod_name       string
		old_mod_name_bytes []byte

		mod_tidy_cmds = []string{
			"mod",
			"tidy",
		}
		mod_regexp, err = regexp.Compile(`^[\w][\w\d\/\-]+$`)
	)
	if !mod_regexp.Match([]byte(new_mod_name)) {
		return fmt.Errorf("new mod name `%s` is failed", new_mod_name)
	}
	new_mod_name_bytes = []byte(new_mod_name)

	if old_mod_name, err = pkg.GetGoModName(); err != nil {
		return err
	}
	old_mod_name_bytes = []byte(old_mod_name)

	// 1.处理 go.mod
	if f_content, err := os.ReadFile(GOMOD_FILE); err != nil {
		return err
	} else {
		f_content = bytes.Replace(f_content, old_mod_name_bytes, new_mod_name_bytes, -1)
		if err = os.WriteFile(GOMOD_FILE, f_content, 0666); err != nil {
			return fmt.Errorf("save go.mod error:%W", err)
		}
	}
	// 2.处理 go 文件
	if rewrite_files, err = pkg.ListFile(".", `.*?\.go.*?`); err != nil {
		return err
	}
	rewrite_replaces := map[string]string{
		// import "mod_name"
		fmt.Sprintf(`%s"`, old_mod_name): fmt.Sprintf(`%s"`, new_mod_name),
		// import "mod_name/sub_mod"
		fmt.Sprintf(`%s/`, old_mod_name): fmt.Sprintf(`%s/`, new_mod_name),
		// mod_name.XXX
		fmt.Sprintf(`%s.`, filepath.Base((old_mod_name))): fmt.Sprintf(`%s.`, filepath.Base((new_mod_name))),
	}
	for _, gofile := range rewrite_files {
		var (
			f_content, err = os.ReadFile(gofile)
		)
		if err != nil {
			return fmt.Errorf("read-go-file-error:%W", err)
		}
		for old_str, new_str := range rewrite_replaces {
			f_content = bytes.ReplaceAll(f_content, []byte(old_str), []byte(new_str))
		}
		if err = os.WriteFile(gofile, f_content, 0666); err != nil {
			return fmt.Errorf("save go file error:%W", err)
		}
	}
	fmt.Printf("> go %s\n", strings.Join(mod_tidy_cmds, " "))
	if err = exec.Command("go", mod_tidy_cmds...).Run(); err != nil {
		return fmt.Errorf("run go mod tidy error:%W", err)
	}
	return err
}
