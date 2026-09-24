package xos

import (
	"context"
	"os"
	"os/exec"
)

func Shell(ctx context.Context, cmd string, args ...string) ([]byte, error) {

	shellCmd := exec.CommandContext(ctx, cmd, args...)
	shellCmd.Dir = "."
	shellCmd.Env = os.Environ()

	return shellCmd.CombinedOutput()
}
