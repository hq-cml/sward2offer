/*
 * 岛屿数量
 * 入参是一个二维数组，分别是1和0，1标识陆地，0标识水
 * 所有连成片（上下左右）的陆地，算一块岛屿
 */
package _200_numbers_of_island

// 思路：仍然是dfs深度遍历的变种
//      遍历整个数组，一旦遇到一块陆地，则进行dfs深度，将所有的连接的陆地连城一块
func IslandNum(arr [][]byte) int {
	// 初步校验
	rows := len(arr)
	if rows < 1 {
		return 0
	}
	cols := len(arr[0])
	if cols < 1 {
		return 0
	}

	// 已访问标记
	var visited [][]bool
	for i:=0; i<rows; i++ {
		visited = append(visited, make([]bool, cols))
	}

	// 逐个探测
	cnt := 0
	for i:=0; i<rows; i++ {
		for j:=0; j<cols; j++ {
			// 奔着新陆地
			dsfIsland(arr, rows, cols, i, j, visited, &cnt, true)
		}
	}
	return cnt
}

func dsfIsland(arr [][]byte, rows, cols, i, j int, visited [][]bool, cnt *int, new bool) {
	// 边界
	if i < 0 || i >= rows {
		return
	}
	if j < 0 || j >= cols {
		return
	}

	// 已访问
	if visited[i][j] {
		return
	}

	// 访问标记
	visited[i][j] = true

	// 遇到水
	if arr[i][j] == 0 {
		return
	}

	// 遇到陆地
	if new {
		*cnt ++ // 是新陆地，岛屿自增
	}
	dsfIsland(arr, rows, cols, i+1, j, visited, cnt, false) // 非新陆地
	dsfIsland(arr, rows, cols, i, j+1, visited, cnt, false)
	dsfIsland(arr, rows, cols, i-1, j, visited, cnt, false)
	dsfIsland(arr, rows, cols, i, j-1, visited, cnt, false)
}
