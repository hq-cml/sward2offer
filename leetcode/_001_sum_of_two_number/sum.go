/*
 * 给定一个整数数组 nums和一个整数目标值 target，该数组中找出和为目标值 target 的那两个整数
 * 并返回它们的数组下标。
 * 你可以假设每种输入只会对应一个答案。但是，数组中同一个元素在答案里不能重复出现。
 *
 * 示例 1：
 * 输入：nums = [2,7,11,15], target = 9
 * 输出：[0,1]
 * 解释：因为 nums[0] + nums[1] == 9 ，返回 [0, 1] 。
 * 示例 2：
 *
 * 输入：nums = [3,2,4], target = 6
 * 输出：[1,2]
 * 示例 3：
 *
 * 输入：nums = [3,3], target = 6
 * 输出：[0,1]
*/
package _001_sum_of_two_number

// 思路1：最直观的是穷举，复杂度n^2，显然不符合要求
// 思路2：想到了类似c57的递增数组的和的问题，但是这需要先排序，则复杂度是n*logn
// 思路3：空间换时间，引入map，key是已经出现的数组成员，val则是成员的下标
//
//	通过这样的处理，转化成为了寻找给定值的问题，整体复杂度降到了N
func Sum(arr []int, sum int) (int, int) {
	if len(arr) < 2 {
		return -1, -1
	}

	m := map[int]int{}
	for i := 0; i < len(arr); i++ {
		need := sum - arr[i] // 转化为了O(1)复杂度寻找need的问题
		if _, ok := m[need]; ok {
			// 如果找到则直接返回
			return m[need], i
		} else {
			// 如果没找到，则将当前成员极其下标存入map
			m[arr[i]] = i
		}
	}
	return -1, -1
}
