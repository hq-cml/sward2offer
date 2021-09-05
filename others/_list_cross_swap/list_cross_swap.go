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
func CrossSwap(l *common.ListNode) *common.ListNode {
	if l == nil || l.Next == nil {
		return l
	}

	head := l.Next //先将新的head保存

	p1 := l
	p2 := p1.Next
	p3 := p2.Next //保留第三个节点，防止断裂
	for p1 != nil && p2 != nil {
		p2.Next = p1 //节点1和节点2互换指针
		p1.Next = p3 //p1的指向，暂时指向p3（在后面p3和p4交换之后，p1可能还要指向p4）
		tmp := p1    //暂时保存p1的指针，后面p1指向p4的时候，会需要用到

		p1 = p3 //p1后移
		if p1 != nil {
			p2 = p1.Next //p2后移，其实指向了p4

			if p2 != nil { //如果p4非空，则需要将之前的p1指向p4
				tmp.Next = p2
				p3 = p2.Next //p3继续后移，防止断裂
			}
		}
	}

	return head
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
