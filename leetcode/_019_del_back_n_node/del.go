/*
 *  删除列表的倒数第n个节点（倒数第1个节点，是最后一个节点）
 */
package _019_del_back_n_node

import (
	"errors"
	"github.com/hq-cml/sward2offer/common"
)

// 双指针策略
// 删除倒数第N个，其实应该找到倒数第N+1个，然后删除它的下一个节点
func DelBackN(head *common.ListNode, n uint32) (*common.ListNode, error) {
	if head == nil {
		return nil, errors.New("nil ")
	}
	if n <= 0 {
		return nil, errors.New("n <= 0")
	}
	n = n + 1 // 倒数第N+1
	p1, p2 := head, head
	for n > 0 && p2 != nil {
		p2 = p2.Next
		n--
	}

	if n == 0 {
		// 平移P2和P1
		for p2 != nil {
			p1 = p1.Next
			p2 = p2.Next
		}
		p1.Next = p1.Next.Next
		return head, nil
	} else if n == 1 {
		// 倒数长度为n，删除倒数第n个，即删除第一个
		return head.Next, nil
	} else {
		return nil, errors.New("n is to big")
	}
}
