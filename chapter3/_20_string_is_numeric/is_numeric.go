/*
// 面试题20：表示数值的字符串
// 题目：请实现一个函数用来判断字符串是否表示数值（包括整数和小数）。例如，
// 字符串“+100”、“5e2”、“-123”、“3.1416”及“-1E-16”都表示数值，但“12e”、
// “1a3.14”、“1.2.3”、“+-5”及“12e+5.4”都不是


这种题目写得出来的，都不太容易。。。
情况特别的多，很不容易。
主要思路是将数字分成3大块，整数部分，小数部分，指数部分，其中难点在于他们都不一定是必须存在的，所以要能够逐个判断并移动指针。
 */
package _20_string_is_numeric

//A.Be|EC
//Ae|EC
//.Be|EC
//A.B
//A
//.B;
func IsNumeric(str string) bool {
	if len(str) == 0 {
		return false
	}

	//先扫描整数部分
	b, ok := scanSignedInt([]byte(str))
	if ok {
		return true
	}
	if b[0] != '.' && b[0] != 'e' && b[0] != 'E'{
		return false
	} else if b[0] == '.' {
		//存在小数点，扫描小数部分
		b = b[1:]
		b, ok = scanUnSignedInt(b)
		if !ok {
			return false
		}
	}

	//扫描指数部分
	if len(b) > 0 &&
		(b[0] == 'e' || b[0] == 'E') {
		b = b[1:]
		b, ok = scanSignedInt(b)
		if !ok {
			return false
		}
	}

	if len(b) > 0 {
		return false
	}

	return true
}

//扫描无符号型字符串整形
func scanUnSignedInt(str []byte) ([]byte, bool) {
	length := len(str)
	var i int
	for i=0; i<length; i++ {
		if !(str[i] >= '0' && str[i] <= '9') {
			return str[i:], false
		}
	}

	return str[i:], true
}

//扫描有符号型字符串整形
func scanSignedInt(str []byte) ([]byte, bool) {
	if len(str) == 0 {
		return str, false
	}
	if (str[0] >= '0' && str[0] <= '9') ||
		str[0] == '+' ||
		str[0] == '-' {
		return scanUnSignedInt(str[1:])
	}
	return nil, false
}