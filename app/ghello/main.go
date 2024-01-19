package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(fmt.Sprintf("Hello -- %s", time.Now().Format("2006-01-02 15:04:05.000")))
}
