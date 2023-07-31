package basic

import (
	"github.com/hq-cml/sward2offer/common"
)

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

//  Quick Sort
func QuickSort(arr []int) {
	length := len(arr)
	if length == 0 || length == 1 {
		return
	}

	// 对arr进行左右的分区
	idx := partition(arr)
	if idx == -1 {
		// 不应该出现的异常情况
		return
	}

	// 分别递归左右
	if idx != 0 {
		QuickSort(arr[:idx])
	}
	if idx != length - 1 {
		QuickSort(arr[idx+1:])
	}
}

func partition(arr []int) int {
	// 边界校验
	if len(arr) == 0 {
		return -1
	}
	if len(arr) == 1 {
		return 0
	}

	// 第一个值作为基准
	val := arr[0]
	beg := 0
	end := len(arr) -1

	// 进行多轮次向心逼近，直到相遇
	for beg < end {
		// 从后向前逼近
		for beg < end && arr[end]>=val {
			end --
		}
		arr[beg], arr[end] = arr[end], arr[beg]
		// 从前向后逼近
		for beg < end && arr[beg]<=val {
			beg ++
		}
		arr[beg], arr[end] = arr[end], arr[beg]
	}
	// beg和end已经相遇
	return beg
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

// 选择排序
// 数组长度N，执行N-1轮选择
// 第i轮（i=0,1,...N-2)，选择[i, end]最小的值，放在i的位置，下一轮从i+1开始
func SelectSort(arr []int) {
	n := len(arr)
	if n==0 || n == 1 {
		return
	}

	//执行N-1轮选择
	for i:=0; i<n-1; i++ {
		minIdx := i
		for j:=i; j<n; j++ {
			if arr[minIdx] > arr[j] {
				minIdx = j
			}
		}
		// 换到i的位置
		arr[i], arr[minIdx] = arr[minIdx], arr[i]
	}
}

//插入排序
//从第一个元素开始，该元素可以认为已经被排序；
//认为第i（i=1,...N-1)个元素之前都是排序好的，
//则将i+1个元素，逐个试探出前i个元素的适合位置，插入进去
//循环往复
func InsertSort(arr []int) {
	length := len(arr)
	if length <= 1 {
		return
	}
	// 初始认为，第一个元素自然是一个有序的，从第二个开始插入
	for i:=1; i<length; i++ {
		for j:=i-1; j>=0; j-- {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

// 归并排序，递归
func MergeSort(arr []int) []int {
	length := len(arr)
	if length <= 1 {
		return arr
	}

	midIdx := length/2
	arr1 := MergeSort(arr[0:midIdx])
	arr2 := MergeSort(arr[midIdx:])

	var ret []int
	for len(arr1) > 0 && len(arr2)>0 {
		if arr1[0]<arr2[0] {
			ret = append(ret, arr1[0])
			arr1 = arr1[1:]
		} else {
			ret = append(ret, arr2[0])
			arr2 = arr2[1:]
		}
	}
	if len(arr1) > 0 {
		ret = append(ret, arr1...)
	}
	if len(arr2) >0 {
		ret = append(ret, arr2...)
	}
	return ret
}

// 堆排序，自实现
// 基本步骤，就是如下的无限循环：
// 建堆，交换....建堆，交换...建堆，交换...建堆，交换
func SelfHeapSort(arr []int) {
	l := len(arr)
	// 循环建堆，每次建堆完成，首元素都是最大的
	// 此时将首元素和最后一个元素交换，下一轮就可以排除一个元素了
	for i:=1; i<l; i++ {
		MaxHeapify(arr, l - i)
		// 将最大元素放置目前的最后，下一轮建堆将不再考虑这个最后一个
		arr[0], arr[l-i] = arr[l-i], arr[0]
	}
}

// 一般性数组建堆，无前提条件
// heapify是有前提的，但是一般的数组没有这个前提
// 所以从数组的最后一个元素开始建堆，这样巧妙解决这个问题
func MaxHeapify(arr []int, boundaryIdx int) {
	for i:=len(arr)-1; i>=0; i-- {
		heapify(arr, i, boundaryIdx)
	}
}

// 建堆，递归方式建立大顶堆
// 默认前提是除了堆顶currIdx，其他的已经子节点已经完成建堆
// 参数currIdx：堆顶元素下标
// 参数boundaryIdx：作为堆结构边界，越界保护
func heapify(arr []int, currIdx, boundaryIdx int) {
	// 递归结束条件
	if currIdx >= boundaryIdx {
		return
	}

	leftIdx := currIdx*2 + 1 //左下标
	rightIdx := currIdx*2 + 2 //右下标

	maxIdx := currIdx
	if leftIdx <= boundaryIdx && arr[maxIdx] < arr[leftIdx] {
		maxIdx = leftIdx
	}
	if rightIdx <= boundaryIdx && arr[maxIdx] < arr[rightIdx] {
		maxIdx = rightIdx
	}

	// 如果currIdx不满足堆的条件，意味着需要进行交换
	if maxIdx != currIdx {
		arr[maxIdx], arr[currIdx] = arr[currIdx], arr[maxIdx]

		// 递归的继续
		heapify(arr, maxIdx, boundaryIdx)
	}
}






















