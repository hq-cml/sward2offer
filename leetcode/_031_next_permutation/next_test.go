package _031_next_permutation

import (
	"fmt"
	"testing"
)

func TestNextPermutation(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "case1",
			args: args{
				arr: []int{1, 2, 3},
			},
		},
		{
			name: "case2",
			args: args{
				arr: []int{3, 2, 1},
			},
		},
		{
			name: "case3",
			args: args{
				arr: []int{1, 3, 2},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}

func TestNextPermutation1(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "case1",
			args: args{
				arr: []int{1, 2, 3},
			},
		},
		{
			name: "case2",
			args: args{
				arr: []int{3, 2, 1},
			},
		},
		{
			name: "case3",
			args: args{
				arr: []int{1, 3, 2},
			},
		},
		{
			name: "case4",
			args: args{
				arr: []int{2, 1, 4, 3},
			},
		},
		{
			name: "case5",
			args: args{
				arr: []int{4, 2, 5, 3, 1},
			},
		},
	}
	for _, tt := range tests {
		NextPermutation(tt.args.arr)
		fmt.Println(tt.args.arr)
	}
}
