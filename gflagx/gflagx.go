package gflagx

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strings"

	"github.com/spf13/pflag"
)

type Flagx struct {
	*pflag.FlagSet
	help_flag bool
	version   string
}

func NewFlagx(appName ...string) *Flagx {
	app := filepath.Base(os.Args[0])
	if len(appName) > 0 && len(appName[0]) > 0 {
		app = appName[0]
	}
	fx := &Flagx{}
	fx.FlagSet = pflag.NewFlagSet(app, pflag.ContinueOnError)
	fx.FlagSet.BoolVarP(&fx.help_flag, "help", "h", false, "show help messgae")
	fx.FlagSet.Usage = func() {
		var (
			out      = bytes.NewBuffer(nil)
			usages   = strings.Split(strings.Trim(fx.FlagSet.FlagUsages(), "\n"), "\n")
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
		if len(fx.version) > 0 {
			fmt.Fprintln(out, "Version:")
			fmt.Fprintf(out, "   %s", fx.version)
		}
		fmt.Println(out.String())
	}
	return fx
}

func (gfx *Flagx) SetVersion(version string) *Flagx {
	gfx.version = version
	return gfx
}

func (gfx *Flagx) GetVersion() string {
	return gfx.version
}

func (gfx *Flagx) HasSetHelpFlag() bool {
	return gfx.help_flag
}
