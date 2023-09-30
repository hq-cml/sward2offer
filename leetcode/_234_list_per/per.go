/**
 * 链表回文
 */
package _234_list_per

import "github.com/hq-cml/sward2offer/common"

// 下面的官方思路，要么太傻缺，要么想不到，自己实现一个，依赖辅助栈
func IsPalindrome0(head *common.ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}

	var stk []int
	p := head
	for p != nil {
		stk = append(stk, p.Val)
		p = p.Next
	}

	p = head
	for p != nil {
		top := stk[len(stk)-1]
		if top != p.Val {
			return false
		}

		// pop
		stk = stk[:len(stk)-1]
		p = p.Next
	}
	return true
}

// 思路1：快慢指针找到中点，然后翻转后半部分，再然后来进行对比
func IsPalindrome(head *common.ListNode) bool {
	if head == nil {
		return true
	}

	// 找到前半部分链表的尾节点并反转后半部分链表
	firstHalfEnd := endOfFirstHalf(head)
	secondHalfStart := reverseList(firstHalfEnd.Next)

	// 判断是否回文
	p1 := head
	p2 := secondHalfStart
	result := true
	for result && p2 != nil {
		if p1.Val != p2.Val {
			result = false
		}
		p1 = p1.Next
		p2 = p2.Next
	}

	// 还原链表并返回结果
	firstHalfEnd.Next = reverseList(secondHalfStart)
	return result
}

func reverseList(head *common.ListNode) *common.ListNode {
	var prev, cur *common.ListNode = nil, head
	for cur != nil {
		nextTmp := cur.Next
		cur.Next = prev
		prev = cur
		cur = nextTmp
	}
	return prev
}

func endOfFirstHalf(head *common.ListNode) *common.ListNode {
	fast := head
	slow := head
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}

// 思路2：递归，这个递归思路真是惊为天人，可以参考链表倒着打印这块思路来理解
func IsPalindromeRecurse(head *common.ListNode) bool {
	frontPointer := head
	var recursivelyCheck func(*common.ListNode) bool
	recursivelyCheck = func(curNode *common.ListNode) bool {
		if curNode != nil {
			if !recursivelyCheck(curNode.Next) {
				return false
			}
			if curNode.Val != frontPointer.Val {
				return false
			}
			frontPointer = frontPointer.Next
		}
		return true
	}
	return recursivelyCheck(head)
}
