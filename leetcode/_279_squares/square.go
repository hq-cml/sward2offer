/*
*

	*给你一个整数 n ，返回 和为 n 的完全平方数的最少数量 。

完全平方数 是一个整数，其值等于另一个整数的平方；换句话说，其值等于一个整数自乘的积。例如，1、4、9 和 16 都是完全平方数，而 3 和 11 不是。

示例 1：

输入：n = 12
输出：3
解释：12 = 4 + 4 + 4
示例 2：

输入：n = 13
输出：2
解释：13 = 4 + 9

提示：

1 <= n <= 104
*/
package _279_squares

import "math"

// 思路：这题其实说白了，就可以翻译成：1,4,9,16...备选
//
//	给一个数n作为和，用尽量少的上面的数，能凑出来，问最少多少个数能凑出来
//	注意：不存在凑不出来的情况，因为有1存在
//	显然，这是一个动态规划的思路，动态转移方程
//	dp[i]表示，和为i的数，需要最少的平方数
//	dp[i] = 1 + Min(dp[1], dp[2] ... dp[i - j*j])，其中 i> j*j
func NumSquares(n int) int {
	dp := make([]int, n+1)
	// i=1...n，逐个得出dp[i]值
	for i := 1; i <= n; i++ {
		min := math.MaxInt32
		// 逐个尝试j，j*j <i，所以dp[i] = 1 + Min(dp[1], dp[2] ... dp[i - j*j])
		for j := 1; j*j <= i; j++ {
			min = Min(min, dp[i-j*j])
		}
		dp[i] = min + 1
	}
	return dp[n]
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
