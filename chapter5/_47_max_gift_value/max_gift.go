/*
 * 面试题47：礼物的最大价值
 * 题目：在一个m×n的棋盘的每一格都放有一个礼物，每个礼物都有一定的价值
 * （价值大于0）。你可以从棋盘的左上角开始拿格子里的礼物，并每次向右或
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
//坐标为(i,j)的最大礼物值为f(i,j)，因为坐标点(i,j)必然是由(i-1,j)或者(i,j-1)走过来的，所以
//f(i,j) = gift[i,j] + max[f(i-1,j), f(i,j-1)]
//难度：4*
func giftRecurse(values []int, rows, cols int, i, j int) int {
	//先计算好最底层值
	if i < 0 || j < 0 { //越界
		return 0
	}

	//递归结束条件：左上角礼物价值
	if i == 0 && j == 0 {
		return values[0]
	}

	//自顶向下，递归下去
	lastGift := values[i*cols+j] //右下角的礼物价值
	return lastGift + max(
		giftRecurse(values, rows, cols, i-1, j),
		giftRecurse(values, rows, cols, i, j-1),
	)
}

//思路2；循环方案（贪心）
//递归的逆向思维，用一个数组来缓存局部的最大值
//最妙的是第0行和第0列的礼物值，由于路径是唯一的，可以直接求出作为循环边界
//难度：5*
func giftLoop(values []int, rows, cols int, i, j int) int {
	maxValues := make([]int, rows*cols)
	//直接从最底层，自底向上的计算，逐层计算并缓存
	maxValues[0] = values[0]
	//填充第0行，作为循环边界（这个路径是必然的）
	for idx := 1; idx <= j; idx++ {
		maxValues[idx] = maxValues[idx-1] + values[idx]
	}
	//填充第0列，作为循环边界（这个路径是必然的）
	for idx := 1; idx <= i; idx++ {
		maxValues[idx*cols] = maxValues[(idx-1)*cols] + values[idx*cols]
	}

	//循环计算剩下的非0行和0列
	for row := 1; row <= i; row++ { //从第1行开始
		for col := 1; col <= j; col++ { //从第1列开始
			currIdx := row*cols+col
			currGift := values[currIdx]          //当前格子的礼物价值
			maxValues[currIdx] = currGift + max( //当前格子最大价值，取决于当前价值和他前面的最大价值
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
