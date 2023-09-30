package _148_list_sort

import (
	"github.com/hq-cml/sward2offer/common"
	"testing"
)

func TestMergeListSort(t *testing.T) {
	type args struct {
		head *common.ListNode
	}
	head := common.NewList()
	head = head.PushNode(6)
	head = head.PushNode(1)
	head = head.PushNode(5)
	head = head.PushNode(2)
	head = head.PushNode(4)
	head = head.PushNode(3)
	head.Foreach(common.NodePrint)

	tests := []struct {
		name string
		args args
		want *common.ListNode
	}{
		{
			name: "case1",
			args: args{
				head: head,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MergeListSort(tt.args.head)
			got.Foreach(common.NodePrint)
		})
	}
}

func TestMergeListSort2(t *testing.T) {
	type args struct {
		head *common.ListNode
	}
	head := common.NewList()
	head = head.PushNode(7)
	head = head.PushNode(6)
	head = head.PushNode(1)
	head = head.PushNode(5)
	head = head.PushNode(2)
	head = head.PushNode(4)
	head = head.PushNode(3)
	head.Foreach(common.NodePrint)

	tests := []struct {
		name string
		args args
		want *common.ListNode
	}{
		{
			name: "case1",
			args: args{
				head: head,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MergeListSort(tt.args.head)
			got.Foreach(common.NodePrint)
		})
	}
}
