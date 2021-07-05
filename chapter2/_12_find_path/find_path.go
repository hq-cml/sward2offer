package _12_find_path

/*

// 面试题12：矩阵中的路径
// 题目：请设计一个函数，用来判断在一个矩阵中是否存在一条包含某字符串所有
// 字符的路径。路径可以从矩阵中任意一格开始，每一步可以在矩阵中向左、右、
// 上、下移动一格。如果一条路径经过了矩阵的某一格，那么该路径不能再次进入
// 该格子。例如在下面的3×4的矩阵中包含一条字符串“bfce”的路径（路径中的字
// 母用下划线标出）。但矩阵中不包含字符串“abfb”的路径，因为字符串的第一个
// 字符b占据了矩阵中的第一行第二个格子之后，路径不能再次进入这个格子。
// A B T G
// C F C S
// J D E H

回溯法通常，都是利用栈或者递归来实现！
 a b t g
 c f c s
 j d e h

bfce
abfb
*/

func FindPath(arr []byte, rows, cols int, find string) bool {
	if rows < 1 || cols <1 || len(arr) != rows*cols  {
		return false
	}
	visited := make([]bool, rows * cols)
	idx := 0
	//可从任意一个格子开始，任意一般性
	for i := 0; i < rows; i ++ {
		for j := 0; j < cols; j ++ {
			if findPath(arr, rows, cols, i, j, find, idx, visited) {
				return true
			}
		}
	}

	return false
}

func findPath(arr []byte, rows, cols int, row, col int, find string, idx int, visited []bool) bool {
	if idx + 1 > len(find) {
		return true
	}

	if row < 0 || row >= rows || col < 0 || col >= cols {
		return false
	}

	if arr[row*cols + col] != find[idx] || visited[row*cols + col] {
		return false
	}

	visited[row*cols + col] = true
	ok := findPath(arr, rows, cols, row, col+1, find, idx + 1, visited) ||
		findPath(arr, rows, cols, row, col-1, find, idx + 1, visited) ||
		findPath(arr, rows, cols, row+1, col, find, idx + 1, visited) ||
		findPath(arr, rows, cols, row-1, col, find, idx + 1, visited)
	if ok {
		return true
	} else {
		visited[row*cols + col] = false
		return false
	}
}