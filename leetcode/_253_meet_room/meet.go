/**
 * 预定会议室
 * 给一个输入，就是一堆会议的开始和节数时间点，要求使用最少的会议室（因为只要两个会议不重叠，就可以使用同一个会议室）
 */
package _253_meet_room

import (
	"sort"
)

// 这题思路非常巧妙，可以画线段图来分析叠加情况
// 最终抽象出来发现只要在数轴上计算出每个时间点的会议开始点和释放点
func Meet(intervals [][]int) int {
	m := map[int]int{} // timepoint => diff
	for _, v := range intervals {
		m[v[0]] += 1 // 会议开始
		m[v[1]] -= 1 // 会议释放
	}

	// 对所有的时间点排序，每个时间点上都累计了开始点和释放点
	var keys []int
	for k, _ := range m {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	// 将每个时间点的开始和释放数累计，就可到了当前点具体应该有几个会议室够用，求出最大的那个，就是最少需要的会议室数量
	sum := 0
	max := 0
	for _, k := range keys {
		sum += m[k]
		if sum > max {
			max = sum
		}
	}

	//fmt.Println(m)
	return max
}
