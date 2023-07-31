/*
 * 以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。
 * 请你合并所有重叠的区间，并返回一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间。
 *
 * 输入：intervals = [[1,3],[2,6],[8,10],[15,18]]
 * 输出：[[1,6],[8,10],[15,18]]
 * 解释：区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].
 *
 * 输入：intervals = [[1,4],[4,5]]
 * 输出：[[1,5]]
 * 解释：区间 [1,4] 和 [4,5] 可被视为重叠区间。
 */
package _056_merge_range

import (
	"sort"
)

func Merge(arr [][2]int) [][2]int {
	if len(arr) < 2 {
		return arr
	}
	// 先排序
	sort.Slice(arr, func(i, j int) bool {
		return arr[i][0] < arr[j][0]
	})

	var ret [][2]int
	ret = append(ret, arr[0])
	currIdx := 0
	for i := 1; i < len(arr); i++ {
		if ret[currIdx][1] >= arr[i][0] {
			// 合并
			if arr[i][1] > ret[currIdx][1] {
				ret[currIdx][1] = arr[i][1]
			}
		} else {
			// 不合并
			ret = append(ret, arr[i])
			currIdx++
		}
	}
	return ret
}
