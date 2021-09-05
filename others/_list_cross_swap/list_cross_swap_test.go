package _list_cross_swap

import (
	"github.com/hq-cml/sward2offer/common"
	"testing"
)

func TestCrossSwap(t *testing.T) {
	type args struct {
		l *common.ListNode
	}
	l1 := common.NewList()
	l1 = l1.PushNode(1)
	l1 = l1.PushNode(2)

	l2 := common.NewList()
	l2 = l2.PushNode(1)
	l2 = l2.PushNode(2)
	l2 = l2.PushNode(3)

	l3 := common.NewList()
	l3 = l3.PushNode(1)
	l3 = l3.PushNode(2)
	l3 = l3.PushNode(3)
	l3 = l3.PushNode(4)
	l3 = l3.PushNode(5)

	l4 := common.NewList()
	l4 = l4.PushNode(1)
	l4 = l4.PushNode(2)
	l4 = l4.PushNode(3)
	l4 = l4.PushNode(4)
	l4 = l4.PushNode(5)
	l4 = l4.PushNode(6)

	tests := []struct {
		name string
		args args
		want *common.ListNode
	}{
		{
			name: "case1",
			args: args{
				l: l1,
			},
			want: nil,
		},
		{
			name: "case2",
			args: args{
				l: l2,
			},
			want: nil,
		},
		{
			name: "case3",
			args: args{
				l: l3,
			},
			want: nil,
		},
		{
			name: "case4",
			args: args{
				l: l4,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CrossSwap(tt.args.l)
			got.Foreach(common.NodePrint)
		})
	}
}

func TestCrossSwap2(t *testing.T) {
	type args struct {
		l *common.ListNode
	}
	l1 := common.NewList()
	l1 = l1.PushNode(1)
	l1 = l1.PushNode(2)

	l2 := common.NewList()
	l2 = l2.PushNode(1)
	l2 = l2.PushNode(2)
	l2 = l2.PushNode(3)

	l3 := common.NewList()
	l3 = l3.PushNode(1)
	l3 = l3.PushNode(2)
	l3 = l3.PushNode(3)
	l3 = l3.PushNode(4)
	l3 = l3.PushNode(5)

	l4 := common.NewList()
	l4 = l4.PushNode(1)
	l4 = l4.PushNode(2)
	l4 = l4.PushNode(3)
	l4 = l4.PushNode(4)
	l4 = l4.PushNode(5)
	l4 = l4.PushNode(6)

	tests := []struct {
		name string
		args args
		want *common.ListNode
	}{
		{
			name: "case1",
			args: args{
				l: l1,
			},
			want: nil,
		},
		{
			name: "case2",
			args: args{
				l: l2,
			},
			want: nil,
		},
		{
			name: "case3",
			args: args{
				l: l3,
			},
			want: nil,
		},
		{
			name: "case4",
			args: args{
				l: l4,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CrossSwapSmart(tt.args.l)
			got.Foreach(common.NodePrint)
		})
	}
}