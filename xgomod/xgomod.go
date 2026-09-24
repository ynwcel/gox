package xgomod

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const (
	FILE_NAME = "go.mod"
)

func GetName() (string, error) {
	var (
		gomod_file    = filepath.Clean(fmt.Sprintf("%s/%s", filepath.Dir("."), FILE_NAME))
		gomod_content []byte
		gomod_name    string
		err           error
	)
	if gomod_content, err = os.ReadFile(gomod_file); err != nil {
		return gomod_name, fmt.Errorf("read go.mod failed:%w", err)
	}
	lines := bytes.Split(gomod_content, []byte("\n"))
	for _, line := range lines {
		if strings.Index(string(line), "module ") == 0 {
			gomod_name = string(bytes.TrimSpace(bytes.Split(line, []byte("module "))[1]))
			break
		}
	}
	return gomod_name, nil
}

func MustGetGoMod() string {
	if v, err := GetName(); err != nil {
		panic(err)
	} else {
		return v
	}
}
