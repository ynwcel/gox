package gnumx

import (
	"testing"
	"time"
)

func testConvert(t *testing.T, data int64) {
	types := map[string]string{
		"ConvertType2":       ConvertType2,
		"ConvertType8":       ConvertType8,
		"ConvertType16Lower": ConvertType16Lower,
		"ConvertType16Upper": ConvertType16Upper,
		"ConvertType26Lower": ConvertType26Lower,
		"ConvertType26Upper": ConvertType26Upper,
		"ConvertType32Lower": ConvertType32Lower,
		"ConvertType32Upper": ConvertType32Upper,
		"ConvertType36Lower": ConvertType36Lower,
		"ConvertType36Upper": ConvertType36Upper,
		"ConvertType52":      ConvertType52,
		"ConvertType58":      ConvertType58,
		"ConvertType62":      ConvertType62,
	}
	for tpe_name, tpe := range types {
		e := NewConvert(tpe)
		v := e.Encode(data)
		r, err := e.Decode(v)
		if err != nil {
			t.Error(err)
		}
		t.Logf("[%-20s],source=%-20d,encode=%-20s,decode=%-20d", tpe_name, data, v, r)
	}
}

func TestAll(t *testing.T) {
	datas := []int64{
		0, 1, 2, 4, 8,
		16, 32, 36, 52, 58, 62,
		202308070950,
		1257894000,
		time.Now().UnixNano(),
	}
	for _, v := range datas {
		testConvert(t, v)
	}
}
