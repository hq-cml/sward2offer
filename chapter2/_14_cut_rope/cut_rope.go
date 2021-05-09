package _14_cut_rope

//动态规划，自底向上
func MaxCut(length int) int {
    //一些特殊情况
    if length <= 1 {
        return 0
    }
    if length == 2 {
        return 1
    }
    if length == 3 {
        return 2
    }

    arr := make([]int, length+1)
    //当长度小于3的时候，切不如不切
    arr[1] = 1
    arr[2] = 2
    arr[3] = 3

    for i := 4; i <= length; i++ {
        max := 0
        for j := 1; j < i; j ++ {
            if arr[j] * arr[i-j] > max {
                max = arr[j] * arr[i-j]
            }
        }
        arr[i] = max
    }

    return arr[length]
}