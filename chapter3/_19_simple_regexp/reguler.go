package _19_simple_regexp

func Regexp(str, pattern string) bool {
	if len(str) == 0 || len(pattern) == 0 {
		return false
	}

	return recurseReg([]byte(str), []byte(pattern))
}

func recurseReg(str, pattern []byte) bool {
	if len(str) == 0 && len(pattern) == 0 {
		return true
	}

	//字符串已经结束，模式串还未结束，则必然匹配失败
	if len(str) != 0 && len(pattern) == 0 {
		return false
	}

	//模式串不能直接是*打头
	if pattern[0] == '*' {
		return false
	}

	//判断是否模式串第二个字符是*
	if len(pattern) > 1 && pattern[1] == '*' {
		//模式串的第二个字符是 *，处理相对比较复杂
		if str[0] == pattern[0] ||                      // 如果第一个字符相等
			(pattern[0] == '.' && len(str) > 0 ) {      // 或者模式串第一个字符是.，这也等同于相等
			return recurseReg(str[1:], pattern[2:]) ||  // 字符串匹配一个，模式串跳过*
					recurseReg(str[1:], pattern) ||     // 字符串匹配一个，模式串维持原状
					recurseReg(str, pattern[2:])        // 字符串维持，模式串直接忽略*和器前面字符
		} else {
			return recurseReg(str, pattern[2:])         // 第一个字符不相等，则直接利用*标识0个字符的功效
		}
	} else {
		//模式串的第二个字符不是 *
		if len(str) > 0 && len(pattern) > 0 {           // 必须保证此时字符串和模式串都未完结
			if str[0] == pattern[0] ||                  // 如果第一个字符相等
				pattern[0] == '.' {                     // 或者模式串第一个字符是.，这也等同于相等
				return recurseReg(str[1:], pattern[1:]) // 同时向后平移
			}
		}
	}

	return false
}