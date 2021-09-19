/*
 * fw面试题：
 * 给定一个字符串，寻找不含交集的字符子串。
 * 比如：
    输入arrab，返回[arra，b]；
    输入arraba，则仅返回：[arraba]
    输入abc，则返回：[abc]
    输入arrarbc，则仅返回：[arrar, bc]
 */
package _find_substr_no_intersection

//这个题
// 最朴素的思路是穷举，显然不合理
// 主要是空间换时间的思路，利用一个map，存储所有字母最后出现的位置
// 这样就可以快速的判断是否存在交集
//难度： 4*
func FindSubStr(str string) []string {
	length := len(str)

	if length == 0 {
		return nil
	}

	//找到每个字符，出现的最后位置
	flags := map[byte]int{}
	for idx := length - 1; idx >= 0; idx-- {
		c := str[idx]
		_, ok := flags[c]
		if !ok {
			flags[c] = idx
		}
	}

	ret := []string{}
	org := -1
	tmp := []byte{}
	for idx := 0; idx < length; idx++ {
		org = idx
		end, ok := flags[str[idx]]
		if ok && idx != end {
			if len(tmp) > 0 {
				ret = append(ret, string(tmp)) //零散字母拼接
			}
			for idx != end {          		//idx追赶end边界
				idx++
				tmp, ok := flags[str[idx]]
				if ok && tmp > end {        //end边界后移
					end = tmp
				}
			}
			ret = append(ret, str[org:idx+1])
		} else {
			tmp = append(tmp, str[idx])     //零散字幕暂存
		}
	}

	if len(tmp) > 0 {                    //剩余零散字母拼接
		ret = append(ret, string(tmp))
	}
	return ret
}

