/*
 * tt面试题：链表的归并排序
 */
package _merge_order_list

import (
    "github.com/hq-cml/sward2offer/chapter3/_25_merge_order_list"
    "github.com/hq-cml/sward2offer/common"
)

//归并排序的思路：
//先拆，分而治之，然后在逐个合并
//合并的部分，用到了之前“有序链表合并”的代码
func MergeListSort(head *common.ListNode) *common.ListNode {
    if head == nil || head.Next == nil {
        return head
    }

    //首先找到链表的中间节点，一分为二
    p1 := head
    p2 := head.Next.Next
    for p2 != nil {
        p1 = p1.Next
        p2 = p2.Next
        if p2 != nil {
            p2 = p2.Next
        }
    }

    //此时p1是中间节点
    head2 := p1.Next
    p1.Next = nil

    return _25_merge_order_list.Merge(
        MergeListSort(head),
        MergeListSort(head2))
}