/*
 * 面试题44：数字序列中某一位的数字
 * 题目：数字以0123456789101112131415…的格式序列化到一个字符序列中。在这
 * 个序列中，第5位（从0开始计数）是5，第13位是1，第19位是4，等等。请写一
 * 个函数求任意位对应的数字。
 */
package _44_digit_in_sequence

import "math"

//思路：
//观察：1位数，占10个字符 => 2位数，占180个 => 3位数，占2700个。。。。
//先计算出第n位数组属于几位数，然后在从这个位数从头开始算。。
//这个题目相当恶心，因为idx从0开始计数，所以边界非常烧脑，要保持idx的计数方式统一，否则会乱掉
//难度：5*
func GetDigit(idx int) int {
	n := 1 //n标识几位数

	//首先判断idx位于几位数
	for {
		if idx-getWidthOfN(n) == 0 {
			//正好位于n位数的第一个数字的最高位，比如idx=10,190,2890...
			//烧脑，idx从0开始，所以这里其实是真正字符串的第11,191,2891个字符
			return 1
		} else if idx-getWidthOfN(n) > 0 {
			idx -= getWidthOfN(n)
			n++
		} else {
			break
		}
	}

	//还剩下的可用的位数，idx仍然从0开始计数，必须要统一
	num := idx / n //相隔完整的num个数字
	idx = idx % n
	var beg int //n位数的开始
	if n == 1 {
		beg = 0
	} else {
		beg = int(math.Pow10(n - 1))
	}
	beg = beg + num

	//到这里
	//就是求beg这个n位数字的第idx位，idx从0开始计数
	cnt := n - idx //倒着求
	ret := 0
	for cnt > 0 {
		ret = beg % 10
		beg = beg / 10
		cnt--
	}

	return ret
}

//获取N位数的宽度
//比如1位数：10；2位数：180；3位数：2700...
func getWidthOfN(n int) int {
	if n < 1 {
		return 0
	}

	if n == 1 {
		return 10
	}

	return int(math.Pow10(n)-math.Pow10(n-1)) * n
}
