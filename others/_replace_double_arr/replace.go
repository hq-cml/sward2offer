/*
  - hs面试题
  - 给定一个二维数组, 一个坐标(x, y), src, dst, 要求实现一个函数:
  - 如果数字==src的替换为dst, 并且如果发生了替换, 则周围(上下左右）也会发生该替换行为.
  - 例子
  - 给定一个数组input如下，参数: src = 2 , dst = 3, xy = (2,2):
    0 2 0 0 0
    0 2 2 2 0
    0 1 2 2 0
    0 2 0 0 0
  - 输出以下数组output:
    0 3 0 0 0
    0 3 3 3 0
    0 1 3 3 0
    0 2 0 0 0
*/
package replace_double_arr

// 二维数组递归循环替换
// 这个题不是特别难，关键是各种边界条件的判断
// 不能越界，并且，同时要考虑检测过得元素不能二次检测
func Replace(arr [][]int, src, dst int, i, j int) {
	rows := len(arr)
	if rows == 0 {
		return
	}
	cols := len(arr[0])
	if cols == 0 {
		return
	}

	if i < 0 || i > rows-1 {
		return
	}
	if j < 0 || j > cols-1 {
		return
	}

	if arr[i][j] != src {
		return
	}

	uniq := make([][]bool, rows)
	for t := 0; t < rows; t++ {
		uniq[t] = make([]bool, cols)
	}

	replace(arr, src, dst, i, j, rows, cols, uniq)
}

// 递归
func replace(arr [][]int, src, dst int, i, j int, rows, cols int, uniq [][]bool) {
	if i < 0 || i > rows-1 {
		return
	}
	if j < 0 || j > cols-1 {
		return
	}
	if arr[i][j] != src {
		return
	}
	if uniq[i][j] {
		return
	}

	arr[i][j] = dst
	uniq[i][j] = true // 防止无限递归

	replace(arr, src, dst, i-1, j, rows, cols, uniq)
	replace(arr, src, dst, i, j-1, rows, cols, uniq)
	replace(arr, src, dst, i+1, j, rows, cols, uniq)
	replace(arr, src, dst, i, j+1, rows, cols, uniq)
}
