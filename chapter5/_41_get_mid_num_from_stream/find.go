/*
 * 面试题41：数据流中的中位数
 * 题目：如何得到一个数据流中的中位数？如果从数据流中读出奇数个数值，那么
 * 中位数就是所有数值排序之后位于中间的数值。如果从数据流中读出偶数个数值，
 * 那么中位数就是所有数值排序之后中间两个数的平均值。
 */
package _41_get_mid_num_from_stream

import (
	"errors"
	"fmt"
	"github.com/hq-cml/sward2offer/common"
)

//思路：关键核心点在于维护一个容器，存储已经读出来的数据，且能够快速获取到中位数
//思路1：
// 选用一个排序数组，插入排序，时间复杂度O（N^2），但是读取复杂度是O(1)。
//思路2：
// 利用两个堆
// 左侧大顶堆，右侧小顶堆，先左后右，轮次放置，左侧整体都要小于右侧
// 同时两个堆数量维持均衡，即，len(左堆) - len(右堆) <= 1
// 则中位数只会是是左侧堆头（奇数个时）或者两个堆头的平均数（偶数个时）
// 时间复杂度O（N * LogN）
func FindMidNum(arr []int) (float64, error) {
	if len(arr) == 0 {
		return 0, errors.New("Invalid input")
	}

	var left bool
	leftHeap := common.NewHeapInt(nil, false) //左堆是大顶堆
	rightHeap := common.NewHeapInt(nil, true) //右堆是小顶堆

	for k, v := range arr {
		if k%2 == 0 {
			left = true
		} else {
			left = false
		}

		//始终保持：len(左堆) - len(右堆) <= 1
		//如果要放入左侧，则应该先检查右侧；反之，应该先检查左侧
		if left {
			//初步计划，放左侧
			if leftHeap.Len() == 0 { //左堆为空，表示刚刚开始，则直接push
				leftHeap.Push(v)
			} else if v > rightHeap.Top() { //待放入的值，大于右侧最小值，所以他应该放入右侧，所以只能从右侧换一个出来放入左侧（不会出现rightHeap.Top==nil的情况）
				rightHeap.Push(v)
				leftHeap.Push(rightHeap.Pop())
			} else { //待放入的值，小于右侧最小值，则可以直接放入左侧
				leftHeap.Push(v)
			}
		} else {
			//初步计划，放右侧
			if v < leftHeap.Top() { //待放入值，小于左侧的最大值，所以他应该放到左侧，所以只能从左侧换一个值出来放入右侧（不会出现leftHeap.Top==nil的情况）
				leftHeap.Push(v)
				rightHeap.Push(leftHeap.Pop())
			} else { //待放入的值，大于左侧最大值，则可以直接放入右侧
				rightHeap.Push(v)
			}
		}
	}

	//求中位数
	//两个堆数量维持均衡，即，len(左堆) - len(右堆) <= 1
	if leftHeap.Len() == rightHeap.Len() {
		return float64(rightHeap.Top()+leftHeap.Top()) / 2, nil
	} else if leftHeap.Len()-1 == rightHeap.Len() {
		return float64(leftHeap.Top()), nil
	}
	return 0, fmt.Errorf("Wrong! leftMax:%v, rightMin:%v",
		leftHeap.Len(), rightHeap.Len())
}
