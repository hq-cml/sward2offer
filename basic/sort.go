package basic

import "github.com/hq-cml/sward2offer/common"

// 冒泡排序
// 1. 总轮数是len-1，所以外层应该是1开始，到len-1
// 2. 第一轮，内层比较次数也是len-1，内层是从0开始，到len-i
func BubbleSort(arr []int) {
	if len(arr) == 0 {
		return
	}

	for i := 1; i < len(arr); i++ {
		for j := 0; j < len(arr)-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j] //swap
			}
		}
	}
}

// 快速排序
func QuickSort(arr []int) {
	if len(arr) == 0 {
		return
	}

	mid := Partition(arr, 0, len(arr)-1)
	if mid == -1 { //mid==-1，所以已经到了极限
		return
	}

	//递归
	QuickSort(arr[0:mid])  //不包括mid
	QuickSort(arr[mid+1:]) //不包括mid
}

// 堆排序
func HeapSort(arr []int) []int {
	if len(arr) == 0 {
		return arr
	}

	s := []int{}
	heap := common.NewHeapInt(s, true) //小顶堆
	for _, v := range arr {
		heap.Push(v)
	}

	ret := []int{}
	for heap.Len() > 0 {
		ret = append(ret, heap.Remove(0)) //逐个从小顶堆移出
	}
	return ret
}
