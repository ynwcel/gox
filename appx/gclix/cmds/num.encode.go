package cmds

import (
	"fmt"
	"strconv"

	"github.com/urfave/cli/v2"
	"github.com/ynwcel/gox/xnum"
)

var numEncodeCmd = &cli.Command{
	Name:   "num-encode",
	Usage:  fmt.Sprintf("%s num-encode <int-number>", appName),
	Action: numEncodeAction,
}

func numEncodeAction(ctx *cli.Context) error {
	if ctx.NArg() < 1 {
		cli.ShowSubcommandHelp(ctx)
		return nil
	}
	if num_int, err := strconv.Atoi(ctx.Args().Get(0)); err != nil {
		return err
	} else {
		converts := []map[string]xnum.NumConverter{
			{"To2": xnum.NewConvertTo2()},
			{"To8": xnum.NewConvertTo8()},
			{"To10": xnum.NewConvertTo10()},
			{"To16Lower": xnum.NewConvertTo16Lower()},
			{"To16Upper": xnum.NewConvertTo16Upper()},
			{"To26Lower": xnum.NewConvertTo26Lower()},
			{"To26Upper": xnum.NewConvertTo26Upper()},
			{"To32Lower": xnum.NewConvertTo32Lower()},
			{"To32Upper": xnum.NewConvertTo32Upper()},
			{"To36Lower": xnum.NewConvertTo36Lower()},
			{"To36Upper": xnum.NewConvertTo36Upper()},
			{"To52": xnum.NewConvertTo52()},
			{"To58": xnum.NewConvertTo58()},
			{"To62": xnum.NewConvertTo62()},
		}
		fmt.Println("Encode", num_int, ":")
		for _, info := range converts {
			for name, convert := range info {
				fmt.Printf("    %-10s: %s\n", name, convert.Encode(int64(num_int)))
			}
		}
	}
	return nil
}
