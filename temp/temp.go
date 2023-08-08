package temp

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
	uniq[i][j] = true

	replace(arr, src, dst, i-1, j, rows, cols, uniq)
	replace(arr, src, dst, i, j-1, rows, cols, uniq)
	replace(arr, src, dst, i+1, j, rows, cols, uniq)
	replace(arr, src, dst, i, j+1, rows, cols, uniq)
}
