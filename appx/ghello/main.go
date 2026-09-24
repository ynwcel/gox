package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Printf("Hello -- %s\n", time.Now().Format("2006-01-02 15:04:05.000"))
}
