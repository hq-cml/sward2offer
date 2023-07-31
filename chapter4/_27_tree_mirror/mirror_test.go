package _27_tree_mirror

import (
	"fmt"
	"github.com/hq-cml/sward2offer/common"
	"testing"
)

func TestMirror(t *testing.T) {
	type args struct {
		root *common.TreeNode
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "case1",
			args: args{
				root: common.NewNodeWithChild(1,
					common.NewNodeWithChild(2,
						common.NewNode(4),
						common.NewNode(5)),
					common.NewNodeWithChild(3,
						nil,
						common.NewNode(6))),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			common.Pre(tt.args.root)
			fmt.Println()
			common.Mid(tt.args.root)
			fmt.Println()

			tt.args.root = Mirror(tt.args.root)

			common.Pre(tt.args.root)
			fmt.Println()
			common.Mid(tt.args.root)
			fmt.Println()
		})
	}
}