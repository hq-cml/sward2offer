/*
 * 实现获取下一个排列的函数，算法需要将给定数字序列重新排列成字典序中下一个更大的排列。
 * 如果不存在下一个更大的排列，则将数字重新排列成最小的排列（即升序排列）。
 * 必须原地修改，只允许使用额外常数空间。
 *
 * 以下是一些例子，输入位于左侧列，其相应输出位于右侧列。
 * 1,2,3 → 1,3,2
 * 3,2,1 → 1,2,3
 * 1,3,2 → 2,1,3
 * 4,2,5,1->4 1 2 5
 * 4,2,5,3,1->4 3 1 2 5
*/
package _031_next_permutation

import "sort"

// 思路，举例子
// 1. 直接从后往前扫描,如果一路都是递增，说明没有下一个更大值，则整个序列逆序
// 2. 如果不是一路递增的，则必然能找到第一个下坠的数，将其记为 a. (例子：4,2,5,3,1，找到第一个下坠的数是2）
// 3. 然后用这位数去和它之后的数进行比较,找到比它大的数中最小的一位b. (例子：4,2,5,3,1，找到2之后比2大最小的一个数字，3）
// 4. 然后交换a和b的位置(例子：4,2,5,3,1，交换2和3，得到4,3,5,2,1）
// 5. 再将a之后的数字升序排列,就是最终的结果. (例子：4,2,5,3,1，将剩余的5,2,1升序排序，则得到最终的4,3,1,2,5）
func NextPermutation(arr []int) {
	if len(arr) <= 1 {
		return
	}
	// 从后向前，找到第一对顺序数（下坠）
	first := -1
	for i := len(arr) - 1; i > 0; i-- {
		if arr[i] > arr[i-1] {
			first = i - 1
			break
		}
	}

	// 如果不存在顺序数，完全逆序，则意味着需要整体reverse
	if first == -1 {
		// reverse
		sort.Ints(arr)
		return
	}

	// 找到first之后最小的一个数（但是这个数不能比first小），swap
	spIdx := first + 1
	min := arr[spIdx]
	for i := spIdx + 1; i < len(arr); i++ {
		if arr[i] > arr[first] && min > arr[i] {
			spIdx = i
			min = arr[spIdx]
		}
	}
	arr[first], arr[spIdx] = arr[spIdx], arr[first]

	// 将first后面的值，升序排列
	sort.Ints(arr[first+1:])
}
