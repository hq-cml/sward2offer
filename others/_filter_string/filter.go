/*
 * mt面试题：一个字符串，递归的去除b，再递归去除'ac'
 * 比如：
 * acbdc => acdc => dc;
 * abcd => acd => d;
 * aabccx => aaccx => acx => x;。
 */
package _filter_string

import "github.com/hq-cml/sward2offer/common"

//利用栈，来实现递归的消除效果
func Filter(str string) string {
	mystack := common.NewStack(false)
	for _, c := range str {
		if c == 'b' { //直接去除
			continue
		}
		if mystack.Len() == 0 {
			mystack.Push(byte(c))
		} else {
			t, _ := mystack.Top()
			tc := t.(byte)
			if tc == 'a' && c == 'c' {
				mystack.Pop()
			} else {
				mystack.Push(byte(c))
			}
		}
	}

	//此时mystack已经是过滤完毕的字符了，但是倒序
	if mystack.Len() == 0 {
		return ""
	} else {
		//将栈中元素，倒过来返回，利用另一个辅助栈
		tmpStack := common.NewStack(false)
		for mystack.Len() != 0 {
			c, _ := mystack.Pop()
			tmpStack.Push(c.(byte))
		}
		s := []byte{}
		for tmpStack.Len() != 0 {
			cc, _ := tmpStack.Pop()
			c := cc.(byte)
			s = append(s, c)
		}
		return string(s)
	}
}
