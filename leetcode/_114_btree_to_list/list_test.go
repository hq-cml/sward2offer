package _114_btree_to_list

import (
	"fmt"
	"github.com/hq-cml/sward2offer/common"
	"testing"
)

func TestTransfer(t *testing.T) {
	type args struct {
		root *common.TreeNode
	}
	tests := []struct {
		name string
		args args
		want *common.TreeNode
	}{
		{
			name: "case1",
			args: args{
				root: common.NewNodeWithChild(1,
					common.NewNodeWithChild(2, nil, nil),
					common.NewNodeWithChild(3, nil, nil)),
			},
			want: nil,
		},
		{
			name: "case2",
			args: args{
				root: common.NewNodeWithChild(3,
					common.NewNodeWithChild(1,
						common.NewNode(3),
						nil),
					common.NewNodeWithChild(4, common.NewNode(1),
						common.NewNode(5))),
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Transfer(tt.args.root)
			for got != nil {
				fmt.Println(got.Val)
				got = got.Right
			}
		})
	}
}
