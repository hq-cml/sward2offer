/*
 * 一个机器人位于一个 m x n网格的左上角 （起始点在下图中标记为 “Start” ）。
 * 机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为 “Finish” ）。
 * 问总共有多少条不同的路径？
 * 示例 1：
 * 输入：m = 3, n = 2 => 输出：3
 * 解释：
 * 从左上角开始，总共有 3 条路径可以到达右下角。
 * 1. 向右 -> 向下 -> 向下
 * 2. 向下 -> 向下 -> 向右
 * 3. 向下 -> 向右 -> 向下
 *
 * 示例 2：
 * 输入：m = 7, n = 3 => 输出：28
 *
 * 示例 4：
 * 输入：m = 3, n = 3 => 输出：6
 */
package _62_different_path

// 动态规划
func CalcPath(m, n int) int {
	if m <= 0 && n <= 0 {
		return 0
	}
	return dp(m-1, n-1)
}

// 无需多言
func dp(i, j int) int {
	if i < 0 || j < 0 {
		return 0
	}
	if i == 0 && j == 0 {
		return 0
	}
	if i == 0 && j == 1 {
		return 1
	}
	if i == 1 && j == 0 {
		return 1
	}
	return dp(i-1, j) + dp(i, j-1)
}
