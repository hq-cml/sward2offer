package _07_build_tree

import (
	"fmt"
	"github.com/hq-cml/sward2offer/common"
	"testing"
)

func TestBuildTree(t *testing.T) {
	type args struct {
		pre []int
		mid []int
	}
	tests := []struct {
		name    string
		args    args
		want    *common.Node
		wantErr bool
	}{
		{
			name:    "case1",
			args:    args{
				pre: []int{1, 2, 4, 7, 3, 5, 6, 8},
				mid: []int{4, 7, 2, 1, 5, 3, 8, 6},
			},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := BuildTree(tt.args.pre, tt.args.mid)
			if (err != nil) != tt.wantErr {
				t.Errorf("BuildTree() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			common.Pre(got); fmt.Println()
			common.Mid(got); fmt.Println()
			common.Post(got); fmt.Println()
			//if !reflect.DeepEqual(got, tt.want) {
			//	t.Errorf("BuildTree() got = %v, want %v", got, tt.want)
			//}
		})
	}
}