/*
*

	*给定一个经过编码的字符串，返回它解码后的字符串。

编码规则为: k[encoded_string]，表示其中方括号内部的 encoded_string 正好重复 k 次。注意 k 保证为正整数。

你可以认为输入字符串总是有效的；输入字符串中没有额外的空格，且输入的方括号总是符合格式要求的。

此外，你可以认为原始数据不包含数字，所有的数字只表示重复的次数 k ，例如不会出现像 3a 或 2[4] 的输入。

示例 1：

输入：s = "3[a]2[bc]"
输出："aaabcbc"
示例 2：

输入：s = "3[a2[c]]"
输出："accaccacc"
示例 3：

输入：s = "2[abc]3[cd]ef"
输出："abcabccdcdcdef"
示例 4：

输入：s = "abc3[cd]xyz"
输出："abccdcdcdxyz"
*/
package _394_string_decode

import "fmt"

func DecodeString(s string) string {
	cnt, headLhen, beg, end := findBracket(s)
	if beg == -1 && end == -1 {
		return s
	}
	fmt.Println(cnt, headLhen, beg, end)
	subStr := s[beg+1 : end]
	left := s[:beg-headLhen]
	right := s[end+1:]
	newStr := left
	for i := 0; i < cnt; i++ {
		newStr += subStr
	}
	newStr += right
	//fmt.Println(newStr)
	return DecodeString(newStr)
}

func findBracket(str string) (int, int, int, int) {
	stk := []byte{}
	beg := -1
	end := -1
	for idx, c := range []byte(str) {
		if c == '[' {
			if len(stk) == 0 {
				beg = idx
			}
			stk = append(stk, c) // push
		} else if c == ']' {
			stk = stk[:len(stk)-1] // pop
			if len(stk) == 0 {
				end = idx
				break
			}
		}
	}

	// 计算长度
	cnt := 0
	headLen := 0
	for i := beg - 1; i >= 0; i-- {
		if str[i] < '0' || str[i] > '9' {
			break
		}
		headLen++
	}
	for i := beg - headLen; i < beg; i++ {
		cnt = cnt*10 + int(str[i]-'0')
	}
	return cnt, headLen, beg, end
}
