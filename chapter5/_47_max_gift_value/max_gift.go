/*
// 面试题47：礼物的最大价值
// 题目：在一个m×n的棋盘的每一格都放有一个礼物，每个礼物都有一定的价值
// （价值大于0）。你可以从棋盘的左上角开始拿格子里的礼物，并每次向左或
// 者向下移动一格直到到达棋盘的右下角。给定一个棋盘及其上面的礼物，请计
// 算你最多能拿到多少价值的礼物？


这种题一般都是动态规划，并且递归的分析问题，但是循环的去解决。
 */
package _47_max_gift_value

import (
	"math"
)

//这个和上一题很相似，只不过是1维变2维了
//动态规划
func MaxGiftValue(values []int, rows, cols int, recurse bool) int {
	if len(values) != rows * cols {
		return 0
	}

	if recurse {
		return giftRecurse(values, rows, cols, rows-1, cols-1)
	} else {
		return giftLoop(values, rows, cols, rows-1, cols-1)
	}
}

//递归方案
//f(i,j) = max{f(i-1,j), f(i,j-1)} + gift[i,j]
func giftRecurse(values []int, rows, cols int, i, j int) int {
	//先计算好最底层值
	if i < 0 || j < 0 {
		return 0
	}
	if i == 0 && j == 0 {
		return values[0]
	}

	//自顶向下，递归下去
	return max(giftRecurse(values, rows, cols, i-1, j),
		giftRecurse(values, rows, cols, i, j-1)) + values[i*cols + j]
}

//循环方案
//递归的逆向思维，用一个数组来缓存局部的最大值
func giftLoop(values []int, rows, cols int, i, j int) int {
	maxValues := make([]int, rows * cols)
	//直接从最底层，自底向上的计算，逐层计算并缓存
	maxValues[0] = values[0]
	//填充第0行
	for idx := 1; idx <= j; idx ++ {
		maxValues[idx] = maxValues[idx-1] + values[idx]
	}
	//填充第0列
	for idx := 1; idx <= i; idx ++ {
		maxValues[idx*cols] = maxValues[(idx-1)*cols] + values[idx*cols]
	}

	//循环计算剩下的非0行和0列
	for row := 1; row <= i; row ++ {
		for col := 1; col <= j; col ++ {
			maxValues[row*cols + col] = max(
				maxValues[(row-1)*cols + col], maxValues[row*cols + (col-1)]) +
				values[row*cols + col]
		}
	}

	return maxValues[i*cols + j]
}

func max(a, b int) int {
	return int(math.Max(float64(a), float64(b)))
}

