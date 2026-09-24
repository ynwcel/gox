package xos

import (
	"runtime"
)

const (
	NAME_OF_WINDOWS = "windows"
	NAME_OF_LINUX   = "linux"
	NAME_OF_DARWIN  = "darwin"
)

func IsWindows() bool {
	return runtime.GOOS == NAME_OF_WINDOWS
}

func IsLinux() bool {
	return runtime.GOOS == NAME_OF_LINUX
}

func IsDarwin() bool {
	return runtime.GOOS == NAME_OF_DARWIN
}
