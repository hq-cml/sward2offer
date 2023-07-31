package _199_btree_right_side

import (
	"github.com/hq-cml/sward2offer/common"
	"testing"
)

func TestRightSide(t *testing.T) {
	type args struct {
		root *common.TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
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
						common.NewNode(4))),
			},
			want: 134,
		},
		{
			name: "case2",
			args: args{
				root: common.NewNodeWithChild(1,
					common.NewNodeWithChild(2,
						common.NewNode(4),
						common.NewNode(5)),
					common.NewNodeWithChild(3,
						nil,
						nil)),
			},
			want: 135,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RightSide(tt.args.root); got != tt.want {
				t.Errorf("RightSide() = %v, want %v", got, tt.want)
			}
		})
	}
}