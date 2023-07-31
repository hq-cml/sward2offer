/*
 * 面试题66：构建乘积数组
 * 题目：给定一个数组A[0, 1, …, n-1]，请构建一个数组B[0, 1, …, n-1]，其
 * 中B中的元素B[i] =A[0]×A[1]×… ×A[i-1]×A[i+1]×…×A[n-1]。不能使用除法。
 */
package _66_multiply_array

//思路：
//思路1：如果可以用除法，是比较简单的，乘起来之后每行分别除以对应数，复杂度O(n)。但是题目要求不能用除法
//思路2：常规思路就是两次循环，则复杂度O(n^2)
//思路3：
// 例：
// 给定 nums = [1,2,3,4]，返回 [24,12,8,6]。
// 算法：
// 1. 使用两个临时数组 left 和 right
// 2. 将 left[i] 的值设为 nums[0] 到 nums[i-1] 的乘积
// 3. 将 right[i] 的值设为 nums[i+1] 到 nums[n-1] 的乘积
// 4. 将 res[i] 的值设为 left[i]*right[i]
// 5. 返回 res
//难度：4*
func Multiply(src []int) []int {
	//特殊情况
	if len(src) == 0 {
		return src
	}
	n := len(src)

	//最终结果容器
	left := make([]int, n)
	left[0] = 1
	right := make([]int, n)
	right[n-1] = 1

	for i:=1; i<n; i++ {
		left[i] = left[i-1] * src[i-1]
	}
	for i:=n-2; i>=0; i-- {
		right[i] = right[i+1] * src[i+1]
	}

	// 最终结果
	dst := make([]int, n)
	for i:=0; i<n; i++ {
		dst[i] = left[i] * right[i]
	}
	return dst
}
