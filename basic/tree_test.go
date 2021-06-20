package basic

import (
    "github.com/hq-cml/sward2offer/common"
    "reflect"
    "testing"
)

func TestFindPath(t *testing.T) {
    type args struct {
        root *common.TreeNode
        need int
        path []int
    }
    tests := []struct {
        name  string
        args  args
        want  []int
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
                need: 5,
                path: []int{},
            },
            want:  []int{1,2,5},
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
                need: 6,
                path: []int{},
            },
            want:  []int{1,3,6},
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
                need: 8,
                path: []int{},
            },
            want:  nil,
            want1: false,
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got, got1 := FindPath(tt.args.root, tt.args.need, tt.args.path)
            if !reflect.DeepEqual(got, tt.want) {
                t.Errorf("FindPath() got = %v, want %v", got, tt.want)
            }
            if got1 != tt.want1 {
                t.Errorf("FindPath() got1 = %v, want %v", got1, tt.want1)
            }
        })
    }
}