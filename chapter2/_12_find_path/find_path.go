package _12_find_path

/*
 a b t g
 c f c s
 j d e h

bfce
abfb
*/

func FindPath(arr []byte, rows, cols int, find string) bool {
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