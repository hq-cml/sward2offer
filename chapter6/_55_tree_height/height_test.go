package _55_tree_height

import (
    "github.com/hq-cml/sward2offer/common"
    "testing"
)

func TestTreeHeight(t *testing.T) {
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
                root: common.NewNodeWithChild(5,
                    common.NewNodeWithChild(3, common.NewNode(2), common.NewNode(4)),
                    common.NewNodeWithChild(7, common.NewNode(6), common.NewNode(8)),
                ),
            },
            want:3,
        },
        {
            name: "case2",
            args: args{
                root: common.NewNodeWithChild(5,
                    common.NewNodeWithChild(3, common.NewNode(2), common.NewNode(4)),
                    common.NewNodeWithChild(7, common.NewNode(6),
                        common.NewNodeWithChild(8, common.NewNode(8), nil)),
                ),
            },
            want:4,
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := TreeHeight(tt.args.root); got != tt.want {
                t.Errorf("TreeHeight() = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestCheckBalance(t *testing.T) {
    type args struct {
        root *common.TreeNode
    }
    tests := []struct {
        name string
        args args
        want bool
    }{
        {
            name: "case1",
            args: args{
                root: common.NewNodeWithChild(5,
                    common.NewNodeWithChild(3, common.NewNode(2), common.NewNode(4)),
                    common.NewNodeWithChild(7, common.NewNode(6), common.NewNode(8)),
                ),
            },
            want:true,
        },
        {
            name: "case2",
            args: args{
                root: common.NewNodeWithChild(5,
                    common.NewNodeWithChild(3, common.NewNode(2), common.NewNode(4)),
                    common.NewNodeWithChild(7, common.NewNode(6),
                        common.NewNodeWithChild(8, common.NewNode(8), nil)),
                ),
            },
            want:true,
        },
        {
            name: "case3",
            args: args{
                root: common.NewNodeWithChild(5,
                    common.NewNodeWithChild(3, common.NewNode(2), common.NewNode(4)),
                    common.NewNodeWithChild(7, common.NewNode(6),
                        common.NewNodeWithChild(8, common.NewNodeWithChild(8,
                            common.NewNode(9), nil), nil)),
                ),
            },
            want:false,
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := CheckBalance(tt.args.root); got != tt.want {
                t.Errorf("CheckBalance() = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestCheckBalace2(t *testing.T) {
    type args struct {
        root *common.TreeNode
    }
    tests := []struct {
        name string
        args args
        want bool
    }{
        {
            name: "case1",
            args: args{
                root: common.NewNodeWithChild(5,
                    common.NewNodeWithChild(3, common.NewNode(2), common.NewNode(4)),
                    common.NewNodeWithChild(7, common.NewNode(6), common.NewNode(8)),
                ),
            },
            want:true,
        },
        {
            name: "case2",
            args: args{
                root: common.NewNodeWithChild(5,
                    common.NewNodeWithChild(3, common.NewNode(2), common.NewNode(4)),
                    common.NewNodeWithChild(7, common.NewNode(6),
                        common.NewNodeWithChild(8, common.NewNode(8), nil)),
                ),
            },
            want:true,
        },
        {
            name: "case3",
            args: args{
                root: common.NewNodeWithChild(5,
                    common.NewNodeWithChild(3, common.NewNode(2), common.NewNode(4)),
                    common.NewNodeWithChild(7, common.NewNode(6),
                        common.NewNodeWithChild(8, common.NewNodeWithChild(8,
                            common.NewNode(9), nil), nil)),
                ),
            },
            want:false,
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := CheckBalace2(tt.args.root); got != tt.want {
                t.Errorf("CheckBalace2() = %v, want %v", got, tt.want)
            }
        })
    }
}