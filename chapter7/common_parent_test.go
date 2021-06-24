package chapter7

import (
    "github.com/hq-cml/sward2offer/common"
    "testing"
)

func TestFindCommonParent(t *testing.T) {
    type args struct {
        root *common.TreeNode
        num1 int
        num2 int
    }
    tests := []struct {
        name  string
        args  args
        want  int
        want1 bool
    }{
        {
            name:  "case1",
            args:  args{
                root: common.NewNodeWithChild(1,
                    common.NewNodeWithChild(2,
                        common.NewNode(4),
                        common.NewNode(5)),
                    common.NewNodeWithChild(3,
                        nil,
                        common.NewNode(6))),
                num1: 5,
                num2: 6,
            },
            want:  1,
            want1: true,
        },
        {
            name:  "case2",
            args:  args{
                root: common.NewNodeWithChild(1,
                    common.NewNodeWithChild(2,
                        common.NewNode(4),
                        common.NewNode(5)),
                    common.NewNodeWithChild(3,
                        nil,
                        common.NewNode(6))),
                num1: 4,
                num2: 5,
            },
            want:  2,
            want1: true,
        },
        {
            name:  "case3",
            args:  args{
                root: common.NewNodeWithChild(1,
                    common.NewNodeWithChild(2,
                        common.NewNode(4),
                        common.NewNode(5)),
                    common.NewNodeWithChild(3,
                        nil,
                        common.NewNode(6))),
                num1: 4,
                num2: 8,
            },
            want:  0,
            want1: false,
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got, got1 := FindCommonParent(tt.args.root, tt.args.num1, tt.args.num2)
            if got != tt.want {
                t.Errorf("FindCommonParent() got = %v, want %v", got, tt.want)
            }
            if got1 != tt.want1 {
                t.Errorf("FindCommonParent() got1 = %v, want %v", got1, tt.want1)
            }
        })
    }
}