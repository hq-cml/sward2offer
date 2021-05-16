package _17_print_1_to_n

import (
	"fmt"
	"strings"
)

//利用字符串数字模拟加1算法
//核心是判断什么时候到头
func PrintN_1(n int) {
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
	fmt.Println(strings.TrimLeft(s, "0"))
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
