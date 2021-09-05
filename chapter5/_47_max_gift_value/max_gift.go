/*
 * 面试题47：礼物的最大价值
 * 题目：在一个m×n的棋盘的每一格都放有一个礼物，每个礼物都有一定的价值
 * （价值大于0）。你可以从棋盘的左上角开始拿格子里的礼物，并每次向左或
 * 者向下移动一格直到到达棋盘的右下角。给定一个棋盘及其上面的礼物，请计
 * 算你最多能拿到多少价值的礼物？
 */
package _47_max_gift_value

import (
	"math"
)

//思路：
//这种最优解的题目，通常都是动态规划，并且递归的分析问题，但是循环的去解决。
//这个和上一题很相似，只不过是1维变2维了
func MaxGiftValue(values []int, rows, cols int, recurse bool) int {
	//校验
	if len(values) != rows*cols {
		return 0
	}

	//是否递归的进行
	if recurse {
		return giftRecurse(values, rows, cols, rows-1, cols-1) // rows-1, cols-1 => 右下角
	} else {
		return giftLoop(values, rows, cols, rows-1, cols-1)
	}
}

//思路1：递归方案
//f(i,j) = gift[i,j] + max[f(i-1,j), f(i,j-1)]
//难度：4*
func giftRecurse(values []int, rows, cols int, i, j int) int {
	//先计算好最底层值
	if i < 0 || j < 0 { //越界
		return 0
	}
	//if i == 0 && j == 0 { //左上角礼物价值
	//	return values[0]
	//}

	//自顶向下，递归下去
	lastGift := values[i*cols+j] //右下角的礼物价值
	return lastGift + max(
		giftRecurse(values, rows, cols, i-1, j),
		giftRecurse(values, rows, cols, i, j-1),
	)
}

//思路2；循环方案
//递归的逆向思维，用一个数组来缓存局部的最大值
//难度：5*
func giftLoop(values []int, rows, cols int, i, j int) int {
	maxValues := make([]int, rows*cols)
	//直接从最底层，自底向上的计算，逐层计算并缓存
	maxValues[0] = values[0]
	//填充第0行
	for idx := 1; idx <= j; idx++ {
		maxValues[idx] = maxValues[idx-1] + values[idx]
	}
	//填充第0列
	for idx := 1; idx <= i; idx++ {
		maxValues[idx*cols] = maxValues[(idx-1)*cols] + values[idx*cols]
	}

	//循环计算剩下的非0行和0列
	for row := 1; row <= i; row++ { //从第二行开始
		for col := 1; col <= j; col++ { //从第二列开始
			currGift := values[row*cols+col]          //当前格子的礼物价值
			maxValues[row*cols+col] = currGift + max( //当前格子最大价值，取决于当前价值和他前面的最大价值
				maxValues[(row-1)*cols+col],
				maxValues[row*cols+(col-1)],
			)
		}
	}

	return maxValues[i*cols+j]
}

func max(a, b int) int {
	return int(math.Max(float64(a), float64(b)))
}
