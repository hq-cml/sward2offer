package basic

//分区函数
//返回：一次分区完毕之后的中间Idx值，保证Idx前面的均小于Idx元素值，后面的均大于
func Partition(arr []int, begIdx, endIdx int) int {
	if len(arr) == 0 {
		return -1
	}
	if begIdx > endIdx || begIdx < 0 || endIdx >= len(arr) {
		return -1
	}

	val := arr[begIdx]

	for endIdx > begIdx {
		for arr[endIdx] >= val && endIdx > begIdx {
			endIdx--
		}
		//swap
		arr[begIdx], arr[endIdx] = arr[endIdx], arr[begIdx]

		for arr[begIdx] <= val && endIdx > begIdx {
			begIdx++
		}
		//swap
		arr[begIdx], arr[endIdx] = arr[endIdx], arr[begIdx]
	}

	return begIdx
}
