/*
 * 给定一个整数数组prices，其中第  prices[i] 表示第 i 天的股票价格 。
 * 设计一个算法计算出最大利润。在满足以下约束条件下，你可以尽可能地完成更多的交易（多次买卖一支股票）:
 * 卖出股票后，你无法在第二天买入股票 (即冷冻期为 1 天)。
 * 注意：你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。
 *
 * 示例 1:
 * 输入: prices = [1,2,3,0,2]
 * 输出: 3
 * 解释: 对应的交易状态为: [买入, 卖出, 冷冻期, 买入, 卖出]
 * 示例 2:
 * 输入: prices = [1]
 * 输出: 0
 */
package _309_stock

// 思路，动态规划，难点在于动态转移方程
//
//	本题因为存在冷冻期的问题，所以状态转移方程是一个更加复杂的设计
//	dp[i] 表示第i天的最大收益，它是两种子情况的Max
//	dp[i][0] -- 第i天未持有股票
//	dp[i][1] -- 第i天持有股票
//	分析可知，无论是否持有股票，又分为两种子情况，综上：
//	dp[i][0] = Max( dp[i-1][1]+price[i],      ## 当天卖出的（前一天持有，这里是+price，因为卖出是挣钱，别搞混了）
//	                dp[i-1][0] )              ## 前一天就没持有
//	dp[i][1] = Max( dp[i-2][0] - price[i],    ## 当天买入（前2天是不持仓的，因为有冷冻期）
//	                dp[i-1][1])               ## 前一天就持有
//	初始：即第0天
//	dp[0][0] = 0
//	dp[0][1] = -price[0]
func MaxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	dp := make([][2]int, len(prices))
	dp[0][0] = 0
	dp[0][1] = 0 - prices[0]

	for i := 1; i < len(prices); i++ {
		dp[i][0] = Max(dp[i-1][1]+prices[i], dp[i-1][0])
		dp[i][1] = Max(func(i int) int {
			if i >= 2 {
				return (dp[i-2][0] - prices[i])
			} else {
				return (0 - prices[i])
			}
		}(i), dp[i-1][1])
	}

	var max int
	for _, v := range dp {
		if v[0] > max {
			max = v[0]
		}
		if v[1] > max {
			max = v[1]
		}
	}
	return max
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
