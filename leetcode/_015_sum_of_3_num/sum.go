/*
 * 给你一个整数数组 nums ，判断是否存在三元组 [nums[i], nums[j], nums[k]]
 * 满足i != j != k。 同时还满足 nums[i] + nums[j] + nums[k] == 0 。请
 * 你返回所有和为 0 且不重复的三元组。 注意：答案中不可以包含重复的三元组。
 */
package _015_sum_of_3_num

import (
	"fmt"
	"sort"
)

// 思路：先整体排序，然后双指针逼近
// 时间复杂度：排序O(n*logn) + 循环逼近：n^2 => O(n^2)
func Find3Num(arr []int, sum int) {
	if len(arr) < 3 {
		return
	}
	// 排序
	sort.Ints(arr)
	uniq := map[string]struct{}{}
	for idx, v := range arr {
		tmp := find2Num(arr, sum-v, idx) // 三个值，每次挑出一个值，然后指针逼近寻找另两个
		if len(tmp) != 0 {
			for _, t := range tmp {
				key := genUnionKey(v, t)
				if _, ok := uniq[key]; !ok {
					fmt.Println(v, t[0], t[1])
					uniq[key] = struct{}{}
				}
			}
		}
	}
}

// 在一个有序的数组中，找到两个数的和为sum，且要避开黑名单idx
// 采用双指针逼近的策略
func find2Num(arr []int, sum int, blackIdx int) [][2]int {
	var ret [][2]int
	i := 0
	j := len(arr) - 1
	for i < j {
		if i == blackIdx {
			i++
			continue
		}
		if j == blackIdx {
			j--
			continue
		}
		if arr[i]+arr[j] > sum {
			j--
		} else if arr[i]+arr[j] < sum {
			i++
		} else {
			ret = append(ret, [2]int{arr[i], arr[j]})
			i++ // 随意
		}
	}
	return ret
}

func genUnionKey(n int, a [2]int) string {
	if n <= a[0] {
		return fmt.Sprintf("%d_%d_%d", n, a[0], a[1])
	} else if n >= a[1] {
		return fmt.Sprintf("%d_%d_%d", a[0], a[1], n)
	} else {
		return fmt.Sprintf("%d_%d_%d", a[0], n, a[1])
	}
}
