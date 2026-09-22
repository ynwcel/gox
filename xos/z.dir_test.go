package xos

import "testing"

func TestDeepGlobV1(t *testing.T) {
	if files, err := DeepGlob("../", `*_test.go`); err != nil {
		t.Error(err)
	} else {
		for _, f := range files {
			t.Log(f)
		}
	}
}
func TestDeepGlobV2(t *testing.T) {
	if files, err := DeepGlob(".", `*.go*`); err != nil {
		t.Error(err)
	} else {
		for _, f := range files {
			t.Log(f)
		}
	}
}
