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
