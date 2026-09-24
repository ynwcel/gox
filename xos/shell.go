package xos

import (
	"context"
	"os"
	"os/exec"
)

func Shell(ctx context.Context, cmd string, args ...string) ([]byte, error) {
	shellCmd := exec.Cmd{
		Env:  os.Environ(),
		Dir:  ".",
		Path: cmd,
		Args: args,
	}
	return shellCmd.CombinedOutput()
}
