/*
 * 面试题64：求1+2+…+n
 * 题目：求1+2+…+n，要求不能使用乘除法、for、while、if、else、switch、case
 * 等关键字及条件判断语句（A ? B : C）
 */
package _64_sum_1_to_n

import "math"

//思路：
//这题的第一反应，肯定是递归，但是递归需要退出条件，由于不能用if else，所以难点在这儿

//思路1：
//利用等差数列求和：
//1+2+..+n = (1+n)n/2 = (n^2 + n) / 2
//平方用pow函数代替，除以2用右移一位代替
//难度：4*
func Sum1ton_ArithmeticProgression(n int) int {
	return (int(math.Pow(float64(n), 2)) + n) >> 1
}

//递归
//利用&&的特性
//难度：4*
func Sum1ton_Recurse(n int) int {
	sum := 0
	add(&sum, n)
	return sum
}

//用sum指针来存储最后值，不能作为返回值
//必须将bool作为唯一返回值，用于&&操作
func add(sum *int, n int) bool {
	*sum += n
	return n > 0 && add(sum, n-1) //只要n>=1，就继续递归下去
}
