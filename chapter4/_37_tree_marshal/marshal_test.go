package _37_tree_marshal

import (
    "fmt"
    "github.com/hq-cml/sward2offer/common"
    "testing"
)

func TestTreeMarshal(t *testing.T) {
    type args struct {
        root *common.TreeNode
    }
    tests := []struct {
        name string
        args args
        want string
    }{
        {
           name: "case1",
           args: args{
               root: common.NewNodeWithChild(10,
                   common.NewNodeWithChild(6,
                       common.NewNode(4),
                       common.NewNode(8)),
                   common.NewNodeWithChild(14,
                       common.NewNode(12),
                       common.NewNode(16))),
           },
        },
        {
            name: "case2",
            args: args{
                root: common.NewNodeWithChild(1,
                    common.NewNode(2),
                    common.NewNode(3)),
            },
        },
        {
            name: "case3",
            args: args{
                root: common.NewNodeWithChild(1,
                    nil,
                    nil),
            },
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := TreeMarshal(tt.args.root)
            fmt.Println(got)
            root := TreeUnMarshal(got)
            common.Pre(root)
            fmt.Println()
        })
    }
}