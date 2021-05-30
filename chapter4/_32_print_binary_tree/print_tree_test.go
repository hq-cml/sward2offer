package _32_print_binary_tree

import (
    "github.com/hq-cml/sward2offer/common"
    "testing"
)

func TestPrintTreeIn1Line(t *testing.T) {
    type args struct {
        root *common.TreeNode
    }
    tests := []struct {
        name    string
        args    args
        wantErr bool
    }{
        {
            name:    "case1",
            args:    args{
                common.NewNodeWithChild(1,
                    common.NewNodeWithChild(2,
                        common.NewNode(4),
                        common.NewNode(5)),
                    common.NewNodeWithChild(3,
                        nil,
                        common.NewNode(6))),
            },
            wantErr: false,
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if err := PrintTreeIn1Line(tt.args.root); (err != nil) != tt.wantErr {
                t.Errorf("PrintTreeIn1Line() error = %v, wantErr %v", err, tt.wantErr)
            }
        })
    }
}

func TestPrintTreeInMultiLine(t *testing.T) {
    type args struct {
        root *common.TreeNode
    }
    tests := []struct {
        name    string
        args    args
        wantErr bool
    }{
        {
            name:    "case1",
            args:    args{
                common.NewNodeWithChild(1,
                    common.NewNodeWithChild(2,
                        common.NewNode(4),
                        common.NewNode(5)),
                    common.NewNodeWithChild(3,
                        nil,
                        common.NewNode(6))),
            },
            wantErr: false,
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if err := PrintTreeInMultiLine(tt.args.root); (err != nil) != tt.wantErr {
                t.Errorf("PrintTreeInMultiLine() error = %v, wantErr %v", err, tt.wantErr)
            }
        })
    }
}