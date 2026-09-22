package xnum

import (
	"fmt"
	"math"
	"strings"
)

type NumConverter interface {
	Encode(int64) string
	Decode(string) (int64, error)
	MustDecode(string) int64
}

type numConvert struct {
	runes []rune
	nmaps map[rune]int
}

func newNumConvert(convert_type string) *numConvert {
	r := &numConvert{
		runes: []rune(convert_type),
	}
	r.nmaps = make(map[rune]int, len(r.runes))
	for i := 0; i < len(r.runes); i++ {
		r.nmaps[r.runes[i]] = i
	}
	return r
}

func (e *numConvert) Encode(num int64) string {
	sign := ""
	if num < 0 {
		sign = "-"
		num *= -1
	}

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
	return fmt.Sprintf("%s%s", sign, string(result))
}

func (e *numConvert) Decode(val string) (int64, error) {
	var (
		base              = int64(len(e.runes))
		r           int64 = 0
		err         error
		is_negative = strings.Index(val, "-") == 0
		vrunes      = []rune(strings.TrimLeft(val, "-"))
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
	if is_negative {
		r *= -1
	}
	return r, err
}

func (e *numConvert) MustDecode(val string) int64 {
	if result, err := e.Decode(val); err != nil {
		panic(err)
	} else {
		return result
	}
}

func NewConvertTo2() NumConverter {
	return newNumConvert("01")
}

func NewConvertTo8() NumConverter {
	return newNumConvert("01234567")
}

func NewConvertTo10() NumConverter {
	return newNumConvert("0123456789")
}

func NewConvertTo16Lower() NumConverter {
	return newNumConvert("0123456789abcdef")
}

func NewConvertTo16Upper() NumConverter {
	return newNumConvert("0123456789ABCDEF")
}

func NewConvertTo26Lower() NumConverter {
	return newNumConvert("abcdefghijklmnopqrstuvwxyz")
}

func NewConvertTo26Upper() NumConverter {
	return newNumConvert("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
}

func NewConvertTo32Lower() NumConverter {
	return newNumConvert("0123456789abcdefghjkmnpqrstvwxyz")
}

func NewConvertTo32Upper() NumConverter {
	return newNumConvert("0123456789ABCDEFGHJKMNPQRSTVWXYZ")
}

func NewConvertTo36Lower() NumConverter {
	return newNumConvert("0123456789abcdefghijklmnopqrstuvwxyz")
}

func NewConvertTo36Upper() NumConverter {
	return newNumConvert("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ")
}
func NewConvertTo52() NumConverter {
	return newNumConvert("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
}
func NewConvertTo58() NumConverter {
	return newNumConvert("123456789abcdefghijkmnopqrstuvwxyzABCDEFGHJKLMNPQRSTUVWXYZ")
}
func NewConvertTo62() NumConverter {
	return newNumConvert("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
}
