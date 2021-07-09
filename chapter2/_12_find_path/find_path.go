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
package _12_find_path

//回溯法
//通常，都是利用栈或者递归来实现！
func FindPath(arr []byte, rows, cols int, find string) bool {
	if rows < 1 || cols <1 || len(arr) != rows*cols  {
		return false
	}
	visited := make([]bool, rows * cols)
	idx := 0
	//可从任意一个格子开始，任意一般性
	for i := 0; i < rows; i ++ {
		for j := 0; j < cols; j ++ {
			if findPath(arr, rows, cols, i, j, find, idx, visited) { //发现任意一条路，即可退出
				return true
			}
		}
	}

	return false
}

//递归
//idx表示本轮要检测的是find的第idx位字符(idx从0开始)
//row，col表示本次探测，是针对arr数组的哪一个元素
func findPath(arr []byte, rows, cols int, row, col int, find string, idx int, visited []bool) bool {
	//探测完毕，完全成功，符合要求
	if idx + 1 > len(find) {
		return true
	}

	//异常扁鹊
	if row < 0 || row >= rows || col < 0 || col >= cols {
		return false
	}

	//如果探测不成功或者探测的目标已经访问过，则探测失败，返回false
	if arr[row*cols + col] != find[idx] || visited[row*cols + col] {
		return false
	}

	//阶段性成功，标记已经访问的路径
	visited[row*cols + col] = true

	//沿着当前道路，继续向上、下、左、右去探测，只要有一条路ok了，就说明当前位置是ok的
	ok := findPath(arr, rows, cols, row, col+1, find, idx + 1, visited) ||
		findPath(arr, rows, cols, row, col-1, find, idx + 1, visited) ||
		findPath(arr, rows, cols, row+1, col, find, idx + 1, visited) ||
		findPath(arr, rows, cols, row-1, col, find, idx + 1, visited)

	if ok {
		//任意一条成功，则返回成功
		return true
	} else {
		//上下左右均失败，则说明当前点不是一个合适的路径，探测失败，并且需要回退
		visited[row*cols + col] = false
		return false
	}
}