package _021_merge_order_list

import (
	"fmt"
	"github.com/hq-cml/sward2offer/common"
	"testing"
)

func TestMerge(t *testing.T) {
	type args struct {
		head1 *common.ListNode
		head2 *common.ListNode
	}

	var l1 *common.ListNode
	var l2 *common.ListNode
	var l3 *common.ListNode
	tests := []struct {
		name string
		args args
		want *common.ListNode
	}{
		{
			name: "case1",
			args: args{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l1, l2, _ = gen()
			got := Merge(l1, l2)
			got.Foreach(common.NodePrint)
			l1, _, l3 = gen()
			got = Merge(l1, l3)
			got.Foreach(common.NodePrint)

			l1, l2, _ = gen()
			got2 := MergeRecurs(l1, l2)
			got2.Foreach(common.NodePrint)
			l1, _, l3 = gen()
			got = MergeRecurs(l1, l3)
			got.Foreach(common.NodePrint)

		})
		fmt.Println()
	}
}

func gen() (*common.ListNode, *common.ListNode, *common.ListNode) {
	var l1 *common.ListNode
	l1 = l1.PushNode(1)
	l1 = l1.PushNode(3)
	l1 = l1.PushNode(5)
	var l2 *common.ListNode
	l2 = l2.PushNode(2)
	l2 = l2.PushNode(4)
	l2 = l2.PushNode(6)
	l2 = l2.PushNode(8)
	l2 = l2.PushNode(10)
	var l3 *common.ListNode
	l3 = l3.PushNode(1)
	l3 = l3.PushNode(3)
	l3 = l3.PushNode(5)
	return l1, l2, l3
}
