/*
 * 面试题25：合并两个排序的链表
 * 题目：输入两个递增排序的链表，合并这两个链表并使新链表中的结点仍然是按
 * 照递增排序的。例如输入图3.11中的链表1和链表2，则合并之后的升序链表如链
 * 表3所示。
 */
package _25_merge_order_list

import "github.com/hq-cml/sward2offer/common"

//思路1：普通写法
//难度：2*
func Merge(head1, head2 *common.ListNode) *common.ListNode {
	if head1 == nil {
		return head2
	}
	if head2 == nil {
		return head1
	}

	//处理确定第一个节点
	var head *common.ListNode
	if head1.Val <= head2.Val {
		head = head1
		head1 = head1.Next
	} else {
		head = head2
		head2 = head2.Next
	}

	//两条队列均非空的时候
	p := head
	for head1 != nil && head2 != nil {
		if head1.Val <= head2.Val {
			p.Next = head1
			head1 = head1.Next
		} else {
			p.Next = head2
			head2 = head2.Next
		}
		p = p.Next
	}

	//剩余节点处理
	if head1 != nil {
		p.Next = head1
	} else {
		p.Next = head2
	}

	return head
}

//思路2：递归，很精妙的递归写法，但是我觉得实际面试中，很难这么想到。
//难度：3*
func MergeRecurs(head1, head2 *common.ListNode) *common.ListNode {
	if head1 == nil {
		return head2
	}
	if head2 == nil {
		return head1
	}

	var head *common.ListNode
	if head1.Val <= head2.Val {
		head = head1
		head1 = head1.Next
	} else {
		head = head2
		head2 = head2.Next
	}

	//递归处理余下部分
	head.Next = MergeRecurs(head1, head2) //妙！
	return head
}
