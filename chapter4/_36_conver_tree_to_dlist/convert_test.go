package _36_conver_tree_to_dlist

import (
    "fmt"
    "github.com/hq-cml/sward2offer/common"
    "testing"
)

func TestConvert(t *testing.T) {
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
                root: common.NewNodeWithChild(10,
                common.NewNodeWithChild(6,
                    common.NewNode(4),
                    common.NewNode(8)),
                common.NewNodeWithChild(14,
                    common.NewNode(12),
                    common.NewNode(16))),
            },
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            common.Mid(tt.args.root)
            fmt.Println("\n######################")
            got := Convert(tt.args.root)
            print(got)
        })
    }
}

func print(node *common.TreeNode) {
    var tail *common.TreeNode
    for node != nil {
        fmt.Print(node.Val, " ")
        tail = node
        node = node.Right
    }
    fmt.Println("\n------------------------")
    for tail != nil {
        fmt.Print(tail.Val, " ")
        tail = tail.Left
    }
    fmt.Println()
}