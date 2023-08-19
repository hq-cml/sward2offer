/*
 * yzh面试题：括号匹配
 * 给定一个字符串，判断是否是合法的括号匹配，括号有3种：(), [], {}
 * PS：同时也是 leetcode 20
 */
package _bracket_match

// 括号匹配
func Match(str string) bool {
	if len(str) == 0 {
		return true
	}
	var stk []byte
	for _, c := range []byte(str) {
		switch c {
		case '(', '[', '{':
			stk = append(stk, c) // push
		case ')', ']', '}':
			if len(stk) == 0 {
				return false
			}
			top := stk[len(stk) - 1]
			if (top == '(' && c == ')' ) ||
				(top == '[' && c == ']' ) ||
				(top == '{' && c == '}' ) {
				stk = stk[:len(stk)-1] // pop
			} else {
				return false
			}
		default:
			return false
		}
	}

	//特殊case比如：{{[[
	if len(stk) > 0 {
		return false
	}
	return true
}
