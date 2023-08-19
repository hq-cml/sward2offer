/*
 * 只出现一次的数字
 * 给你一个非空整数数组 nums ，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。
 *
 * 你必须设计并实现线性时间复杂度的算法来解决此问题，且该算法只使用常量额外空间。
 * 输入：nums = [2,2,1] => 2
 * 输入：nums = [4,1,2,1,2] => 4
 * 输入：nums = [1] => 1
 */
package _136_only_one_num

// 编程之美，直接利用异或
func SingleNumber(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	num := nums[0]
	for i := 1; i < len(nums); i++ {
		num = num ^ nums[i]
	}
	return num
}
