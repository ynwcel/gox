package pkg

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func GetGoModName() (string, error) {
	var (
		gomod_file    = fmt.Sprintf("%s/go.mod", filepath.Dir("."))
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
