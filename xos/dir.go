package xos

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/ynwcel/gox/xfs"
)

func MustDeepGlob(dir, pattern string) []string {
	if result, err := DeepGlob(dir, pattern); err != nil {
		panic(err)
	} else {
		return result
	}
}

func DeepGlob(dir, pattern string) ([]string, error) {
	if result, err := xfs.DeepGlob(os.DirFS(dir), pattern); err != nil {
		return nil, err
	} else {
		fmt_result := make([]string, len(result))
		for idx, file := range result {
			fmt_result[idx] = filepath.Clean(fmt.Sprintf("%s/%s", dir, file))
		}
		return fmt_result, nil
	}
}
