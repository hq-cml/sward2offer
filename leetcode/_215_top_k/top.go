/**
 * Top k，返回数组中，第k大的数字
 */
package _215_top_k

import (
	"sort"
)

// 思路1：直接排序
func FindKthLargest(nums []int, k int) int {
	if len(nums) < k {
		return -1
	}
	sort.Ints(nums)
	return nums[len(nums)-k]
}

// 思路2：堆排序
func FindKthHeap(nums []int, k int) int {
	heapSize := len(nums)
	// 大顶堆
	buildMaxHeap(nums, heapSize)
	// 现在有一个大顶堆，需要一个第k大的节点，则只需要
	// 将head移除k-1次（每次移除后需要重新建堆）
	// 则剩下的head就是第k个最大数
	for i := len(nums) - 1; i >= len(nums)-k+1; i-- {
		nums[0], nums[i] = nums[i], nums[0]
		heapSize--
		maxHeapify(nums, 0, heapSize)
	}
	return nums[0]
}

// 构建一个大顶堆
func buildMaxHeap(a []int, heapSize int) {
	for i := heapSize / 2; i >= 0; i-- {
		maxHeapify(a, i, heapSize)
	}
}

// 在[i+1:]已经构建完成的情况下，构建i节点成为一个大顶堆
func maxHeapify(a []int, i, heapSize int) {
	l, r, largest := i*2+1, i*2+2, i
	if l < heapSize && a[l] > a[largest] {
		largest = l
	}
	if r < heapSize && a[r] > a[largest] {
		largest = r
	}
	if largest != i {
		a[i], a[largest] = a[largest], a[i]
		maxHeapify(a, largest, heapSize)
	}
}
