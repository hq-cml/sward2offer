/*
 * 面试题43：从1到n整数中1出现的次数
 * 题目：输入一个整数n，求从1到n这n个整数的十进制表示中1出现的次数。例如
 * 输入12，从1到12这些整数中包含1 的数字有1，10，11和12，1一共出现了5次。
 */
package _43_numberof1_between_1_n

import (
	"math"
	"strconv"
)

//思路1：
//复杂度O(logn)
//这种题，非常抽象，只能欣赏欣赏了
//难度:6*
func Calc(n int) int {
	if n <= 0 {
		return 0
	}

	return numberOf1([]byte(strconv.Itoa(n)))
}

//递归
//TODO 看不懂了
func numberOf1(str []byte) int {
	//退出条件
	if len(str) == 0 {
		return 0
	}

	//str[0]是数字最高位
	firstChar := str[0] - '0'
	length := len(str)

	//递归退出条件
	if length == 1 && firstChar == 0 { //字符串只剩下"0"，退出
		return 0
	}
	if length == 1 && firstChar > 0 { // why?
		return 1
	}

	numFirstDigit := 0
	if firstChar == 1 {
		t, _ := strconv.Atoi(string(str[1:]))
		numFirstDigit = t + 1
	} else { //firstChar > 1
		numFirstDigit = int(math.Pow10(length - 1))
	}

	numOtherDigit := int(firstChar) * (length - 1) * int(math.Pow10(length-2))

	numRecurse := numberOf1(str[1:])

	return numFirstDigit + numOtherDigit + numRecurse
}

//思路2：
//朴素思路：复杂度为O(N^2)，用于验证Calc1
func Calc2(n int) int {
	calc := func(n int) int { //匿名小函数，计算一个数字中1的
		num := 0
		for n != 0 {
			if n%10 == 1 {
				num++
			}
			n = n / 10
		}
		return num
	}

	//循环计算，将1的个数累积
	num := 0
	for n != 0 {
		num += calc(n)
		n--
	}
	return num
}
