package _34_find_path_num

import (
    "github.com/hq-cml/sward2offer/common"
    "testing"
)

func TestFindPath(t *testing.T) {
    type args struct {
        root *common.TreeNode
        num  int
    }
    tests := []struct {
        name string
        args args
    }{
        {
            name: "case1",
            args: args{
                root: common.NewNodeWithChild(10,
                    common.NewNodeWithChild(5,
                        common.NewNode(4),
                        common.NewNode(7)),
                    common.NewNode(12)),
                num:  22,
            },
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            FindPath(tt.args.root, tt.args.num)
        })
    }
}