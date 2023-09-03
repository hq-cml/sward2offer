/*
 * 给定一个非负整数数组 nums ，你最初位于数组的 第一个下标 。
 *
 * 数组中的每个元素代表你在该位置可以跳跃的最大长度。
 * 判断你是否能够到达最后一个下标。
 *
 * 输入：nums = [2,3,1,1,4]
 * 输出：true
 * 解释：可以先跳 1 步，从下标 0 到达下标 1, 然后再从下标 1 跳 3 步到达最后一个下标。
 *
 * 输入：nums = [3,2,1,0,4]
 * 输出：false
 * 解释：无论怎样，总会到达下标为 3 的位置。但该下标的最大跳跃长度是 0 ， 所以永远不可能到达最后一个下标。
 */
package _055_jump_game

// 思路：递归思想
// Tips：这题的另一个巧妙之处是failedIndexes，如果没有这个设置，一些case会超时
func Jump(arr []int) bool {
	if len(arr) == 0 {
		return false
	}
	failedIndexes := make([]bool, len(arr)) //失败黑名单，缓存结果进行加速
	target := len(arr) - 1
	return jump(arr, 0, target, failedIndexes)
}

// 无需多言
func jump(arr []int, begIdx int, targetIdx int, failedIndexes []bool) bool {
	if failedIndexes[begIdx] {
		return false
	}
	if begIdx == targetIdx {
		return true
	}

	curr := arr[begIdx]
	for i := 1; i <= curr; i++ {
		if jump(arr, begIdx+i, targetIdx, failedIndexes) {
			return true
		}
	}
	failedIndexes[begIdx] = true
	return false
}
