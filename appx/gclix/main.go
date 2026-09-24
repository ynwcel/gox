package main

import "github.com/ynwcel/gox/appx/gclix/cmds"

var (
	buildVersion = ""
)

func main() {
	if err := cmds.RunWithVersion(buildVersion); err != nil {
		panic(err)
	}
}
