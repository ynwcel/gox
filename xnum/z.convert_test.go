package xnum

import (
	"testing"
)

func TestConvert(t *testing.T) {
	var (
		nums = []int{
			0,
			1,
			2,
			3,
			4,
			7,
			8,
			10,
			16,
			32,
			1024,
			-1024,
		}
		converts = []map[string]NumConverter{
			{"to2": NewConvertTo2()},
			{"to8": NewConvertTo8()},
			{"to16lower": NewConvertTo16Lower()},
			{"to16uppper": NewConvertTo16Upper()},
		}
	)
	for _, num := range nums {
		for _, info := range converts {
			for name, convert := range info {
				target := convert.Encode(int64(num))
				if source, err := convert.Decode(target); err != nil {
					t.Error(err)
				} else {
					t.Log(num, name, "encode =", target, "decode=", source)
				}
			}
		}
	}
}
