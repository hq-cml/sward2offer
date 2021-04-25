package _04_find_double_arr

/*
 1  2  8  9
 2  4  9  12
 4  7  10 13
 6  8  11 15
 */
func FindDoubleArray(arr []int, row, col int, needFind int) bool {
    if len(arr) != row * col {
        return false
    }

    i := 0
    j := col - 1
    for i < row || j >= 0 {
        curr := arr[i*row + j]
        if curr == needFind {
            return true
        } else if curr > needFind {
            j --
        } else {
            i ++
        }
    }

    return false
}