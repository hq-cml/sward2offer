/**
 * 给你一个整数数组 coins ，表示不同面额的硬币；以及一个整数 amount ，表示总金额。
 * 计算并返回可以凑成总金额所需的 最少的硬币个数 。如果没有任何一种硬币组合能组成总金额，返回 -1 。
 * 你可以认为每种硬币的数量是无限的。
 *
 * 示例 1：
 * 输入：coins = [1, 2, 5], amount = 11
 * 输出：3
 * 解释：11 = 5 + 5 + 1
 *
 * 示例 2：
 * 输入：coins = [2], amount = 3
 * 输出：-1
 *
 * 示例 3：
 * 输入：coins = [1], amount = 0
 * 输出：0
 */
package _322_coin_changes

import "math"

// 正确思路：动态规划
//
//	dp[i]表示，要凑出金额i，需要的最少钞票数，则假设钞票面额a,b,c 3种
//	dp[i] = 1 + Min(dp[i-a], dp[i-b], dp[i-c])
func CoinChange(coins []int, amount int) int {
	if len(coins) == 0 {
		return 0
	}
	dp := make([]int, amount+1)
	dp[0] = 0
	for i := 1; i <= amount; i++ {
		min := math.MaxInt
		for _, coin := range coins {
			if coin < i {
				if dp[i-coin] == -1 {
					continue
				}
				if 1+dp[i-coin] < min {
					min = 1 + dp[i-coin]
				}
			} else if coin == i {
				min = 1
				break
			} else {
				continue
			}
		}
		if min != math.MaxInt {
			dp[i] = min
		} else {
			dp[i] = -1
		}
	}
	return dp[amount]
}

// 错误思路：谈心算法
// 总是希望找到最大的面值并使用，这种想法直观上对，实际不对
// 例子：{10, 5, 7, 2}
//      谈心算法得到：10+5+2
//      实际正解：7+7
//func CoinChange(coins []int, amount int) int {
//	if len(coins) == 0 {
//		return 0
//	}
//	if amount == 0 {
//		return 0
//	}
//
//	// 逆序
//	sort.Ints(coins)
//	i, j := 0, len(coins)-1
//	for i < j {
//		coins[i], coins[j] = coins[j], coins[i]
//		i++
//		j--
//	}
//
//	//
//	var cnt int
//	if change(coins, amount, &cnt) {
//		return cnt
//	} else {
//		return -1
//	}
//}
//
//func change(coins []int, need int, cnt *int) bool {
//	for _, v := range coins {
//		if v > need {
//			continue
//		}
//		if v == need {
//			*cnt++
//			return true
//		}
//		// v < need，尝试dfs
//		*cnt++
//		if change(coins, need-v, cnt) {
//			return true
//		}
//		*cnt--
//	}
//	return false
//}
