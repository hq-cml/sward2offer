package _brother_2_child

import (
    "fmt"
    "github.com/hq-cml/sward2offer/common"
    "testing"
)

func TestChange2ChildRecurse(t *testing.T) {
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
                        common.NewNode(5),
                    ),
                    common.NewNodeWithChild(3,
                        common.NewNode(6),
                        common.NewNode(7),
                    ),
                ),
            },
        },
    }
    for _, tt := range tests {
        common.Pre(tt.args.root)
        fmt.Println()
        common.Mid(tt.args.root)
        fmt.Println()
        common.Post(tt.args.root)
        fmt.Println()

        fmt.Println()

        Change2ChildNoRecursion(tt.args.root)
        common.Pre(tt.args.root)
        fmt.Println()
        common.Mid(tt.args.root)
        fmt.Println()
        common.Post(tt.args.root)
        fmt.Println()
    }
}