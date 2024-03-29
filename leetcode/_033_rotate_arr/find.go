/*
 * 整数数组 nums 按升序排列，数组中的值 互不相同 。
 * 在传递给函数之前，nums 在预先未知的某个下标 k（0 <= k < nums.length）上进行了 旋转，
 * 使数组变为 [nums[k], nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]]（下标 从 0 开始 计数）。
 * 例如， [0,1,2,4,5,6,7] 在下标 3 处经旋转后可能变为[4,5,6,7,0,1,2] 。
 * 给你 旋转后 的数组 nums 和一个整数 target ，如果 nums 中存在这个目标值 target ，则返回它的下标，否则返回-1。
 * 你必须设计一个时间复杂度为 O(log n) 的算法解决此问题。
 *
 * eg1: nums = [4,5,6,7,0,1,2], target = 0 => 4
 * eg2: nums = [4,5,6,7,0,1,2], target = 3 => -1
 * eg3:
 */
package _033_rotate_arr

// 二分查找的变种
func Find(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	beg := 0
	end := len(nums) - 1
	for beg <= end {
		mid := (end-beg)/2 + beg
		if nums[mid] == target {
			return mid
		} else {
			// 如果nums[mid] != target，则
			// 首先，通过比较nums[mid]和nums[beg]来判断旋转发生在前半截还是后半截
			// 其次，利用了递增特性，再来确定目标存在于前半截还是后半截
			if nums[mid] > nums[beg] {
				// 前半截是顺序递增的（旋转发生在后半截）
				// 利用递增特性，判断target是否可能出现在这个区间内
				if target >= nums[beg] && target <= nums[mid-1] {
					end = mid - 1
				} else {
					beg = mid + 1
				}
			} else {
				// 后半截是顺序递增的（旋转发生在前半截）
				// 利用递增特性，判断target是否可能出现在这个区间内
				if mid+1 <= len(nums)-1 && target >= nums[mid+1] && target <= nums[end] {
					beg = mid + 1
				} else {
					end = mid - 1
				}
			}
		}
	}
	return -1
}
