/*
*
  - 给你一个含 n 个整数的数组 nums ，其中 nums[i] 在区间 [1, n] 内。请你找出所有在 [1, n] 范围内但没有出现在 nums 中的数字，并以数组的形式返回结果。

示例 1：

输入：nums = [4,3,2,7,8,2,3,1]
输出：[5,6]
示例 2：

输入：nums = [1,1]
输出：[2]
*/
package _448_find_num

func FindDisappearedNumbers(nums []int) []int {
	n := len(nums)
	arr := make([]int, n+1)
	for _, v := range nums {
		arr[v] = 1
	}
	var ret []int
	for k, v := range arr {
		if k == 0 {
			continue
		}
		if v == 0 {
			ret = append(ret, k)
		}
	}
	return ret
}
