package gnumx

import (
	"fmt"
	"math"
)

const (
	ConvertType2       = "01"
	ConvertType8       = "01234567"
	ConvertType10      = "0123456789"
	ConvertType16Lower = "0123456789abcdef"
	ConvertType16Upper = "0123456789ABCDEF"
	ConvertType26Lower = "abcdefghijklmnopqrstuvwxyz"
	ConvertType26Upper = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	ConvertType32Lower = "0123456789abcdefghjkmnpqrstvwxyz"
	ConvertType32Upper = "0123456789ABCDEFGHJKMNPQRSTVWXYZ"
	ConvertType36Lower = "0123456789abcdefghijklmnopqrstuvwxyz"
	ConvertType36Upper = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	ConvertType52      = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	ConvertType58      = "123456789abcdefghijkmnopqrstuvwxyzABCDEFGHJKLMNPQRSTUVWXYZ"
	ConvertType62      = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

type Convert struct {
	runes []rune
	nmaps map[rune]int
}

func NewConvert(convert_type string) *Convert {
	r := &Convert{
		runes: []rune(convert_type),
	}
	r.nmaps = make(map[rune]int, len(r.runes))
	for i := 0; i < len(r.runes); i++ {
		r.nmaps[r.runes[i]] = i
	}
	return r
}

func (e *Convert) Encode(num int64) string {
	base := int64(len(e.runes))
	r := make([]rune, 0, 64)

	for {
		s, y := num/int64(base), num%int64(base)
		r = append(r, e.runes[y])
		if s < int64(base) {
			if s != 0 {
				r = append(r, e.runes[s])
			}
			break
		} else {
			num = s
		}
	}
	result := make([]rune, len(r))
	for i := len(r) - 1; i >= 0; i-- {
		result[len(r)-i-1] = r[i]
	}
	return string(result)
}

func (e *Convert) Decode(val string) (int64, error) {
	var (
		base         = int64(len(e.runes))
		vrunes       = []rune(val)
		r      int64 = 0
		err    error
	)
	for i, vlen := 0, len(vrunes); i < vlen; i++ {
		cur_rune := vrunes[i]
		if num, ok := e.nmaps[cur_rune]; ok {
			r = r + int64(num)*int64(math.Pow(float64(base), float64(vlen)-float64(i+1)))
		} else {
			r = 0
			err = fmt.Errorf("invalid letter `%s`", string(cur_rune))
			break
		}
	}
	return r, err
}
