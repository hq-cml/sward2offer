package common

import "container/heap"

const (
	HEAP_TYPE_MIN = iota //0，小顶堆
	HEAP_TYPE_MAX        //1，大顶堆
)

//堆实现：typ标志是大顶堆还是小顶堆
type HeapInt struct {
	dataArr []int
	typ     int
}

/*
 * New...
 */
func NewMinHeapInt(arr []int) *HeapInt {
	h := &HeapInt{
		typ:     HEAP_TYPE_MIN,
		dataArr: arr,
	}
	heap.Init(h) //初始化
	return h
}

func NewMaxHeapInt(arr []int) *HeapInt {
	h := &HeapInt{
		typ:     HEAP_TYPE_MAX,
		dataArr: arr,
	}
	heap.Init(h) //初始化
	return h
}

/*
 * 实现heap.Interface
 */
func (h *HeapInt)Len() int {
	return len(h.dataArr)
}

func (h *HeapInt)Swap(i, j int) {
	h.dataArr[i], h.dataArr[j] = h.dataArr[j], h.dataArr[i]
}

func (h *HeapInt)Less(i, j int) bool {
	if h.typ == HEAP_TYPE_MIN {
		return h.dataArr[i] < h.dataArr[j]
	} else {
		return h.dataArr[i] > h.dataArr[j]
	}
}

//实现heap.Interface，对外莫用
// add x as element Len()
func (h *HeapInt) Push(x interface{}) {
	h.dataArr = append(h.dataArr, x.(int))
}

//实现heap.Interface，对外莫用
// remove and return element Len() - 1.
func (h *HeapInt) Pop() interface{} {
	x := h.dataArr[len(h.dataArr)-1]
	h.dataArr = h.dataArr[0 : len(h.dataArr)-1]
	return x
}

/*
 * 其他函数，对外用
 */
func (h *HeapInt)Get(idx int) int {
	return h.dataArr[idx]
}

func (h *HeapInt)Put(i int) {
	heap.Push(h, i)
}

func (h *HeapInt)Remove(idx int) int {
	x := h.Get(idx)
	heap.Remove(h, idx)
	return x
}

func (h *HeapInt)Arr() []int {
	arr := make([]int, h.Len())
	copy(arr, h.dataArr)
	return arr
}