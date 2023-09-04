/*
 * 给定一个未排序的整数数组 nums ，找出数字连续的最长序列（不要求序列元素在原数组中连续）的长度。
 * 请你设计并实现时间复杂度为O(n) 的算法解决此问题。
 *
 * 示例 1：
 * 输入：nums = [100,4,200,1,3,2]=>4
 * 解释：最长数字连续序列是 [1, 2, 3, 4]。它的长度为 4。
 *
 * 示例 2：
 * 输入：nums = [0,3,7,2,5,8,4,6,0,1]=> 输出：9
 */
package _128_max_list

// 思路：整个数轴上，分成了一个个片段，要找到最长的片段
//  1. 空间换时间，用一个map来判断元素的存在性
//  2. 如何判断某个数字是片段头？如果v-1不存在于map中，就是片段头
func LongestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	// 空间换时间
	m := map[int]struct{}{}
	for _, v := range nums {
		m[v] = struct{}{}
	}

	max := 0
	for v, _ := range m {
		_, ok := m[v-1]
		if ok {
			continue
		}
		// 起点，开始计算这个片段的长度
		cnt := 1
		for {
			v++
			if _, ok1 := m[v]; ok1 {
				cnt++
			} else {
				break
			}
		}
		if cnt > max {
			max = cnt
		}
	}
	return max
}
