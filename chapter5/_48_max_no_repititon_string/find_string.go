/*
 *

 */
package _48_max_no_repititon_string

//arabcacfr
//寻找最长不重复子串，暴力方法
//难度：1*
func FindMaxNoRepititionString(str string) (int, int) {
	var maxCnt = 0;
	var index = 0;
	for idx:=0; idx < len(str); idx ++ {
		cnt := 0
		m := map[byte]struct{}{}
		for i := range str[idx:] {
			_, ok := m[str[i]]
			if !ok {
				m[str[i]] = struct{}{}
				cnt ++
			} else {
				if cnt > maxCnt {
					maxCnt = cnt
					index = idx
				}
				break
			}
		}
		if cnt > maxCnt {
			maxCnt = cnt
			index = idx
		}
	}

	return index, maxCnt
}

//寻找最长不重复子串（滑动窗口）
//左右两个边界指针，不断的向右移动
//如果不重复，则右边界后移，如果重复，则左边界后移
//难度：4*
func FindMaxNoRepititionStrSlideWindow(str string) (int, int) {
	var maxCnt = 0;
	var index = 0;
	length := len(str)
	i, j := 0, 0
	m := map[byte]struct{}{}
	for i < length && j < length {
		c := str[i]
		_, ok := m[c]
		if !ok {
			//窗口右边界右移
			i++
			m[c] = struct{}{}
			if maxCnt < (i-j) {
				maxCnt = (i-j)
				index = j
			}
		} else {
			//窗口左边界右移
			delete(m, str[j])
			j++
		}
	}

	return index, maxCnt
}

//作者的方法（动态规划）
//f(i)表示第一个字符结尾的，不包含重复字符的，子字符串最大长度，则：
//	如果第i个字符未出现过=>f(i) = f(i-1) + 1
//  如果第i个字符出现过，记录第i个字符和它上次出现的距离为d：
//     d<=f(i-1)，即第i个字符，出现在f(i-1)对应的长度之中=>f(i) = d
//     d >f(i-1)，第i个字符，出现在很早的位置=>f(i) = f(i-1) + 1
//难度：5*
func FindMaxNoRepititionStrDynamicPlan(str string) (int, int) {
	var currLen int
	var maxLen int
	var index = 0;

	charPosition := make([]int, 26)
	for i:=0; i<26; i++ {
		charPosition[i] = -1    //空间换时间，charPosition记录了每个字母的最后出现位置
	}

	for i:=0; i<len(str); i++ {
		byteI := str[i] - 'a'
		d := i - charPosition[byteI]
		if charPosition[byteI] == -1 { //第i字符从未出现过
			currLen ++
		} else if d > currLen {		   //第i字符曾经出现过，且第i个字符，出现在很早的位置
			currLen ++
		} else {					   //第i字符曾经出现过，且第i个字符，出现在f(i-1)对应的长度之中
			if currLen >maxLen {
				maxLen = currLen
				index = i - currLen
			}
			currLen = d
		}
		charPosition[byteI] = i //记录第i字符最新出现位置
	}

	if currLen >maxLen {
		maxLen = currLen
		index = len(str) - 1 - currLen
	}
	return index, maxLen
}



