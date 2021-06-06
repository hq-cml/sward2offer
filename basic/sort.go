package basic

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
