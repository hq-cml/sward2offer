/*
// 面试题4：二维数组中的查找
// 题目：在一个二维数组中，每一行都按照从左到右递增的顺序排序，每一列都按
// 照从上到下递增的顺序排序。请完成一个函数，输入这样的一个二维数组和一个
// 整数，判断数组中是否含有该整数。
 */
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