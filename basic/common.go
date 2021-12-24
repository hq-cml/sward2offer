package basic

// 分区函数
// 返回：一次分区完毕之后的中间Idx值，保证Idx前面的均小于Idx元素值，后面的均大于
func Partition(arr []int, begIdx, endIdx int) int {
	// 边界校验
	if len(arr) == 0 {
		return -1
	}
	if begIdx > endIdx || begIdx < 0 || endIdx >= len(arr) {
		return -1
	}

	// 第一个值作为基准
	val := arr[begIdx]

	// 进行多轮次前后逼近，直到相遇
	for endIdx > begIdx {
		// 从后向前逼近
		for arr[endIdx] >= val && endIdx > begIdx {
			endIdx--
		}
		if endIdx != begIdx { // 说明必然 arr[endIdx] < val
			arr[begIdx], arr[endIdx] = arr[endIdx], arr[begIdx] //swap
		}

		// 从前向后逼近
		for arr[begIdx] <= val && endIdx > begIdx {
			begIdx++
		}
		if begIdx != endIdx { // 说明必然 arr[begIdx] > val
			arr[begIdx], arr[endIdx] = arr[endIdx], arr[begIdx] //swap
		}
	}

	return begIdx
}
