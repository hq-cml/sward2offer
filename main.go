package main

func Partition(arr []int, beg, end int) int {
	if len(arr) == 0 {
		return -1
	}
	if beg < 0 || end >= len(arr) || beg > end {
		return -1
	}

	val := arr[beg]

	for beg < end {
		for val <= arr[end] && beg < end {
			end--
		}
		if beg < end { //swap
			arr[beg], arr[end] = arr[end], arr[beg]
		}

		for val >= arr[beg] && beg < end {
			beg++
		}
		if beg < end { //swap
			arr[beg], arr[end] = arr[end], arr[beg]
		}
	}

	return beg
}

func main() {

}
