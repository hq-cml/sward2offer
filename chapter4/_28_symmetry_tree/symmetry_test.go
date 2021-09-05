package _28_symmetry_tree

import (
	"github.com/hq-cml/sward2offer/common"
	"testing"
)

func TestSymmetry(t *testing.T) {
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
				root: common.NewNodeWithChild(1,
					common.NewNodeWithChild(2,
						common.NewNode(4),
						common.NewNode(5)),
					common.NewNodeWithChild(3,
						nil,
						common.NewNode(6))),
			},
			want: false,
		},
		{
			name: "case2",
			args: args{
				root: common.NewNodeWithChild(1,
					common.NewNodeWithChild(2,
						common.NewNode(4),
						common.NewNode(5)),
					common.NewNodeWithChild(2,
						common.NewNode(5),
						common.NewNode(4))),
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Symmetry(tt.args.root); got != tt.want {
				t.Errorf("Symmetry() = %v, want %v", got, tt.want)
			}
		})
	}
}
