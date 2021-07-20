/*
// 面试题64：求1+2+…+n
// 题目：求1+2+…+n，要求不能使用乘除法、for、while、if、else、switch、case
// 等关键字及条件判断语句（A?B:C）。

这题如果可以用三目运算符，那么很简单，直接递归，关键是连三目运算符也不让用。。。
作者给的方法都是利用了C++特性，唯一一个可以接受的是利用函数指针的，但这也其实用到了C语言的0的取反是1的特性，不通用
 */
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