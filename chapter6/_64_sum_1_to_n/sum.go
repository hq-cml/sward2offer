package _64_sum_1_to_n

import "math"

//利用等差数列求和：
//1+2+..+n = (1+n)n/2 = (n^2 + n) / 2
//平方用pow函数代替，/2用右移一位代替
func Sum1ton_ArithmeticProgression(n int) int {
	if n < 1 {
		return 0
	}
	return (int(math.Pow(float64(n), 2)) + n ) >> 1
}

//递归
//利用&&的特性
func Sum1ton_Recurse(n int) int {
	if n < 1 {
		return 0
	}
	sum := 0
	var r func(sum *int, n int) bool
	r = func(sum *int, n int) bool { //匿名函数
		*sum += n
		return n>0 && r(sum, n-1)
	}

	r(&sum, n)
	return sum
}