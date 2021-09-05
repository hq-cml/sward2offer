/*
 * 面试题15：二进制中1的个数
 * 题目：请实现一个函数，输入一个整数，输出该数二进制表示中1的个数。例如
 * 把9表示成二进制是1001，有2位是1。因此如果输入9，该函数输出2。
 */
package _15_binary_1_cnt

//三个思路：
//1. 最常规的是就是不断右移和1进行&操作，但是如果是有符号的负数，就会死循环
//2. 换一种思路，不断左移1，然后与数字进行与操作，这样不会死循环，但是效率略低
//3. 神思路，不断减1做&运算，效率高
//难度：2*

//利用1不断左移，按位与操作，统计个数
func Calc1(n int64) int {
	var x int64 = 1
	cnt := 0
	for x != 0 {
		if x&n != 0 {
			cnt++
		}
		x = x << 1
	}
	return cnt
}

//利用减1后与自己操作，速度更快
func Calc2(n int64) int {
	cnt := 0
	for n != 0 {
		cnt++
		n = n & (n - 1)
	}
	return cnt
}
