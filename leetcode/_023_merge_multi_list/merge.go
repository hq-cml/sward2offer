/*
 * 多个有序列表，合并成1个
 * 输入：lists = [[1,4,5],[1,3,4],[2,6]]
 * 输出：[1,1,2,3,4,4,5,6]
 */
package _023_merge_multi_list

import (
	"github.com/hq-cml/sward2offer/common"
	"math"
)

// 思路：
// 1. 两个有序列表合并的升级版本，则可以先实现两个列表合并，然后循环处理
// 2. 直接模仿两个列表合并的思想，实现K个列表合并
func MergeMulti(arr []*common.ListNode) *common.ListNode {
	if len(arr) == 0 {
		return nil
	}
	min := math.MaxInt64
	minIdx := -1
	for idx, l := range arr {
		if l == nil {
			continue
		}
		if min > l.Val {
			min = l.Val
			minIdx = idx
		}
	}
	if minIdx == -1 {
		return nil
	}

	head := arr[minIdx]
	arr[minIdx] = arr[minIdx].Next
	head.Next = MergeMulti(arr)
	return head
}
