package temp

import (
	"github.com/hq-cml/sward2offer/common"
)

func Reverse(head *common.ListNode) *common.ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	p1 := head
	p2 := head.Next
	head.Next = nil
	for p2 != nil {
		p3 := p2.Next
		p2.Next = p1
		p1 = p2
		p2 = p3
	}
	return p1
}
