/*
	*
	 - 给你一个整数数组 nums ，你需要找出一个 连续子数组 ，如果对这个子数组进行升序排序，那么整个数组都会变为升序排序。

请你找出符合题意的 最短 子数组，并输出它的长度。

示例 1：

输入：nums = [2,6,4,8,10,9,15]
输出：5
解释：你只需要对 [6, 4, 8, 10, 9] 进行升序排序，那么整个表都会变为升序排序。
示例 2：

输入：nums = [1,2,3,4]
输出：0
示例 3：

输入：nums = [1]
输出：0
*/
package _581_shortest_subarr

import (
	"fmt"
	"sort"
)

// 思路：这题就是采用了朴素解法，copy和排序，前后向中间逼近
// TODO 可以进一步优化
func FindUnsortedSubarray(nums []int) int {
	if len(nums) <= 1 {
		return 0
	}

	cop := make([]int, len(nums))
	copy(cop, nums)
	sort.Ints(cop)
	fmt.Println(cop)
	i, j := 0, len(nums)-1
	for i < len(nums) {
		if nums[i] != cop[i] {
			break
		}
		i++
	}
	if i == len(nums) {
		return 0
	}
	for j >= 0 {
		if nums[j] != cop[j] {
			break
		}
		j--
	}
	fmt.Println(i)
	fmt.Println(j)
	return j - i + 1
}
