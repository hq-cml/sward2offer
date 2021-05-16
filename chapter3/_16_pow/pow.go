package _16_pow

import (
	"fmt"
	"math"
)

const (
	MIN = 0.0000001
)

func equalZero(f float64) bool {
	if math.Abs(f) <= MIN {
		return true
	}
	return false
}

//常规思路
func Pow1(f float64, exp int) (float64, error) {
	if exp == 0 {
		return 1, nil
	}

	if equalZero(f) {
		if exp > 0 {
			return 0, nil
		}
		return 0, fmt.Errorf("exp should not be 0 when f == 0")
	}

	minus := false
	if exp < 0 {
		minus = true
		exp = exp * -1
	}

	var r float64 = 1
	for exp > 0 {
		r *= f
		exp--
	}

	if minus {
		return 1/r, nil
	} else {
		return r, nil
	}
}

//效率比Pow1更高
func Pow2(f float64, exp int) (float64, error) {
	if exp == 0 {
		return 1, nil
	}

	if equalZero(f) {
		if exp > 0 {
			return 0, nil
		}
		return 0, fmt.Errorf("exp should not be 0 when f == 0")
	}

	minus := false
	if exp < 0 {
		minus = true
		exp = exp * -1
	}

	r := powRecurse(f, exp)

	if minus {
		return 1/r, nil
	} else {
		return r, nil
	}
}

func powRecurse(f float64, exp int) float64 {
	if exp == 0 {
		return 1
	}
	if exp == 1 {
		return f
	}

	r := powRecurse(f, exp>>1)
	r = r*r
	if exp & 1 == 1 {
		return r*f
	}
	return r
}