/*
 * 给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 target  的那 两个 整数，并返回它们的数组下标。
 * 你可以假设每种输入只会对应一个答案。但是，数组中同一个元素在答案里不能重复出现。
 *
 * 例：
 * 输入：nums = [2,7,11,15], target = 9 输出：[0,1]
 * 输入：nums = [3,2,4], target = 6 输出：[1,2]
 * 输入：nums = [3,3], target = 6 输出：[0,1]
 */
package _01_two_sum

//这个题容易和滑动指针混淆，其实题目未提及数组一定是排序的，且要求返回下标，所以用滑动指针并不合适
//思路1：暴力法
//思路2：空间换时间，利用map，来加速寻找的速度，O(N)
func Sum(nums []int, sum int) []int {
	if len(nums) < 2 {
		return nil
	}
	m := map[int]int{}
	for idx, v := range nums {
		need := sum - v
		if i, ok := m[need]; ok {
			return []int{i, idx}
		}
		m[v] = idx
	}
	return nil
}
