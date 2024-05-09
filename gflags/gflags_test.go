package gflags

import (
	"testing"
)

var test_flags = []string{}

func TestRun(t *testing.T) {
	var (
		int_flag  int
		str_flag  string
		bool_flag bool
		strs_flag []string
		cmd       = New().SetVersion("0.0.1")
	)
	cmd.IntVarP(&int_flag, "int", "i", 0, "set int_flag value")
	cmd.StringVarP(&str_flag, "str", "s", "", "set str_flag value")
	cmd.BoolVarP(&bool_flag, "bool", "b", false, "set bool_flag value")
	cmd.StringSliceVar(&strs_flag, "strs", []string{}, "set str slices value")
	if err := cmd.Parse(test_flags); err != nil {
		t.Error(err)
	}
	if cmd.HasSetHelpFlag() {
		cmd.Usage()
	}
}

func TestMain(m *testing.M) {
	test_flags = append(test_flags, "--help")
	m.Run()
}
