package _1448_good_node

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
			if got := Dfs(tt.args.root); got != tt.want {
				t.Errorf("Dfs() = %v, want %v", got, tt.want)
			}
		})
	}
}