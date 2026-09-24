package xos

import (
	"bufio"
	"fmt"
	"os"
)

func Input(tip string) string {
	fmt.Println("> ", tip)
	var r = bufio.NewReader(os.Stdin)
	for {
		fmt.Println("<:")
		if line, err := r.ReadBytes('\n'); err == nil {
			return string(line)
		}
	}
}
