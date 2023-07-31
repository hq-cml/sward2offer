/*
 * mb面试题
 * 链表节点交叉变换
 * 1->2->3->4->5
 * 2->1->4->3->5
 */
package _list_cross_swap

import "github.com/hq-cml/sward2offer/common"

//思路1：利用指针的来回变更指向实现，难度不大，主要是需要想清楚
//难度：4*
func CrossSwap(head *common.ListNode) *common.ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	ret := head.Next // 最终结果，先预存等待返回

	p1 := head
	p2 := p1.Next

	for p1 != nil && p2 != nil {
		// p3保存后续处理节点
		p3 := p2.Next

		// 反转1和2
		p2.Next = p1
		// p1有三种情况，取决于原来p2的情况
		if p3 == nil {
			p1.Next = nil // p2已经是结尾了
		} else if p3.Next == nil {
			p1.Next = p3 // p2是倒数第二个节点
		} else {
			p1.Next = p3.Next // p2是普通节点
		}

		// 后移
		p1 = p3
		if p1 != nil {
			p2 = p1.Next
		}
	}
	return ret
}

//思路2：一个更加聪明的办法，直接交换节点的值，节点的指针不动（类似于O(1)复杂度删除节点的思路）
//难度：3*
func CrossSwapSmart(l *common.ListNode) *common.ListNode {
	if l == nil || l.Next == nil {
		return l
	}

	p1 := l
	p2 := p1.Next

	for p1 != nil && p2 != nil {
		p1.Val, p2.Val = p2.Val, p1.Val

		p1 = p2.Next
		if p1 != nil {
			p2 = p1.Next
		}
	}

	return l
}
