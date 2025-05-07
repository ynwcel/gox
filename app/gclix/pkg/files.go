package pkg

import (
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
)

func FileExists(filename string) bool {
	if _, err := os.Stat(filename); err != nil && os.IsNotExist(err) {
		return false
	}
	return true
}

func IsDir(dirname string) bool {
	if state, err := os.Stat(dirname); err == nil && state.IsDir() {
		return true
	}
	return false
}

func PutContent(filename string, content string) (int, error) {
	ofile, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		return 0, err
	}
	defer ofile.Close()
	return fmt.Fprint(ofile, content)
}

func CopyFile(src, target string) (bool, error) {
	ofile, err := os.Open(src)
	if err != nil {
		return false, err
	}
	defer ofile.Close()
	wfile, err := os.OpenFile(target, os.O_CREATE|os.O_WRONLY, 0766)
	if err != nil {
		return false, err
	}
	defer wfile.Close()
	if _, err := io.Copy(wfile, ofile); err == nil {
		return true, nil
	} else {
		return false, err
	}
}

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
