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