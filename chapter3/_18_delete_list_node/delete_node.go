/*
 * 面试题18（一）：在O(1)时间删除链表结点
 * 题目：给定单向链表的头指针和一个结点指针，定义一个函数在O(1)时间删除该结点。
 *
 *面试题18（二）：删除链表中重复的结点
 *题目：在一个排序的链表中，如何删除重复的结点？例如，在图3.4（a）中重复
 *结点被删除之后，链表如图3.4（b）所示。
 */

package _18_delete_list_node

import "github.com/hq-cml/sward2offer/common"

//O(1)复杂度删除指定节点
//思路：利用替换思想，巧妙删除。一些异常情况需要考虑：
//如果删除的是末尾节点：退化成普通删除
//仅单个节点：删除头，返回NULL
//总体来说，复杂度仍然是O(1)
func DeleteNodeInO1(head *common.ListNode, del *common.ListNode) (*common.ListNode) {
    if head == nil || del == nil {
        return head
    }

    if head == del { //待删除的是头结点
        return head.Next
    } else if del.Next == nil { //待删除的节点是末尾
        p1 := head
        p2 := head.Next
        for p2 != del {
            p1 = p2
            p2 = p2.Next
        }
        p1.Next = nil
    } else { //替换
        del.Val = del.Next.Val
        del.Next = del.Next.Next
    }
    return head
}

//删除有序链表的重复节点
//这道题的作者的意思是删除重复的节点，一个不留，我觉得不合理，改成了过滤重复。
func DeleteDuplicate(head *common.ListNode) (*common.ListNode){
    if head == nil || head.Next == nil {
        return head
    }

    p1 := head
    p2 := head.Next
    for {
        for p2 != nil && p1.Val == p2.Val {
            p2 = p2.Next
            p1.Next = p2
        }
        if p2 == nil {
            break
        }
        p1 = p2
        p2 = p2.Next
    }

    return head
}