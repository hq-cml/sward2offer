/*
 * 面试题17：打印1到最大的n位数
 * 题目：输入数字n，按顺序打印出从1最大的n位十进制数。比如输入3，则
 * 打印出1、2、3一直到最大的3位数即999。
 */
package _17_print_1_to_n

import (
	"fmt"
	"strings"
)

//打印一个字符串，字符串全是数字，所以需要考虑0字符
func myPrint(s string) {
	s = strings.TrimLeft(s, "0")
	if len(s) == 0 {
		fmt.Println("0")
	}else {
		fmt.Println(s)
	}
}

//思路1：
//用字符串来模拟数字的打印。主要是考虑思维的全面性，因为要考虑大数的问题，所以就涉及到进位的问题，很复杂。
//利用字符串数字模拟加1算法，核心是判断什么时候到头
func PrintN_1(n int) {
	if n <= 0 {
		return
	}

	//初始化
	str := make([]byte, n)
	str[n-1] = '1'
	for i:=0; i<n-1; i++ {
		str[i] = '0'
	}

	//打印第一个数
	myPrint(string(str))
	for !mockIncr(str) { //未溢出，就一直打印
		myPrint(string(str))
	}
}

// 模拟字符串数字 + 1
// 返回值标志是否达到最大值溢出
func mockIncr(str []byte) bool {
	n := len(str)
	for i := n-1; i >= 0; i-- { //从个位数开始（数组的最后一位）
		if str[i] == '9' {
			if i > 0 {
				str[i] = '0'
			} else {
				return true
			}
		} else {
			str[i] = str[i] + 1
			break
		}
	}
	return false
}


//思路2：
//利用递归做全排列
//写出来之后，会感觉很简单
//实际写的过程中，非常烧脑，主要是字符串的低位和字符索引的低位正好是反的，
//所以一不小心就会出错，这种情况最好就是在纸上画出来，将idx标注出来
func PrintN_2(n int) {
	if n == 0 {
		return
	}
	str := make([]byte, n)
	recursePrint(str, 0)  //从数字的最高位（n位）开始
}

//牛逼，仅限于欣赏
func recursePrint(str []byte, idx int) {
	if idx == len(str)-1 {
		// 结束条件：
		// idx的最高位，正好是数字的个位数（最低位），0到9遍历
		for i:=0; i < 10; i++ {
			str[idx] = '0' + byte(i)
			myPrint(string(str))
		}
	} else {
		for i:=0; i < 10; i++ {
			str[idx] = '0' + byte(i)   //第i位（只要i不是个位），0到9遍历
			recursePrint(str, idx + 1) //递归下去
		}
	}
}