package pkg

import (
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
)

func ListFile(dirpath string, pattern string) ([]string, error) {
	var (
		dirfs       = os.DirFS(dirpath)
		file_regexp *regexp.Regexp
		curfiles    []string
		result      []string
		err         error
	)
	if file_regexp, err = regexp.Compile(pattern); err != nil {
		return nil, err
	}
	if curfiles, err = fs.Glob(dirfs, "*"); err != nil {
		return nil, err
	}
	result = make([]string, 0, len(curfiles))
	for _, file := range curfiles {
		fpath := filepath.Join(dirpath, file)
		if finfo, err := fs.Stat(dirfs, file); err != nil {
			return nil, err
		} else {
			if file_regexp.Match([]byte(finfo.Name())) {
				result = append(result, fpath)
			}
			if finfo.IsDir() {
				if subfs, err := ListFile(fpath, pattern); err != nil {
					return nil, err
				} else {
					result = append(result, subfs...)
				}
			}
		}
	}
	return result, nil
}
