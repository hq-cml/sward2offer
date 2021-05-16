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