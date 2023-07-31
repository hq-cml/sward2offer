package _039_pick_sum

import (
	"fmt"
	"testing"
)

func TestFind(t *testing.T) {
	type args struct {
		arr    []int
		target int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "case1",
			args: args{
				arr:    []int{7, 2, 3, 6},
				target: 7,
			},
			want: nil,
		}, {
			name: "case1",
			args: args{
				arr:    []int{2, 5, 3},
				target: 8,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Find(tt.args.arr, tt.args.target)
			fmt.Println(got)
		})
	}
}
