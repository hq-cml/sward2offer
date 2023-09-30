/**
 * 最小值，方正证券题目
 */
package _064_max_gift

import "math"

func MinGift(values [][]int) int {
	rows := len(values)
	if rows == 0 {
		return 0
	}
	cols := len(values[0])
	if cols == 0 {
		return 0
	}

	return minGift(values, rows-1, cols-1)
}

func minGift(values [][]int, i, j int) int {
	if i < 0 || j < 0 {
		// 这里巨坑，很容易就思考错误了。不是返回0，而应该是返回以一个最大值！！因为整个题目求得是最小值，0可能会被错误的参与计算
		return math.MaxInt
	}
	if i == 0 && j == 0 {
		return values[0][0]
	}
	min := Min(minGift(values, i-1, j), minGift(values, i, j-1))
	return values[i][j] + min
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
