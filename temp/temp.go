package temp

// 利用二分查找找到上边界
func findFirstK(arr []int, k int) int {
	beg, end := 0, len(arr)-1
	var idx int = -1
	for beg <= end {
		mid := (end-beg)/2 + beg
		if arr[mid] == k {
			idx = mid
			end = end - 1
		} else if arr[mid] > k {
			end = mid - 1
		} else {
			beg = mid + 1
		}
	}
	return idx
}
