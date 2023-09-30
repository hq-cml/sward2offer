/*
 * 给你一个整数数组 nums 和一个整数 k ，请你统计并返回该数组中和为 k 的连续子数组的个数 。
 * 示例 1：
 * 输入：nums = [1,1,1], k = 2
 * 输出：2
 *
 * 示例 2：
 * 输入：nums = [1,2,3], k = 3
 * 输出：2
 */
package _560_sub_array_sum

// 思路：这题直观的思路是利用滑动窗口，实际上由于元素是无序的，利用滑动窗口解决不了问题
// 利用前缀和数组的思想：
// 1, 2, 3  => 0, 1, 3, 6  (0是表示作为边界，这相当于就似乎一个累计值的数组）
// 从后向前遍历这个数组，只要数组中的两个值的差 == k，说明就存在一个满足条件的连续数组
// 其实这种做法复杂度仍然是N^2，本题还可以进一步优化
func subarraySum(nums []int, k int) int {
	if len(nums) == 0 {
		return 0
	}
	// 前缀和数组
	prefixSums := make([]int, len(nums)+1)
	for i := 1; i <= len(nums); i++ {
		prefixSums[i] = prefixSums[i-1] + nums[i-1]
	}

	// 从后向前
	var ret int
	for j := len(nums); j >= 0; j-- {
		for i := 0; i < j; i++ {
			if prefixSums[j]-prefixSums[i] == k {
				ret++
			}
		}
	}
	return ret
}
