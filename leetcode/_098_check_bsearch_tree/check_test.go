package _098_check_bsearch_tree

import (
	"github.com/hq-cml/sward2offer/common"
	"testing"
)

func Test_isValidBST(t *testing.T) {
	type args struct {
		root *common.TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case1",
			args: args{
				root: common.NewNodeWithChild(5,
					common.NewNodeWithChild(4,
						nil,
						nil),
					common.NewNodeWithChild(6,
						common.NewNode(3),
						common.NewNode(7))),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsValidBST(tt.args.root); got != tt.want {
				t.Errorf("isValidBST() = %v, want %v", got, tt.want)
			}
		})
	}
}
