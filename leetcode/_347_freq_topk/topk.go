/*
*

	*给你一个整数数组 nums 和一个整数 k ，请你返回其中出现频率前 k 高的元素。你可以按 任意顺序 返回答案。

示例 1:

输入: nums = [1,1,1,2,2,3], k = 2
输出: [1,2]
示例 2:

输入: nums = [1], k = 1
输出: [1]
*/
package _347_freq_topk

// 思路1：这题很明显可以使用堆排序，不过性能还是不够高
// 思路2：桶排序思想!
func TopKFrequent(nums []int, k int) []int {
	// 先统计词频
	freq := make(map[int]int)
	for _, v := range nums {
		freq[v]++
	}

	// 将词频翻转成一个list，list下标就是频率，list的值是满足这个词频的所有的元素（可能是多个）
	// 巧妙的点是，这个list的长度应该怎么定？这里直接用了len(nums)+1，因为最极端的情况就是所有的元素都一样
	list := make([][]int, len(nums)+1)
	for key, cnt := range freq {
		list[cnt] = append(list[cnt], key)
	}

	// 遍历得到得到最终结果
	var ret []int
	for i := len(nums); i >= 0; i-- {
		if len(list[i]) == 0 { // 如果某个元素从未出现过，则应该忽略
			continue
		}
		subList := list[i]
		for _, key := range subList {
			if k > 0 {
				ret = append(ret, key)
				k--
			}

		}
	}
	return ret
}
