package cmds

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
	"github.com/ynwcel/gox/xnum"
)

var (
	arg_type     string
	numDecodeCmd = &cli.Command{
		Name:  "num-decode",
		Usage: fmt.Sprintf("%s num-decode --type=xxx <value>", appName),
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "type",
				DefaultText: "2/8/16Lower/16Upper/26Lower/26Upper/32Lower/32Upper/36Lower/36Upper/52/58/62",
				Value:       "10",
				Destination: &arg_type,
			},
		},
		Action: numDecodeAction,
	}
)

func numDecodeAction(ctx *cli.Context) error {
	if ctx.NArg() < 1 {
		cli.ShowSubcommandHelp(ctx)
		return nil
	}
	var (
		type_name = fmt.Sprintf("To%s", arg_type)
		src_value = ctx.Args().Get(0)
	)
	converts := map[string]xnum.NumConverter{
		"To2":       xnum.NewConvertTo2(),
		"To8":       xnum.NewConvertTo8(),
		"To10":      xnum.NewConvertTo10(),
		"To16Lower": xnum.NewConvertTo16Lower(),
		"To16Upper": xnum.NewConvertTo16Upper(),
		"To26Lower": xnum.NewConvertTo26Lower(),
		"To26Upper": xnum.NewConvertTo26Upper(),
		"To32Lower": xnum.NewConvertTo32Lower(),
		"To32Upper": xnum.NewConvertTo32Upper(),
		"To36Lower": xnum.NewConvertTo36Lower(),
		"To36Upper": xnum.NewConvertTo36Upper(),
		"To52":      xnum.NewConvertTo52(),
		"To58":      xnum.NewConvertTo58(),
		"To62":      xnum.NewConvertTo62(),
	}
	if convert, ok := converts[type_name]; !ok {
		return fmt.Errorf("decode type[=%s] not found!", arg_type)
	} else if int_64, err := convert.Decode(src_value); err != nil {
		return err
	} else {
		fmt.Fprintf(os.Stdout, "Decode %s/%s  = %d", src_value, arg_type, int_64)
	}
	return nil
}
