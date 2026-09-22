package xos

import (
	"io"
	"os"
)

func MustCopyFile(src string, target string) bool {
	if result, err := CopyFile(src, target); err != nil {
		panic(err)
	} else {
		return result
	}
}

func CopyFile(src string, target string) (bool, error) {
	rfile, err := os.Open(src)
	if err != nil {
		return false, err
	}
	if _, err := PutContent(target, rfile); err != nil {
		return true, nil
	} else {
		return false, err
	}
}

func MustPutContent(filename string, content io.Reader) int64 {
	if result, err := PutContent(filename, content); err != nil {
		panic(err)
	} else {
		return result
	}
}

func PutContent(filename string, content io.Reader) (int64, error) {
	ofile, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, os.ModePerm)
	if err != nil {
		return 0, err
	}
	defer ofile.Close()
	return io.Copy(ofile, content)
}
