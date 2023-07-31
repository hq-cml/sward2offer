package _good_node_nums

import (
	"github.com/hq-cml/sward2offer/common"
	"testing"
)

func TestDfs(t *testing.T) {
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
				root: common.NewNodeWithChild(3,
					common.NewNodeWithChild(1,
						common.NewNode(3),
						nil),
					common.NewNodeWithChild(4,common.NewNode(1),
						common.NewNode(5))),
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Good(tt.args.root); got != tt.want {
				t.Errorf("Good() = %v, want %v", got, tt.want)
			}
		})
	}
}