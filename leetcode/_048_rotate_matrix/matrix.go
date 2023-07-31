/*
  - 给定一个 n×n 的二维矩阵matrix 表示一个图像。请你将图像顺时针旋转 90 度。
  - 你必须在 原地 旋转图像，这意味着你需要直接修改输入的二维矩阵。请不要 使用另一个矩阵来旋转图像。

输入：matrix =   输出：

	[[1,2,3],      [[7,4,1],
	[4,5,6],       [8,5,2],
	[7,8,9]]       [9,6,3]]

输入：matrix =       输出：

	[[5,1,9,11],    [[15,13,2,5],
	 [2,4,8,10],    [14,3,4,1],
	 [13,3,6,7],    [12,6,8,9],
	[15,14,12,16]]  [16,7,10,11]]
*/
package _048_rotate_matrix

// 思路：从外到内，一层层的旋转
func Matrix(matrix [][]int) {
	n := len(matrix)
	circles := n / 2 // 层数
	if n%2 == 1 {
		circles++
	}
	for i := 0; i < circles; i++ {
		rotateCircle(matrix, n, i)
	}
}

// 单独处理一层的旋转
// 非常烧脑的是四个坐标的计算
func rotateCircle(matrix [][]int, dim, circle int) {
	beg := circle
	end := dim - circle - 1
	// 对第一行（除最后一个元素），每个元素循环移动4次，则本层结束
	for idx := beg; idx < end; idx++ {
		x0, y0 := beg, idx          // 行固定
		x1, y1 := x0+(y0-beg), end  // 列固定
		x2, y2 := end, end-(x1-beg) // 行固定
		x3, y3 := end-(end-y2), beg // 列固定
		// 循环移动
		matrix[x0][y0], matrix[x1][y1], matrix[x2][y2], matrix[x3][y3] =
			matrix[x3][y3], matrix[x0][y0], matrix[x1][y1], matrix[x2][y2]
	}
}
