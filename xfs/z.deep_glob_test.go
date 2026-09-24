package xfs

import (
	"os"
	"testing"
)

func TestDeepGlob(t *testing.T) {
	if files, err := DeepGlob(os.DirFS("../"), `*_test.go`); err != nil {
		t.Error(err)
	} else {
		for _, f := range files {
			t.Log(f)
		}
	}
}
