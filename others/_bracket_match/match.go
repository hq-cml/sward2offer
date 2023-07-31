/*
 * yzh面试题：括号匹配
 * 给定一个字符串，判断是否是合法的括号匹配，括号有3种：(), [], {}
 * PS：同时也是 leetcode 20
 */
package _bracket_match

import (
	"github.com/hq-cml/sward2offer/common"
)

func Match(str string) bool {
	stk := common.NewStack(false)
	for _, c := range str {
		switch c {
		case '{', '[', '(':
			stk.Push(byte(c))
		case ')', ']', '}':
			if stk.Len() == 0 {
				return false
			}
			top, _ := stk.Top()
			tc := top.(byte)
			if (tc == '(' && c == ')') ||
				(tc == '[' && c == ']') ||
				(tc == '{' && c == '}') {
				stk.Pop()
			} else {
				return false
			}
		}
	}

	if stk.Len() > 0 {
		return false //比如：{{[[
	}
	return true
}
