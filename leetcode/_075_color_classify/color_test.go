package _075_color_classify

import (
	"fmt"
	"testing"
)

func TestClassify(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "case1",
			args: args{
				[]int{2, 1, 2, 0, 2, 1, 0},
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Classify(tt.args.arr)
			fmt.Println(got)
		})
	}
}

func Test_classify(t *testing.T) {
	type args struct {
		arr      []int
		startIdx int
		endIdx   int
		delima   int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case1",
			args: args{
				arr:      []int{2, 1, 2, 0, 2, 1, 0},
				startIdx: 0,
				endIdx:   6,
				delima:   0,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := classify(tt.args.arr, tt.args.startIdx, tt.args.endIdx, tt.args.delima)
			fmt.Println(tt.args.arr)
			fmt.Println(got)
		})
	}
}

func TestClassify2(t *testing.T) {
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
				arr: []int{2, 1, 2, 0, 2, 1, 0},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Classify2(tt.args.arr)
			fmt.Println(tt.args.arr)
		})
	}
}
