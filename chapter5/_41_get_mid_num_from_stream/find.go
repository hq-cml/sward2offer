/*
// 面试题41：数据流中的中位数
// 题目：如何得到一个数据流中的中位数？如果从数据流中读出奇数个数值，那么
// 中位数就是所有数值排序之后位于中间的数值。如果从数据流中读出偶数个数值，
// 那么中位数就是所有数值排序之后中间两个数的平均值。


关键核心点在于维护一个已经读出来的数据的容器的选择，我觉得面试中，就直接选用一个排序数组就行了，作者选择了堆，这些都太复杂，不太适合再面试的时候搞的。
用排序的数组，插入排序，时间复杂度O（N），但是读取复杂度是O(1)。
 */
package _41_get_mid_num_from_stream

import (
    "errors"
    "fmt"
    "github.com/hq-cml/sward2offer/common"
)

//利用两个堆
// 左侧大顶堆，右侧小顶堆，先左后右，轮次放置，同时两个堆数量维持均衡，数量差不超过1
// 则中位数只会是是左侧堆头或者两个堆头的平均数
func FindMidNum1(arr []int) (float64, error) {
    if len(arr) == 0 {
        return 0, errors.New("Invalid input")
    }

    var left bool
    maxLeftHeap := common.NewHeapInt(nil, false)
    minRightHeap := common.NewHeapInt(nil, true)

    for k, v := range arr {
        if k % 2 == 0 {
            left = true
        } else {
            left = false
        }

        if left {
            if maxLeftHeap.Len() == 0 {
                maxLeftHeap.Push(v)
            } else if v <= maxLeftHeap.Top() {
                maxLeftHeap.Push(v)
            } else {
                minRightHeap.Push(v)
                maxLeftHeap.Push(minRightHeap.Pop())
            }
        } else {
            if v >= maxLeftHeap.Top() {
                minRightHeap.Push(v)
            } else {
                maxLeftHeap.Push(v)
                minRightHeap.Push(maxLeftHeap.Pop())
            }
        }
    }

    if maxLeftHeap.Len() == minRightHeap.Len() {
        return float64(minRightHeap.Top() + maxLeftHeap.Top())/2, nil
    } else if maxLeftHeap.Len()-1 == minRightHeap.Len() {
        return float64(maxLeftHeap.Top()), nil
    }
    return 0, fmt.Errorf("Wrong! leftMax:%v, rightMin:%v",
        maxLeftHeap.Len(), minRightHeap.Len())
}