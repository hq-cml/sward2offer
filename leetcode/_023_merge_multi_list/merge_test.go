package _023_merge_multi_list

import (
	"github.com/hq-cml/sward2offer/common"
	"testing"
)

func TestMergeMulti(t *testing.T) {
	type args struct {
		arr []*common.ListNode
	}
	var l1 *common.ListNode
	l1 = l1.PushNode(1)
	l1 = l1.PushNode(4)
	l1 = l1.PushNode(5)

	var l2 *common.ListNode
	l2 = l2.PushNode(1)
	l2 = l2.PushNode(3)
	l2 = l2.PushNode(4)

	var l3 *common.ListNode
	l3 = l3.PushNode(2)
	l3 = l3.PushNode(6)

	tests := []struct {
		name string
		args args
		want *common.ListNode
	}{
		{
			name: "case1",
			args: args{
				arr: []*common.ListNode{
					l1, l2, l3,
				},
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MergeMulti(tt.args.arr)
			got.Foreach(common.NodePrint)
		})
	}
}
