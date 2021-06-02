package basic

import "fmt"

func QuickSort(arr []int) {
    if len(arr) == 0{
        return
    }

    mid := Partition(arr, 0, len(arr)-1)
    if mid == -1 {
        return
    }
    fmt.Println("X----",mid)
    QuickSort(arr[0:mid])
    //if mid < len(arr) -1 {
        QuickSort(arr[mid+1:])
   // }

}

//一次分离
func Partition(arr []int, begIdx, endIdx int) (int){
    if len(arr) == 0 {
        return -1
    }
    //if begIdx > endIdx || begIdx < 0 || endIdx >= len(arr) {
    //    return -1
    //}

    val := arr[begIdx]

    for endIdx > begIdx {
        for arr[endIdx] >= val && endIdx > begIdx {
            endIdx --
            continue
        }
        //swap
        arr[begIdx], arr[endIdx] = arr[endIdx], arr[begIdx]

        for arr[begIdx] <= val && endIdx > begIdx {
            begIdx ++
            continue
        }
        //swap
        arr[begIdx], arr[endIdx] = arr[endIdx], arr[begIdx]
    }

    return begIdx
}