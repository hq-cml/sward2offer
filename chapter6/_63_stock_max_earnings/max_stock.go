/*
// 面试题63：股票的最大利润
// 题目：假设把某股票的价格按照时间先后顺序存储在数组中，请问买卖交易该股
// 票可能获得的利润是多少？例如一只股票在某些时间节点的价格为{9, 11, 8, 5,
// 7, 12, 16, 14}。如果我们能在价格为5的时候买入并在价格为16时卖出，则能
// 收获最大的利润11。

这题很容易被懵住，遍历一遍数组，求出最大值和最小值，然后max-min得到结果。
这是错的，因为存在顺序！！必须要先买入，才可能有卖出。
仍然采用遍历一遍的思路，但是，不断保存当前最小的值和当前最大的diff值。
 */
package _63_stock_max_earnings

//股票最大收益
//代码不复杂，但是不容易想到，O(1)的复杂度实现
//遍历价格数组，保留到当前价格之前的最小价格以及最大的价格收益差
//难度：4*
func MaxStock(price []int) int {
	if len(price) < 2 {
		return 0
	}

	min := price[0]
	maxDiff := price[1] - min
	for i:=2; i < len(price); i++ {
		if price[i-1] < min {
			min = price[i-1]
		}

		currentDiff := price[i] - min
		if currentDiff > maxDiff {
			maxDiff = currentDiff
		}
	}

	return maxDiff
}
