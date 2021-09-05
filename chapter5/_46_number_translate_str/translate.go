/*
 * 面试题46：把数字翻译成字符串
 * 题目：给定一个数字，我们按照如下规则把它翻译为字符串：0翻译成"a"，1翻
 * 译成"b"，……，11翻译成"l"，……，25翻译成"z"。一个数字可能有多个翻译。例
 * 如12258有5种不同的翻译，它们分别是"bccfi"、"bwfi"、"bczi"、"mcfi"和
 * "mzi"。请编程实现一个函数用来计算一个数字有多少种不同的翻译方法。
 */
package _46_number_translate_str

import (
	"errors"
	"strconv"
)

//思路：
//这是一个递归思想，有点类似于斐波那契数列
//难度：5*

//0-a; 25-z
func Translate(n int) (int, error) {
	if n < 0 {
		return 0, errors.New("Invalid n")
	}

	s := strconv.Itoa(n) //先转成字符串
	return translateRecurse(s), nil
}

//思路1：递归
//f(i) = f(i+1) + p(i, i+1) * f(i+2)
//其中
//   f(i)表示长度为i的字符串的可能性数量
//   p(i, i+1)表示第i和i+1位合并之后，是否在[10, 25]之间，在则 == 1，不在则 == 0
//原因：
//   如果是"05123"，则只能跨一步；如果是"35123"，同样只能跨一步；
//   如果是"15123"，则可能跨1步，也可能跨2步
func translateRecurse(org string) int {
	//递归的退出条件
	if org == "" {
		return 0 //可能性数
	}
	if len(org) == 1 {
		return 1
	}
	if len(org) == 2 {
		if check(org) {
			return 2 //10-25之间，则可能性是2种
		} else {
			return 1 //否则，只有一种可能
		}
	}

	//因子
	//如果前两个字符在10-25之间，factor==1，否则factor==0
	factor := 1
	if !check(org[:2]) {
		factor = 0
	}

	//其实，根据公式，本质上就是想得到f(0)
	f0 := translateRecurse(org[1:]) + factor*translateRecurse(org[2:])
	return f0
}

//思路2：利用循环代替递归
//递归的方式相对清晰，但是却存在很多重复计算，导致性能问题
//利用循环可以自底向上解决问题，效率更高
//利用一个数组来缓存计算的各个中间结果
func translateLoop(org string) int {
	length := len(org)

	if length == 1 {
		return 1
	}
	if length == 2 {
		if check(org) {
			return 2 //10-25之间，则可能性是2种
		} else {
			return 1 //否则，只有一种可能
		}
	}

	numbers := make([]int, length)
	numbers[length-1] = 1 //递归结束条件，在循环中作为边界条件
	numbers[length-2] = func() int {
		if check(org[length-2:]) { //如果还剩两位，有两种可能
			return 2
		} else {
			return 1
		}
	}()

	//循环式倒过来循环的
	//从最后一位开始
	for idx := length - 3; idx >= 0; idx-- {
		if check(org[idx : idx+2]) {
			numbers[idx] = numbers[idx+1] + numbers[idx+2]
		} else {
			numbers[idx] = numbers[idx+1]
		}
	}

	return numbers[0]
}

//判断一个2字符的串，是否在[10, 25]之间
func check(str string) bool {
	if len(str) != 2 {
		return false
	}
	i, _ := strconv.Atoi(str)
	if i >= 10 && i <= 25 {
		return true
	} else {
		return false
	}
}
