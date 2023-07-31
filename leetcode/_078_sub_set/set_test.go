package _078_sub_set

import (
	"fmt"
	"testing"
)

func Test_pickSet(t *testing.T) {
	type args struct {
		arr []int
		n   int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "case1",
			args: args{
				arr: []int{1, 2, 3},
				n:   2,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := pickSet(tt.args.arr, tt.args.n)
			fmt.Println(got)
		})
	}
}

func TestPickSet(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "case1",
			args: args{
				arr: []int{1, 2, 3},
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := PickSet(tt.args.arr)
			fmt.Println(got)
		})
	}
}
