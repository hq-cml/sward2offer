package common

import "container/heap"

const (
	HEAP_TYPE_MIN = iota //0，小顶堆
	HEAP_TYPE_MAX        //1，大顶堆
)

//对外
type HeapInt struct{
	*heapInt
}

//min:是否是小顶堆
func NewHeapInt(arr []int, min bool) *HeapInt {
	var p *heapInt
	if min {
		p = newMinHeapInt(arr)
	} else {
		p = newMaxHeapInt(arr)
	}
	return &HeapInt{
		p,
	}
}

func (h *HeapInt)Len() int {
	return h.heapInt.Len()
}

func (h *HeapInt)Push(i int) {
	heap.Push(h.heapInt, i)
}

func (h *HeapInt)Pop() int {
	x := h.Get(0)
	heap.Remove(h.heapInt, 0)
	return x
}

func (h *HeapInt)Top() int {
	return h.Get(0)
}

func (h *HeapInt)Get(idx int) int {
	return h.heapInt.dataArr[idx]
}

func (h *HeapInt)Remove(idx int) int {
	x := h.Get(idx)
	heap.Remove(h.heapInt, idx)
	return x
}

func (h *HeapInt)All() []int {
	arr := make([]int, h.Len())
	copy(arr, h.heapInt.dataArr)
	return arr
}

/*
 * 实际的堆实现：typ标志是大顶堆还是小顶堆
 */
type heapInt struct {
	dataArr []int
	typ     int
}

/*
 * New...
 */
func newMinHeapInt(arr []int) *heapInt {
	h := &heapInt{
		typ:     HEAP_TYPE_MIN,
		dataArr: arr,
	}
	heap.Init(h) //初始化
	return h
}

func newMaxHeapInt(arr []int) *heapInt {
	h := &heapInt{
		typ:     HEAP_TYPE_MAX,
		dataArr: arr,
	}
	heap.Init(h) //初始化
	return h
}

/*
 * 实现heap.Interface
 */
func (h *heapInt)Len() int {
	return len(h.dataArr)
}

func (h *heapInt)Swap(i, j int) {
	h.dataArr[i], h.dataArr[j] = h.dataArr[j], h.dataArr[i]
}

func (h *heapInt)Less(i, j int) bool {
	if h.typ == HEAP_TYPE_MIN {
		return h.dataArr[i] < h.dataArr[j]
	} else {
		return h.dataArr[i] > h.dataArr[j]
	}
}

//实现heap.Interface
// add x as element Len()
func (h *heapInt) Push(x interface{}) {
	h.dataArr = append(h.dataArr, x.(int))
}

//实现heap.Interface
// remove and return element Len() - 1.
func (h *heapInt) Pop() interface{} {
	x := h.dataArr[len(h.dataArr)-1]
	h.dataArr = h.dataArr[0 : len(h.dataArr)-1]
	return x
}
