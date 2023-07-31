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
//大概意思是，每一位都进行递归运算，分别算出每一位可能出现1的次数
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
	if length == 1 && firstChar > 0 { // 只有1位，且>0，则必然有一个1
		return 1
	}

	numFirstDigit := 0
	// firstChar不可能是0，只会大于0的某个数字
	if firstChar == 1 {
		// 如果==1，则该位出现1的个数，必然是下面所有数的和（例如：123，则出现23+1次）
		t, _ := strconv.Atoi(string(str[1:]))
		numFirstDigit = t + 1
	} else { //firstChar > 1
		// 如果>1，则该位出现1的个数，是10的len-1次方哥（例如：323，则出现100-199，共100次）
		numFirstDigit = int(math.Pow10(length - 1))
	}

	// 除了最高位的1，其他位的1的个数
	// 类似于上面的else逻辑，有点烧脑，例如：323，则会出现 3*2*10次的1
	numOtherDigit := int(firstChar) * (length - 1) * int(math.Pow10(length-2))

	return numFirstDigit +  // 最高位
		numOtherDigit +     // 其他位
		numberOf1(str[1:])  // 递归向下一位
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
