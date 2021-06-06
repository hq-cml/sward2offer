package basic

//分区函数
func Partition(arr []int, begIdx, endIdx int) (int){
    if len(arr) == 0 {
        return -1
    }
    if begIdx > endIdx || begIdx < 0 || endIdx >= len(arr) {
        return -1
    }

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
