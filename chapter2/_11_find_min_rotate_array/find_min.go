/*
 // 面试题11：旋转数组的最小数字
// 题目：把一个数组最开始的若干个元素搬到数组的末尾，我们称之为数组的旋转。
// 输入一个递增排序的数组的一个旋转，输出旋转数组的最小元素。例如数组
// {3, 4, 5, 1, 2}为{1, 2, 3, 4, 5}的一个旋转，该数组的最小值为1。

最直观的是遍历，但显然不是好的答案。
一个变种的二分查找，利用递归，需要注意一些边界的情况。
 */

package _11_find_min_rotate_array

import (
    "errors"
    "math"
)

func FindMin(arr []int) (int, error) {
    if len(arr) == 0 {
        return -1, errors.New("Invalid arr")
    }
    if len(arr) == 1 {
        return arr[0], nil
    }

    p1 := 0
    p2 := len(arr) - 1

    if arr[p1] < arr[p2] { //特例，旋转0个（未旋转）
        return arr[p1], nil
    }

    for p1 < (p2-1) {
        mid := p1 + (p2 - p1) >> 1
        if arr[p1] == arr[mid] && arr[p2] == arr[mid] {
            //特例：1, 1, 1, 0, 1
            return findMinInOrder(arr, p1, p2), nil
        }
        if arr[mid] <= arr[p2] {
            p2 = mid
        } else {
            p1 = mid
        }
    }
    return arr[p2], nil
}

func findMinInOrder(arr []int, idx1, idx2 int) int {
    min := math.MaxInt64
    for i := idx1; i < idx2; i++ {
        if arr[i] < min {
            min = arr[i]
        }
    }
    return min
}