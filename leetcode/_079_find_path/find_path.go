/*
 * 面试题12：矩阵中的路径
 * 题目：请设计一个函数，用来判断在一个矩阵中是否存在一条包含某字符串所有
 * 字符的路径。路径可以从矩阵中任意一格开始，每一步可以在矩阵中向左、右、
 * 上、下移动一格。如果一条路径经过了矩阵的某一格，那么该路径不能再次进入
 * 该格子。例如在下面的3×4的矩阵中包含一条字符串“bfce”的路径（路径中的字
 * 母用下划线标出）。但矩阵中不包含字符串“abfb”的路径，因为字符串的第一个
 * 字符b占据了矩阵中的第一行第二个格子之后，路径不能再次进入这个格子。
 * A B T G
 * C F C S
 * J D E H
 */
package _079_find_path

// 回溯法
// 通常，都是利用栈或者递归来实现！
// 难度：5*
func Exist(board [][]byte, word string) bool {
	rows := len(board)
	if rows == 0 {
		return false
	}
	cols := len(board[0])
	if cols == 0 {
		return false
	}

	visited := make([][]bool, rows)
	for i := 0; i < rows; i++ {
		visited[i] = make([]bool, cols)
	}

	//可从任意一个格子开始，任意一般性
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if dfsExist(board, word, rows, cols, i, j, visited) {
				return true
			}
		}
	}
	return false
}

// 递归
func dfsExist(board [][]byte, word string, rows, cols int, i, j int, visited [][]bool) bool {
	//探测完毕，完全成功，符合要求
	if len(word) == 0 {
		return true
	}

	//异常边界
	if i < 0 || j < 0 || i >= rows || j >= cols {
		return false
	}

	//探测的目标已经访问过，则探测失败，返回false
	if visited[i][j] {
		return false
	}

	// 对比当前字符
	if word[0] != board[i][j] {
		return false
	}

	//当前这个字符，阶段性成功，标记已经访问的路径
	visited[i][j] = true

	//沿着当前道路，继续向上、下、左、右去探测，只要有一条路ok了，就说明当前位置是ok的
	if dfsExist(board, word[1:], rows, cols, i+1, j, visited) ||
		dfsExist(board, word[1:], rows, cols, i, j+1, visited) ||
		dfsExist(board, word[1:], rows, cols, i-1, j, visited) ||
		dfsExist(board, word[1:], rows, cols, i, j-1, visited) {
		//任意一条成功，则返回成功
		return true
	} else {
		//上下左右均失败，则说明当前点不是一个合适的路径，探测失败，并且需要回退
		visited[i][j] = false
		return false
	}
}
