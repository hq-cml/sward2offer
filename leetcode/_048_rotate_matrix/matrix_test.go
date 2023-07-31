package _048_rotate_matrix

import (
	"fmt"
	"testing"
)

func TestMatrix(t *testing.T) {
	type args struct {
		matrix [][]int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "case0",
			args: args{
				[][]int{
					{1, 2},
					{3, 4},
				},
			},
		},
		{
			name: "case1",
			args: args{
				[][]int{
					{1, 2, 3},
					{4, 5, 6},
					{7, 8, 9},
				},
			},
		},
		{
			name: "case2",
			args: args{
				[][]int{
					{5, 1, 9, 11},
					{2, 4, 8, 10},
					{13, 3, 6, 7},
					{15, 14, 12, 16},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Matrix(tt.args.matrix)
			for _, l := range tt.args.matrix {
				fmt.Println(l)
			}
		})
	}
}
