/*
 * 你是一个专业的小偷，计划偷窃沿街的房屋。每间房内都藏有一定的现金，影响你偷窃的唯一制约因素就是相邻的房屋装有相互连通的防盗系统，如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警。
 * 给定一个代表每个房屋存放金额的非负整数数组，计算你 不触动警报装置的情况下 ，一夜之内能够偷窃到的最高金额。
 *
 * 示例 1：
 * 输入：[1,2,3,1]
 * 输出：4
 * 解释：偷窃 1 号房屋 (金额 = 1) ，然后偷窃 3 号房屋 (金额 = 3)。
 * 偷窃到的最高金额 = 1 + 3 = 4 。
 *
 * 示例 2：
 * 输入：[2,7,9,3,1]
 * 输出：12
 * 解释：偷窃 1 号房屋 (金额 = 2), 偷窃 3 号房屋 (金额 = 9)，接着偷窃 5 号房屋 (金额 = 1)。
 * 偷窃到的最高金额 = 2 + 9 + 1 = 12 。
 */
package _198_rob

// 非递归
// 思路：动态规划，假设dp[i]，表示到第i家为止，能够偷到的最大值（i=0,1..., n-1）
//
//	则dp[i]可能有两种情况：
//	1. 第i家偷了，则dp[i] = nums[i] + dp[i-2]  //因为不能相邻的偷，所以i-2
//	2. 第i家不偷，则dp[i] = dp[i-1]
//	3. 所以，dp[i]应该是1和2的取Max值
func Rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	dp[1] = Max(nums[0], nums[1])
	for i := 2; i < len(nums); i++ {
		dp[i] = Max((nums[i] + dp[i-2]), dp[i-1])
	}
	return dp[len(nums)-1]
}

// 递归写法
func RobRecurse(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 {
		return nums[0]
	}

	if len(nums) == 2 {
		return Max(nums[0], nums[1])
	}

	return Max((nums[len(nums)-1] + RobRecurse(nums[:len(nums)-2])),
		RobRecurse(nums[:len(nums)-1]))
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
