/*
 * 给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。
 * 请注意 ，必须在不复制数组的情况下原地对数组进行操作。
 *
 * 示例 1:
 * 输入: nums = [0,1,0,3,12]
 * 输出: [1,3,12,0,0]
 * 示例 2:
 *
 * 输入: nums = [0]
 * 输出: [0]
 */
package _283_move_zero

// 思路：快慢指针策略
func MoveZeroes(nums []int) {
	fast := 0
	slow := 0
	for fast < len(nums) {
		if nums[fast] == 0 { // 快指针找到第一个不是0的数字
			fast++
			continue
		}

		// 将不是0的赋值给slow，然后快慢同时后移
		nums[slow] = nums[fast]
		slow++
		fast++
	}

	// 将剩余的slow后面全部置为0
	for i := slow; i < len(nums); i++ {
		nums[i] = 0
	}
}
