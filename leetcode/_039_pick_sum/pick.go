/*
 * 给你一个 无重复元素 的整数数组candidates(数组全部元素>=0） 和一个目标整数target(>0)
 * 找出candidates中可以使数字和为目标数target 的 所有不同组合 ，并以列表形式返回。
 * candidates 中的 同一个 数字可以 无限制重复被选取 。如果至少一个数字的被选数量不同，则两种组合是不同的。
 * 例1：
 * 输入：candidates = [6,2,7,3], target = 7
 * 输出：[[2,2,3],[7]]
 * 例2：
 * 输入: candidates = [3,2,5], target = 8
 * 输出: [[2,2,2,2],[2,3,3],[3,5]]
*/
package _039_pick_sum

import "sort"

// 思路：这是一个DFS的经典应用，有点类似于全排列的求解
// 另外就是对于结果要求不重复，这里的处理很巧妙
// 首先将待选数组排序，进而保证结果的录入选取只能从当前值（因为允许多次使用）往后
// 这样结果成员是由小到大的，自然就不会出现重复的情况了
func FindAllCombin(arr []int, target int) [][]int {
	if len(arr) == 0 {
		return nil
	}
	var ret [][]int
	var curr []int
	sort.Ints(arr) // 先排序

	// dfs
	var find func(arr []int, begIdx int, crr []int, target int)
	find = func(arr []int, begIdx int, curr []int, target int) {
		// 每个成员都有尝试的必要
		for idx, v := range arr {
			// idx只能>=begIdx，进而保证了不出现重复
			if idx < begIdx {
				continue
			}

			// 如果找到need，则得到一组结果，否则，继续dfs
			if v == target {
				ret = append(ret, append(curr, v))
			} else if v < target {
				// target > v才有进一步找的必要（因为数组是递增的正整数）
				find(arr, idx, append(curr, v), target-v)
			} else {
				continue
			}
		}
	}
	find(arr, 0, curr, target)
	return ret
}
