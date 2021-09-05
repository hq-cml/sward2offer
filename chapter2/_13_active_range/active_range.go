/*
 * 面试题13：机器人的运动范围
 * 题目：地上有一个m行n列的方格。一个机器人从坐标(0, 0)的格子开始移动，它
 * 每一次可以向左、右、上、下移动一格，但不能进入行坐标和列坐标的数位之和
 * 大于k的格子。例如，当k为18时，机器人能够进入方格(35, 37)，因为3+5+3+7=18。
 * 但它不能进入方格(35, 38)，因为3+5+3+8=19。请问该机器人能够到达多少个格子？
 */
package _13_active_range

//思路:
//上一题类似，这种路径探测类的场景，用栈或者递归来实现
//难度：5*
func CalcRange(rows, cols int, max int) int {
	if rows < 1 || cols < 1 || max < 0 {
		return 0
	}
	visited := make([]bool, rows*cols)

	return calcRange(rows, cols, 0, 0, max, visited)
}

//递归
func calcRange(rows, cols int, row, col int, max int, visited []bool) int {
	//退出条件：坐标数位之和大于给定值
	if calc(row)+calc(col) > max {
		return 0
	}

	//越界条件
	if row < 0 || row >= rows || col < 0 || col >= cols {
		return 0
	}

	//如果当前坐标已经探测过，则退出
	if visited[row*cols+col] {
		return 0
	}

	//阶段性成功，当前坐标记录为已访问
	visited[row*cols+col] = true

	//在当前的基础上，分别向上下左右四个方向继续探测，且将探测的结果相加
	cnt := 1 + calcRange(rows, cols, row, col+1, max, visited) +
		calcRange(rows, cols, row, col-1, max, visited) +
		calcRange(rows, cols, row+1, col, max, visited) +
		calcRange(rows, cols, row-1, col, max, visited)

	//最终结果
	return cnt
}

//求数位之和
func calc(n int) int {
	sum := 0
	for n != 0 {
		sum += n % 10
		n = n / 10
	}
	return sum
}
