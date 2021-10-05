package _15_3sum

import (
	"reflect"
	"testing"
)

func TestThreeSum(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "case1",
			args: args{
				nums: []int{},
			},
			want: [][]int{},
		},
		{
			name: "case2",
			args: args{
				nums: []int{0},
			},
			want: [][]int{},
		},
		{
			name: "case3",
			args: args{
				nums: []int{-1, 0, 1, 2, -1, -4},
			},
			want: [][]int{
				{-1, -1, 2},
				{-1, 0, 1},
			},
		},
		{
			name: "case4",
			args: args{
				nums: []int{-1, 0, 1, 2, -1, -4, -1, -1, -1},
			},
			want: [][]int{
				{-1, -1, 2},
				{-1, 0, 1},
			},
		},
		{
			name: "case5",
			args: args{
				nums: []int{-1, 0, 1, 1, 1, 2, -1, 2, 2, 2},
			},
			want: [][]int{
				{-1, -1, 2},
				{-1, 0, 1},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ThreeSum(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ThreeSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
