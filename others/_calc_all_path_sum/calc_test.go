package _calc_all_path_sum

import (
	"github.com/hq-cml/sward2offer/common"
	"testing"
)

func TestDfsCalc(t *testing.T) {
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
					common.NewNode(3),
					common.NewNode(2)),
			},
			want: 25,
		},

		{
			name: "case2",
			args: args{
				root: common.NewNodeWithChild(3,
					common.NewNodeWithChild(1,
						common.NewNode(3),
						nil),
					common.NewNodeWithChild(4,common.NewNode(1),
						common.NewNode(5))),
			},
			want: 999,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DfsCalc(tt.args.root); got != tt.want {
				t.Errorf("DfsCalc() = %v, want %v", got, tt.want)
			}
		})
	}
}