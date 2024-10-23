package main

import (
	"github.com/ynwcel/gox/app/gclix/cmds"
)

var (
	appVersion = "0.0.1"
)

func main() {
	if err := cmds.RunWithVersion(appVersion); err != nil {
		panic(err)
	}
}
