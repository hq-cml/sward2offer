/*
*
  - 给你一个 只包含正整数 的 非空 数组 nums 。请你判断是否可以将这个数组分割成两个子集，使得两个子集的元素和相等。

示例 1：

输入：nums = [1,5,11,5]
输出：true
解释：数组可以分割成 [1, 5, 5] 和 [11] 。
示例 2：

输入：nums = [1,2,3,5]
输出：false
解释：数组不能分割成两个元素和相等的子集。

提示：

1 <= nums.length <= 200
1 <= nums[i] <= 100
*/
package _416_split_subset

// 这题可以转换思路，假设集合的整体和是sum，那么问题转化成为能否找到有一个子集，和是sum/2
// 思路：动态规划
//
//		dp[i][j]-- 表示[0,i]这个区间所有整数，是否能够凑出一部分，使得和 == j
//		则：
//		1. nums[i]未选中：dp[i][j] = dp[i-1][j]
//		2. nums[i]被选中，分两种情况：
//	         2.1 nums[i] == j， dp[i][j] = true
//	         2.2 nums[i] < j，dp[i][j] = dp[i-1][j-nums[i]]
//
// 这个真是好难
func canPartition(nums []int) bool {
	// 求和
	var sum int
	for _, v := range nums {
		sum += v
	}
	if sum%2 == 1 { // 总和不能为奇数
		return false
	}
	target := sum / 2

	// 问题转化为数组中凑一个子数组，使得和为target
	dp := make([][]bool, len(nums))
	for i := 0; i < len(nums); i++ {
		dp[i] = make([]bool, target+1)
	}

	// 初始值，第0行
	if nums[0] <= target {
		dp[0][nums[0]] = true
	}

	// 遍历填值
	for i := 1; i < len(nums); i++ {
		for j := 0; j <= target; j++ {
			// 当前不选用
			dp[i][j] = dp[i-1][j]

			// 当前选用
			if nums[i] == j {
				dp[i][j] = true
				continue
			}
			if nums[i] < j {
				tmp := dp[i-1][j] || dp[i-1][j-nums[i]]
				dp[i][j] = tmp || dp[i][j]
			}
		}
	}
	return dp[len(nums)-1][target]
}
