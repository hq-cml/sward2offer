/*
 * 面试题4：二维数组中的查找
 * 题目：在一个二维数组中，每一行都按照从左到右递增的顺序排序，每一列都按
 * 照从上到下递增的顺序排序。请完成一个函数，输入这样的一个二维数组和一个
 * 整数，判断数组中是否含有该整数。
 *
 *   1  2   8  9
 *   2  4   9 12
 *   4  7  10 13
 *   6  8  11 15
 */
package _240_find_in_darr

// 最朴素的思路就会按行扫描，则时间复杂度为O(M*N)
// 充分利用题目对于二维数组预设条件
// 从右上角开始探测，假设右上角数字为i，需要寻找的为need
// 只要i>need，则该列肯定都>need，所以左移
// 只要i<need，则改行肯定都<need，所以下移
// 每一轮都是不断地先向左，再向下，这样最终如果能够找到，那么就是在数组的右上角，否则找不到。
func FindDoubleArray(arr []int, row, col int, needFind int) bool {
	if len(arr) != row*col {
		return false
	}

	i := 0 //i行j列，从右上角开始
	j := col - 1
	for i < row || j >= 0 {
		curr := arr[i*row+j]
		if curr == needFind {
			return true
		} else if curr > needFind {
			j--
		} else {
			i++
		}
	}

	return false
}
