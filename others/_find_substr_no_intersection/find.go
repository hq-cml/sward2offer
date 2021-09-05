/*
 * fw面试题：
 * 给定一个字符串，寻找不含交集的字符子串。
 * 比如：输入arrab，返回[arra，b]；输入arraba，则仅返回：[arraba]
 */
package _find_substr_no_intersection

func FindSubStr1(str string) []string {
	length := len(str)

	if length == 0 {
		return nil
	}

	flags := map[byte]int{}
	for i := length - 1; i >= 0; i-- {
		c := str[i]
		_, ok := flags[c]
		if !ok {
			flags[c] = i
		}
	}

	ret := []string{}
	org := 0
	for i := 0; i < length; i++ {
		org = i
		j, ok := flags[str[i]]
		if ok && i != j {
			for i != j {
				i++

				tmp, ok := flags[str[i]]
				if ok && tmp > j {
					j = tmp
				}
			}
			ret = append(ret, str[org:i+1])

		} else {
			ret = append(ret, string(str[i]))
		}
	}

	return ret
}

func FindSubStr2(str string) []string {
	length := len(str)

	if length == 0 {
		return nil
	}

	flags := map[byte]int{}
	for i := length - 1; i >= 0; i-- {
		c := str[i]
		_, ok := flags[c]
		if !ok {
			flags[c] = i
		}
	}

	ret := []string{}
	for i := 0; i < length; i++ {
		c := str[i]
		end, ok := flags[c]
		if ok {
			ret = append(ret, str[i:end+1])
			i = end
		} else {
			ret = append(ret, string(str[i]))
		}
	}

	return ret
}
