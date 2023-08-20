/*
*

	*在一个由 '0' 和 '1' 组成的二维矩阵内，找到只包含 '1' 的最大正方形，并返回其面积。

输入：matrix = [["1","0","1","0","0"],["1","0","1","1","1"],["1","1","1","1","1"],["1","0","0","1","0"]]
输出：4

输入：matrix = [["0","1"],["1","0"]]
输出：1
示例 3：

输入：matrix = [["0"]]
输出：0
*/
package _221_max_area

// 思路：这题明显就应该是动态规划，难点在于动态转移方程规律的找寻
//
//	dp[i][j]表示以此节点作为右下角的正方形的最大边长，则
//那么如何计算 dp 中的每个元素值呢？对于每个位置 (i,j)(i, j)(i,j)，检查在矩阵中该位置的值：
//
//如果该位置的值是 0，则 dp(i,j)=0，因为当前位置不可能在由 1 组成的正方形中；
//
//如果该位置的值是 1，则 dp(i,j)的值由其上方、左方和左上方的三个相邻位置的 dp 值决定。具体而言，当前位置的元素值等于三个相邻位置的元素中的最小值加 111，状态转移方程如下：
//
//dp(i,j)=Min(dp(i−1,j),dp(i−1,j−1),dp(i,j−1))+1

func MaximalSquare(matrix [][]byte) int {
	dp := make([][]int, len(matrix))
	maxSide := 0

	// dp初始值
	for i := 0; i < len(matrix); i++ {
		dp[i] = make([]int, len(matrix[i]))
		for j := 0; j < len(matrix[i]); j++ {
			dp[i][j] = int(matrix[i][j] - '0')
			if dp[i][j] == 1 {
				maxSide = 1
			}
		}
	}

	// dp迭代计算
	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[i]); j++ {
			if matrix[i][j] == '1' {
				dp[i][j] = min(min(dp[i-1][j], dp[i][j-1]), dp[i-1][j-1]) + 1
				if dp[i][j] > maxSide {
					maxSide = dp[i][j]
				}
			}
		}
	}

	// 返回最长边对应的面积
	return maxSide * maxSide
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
