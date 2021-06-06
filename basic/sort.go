package basic

import "github.com/hq-cml/sward2offer/common"

//快速排序
func QuickSort(arr []int) {
    if len(arr) == 0{
        return
    }

    mid := Partition(arr, 0, len(arr)-1)
    if mid == -1 {
        return
    }

    //递归
    QuickSort(arr[0:mid])
    QuickSort(arr[mid+1:])
}

//堆排序
func HeapSort(arr []int) []int {
    if len(arr) == 0{
        return arr
    }

    s := []int{}
    heap := common.NewHeapInt(s, true)
    for _, v := range arr {
        heap.Push(v)
    }

    ret := []int{}
    for heap.Len() > 0 {
        ret = append(ret, heap.Remove(0))
    }
    return ret
}