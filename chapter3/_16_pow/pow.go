/*
 * 面试题16：数值的整数次方
 * 题目：实现函数double Power(double base, int exponent)，求base的exponent
 * 次方。不得使用库函数，同时不需要考虑大数问题。
 */
package _16_pow

import (
	"fmt"
	"math"
)

const (
	MIN = 0.0000001
)

// 浮点数判断0（利用绝对值）
func equalZero(f float64) bool {
	if math.Abs(f) <= MIN {
		return true
	}
	return false
}

//分析：主要考察考虑场景的完整性：比如底数为0，或者指数为负数的情况。
//常规思路
func Pow1(f float64, exp int) (float64, error) {
	// 指数为0，返回1
	if exp == 0 {
		return 1, nil
	}

	//底数为0
	if equalZero(f) {
		if exp > 0 {
			return 0, nil
		}
		return 0, fmt.Errorf("exp should not be 0 when f == 0")
	}

	//指数是否为正
	minus := false
	if exp < 0 {
		minus = true
		exp = exp * -1
	}

	//计算
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
//利用递归的方式来提速
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

//原理：
// 如果exp=32，则其实可以写成 (f^16)^2 = ((f^8)^2)^2 ...递归
// 如果exp=31，则其实可以写成 (f^15)^2 * f = ((f^7)^2* f)^2 * f ...递归
// 更加抽象的，可以将exp抽象成为奇数或者偶数
// 这样，将乘法的次数，对数递减logN
func powRecurse(f float64, exp int) float64 {
	//递归结束条件
	if exp == 0 {
		return 1
	}
	if exp == 1 {
		return f
	}

	r := powRecurse(f, exp>>1)  //向下递归
	r = r*r           // 求平方
	if (exp & 1) == 1 { // exp 是奇数，多乘一个r
		return r*f
	} else {
		return r      // exp 是偶数
	}
}