package _160_list_first_common_node

import (
	"fmt"
	"github.com/hq-cml/sward2offer/common"
	"testing"
)

func TestFirstCommonNode(t *testing.T) {
	type args struct {
		l1 *common.ListNode
		l2 *common.ListNode
	}
	l1 := common.NewList()
	l2 := common.NewList()

	l1 = l1.PushNode(1)
	l1 = l1.PushNode(2)
	l1 = l1.PushNode(3)

	l2 = l2.PushNode(10)
	l2 = l2.PushNode(20)

	node := &common.ListNode{
		Val:  4,
		Next: nil,
	}

	l1 = l1.PushNodePtr(node)
	l2 = l2.PushNodePtr(node)

	l1 = l1.PushNode(5)
	l1 = l1.PushNode(6)

	l1.Foreach(common.NodePrint)
	l2.Foreach(common.NodePrint)

	tests := []struct {
		name string
		args args
		want *common.ListNode
	}{
		{
			name: "case1",
			args: args{
				l1: l1,
				l2: l2,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FirstCommonNode(tt.args.l1, tt.args.l2)
			fmt.Println("Commont:", got.Val)
		})
	}
}
