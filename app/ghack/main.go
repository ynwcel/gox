package main

import (
	"github.com/gox/app/ghack/clix"
)

var (
	appVersion = "0.0.1"
)

func main() {
	if err := clix.RunWithVersion(appVersion); err != nil {
		panic(err)
	}
}
