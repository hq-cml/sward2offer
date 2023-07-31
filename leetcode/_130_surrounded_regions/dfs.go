/*
 * 类似于围棋
 * 入参是一个二维数组，分别是X和O，将不在边上的O，全部替换成X
 *
 * 例子1：
 * XXXX      XXXX
 * XOOX  =>  XXXX
 * XXOX      XXXX
 * XOXX      XOXX
 *
 * 例子2：
 * XXXXXX     XXXXXX
 * XOOXOX =>  XXXXOX
 * XXOXOX     XXXXOX
 * XOXXOX     XOXXOX
 */
package _130_surrounded_regions

// 思路：
// 1. 分析可得，和四个边中O以及相邻的O，将不被替换，其他的O会被替换
// 2. 从4边开始，将4边中出现的O，进行深度遍历，全部标记为中间态字符$
// 3. 扫描整个二维数组，将剩下的O，设置成为X；将中间字符$，设置成为O
func Solve(arr [][]byte) {
	rows := len(arr) // m行
	if rows <= 2 {   // 至少需要3行，才可能需要替换
		return
	}
	cols := len(arr[0]) // n列
	if cols <= 3 {
		return // 小于1列，必然不可替换
	}

	var visited [][]bool
	for i:=0; i<rows; i++ {
		visited = append(visited, make([]bool, cols))
	}

	// 从四条边开始，逐个探测
	for j:=0; j<cols; j++ {
		dfsSolve(arr, rows, cols, 0, j, visited)
	}
	for i:=1; i<rows; i++ {
		dfsSolve(arr, rows, cols, i, cols-1, visited)
	}
	for j:=cols-2; j>=0; j-- {
		dfsSolve(arr, rows, cols, rows-1, j, visited)
	}
	for i:=rows-1; i>=1; i-- {
		dfsSolve(arr, rows, cols, i, 0, visited)
	}

	// 最终替换
	for i:=0; i<rows; i++ {
		for j:=0; j<cols; j++ {
			if arr[i][j] == 'T' {
				arr[i][j] = 'O'
			} else {
				arr[i][j] = 'X'
			}
		}
	}
}

// 深度遍历：从坐标x, y开始检测替换
func dfsSolve(arr [][]byte, rows, cols int, x, y int, visited [][]bool) {
	// 越界
	if x < 0 || x > rows-1 || y < 0 || y > cols-1 {
		return
	}

	if visited[x][y] {
		return
	}
	visited[x][y] = true

	// 无需替换
	if arr[x][y] != 'O' {
		return
	}

	arr[x][y] = 'T' // 临时标记，最终无需替换的O

	// 上下左右dfs
	dfsSolve(arr, rows, cols, x-1, y, visited)
	dfsSolve(arr, rows, cols, x+1, y, visited)
	dfsSolve(arr, rows, cols, x, y-1, visited)
	dfsSolve(arr, rows, cols, x, y+1, visited)
}

