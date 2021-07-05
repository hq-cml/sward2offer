/*
// 面试题13：机器人的运动范围
// 题目：地上有一个m行n列的方格。一个机器人从坐标(0, 0)的格子开始移动，它
// 每一次可以向左、右、上、下移动一格，但不能进入行坐标和列坐标的数位之和
// 大于k的格子。例如，当k为18时，机器人能够进入方格(35, 37)，因为3+5+3+7=18。
// 但它不能进入方格(35, 38)，因为3+5+3+8=19。请问该机器人能够到达多少个格子？




 */

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
