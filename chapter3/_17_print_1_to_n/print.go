/*
// 面试题17：打印1到最大的n位数
// 题目：输入数字n，按顺序打印出从1最大的n位十进制数。比如输入3，则
// 打印出1、2、3一直到最大的3位数即999。

这个题我第一次没搞出来，第二次瘦了点启发，才勉强搞出来。
这种题主要是考虑思维的全面性，因为要考虑大数的问题，所以需要考虑到用字符串来模拟数字的打印。
所以就涉及到进位的问题，很复杂。
这道题作者还给了一个思路，是用递归做全排列，但是我没太看懂，后面再说吧。
 */

package _17_print_1_to_n

import (
	"fmt"
	"strings"
)

//利用字符串数字模拟加1算法
//核心是判断什么时候到头
func PrintN_1(n int) {
	if n <= 0 {
		return
	}
	str := make([]byte, n)
	str[n-1] = '1'
	for i:=0; i<n-1; i++ {
		str[i] = '0'
	}

	myPrint(string(str))
	for !mockIncr(str) {
		myPrint(string(str))
	}
}

func myPrint(s string) {
	s = strings.TrimLeft(s, "0")
	if len(s) == 0 {
		fmt.Println("0")
	}else {
		fmt.Println(s)
	}
}

// 返回值标志是否溢出
func mockIncr(str []byte) bool {
	n := len(str)
	for i := n-1; i >= 0; i-- {
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

//利用递归实现
//写出来之后，会感觉很简单
//实际写的过程中，非常烧脑，主要是字符串的低位和字符索引的低位
//正好是反的，所以一不小心就会出错，这种情况最好就是在纸上画出来，将idx标注出来
func PrintN_2(n int) {
	if n == 0 {
		return
	}
	str := make([]byte, n)
	recursePrint(str, 0)
}

func recursePrint(str []byte, idx int) {
	if idx == len(str)-1 {
		// 结束条件：
		// idx的最高位，正好是数字的最低位
		for i:=0; i < 10; i++ {
			str[idx] = '0' + byte(i)
			myPrint(string(str))
		}
	} else {
		for i:=0; i < 10; i++ {
			str[idx] = '0' + byte(i)
			recursePrint(str, idx + 1)
		}
	}
}