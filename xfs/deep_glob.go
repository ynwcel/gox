package xfs

import (
	"fmt"
	"io/fs"
)

func MustDeepGlob(dirfs fs.FS, pattern string) []string {
	if result, err := DeepGlob(dirfs, pattern); err != nil {
		panic(err)
	} else {
		return result
	}
}

func DeepGlob(dirfs fs.FS, pattern string) ([]string, error) {
	var (
		curfiles []string
		result   []string
		err      error
	)
	if curfiles, err = fs.Glob(dirfs, pattern); err != nil {
		return nil, err
	} else {
		result = append(result, curfiles...)
	}
	if subfiles, err := fs.Glob(dirfs, `*`); err != nil {
		return nil, err
	} else {
		for _, entry := range subfiles {
			finfo, err := fs.Stat(dirfs, entry)
			if err != nil {
				return nil, err
			} else if !finfo.IsDir() {
				continue
			}
			if sub_fs, err := fs.Sub(dirfs, finfo.Name()); err != nil {
				return nil, err
			} else if sub_files, err := DeepGlob(sub_fs, pattern); err != nil {
				return nil, err
			} else {
				for _, path := range sub_files {
					result = append(result, fmt.Sprintf("%s/%s", finfo.Name(), path))
				}
			}
		}
	}
	return result, nil
}
