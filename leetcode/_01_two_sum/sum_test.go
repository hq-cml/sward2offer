package _01_two_sum

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	type args struct {
		nums []int
		sum  int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "case1",
			args: args{
				nums: []int{2, 7, 11, 15},
				sum:  9,
			},
			want: []int{0, 1},
		},
		{
			name: "case2",
			args: args{
				nums: []int{3, 2, 4},
				sum:  6,
			},
			want: []int{1, 2},
		},
		{
			name: "case3",
			args: args{
				nums: []int{3, 3},
				sum:  6,
			},
			want: []int{0, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Sum(tt.args.nums, tt.args.sum); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Sum() = %v, want %v", got, tt.want)
			}
		})
	}
}
