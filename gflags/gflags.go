package gflags

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strings"

	"github.com/spf13/pflag"
)

type gflagx struct {
	*pflag.FlagSet
	help_flag bool
	version   string
}

func New(appName ...string) *gflagx {
	app := filepath.Base(os.Args[0])
	if len(appName) > 0 && len(appName[0]) > 0 {
		app = appName[0]
	}
	cx := &gflagx{}
	cx.FlagSet = pflag.NewFlagSet(app, pflag.ContinueOnError)
	cx.FlagSet.BoolVarP(&cx.help_flag, "help", "h", false, "show help messgae")
	cx.FlagSet.Usage = func() {
		var (
			out      = bytes.NewBuffer(nil)
			usages   = strings.Split(strings.Trim(cx.FlagSet.FlagUsages(), "\n"), "\n")
			help_idx = slices.IndexFunc(usages, func(line string) bool {
				return strings.Contains(line, "--help")
			})
		)
		fmt.Fprintln(out, "Usage:")
		fmt.Fprintf(out, "   %s [options]\n", app)
		fmt.Fprintln(out, "Flags:")
		for idx, u := range usages {
			if idx == help_idx {
				continue
			}
			fmt.Fprintln(out, strings.TrimRight(u, "\n"))
		}
		fmt.Fprintln(out, usages[help_idx])
		if len(cx.version) > 0 {
			fmt.Fprintln(out, "Version:")
			fmt.Fprintf(out, "   %s", cx.version)
		}
		fmt.Println(out.String())
	}
	return cx
}

func (gfx *gflagx) SetVersion(version string) *gflagx {
	gfx.version = version
	return gfx
}

func (gfx *gflagx) GetVersion(version string) string {
	return gfx.version
}

func (gfx *gflagx) HasSetHelpFlag() bool {
	return gfx.help_flag
}
