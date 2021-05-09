package _13_active_range

func CalcRange(rows, cols int, max int) int {
	if rows < 1 || cols <1 || max < 0 {
		return 0
	}
	visited := make([]bool, rows * cols)

	return calcRange(rows, cols, 0, 0, max, visited)
}

func calcRange(rows, cols int, row, col int, max int, visited []bool) int {
	if calc(row) + calc(col) > max {
		return 0
	}
	if row < 0 || row >= rows || col < 0 || col >= cols {
		return 0
	}
	if visited[row*cols + col] {
		return 0
	}

	visited[row*cols + col] = true
	cnt := 1 + calcRange(rows, cols, row, col+1, max, visited) +
		calcRange(rows, cols, row, col-1, max, visited) +
		calcRange(rows, cols, row+1, col, max,  visited) +
		calcRange(rows, cols, row-1, col, max,  visited)

	return cnt
}

func calc(n int) int {
	sum := 0
	for n != 0 {
		sum += n % 10
		n = n / 10
	}
	return sum
}
