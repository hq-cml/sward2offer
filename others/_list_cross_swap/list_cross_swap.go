/*
 * mb面试题
 * 链表节点交叉变换
 * 1->2->3->4->5
 * 2->1->4->3->5
 */
package _list_cross_swap

import "github.com/hq-cml/sward2offer/common"

func CrossSwap(l *common.ListNode) (*common.ListNode) {
	if l == nil || l.Next == nil {
		return l
	}

	head := l.Next //先将新的head保存

	p1 := l
	p2 := p1.Next
	p3 := p2.Next
	for p1 != nil && p2 != nil {
		p2.Next = p1

	}


	return head
}