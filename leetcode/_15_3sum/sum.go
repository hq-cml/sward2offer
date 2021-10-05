/*
 * 给一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？请你找出所有和为 0 且不重复的三元组。
 * 注意：答案中不可以包含重复的三元组。
 *
 * 示例 1：
 * 输入：nums = [-1,0,1,2,-1,-4] 输出：[[-1,-1,2],[-1,0,1]]
 *
 * 示例 2：
 * 输入：nums = [] 输出：[]
 *
 * 示例 3：
 * 输入：nums = [0]输出：[]
 */
package _15_3sum

import (
	"sort"
)

//难度：4*
//思路：排序 + 双指针策略
//本题的难点在于如何去除重复解。
//
//对数组进行排序。
//遍历排序后数组：
//若 nums[i]>0：因为已经排序好，所以后面不可能有三个数加和等于 0，直接返回结果。
//对于重复元素：跳过，避免出现重复解
//令左指针 L=i+1，右指针 R=n-1，当 L<R 时，执行循环：
//当 nums[i]+nums[L]+nums[R]==0，执行循环，判断左界和右界是否和下一位置重复，去除重复解。并同时将 L,R 移到下一位置，寻找新的解
//若和大于 0，说明 nums[R] 太大，R 左移
//若和小于 0，说明 nums[L] 太小，L 右移
//复杂度分析
//时间复杂度：O(n^2)，数组排序 O(NlogN)，遍历数组 O(n)，双指针遍历 O(n)，总体 :O(NlogN)+O(n)∗O(n)，
//空间复杂度：O(1)

func ThreeSum(nums []int) [][]int {
	//结果
	ret := make([][]int, 0)

	n := len(nums)
	if n < 3 {
		return ret
	}

	//排序
	sort.Ints(nums)

	// 枚举每一个元素，作为第一个数字a
	// 然后用两个指针，模拟数字b和c，进行探测
	for i := 0; i < n; i++ {
		a := nums[i]
		if a > 0 {
			//若 nums[i]>0：因为已经排序好，所以后面不可能有三个数加和等于 0，直接返回
			break
		}
		if i > 0 && a == nums[i-1] {
			continue //去除重复的情况
		}
		// 不断调整左右两个指针，
		left := i + 1
		right := n - 1
		for left < right {
			b := nums[left]
			c := nums[right]
			if a+b+c == 0 {
				ret = append(ret, []int{a, b, c})
				// 打破平衡， 并且要考虑去重的问题
				for left < right && b == nums[left+1] {
					left++ // 去重
				}
				left++
				for left < right && c == nums[right-1] {
					right-- // 去重
				}
				right--

			} else if a+b+c < 0 {
				left++
			} else {
				right--
			}
		}
	}
	return ret
}
