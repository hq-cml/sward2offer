/*
 * 面试题19：正则表达式匹配
 * 题目：请实现一个函数用来匹配包含'.'和'*'的正则表达式。模式中的字符'.'
 * 表示任意一个字符，而'*'表示它前面的字符可以出现任意次（含0次）。在本题
 * 中，匹配是指字符串的所有字符匹配整个模式。例如，字符串"aaa"与模式"a.a"
 * 和"ab*ac*a"匹配，但与"aa.a"及"ab*a"均不匹配。
 */

package _19_simple_regexp

//思路：
//利用递归来实现状态机，需要想清楚各种情况
//遇到这种题目，也就只能欣赏欣赏了。
//难度：5*
func Regexp(str, pattern string) bool {
	if len(str) == 0 && len(pattern) == 0 {
		return true
	}

	if len(str) == 0 || len(pattern) == 0 {
		return false
	}

	return recurseReg([]byte(str), []byte(pattern))
}

//递归
func recurseReg(str, pattern []byte) bool {
	//退出条件，字符串与模式串同时结束，返回true
	if len(str) == 0 && len(pattern) == 0 {
		return true
	}

	//字符串未结束，模式串已经结束，则必然匹配失败
	if len(str) != 0 && len(pattern) == 0 {
		return false
	}

	//模式串不能直接是*打头，因为它标识它前面的字符串出现的次数
	if pattern[0] == '*' {
		return false
	}

	//判断是否模式串第二个字符是*
	if len(pattern) > 1 && pattern[1] == '*' {
		//模式串的第二个字符是 *，处理相对比较复杂
		if str[0] == pattern[0] || // 如果第一个字符相等
			(pattern[0] == '.' && len(str) > 0) { // 或者模式串第一个字符是.，这也等同于相等
			return recurseReg(str[1:], pattern[2:]) || // 字符串匹配一个，模式串跳过*
				recurseReg(str[1:], pattern) || // 字符串匹配一个，模式串维持原状，因为*可能表示字符
				recurseReg(str, pattern[2:]) // 字符串维持，模式串直接忽略*和器前面字符
		} else {
			return recurseReg(str, pattern[2:]) // 第一个字符不相等，则直接利用*表示0个字符的功效
		}
	} else {
		//模式串的第二个字符不是 *
		if len(str) > 0 && len(pattern) > 0 { // 必须保证此时字符串和模式串都未完结
			if str[0] == pattern[0] || // 如果第一个字符相等
				pattern[0] == '.' { // 或者模式串第一个字符是.，这也等同于相等
				return recurseReg(str[1:], pattern[1:]) // 同时向后平移
			}
		}
	}

	return false
}
