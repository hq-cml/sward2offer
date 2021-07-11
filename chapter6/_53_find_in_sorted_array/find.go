/*
// 面试题53（一）：数字在排序数组中出现的次数
// 题目：统计一个数字在排序数组中出现的次数。例如输入排序数组{1, 2, 3, 3,
// 3, 3, 4, 5}和数字3，由于3在这个数组中出现了4次，因此输出4。

// 面试题53（二）：0到n-1中缺失的数字
// 题目：一个长度为n-1的递增排序数组中的所有数字都是唯一的，并且每个数字
// 都在范围0到n-1之内。在范围0到n-1的n个数字中有且只有一个数字不在该数组
// 中，请找出这个数字。

// 面试题53（三）：数组中数值和下标相等的元素
// 题目：假设一个单调递增的数组里的每个元素都是整数并且是唯一的。请编程实
// 现一个函数找出数组中任意一个数值等于其下标的元素。例如，在数组{-3, -1,
// 1, 3, 5}中，数字3和它的下标相等。

 * 在有序数组中的各种查找
 * 因为有序，所以二分查找是最直接的想法，并且是各种变种




这个题目最朴素的想法是从头遍历，但是显然不合要求
第二个思路是利用二分查找，找到第一个3，然后从两边扩散，不断减小和增大探测，直到找到边界。但是这么做的复杂度仍然是O(n)，所以仍然不和要求
作者给出的思路是，一个二分查找的变种，能够分别以O(lgN)的复杂度，找到3的上下边界。然后相减得到长度。



这题仍然是二分查找的变种，按照题目的描述，在缺失数据之前，所有的元素等于下标，从缺失元素开始，下标!=元素。所以，仍然可以利用二分查找来发现第一个元素!=下标的成员。


经过分析，这仍然是二分查找的变种


 */
package _53_find_in_sorted_array

//题目1：数字在排序数组中出现的次数
//这个题目最朴素的想法是从头遍历，但是显然不合要求
//第二个思路是利用二分查找，找到第一个3，然后从两边扩散，不断减小和增大探测，直到找到边界。但是这么做的复杂度仍然是O(n)，所以仍然不和要求
//作者给出的思路是，一个二分查找的变种，能够分别以O(lgN)的复杂度，找到3的上下边界。然后相减得到长度。
//难度：3*
func CalcCnt(arr []int, k int) int {
    if len(arr) == 0 {
        return -1
    }

    return findLastK(arr, k) -
        findFirstK(arr, k) + 1
}

func findFirstK(arr []int ,k int) int {
    i := 0
    j := len(arr) - 1
    firstIdx := -1
    for i <= j {
        mid := (i+j) / 2
        if arr[mid] == k {
            firstIdx = mid
            j = mid - 1
        } else if arr[mid] < k {
            i = mid + 1
        } else {
            j = mid - 1
        }
    }

    return firstIdx
}

func findLastK(arr []int ,k int) int {
    i := 0
    j := len(arr) - 1
    firstIdx := -1
    for i <= j {
        mid := (i+j) / 2
        if arr[mid] == k {
            firstIdx = mid
            i = mid + 1
        } else if arr[mid] < k {
            i = mid + 1
        } else {
            j = mid - 1
        }
    }

    return firstIdx
}

//题目2:找到0~n-1中缺失的数字
//仍然是二分查找的变种，按照题目的描述，在缺失数据之前，所有的元素等于下标，
//从缺失元素开始，下标!=元素。所以，仍然可以利用二分查找来发现第一个元素!=下标的成员。
func FindNumberLost(arr []int) int {
    i := 0
    j := len(arr) - 1
    firstIdx := -1
    for i <= j {
        mid := (i+j) / 2
        if arr[mid] == mid {
            i = mid + 1
        } else {
            firstIdx = mid
            j = mid - 1
        }
    }

    return firstIdx
}

//题目3:找到和下标相等的元素
//仍然是二分法查找的变种
func FindNumberSameAsIndex(arr []int) int {
    i := 0
    j := len(arr) - 1
    for i <= j {
        mid := (i+j) / 2
        if arr[mid] == mid {
            return mid
        } else if arr[mid] > mid {
            j = mid - 1
        } else {
            i = mid + 1
        }
    }

    return -1
}
