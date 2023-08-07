/*
 * 面试题63：股票的最大利润
 * 题目：假设把某股票的价格按照时间先后顺序存储在数组中，请问买卖交易该股
 * 票可能获得的利润是多少？例如一只股票在某些时间节点的价格为{9, 11, 8, 5, 7, 12, 16, 14}。
 * 如果我们能在价格为5的时候买入并在价格为16时卖出，则能
 * 收获最大的利润11。
 */
package _63_stock_max_earnings

// 思路：
// 一个错误：遍历一遍数组，求出最大值和最小值，然后max-min得到结果。
// 这是错的，因为存在顺序问题！！必须要先买入，才可能有卖出。
// 思路1：暴力法，但是这样要O(N^2)
// 思路2：仍然采用遍历一遍的思路，但是，不断保存当前最小的值和当前最大的diff值。
//
//	代码不复杂，但是不容易想到，O(N)的复杂度实现
//
// 难度：4*
func MaxStock(price []int) int {
	// 至少2个交易日才能够出现价差
	if len(price) < 2 {
		return 0
	}

	currMin := price[0]
	maxDiff := 0
	for i := 1; i < len(price); i++ {
		// 当日价格减去历史最小值，得最大收益
		if (price[i] - currMin) > maxDiff {
			maxDiff = price[i] - currMin
		}
		// 记录历史最低值
		if price[i] < currMin {
			currMin = price[i]
		}
	}
	return maxDiff
}
