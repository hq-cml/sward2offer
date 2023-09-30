/*
 * 面试题48：最长不含重复字符的子字符串
 * 题目：请从字符串中找出一个最长的不包含重复字符的子字符串，计算该最长子
 * 字符串的长度。假设字符串中只包含从'a'到'z'的字符。
 *
 * PS：本题也是leetcode 3原题
 */
package _003_max_norep_str

// 思路1：
// 暴力方法，穷举
// 从每个字符开始，向后延长直到遇到重复，时间复杂度O(N^2)
// 难度：2*
func FindMaxNoRepititionString(str string) (int, int) {
	var maxCnt = 0
	var index = 0
	for idx := 0; idx < len(str); idx++ { //从每个字符分别开始探测
		cnt := 0
		m := map[byte]struct{}{}
		for i := range str[idx:] {
			_, ok := m[str[i]]
			if !ok {
				m[str[i]] = struct{}{}
				cnt++
			} else { //一旦遇到重复，则探测结束
				if cnt > maxCnt {
					maxCnt = cnt
					index = idx
				}
				break
			}
		}
		if cnt > maxCnt { //可能到了末尾，也没出现重复
			maxCnt = cnt
			index = idx
		}
	}

	return index, maxCnt
}

// 思路2：
// 滑动窗口，左右两个边界指针，不断的向右移动
// 如果不重复，则右边界后移，如果重复，则左边界后移
// 利用一个map来判断是否出现重复
// 时间复杂度O(N)
// 难度：4*
func FindMaxNoRepititionStrSlideWindow(str string) (int, int) {
	var maxLen = 0 //最终最大的子串长度
	var index = 0  //最终子串开始位置
	length := len(str)
	j, i := 0, 0 //两指针，均从头开始：i左边界；j右边界
	m := map[byte]struct{}{}
	for j < length && i < length {
		c := str[j]
		_, ok := m[c]
		if !ok {
			//窗口右边界右移
			j++
			m[c] = struct{}{}
			if maxLen < (j - i) {
				maxLen = (j - i)
				index = i
			}
		} else {
			//窗口左边界右移
			delete(m, str[i])
			i++
		}
	}
	//abcdeefghij
	return index, maxLen
}

// 滑动窗口的变种实现，利用一个map来保存所有的字符的位置
// 我自己的方法
func FindMaxNoRep(str string) (int, int) {
	if len(str) == 0 {
		return 0, 0
	}

	uniq := map[byte]int{}
	idx := 0     // 最终返回值，最长子串的idx
	max := -1    // 最终返回值，最长子串的长度
	currIdx := 0 // 当前子串的开始idx
	for i, c := range []byte(str) {
		if repIdx, ok := uniq[c]; !ok {
			uniq[c] = i
			length := i - currIdx + 1
			if length > max {
				max = length
				idx = currIdx // 将当前开始idx，作为返回结果
			}
		} else {
			// 将重复的子串部分，全部清除掉
			for j := currIdx; j <= repIdx; j++ {
				delete(uniq, str[j])
			}
			uniq[c] = i          // 重复部分已清除，新字符入uniq
			currIdx = repIdx + 1 // 当前子串的起始位置从清除后的位置开始
		}
	}

	return idx, max
}

// 思路3：
// 作者的方法（动态规划）
// f(i)表示第i个字符结尾的，不包含重复字符的，子字符串最大长度，则：
//
//		如果第i个字符未出现过，则f(i) = f(i-1) + 1
//	 如果第i个字符出现过，记录第i个字符和它上次出现的距离为d：
//	    d<=f(i-1)，即第i个字符，出现在f(i-1)对应的长度之中，则f(i) = d
//	    d >f(i-1)，即第i个字符，出现在很早的位置，则f(i) = f(i-1) + 1
//
// 其实，本质上这个方法和上面思路2是类似的，只是更加抽象化
// 难度：5*
func FindMaxNoRepititionStrDynamicPlan(str string) (int, int) {
	var currLen int
	var maxLen int
	var index = 0

	// 本质上是当做一个map在使用，用于判断字符是否出现过，如果出现过，则记录上一次出现位置
	charPosition := make([]int, 26)
	for i := 0; i < 26; i++ {
		charPosition[i] = -1 //空间换时间，charPosition记录了每个字母的最后出现位置
	}

	for idx := 0; idx < len(str); idx++ {
		byteI := str[idx] - 'a'

		if charPosition[byteI] == -1 { //第idx字符从未出现过
			currLen++
		} else {
			d := idx - charPosition[byteI]
			if d > currLen { //第idx字符曾经出现过，且第idx个字符，出现在很早的位置
				currLen++
			} else { //第idx字符曾经出现过，且第idx个字符，出现在f(idx-1)对应的长度之中
				if currLen > maxLen {
					maxLen = currLen
					index = idx - currLen
				}
				currLen = d
			}
		}
		charPosition[byteI] = idx //记录第i字符最新出现位置
	}

	//不重复子串，出现在了末尾
	if currLen > maxLen {
		maxLen = currLen
		index = len(str) - currLen
	}
	return index, maxLen
}
