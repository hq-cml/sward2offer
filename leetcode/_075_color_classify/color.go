/*
 * 给定一个包含红色、白色和蓝色、共n 个元素的数组nums，原地对它们进行排序，使得相同颜色的元素相邻，
 * 并按照红色、白色、蓝色顺序排列。我们使用整数 0、1 和 2 分别表示红色、白色和蓝色。
 * 必须在不使用库内置的 sort 函数的情况下解决这个问题。
 * 示例 1：
 * 输入：nums = [2,0,2,1,1,0]
 * 输出：[0,0,1,1,2,2]
 *
 * 示例 2：
 * 输入：nums = [2,0,1]
 * 输出：[0,1,2]
 */
package _075_color_classify

// 思路1：类似于基数排序，直接统计每个数字的个数，然后再直接更改，只需要扫描2遍
func Classify(nums []int) []int {
	cnt0 := 0
	cnt1 := 0
	cnt2 := 0
	for _, v := range nums {
		switch v {
		case 0:
			cnt0++
		case 1:
			cnt1++
		case 2:
			cnt2++
		default:
			return nil
		}
	}
	i := 0
	for cnt0 > 0 {
		nums[i] = 0
		i++
		cnt0--
	}
	for cnt1 > 0 {
		nums[i] = 1
		i++
		cnt1--
	}
	for cnt2 > 0 {
		nums[i] = 2
		i++
		cnt2--
	}
	return nums
}

// 思路2：类似于快排的思路，分两次分类
// 第一次将全部的0移动到最左侧，第二次将所有的1排到中间
func Classify2(arr []int) {
	idx := classify(arr, 0, len(arr)-1, 0)
	classify(arr, idx, len(arr)-1, 1)
}

// 分割，delima表示分割数
// 分割完一次，应该返回idx，则[startIdx, idx)应该全部==delima
func classify(arr []int, startIdx, endIdx int, delima int) int {
	if startIdx >= endIdx {
		return startIdx
	}

	for startIdx < endIdx {
		for startIdx < endIdx && arr[startIdx] == delima {
			startIdx++
		}
		for startIdx < endIdx && arr[endIdx] > delima {
			endIdx--
		}
		arr[startIdx], arr[endIdx] = arr[endIdx], arr[startIdx]
	}

	return startIdx
}
