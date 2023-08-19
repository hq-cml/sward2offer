package _traverse_btree_norecurse

import (
	"github.com/hq-cml/sward2offer/common"
	"testing"
)

func TestPre(t *testing.T) {
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
			PreTree(tt.args.root)
		})
	}
}

func TestMid(t *testing.T) {
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
			MidTree(tt.args.root)
		})
	}
}

func TestPost(t *testing.T) {
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
			PostTree(tt.args.root)
		})
	}
}
