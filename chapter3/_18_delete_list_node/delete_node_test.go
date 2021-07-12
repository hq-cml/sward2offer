package _18_delete_list_node

import (
	"fmt"
	"github.com/hq-cml/sward2offer/common"
	"testing"
)

func TestDeleteNodeInO1(t *testing.T) {
	type args struct {
		head *common.ListNode
		del  *common.ListNode
	}

	l1 := common.NewList()

	l1 = l1.PushNode(1)
	l1 = l1.PushNode(2)
	node := &common.ListNode{
		Val:  3,
	}
	l1 = l1.PushNodePtr(node)
	l1 = l1.PushNode(4)

	tests := []struct {
		name string
		args args
		want *common.ListNode
	}{
		{
			name: "case1",
			args: args{
				head: l1,
				del:  node,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.args.head.Foreach(common.NodePrint)
			DeleteNodeInO1(tt.args.head, tt.args.del)
			tt.args.head.Foreach(common.NodePrint)
		})
	}
}

func TestDeleteDuplicate(t *testing.T) {
	type args struct {
		head *common.ListNode
	}

	l1 := common.NewList()
	l1 = l1.PushNode(1)
	l1 = l1.PushNode(2)
	l1 = l1.PushNode(2)
	l1 = l1.PushNode(4)

	l2 := common.NewList()
	l2 = l2.PushNode(1)
	l2 = l2.PushNode(2)
	l2 = l2.PushNode(2)
	l2 = l2.PushNode(2)
	l2 = l2.PushNode(2)
	l2 = l2.PushNode(2)
	l2 = l2.PushNode(4)
	l2 = l2.PushNode(4)
	l2 = l2.PushNode(4)
	l2 = l2.PushNode(5)

	tests := []struct {
		name string
		args args
		want *common.ListNode
	}{
		{
			name: "case1",
			args: args{
				head: l1,
			},
			want: nil,
		},
		{
			name: "case2",
			args: args{
				head: l2,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.args.head.Foreach(common.NodePrint)
			DeleteDuplicate(tt.args.head)
			tt.args.head.Foreach(common.NodePrint)
			fmt.Println()
		})
	}
}