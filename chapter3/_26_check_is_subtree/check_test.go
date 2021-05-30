package _26_check_is_subtree

import (
    "github.com/hq-cml/sward2offer/common"
    "testing"
)

func TestCheck(t *testing.T) {
    type args struct {
        root1 *common.TreeNode
        root2 *common.TreeNode
    }
    tree1 := common.NewNodeWithChild(8,
        common.NewNode(9),
        common.NewNode(2),
    )

    tree2 := common.NewNodeWithChild(8,
        common.NewNode(9),
        common.NewNode(2),
    )

    tree3 := common.NewNodeWithChild(8,
        common.NewNode(9),
        common.NewNode(3),
    )

    tree4 := common.NewNodeWithChild(8,
        common.NewNodeWithChild(8,
            common.NewNode(9),
            common.NewNodeWithChild(2,
                common.NewNode(4),
                common.NewNode(7),
            ),
        ),
        common.NewNode(7),
    )

    tests := []struct {
        name string
        args args
        want bool
    }{
        {
            name: "case1",
            args: args{
                root1: tree1,
                root2: tree2,
            },
            want: true,
        },
        {
            name: "case2",
            args: args{
                root1: tree1,
                root2: tree3,
            },
            want: false,
        },
        {
            name: "case3",
            args: args{
                root1: tree4,
                root2: tree1,
            },
            want: true,
        },
        {
            name: "case4",
            args: args{
                root1: tree4,
                root2: tree3,
            },
            want: false,
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := Check(tt.args.root1, tt.args.root2); got != tt.want {
                t.Errorf("Check() = %v, want %v", got, tt.want)
            }
        })
    }
}