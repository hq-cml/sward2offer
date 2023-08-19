/*
*

	*给你一个整数数组 nums ，请你找出数组中乘积最大的非空连续子数组（该子数组中至少包含一个数字），并返回该子数组所对应的乘积。

测试用例的答案是一个 32-位 整数。

子数组 是数组的连续子序列。

示例 1:

输入: nums = [2,3,-2,4]
输出: 6
解释: 子数组 [2,3] 有最大乘积 6。
示例 2:

输入: nums = [-2,0,-1]
输出: 0
解释: 结果不能为 2, 因为 [-2,-1] 不是子数组。
*/
package _152_mul_max

// 思路：和leetcode 53类似
//
//	区别在于乘积可能时而正数时而负数，所以这里除了累计最大值，也要累计最小值，说不定一乘，就正了
func MaxProduct(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	preMax := nums[0] //最大值
	preMin := nums[0] //最小值
	ret := nums[0]
	for i := 1; i < len(nums); i++ {
		a := preMax * nums[i] //累计值
		b := preMin * nums[i]

		preMax = Max(nums[i], Max(a, b)) // 从累计值和当前值进行比较，选择大的，即如果出现比当前值还小，就从当前值开始
		preMin = Min(nums[i], Min(a, b))

		ret = Max(ret, preMax)
	}

	return ret
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
