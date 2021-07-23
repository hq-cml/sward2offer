/*
 * 面试题48：最长不含重复字符的子字符串
 * 题目：请从字符串中找出一个最长的不包含重复字符的子字符串，计算该最长子
 * 字符串的长度。假设字符串中只包含从'a'到'z'的字符。
 */
package _48_max_no_repititon_string

//arabcacfr

//思路1：
//暴力方法，穷举
//从每个字符开始，向后延长直到遇到重复，时间复杂度O(N^2)
//难度：1*
func FindMaxNoRepititionString(str string) (int, int) {
	var maxCnt = 0;
	var index = 0;
	for idx:=0; idx < len(str); idx ++ { //从每个字符分别开始探测
		cnt := 0
		m := map[byte]struct{}{}
		for i := range str[idx:] {
			_, ok := m[str[i]]
			if !ok {
				m[str[i]] = struct{}{}
				cnt ++
			} else {                     //一旦遇到重复，则探测结束
				if cnt > maxCnt {
					maxCnt = cnt
					index = idx
				}
				break
			}
		}
		if cnt > maxCnt {                //可能到了末尾，也没出现重复
			maxCnt = cnt
			index = idx
		}
	}

	return index, maxCnt
}

//思路2：
//滑动窗口，左右两个边界指针，不断的向右移动
//如果不重复，则右边界后移，如果重复，则左边界后移
//时间复杂度O(N)
//难度：4*
func FindMaxNoRepititionStrSlideWindow(str string) (int, int) {
	var maxCnt = 0;
	var index = 0;
	length := len(str)
	i, j := 0, 0                     //两指针，均从头开始：i右边界；j左边界
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
    //abcdeefghij
	return index, maxCnt
}

//思路3：
//作者的方法（动态规划）
//f(i)表示第i个字符结尾的，不包含重复字符的，子字符串最大长度，则：
//	如果第i个字符未出现过，则f(i) = f(i-1) + 1
//  如果第i个字符出现过，记录第i个字符和它上次出现的距离为d：
//     d<=f(i-1)，即第i个字符，出现在f(i-1)对应的长度之中，则f(i) = d
//     d >f(i-1)，即第i个字符，出现在很早的位置，则f(i) = f(i-1) + 1
//难度：5*
func FindMaxNoRepititionStrDynamicPlan(str string) (int, int) {
	var currLen int
	var maxLen int
	var index = 0;

	charPosition := make([]int, 26)
	for i:=0; i<26; i++ {
		charPosition[i] = -1           //空间换时间，charPosition记录了每个字母的最后出现位置
	}

	for idx :=0; idx < len(str); idx++ {
		byteI := str[idx] - 'a'

		if charPosition[byteI] == -1 {  //第idx字符从未出现过
			currLen ++
		} else {
			d := idx - charPosition[byteI]
			if d > currLen {		    //第idx字符曾经出现过，且第idx个字符，出现在很早的位置
				currLen ++
			} else {					//第idx字符曾经出现过，且第idx个字符，出现在f(idx-1)对应的长度之中
				if currLen >maxLen {
					maxLen = currLen
					index = idx - currLen
				}
				currLen = d
			}
		}
		charPosition[byteI] = idx       //记录第i字符最新出现位置
	}

	//不重复子串，出现在了末尾
	if currLen > maxLen {
		maxLen = currLen
		index = len(str) - currLen
	}
	return index, maxLen
}



