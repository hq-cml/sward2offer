package _54_k_th_Node

import (
	"fmt"
	"github.com/hq-cml/sward2offer/common"
	"testing"
)

func TestKthNode(t *testing.T) {
	type args struct {
		root *common.TreeNode
		k    int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "case1",
			args: args{
				root: common.NewNodeWithChild(5,
					common.NewNodeWithChild(3, common.NewNode(2), common.NewNode(4)),
					common.NewNodeWithChild(7, common.NewNode(6), common.NewNode(8)),
				),
				k: 3,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := KthNode(tt.args.root, tt.args.k)
			fmt.Println(got.Val)
		})
	}
}
